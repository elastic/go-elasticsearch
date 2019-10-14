// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package xkcdsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// SearchResults wraps the Elasticsearch search response.
//
type SearchResults struct {
	Total int    `json:"total"`
	Hits  []*Hit `json:"hits"`
}

// Hit wraps the document returned in search response.
//
type Hit struct {
	Document
	URL        string        `json:"url"`
	Sort       []interface{} `json:"sort"`
	Highlights *struct {
		Title      []string `json:"title"`
		Alt        []string `json:"alt"`
		Transcript []string `json:"transcript"`
	} `json:"highlights,omitempty"`
}

// StoreConfig configures the store.
//
type StoreConfig struct {
	Client    *elasticsearch.Client
	IndexName string
}

// Store allows to index and search documents.
//
type Store struct {
	es        *elasticsearch.Client
	indexName string
}

// NewStore returns a new instance of the store.
//
func NewStore(c StoreConfig) (*Store, error) {
	indexName := c.IndexName
	if indexName == "" {
		indexName = "xkcd"
	}

	s := Store{es: c.Client, indexName: indexName}
	return &s, nil
}

// CreateIndex creates a new index with mapping.
//
func (s *Store) CreateIndex(mapping string) error {
	res, err := s.es.Indices.Create(s.indexName, s.es.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("error: %s", res)
	}
	return nil
}

// Create indexes a new document into store.
//
func (s *Store) Create(item *Document) error {
	payload, err := json.Marshal(item)
	if err != nil {
		return err
	}

	ctx := context.Background()
	res, err := esapi.CreateRequest{
		Index:      s.indexName,
		DocumentID: item.ID,
		Body:       bytes.NewReader(payload),
	}.Do(ctx, s.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}
		return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}

	return nil
}

// Exists returns true when a document with id already exists in the store.
//
func (s *Store) Exists(id string) (bool, error) {
	res, err := s.es.Exists(s.indexName, id)
	if err != nil {
		return false, err
	}
	switch res.StatusCode {
	case 200:
		return true, nil
	case 404:
		return false, nil
	default:
		return false, fmt.Errorf("[%s]", res.Status())
	}
}

// Search returns results matching a query, paginated by after.
//
func (s *Store) Search(query string, after ...string) (*SearchResults, error) {
	var results SearchResults

	res, err := s.es.Search(
		s.es.Search.WithIndex(s.indexName),
		s.es.Search.WithBody(s.buildQuery(query, after...)),
	)
	if err != nil {
		return &results, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return &results, err
		}
		return &results, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}

	type envelopeResponse struct {
		Took int
		Hits struct {
			Total struct {
				Value int
			}
			Hits []struct {
				ID         string          `json:"_id"`
				Source     json.RawMessage `json:"_source"`
				Highlights json.RawMessage `json:"highlight"`
				Sort       []interface{}   `json:"sort"`
			}
		}
	}

	var r envelopeResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return &results, err
	}

	results.Total = r.Hits.Total.Value

	if len(r.Hits.Hits) < 1 {
		results.Hits = []*Hit{}
		return &results, nil
	}

	for _, hit := range r.Hits.Hits {
		var h Hit
		h.ID = hit.ID
		h.Sort = hit.Sort
		h.URL = strings.Join([]string{baseURL, h.ID, ""}, "/")

		if err := json.Unmarshal(hit.Source, &h); err != nil {
			return &results, err
		}

		if len(hit.Highlights) > 0 {
			if err := json.Unmarshal(hit.Highlights, &h.Highlights); err != nil {
				return &results, err
			}
		}

		results.Hits = append(results.Hits, &h)
	}

	return &results, nil
}

func (s *Store) buildQuery(query string, after ...string) io.Reader {
	var b strings.Builder

	b.WriteString("{\n")

	if query == "" {
		b.WriteString(searchAll)
	} else {
		b.WriteString(fmt.Sprintf(searchMatch, query))
	}

	if len(after) > 0 && after[0] != "" && after[0] != "null" {
		b.WriteString(",\n")
		b.WriteString(fmt.Sprintf(`	"search_after": %s`, after))
	}

	b.WriteString("\n}")

	// fmt.Printf("%s\n", b.String())
	return strings.NewReader(b.String())
}

const searchAll = `
	"query" : { "match_all" : {} },
	"size" : 25,
	"sort" : { "published" : "desc", "_doc" : "asc" }`

const searchMatch = `
	"query" : {
		"multi_match" : {
			"query" : %q,
			"fields" : ["title^100", "alt^10", "transcript"],
			"operator" : "and"
		}
	},
	"highlight" : {
		"fields" : {
			"title" : { "number_of_fragments" : 0 },
			"alt" : { "number_of_fragments" : 0 },
			"transcript" : { "number_of_fragments" : 5, "fragment_size" : 25 }
		}
	},
	"size" : 25,
	"sort" : [ { "_score" : "desc" }, { "_doc" : "asc" } ]`

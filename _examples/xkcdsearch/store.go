// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package xkcdsearch

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

// SearchResults wraps the Elasticsearch search response.
type SearchResults struct {
	Total int    `json:"total"`
	Hits  []*Hit `json:"hits"`
}

// Hit wraps the document returned in search response.
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
type StoreConfig struct {
	Client    *elasticsearch.TypedClient
	IndexName string
}

// Store allows to index and search documents.
type Store struct {
	es        *elasticsearch.TypedClient
	indexName string
}

// NewStore returns a new instance of the store.
func NewStore(c StoreConfig) (*Store, error) {
	indexName := c.IndexName
	if indexName == "" {
		indexName = "xkcd"
	}

	s := Store{es: c.Client, indexName: indexName}
	return &s, nil
}

// CreateIndex creates a new index with the provided JSON mapping body.
func (s *Store) CreateIndex(mapping string) error {
	_, err := s.es.Indices.Create(s.indexName).
		Raw(strings.NewReader(mapping)).
		Do(context.Background())
	return err
}

// Create indexes a new document into the store.
func (s *Store) Create(item *Document) error {
	_, err := s.es.Create(s.indexName, item.ID).
		Document(item).
		Do(context.Background())
	return err
}

// Exists returns true when a document with id already exists in the store.
func (s *Store) Exists(id string) (bool, error) {
	return s.es.Exists(s.indexName, id).Do(context.Background())
}

// Search returns results matching a query, paginated by after.
//
// `after` is a comma-separated list of sort values (as written by a previous
// response's `sort` field), e.g. "1,2". Each token is parsed as int64, then
// float64, falling back to string.
func (s *Store) Search(query string, after ...string) (*SearchResults, error) {
	var results SearchResults

	req := s.es.Search().Index(s.indexName).Size(25)

	if query == "" {
		req = req.
			Query(esdsl.NewMatchAllQuery()).
			Sort(
				esdsl.NewSortOptions().AddSortOption("published", esdsl.NewFieldSort(sortorder.Desc)),
				esdsl.NewSortOptions().AddSortOption("_doc", esdsl.NewFieldSort(sortorder.Asc)),
			)
	} else {
		req = req.
			Query(esdsl.NewMultiMatchQuery(query).
				Fields("title^100", "alt^10", "transcript").
				Operator(operator.And)).
			Highlight(esdsl.NewHighlight().Fields([]map[string]types.HighlightField{{
				"title":      *esdsl.NewHighlightField().NumberOfFragments(0).HighlightFieldCaster(),
				"alt":        *esdsl.NewHighlightField().NumberOfFragments(0).HighlightFieldCaster(),
				"transcript": *esdsl.NewHighlightField().NumberOfFragments(5).FragmentSize(25).HighlightFieldCaster(),
			}})).
			Sort(
				esdsl.NewSortOptions().Score_(esdsl.NewScoreSort().Order(sortorder.Desc)),
				esdsl.NewSortOptions().AddSortOption("_doc", esdsl.NewFieldSort(sortorder.Asc)),
			)
	}

	if len(after) > 0 && after[0] != "" && after[0] != "null" {
		req = req.SearchAfterValues(parseSearchAfter(after[0]))
	}

	res, err := req.Do(context.Background())
	if err != nil {
		return &results, err
	}

	if res.Hits.Total != nil {
		results.Total = int(res.Hits.Total.Value)
	}

	if len(res.Hits.Hits) < 1 {
		results.Hits = []*Hit{}
		return &results, nil
	}

	for _, hit := range res.Hits.Hits {
		var h Hit
		if hit.Id_ != nil {
			h.ID = *hit.Id_
		}
		h.Sort = make([]interface{}, len(hit.Sort))
		for i, v := range hit.Sort {
			h.Sort[i] = v
		}
		h.URL = baseURL + "/" + h.ID + "/"

		if len(hit.Source_) > 0 {
			if err := json.Unmarshal(hit.Source_, &h); err != nil {
				return &results, err
			}
		}

		if len(hit.Highlight) > 0 {
			h.Highlights = &struct {
				Title      []string `json:"title"`
				Alt        []string `json:"alt"`
				Transcript []string `json:"transcript"`
			}{
				Title:      hit.Highlight["title"],
				Alt:        hit.Highlight["alt"],
				Transcript: hit.Highlight["transcript"],
			}
		}

		results.Hits = append(results.Hits, &h)
	}

	return &results, nil
}

// parseSearchAfter splits a comma-separated token list into typed field
// values, trying int64 then float64 and falling back to string.
func parseSearchAfter(s string) []types.FieldValue {
	parts := strings.Split(s, ",")
	out := make([]types.FieldValue, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if i, err := strconv.ParseInt(p, 10, 64); err == nil {
			out = append(out, i)
			continue
		}
		if f, err := strconv.ParseFloat(p, 64); err == nil {
			out = append(out, f)
			continue
		}
		out = append(out, p)
	}
	return out
}

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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MultisearchBody type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/msearch/types.ts#L70-L204
type MultisearchBody struct {
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`
	Collapse     *FieldCollapse          `json:"collapse,omitempty"`
	// DocvalueFields Array of wildcard (*) patterns. The request returns doc values for field
	// names matching these patterns in the hits.fields property of the response.
	DocvalueFields []FieldAndFormat `json:"docvalue_fields,omitempty"`
	// Explain If true, returns detailed information about score computation as part of a
	// hit.
	Explain *bool `json:"explain,omitempty"`
	// Ext Configuration of search extensions defined by Elasticsearch plugins.
	Ext map[string]json.RawMessage `json:"ext,omitempty"`
	// Fields Array of wildcard (*) patterns. The request returns values for field names
	// matching these patterns in the hits.fields property of the response.
	Fields []FieldAndFormat `json:"fields,omitempty"`
	// From Starting document offset. By default, you cannot page through more than
	// 10,000
	// hits using the from and size parameters. To page through more hits, use the
	// search_after parameter.
	From      *int       `json:"from,omitempty"`
	Highlight *Highlight `json:"highlight,omitempty"`
	// IndicesBoost Boosts the _score of documents from specified indices.
	IndicesBoost []map[string]Float64 `json:"indices_boost,omitempty"`
	// Knn Defines the approximate kNN search to run.
	Knn []KnnSearch `json:"knn,omitempty"`
	// MinScore The minimum `_score` for matching documents.
	// Documents with a lower `_score` are not included in search results and
	// results collected by aggregations.
	MinScore *Float64 `json:"min_score,omitempty"`
	// Pit Limits the search to a point in time (PIT). If you provide a PIT, you
	// cannot specify an <index> in the request path.
	Pit        *PointInTimeReference `json:"pit,omitempty"`
	PostFilter *Query                `json:"post_filter,omitempty"`
	Profile    *bool                 `json:"profile,omitempty"`
	// Query Defines the search definition using the Query DSL.
	Query   *Query    `json:"query,omitempty"`
	Rescore []Rescore `json:"rescore,omitempty"`
	// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
	// precedence over mapped fields with the same name.
	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`
	// ScriptFields Retrieve a script evaluation (based on different fields) for each hit.
	ScriptFields map[string]ScriptField `json:"script_fields,omitempty"`
	SearchAfter  []FieldValue           `json:"search_after,omitempty"`
	// SeqNoPrimaryTerm If true, returns sequence number and primary term of the last modification
	// of each hit. See Optimistic concurrency control.
	SeqNoPrimaryTerm *bool `json:"seq_no_primary_term,omitempty"`
	// Size The number of hits to return. By default, you cannot page through more
	// than 10,000 hits using the from and size parameters. To page through more
	// hits, use the search_after parameter.
	Size *int               `json:"size,omitempty"`
	Sort []SortCombinations `json:"sort,omitempty"`
	// Source_ Indicates which source fields are returned for matching documents. These
	// fields are returned in the hits._source property of the search response.
	Source_ SourceConfig `json:"_source,omitempty"`
	// Stats Stats groups to associate with the search. Each group maintains a statistics
	// aggregation for its associated searches. You can retrieve these stats using
	// the indices stats API.
	Stats []string `json:"stats,omitempty"`
	// StoredFields List of stored fields to return as part of a hit. If no fields are specified,
	// no stored fields are included in the response. If this field is specified,
	// the _source
	// parameter defaults to false. You can pass _source: true to return both source
	// fields
	// and stored fields in the search response.
	StoredFields []string   `json:"stored_fields,omitempty"`
	Suggest      *Suggester `json:"suggest,omitempty"`
	// TerminateAfter Maximum number of documents to collect for each shard. If a query reaches
	// this
	// limit, Elasticsearch terminates the query early. Elasticsearch collects
	// documents
	// before sorting. Defaults to 0, which does not terminate query execution
	// early.
	TerminateAfter *int64 `json:"terminate_after,omitempty"`
	// Timeout Specifies the period of time to wait for a response from each shard. If no
	// response
	// is received before the timeout expires, the request fails and returns an
	// error.
	// Defaults to no timeout.
	Timeout *string `json:"timeout,omitempty"`
	// TrackScores If true, calculate and return document scores, even if the scores are not
	// used for sorting.
	TrackScores *bool `json:"track_scores,omitempty"`
	// TrackTotalHits Number of hits matching the query to count accurately. If true, the exact
	// number of hits is returned at the cost of some performance. If false, the
	// response does not include the total number of hits matching the query.
	// Defaults to 10,000 hits.
	TrackTotalHits TrackHits `json:"track_total_hits,omitempty"`
	// Version If true, returns document version as part of a hit.
	Version *bool `json:"version,omitempty"`
}

func (s *MultisearchBody) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "aggregations", "aggs":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]Aggregations, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return fmt.Errorf("%s | %w", "Aggregations", err)
			}

		case "collapse":
			if err := dec.Decode(&s.Collapse); err != nil {
				return fmt.Errorf("%s | %w", "Collapse", err)
			}

		case "docvalue_fields":
			if err := dec.Decode(&s.DocvalueFields); err != nil {
				return fmt.Errorf("%s | %w", "DocvalueFields", err)
			}

		case "explain":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Explain", err)
				}
				s.Explain = &value
			case bool:
				s.Explain = &v
			}

		case "ext":
			if s.Ext == nil {
				s.Ext = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Ext); err != nil {
				return fmt.Errorf("%s | %w", "Ext", err)
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "from":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "From", err)
				}
				s.From = &value
			case float64:
				f := int(v)
				s.From = &f
			}

		case "highlight":
			if err := dec.Decode(&s.Highlight); err != nil {
				return fmt.Errorf("%s | %w", "Highlight", err)
			}

		case "indices_boost":
			if err := dec.Decode(&s.IndicesBoost); err != nil {
				return fmt.Errorf("%s | %w", "IndicesBoost", err)
			}

		case "knn":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewKnnSearch()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Knn", err)
				}

				s.Knn = append(s.Knn, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Knn); err != nil {
					return fmt.Errorf("%s | %w", "Knn", err)
				}
			}

		case "min_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinScore", err)
				}
				f := Float64(value)
				s.MinScore = &f
			case float64:
				f := Float64(v)
				s.MinScore = &f
			}

		case "pit":
			if err := dec.Decode(&s.Pit); err != nil {
				return fmt.Errorf("%s | %w", "Pit", err)
			}

		case "post_filter":
			if err := dec.Decode(&s.PostFilter); err != nil {
				return fmt.Errorf("%s | %w", "PostFilter", err)
			}

		case "profile":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Profile", err)
				}
				s.Profile = &value
			case bool:
				s.Profile = &v
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		case "rescore":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewRescore()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Rescore", err)
				}

				s.Rescore = append(s.Rescore, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Rescore); err != nil {
					return fmt.Errorf("%s | %w", "Rescore", err)
				}
			}

		case "runtime_mappings":
			if err := dec.Decode(&s.RuntimeMappings); err != nil {
				return fmt.Errorf("%s | %w", "RuntimeMappings", err)
			}

		case "script_fields":
			if s.ScriptFields == nil {
				s.ScriptFields = make(map[string]ScriptField, 0)
			}
			if err := dec.Decode(&s.ScriptFields); err != nil {
				return fmt.Errorf("%s | %w", "ScriptFields", err)
			}

		case "search_after":
			if err := dec.Decode(&s.SearchAfter); err != nil {
				return fmt.Errorf("%s | %w", "SearchAfter", err)
			}

		case "seq_no_primary_term":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SeqNoPrimaryTerm", err)
				}
				s.SeqNoPrimaryTerm = &value
			case bool:
				s.SeqNoPrimaryTerm = &v
			}

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "sort":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(SortCombinations)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}

				s.Sort = append(s.Sort, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Sort); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}
			}

		case "_source":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		source__field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Source_", err)
				}

				switch t {

				case "exclude_vectors", "excludes", "includes":
					o := NewSourceFilter()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Source_", err)
					}
					s.Source_ = o
					break source__field

				}
			}
			if s.Source_ == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Source_); err != nil {
					return fmt.Errorf("%s | %w", "Source_", err)
				}
			}

		case "stats":
			if err := dec.Decode(&s.Stats); err != nil {
				return fmt.Errorf("%s | %w", "Stats", err)
			}

		case "stored_fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "StoredFields", err)
				}

				s.StoredFields = append(s.StoredFields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.StoredFields); err != nil {
					return fmt.Errorf("%s | %w", "StoredFields", err)
				}
			}

		case "suggest":
			if err := dec.Decode(&s.Suggest); err != nil {
				return fmt.Errorf("%s | %w", "Suggest", err)
			}

		case "terminate_after":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TerminateAfter", err)
				}
				s.TerminateAfter = &value
			case float64:
				f := int64(v)
				s.TerminateAfter = &f
			}

		case "timeout":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Timeout", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Timeout = &o

		case "track_scores":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TrackScores", err)
				}
				s.TrackScores = &value
			case bool:
				s.TrackScores = &v
			}

		case "track_total_hits":
			if err := dec.Decode(&s.TrackTotalHits); err != nil {
				return fmt.Errorf("%s | %w", "TrackTotalHits", err)
			}

		case "version":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Version", err)
				}
				s.Version = &value
			case bool:
				s.Version = &v
			}

		}
	}
	return nil
}

// NewMultisearchBody returns a MultisearchBody.
func NewMultisearchBody() *MultisearchBody {
	r := &MultisearchBody{
		Aggregations: make(map[string]Aggregations),
		Ext:          make(map[string]json.RawMessage),
		ScriptFields: make(map[string]ScriptField),
	}

	return r
}

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

// ReindexSource type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/reindex/types.ts#L69-L110
type ReindexSource struct {
	// Index The name of the data stream, index, or alias you are copying from.
	// It accepts a comma-separated list to reindex from multiple sources.
	Index []string `json:"index"`
	// Query The documents to reindex, which is defined with Query DSL.
	Query *Query `json:"query,omitempty"`
	// Remote A remote instance of Elasticsearch that you want to index from.
	Remote          *RemoteSource `json:"remote,omitempty"`
	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`
	// Size The number of documents to index per batch.
	// Use it when you are indexing from remote to ensure that the batches fit
	// within the on-heap buffer, which defaults to a maximum size of 100 MB.
	Size *int `json:"size,omitempty"`
	// Slice Slice the reindex request manually using the provided slice ID and total
	// number of slices.
	Slice *SlicedScroll `json:"slice,omitempty"`
	// Sort A comma-separated list of `<field>:<direction>` pairs to sort by before
	// indexing.
	// Use it in conjunction with `max_docs` to control what documents are
	// reindexed.
	//
	// WARNING: Sort in reindex is deprecated.
	// Sorting in reindex was never guaranteed to index documents in order and
	// prevents further development of reindex such as resilience and performance
	// improvements.
	// If used in combination with `max_docs`, consider using a query filter
	// instead.
	Sort []SortCombinations `json:"sort,omitempty"`
	// SourceFields_ If `true`, reindex all source fields.
	// Set it to a list to reindex select fields.
	SourceFields_ []string `json:"_source,omitempty"`
}

func (s *ReindexSource) UnmarshalJSON(data []byte) error {

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

		case "index":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Index", err)
				}

				s.Index = append(s.Index, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Index); err != nil {
					return fmt.Errorf("%s | %w", "Index", err)
				}
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		case "remote":
			if err := dec.Decode(&s.Remote); err != nil {
				return fmt.Errorf("%s | %w", "Remote", err)
			}

		case "runtime_mappings":
			if err := dec.Decode(&s.RuntimeMappings); err != nil {
				return fmt.Errorf("%s | %w", "RuntimeMappings", err)
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

		case "slice":
			if err := dec.Decode(&s.Slice); err != nil {
				return fmt.Errorf("%s | %w", "Slice", err)
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
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "SourceFields_", err)
				}

				s.SourceFields_ = append(s.SourceFields_, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.SourceFields_); err != nil {
					return fmt.Errorf("%s | %w", "SourceFields_", err)
				}
			}

		}
	}
	return nil
}

// NewReindexSource returns a ReindexSource.
func NewReindexSource() *ReindexSource {
	r := &ReindexSource{}

	return r
}

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
// https://github.com/elastic/elasticsearch-specification/tree/9fcf6a64c550d2e8090c8134867f200b56fd7fc7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// The result of an index action. It is a container that holds either the
// `response` of an executed index operation, or the `request` that would have
// run when the action is simulated.
//
// https://github.com/elastic/elasticsearch-specification/blob/9fcf6a64c550d2e8090c8134867f200b56fd7fc7/specification/watcher/_types/Actions.ts#L268-L287
type IndexResult struct {
	// Request The request that would have been executed. It is only present when the action
	// is simulated.
	Request *IndexResultRequestSummary `json:"request,omitempty"`
	// Response The outcome of the index operation. A single summary is returned when a
	// single document is indexed, or an array when several documents are indexed at
	// once. When a bulk operation ends in `failure` or `partial_failure`, the array
	// includes failed items.
	Response []IndexResultSummary `json:"response,omitempty"`
}

func (s *IndexResult) UnmarshalJSON(data []byte) error {

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

		case "request":
			if err := dec.Decode(&s.Request); err != nil {
				return fmt.Errorf("%s | %w", "Request", err)
			}

		case "response":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewIndexResultSummary()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Response", err)
				}

				s.Response = append(s.Response, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Response); err != nil {
					return fmt.Errorf("%s | %w", "Response", err)
				}
			}

		}
	}
	return nil
}

// NewIndexResult returns a IndexResult.
func NewIndexResult() *IndexResult {
	r := &IndexResult{}

	return r
}

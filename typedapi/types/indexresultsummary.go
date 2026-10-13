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
// https://github.com/elastic/elasticsearch-specification/tree/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/result"
)

// A single item of an index action result. Successful items and failed items
// expose different fields; only `id` and `index` are present in both. Failed
// items appear when a bulk index action ends in `failure` or `partial_failure`.
//
// https://github.com/elastic/elasticsearch-specification/blob/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6/specification/watcher/_types/Actions.ts#L296-L316
type IndexResultSummary struct {
	Created *bool `json:"created,omitempty"`
	// Failed Only present for failed items
	Failed *bool  `json:"failed,omitempty"`
	Id     string `json:"id"`
	Index  string `json:"index"`
	// Message Only present for failed items
	Message *string        `json:"message,omitempty"`
	Result  *result.Result `json:"result,omitempty"`
	Version *int64         `json:"version,omitempty"`
}

func (s *IndexResultSummary) UnmarshalJSON(data []byte) error {

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

		case "created":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Created", err)
				}
				s.Created = &value
			case bool:
				s.Created = &v
			}

		case "failed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Failed", err)
				}
				s.Failed = &value
			case bool:
				s.Failed = &v
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "message":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Message", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Message = &o

		case "result":
			if err := dec.Decode(&s.Result); err != nil {
				return fmt.Errorf("%s | %w", "Result", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewIndexResultSummary returns a IndexResultSummary.
func NewIndexResultSummary() *IndexResultSummary {
	r := &IndexResultSummary{}

	return r
}

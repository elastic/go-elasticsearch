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
)

// PainlessContextSetup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/scripts_painless_execute/types.ts#L27-L46
type PainlessContextSetup struct {
	// Document Document that's temporarily indexed in-memory and accessible from the script.
	Document json.RawMessage `json:"document,omitempty"`
	// Index Index containing a mapping that's compatible with the indexed document.
	// You may specify a remote index by prefixing the index with the remote cluster
	// alias.
	// For example, `remote1:my_index` indicates that you want to run the painless
	// script against the "my_index" index on the "remote1" cluster.
	// This request will be forwarded to the "remote1" cluster if you have
	// configured a connection to that remote cluster.
	//
	// NOTE: Wildcards are not accepted in the index expression for this endpoint.
	// The expression `*:myindex` will return the error "No such remote cluster" and
	// the expression `logs*` or `remote1:logs*` will return the error "index not
	// found".
	Index string `json:"index"`
	// Query Use this parameter to specify a query for computing a score.
	Query *Query `json:"query,omitempty"`
}

func (s *PainlessContextSetup) UnmarshalJSON(data []byte) error {

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

		case "document":
			if err := dec.Decode(&s.Document); err != nil {
				return fmt.Errorf("%s | %w", "Document", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		}
	}
	return nil
}

// NewPainlessContextSetup returns a PainlessContextSetup.
func NewPainlessContextSetup() *PainlessContextSetup {
	r := &PainlessContextSetup{}

	return r
}

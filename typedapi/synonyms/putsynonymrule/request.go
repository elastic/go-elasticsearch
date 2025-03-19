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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package putsynonymrule

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Request holds the request body struct for the package putsynonymrule
//
// https://github.com/elastic/elasticsearch-specification/blob/c75a0abec670d027d13eb8d6f23374f86621c76b/specification/synonyms/put_synonym_rule/SynonymRulePutRequest.ts#L23-L60
type Request struct {

	// Synonyms The synonym rule information definition, which must be in Solr format.
	Synonyms string `json:"synonyms"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putsynonymrule request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "synonyms":
			if err := dec.Decode(&s.Synonyms); err != nil {
				return fmt.Errorf("%s | %w", "Synonyms", err)
			}

		}
	}
	return nil
}

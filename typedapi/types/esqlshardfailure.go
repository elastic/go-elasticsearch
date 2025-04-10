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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// EsqlShardFailure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/esql/_types/EsqlResult.ts#L88-L93
type EsqlShardFailure struct {
	Index  string     `json:"index"`
	Node   *string    `json:"node,omitempty"`
	Reason ErrorCause `json:"reason"`
	Shard  string     `json:"shard"`
}

func (s *EsqlShardFailure) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}

		case "reason":
			if err := dec.Decode(&s.Reason); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}

		case "shard":
			if err := dec.Decode(&s.Shard); err != nil {
				return fmt.Errorf("%s | %w", "Shard", err)
			}

		}
	}
	return nil
}

// NewEsqlShardFailure returns a EsqlShardFailure.
func NewEsqlShardFailure() *EsqlShardFailure {
	r := &EsqlShardFailure{}

	return r
}

// false

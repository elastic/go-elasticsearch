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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TermsPartition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/bucket.ts#L1080-L1089
type TermsPartition struct {
	// NumPartitions The number of partitions.
	NumPartitions int64 `json:"num_partitions"`
	// Partition The partition number for this request.
	Partition int64 `json:"partition"`
}

func (s *TermsPartition) UnmarshalJSON(data []byte) error {

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

		case "num_partitions":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumPartitions", err)
				}
				s.NumPartitions = value
			case float64:
				f := int64(v)
				s.NumPartitions = f
			}

		case "partition":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Partition", err)
				}
				s.Partition = value
			case float64:
				f := int64(v)
				s.Partition = f
			}

		}
	}
	return nil
}

// NewTermsPartition returns a TermsPartition.
func NewTermsPartition() *TermsPartition {
	r := &TermsPartition{}

	return r
}

type TermsPartitionVariant interface {
	TermsPartitionCaster() *TermsPartition
}

func (s *TermsPartition) TermsPartitionCaster() *TermsPartition {
	return s
}

func (s *TermsPartition) TermsIncludeCaster() *TermsInclude {
	o := TermsInclude(s)
	return &o
}

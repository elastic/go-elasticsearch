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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ShardsCapacityIndicatorTierDetail type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/health_report/types.ts#L182-L185
type ShardsCapacityIndicatorTierDetail struct {
	CurrentUsedShards  *int `json:"current_used_shards,omitempty"`
	MaxShardsInCluster int  `json:"max_shards_in_cluster"`
}

func (s *ShardsCapacityIndicatorTierDetail) UnmarshalJSON(data []byte) error {

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

		case "current_used_shards":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CurrentUsedShards = &value
			case float64:
				f := int(v)
				s.CurrentUsedShards = &f
			}

		case "max_shards_in_cluster":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxShardsInCluster = value
			case float64:
				f := int(v)
				s.MaxShardsInCluster = f
			}

		}
	}
	return nil
}

// NewShardsCapacityIndicatorTierDetail returns a ShardsCapacityIndicatorTierDetail.
func NewShardsCapacityIndicatorTierDetail() *ShardsCapacityIndicatorTierDetail {
	r := &ShardsCapacityIndicatorTierDetail{}

	return r
}

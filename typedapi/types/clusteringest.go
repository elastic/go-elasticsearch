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

// ClusterIngest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/cluster/stats/types.ts#L367-L370
type ClusterIngest struct {
	NumberOfPipelines int                         `json:"number_of_pipelines"`
	ProcessorStats    map[string]ClusterProcessor `json:"processor_stats"`
}

func (s *ClusterIngest) UnmarshalJSON(data []byte) error {

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

		case "number_of_pipelines":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfPipelines", err)
				}
				s.NumberOfPipelines = value
			case float64:
				f := int(v)
				s.NumberOfPipelines = f
			}

		case "processor_stats":
			if s.ProcessorStats == nil {
				s.ProcessorStats = make(map[string]ClusterProcessor, 0)
			}
			if err := dec.Decode(&s.ProcessorStats); err != nil {
				return fmt.Errorf("%s | %w", "ProcessorStats", err)
			}

		}
	}
	return nil
}

// NewClusterIngest returns a ClusterIngest.
func NewClusterIngest() *ClusterIngest {
	r := &ClusterIngest{
		ProcessorStats: make(map[string]ClusterProcessor),
	}

	return r
}

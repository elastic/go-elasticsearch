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

// ClusterNodeCount type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L457-L478
type ClusterNodeCount struct {
	CoordinatingOnly    *int `json:"coordinating_only,omitempty"`
	Data                *int `json:"data,omitempty"`
	DataCold            *int `json:"data_cold,omitempty"`
	DataContent         *int `json:"data_content,omitempty"`
	DataFrozen          *int `json:"data_frozen,omitempty"`
	DataHot             *int `json:"data_hot,omitempty"`
	DataWarm            *int `json:"data_warm,omitempty"`
	Index               *int `json:"index,omitempty"`
	Ingest              *int `json:"ingest,omitempty"`
	Master              *int `json:"master,omitempty"`
	Ml                  *int `json:"ml,omitempty"`
	RemoteClusterClient *int `json:"remote_cluster_client,omitempty"`
	Search              *int `json:"search,omitempty"`
	Total               int  `json:"total"`
	Transform           *int `json:"transform,omitempty"`
	VotingOnly          *int `json:"voting_only,omitempty"`
}

func (s *ClusterNodeCount) UnmarshalJSON(data []byte) error {

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

		case "coordinating_only":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CoordinatingOnly", err)
				}
				s.CoordinatingOnly = &value
			case float64:
				f := int(v)
				s.CoordinatingOnly = &f
			}

		case "data":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Data", err)
				}
				s.Data = &value
			case float64:
				f := int(v)
				s.Data = &f
			}

		case "data_cold":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataCold", err)
				}
				s.DataCold = &value
			case float64:
				f := int(v)
				s.DataCold = &f
			}

		case "data_content":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataContent", err)
				}
				s.DataContent = &value
			case float64:
				f := int(v)
				s.DataContent = &f
			}

		case "data_frozen":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataFrozen", err)
				}
				s.DataFrozen = &value
			case float64:
				f := int(v)
				s.DataFrozen = &f
			}

		case "data_hot":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataHot", err)
				}
				s.DataHot = &value
			case float64:
				f := int(v)
				s.DataHot = &f
			}

		case "data_warm":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataWarm", err)
				}
				s.DataWarm = &value
			case float64:
				f := int(v)
				s.DataWarm = &f
			}

		case "index":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Index", err)
				}
				s.Index = &value
			case float64:
				f := int(v)
				s.Index = &f
			}

		case "ingest":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Ingest", err)
				}
				s.Ingest = &value
			case float64:
				f := int(v)
				s.Ingest = &f
			}

		case "master":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Master", err)
				}
				s.Master = &value
			case float64:
				f := int(v)
				s.Master = &f
			}

		case "ml":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Ml", err)
				}
				s.Ml = &value
			case float64:
				f := int(v)
				s.Ml = &f
			}

		case "remote_cluster_client":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RemoteClusterClient", err)
				}
				s.RemoteClusterClient = &value
			case float64:
				f := int(v)
				s.RemoteClusterClient = &f
			}

		case "search":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Search", err)
				}
				s.Search = &value
			case float64:
				f := int(v)
				s.Search = &f
			}

		case "total":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int(v)
				s.Total = f
			}

		case "transform":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Transform", err)
				}
				s.Transform = &value
			case float64:
				f := int(v)
				s.Transform = &f
			}

		case "voting_only":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "VotingOnly", err)
				}
				s.VotingOnly = &value
			case float64:
				f := int(v)
				s.VotingOnly = &f
			}

		}
	}
	return nil
}

// NewClusterNodeCount returns a ClusterNodeCount.
func NewClusterNodeCount() *ClusterNodeCount {
	r := &ClusterNodeCount{}

	return r
}

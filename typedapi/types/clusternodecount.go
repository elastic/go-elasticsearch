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

// ClusterNodeCount type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/stats/types.ts#L348-L367
type ClusterNodeCount struct {
	CoordinatingOnly    int  `json:"coordinating_only"`
	Data                int  `json:"data"`
	DataCold            int  `json:"data_cold"`
	DataContent         int  `json:"data_content"`
	DataFrozen          *int `json:"data_frozen,omitempty"`
	DataHot             int  `json:"data_hot"`
	DataWarm            int  `json:"data_warm"`
	Ingest              int  `json:"ingest"`
	Master              int  `json:"master"`
	Ml                  int  `json:"ml"`
	RemoteClusterClient int  `json:"remote_cluster_client"`
	Total               int  `json:"total"`
	Transform           int  `json:"transform"`
	VotingOnly          int  `json:"voting_only"`
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

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CoordinatingOnly = value
			case float64:
				f := int(v)
				s.CoordinatingOnly = f
			}

		case "data":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Data = value
			case float64:
				f := int(v)
				s.Data = f
			}

		case "data_cold":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DataCold = value
			case float64:
				f := int(v)
				s.DataCold = f
			}

		case "data_content":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DataContent = value
			case float64:
				f := int(v)
				s.DataContent = f
			}

		case "data_frozen":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DataFrozen = &value
			case float64:
				f := int(v)
				s.DataFrozen = &f
			}

		case "data_hot":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DataHot = value
			case float64:
				f := int(v)
				s.DataHot = f
			}

		case "data_warm":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DataWarm = value
			case float64:
				f := int(v)
				s.DataWarm = f
			}

		case "ingest":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Ingest = value
			case float64:
				f := int(v)
				s.Ingest = f
			}

		case "master":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Master = value
			case float64:
				f := int(v)
				s.Master = f
			}

		case "ml":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Ml = value
			case float64:
				f := int(v)
				s.Ml = f
			}

		case "remote_cluster_client":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.RemoteClusterClient = value
			case float64:
				f := int(v)
				s.RemoteClusterClient = f
			}

		case "total":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Total = value
			case float64:
				f := int(v)
				s.Total = f
			}

		case "transform":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Transform = value
			case float64:
				f := int(v)
				s.Transform = f
			}

		case "voting_only":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.VotingOnly = value
			case float64:
				f := int(v)
				s.VotingOnly = f
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

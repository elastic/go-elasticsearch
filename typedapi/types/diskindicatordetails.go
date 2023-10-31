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

// DiskIndicatorDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/health_report/types.ts#L126-L132
type DiskIndicatorDetails struct {
	IndicesWithReadonlyBlock     int64 `json:"indices_with_readonly_block"`
	NodesOverFloodStageWatermark int64 `json:"nodes_over_flood_stage_watermark"`
	NodesOverHighWatermark       int64 `json:"nodes_over_high_watermark"`
	NodesWithEnoughDiskSpace     int64 `json:"nodes_with_enough_disk_space"`
	NodesWithUnknownDiskStatus   int64 `json:"nodes_with_unknown_disk_status"`
}

func (s *DiskIndicatorDetails) UnmarshalJSON(data []byte) error {

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

		case "indices_with_readonly_block":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndicesWithReadonlyBlock = value
			case float64:
				f := int64(v)
				s.IndicesWithReadonlyBlock = f
			}

		case "nodes_over_flood_stage_watermark":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NodesOverFloodStageWatermark = value
			case float64:
				f := int64(v)
				s.NodesOverFloodStageWatermark = f
			}

		case "nodes_over_high_watermark":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NodesOverHighWatermark = value
			case float64:
				f := int64(v)
				s.NodesOverHighWatermark = f
			}

		case "nodes_with_enough_disk_space":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NodesWithEnoughDiskSpace = value
			case float64:
				f := int64(v)
				s.NodesWithEnoughDiskSpace = f
			}

		case "nodes_with_unknown_disk_status":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NodesWithUnknownDiskStatus = value
			case float64:
				f := int64(v)
				s.NodesWithUnknownDiskStatus = f
			}

		}
	}
	return nil
}

// NewDiskIndicatorDetails returns a DiskIndicatorDetails.
func NewDiskIndicatorDetails() *DiskIndicatorDetails {
	r := &DiskIndicatorDetails{}

	return r
}

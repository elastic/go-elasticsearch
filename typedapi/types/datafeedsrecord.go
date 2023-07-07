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
// https://github.com/elastic/elasticsearch-specification/tree/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/datafeedstate"
)

// DatafeedsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c/specification/cat/ml_datafeeds/types.ts#L22-L83
type DatafeedsRecord struct {
	// AssignmentExplanation why the datafeed is or is not assigned to a node
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// BucketsCount bucket count
	BucketsCount *string `json:"buckets.count,omitempty"`
	// Id the datafeed_id
	Id *string `json:"id,omitempty"`
	// NodeAddress network address of the assigned node
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId ephemeral id of the assigned node
	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`
	// NodeId id of the assigned node
	NodeId *string `json:"node.id,omitempty"`
	// NodeName name of the assigned node
	NodeName *string `json:"node.name,omitempty"`
	// SearchBucketAvg the average search time per bucket (millisecond)
	SearchBucketAvg *string `json:"search.bucket_avg,omitempty"`
	// SearchCount number of searches ran by the datafeed
	SearchCount *string `json:"search.count,omitempty"`
	// SearchExpAvgHour the exponential average search time per hour (millisecond)
	SearchExpAvgHour *string `json:"search.exp_avg_hour,omitempty"`
	// SearchTime the total search time
	SearchTime *string `json:"search.time,omitempty"`
	// State the datafeed state
	State *datafeedstate.DatafeedState `json:"state,omitempty"`
}

func (s *DatafeedsRecord) UnmarshalJSON(data []byte) error {

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

		case "assignment_explanation", "ae":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AssignmentExplanation = &o

		case "buckets.count", "bc", "bucketsCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BucketsCount = &o

		case "id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id = &o

		case "node.address", "na", "nodeAddress":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeAddress = &o

		case "node.ephemeral_id", "ne", "nodeEphemeralId":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeEphemeralId = &o

		case "node.id", "ni", "nodeId":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeId = &o

		case "node.name", "nn", "nodeName":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeName = &o

		case "search.bucket_avg", "sba", "searchBucketAvg":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchBucketAvg = &o

		case "search.count", "sc", "searchCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchCount = &o

		case "search.exp_avg_hour", "seah", "searchExpAvgHour":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchExpAvgHour = &o

		case "search.time", "st", "searchTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchTime = &o

		case "state", "s":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDatafeedsRecord returns a DatafeedsRecord.
func NewDatafeedsRecord() *DatafeedsRecord {
	r := &DatafeedsRecord{}

	return r
}

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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

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
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cat/ml_datafeeds/types.ts#L22-L87
type DatafeedsRecord struct {
	// AssignmentExplanation For started datafeeds only, contains messages relating to the selection of a
	// node.
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// BucketsCount The number of buckets processed.
	BucketsCount *string `json:"buckets.count,omitempty"`
	// Id The datafeed identifier.
	Id *string `json:"id,omitempty"`
	// NodeAddress The network address of the assigned node.
	// For started datafeeds only, this information pertains to the node upon which
	// the datafeed is started.
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId The ephemeral identifier of the assigned node.
	// For started datafeeds only, this information pertains to the node upon which
	// the datafeed is started.
	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`
	// NodeId The unique identifier of the assigned node.
	// For started datafeeds only, this information pertains to the node upon which
	// the datafeed is started.
	NodeId *string `json:"node.id,omitempty"`
	// NodeName The name of the assigned node.
	// For started datafeeds only, this information pertains to the node upon which
	// the datafeed is started.
	NodeName *string `json:"node.name,omitempty"`
	// SearchBucketAvg The average search time per bucket, in milliseconds.
	SearchBucketAvg *string `json:"search.bucket_avg,omitempty"`
	// SearchCount The number of searches run by the datafeed.
	SearchCount *string `json:"search.count,omitempty"`
	// SearchExpAvgHour The exponential average search time per hour, in milliseconds.
	SearchExpAvgHour *string `json:"search.exp_avg_hour,omitempty"`
	// SearchTime The total time the datafeed spent searching, in milliseconds.
	SearchTime *string `json:"search.time,omitempty"`
	// State The status of the datafeed.
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

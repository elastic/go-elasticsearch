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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// DataStream type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/_types/DataStream.ts#L32-L52
type DataStream struct {
	AllowCustomRouting *bool                      `json:"allow_custom_routing,omitempty"`
	Generation         int                        `json:"generation"`
	Hidden             bool                       `json:"hidden"`
	IlmPolicy          *string                    `json:"ilm_policy,omitempty"`
	Indices            []DataStreamIndex          `json:"indices"`
	Lifecycle          *DataLifecycleWithRollover `json:"lifecycle,omitempty"`
	Meta_              Metadata                   `json:"_meta,omitempty"`
	Name               string                     `json:"name"`
	Replicated         *bool                      `json:"replicated,omitempty"`
	Status             healthstatus.HealthStatus  `json:"status"`
	System             *bool                      `json:"system,omitempty"`
	Template           string                     `json:"template"`
	TimestampField     DataStreamTimestampField   `json:"timestamp_field"`
}

func (s *DataStream) UnmarshalJSON(data []byte) error {

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

		case "allow_custom_routing":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowCustomRouting = &value
			case bool:
				s.AllowCustomRouting = &v
			}

		case "generation":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Generation = value
			case float64:
				f := int(v)
				s.Generation = f
			}

		case "hidden":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Hidden = value
			case bool:
				s.Hidden = v
			}

		case "ilm_policy":
			if err := dec.Decode(&s.IlmPolicy); err != nil {
				return err
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "lifecycle":
			if err := dec.Decode(&s.Lifecycle); err != nil {
				return err
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "replicated":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Replicated = &value
			case bool:
				s.Replicated = &v
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return err
			}

		case "system":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.System = &value
			case bool:
				s.System = &v
			}

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return err
			}

		case "timestamp_field":
			if err := dec.Decode(&s.TimestampField); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataStream returns a DataStream.
func NewDataStream() *DataStream {
	r := &DataStream{}

	return r
}

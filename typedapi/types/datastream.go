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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// DataStream type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/indices/_types/DataStream.ts#L32-L96
type DataStream struct {
	// AllowCustomRouting If `true`, the data stream allows custom routing on write request.
	AllowCustomRouting *bool `json:"allow_custom_routing,omitempty"`
	// Generation Current generation for the data stream. This number acts as a cumulative
	// count of the stream’s rollovers, starting at 1.
	Generation int `json:"generation"`
	// Hidden If `true`, the data stream is hidden.
	Hidden bool `json:"hidden"`
	// IlmPolicy Name of the current ILM lifecycle policy in the stream’s matching index
	// template.
	// This lifecycle policy is set in the `index.lifecycle.name` setting.
	// If the template does not include a lifecycle policy, this property is not
	// included in the response.
	// NOTE: A data stream’s backing indices may be assigned different lifecycle
	// policies. To retrieve the lifecycle policy for individual backing indices,
	// use the get index settings API.
	IlmPolicy *string `json:"ilm_policy,omitempty"`
	// Indices Array of objects containing information about the data stream’s backing
	// indices.
	// The last item in this array contains information about the stream’s current
	// write index.
	Indices []DataStreamIndex `json:"indices"`
	// Lifecycle Contains the configuration for the data lifecycle management of this data
	// stream.
	Lifecycle *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`
	// Meta_ Custom metadata for the stream, copied from the `_meta` object of the
	// stream’s matching index template.
	// If empty, the response omits this property.
	Meta_ Metadata `json:"_meta,omitempty"`
	// Name Name of the data stream.
	Name string `json:"name"`
	// Replicated If `true`, the data stream is created and managed by cross-cluster
	// replication and the local cluster can not write into this data stream or
	// change its mappings.
	Replicated *bool `json:"replicated,omitempty"`
	// Status Health status of the data stream.
	// This health status is based on the state of the primary and replica shards of
	// the stream’s backing indices.
	Status healthstatus.HealthStatus `json:"status"`
	// System If `true`, the data stream is created and managed by an Elastic stack
	// component and cannot be modified through normal user interaction.
	System *bool `json:"system,omitempty"`
	// Template Name of the index template used to create the data stream’s backing indices.
	// The template’s index pattern must match the name of this data stream.
	Template string `json:"template"`
	// TimestampField Information about the `@timestamp` field in the data stream.
	TimestampField DataStreamTimestampField `json:"timestamp_field"`
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

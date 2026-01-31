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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

// DataStreamLifecycleWithRollover type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/indices/_types/DataStreamLifecycle.ts#L53-L64
type DataStreamLifecycleWithRollover struct {
	// DataRetention If defined, every document added to this data stream will be stored at least
	// for this time frame.
	// Any time after this duration the document could be deleted.
	// When empty, every document in this data stream will be stored indefinitely.
	DataRetention Duration `json:"data_retention,omitempty"`
	// Downsampling The list of downsampling rounds to execute as part of this downsampling
	// configuration
	Downsampling []DownsamplingRound `json:"downsampling,omitempty"`
	// DownsamplingMethod The method used to downsample the data. There are two options `aggregate` and
	// `last_value`. It requires
	// `downsampling` to be defined. Defaults to `aggregate`.
	DownsamplingMethod *samplingmethod.SamplingMethod `json:"downsampling_method,omitempty"`
	// Enabled If defined, it turns data stream lifecycle on/off (`true`/`false`) for this
	// data stream. A data stream lifecycle
	// that's disabled (enabled: `false`) will have no effect on the data stream.
	Enabled *bool `json:"enabled,omitempty"`
	// Rollover The conditions which will trigger the rollover of a backing index as
	// configured by the cluster setting `cluster.lifecycle.default.rollover`.
	// This property is an implementation detail and it will only be retrieved when
	// the query param `include_defaults` is set to true.
	// The contents of this field are subject to change.
	Rollover *DataStreamLifecycleRolloverConditions `json:"rollover,omitempty"`
}

func (s *DataStreamLifecycleWithRollover) UnmarshalJSON(data []byte) error {

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

		case "data_retention":
			if err := dec.Decode(&s.DataRetention); err != nil {
				return fmt.Errorf("%s | %w", "DataRetention", err)
			}

		case "downsampling":
			if err := dec.Decode(&s.Downsampling); err != nil {
				return fmt.Errorf("%s | %w", "Downsampling", err)
			}

		case "downsampling_method":
			if err := dec.Decode(&s.DownsamplingMethod); err != nil {
				return fmt.Errorf("%s | %w", "DownsamplingMethod", err)
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		case "rollover":
			if err := dec.Decode(&s.Rollover); err != nil {
				return fmt.Errorf("%s | %w", "Rollover", err)
			}

		}
	}
	return nil
}

// NewDataStreamLifecycleWithRollover returns a DataStreamLifecycleWithRollover.
func NewDataStreamLifecycleWithRollover() *DataStreamLifecycleWithRollover {
	r := &DataStreamLifecycleWithRollover{}

	return r
}

type DataStreamLifecycleWithRolloverVariant interface {
	DataStreamLifecycleWithRolloverCaster() *DataStreamLifecycleWithRollover
}

func (s *DataStreamLifecycleWithRollover) DataStreamLifecycleWithRolloverCaster() *DataStreamLifecycleWithRollover {
	return s
}

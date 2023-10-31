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

// DataStreamLifecycleExplain type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/indices/explain_data_lifecycle/IndicesExplainDataLifecycleResponse.ts#L31-L41
type DataStreamLifecycleExplain struct {
	Error                   *string                          `json:"error,omitempty"`
	GenerationTime          Duration                         `json:"generation_time,omitempty"`
	Index                   string                           `json:"index"`
	IndexCreationDateMillis *int64                           `json:"index_creation_date_millis,omitempty"`
	Lifecycle               *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`
	ManagedByLifecycle      bool                             `json:"managed_by_lifecycle"`
	RolloverDateMillis      *int64                           `json:"rollover_date_millis,omitempty"`
	TimeSinceIndexCreation  Duration                         `json:"time_since_index_creation,omitempty"`
	TimeSinceRollover       Duration                         `json:"time_since_rollover,omitempty"`
}

func (s *DataStreamLifecycleExplain) UnmarshalJSON(data []byte) error {

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

		case "error":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Error = &o

		case "generation_time":
			if err := dec.Decode(&s.GenerationTime); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "index_creation_date_millis":
			if err := dec.Decode(&s.IndexCreationDateMillis); err != nil {
				return err
			}

		case "lifecycle":
			if err := dec.Decode(&s.Lifecycle); err != nil {
				return err
			}

		case "managed_by_lifecycle":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ManagedByLifecycle = value
			case bool:
				s.ManagedByLifecycle = v
			}

		case "rollover_date_millis":
			if err := dec.Decode(&s.RolloverDateMillis); err != nil {
				return err
			}

		case "time_since_index_creation":
			if err := dec.Decode(&s.TimeSinceIndexCreation); err != nil {
				return err
			}

		case "time_since_rollover":
			if err := dec.Decode(&s.TimeSinceRollover); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataStreamLifecycleExplain returns a DataStreamLifecycleExplain.
func NewDataStreamLifecycleExplain() *DataStreamLifecycleExplain {
	r := &DataStreamLifecycleExplain{}

	return r
}

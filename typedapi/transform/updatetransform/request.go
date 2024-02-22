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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package updatetransform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package updatetransform
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/transform/update_transform/UpdateTransformRequest.ts#L31-L105
type Request struct {

	// Description Free text description of the transform.
	Description *string `json:"description,omitempty"`
	// Dest The destination for the transform.
	Dest *types.TransformDestination `json:"dest,omitempty"`
	// Frequency The interval between checks for changes in the source indices when the
	// transform is running continuously. Also determines the retry interval in
	// the event of transient failures while the transform is searching or
	// indexing. The minimum value is 1s and the maximum is 1h.
	Frequency types.Duration `json:"frequency,omitempty"`
	// Meta_ Defines optional transform metadata.
	Meta_ types.Metadata `json:"_meta,omitempty"`
	// RetentionPolicy Defines a retention policy for the transform. Data that meets the defined
	// criteria is deleted from the destination index.
	RetentionPolicy types.RetentionPolicyContainer `json:"retention_policy,omitempty"`
	// Settings Defines optional transform settings.
	Settings *types.Settings `json:"settings,omitempty"`
	// Source The source of the data for the transform.
	Source *types.TransformSource `json:"source,omitempty"`
	// Sync Defines the properties transforms require to run continuously.
	Sync *types.SyncContainer `json:"sync,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Updatetransform request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "dest":
			if err := dec.Decode(&s.Dest); err != nil {
				return err
			}

		case "frequency":
			if err := dec.Decode(&s.Frequency); err != nil {
				return err
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return err
			}

		case "retention_policy":
			if err := dec.Decode(&s.RetentionPolicy); err != nil {
				return err
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return err
			}

		case "source":
			if err := dec.Decode(&s.Source); err != nil {
				return err
			}

		case "sync":
			if err := dec.Decode(&s.Sync); err != nil {
				return err
			}

		}
	}
	return nil
}

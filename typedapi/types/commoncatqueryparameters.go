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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CommonCatQueryParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/_spec_utils/behaviors.ts#L86-L132
type CommonCatQueryParameters struct {
	// Format Specifies the format to return the columnar data in, can be set to
	// `text`, `json`, `cbor`, `yaml`, or `smile`.
	Format *string `json:"format,omitempty"`
	// H List of columns to appear in the response. Supports simple wildcards.
	H []string `json:"h,omitempty"`
	// Help When set to `true` will output available columns. This option
	// can't be combined with any other query string option.
	Help *bool `json:"help,omitempty"`
	// Local If `true`, the request computes the list of selected nodes from the
	// local cluster state. If `false` the list of selected nodes are computed
	// from the cluster state of the master node. In both cases the coordinating
	// node will send requests for further information to each selected node.
	Local *bool `json:"local,omitempty"`
	// MasterTimeout Period to wait for a connection to the master node.
	MasterTimeout Duration `json:"master_timeout,omitempty"`
	// S List of columns that determine how the table should be sorted.
	// Sorting defaults to ascending and can be changed by setting `:asc`
	// or `:desc` as a suffix to the column name.
	S []string `json:"s,omitempty"`
	// V When set to `true` will enable verbose output.
	V *bool `json:"v,omitempty"`
}

func (s *CommonCatQueryParameters) UnmarshalJSON(data []byte) error {

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

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "h":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "H", err)
				}

				s.H = append(s.H, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.H); err != nil {
					return fmt.Errorf("%s | %w", "H", err)
				}
			}

		case "help":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Help", err)
				}
				s.Help = &value
			case bool:
				s.Help = &v
			}

		case "local":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Local", err)
				}
				s.Local = &value
			case bool:
				s.Local = &v
			}

		case "master_timeout":
			if err := dec.Decode(&s.MasterTimeout); err != nil {
				return fmt.Errorf("%s | %w", "MasterTimeout", err)
			}

		case "s":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "S", err)
				}

				s.S = append(s.S, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.S); err != nil {
					return fmt.Errorf("%s | %w", "S", err)
				}
			}

		case "v":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "V", err)
				}
				s.V = &value
			case bool:
				s.V = &v
			}

		}
	}
	return nil
}

// NewCommonCatQueryParameters returns a CommonCatQueryParameters.
func NewCommonCatQueryParameters() *CommonCatQueryParameters {
	r := &CommonCatQueryParameters{}

	return r
}

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
)

// ElasticsearchVersionInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/Base.ts#L54-L64
type ElasticsearchVersionInfo struct {
	BuildDate                        DateTime `json:"build_date"`
	BuildFlavor                      string   `json:"build_flavor"`
	BuildHash                        string   `json:"build_hash"`
	BuildSnapshot                    bool     `json:"build_snapshot"`
	BuildType                        string   `json:"build_type"`
	Int                              string   `json:"number"`
	LuceneVersion                    string   `json:"lucene_version"`
	MinimumIndexCompatibilityVersion string   `json:"minimum_index_compatibility_version"`
	MinimumWireCompatibilityVersion  string   `json:"minimum_wire_compatibility_version"`
}

func (s *ElasticsearchVersionInfo) UnmarshalJSON(data []byte) error {

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

		case "build_date":
			if err := dec.Decode(&s.BuildDate); err != nil {
				return err
			}

		case "build_flavor":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildFlavor = o

		case "build_hash":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildHash = o

		case "build_snapshot":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.BuildSnapshot = value
			case bool:
				s.BuildSnapshot = v
			}

		case "build_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildType = o

		case "number":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Int = o

		case "lucene_version":
			if err := dec.Decode(&s.LuceneVersion); err != nil {
				return err
			}

		case "minimum_index_compatibility_version":
			if err := dec.Decode(&s.MinimumIndexCompatibilityVersion); err != nil {
				return err
			}

		case "minimum_wire_compatibility_version":
			if err := dec.Decode(&s.MinimumWireCompatibilityVersion); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewElasticsearchVersionInfo returns a ElasticsearchVersionInfo.
func NewElasticsearchVersionInfo() *ElasticsearchVersionInfo {
	r := &ElasticsearchVersionInfo{}

	return r
}

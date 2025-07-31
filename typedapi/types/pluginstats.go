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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// PluginStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Stats.ts#L205-L215
type PluginStats struct {
	Classname            string   `json:"classname"`
	Description          string   `json:"description"`
	ElasticsearchVersion string   `json:"elasticsearch_version"`
	ExtendedPlugins      []string `json:"extended_plugins"`
	HasNativeController  bool     `json:"has_native_controller"`
	JavaVersion          string   `json:"java_version"`
	Licensed             bool     `json:"licensed"`
	Name                 string   `json:"name"`
	Version              string   `json:"version"`
}

func (s *PluginStats) UnmarshalJSON(data []byte) error {

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

		case "classname":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Classname", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Classname = o

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = o

		case "elasticsearch_version":
			if err := dec.Decode(&s.ElasticsearchVersion); err != nil {
				return fmt.Errorf("%s | %w", "ElasticsearchVersion", err)
			}

		case "extended_plugins":
			if err := dec.Decode(&s.ExtendedPlugins); err != nil {
				return fmt.Errorf("%s | %w", "ExtendedPlugins", err)
			}

		case "has_native_controller":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "HasNativeController", err)
				}
				s.HasNativeController = value
			case bool:
				s.HasNativeController = v
			}

		case "java_version":
			if err := dec.Decode(&s.JavaVersion); err != nil {
				return fmt.Errorf("%s | %w", "JavaVersion", err)
			}

		case "licensed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Licensed", err)
				}
				s.Licensed = value
			case bool:
				s.Licensed = v
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewPluginStats returns a PluginStats.
func NewPluginStats() *PluginStats {
	r := &PluginStats{}

	return r
}

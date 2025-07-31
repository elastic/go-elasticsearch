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

// ElasticsearchVersionInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Base.ts#L76-L118
type ElasticsearchVersionInfo struct {
	// BuildDate The Elasticsearch Git commit's date.
	BuildDate DateTime `json:"build_date"`
	// BuildFlavor The build flavor. For example, `default`.
	BuildFlavor string `json:"build_flavor"`
	// BuildHash The Elasticsearch Git commit's SHA hash.
	BuildHash string `json:"build_hash"`
	// BuildSnapshot Indicates whether the Elasticsearch build was a snapshot.
	BuildSnapshot bool `json:"build_snapshot"`
	// BuildType The build type that corresponds to how Elasticsearch was installed.
	// For example, `docker`, `rpm`, or `tar`.
	BuildType string `json:"build_type"`
	// Int The Elasticsearch version number.
	//
	// ::: IMPORTANT: For Serverless deployments, this static value is always
	// `8.11.0` and is used solely for backward compatibility with legacy clients.
	//
	//	Serverless environments are versionless and automatically upgraded, so this
	//
	// value can be safely ignored.
	Int string `json:"number"`
	// LuceneVersion The version number of Elasticsearch's underlying Lucene software.
	LuceneVersion string `json:"lucene_version"`
	// MinimumIndexCompatibilityVersion The minimum index version with which the responding node can read from disk.
	MinimumIndexCompatibilityVersion string `json:"minimum_index_compatibility_version"`
	// MinimumWireCompatibilityVersion The minimum node version with which the responding node can communicate.
	// Also the minimum version from which you can perform a rolling upgrade.
	MinimumWireCompatibilityVersion string `json:"minimum_wire_compatibility_version"`
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
				return fmt.Errorf("%s | %w", "BuildDate", err)
			}

		case "build_flavor":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BuildFlavor", err)
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
				return fmt.Errorf("%s | %w", "BuildHash", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BuildHash = o

		case "build_snapshot":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildSnapshot", err)
				}
				s.BuildSnapshot = value
			case bool:
				s.BuildSnapshot = v
			}

		case "build_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BuildType", err)
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
				return fmt.Errorf("%s | %w", "Int", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Int = o

		case "lucene_version":
			if err := dec.Decode(&s.LuceneVersion); err != nil {
				return fmt.Errorf("%s | %w", "LuceneVersion", err)
			}

		case "minimum_index_compatibility_version":
			if err := dec.Decode(&s.MinimumIndexCompatibilityVersion); err != nil {
				return fmt.Errorf("%s | %w", "MinimumIndexCompatibilityVersion", err)
			}

		case "minimum_wire_compatibility_version":
			if err := dec.Decode(&s.MinimumWireCompatibilityVersion); err != nil {
				return fmt.Errorf("%s | %w", "MinimumWireCompatibilityVersion", err)
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

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
// https://github.com/elastic/elasticsearch-specification/tree/1cdb0310e2ed76dea2491f2551fca3f9143b6dc4

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// BuildInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1cdb0310e2ed76dea2491f2551fca3f9143b6dc4/specification/prometheus/status_buildinfo/types.ts#L23-L40
type BuildInfo struct {
	// Application The name of the application. It is always `Elasticsearch`.
	Application string `json:"application"`
	// BuildDate The date of the Elasticsearch build.
	BuildDate DateTime `json:"buildDate"`
	// Revision The commit hash of the Elasticsearch build.
	Revision string `json:"revision"`
	// Version The Elasticsearch version.
	Version string `json:"version"`
}

func (s *BuildInfo) UnmarshalJSON(data []byte) error {

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

		case "application":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Application", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Application = o

		case "buildDate":
			if err := dec.Decode(&s.BuildDate); err != nil {
				return fmt.Errorf("%s | %w", "BuildDate", err)
			}

		case "revision":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Revision", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Revision = o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewBuildInfo returns a BuildInfo.
func NewBuildInfo() *BuildInfo {
	r := &BuildInfo{}

	return r
}

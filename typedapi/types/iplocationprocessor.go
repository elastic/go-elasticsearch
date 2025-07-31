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

// IpLocationProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Processors.ts#L478-L512
type IpLocationProcessor struct {
	// DatabaseFile The database filename referring to a database the module ships with
	// (GeoLite2-City.mmdb, GeoLite2-Country.mmdb, or GeoLite2-ASN.mmdb) or a custom
	// database in the ingest-geoip config directory.
	DatabaseFile *string `json:"database_file,omitempty"`
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// DownloadDatabaseOnPipelineCreation If `true` (and if `ingest.geoip.downloader.eager.download` is `false`), the
	// missing database is downloaded when the pipeline is created.
	// Else, the download is triggered by when the pipeline is used as the
	// `default_pipeline` or `final_pipeline` in an index.
	DownloadDatabaseOnPipelineCreation *bool `json:"download_database_on_pipeline_creation,omitempty"`
	// Field The field to get the ip address from for the geographical lookup.
	Field string `json:"field"`
	// FirstOnly If `true`, only the first found IP location data will be returned, even if
	// the field contains an array.
	FirstOnly *bool `json:"first_only,omitempty"`
	// If Conditionally execute the processor.
	If *string `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If `true` and `field` does not exist, the processor quietly exits without
	// modifying the document.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Properties Controls what properties are added to the `target_field` based on the IP
	// location lookup.
	Properties []string `json:"properties,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TargetField The field that will hold the geographical information looked up from the
	// MaxMind database.
	TargetField *string `json:"target_field,omitempty"`
}

func (s *IpLocationProcessor) UnmarshalJSON(data []byte) error {

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

		case "database_file":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DatabaseFile", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DatabaseFile = &o

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
			s.Description = &o

		case "download_database_on_pipeline_creation":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DownloadDatabaseOnPipelineCreation", err)
				}
				s.DownloadDatabaseOnPipelineCreation = &value
			case bool:
				s.DownloadDatabaseOnPipelineCreation = &v
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "first_only":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "FirstOnly", err)
				}
				s.FirstOnly = &value
			case bool:
				s.FirstOnly = &v
			}

		case "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "If", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.If = &o

		case "ignore_failure":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreFailure", err)
				}
				s.IgnoreFailure = &value
			case bool:
				s.IgnoreFailure = &v
			}

		case "ignore_missing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreMissing", err)
				}
				s.IgnoreMissing = &value
			case bool:
				s.IgnoreMissing = &v
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
			}

		case "properties":
			if err := dec.Decode(&s.Properties); err != nil {
				return fmt.Errorf("%s | %w", "Properties", err)
			}

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tag", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tag = &o

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return fmt.Errorf("%s | %w", "TargetField", err)
			}

		}
	}
	return nil
}

// NewIpLocationProcessor returns a IpLocationProcessor.
func NewIpLocationProcessor() *IpLocationProcessor {
	r := &IpLocationProcessor{}

	return r
}

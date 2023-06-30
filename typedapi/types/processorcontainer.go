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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// ProcessorContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ingest/_types/Processors.ts#L28-L67
type ProcessorContainer struct {
	Append          *AppendProcessor          `json:"append,omitempty"`
	Attachment      *AttachmentProcessor      `json:"attachment,omitempty"`
	Bytes           *BytesProcessor           `json:"bytes,omitempty"`
	Circle          *CircleProcessor          `json:"circle,omitempty"`
	Convert         *ConvertProcessor         `json:"convert,omitempty"`
	Csv             *CsvProcessor             `json:"csv,omitempty"`
	Date            *DateProcessor            `json:"date,omitempty"`
	DateIndexName   *DateIndexNameProcessor   `json:"date_index_name,omitempty"`
	Dissect         *DissectProcessor         `json:"dissect,omitempty"`
	DotExpander     *DotExpanderProcessor     `json:"dot_expander,omitempty"`
	Drop            *DropProcessor            `json:"drop,omitempty"`
	Enrich          *EnrichProcessor          `json:"enrich,omitempty"`
	Fail            *FailProcessor            `json:"fail,omitempty"`
	Foreach         *ForeachProcessor         `json:"foreach,omitempty"`
	Geoip           *GeoIpProcessor           `json:"geoip,omitempty"`
	Grok            *GrokProcessor            `json:"grok,omitempty"`
	Gsub            *GsubProcessor            `json:"gsub,omitempty"`
	Inference       *InferenceProcessor       `json:"inference,omitempty"`
	Join            *JoinProcessor            `json:"join,omitempty"`
	Json            *JsonProcessor            `json:"json,omitempty"`
	Kv              *KeyValueProcessor        `json:"kv,omitempty"`
	Lowercase       *LowercaseProcessor       `json:"lowercase,omitempty"`
	Pipeline        *PipelineProcessor        `json:"pipeline,omitempty"`
	Remove          *RemoveProcessor          `json:"remove,omitempty"`
	Rename          *RenameProcessor          `json:"rename,omitempty"`
	Script          Script                    `json:"script,omitempty"`
	Set             *SetProcessor             `json:"set,omitempty"`
	SetSecurityUser *SetSecurityUserProcessor `json:"set_security_user,omitempty"`
	Sort            *SortProcessor            `json:"sort,omitempty"`
	Split           *SplitProcessor           `json:"split,omitempty"`
	Trim            *TrimProcessor            `json:"trim,omitempty"`
	Uppercase       *UppercaseProcessor       `json:"uppercase,omitempty"`
	Urldecode       *UrlDecodeProcessor       `json:"urldecode,omitempty"`
	UserAgent       *UserAgentProcessor       `json:"user_agent,omitempty"`
}

func (s *ProcessorContainer) UnmarshalJSON(data []byte) error {

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

		case "append":
			if err := dec.Decode(&s.Append); err != nil {
				return err
			}

		case "attachment":
			if err := dec.Decode(&s.Attachment); err != nil {
				return err
			}

		case "bytes":
			if err := dec.Decode(&s.Bytes); err != nil {
				return err
			}

		case "circle":
			if err := dec.Decode(&s.Circle); err != nil {
				return err
			}

		case "convert":
			if err := dec.Decode(&s.Convert); err != nil {
				return err
			}

		case "csv":
			if err := dec.Decode(&s.Csv); err != nil {
				return err
			}

		case "date":
			if err := dec.Decode(&s.Date); err != nil {
				return err
			}

		case "date_index_name":
			if err := dec.Decode(&s.DateIndexName); err != nil {
				return err
			}

		case "dissect":
			if err := dec.Decode(&s.Dissect); err != nil {
				return err
			}

		case "dot_expander":
			if err := dec.Decode(&s.DotExpander); err != nil {
				return err
			}

		case "drop":
			if err := dec.Decode(&s.Drop); err != nil {
				return err
			}

		case "enrich":
			if err := dec.Decode(&s.Enrich); err != nil {
				return err
			}

		case "fail":
			if err := dec.Decode(&s.Fail); err != nil {
				return err
			}

		case "foreach":
			if err := dec.Decode(&s.Foreach); err != nil {
				return err
			}

		case "geoip":
			if err := dec.Decode(&s.Geoip); err != nil {
				return err
			}

		case "grok":
			if err := dec.Decode(&s.Grok); err != nil {
				return err
			}

		case "gsub":
			if err := dec.Decode(&s.Gsub); err != nil {
				return err
			}

		case "inference":
			if err := dec.Decode(&s.Inference); err != nil {
				return err
			}

		case "join":
			if err := dec.Decode(&s.Join); err != nil {
				return err
			}

		case "json":
			if err := dec.Decode(&s.Json); err != nil {
				return err
			}

		case "kv":
			if err := dec.Decode(&s.Kv); err != nil {
				return err
			}

		case "lowercase":
			if err := dec.Decode(&s.Lowercase); err != nil {
				return err
			}

		case "pipeline":
			if err := dec.Decode(&s.Pipeline); err != nil {
				return err
			}

		case "remove":
			if err := dec.Decode(&s.Remove); err != nil {
				return err
			}

		case "rename":
			if err := dec.Decode(&s.Rename); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "set":
			if err := dec.Decode(&s.Set); err != nil {
				return err
			}

		case "set_security_user":
			if err := dec.Decode(&s.SetSecurityUser); err != nil {
				return err
			}

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return err
			}

		case "split":
			if err := dec.Decode(&s.Split); err != nil {
				return err
			}

		case "trim":
			if err := dec.Decode(&s.Trim); err != nil {
				return err
			}

		case "uppercase":
			if err := dec.Decode(&s.Uppercase); err != nil {
				return err
			}

		case "urldecode":
			if err := dec.Decode(&s.Urldecode); err != nil {
				return err
			}

		case "user_agent":
			if err := dec.Decode(&s.UserAgent); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewProcessorContainer returns a ProcessorContainer.
func NewProcessorContainer() *ProcessorContainer {
	r := &ProcessorContainer{}

	return r
}

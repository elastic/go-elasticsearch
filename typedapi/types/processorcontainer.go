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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// ProcessorContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ingest/_types/Processors.ts#L28-L67
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

// NewProcessorContainer returns a ProcessorContainer.
func NewProcessorContainer() *ProcessorContainer {
	r := &ProcessorContainer{}

	return r
}

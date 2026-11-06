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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/documentextractionoutputformat"
)

// Task settings for the `document_extraction` task type of the `elastic`
// service. They can be set when creating the inference endpoint and overridden
// per request. Request values take precedence, and `parameters` are merged per
// key.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/document_extraction/DocumentExtractionRequest.ts#L103-L118
type DocumentExtractionTaskSettings struct {
	// OutputFormat The output format of the extracted content. Only `markdown` is supported.
	OutputFormat *documentextractionoutputformat.DocumentExtractionOutputFormat `json:"output_format,omitempty"`
	// Parameters A free-form object of extraction options that is forwarded to the Elastic
	// Inference Service as-is. The accepted parameters are defined by the Elastic
	// Inference Service.
	Parameters map[string]json.RawMessage `json:"parameters,omitempty"`
}

// NewDocumentExtractionTaskSettings returns a DocumentExtractionTaskSettings.
func NewDocumentExtractionTaskSettings() *DocumentExtractionTaskSettings {
	r := &DocumentExtractionTaskSettings{
		Parameters: make(map[string]json.RawMessage),
	}

	return r
}

type DocumentExtractionTaskSettingsVariant interface {
	DocumentExtractionTaskSettingsCaster() *DocumentExtractionTaskSettings
}

func (s *DocumentExtractionTaskSettings) DocumentExtractionTaskSettingsCaster() *DocumentExtractionTaskSettings {
	return s
}

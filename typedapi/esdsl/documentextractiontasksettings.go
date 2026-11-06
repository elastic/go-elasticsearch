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

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/documentextractionoutputformat"
)

type _documentExtractionTaskSettings struct {
	v *types.DocumentExtractionTaskSettings
}

func NewDocumentExtractionTaskSettings() *_documentExtractionTaskSettings {

	return &_documentExtractionTaskSettings{v: types.NewDocumentExtractionTaskSettings()}

}

func (s *_documentExtractionTaskSettings) OutputFormat(outputformat documentextractionoutputformat.DocumentExtractionOutputFormat) *_documentExtractionTaskSettings {

	s.v.OutputFormat = &outputformat
	return s
}

func (s *_documentExtractionTaskSettings) Parameters(parameters map[string]json.RawMessage) *_documentExtractionTaskSettings {

	s.v.Parameters = parameters
	return s
}

func (s *_documentExtractionTaskSettings) AddParameter(key string, value json.RawMessage) *_documentExtractionTaskSettings {

	var tmp map[string]json.RawMessage
	if s.v.Parameters == nil {
		s.v.Parameters = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Parameters
	}

	tmp[key] = value

	s.v.Parameters = tmp
	return s
}

func (s *_documentExtractionTaskSettings) DocumentExtractionTaskSettingsCaster() *types.DocumentExtractionTaskSettings {
	return s.v
}

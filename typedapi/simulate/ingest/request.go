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

package ingest

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package ingest
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/simulate/ingest/SimulateIngestRequest.ts#L29-L100
type Request struct {

	// ComponentTemplateSubstitutions A map of component template names to substitute component template definition
	// objects.
	ComponentTemplateSubstitutions map[string]types.ComponentTemplateNode `json:"component_template_substitutions,omitempty"`
	// Docs Sample documents to test in the pipeline.
	Docs []types.Document `json:"docs"`
	// IndexTemplateSubstitutions A map of index template names to substitute index template definition
	// objects.
	IndexTemplateSubstitutions map[string]types.IndexTemplate `json:"index_template_substitutions,omitempty"`
	MappingAddition            *types.TypeMapping             `json:"mapping_addition,omitempty"`
	// PipelineSubstitutions Pipelines to test.
	// If you don’t specify the `pipeline` request path parameter, this parameter is
	// required.
	// If you specify both this and the request path parameter, the API only uses
	// the request path parameter.
	PipelineSubstitutions map[string]types.IngestPipeline `json:"pipeline_substitutions,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		ComponentTemplateSubstitutions: make(map[string]types.ComponentTemplateNode, 0),
		IndexTemplateSubstitutions:     make(map[string]types.IndexTemplate, 0),
		PipelineSubstitutions:          make(map[string]types.IngestPipeline, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Ingest request: %w", err)
	}

	return &req, nil
}

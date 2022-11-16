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
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// PipelineProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/ingest/_types/Processors.ts#L306-L309
type PipelineProcessor struct {
	Description           *string              `json:"description,omitempty"`
	If                    *string              `json:"if,omitempty"`
	IgnoreFailure         *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissingPipeline *bool                `json:"ignore_missing_pipeline,omitempty"`
	Name                  string               `json:"name"`
	OnFailure             []ProcessorContainer `json:"on_failure,omitempty"`
	Tag                   *string              `json:"tag,omitempty"`
}

// NewPipelineProcessor returns a PipelineProcessor.
func NewPipelineProcessor() *PipelineProcessor {
	r := &PipelineProcessor{}

	return r
}

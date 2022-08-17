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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// MlInferenceIngestProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L196-L201
type MlInferenceIngestProcessor struct {
	NumDocsProcessed MlInferenceIngestProcessorCount `json:"num_docs_processed"`
	NumFailures      MlInferenceIngestProcessorCount `json:"num_failures"`
	Pipelines        MlCounter                       `json:"pipelines"`
	TimeMs           MlInferenceIngestProcessorCount `json:"time_ms"`
}

// MlInferenceIngestProcessorBuilder holds MlInferenceIngestProcessor struct and provides a builder API.
type MlInferenceIngestProcessorBuilder struct {
	v *MlInferenceIngestProcessor
}

// NewMlInferenceIngestProcessor provides a builder for the MlInferenceIngestProcessor struct.
func NewMlInferenceIngestProcessorBuilder() *MlInferenceIngestProcessorBuilder {
	r := MlInferenceIngestProcessorBuilder{
		&MlInferenceIngestProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the MlInferenceIngestProcessor struct
func (rb *MlInferenceIngestProcessorBuilder) Build() MlInferenceIngestProcessor {
	return *rb.v
}

func (rb *MlInferenceIngestProcessorBuilder) NumDocsProcessed(numdocsprocessed *MlInferenceIngestProcessorCountBuilder) *MlInferenceIngestProcessorBuilder {
	v := numdocsprocessed.Build()
	rb.v.NumDocsProcessed = v
	return rb
}

func (rb *MlInferenceIngestProcessorBuilder) NumFailures(numfailures *MlInferenceIngestProcessorCountBuilder) *MlInferenceIngestProcessorBuilder {
	v := numfailures.Build()
	rb.v.NumFailures = v
	return rb
}

func (rb *MlInferenceIngestProcessorBuilder) Pipelines(pipelines *MlCounterBuilder) *MlInferenceIngestProcessorBuilder {
	v := pipelines.Build()
	rb.v.Pipelines = v
	return rb
}

func (rb *MlInferenceIngestProcessorBuilder) TimeMs(timems *MlInferenceIngestProcessorCountBuilder) *MlInferenceIngestProcessorBuilder {
	v := timems.Build()
	rb.v.TimeMs = v
	return rb
}

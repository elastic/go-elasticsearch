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

// DataframeAnalysisAnalyzedFields type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L238-L244
type DataframeAnalysisAnalyzedFields struct {
	// Excludes An array of strings that defines the fields that will be included in the
	// analysis.
	Excludes []string `json:"excludes"`
	// Includes An array of strings that defines the fields that will be excluded from the
	// analysis. You do not need to add fields with unsupported data types to
	// excludes, these fields are excluded from the analysis automatically.
	Includes []string `json:"includes"`
}

// DataframeAnalysisAnalyzedFieldsBuilder holds DataframeAnalysisAnalyzedFields struct and provides a builder API.
type DataframeAnalysisAnalyzedFieldsBuilder struct {
	v *DataframeAnalysisAnalyzedFields
}

// NewDataframeAnalysisAnalyzedFields provides a builder for the DataframeAnalysisAnalyzedFields struct.
func NewDataframeAnalysisAnalyzedFieldsBuilder() *DataframeAnalysisAnalyzedFieldsBuilder {
	r := DataframeAnalysisAnalyzedFieldsBuilder{
		&DataframeAnalysisAnalyzedFields{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalysisAnalyzedFields struct
func (rb *DataframeAnalysisAnalyzedFieldsBuilder) Build() DataframeAnalysisAnalyzedFields {
	return *rb.v
}

// Excludes An array of strings that defines the fields that will be included in the
// analysis.

func (rb *DataframeAnalysisAnalyzedFieldsBuilder) Excludes(excludes ...string) *DataframeAnalysisAnalyzedFieldsBuilder {
	rb.v.Excludes = excludes
	return rb
}

// Includes An array of strings that defines the fields that will be excluded from the
// analysis. You do not need to add fields with unsupported data types to
// excludes, these fields are excluded from the analysis automatically.

func (rb *DataframeAnalysisAnalyzedFieldsBuilder) Includes(includes ...string) *DataframeAnalysisAnalyzedFieldsBuilder {
	rb.v.Includes = includes
	return rb
}

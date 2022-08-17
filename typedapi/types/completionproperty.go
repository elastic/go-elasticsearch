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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// CompletionProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/specialized.ts#L26-L34
type CompletionProperty struct {
	Analyzer                   *string                        `json:"analyzer,omitempty"`
	Contexts                   []SuggestContext               `json:"contexts,omitempty"`
	CopyTo                     *Fields                        `json:"copy_to,omitempty"`
	DocValues                  *bool                          `json:"doc_values,omitempty"`
	Dynamic                    *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields                     map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove                *int                           `json:"ignore_above,omitempty"`
	LocalMetadata              *Metadata                      `json:"local_metadata,omitempty"`
	MaxInputLength             *int                           `json:"max_input_length,omitempty"`
	Meta                       map[string]string              `json:"meta,omitempty"`
	PreservePositionIncrements *bool                          `json:"preserve_position_increments,omitempty"`
	PreserveSeparators         *bool                          `json:"preserve_separators,omitempty"`
	Properties                 map[PropertyName]Property      `json:"properties,omitempty"`
	SearchAnalyzer             *string                        `json:"search_analyzer,omitempty"`
	Similarity                 *string                        `json:"similarity,omitempty"`
	Store                      *bool                          `json:"store,omitempty"`
	Type                       string                         `json:"type,omitempty"`
}

// CompletionPropertyBuilder holds CompletionProperty struct and provides a builder API.
type CompletionPropertyBuilder struct {
	v *CompletionProperty
}

// NewCompletionProperty provides a builder for the CompletionProperty struct.
func NewCompletionPropertyBuilder() *CompletionPropertyBuilder {
	r := CompletionPropertyBuilder{
		&CompletionProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "completion"

	return &r
}

// Build finalize the chain and returns the CompletionProperty struct
func (rb *CompletionPropertyBuilder) Build() CompletionProperty {
	return *rb.v
}

func (rb *CompletionPropertyBuilder) Analyzer(analyzer string) *CompletionPropertyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *CompletionPropertyBuilder) Contexts(contexts []SuggestContextBuilder) *CompletionPropertyBuilder {
	tmp := make([]SuggestContext, len(contexts))
	for _, value := range contexts {
		tmp = append(tmp, value.Build())
	}
	rb.v.Contexts = tmp
	return rb
}

func (rb *CompletionPropertyBuilder) CopyTo(copyto *FieldsBuilder) *CompletionPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *CompletionPropertyBuilder) DocValues(docvalues bool) *CompletionPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *CompletionPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *CompletionPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *CompletionPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *CompletionPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *CompletionPropertyBuilder) IgnoreAbove(ignoreabove int) *CompletionPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *CompletionPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *CompletionPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *CompletionPropertyBuilder) MaxInputLength(maxinputlength int) *CompletionPropertyBuilder {
	rb.v.MaxInputLength = &maxinputlength
	return rb
}

func (rb *CompletionPropertyBuilder) Meta(value map[string]string) *CompletionPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *CompletionPropertyBuilder) PreservePositionIncrements(preservepositionincrements bool) *CompletionPropertyBuilder {
	rb.v.PreservePositionIncrements = &preservepositionincrements
	return rb
}

func (rb *CompletionPropertyBuilder) PreserveSeparators(preserveseparators bool) *CompletionPropertyBuilder {
	rb.v.PreserveSeparators = &preserveseparators
	return rb
}

func (rb *CompletionPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *CompletionPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *CompletionPropertyBuilder) SearchAnalyzer(searchanalyzer string) *CompletionPropertyBuilder {
	rb.v.SearchAnalyzer = &searchanalyzer
	return rb
}

func (rb *CompletionPropertyBuilder) Similarity(similarity string) *CompletionPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *CompletionPropertyBuilder) Store(store bool) *CompletionPropertyBuilder {
	rb.v.Store = &store
	return rb
}

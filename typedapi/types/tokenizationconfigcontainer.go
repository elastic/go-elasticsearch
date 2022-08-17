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

// TokenizationConfigContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L97-L114
type TokenizationConfigContainer struct {
	// Bert Indicates BERT tokenization and its options
	Bert *NlpBertTokenizationConfig `json:"bert,omitempty"`
	// Mpnet Indicates MPNET tokenization and its options
	Mpnet *NlpBertTokenizationConfig `json:"mpnet,omitempty"`
	// Roberta Indicates RoBERTa tokenization and its options
	Roberta *NlpRobertaTokenizationConfig `json:"roberta,omitempty"`
}

// TokenizationConfigContainerBuilder holds TokenizationConfigContainer struct and provides a builder API.
type TokenizationConfigContainerBuilder struct {
	v *TokenizationConfigContainer
}

// NewTokenizationConfigContainer provides a builder for the TokenizationConfigContainer struct.
func NewTokenizationConfigContainerBuilder() *TokenizationConfigContainerBuilder {
	r := TokenizationConfigContainerBuilder{
		&TokenizationConfigContainer{},
	}

	return &r
}

// Build finalize the chain and returns the TokenizationConfigContainer struct
func (rb *TokenizationConfigContainerBuilder) Build() TokenizationConfigContainer {
	return *rb.v
}

// Bert Indicates BERT tokenization and its options

func (rb *TokenizationConfigContainerBuilder) Bert(bert *NlpBertTokenizationConfigBuilder) *TokenizationConfigContainerBuilder {
	v := bert.Build()
	rb.v.Bert = &v
	return rb
}

// Mpnet Indicates MPNET tokenization and its options

func (rb *TokenizationConfigContainerBuilder) Mpnet(mpnet *NlpBertTokenizationConfigBuilder) *TokenizationConfigContainerBuilder {
	v := mpnet.Build()
	rb.v.Mpnet = &v
	return rb
}

// Roberta Indicates RoBERTa tokenization and its options

func (rb *TokenizationConfigContainerBuilder) Roberta(roberta *NlpRobertaTokenizationConfigBuilder) *TokenizationConfigContainerBuilder {
	v := roberta.Build()
	rb.v.Roberta = &v
	return rb
}

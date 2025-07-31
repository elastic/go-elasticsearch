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

package types

import (
	"encoding/json"
	"fmt"
)

// TokenizationConfigContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/inference.ts#L135-L158
type TokenizationConfigContainer struct {
	AdditionalTokenizationConfigContainerProperty map[string]json.RawMessage `json:"-"`
	// Bert Indicates BERT tokenization and its options
	Bert *NlpBertTokenizationConfig `json:"bert,omitempty"`
	// BertJa Indicates BERT Japanese tokenization and its options
	BertJa *NlpBertTokenizationConfig `json:"bert_ja,omitempty"`
	// Mpnet Indicates MPNET tokenization and its options
	Mpnet *NlpBertTokenizationConfig `json:"mpnet,omitempty"`
	// Roberta Indicates RoBERTa tokenization and its options
	Roberta    *NlpRobertaTokenizationConfig `json:"roberta,omitempty"`
	XlmRoberta *XlmRobertaTokenizationConfig `json:"xlm_roberta,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s TokenizationConfigContainer) MarshalJSON() ([]byte, error) {
	type opt TokenizationConfigContainer
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalTokenizationConfigContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalTokenizationConfigContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewTokenizationConfigContainer returns a TokenizationConfigContainer.
func NewTokenizationConfigContainer() *TokenizationConfigContainer {
	r := &TokenizationConfigContainer{
		AdditionalTokenizationConfigContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}

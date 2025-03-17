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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"encoding/json"
	"fmt"
)

// RetrieverContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/_types/Retriever.ts#L28-L42
type RetrieverContainer struct {
	AdditionalRetrieverContainerProperty map[string]json.RawMessage `json:"-"`
	// Knn A retriever that replaces the functionality  of a knn search.
	Knn *KnnRetriever `json:"knn,omitempty"`
	// Rrf A retriever that produces top documents from reciprocal rank fusion (RRF).
	Rrf *RRFRetriever `json:"rrf,omitempty"`
	// Rule A retriever that replaces the functionality of a rule query.
	Rule *RuleRetriever `json:"rule,omitempty"`
	// Standard A retriever that replaces the functionality of a traditional query.
	Standard *StandardRetriever `json:"standard,omitempty"`
	// TextSimilarityReranker A retriever that reranks the top documents based on a reranking model using
	// the InferenceAPI
	TextSimilarityReranker *TextSimilarityReranker `json:"text_similarity_reranker,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s RetrieverContainer) MarshalJSON() ([]byte, error) {
	type opt RetrieverContainer
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
	for key, value := range s.AdditionalRetrieverContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalRetrieverContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewRetrieverContainer returns a RetrieverContainer.
func NewRetrieverContainer() *RetrieverContainer {
	r := &RetrieverContainer{
		AdditionalRetrieverContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}

// true

type RetrieverContainerVariant interface {
	RetrieverContainerCaster() *RetrieverContainer
}

func (s *RetrieverContainer) RetrieverContainerCaster() *RetrieverContainer {
	return s
}

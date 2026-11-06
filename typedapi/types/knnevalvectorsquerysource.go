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
)

// Uses the query vectors in the request.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/knn_eval/types.ts#L48-L57
type KnnEvalVectorsQuerySource struct {
	From string `json:"from,omitempty"`
	// Vectors The query vectors, from 1 to 1000 entries with unique identifiers.
	Vectors []KnnEvalQueryVector `json:"vectors"`
}

// MarshalJSON override marshalling to include literal value
func (s KnnEvalVectorsQuerySource) MarshalJSON() ([]byte, error) {
	type innerKnnEvalVectorsQuerySource KnnEvalVectorsQuerySource
	tmp := innerKnnEvalVectorsQuerySource{
		From:    s.From,
		Vectors: s.Vectors,
	}

	tmp.From = "vectors"

	return json.Marshal(tmp)
}

// NewKnnEvalVectorsQuerySource returns a KnnEvalVectorsQuerySource.
func NewKnnEvalVectorsQuerySource() *KnnEvalVectorsQuerySource {
	r := &KnnEvalVectorsQuerySource{}

	return r
}

type KnnEvalVectorsQuerySourceVariant interface {
	KnnEvalVectorsQuerySourceCaster() *KnnEvalVectorsQuerySource
}

func (s *KnnEvalVectorsQuerySource) KnnEvalVectorsQuerySourceCaster() *KnnEvalVectorsQuerySource {
	return s
}

func (s *KnnEvalVectorsQuerySource) KnnEvalQuerySourceCaster() *KnnEvalQuerySource {
	if s == nil {
		return nil
	}
	o := KnnEvalQuerySource(s)
	return &o
}

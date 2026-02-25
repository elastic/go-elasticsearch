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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package densevectorsimilarity
package densevectorsimilarity

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/mapping/DenseVectorProperty.ts#L87-L132
type DenseVectorSimilarity struct {
	Name string
}

var (

	// Cosine Computes the cosine similarity. During indexing Elasticsearch automatically
	// normalizes vectors with `cosine` similarity to unit length. This allows to
	// internally use `dot_product` for computing similarity, which is more
	// efficient. Original un-normalized vectors can be still accessed through
	// scripts.
	//
	// The document `_score` is computed as `(1 + cosine(query, vector)) / 2`.
	//
	// The `cosine` similarity does not allow vectors with zero magnitude, since
	// cosine is not defined in this case.
	Cosine = DenseVectorSimilarity{"cosine"}

	// Dotproduct Computes the dot product of two unit vectors. This option provides an
	// optimized way to perform cosine similarity. The constraints and computed
	// score are defined by `element_type`.
	//
	// When `element_type` is `float`, all vectors must be unit length, including
	// both document and query vectors.
	//
	// The document `_score` is computed as `(1 + dot_product(query, vector)) / 2`.
	//
	// When `element_type` is `byte`, all vectors must have the same length
	// including both document and query vectors or results will be inaccurate.
	//
	// The document `_score` is computed as `0.5 + (dot_product(query, vector) /
	// (32768 * dims))` where `dims` is the number of dimensions per vector.
	Dotproduct = DenseVectorSimilarity{"dot_product"}

	// L2norm Computes similarity based on the `L2` distance (also known as Euclidean
	// distance) between the vectors.
	//
	// The document `_score` is computed as `1 / (1 + l2_norm(query, vector)^2)`.
	//
	// For `bit` vectors, instead of using `l2_norm`, the `hamming` distance between
	// the vectors is used.
	//
	// The `_score` transformation is `(numBits - hamming(a, b)) / numBits`.
	L2norm = DenseVectorSimilarity{"l2_norm"}

	// Maxinnerproduct Computes the maximum inner product of two vectors. This is similar to
	// `dot_product`, but doesn't require vectors to be normalized. This means that
	// each vector’s magnitude can significantly effect the score.
	//
	// The document `_score` is adjusted to prevent negative values. For
	// `max_inner_product` values `< 0`, the `_score` is `1 / (1 + -1 *
	// max_inner_product(query, vector))`. For non-negative `max_inner_product`
	// results the `_score` is calculated `max_inner_product(query, vector) + 1`.
	Maxinnerproduct = DenseVectorSimilarity{"max_inner_product"}
)

func (d DenseVectorSimilarity) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DenseVectorSimilarity) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "cosine":
		*d = Cosine
	case "dot_product":
		*d = Dotproduct
	case "l2_norm":
		*d = L2norm
	case "max_inner_product":
		*d = Maxinnerproduct
	default:
		*d = DenseVectorSimilarity{string(text)}
	}

	return nil
}

func (d DenseVectorSimilarity) String() string {
	return d.Name
}

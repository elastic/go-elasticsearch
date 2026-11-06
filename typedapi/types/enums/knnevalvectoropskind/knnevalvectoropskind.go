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

// Package knnevalvectoropskind
package knnevalvectoropskind

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/knn_eval/types.ts#L118-L127
type KnnEvalVectorOpsKind struct {
	Name string
}

var (

	// Fullprecisionscan Full-precision comparisons across all the matching vectors.
	Fullprecisionscan = KnnEvalVectorOpsKind{"full_precision_scan"}

	// Quantizedvisitplusrescore Quantized visits plus full-precision rescoring.
	Quantizedvisitplusrescore = KnnEvalVectorOpsKind{"quantized_visit_plus_rescore"}
)

func (k KnnEvalVectorOpsKind) MarshalText() (text []byte, err error) {
	return []byte(k.String()), nil
}

func (k *KnnEvalVectorOpsKind) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "full_precision_scan":
		*k = Fullprecisionscan
	case "quantized_visit_plus_rescore":
		*k = Quantizedvisitplusrescore
	default:
		*k = KnnEvalVectorOpsKind{string(text)}
	}

	return nil
}

func (k KnnEvalVectorOpsKind) String() string {
	return k.Name
}

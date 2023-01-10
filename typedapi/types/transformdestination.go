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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// TransformDestination type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/transform/_types/Transform.ts#L34-L45
type TransformDestination struct {
	// Index The destination index for the transform. The mappings of the destination
	// index are deduced based on the source
	// fields when possible. If alternate mappings are required, use the create
	// index API prior to starting the
	// transform.
	Index *string `json:"index,omitempty"`
	// Pipeline The unique identifier for an ingest pipeline.
	Pipeline *string `json:"pipeline,omitempty"`
}

// NewTransformDestination returns a TransformDestination.
func NewTransformDestination() *TransformDestination {
	r := &TransformDestination{}

	return r
}

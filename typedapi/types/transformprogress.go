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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// TransformProgress type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/transform/get_transform_stats/types.ts#L40-L46
type TransformProgress struct {
	DocsIndexed     int64   `json:"docs_indexed"`
	DocsProcessed   int64   `json:"docs_processed"`
	DocsRemaining   int64   `json:"docs_remaining"`
	PercentComplete float64 `json:"percent_complete"`
	TotalDocs       int64   `json:"total_docs"`
}

// NewTransformProgress returns a TransformProgress.
func NewTransformProgress() *TransformProgress {
	r := &TransformProgress{}

	return r
}

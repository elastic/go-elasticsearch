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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package types

// Allows specifying one or multiple items for the `embedding` task, which
// should result in a single embedding vector. Support for multiple items is
// available in Elasticsearch 9.5.0 and later.
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/inference/_types/CommonTypes.ts#L641-L647
type EmbeddingContentObjectGroup []EmbeddingContentObjectItem

type EmbeddingContentObjectGroupVariant interface {
	EmbeddingContentObjectGroupCaster() *EmbeddingContentObjectGroup
}

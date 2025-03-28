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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _alibabaCloudTaskSettings struct {
	v *types.AlibabaCloudTaskSettings
}

func NewAlibabaCloudTaskSettings() *_alibabaCloudTaskSettings {

	return &_alibabaCloudTaskSettings{v: types.NewAlibabaCloudTaskSettings()}

}

// For a `sparse_embedding` or `text_embedding` task, specify the type of input
// passed to the model.
// Valid values are:
//
// * `ingest` for storing document embeddings in a vector database.
// * `search` for storing embeddings of search queries run against a vector
// database to find relevant documents.
func (s *_alibabaCloudTaskSettings) InputType(inputtype string) *_alibabaCloudTaskSettings {

	s.v.InputType = &inputtype

	return s
}

// For a `sparse_embedding` task, it affects whether the token name will be
// returned in the response.
// It defaults to `false`, which means only the token ID will be returned in the
// response.
func (s *_alibabaCloudTaskSettings) ReturnToken(returntoken bool) *_alibabaCloudTaskSettings {

	s.v.ReturnToken = &returntoken

	return s
}

func (s *_alibabaCloudTaskSettings) AlibabaCloudTaskSettingsCaster() *types.AlibabaCloudTaskSettings {
	return s.v
}

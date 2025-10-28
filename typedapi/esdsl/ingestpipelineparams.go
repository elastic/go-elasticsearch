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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ingestPipelineParams struct {
	v *types.IngestPipelineParams
}

func NewIngestPipelineParams(extractbinarycontent bool, name string, reducewhitespace bool, runmlinference bool) *_ingestPipelineParams {

	tmp := &_ingestPipelineParams{v: types.NewIngestPipelineParams()}

	tmp.ExtractBinaryContent(extractbinarycontent)

	tmp.Name(name)

	tmp.ReduceWhitespace(reducewhitespace)

	tmp.RunMlInference(runmlinference)

	return tmp

}

func (s *_ingestPipelineParams) ExtractBinaryContent(extractbinarycontent bool) *_ingestPipelineParams {

	s.v.ExtractBinaryContent = extractbinarycontent

	return s
}

func (s *_ingestPipelineParams) Name(name string) *_ingestPipelineParams {

	s.v.Name = name

	return s
}

func (s *_ingestPipelineParams) ReduceWhitespace(reducewhitespace bool) *_ingestPipelineParams {

	s.v.ReduceWhitespace = reducewhitespace

	return s
}

func (s *_ingestPipelineParams) RunMlInference(runmlinference bool) *_ingestPipelineParams {

	s.v.RunMlInference = runmlinference

	return s
}

func (s *_ingestPipelineParams) IngestPipelineParamsCaster() *types.IngestPipelineParams {
	return s.v
}

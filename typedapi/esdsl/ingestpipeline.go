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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ingestPipeline struct {
	v *types.IngestPipeline
}

func NewIngestPipeline() *_ingestPipeline {

	return &_ingestPipeline{v: types.NewIngestPipeline()}

}

func (s *_ingestPipeline) Deprecated(deprecated bool) *_ingestPipeline {

	s.v.Deprecated = &deprecated

	return s
}

func (s *_ingestPipeline) Description(description string) *_ingestPipeline {

	s.v.Description = &description

	return s
}

func (s *_ingestPipeline) Meta_(metadata types.MetadataVariant) *_ingestPipeline {

	s.v.Meta_ = *metadata.MetadataCaster()

	return s
}

func (s *_ingestPipeline) OnFailure(onfailures ...types.ProcessorContainerVariant) *_ingestPipeline {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_ingestPipeline) Processors(processors ...types.ProcessorContainerVariant) *_ingestPipeline {

	for _, v := range processors {

		s.v.Processors = append(s.v.Processors, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_ingestPipeline) Version(versionnumber int64) *_ingestPipeline {

	s.v.Version = &versionnumber

	return s
}

func (s *_ingestPipeline) IngestPipelineCaster() *types.IngestPipeline {
	return s.v
}

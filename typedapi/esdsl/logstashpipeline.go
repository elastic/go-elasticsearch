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

type _logstashPipeline struct {
	v *types.LogstashPipeline
}

func NewLogstashPipeline(description string, pipeline string, pipelinemetadata types.PipelineMetadataVariant, pipelinesettings types.PipelineSettingsVariant, username string) *_logstashPipeline {

	tmp := &_logstashPipeline{v: types.NewLogstashPipeline()}

	tmp.Description(description)

	tmp.Pipeline(pipeline)

	tmp.PipelineMetadata(pipelinemetadata)

	tmp.PipelineSettings(pipelinesettings)

	tmp.Username(username)

	return tmp

}

func (s *_logstashPipeline) Description(description string) *_logstashPipeline {

	s.v.Description = description

	return s
}

func (s *_logstashPipeline) LastModified(datetime types.DateTimeVariant) *_logstashPipeline {

	s.v.LastModified = *datetime.DateTimeCaster()

	return s
}

func (s *_logstashPipeline) Pipeline(pipeline string) *_logstashPipeline {

	s.v.Pipeline = pipeline

	return s
}

func (s *_logstashPipeline) PipelineMetadata(pipelinemetadata types.PipelineMetadataVariant) *_logstashPipeline {

	s.v.PipelineMetadata = *pipelinemetadata.PipelineMetadataCaster()

	return s
}

func (s *_logstashPipeline) PipelineSettings(pipelinesettings types.PipelineSettingsVariant) *_logstashPipeline {

	s.v.PipelineSettings = *pipelinesettings.PipelineSettingsCaster()

	return s
}

func (s *_logstashPipeline) Username(username string) *_logstashPipeline {

	s.v.Username = username

	return s
}

func (s *_logstashPipeline) LogstashPipelineCaster() *types.LogstashPipeline {
	return s.v
}

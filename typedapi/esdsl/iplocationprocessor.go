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

type _ipLocationProcessor struct {
	v *types.IpLocationProcessor
}

// Currently an undocumented alias for GeoIP Processor.
func NewIpLocationProcessor() *_ipLocationProcessor {

	return &_ipLocationProcessor{v: types.NewIpLocationProcessor()}

}

func (s *_ipLocationProcessor) DatabaseFile(databasefile string) *_ipLocationProcessor {

	s.v.DatabaseFile = &databasefile

	return s
}

func (s *_ipLocationProcessor) DownloadDatabaseOnPipelineCreation(downloaddatabaseonpipelinecreation bool) *_ipLocationProcessor {

	s.v.DownloadDatabaseOnPipelineCreation = &downloaddatabaseonpipelinecreation

	return s
}

func (s *_ipLocationProcessor) Field(field string) *_ipLocationProcessor {

	s.v.Field = field

	return s
}

func (s *_ipLocationProcessor) FirstOnly(firstonly bool) *_ipLocationProcessor {

	s.v.FirstOnly = &firstonly

	return s
}

func (s *_ipLocationProcessor) IgnoreMissing(ignoremissing bool) *_ipLocationProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_ipLocationProcessor) Properties(properties ...string) *_ipLocationProcessor {

	for _, v := range properties {

		s.v.Properties = append(s.v.Properties, v)

	}
	return s
}

func (s *_ipLocationProcessor) TargetField(field string) *_ipLocationProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_ipLocationProcessor) Description(description string) *_ipLocationProcessor {

	s.v.Description = &description

	return s
}

func (s *_ipLocationProcessor) If(if_ types.ScriptVariant) *_ipLocationProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_ipLocationProcessor) IgnoreFailure(ignorefailure bool) *_ipLocationProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_ipLocationProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_ipLocationProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_ipLocationProcessor) Tag(tag string) *_ipLocationProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_ipLocationProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.IpLocation = s.v

	return container
}

func (s *_ipLocationProcessor) IpLocationProcessorCaster() *types.IpLocationProcessor {
	return s.v
}

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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _geoIpProcessor struct {
	v *types.GeoIpProcessor
}

// The `geoip` processor adds information about the geographical location of an
// IPv4 or IPv6 address.
func NewGeoIpProcessor() *_geoIpProcessor {

	return &_geoIpProcessor{v: types.NewGeoIpProcessor()}

}

func (s *_geoIpProcessor) DatabaseFile(databasefile string) *_geoIpProcessor {

	s.v.DatabaseFile = &databasefile

	return s
}

func (s *_geoIpProcessor) Description(description string) *_geoIpProcessor {

	s.v.Description = &description

	return s
}

func (s *_geoIpProcessor) DownloadDatabaseOnPipelineCreation(downloaddatabaseonpipelinecreation bool) *_geoIpProcessor {

	s.v.DownloadDatabaseOnPipelineCreation = &downloaddatabaseonpipelinecreation

	return s
}

func (s *_geoIpProcessor) Field(field string) *_geoIpProcessor {

	s.v.Field = field

	return s
}

func (s *_geoIpProcessor) FirstOnly(firstonly bool) *_geoIpProcessor {

	s.v.FirstOnly = &firstonly

	return s
}

func (s *_geoIpProcessor) If(if_ types.ScriptVariant) *_geoIpProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_geoIpProcessor) IgnoreFailure(ignorefailure bool) *_geoIpProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_geoIpProcessor) IgnoreMissing(ignoremissing bool) *_geoIpProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_geoIpProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_geoIpProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_geoIpProcessor) Properties(properties ...string) *_geoIpProcessor {

	for _, v := range properties {

		s.v.Properties = append(s.v.Properties, v)

	}
	return s
}

func (s *_geoIpProcessor) Tag(tag string) *_geoIpProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_geoIpProcessor) TargetField(field string) *_geoIpProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_geoIpProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Geoip = s.v

	return container
}

func (s *_geoIpProcessor) GeoIpProcessorCaster() *types.GeoIpProcessor {
	return s.v
}

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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _ipLocationProcessor struct {
	v *types.IpLocationProcessor
}

// Currently an undocumented alias for GeoIP Processor.
func NewIpLocationProcessor() *_ipLocationProcessor {

	return &_ipLocationProcessor{v: types.NewIpLocationProcessor()}

}

// The database filename referring to a database the module ships with
// (GeoLite2-City.mmdb, GeoLite2-Country.mmdb, or GeoLite2-ASN.mmdb) or a custom
// database in the ingest-geoip config directory.
func (s *_ipLocationProcessor) DatabaseFile(databasefile string) *_ipLocationProcessor {

	s.v.DatabaseFile = &databasefile

	return s
}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_ipLocationProcessor) Description(description string) *_ipLocationProcessor {

	s.v.Description = &description

	return s
}

// If `true` (and if `ingest.geoip.downloader.eager.download` is `false`), the
// missing database is downloaded when the pipeline is created.
// Else, the download is triggered by when the pipeline is used as the
// `default_pipeline` or `final_pipeline` in an index.
func (s *_ipLocationProcessor) DownloadDatabaseOnPipelineCreation(downloaddatabaseonpipelinecreation bool) *_ipLocationProcessor {

	s.v.DownloadDatabaseOnPipelineCreation = &downloaddatabaseonpipelinecreation

	return s
}

// The field to get the ip address from for the geographical lookup.
func (s *_ipLocationProcessor) Field(field string) *_ipLocationProcessor {

	s.v.Field = field

	return s
}

// If `true`, only the first found IP location data will be returned, even if
// the field contains an array.
func (s *_ipLocationProcessor) FirstOnly(firstonly bool) *_ipLocationProcessor {

	s.v.FirstOnly = &firstonly

	return s
}

// Conditionally execute the processor.
func (s *_ipLocationProcessor) If(if_ types.ScriptVariant) *_ipLocationProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_ipLocationProcessor) IgnoreFailure(ignorefailure bool) *_ipLocationProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist, the processor quietly exits without
// modifying the document.
func (s *_ipLocationProcessor) IgnoreMissing(ignoremissing bool) *_ipLocationProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_ipLocationProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_ipLocationProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Controls what properties are added to the `target_field` based on the IP
// location lookup.
func (s *_ipLocationProcessor) Properties(properties ...string) *_ipLocationProcessor {

	for _, v := range properties {

		s.v.Properties = append(s.v.Properties, v)

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_ipLocationProcessor) Tag(tag string) *_ipLocationProcessor {

	s.v.Tag = &tag

	return s
}

// The field that will hold the geographical information looked up from the
// MaxMind database.
func (s *_ipLocationProcessor) TargetField(field string) *_ipLocationProcessor {

	s.v.TargetField = &field

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

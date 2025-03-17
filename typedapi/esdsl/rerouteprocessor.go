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

type _rerouteProcessor struct {
	v *types.RerouteProcessor
}

// Routes a document to another target index or data stream.
// When setting the `destination` option, the target is explicitly specified and
// the dataset and namespace options can’t be set.
// When the `destination` option is not set, this processor is in a data stream
// mode. Note that in this mode, the reroute processor can only be used on data
// streams that follow the data stream naming scheme.
func NewRerouteProcessor() *_rerouteProcessor {

	return &_rerouteProcessor{v: types.NewRerouteProcessor()}

}

// Field references or a static value for the dataset part of the data stream
// name.
// In addition to the criteria for index names, cannot contain - and must be no
// longer than 100 characters.
// Example values are nginx.access and nginx.error.
//
// Supports field references with a mustache-like syntax (denoted as {{double}}
// or {{{triple}}} curly braces).
// When resolving field references, the processor replaces invalid characters
// with _. Uses the <dataset> part
// of the index name as a fallback if all field references resolve to a null,
// missing, or non-string value.
//
// default {{data_stream.dataset}}
func (s *_rerouteProcessor) Dataset(datasets ...string) *_rerouteProcessor {

	s.v.Dataset = make([]string, len(datasets))
	s.v.Dataset = datasets

	return s
}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_rerouteProcessor) Description(description string) *_rerouteProcessor {

	s.v.Description = &description

	return s
}

// A static value for the target. Can’t be set when the dataset or namespace
// option is set.
func (s *_rerouteProcessor) Destination(destination string) *_rerouteProcessor {

	s.v.Destination = &destination

	return s
}

// Conditionally execute the processor.
func (s *_rerouteProcessor) If(if_ types.ScriptVariant) *_rerouteProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_rerouteProcessor) IgnoreFailure(ignorefailure bool) *_rerouteProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// Field references or a static value for the namespace part of the data stream
// name. See the criteria for
// index names for allowed characters. Must be no longer than 100 characters.
//
// Supports field references with a mustache-like syntax (denoted as {{double}}
// or {{{triple}}} curly braces).
// When resolving field references, the processor replaces invalid characters
// with _. Uses the <namespace> part
// of the index name as a fallback if all field references resolve to a null,
// missing, or non-string value.
//
// default {{data_stream.namespace}}
func (s *_rerouteProcessor) Namespace(namespaces ...string) *_rerouteProcessor {

	s.v.Namespace = make([]string, len(namespaces))
	s.v.Namespace = namespaces

	return s
}

// Handle failures for the processor.
func (s *_rerouteProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_rerouteProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_rerouteProcessor) Tag(tag string) *_rerouteProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_rerouteProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Reroute = s.v

	return container
}

func (s *_rerouteProcessor) RerouteProcessorCaster() *types.RerouteProcessor {
	return s.v
}

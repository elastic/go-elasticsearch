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

type _uriPartsProcessor struct {
	v *types.UriPartsProcessor
}

// Parses a Uniform Resource Identifier (URI) string and extracts its components
// as an object.
// This URI object includes properties for the URI’s domain, path, fragment,
// port, query, scheme, user info, username, and password.
func NewUriPartsProcessor() *_uriPartsProcessor {

	return &_uriPartsProcessor{v: types.NewUriPartsProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_uriPartsProcessor) Description(description string) *_uriPartsProcessor {

	s.v.Description = &description

	return s
}

// Field containing the URI string.
func (s *_uriPartsProcessor) Field(field string) *_uriPartsProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_uriPartsProcessor) If(if_ types.ScriptVariant) *_uriPartsProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_uriPartsProcessor) IgnoreFailure(ignorefailure bool) *_uriPartsProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist, the processor quietly exits without
// modifying the document.
func (s *_uriPartsProcessor) IgnoreMissing(ignoremissing bool) *_uriPartsProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// If `true`, the processor copies the unparsed URI to
// `<target_field>.original`.
func (s *_uriPartsProcessor) KeepOriginal(keeporiginal bool) *_uriPartsProcessor {

	s.v.KeepOriginal = &keeporiginal

	return s
}

// Handle failures for the processor.
func (s *_uriPartsProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_uriPartsProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// If `true`, the processor removes the `field` after parsing the URI string.
// If parsing fails, the processor does not remove the `field`.
func (s *_uriPartsProcessor) RemoveIfSuccessful(removeifsuccessful bool) *_uriPartsProcessor {

	s.v.RemoveIfSuccessful = &removeifsuccessful

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_uriPartsProcessor) Tag(tag string) *_uriPartsProcessor {

	s.v.Tag = &tag

	return s
}

// Output field for the URI object.
func (s *_uriPartsProcessor) TargetField(field string) *_uriPartsProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_uriPartsProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.UriParts = s.v

	return container
}

func (s *_uriPartsProcessor) UriPartsProcessorCaster() *types.UriPartsProcessor {
	return s.v
}

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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fingerprintdigest"
)

type _fingerprintProcessor struct {
	v *types.FingerprintProcessor
}

// Computes a hash of the document’s content. You can use this hash for
// content fingerprinting.
func NewFingerprintProcessor() *_fingerprintProcessor {

	return &_fingerprintProcessor{v: types.NewFingerprintProcessor()}

}

func (s *_fingerprintProcessor) Fields(fields ...string) *_fingerprintProcessor {

	s.v.Fields = fields

	return s
}

func (s *_fingerprintProcessor) IgnoreMissing(ignoremissing bool) *_fingerprintProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_fingerprintProcessor) Method(method fingerprintdigest.FingerprintDigest) *_fingerprintProcessor {

	s.v.Method = &method
	return s
}

func (s *_fingerprintProcessor) Salt(salt string) *_fingerprintProcessor {

	s.v.Salt = &salt

	return s
}

func (s *_fingerprintProcessor) TargetField(field string) *_fingerprintProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_fingerprintProcessor) Description(description string) *_fingerprintProcessor {

	s.v.Description = &description

	return s
}

func (s *_fingerprintProcessor) If(if_ types.ScriptVariant) *_fingerprintProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_fingerprintProcessor) IgnoreFailure(ignorefailure bool) *_fingerprintProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_fingerprintProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_fingerprintProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_fingerprintProcessor) Tag(tag string) *_fingerprintProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_fingerprintProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Fingerprint = s.v

	return container
}

func (s *_fingerprintProcessor) FingerprintProcessorCaster() *types.FingerprintProcessor {
	return s.v
}

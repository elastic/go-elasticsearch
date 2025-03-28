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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fingerprintdigest"
)

type _fingerprintProcessor struct {
	v *types.FingerprintProcessor
}

// Computes a hash of the document’s content. You can use this hash for
// content fingerprinting.
func NewFingerprintProcessor() *_fingerprintProcessor {

	return &_fingerprintProcessor{v: types.NewFingerprintProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_fingerprintProcessor) Description(description string) *_fingerprintProcessor {

	s.v.Description = &description

	return s
}

// Array of fields to include in the fingerprint. For objects, the processor
// hashes both the field key and value. For other fields, the processor hashes
// only the field value.
func (s *_fingerprintProcessor) Fields(fields ...string) *_fingerprintProcessor {

	s.v.Fields = fields

	return s
}

// Conditionally execute the processor.
func (s *_fingerprintProcessor) If(if_ types.ScriptVariant) *_fingerprintProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_fingerprintProcessor) IgnoreFailure(ignorefailure bool) *_fingerprintProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If true, the processor ignores any missing fields. If all fields are
// missing, the processor silently exits without modifying the document.
func (s *_fingerprintProcessor) IgnoreMissing(ignoremissing bool) *_fingerprintProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// The hash method used to compute the fingerprint. Must be one of MD5, SHA-1,
// SHA-256, SHA-512, or MurmurHash3.
func (s *_fingerprintProcessor) Method(method fingerprintdigest.FingerprintDigest) *_fingerprintProcessor {

	s.v.Method = &method
	return s
}

// Handle failures for the processor.
func (s *_fingerprintProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_fingerprintProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Salt value for the hash function.
func (s *_fingerprintProcessor) Salt(salt string) *_fingerprintProcessor {

	s.v.Salt = &salt

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_fingerprintProcessor) Tag(tag string) *_fingerprintProcessor {

	s.v.Tag = &tag

	return s
}

// Output field for the fingerprint.
func (s *_fingerprintProcessor) TargetField(field string) *_fingerprintProcessor {

	s.v.TargetField = &field

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

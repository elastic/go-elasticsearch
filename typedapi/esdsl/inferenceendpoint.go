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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _inferenceEndpoint struct {
	v *types.InferenceEndpoint
}

func NewInferenceEndpoint(service string) *_inferenceEndpoint {

	tmp := &_inferenceEndpoint{v: types.NewInferenceEndpoint()}

	tmp.Service(service)

	return tmp

}

// Chunking configuration object
func (s *_inferenceEndpoint) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *_inferenceEndpoint {

	s.v.ChunkingSettings = chunkingsettings.InferenceChunkingSettingsCaster()

	return s
}

// The service type
func (s *_inferenceEndpoint) Service(service string) *_inferenceEndpoint {

	s.v.Service = service

	return s
}

// Settings specific to the service
func (s *_inferenceEndpoint) ServiceSettings(servicesettings json.RawMessage) *_inferenceEndpoint {

	s.v.ServiceSettings = servicesettings

	return s
}

// Task settings specific to the service and task type
func (s *_inferenceEndpoint) TaskSettings(tasksettings json.RawMessage) *_inferenceEndpoint {

	s.v.TaskSettings = tasksettings

	return s
}

func (s *_inferenceEndpoint) InferenceEndpointCaster() *types.InferenceEndpoint {
	return s.v
}

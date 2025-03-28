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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _elasticsearchServiceSettings struct {
	v *types.ElasticsearchServiceSettings
}

func NewElasticsearchServiceSettings(modelid string, numthreads int) *_elasticsearchServiceSettings {

	tmp := &_elasticsearchServiceSettings{v: types.NewElasticsearchServiceSettings()}

	tmp.ModelId(modelid)

	tmp.NumThreads(numthreads)

	return tmp

}

// Adaptive allocations configuration details.
// If `enabled` is true, the number of allocations of the model is set based on
// the current load the process gets.
// When the load is high, a new model allocation is automatically created,
// respecting the value of `max_number_of_allocations` if it's set.
// When the load is low, a model allocation is automatically removed, respecting
// the value of `min_number_of_allocations` if it's set.
// If `enabled` is true, do not set the number of allocations manually.
func (s *_elasticsearchServiceSettings) AdaptiveAllocations(adaptiveallocations types.AdaptiveAllocationsVariant) *_elasticsearchServiceSettings {

	s.v.AdaptiveAllocations = adaptiveallocations.AdaptiveAllocationsCaster()

	return s
}

// The deployment identifier for a trained model deployment.
// When `deployment_id` is used the `model_id` is optional.
func (s *_elasticsearchServiceSettings) DeploymentId(deploymentid string) *_elasticsearchServiceSettings {

	s.v.DeploymentId = &deploymentid

	return s
}

// The name of the model to use for the inference task.
// It can be the ID of a built-in model (for example, `.multilingual-e5-small`
// for E5) or a text embedding model that was uploaded by using the Eland
// client.
func (s *_elasticsearchServiceSettings) ModelId(modelid string) *_elasticsearchServiceSettings {

	s.v.ModelId = modelid

	return s
}

// The total number of allocations that are assigned to the model across machine
// learning nodes.
// Increasing this value generally increases the throughput.
// If adaptive allocations are enabled, do not set this value because it's
// automatically set.
func (s *_elasticsearchServiceSettings) NumAllocations(numallocations int) *_elasticsearchServiceSettings {

	s.v.NumAllocations = &numallocations

	return s
}

// The number of threads used by each model allocation during inference.
// This setting generally increases the speed per inference request.
// The inference process is a compute-bound process; `threads_per_allocations`
// must not exceed the number of available allocated processors per node.
// The value must be a power of 2.
// The maximum value is 32.
func (s *_elasticsearchServiceSettings) NumThreads(numthreads int) *_elasticsearchServiceSettings {

	s.v.NumThreads = numthreads

	return s
}

func (s *_elasticsearchServiceSettings) ElasticsearchServiceSettingsCaster() *types.ElasticsearchServiceSettings {
	return s.v
}

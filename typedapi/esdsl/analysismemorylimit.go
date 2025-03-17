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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _analysisMemoryLimit struct {
	v *types.AnalysisMemoryLimit
}

func NewAnalysisMemoryLimit(modelmemorylimit string) *_analysisMemoryLimit {

	tmp := &_analysisMemoryLimit{v: types.NewAnalysisMemoryLimit()}

	tmp.ModelMemoryLimit(modelmemorylimit)

	return tmp

}

// Limits can be applied for the resources required to hold the mathematical
// models in memory. These limits are approximate and can be set per job. They
// do not control the memory used by other processes, for example the
// Elasticsearch Java processes.
func (s *_analysisMemoryLimit) ModelMemoryLimit(modelmemorylimit string) *_analysisMemoryLimit {

	s.v.ModelMemoryLimit = modelmemorylimit

	return s
}

func (s *_analysisMemoryLimit) AnalysisMemoryLimitCaster() *types.AnalysisMemoryLimit {
	return s.v
}

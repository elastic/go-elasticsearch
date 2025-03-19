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

type _tokenPruningConfig struct {
	v *types.TokenPruningConfig
}

// Optional pruning configuration.
// If enabled, this will omit non-significant tokens from the query in order to
// improve query performance.
// This is only used if prune is set to true.
// If prune is set to true but pruning_config is not specified, default values
// will be used.
func NewTokenPruningConfig() *_tokenPruningConfig {

	return &_tokenPruningConfig{v: types.NewTokenPruningConfig()}

}

// Whether to only score pruned tokens, vs only scoring kept tokens.
func (s *_tokenPruningConfig) OnlyScorePrunedTokens(onlyscoreprunedtokens bool) *_tokenPruningConfig {

	s.v.OnlyScorePrunedTokens = &onlyscoreprunedtokens

	return s
}

// Tokens whose frequency is more than this threshold times the average
// frequency of all tokens in the specified field are considered outliers and
// pruned.
func (s *_tokenPruningConfig) TokensFreqRatioThreshold(tokensfreqratiothreshold int) *_tokenPruningConfig {

	s.v.TokensFreqRatioThreshold = &tokensfreqratiothreshold

	return s
}

// Tokens whose weight is less than this threshold are considered nonsignificant
// and pruned.
func (s *_tokenPruningConfig) TokensWeightThreshold(tokensweightthreshold float32) *_tokenPruningConfig {

	s.v.TokensWeightThreshold = &tokensweightthreshold

	return s
}

func (s *_tokenPruningConfig) SparseVectorQueryCaster() *types.SparseVectorQuery {
	container := types.NewSparseVectorQuery()

	container.PruningConfig = s.v

	return container
}

func (s *_tokenPruningConfig) TokenPruningConfigCaster() *types.TokenPruningConfig {
	return s.v
}

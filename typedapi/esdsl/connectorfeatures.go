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

type _connectorFeatures struct {
	v *types.ConnectorFeatures
}

func NewConnectorFeatures() *_connectorFeatures {

	return &_connectorFeatures{v: types.NewConnectorFeatures()}

}

// Indicates whether document-level security is enabled.
func (s *_connectorFeatures) DocumentLevelSecurity(documentlevelsecurity types.FeatureEnabledVariant) *_connectorFeatures {

	s.v.DocumentLevelSecurity = documentlevelsecurity.FeatureEnabledCaster()

	return s
}

// Indicates whether incremental syncs are enabled.
func (s *_connectorFeatures) IncrementalSync(incrementalsync types.FeatureEnabledVariant) *_connectorFeatures {

	s.v.IncrementalSync = incrementalsync.FeatureEnabledCaster()

	return s
}

// Indicates whether managed connector API keys are enabled.
func (s *_connectorFeatures) NativeConnectorApiKeys(nativeconnectorapikeys types.FeatureEnabledVariant) *_connectorFeatures {

	s.v.NativeConnectorApiKeys = nativeconnectorapikeys.FeatureEnabledCaster()

	return s
}

func (s *_connectorFeatures) SyncRules(syncrules types.SyncRulesFeatureVariant) *_connectorFeatures {

	s.v.SyncRules = syncrules.SyncRulesFeatureCaster()

	return s
}

func (s *_connectorFeatures) ConnectorFeaturesCaster() *types.ConnectorFeatures {
	return s.v
}

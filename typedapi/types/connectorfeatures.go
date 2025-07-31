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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

// ConnectorFeatures type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/connector/_types/Connector.ts#L230-L244
type ConnectorFeatures struct {
	// DocumentLevelSecurity Indicates whether document-level security is enabled.
	DocumentLevelSecurity *FeatureEnabled `json:"document_level_security,omitempty"`
	// IncrementalSync Indicates whether incremental syncs are enabled.
	IncrementalSync *FeatureEnabled `json:"incremental_sync,omitempty"`
	// NativeConnectorApiKeys Indicates whether managed connector API keys are enabled.
	NativeConnectorApiKeys *FeatureEnabled   `json:"native_connector_api_keys,omitempty"`
	SyncRules              *SyncRulesFeature `json:"sync_rules,omitempty"`
}

// NewConnectorFeatures returns a ConnectorFeatures.
func NewConnectorFeatures() *ConnectorFeatures {
	r := &ConnectorFeatures{}

	return r
}

type ConnectorFeaturesVariant interface {
	ConnectorFeaturesCaster() *ConnectorFeatures
}

func (s *ConnectorFeatures) ConnectorFeaturesCaster() *ConnectorFeatures {
	return s
}

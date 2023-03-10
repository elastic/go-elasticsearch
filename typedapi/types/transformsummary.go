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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"encoding/json"
)

// TransformSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/transform/get_transform/types.ts#L33-L61
type TransformSummary struct {
	// Authorization The security privileges that the transform uses to run its queries. If
	// Elastic Stack security features were disabled at the time of the most recent
	// update to the transform, this property is omitted.
	Authorization *TransformAuthorization `json:"authorization,omitempty"`
	// CreateTime The time the transform was created.
	CreateTime *int64 `json:"create_time,omitempty"`
	// Description Free text description of the transform.
	Description *string `json:"description,omitempty"`
	// Dest The destination for the transform.
	Dest      ReindexDestination         `json:"dest"`
	Frequency Duration                   `json:"frequency,omitempty"`
	Id        string                     `json:"id"`
	Latest    *Latest                    `json:"latest,omitempty"`
	Meta_     map[string]json.RawMessage `json:"_meta,omitempty"`
	// Pivot The pivot method transforms the data by aggregating and grouping it.
	Pivot           *Pivot                    `json:"pivot,omitempty"`
	RetentionPolicy *RetentionPolicyContainer `json:"retention_policy,omitempty"`
	// Settings Defines optional transform settings.
	Settings *Settings `json:"settings,omitempty"`
	// Source The source of the data for the transform.
	Source TransformSource `json:"source"`
	// Sync Defines the properties transforms require to run continuously.
	Sync *SyncContainer `json:"sync,omitempty"`
	// Version The version of Elasticsearch that existed on the node when the transform was
	// created.
	Version *string `json:"version,omitempty"`
}

// NewTransformSummary returns a TransformSummary.
func NewTransformSummary() *TransformSummary {
	r := &TransformSummary{}

	return r
}

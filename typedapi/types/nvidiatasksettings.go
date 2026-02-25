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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheretruncatetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/nvidiainputtype"
)

// NvidiaTaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/inference/_types/CommonTypes.ts#L2027-L2047
type NvidiaTaskSettings struct {
	// InputType For a `text_embedding` task, type of input sent to the Nvidia endpoint. Valid
	// values are:
	//
	//   - `ingest`: Mapped to Nvidia's `passage` value in request. Used when
	//     generating embeddings during indexing.
	//   - `search`: Mapped to Nvidia's `query` value in request. Used when
	//     generating embeddings during querying.
	//
	// IMPORTANT: For Nvidia endpoints, if the `input_type` field is not specified,
	// it defaults to `query`.
	InputType *nvidiainputtype.NvidiaInputType `json:"input_type,omitempty"`
	// Truncate For a `text_embedding` task, the method used by the Nvidia model to handle
	// inputs longer than the maximum token length. Valid values are:
	//
	//   - `END`: When the input exceeds the maximum input token length, the end of
	//     the input is discarded.
	//   - `NONE`: When the input exceeds the maximum input token length, an error
	//     is returned.
	//   - `START`: When the input exceeds the maximum input token length, the start
	//     of the input is discarded.
	Truncate *coheretruncatetype.CohereTruncateType `json:"truncate,omitempty"`
}

// NewNvidiaTaskSettings returns a NvidiaTaskSettings.
func NewNvidiaTaskSettings() *NvidiaTaskSettings {
	r := &NvidiaTaskSettings{}

	return r
}

type NvidiaTaskSettingsVariant interface {
	NvidiaTaskSettingsCaster() *NvidiaTaskSettings
}

func (s *NvidiaTaskSettings) NvidiaTaskSettingsCaster() *NvidiaTaskSettings {
	return s
}

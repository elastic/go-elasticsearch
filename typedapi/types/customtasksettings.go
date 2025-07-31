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

import (
	"encoding/json"
)

// CustomTaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/inference/_types/CommonTypes.ts#L1138-L1152
type CustomTaskSettings struct {
	// Parameters Specifies parameters that are required to run the custom service. The
	// parameters depend on the model your custom service uses.
	// For example:
	// ```
	//
	//	"task_settings":{
	//	  "parameters":{
	//	    "input_type":"query",
	//	    "return_token":true
	//	  }
	//	}
	//
	// ```
	Parameters json.RawMessage `json:"parameters,omitempty"`
}

// NewCustomTaskSettings returns a CustomTaskSettings.
func NewCustomTaskSettings() *CustomTaskSettings {
	r := &CustomTaskSettings{}

	return r
}

type CustomTaskSettingsVariant interface {
	CustomTaskSettingsCaster() *CustomTaskSettings
}

func (s *CustomTaskSettings) CustomTaskSettingsCaster() *CustomTaskSettings {
	return s
}

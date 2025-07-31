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

package putwatsonx

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/watsonxservicetype"
)

// Request holds the request body struct for the package putwatsonx
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/inference/put_watsonx/PutWatsonxRequest.ts#L29-L75
type Request struct {

	// Service The type of service supported for the specified task type. In this case,
	// `watsonxai`.
	Service watsonxservicetype.WatsonxServiceType `json:"service"`
	// ServiceSettings Settings used to install the inference model. These settings are specific to
	// the `watsonxai` service.
	ServiceSettings types.WatsonxServiceSettings `json:"service_settings"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putwatsonx request: %w", err)
	}

	return &req, nil
}

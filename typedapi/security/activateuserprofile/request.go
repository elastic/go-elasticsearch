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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package activateuserprofile

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/granttype"
)

// Request holds the request body struct for the package activateuserprofile
//
// https://github.com/elastic/elasticsearch-specification/blob/cbfcc73d01310bed2a480ec35aaef98138b598e5/specification/security/activate_user_profile/Request.ts#L23-L76
type Request struct {

	// AccessToken The user's Elasticsearch access token or JWT.
	// Both `access` and `id` JWT token types are supported and they depend on the
	// underlying JWT realm configuration.
	// If you specify the `access_token` grant type, this parameter is required.
	// It is not valid with other grant types.
	AccessToken *string `json:"access_token,omitempty"`
	// GrantType The type of grant.
	GrantType granttype.GrantType `json:"grant_type"`
	// Password The user's password.
	// If you specify the `password` grant type, this parameter is required.
	// It is not valid with other grant types.
	Password *string `json:"password,omitempty"`
	// Username The username that identifies the user.
	// If you specify the `password` grant type, this parameter is required.
	// It is not valid with other grant types.
	Username *string `json:"username,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Activateuserprofile request: %w", err)
	}

	return &req, nil
}

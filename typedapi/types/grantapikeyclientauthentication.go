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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/grantapikeyclientauthenticationscheme"
)

// GrantApiKeyClientAuthentication type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/security/grant_api_key/types.ts#L67-L76
type GrantApiKeyClientAuthentication struct {
	// Scheme The client authentication scheme.
	Scheme grantapikeyclientauthenticationscheme.GrantApiKeyClientAuthenticationScheme `json:"scheme"`
	// Value The client authentication value, which is what normally follows the scheme in
	// the `ES-Client-Authentication` request header.
	Value string `json:"value"`
}

func (s *GrantApiKeyClientAuthentication) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "scheme":
			if err := dec.Decode(&s.Scheme); err != nil {
				return fmt.Errorf("%s | %w", "Scheme", err)
			}

		case "value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Value", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Value = o

		}
	}
	return nil
}

// NewGrantApiKeyClientAuthentication returns a GrantApiKeyClientAuthentication.
func NewGrantApiKeyClientAuthentication() *GrantApiKeyClientAuthentication {
	r := &GrantApiKeyClientAuthentication{}

	return r
}

type GrantApiKeyClientAuthenticationVariant interface {
	GrantApiKeyClientAuthenticationCaster() *GrantApiKeyClientAuthentication
}

func (s *GrantApiKeyClientAuthentication) GrantApiKeyClientAuthenticationCaster() *GrantApiKeyClientAuthentication {
	return s
}

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
)

// The prompt caching configuration to use for the completion request. Currently
// supported only for `elastic` provider.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/_types/CommonTypes.ts#L239-L252
type CacheControl struct {
	// Ttl The time to live of the cached content, for example `1h`.
	Ttl Duration `json:"ttl,omitempty"`
	// Type The type of cache control, for example `ephemeral`.
	Type *string `json:"type,omitempty"`
}

func (s *CacheControl) UnmarshalJSON(data []byte) error {

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

		case "ttl":
			if err := dec.Decode(&s.Ttl); err != nil {
				return fmt.Errorf("%s | %w", "Ttl", err)
			}

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = &o

		}
	}
	return nil
}

// NewCacheControl returns a CacheControl.
func NewCacheControl() *CacheControl {
	r := &CacheControl{}

	return r
}

type CacheControlVariant interface {
	CacheControlCaster() *CacheControl
}

func (s *CacheControl) CacheControlCaster() *CacheControl {
	return s
}

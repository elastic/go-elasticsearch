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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Http type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L633-L647
type Http struct {
	// Clients Information on current and recently-closed HTTP client connections.
	// Clients that have been closed longer than the
	// `http.client_stats.closed_channels.max_age` setting will not be represented
	// here.
	Clients []Client `json:"clients,omitempty"`
	// CurrentOpen Current number of open HTTP connections for the node.
	CurrentOpen *int `json:"current_open,omitempty"`
	// TotalOpened Total number of HTTP connections opened for the node.
	TotalOpened *int64 `json:"total_opened,omitempty"`
}

func (s *Http) UnmarshalJSON(data []byte) error {

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

		case "clients":
			if err := dec.Decode(&s.Clients); err != nil {
				return err
			}

		case "current_open":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CurrentOpen = &value
			case float64:
				f := int(v)
				s.CurrentOpen = &f
			}

		case "total_opened":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalOpened = &value
			case float64:
				f := int64(v)
				s.TotalOpened = &f
			}

		}
	}
	return nil
}

// NewHttp returns a Http.
func NewHttp() *Http {
	r := &Http{}

	return r
}

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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// HttpInputProxy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/watcher/_types/Input.ts#L67-L70
type HttpInputProxy struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func (s *HttpInputProxy) UnmarshalJSON(data []byte) error {

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

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}

		case "port":
			if err := dec.Decode(&s.Port); err != nil {
				return fmt.Errorf("%s | %w", "Port", err)
			}

		}
	}
	return nil
}

// NewHttpInputProxy returns a HttpInputProxy.
func NewHttpInputProxy() *HttpInputProxy {
	r := &HttpInputProxy{}

	return r
}

type HttpInputProxyVariant interface {
	HttpInputProxyCaster() *HttpInputProxy
}

func (s *HttpInputProxy) HttpInputProxyCaster() *HttpInputProxy {
	return s
}

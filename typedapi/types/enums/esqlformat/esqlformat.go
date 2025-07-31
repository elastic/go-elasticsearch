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

// Package esqlformat
package esqlformat

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/esql/_types/QueryParameters.ts#L20-L29
type EsqlFormat struct {
	Name string
}

var (
	Csv = EsqlFormat{"csv"}

	Json = EsqlFormat{"json"}

	Tsv = EsqlFormat{"tsv"}

	Txt = EsqlFormat{"txt"}

	Yaml = EsqlFormat{"yaml"}

	Cbor = EsqlFormat{"cbor"}

	Smile = EsqlFormat{"smile"}

	Arrow = EsqlFormat{"arrow"}
)

func (e EsqlFormat) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *EsqlFormat) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "csv":
		*e = Csv
	case "json":
		*e = Json
	case "tsv":
		*e = Tsv
	case "txt":
		*e = Txt
	case "yaml":
		*e = Yaml
	case "cbor":
		*e = Cbor
	case "smile":
		*e = Smile
	case "arrow":
		*e = Arrow
	default:
		*e = EsqlFormat{string(text)}
	}

	return nil
}

func (e EsqlFormat) String() string {
	return e.Name
}

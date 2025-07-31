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

// Package esqlclusterstatus
package esqlclusterstatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/esql/_types/EsqlResult.ts#L83-L89
type EsqlClusterStatus struct {
	Name string
}

var (
	Running = EsqlClusterStatus{"running"}

	Successful = EsqlClusterStatus{"successful"}

	Partial = EsqlClusterStatus{"partial"}

	Skipped = EsqlClusterStatus{"skipped"}

	Failed = EsqlClusterStatus{"failed"}
)

func (e EsqlClusterStatus) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *EsqlClusterStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "running":
		*e = Running
	case "successful":
		*e = Successful
	case "partial":
		*e = Partial
	case "skipped":
		*e = Skipped
	case "failed":
		*e = Failed
	default:
		*e = EsqlClusterStatus{string(text)}
	}

	return nil
}

func (e EsqlClusterStatus) String() string {
	return e.Name
}

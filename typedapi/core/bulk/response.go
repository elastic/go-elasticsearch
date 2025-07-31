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

package bulk

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operationtype"
)

// Response holds the response body struct for the package bulk
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/bulk/BulkResponse.ts#L24-L45
type Response struct {

	// Errors If `true`, one or more of the operations in the bulk request did not complete
	// successfully.
	Errors     bool   `json:"errors"`
	IngestTook *int64 `json:"ingest_took,omitempty"`
	// Items The result of each operation in the bulk request, in the order they were
	// submitted.
	Items []map[operationtype.OperationType]types.ResponseItem `json:"items"`
	// Took The length of time, in milliseconds, it took to process the bulk request.
	Took int64 `json:"took"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

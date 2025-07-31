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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package get

import (
	"encoding/json"
)

// Response holds the response body struct for the package get
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/get/GetResponse.ts#L23-L34
type Response struct {

	// Fields If the `stored_fields` parameter is set to `true` and `found` is `true`, it
	// contains the document fields stored in the index.
	Fields map[string]json.RawMessage `json:"fields,omitempty"`
	// Found Indicates whether the document exists.
	Found bool `json:"found"`
	// Id_ The unique identifier for the document.
	Id_      string   `json:"_id"`
	Ignored_ []string `json:"_ignored,omitempty"`
	// Index_ The name of the index the document belongs to.
	Index_ string `json:"_index"`
	// PrimaryTerm_ The primary term assigned to the document for the indexing operation.
	PrimaryTerm_ *int64 `json:"_primary_term,omitempty"`
	// Routing_ The explicit routing, if set.
	Routing_ *string `json:"_routing,omitempty"`
	// SeqNo_ The sequence number assigned to the document for the indexing operation.
	// Sequence numbers are used to ensure an older version of a document doesn't
	// overwrite a newer version.
	SeqNo_ *int64 `json:"_seq_no,omitempty"`
	// Source_ If `found` is `true`, it contains the document data formatted in JSON.
	// If the `_source` parameter is set to `false` or the `stored_fields` parameter
	// is set to `true`, it is excluded.
	Source_ json.RawMessage `json:"_source,omitempty"`
	// Version_ The document version, which is ncremented each time the document is updated.
	Version_ *int64 `json:"_version,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Fields: make(map[string]json.RawMessage, 0),
	}
	return r
}

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

package reindex

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conflicts"
)

// Request holds the request body struct for the package reindex
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/reindex/ReindexRequest.ts#L27-L317
type Request struct {

	// Conflicts Indicates whether to continue reindexing even when there are conflicts.
	Conflicts *conflicts.Conflicts `json:"conflicts,omitempty"`
	// Dest The destination you are copying to.
	Dest types.ReindexDestination `json:"dest"`
	// MaxDocs The maximum number of documents to reindex.
	// By default, all documents are reindexed.
	// If it is a value less then or equal to `scroll_size`, a scroll will not be
	// used to retrieve the results for the operation.
	//
	// If `conflicts` is set to `proceed`, the reindex operation could attempt to
	// reindex more documents from the source than `max_docs` until it has
	// successfully indexed `max_docs` documents into the target or it has gone
	// through every document in the source query.
	MaxDocs *int64 `json:"max_docs,omitempty"`
	// Script The script to run to update the document source or metadata when reindexing.
	Script *types.Script `json:"script,omitempty"`
	Size   *int64        `json:"size,omitempty"`
	// Source The source you are copying from.
	Source types.ReindexSource `json:"source"`
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
		return nil, fmt.Errorf("could not deserialise json into Reindex request: %w", err)
	}

	return &req, nil
}

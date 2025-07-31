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

package executewatch

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package executewatch
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/watcher/execute_watch/WatcherExecuteWatchResponse.ts#L23-L34
type Response struct {

	// Id_ The watch record identifier as it would be stored in the `.watcher-history`
	// index.
	Id_ string `json:"_id"`
	// WatchRecord The watch record document as it would be stored in the `.watcher-history`
	// index.
	WatchRecord types.WatchRecord `json:"watch_record"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

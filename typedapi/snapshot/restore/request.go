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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package restore

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package restore
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/restore/SnapshotRestoreRequest.ts#L25-L50
type Request struct {
	IgnoreIndexSettings []string `json:"ignore_index_settings,omitempty"`

	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`

	IncludeAliases *bool `json:"include_aliases,omitempty"`

	IncludeGlobalState *bool `json:"include_global_state,omitempty"`

	IndexSettings *types.IndexSettings `json:"index_settings,omitempty"`

	Indices *types.Indices `json:"indices,omitempty"`

	Partial *bool `json:"partial,omitempty"`

	RenamePattern *string `json:"rename_pattern,omitempty"`

	RenameReplacement *string `json:"rename_replacement,omitempty"`
}

// RequestBuilder is the builder API for the restore.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Restore request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) IgnoreIndexSettings(ignore_index_settings ...string) *RequestBuilder {
	rb.v.IgnoreIndexSettings = ignore_index_settings
	return rb
}

func (rb *RequestBuilder) IgnoreUnavailable(ignoreunavailable bool) *RequestBuilder {
	rb.v.IgnoreUnavailable = &ignoreunavailable
	return rb
}

func (rb *RequestBuilder) IncludeAliases(includealiases bool) *RequestBuilder {
	rb.v.IncludeAliases = &includealiases
	return rb
}

func (rb *RequestBuilder) IncludeGlobalState(includeglobalstate bool) *RequestBuilder {
	rb.v.IncludeGlobalState = &includeglobalstate
	return rb
}

func (rb *RequestBuilder) IndexSettings(indexsettings *types.IndexSettingsBuilder) *RequestBuilder {
	v := indexsettings.Build()
	rb.v.IndexSettings = &v
	return rb
}

func (rb *RequestBuilder) Indices(indices *types.IndicesBuilder) *RequestBuilder {
	v := indices.Build()
	rb.v.Indices = &v
	return rb
}

func (rb *RequestBuilder) Partial(partial bool) *RequestBuilder {
	rb.v.Partial = &partial
	return rb
}

func (rb *RequestBuilder) RenamePattern(renamepattern string) *RequestBuilder {
	rb.v.RenamePattern = &renamepattern
	return rb
}

func (rb *RequestBuilder) RenameReplacement(renamereplacement string) *RequestBuilder {
	rb.v.RenameReplacement = &renamereplacement
	return rb
}

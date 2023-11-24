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

// Creates or updates a document in an index.
package index

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Index struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      interface{}
	deferred []func(request interface{}) error
	raw      io.Reader

	paramSet int

	id    string
	index string
}

// NewIndex type alias for index.
type NewIndex func(index string) *Index

// NewIndexFunc returns a new instance of Index with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewIndexFunc(tp elastictransport.Interface) NewIndex {
	return func(index string) *Index {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Creates or updates a document in an index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
func New(tp elastictransport.Interface) *Index {
	r := &Index{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Index) Raw(raw io.Reader) *Index {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Index) Request(req interface{}) *Index {
	r.req = req

	return r
}

// Document allows to set the request property with the appropriate payload.
func (r *Index) Document(document interface{}) *Index {
	r.req = document

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Index) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Index: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_doc")
		path.WriteString("/")

		path.WriteString(r.id)

		method = http.MethodPut
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_doc")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Index) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Index query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a index.Response
func (r Index) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the Index headers map.
func (r *Index) Header(key, value string) *Index {
	r.headers.Set(key, value)

	return r
}

// Id Unique identifier for the document.
// API Name: id
func (r *Index) Id(id string) *Index {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index Name of the data stream or index to target.
// API Name: index
func (r *Index) _index(index string) *Index {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// IfPrimaryTerm Only perform the operation if the document has this primary term.
// API name: if_primary_term
func (r *Index) IfPrimaryTerm(ifprimaryterm string) *Index {
	r.values.Set("if_primary_term", ifprimaryterm)

	return r
}

// IfSeqNo Only perform the operation if the document has this sequence number.
// API name: if_seq_no
func (r *Index) IfSeqNo(sequencenumber string) *Index {
	r.values.Set("if_seq_no", sequencenumber)

	return r
}

// OpType Set to create to only index the document if it does not already exist (put if
// absent).
// If a document with the specified `_id` already exists, the indexing operation
// will fail.
// Same as using the `<index>/_create` endpoint.
// Valid values: `index`, `create`.
// If document id is specified, it defaults to `index`.
// Otherwise, it defaults to `create`.
// API name: op_type
func (r *Index) OpType(optype optype.OpType) *Index {
	r.values.Set("op_type", optype.String())

	return r
}

// Pipeline ID of the pipeline to use to preprocess incoming documents.
// If the index has a default ingest pipeline specified, then setting the value
// to `_none` disables the default ingest pipeline for this request.
// If a final pipeline is configured it will always run, regardless of the value
// of this parameter.
// API name: pipeline
func (r *Index) Pipeline(pipeline string) *Index {
	r.values.Set("pipeline", pipeline)

	return r
}

// Refresh If `true`, Elasticsearch refreshes the affected shards to make this operation
// visible to search, if `wait_for` then wait for a refresh to make this
// operation visible to search, if `false` do nothing with refreshes.
// Valid values: `true`, `false`, `wait_for`.
// API name: refresh
func (r *Index) Refresh(refresh refresh.Refresh) *Index {
	r.values.Set("refresh", refresh.String())

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Index) Routing(routing string) *Index {
	r.values.Set("routing", routing)

	return r
}

// Timeout Period the request waits for the following operations: automatic index
// creation, dynamic mapping updates, waiting for active shards.
// API name: timeout
func (r *Index) Timeout(duration string) *Index {
	r.values.Set("timeout", duration)

	return r
}

// Version Explicit version number for concurrency control.
// The specified version must match the current version of the document for the
// request to succeed.
// API name: version
func (r *Index) Version(versionnumber string) *Index {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType Specific version type: `external`, `external_gte`.
// API name: version_type
func (r *Index) VersionType(versiontype versiontype.VersionType) *Index {
	r.values.Set("version_type", versiontype.String())

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to all or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *Index) WaitForActiveShards(waitforactiveshards string) *Index {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// RequireAlias If `true`, the destination must be an index alias.
// API name: require_alias
func (r *Index) RequireAlias(requirealias bool) *Index {
	r.values.Set("require_alias", strconv.FormatBool(requirealias))

	return r
}

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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Updates a document with a script or partial document.
package update

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Update struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	id    string
	index string
}

// NewUpdate type alias for index.
type NewUpdate func(index, id string) *Update

// NewUpdateFunc returns a new instance of Update with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateFunc(tp elastictransport.Interface) NewUpdate {
	return func(index, id string) *Update {
		n := New(tp)

		n.Id(id)

		n.Index(index)

		return n
	}
}

// Updates a document with a script or partial document.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-update.html
func New(tp elastictransport.Interface) *Update {
	r := &Update{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Update) Raw(raw io.Reader) *Update {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Update) Request(req *Request) *Update {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Update) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Update: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_update")
		path.WriteString("/")

		path.WriteString(r.id)

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
func (r Update) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Update query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a update.Response
func (r Update) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Update headers map.
func (r *Update) Header(key, value string) *Update {
	r.headers.Set(key, value)

	return r
}

// Id Document ID
// API Name: id
func (r *Update) Id(v string) *Update {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Index The name of the index
// API Name: index
func (r *Update) Index(v string) *Update {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// IfPrimaryTerm Only perform the operation if the document has this primary term.
// API name: if_primary_term
func (r *Update) IfPrimaryTerm(v string) *Update {
	r.values.Set("if_primary_term", v)

	return r
}

// IfSeqNo Only perform the operation if the document has this sequence number.
// API name: if_seq_no
func (r *Update) IfSeqNo(v string) *Update {
	r.values.Set("if_seq_no", v)

	return r
}

// Lang The script language.
// API name: lang
func (r *Update) Lang(v string) *Update {
	r.values.Set("lang", v)

	return r
}

// Refresh If 'true', Elasticsearch refreshes the affected shards to make this operation
// visible to search, if 'wait_for' then wait for a refresh to make this
// operation
// visible to search, if 'false' do nothing with refreshes.
// API name: refresh
func (r *Update) Refresh(enum refresh.Refresh) *Update {
	r.values.Set("refresh", enum.String())

	return r
}

// RequireAlias If true, the destination must be an index alias.
// API name: require_alias
func (r *Update) RequireAlias(b bool) *Update {
	r.values.Set("require_alias", strconv.FormatBool(b))

	return r
}

// RetryOnConflict Specify how many times should the operation be retried when a conflict
// occurs.
// API name: retry_on_conflict
func (r *Update) RetryOnConflict(i int) *Update {
	r.values.Set("retry_on_conflict", strconv.Itoa(i))

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Update) Routing(v string) *Update {
	r.values.Set("routing", v)

	return r
}

// Timeout Period to wait for dynamic mapping updates and active shards.
// This guarantees Elasticsearch waits for at least the timeout before failing.
// The actual wait time could be longer, particularly when multiple waits occur.
// API name: timeout
func (r *Update) Timeout(v string) *Update {
	r.values.Set("timeout", v)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operations.
// Set to 'all' or any positive integer up to the total number of shards in the
// index
// (number_of_replicas+1). Defaults to 1 meaning the primary shard.
// API name: wait_for_active_shards
func (r *Update) WaitForActiveShards(v string) *Update {
	r.values.Set("wait_for_active_shards", v)

	return r
}

// Source_ Set to false to disable source retrieval. You can also specify a
// comma-separated
// list of the fields you want to retrieve.
// API name: _source
func (r *Update) Source_(v string) *Update {
	r.values.Set("_source", v)

	return r
}

// SourceExcludes_ Specify the source fields you want to exclude.
// API name: _source_excludes
func (r *Update) SourceExcludes_(v string) *Update {
	r.values.Set("_source_excludes", v)

	return r
}

// SourceIncludes_ Specify the source fields you want to retrieve.
// API name: _source_includes
func (r *Update) SourceIncludes_(v string) *Update {
	r.values.Set("_source_includes", v)

	return r
}

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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

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

	req interface{}
	raw io.Reader

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

		n.Index(index)

		return n
	}
}

// Creates or updates a document in an index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html
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

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Index) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

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

	return nil, errorResponse
}

// Header set a key, value pair in the Index headers map.
func (r *Index) Header(key, value string) *Index {
	r.headers.Set(key, value)

	return r
}

// Id Document ID
// API Name: id
func (r *Index) Id(v string) *Index {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Index The name of the index
// API Name: index
func (r *Index) Index(v string) *Index {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// IfPrimaryTerm only perform the index operation if the last operation that has changed the
// document has the specified primary term
// API name: if_primary_term
func (r *Index) IfPrimaryTerm(v string) *Index {
	r.values.Set("if_primary_term", v)

	return r
}

// IfSeqNo only perform the index operation if the last operation that has changed the
// document has the specified sequence number
// API name: if_seq_no
func (r *Index) IfSeqNo(v string) *Index {
	r.values.Set("if_seq_no", v)

	return r
}

// OpType Explicit operation type. Defaults to `index` for requests with an explicit
// document ID, and to `create`for requests without an explicit document ID
// API name: op_type
func (r *Index) OpType(enum optype.OpType) *Index {
	r.values.Set("op_type", enum.String())

	return r
}

// Pipeline The pipeline id to preprocess incoming documents with
// API name: pipeline
func (r *Index) Pipeline(v string) *Index {
	r.values.Set("pipeline", v)

	return r
}

// Refresh If `true` then refresh the affected shards to make this operation visible to
// search, if `wait_for` then wait for a refresh to make this operation visible
// to search, if `false` (the default) then do nothing with refreshes.
// API name: refresh
func (r *Index) Refresh(enum refresh.Refresh) *Index {
	r.values.Set("refresh", enum.String())

	return r
}

// Routing Specific routing value
// API name: routing
func (r *Index) Routing(v string) *Index {
	r.values.Set("routing", v)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Index) Timeout(v string) *Index {
	r.values.Set("timeout", v)

	return r
}

// Version Explicit version number for concurrency control
// API name: version
func (r *Index) Version(v string) *Index {
	r.values.Set("version", v)

	return r
}

// VersionType Specific version type
// API name: version_type
func (r *Index) VersionType(enum versiontype.VersionType) *Index {
	r.values.Set("version_type", enum.String())

	return r
}

// WaitForActiveShards Sets the number of shard copies that must be active before proceeding with
// the index operation. Defaults to 1, meaning the primary shard only. Set to
// `all` for all shard copies, otherwise set to any non-negative value less than
// or equal to the total number of copies for the shard (number of replicas + 1)
// API name: wait_for_active_shards
func (r *Index) WaitForActiveShards(v string) *Index {
	r.values.Set("wait_for_active_shards", v)

	return r
}

// RequireAlias When true, requires destination to be an alias. Default is false
// API name: require_alias
func (r *Index) RequireAlias(b bool) *Index {
	r.values.Set("require_alias", strconv.FormatBool(b))

	return r
}

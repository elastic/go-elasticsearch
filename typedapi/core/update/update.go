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

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

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

		n._id(id)

		n._index(index)

		return n
	}
}

// Updates a document with a script or partial document.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
func New(tp elastictransport.Interface) *Update {
	r := &Update{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
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
func (r *Update) _id(id string) *Update {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index The name of the index
// API Name: index
func (r *Update) _index(index string) *Update {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// IfPrimaryTerm Only perform the operation if the document has this primary term.
// API name: if_primary_term
func (r *Update) IfPrimaryTerm(ifprimaryterm string) *Update {
	r.values.Set("if_primary_term", ifprimaryterm)

	return r
}

// IfSeqNo Only perform the operation if the document has this sequence number.
// API name: if_seq_no
func (r *Update) IfSeqNo(sequencenumber string) *Update {
	r.values.Set("if_seq_no", sequencenumber)

	return r
}

// Lang The script language.
// API name: lang
func (r *Update) Lang(lang string) *Update {
	r.values.Set("lang", lang)

	return r
}

// Refresh If 'true', Elasticsearch refreshes the affected shards to make this operation
// visible to search, if 'wait_for' then wait for a refresh to make this
// operation
// visible to search, if 'false' do nothing with refreshes.
// API name: refresh
func (r *Update) Refresh(refresh refresh.Refresh) *Update {
	r.values.Set("refresh", refresh.String())

	return r
}

// RequireAlias If true, the destination must be an index alias.
// API name: require_alias
func (r *Update) RequireAlias(requirealias bool) *Update {
	r.values.Set("require_alias", strconv.FormatBool(requirealias))

	return r
}

// RetryOnConflict Specify how many times should the operation be retried when a conflict
// occurs.
// API name: retry_on_conflict
func (r *Update) RetryOnConflict(retryonconflict int) *Update {
	r.values.Set("retry_on_conflict", strconv.Itoa(retryonconflict))

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Update) Routing(routing string) *Update {
	r.values.Set("routing", routing)

	return r
}

// Timeout Period to wait for dynamic mapping updates and active shards.
// This guarantees Elasticsearch waits for at least the timeout before failing.
// The actual wait time could be longer, particularly when multiple waits occur.
// API name: timeout
func (r *Update) Timeout(duration string) *Update {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operations.
// Set to 'all' or any positive integer up to the total number of shards in the
// index
// (number_of_replicas+1). Defaults to 1 meaning the primary shard.
// API name: wait_for_active_shards
func (r *Update) WaitForActiveShards(waitforactiveshards string) *Update {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// SourceExcludes_ Specify the source fields you want to exclude.
// API name: _source_excludes
func (r *Update) SourceExcludes_(fields ...string) *Update {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ Specify the source fields you want to retrieve.
// API name: _source_includes
func (r *Update) SourceIncludes_(fields ...string) *Update {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// DetectNoop Set to false to disable setting 'result' in the response
// to 'noop' if no change to the document occurred.
// API name: detect_noop
func (r *Update) DetectNoop(detectnoop bool) *Update {
	r.req.DetectNoop = &detectnoop

	return r
}

// Doc A partial update to an existing document.
// API name: doc
//
// doc should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *Update) Doc(doc interface{}) *Update {
	switch casted := doc.(type) {
	case json.RawMessage:
		r.req.Doc = casted
	default:
		r.deferred = append(r.deferred, func(request *Request) error {
			data, err := json.Marshal(doc)
			if err != nil {
				return err
			}
			r.req.Doc = data
			return nil
		})
	}

	return r
}

// DocAsUpsert Set to true to use the contents of 'doc' as the value of 'upsert'
// API name: doc_as_upsert
func (r *Update) DocAsUpsert(docasupsert bool) *Update {
	r.req.DocAsUpsert = &docasupsert

	return r
}

// Script Script to execute to update the document.
// API name: script
func (r *Update) Script(script types.Script) *Update {
	r.req.Script = script

	return r
}

// ScriptedUpsert Set to true to execute the script whether or not the document exists.
// API name: scripted_upsert
func (r *Update) ScriptedUpsert(scriptedupsert bool) *Update {
	r.req.ScriptedUpsert = &scriptedupsert

	return r
}

// Source_ Set to false to disable source retrieval. You can also specify a
// comma-separated
// list of the fields you want to retrieve.
// API name: _source
func (r *Update) Source_(sourceconfig types.SourceConfig) *Update {
	r.req.Source_ = sourceconfig

	return r
}

// Upsert If the document does not already exist, the contents of 'upsert' are inserted
// as a
// new document. If the document exists, the 'script' is executed.
// API name: upsert
//
// upsert should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *Update) Upsert(upsert interface{}) *Update {
	switch casted := upsert.(type) {
	case json.RawMessage:
		r.req.Upsert = casted
	default:
		r.deferred = append(r.deferred, func(request *Request) error {
			data, err := json.Marshal(upsert)
			if err != nil {
				return err
			}
			r.req.Upsert = data
			return nil
		})
	}

	return r
}

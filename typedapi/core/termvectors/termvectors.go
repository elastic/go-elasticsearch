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

// Returns information and statistics about terms in the fields of a particular
// document.
package termvectors

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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	indexMask = iota + 1

	idMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Termvectors struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	index string
	id    string
}

// NewTermvectors type alias for index.
type NewTermvectors func(index string) *Termvectors

// NewTermvectorsFunc returns a new instance of Termvectors with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewTermvectorsFunc(tp elastictransport.Interface) NewTermvectors {
	return func(index string) *Termvectors {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Returns information and statistics about terms in the fields of a particular
// document.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-termvectors.html
func New(tp elastictransport.Interface) *Termvectors {
	r := &Termvectors{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Termvectors) Raw(raw io.Reader) *Termvectors {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Termvectors) Request(req *Request) *Termvectors {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Termvectors) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Termvectors: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_termvectors")
		path.WriteString("/")

		path.WriteString(r.id)

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_termvectors")

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
func (r Termvectors) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Termvectors query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a termvectors.Response
func (r Termvectors) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Termvectors headers map.
func (r *Termvectors) Header(key, value string) *Termvectors {
	r.headers.Set(key, value)

	return r
}

// Index The index in which the document resides.
// API Name: index
func (r *Termvectors) Index(v string) *Termvectors {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Id The id of the document, when not specified a doc param should be supplied.
// API Name: id
func (r *Termvectors) Id(v string) *Termvectors {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Fields A comma-separated list of fields to return.
// API name: fields
func (r *Termvectors) Fields(v string) *Termvectors {
	r.values.Set("fields", v)

	return r
}

// FieldStatistics Specifies if document count, sum of document frequencies and sum of total
// term frequencies should be returned.
// API name: field_statistics
func (r *Termvectors) FieldStatistics(b bool) *Termvectors {
	r.values.Set("field_statistics", strconv.FormatBool(b))

	return r
}

// Offsets Specifies if term offsets should be returned.
// API name: offsets
func (r *Termvectors) Offsets(b bool) *Termvectors {
	r.values.Set("offsets", strconv.FormatBool(b))

	return r
}

// Payloads Specifies if term payloads should be returned.
// API name: payloads
func (r *Termvectors) Payloads(b bool) *Termvectors {
	r.values.Set("payloads", strconv.FormatBool(b))

	return r
}

// Positions Specifies if term positions should be returned.
// API name: positions
func (r *Termvectors) Positions(b bool) *Termvectors {
	r.values.Set("positions", strconv.FormatBool(b))

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random).
// API name: preference
func (r *Termvectors) Preference(v string) *Termvectors {
	r.values.Set("preference", v)

	return r
}

// Realtime Specifies if request is real-time as opposed to near-real-time (default:
// true).
// API name: realtime
func (r *Termvectors) Realtime(b bool) *Termvectors {
	r.values.Set("realtime", strconv.FormatBool(b))

	return r
}

// Routing Specific routing value.
// API name: routing
func (r *Termvectors) Routing(v string) *Termvectors {
	r.values.Set("routing", v)

	return r
}

// TermStatistics Specifies if total term frequency and document frequency should be returned.
// API name: term_statistics
func (r *Termvectors) TermStatistics(b bool) *Termvectors {
	r.values.Set("term_statistics", strconv.FormatBool(b))

	return r
}

// Version Explicit version number for concurrency control
// API name: version
func (r *Termvectors) Version(v string) *Termvectors {
	r.values.Set("version", v)

	return r
}

// VersionType Specific version type
// API name: version_type
func (r *Termvectors) VersionType(enum versiontype.VersionType) *Termvectors {
	r.values.Set("version_type", enum.String())

	return r
}

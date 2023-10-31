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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

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

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

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

		n._index(index)

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

		req: NewRequest(),
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

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the Termvectors headers map.
func (r *Termvectors) Header(key, value string) *Termvectors {
	r.headers.Set(key, value)

	return r
}

// Index Name of the index that contains the document.
// API Name: index
func (r *Termvectors) _index(index string) *Termvectors {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Id Unique identifier of the document.
// API Name: id
func (r *Termvectors) Id(id string) *Termvectors {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Fields Comma-separated list or wildcard expressions of fields to include in the
// statistics.
// Used as the default list unless a specific field list is provided in the
// `completion_fields` or `fielddata_fields` parameters.
// API name: fields
func (r *Termvectors) Fields(fields ...string) *Termvectors {
	r.values.Set("fields", strings.Join(fields, ","))

	return r
}

// FieldStatistics If `true`, the response includes the document count, sum of document
// frequencies, and sum of total term frequencies.
// API name: field_statistics
func (r *Termvectors) FieldStatistics(fieldstatistics bool) *Termvectors {
	r.values.Set("field_statistics", strconv.FormatBool(fieldstatistics))

	return r
}

// Offsets If `true`, the response includes term offsets.
// API name: offsets
func (r *Termvectors) Offsets(offsets bool) *Termvectors {
	r.values.Set("offsets", strconv.FormatBool(offsets))

	return r
}

// Payloads If `true`, the response includes term payloads.
// API name: payloads
func (r *Termvectors) Payloads(payloads bool) *Termvectors {
	r.values.Set("payloads", strconv.FormatBool(payloads))

	return r
}

// Positions If `true`, the response includes term positions.
// API name: positions
func (r *Termvectors) Positions(positions bool) *Termvectors {
	r.values.Set("positions", strconv.FormatBool(positions))

	return r
}

// Preference Specifies the node or shard the operation should be performed on.
// Random by default.
// API name: preference
func (r *Termvectors) Preference(preference string) *Termvectors {
	r.values.Set("preference", preference)

	return r
}

// Realtime If true, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *Termvectors) Realtime(realtime bool) *Termvectors {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Termvectors) Routing(routing string) *Termvectors {
	r.values.Set("routing", routing)

	return r
}

// TermStatistics If `true`, the response includes term frequency and document frequency.
// API name: term_statistics
func (r *Termvectors) TermStatistics(termstatistics bool) *Termvectors {
	r.values.Set("term_statistics", strconv.FormatBool(termstatistics))

	return r
}

// Version If `true`, returns the document version as part of a hit.
// API name: version
func (r *Termvectors) Version(versionnumber string) *Termvectors {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType Specific version type.
// API name: version_type
func (r *Termvectors) VersionType(versiontype versiontype.VersionType) *Termvectors {
	r.values.Set("version_type", versiontype.String())

	return r
}

// Doc An artificial document (a document not present in the index) for which you
// want to retrieve term vectors.
// API name: doc
//
// doc should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *Termvectors) Doc(doc interface{}) *Termvectors {
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

// Filter Filter terms based on their tf-idf scores.
// API name: filter
func (r *Termvectors) Filter(filter *types.TermVectorsFilter) *Termvectors {

	r.req.Filter = filter

	return r
}

// PerFieldAnalyzer Overrides the default per-field analyzer.
// API name: per_field_analyzer
func (r *Termvectors) PerFieldAnalyzer(perfieldanalyzer map[string]string) *Termvectors {

	r.req.PerFieldAnalyzer = perfieldanalyzer

	return r
}

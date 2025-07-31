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

// Get term vector information.
//
// Get information and statistics about terms in the fields of a particular
// document.
//
// You can retrieve term vectors for documents stored in the index or for
// artificial documents passed in the body of the request.
// You can specify the fields you are interested in through the `fields`
// parameter or by adding the fields to the request body.
// For example:
//
// ```
// GET /my-index-000001/_termvectors/1?fields=message
// ```
//
// Fields can be specified using wildcards, similar to the multi match query.
//
// Term vectors are real-time by default, not near real-time.
// This can be changed by setting `realtime` parameter to `false`.
//
// You can request three types of values: _term information_, _term statistics_,
// and _field statistics_.
// By default, all term information and field statistics are returned for all
// fields but term statistics are excluded.
//
// **Term information**
//
// * term frequency in the field (always returned)
// * term positions (`positions: true`)
// * start and end offsets (`offsets: true`)
// * term payloads (`payloads: true`), as base64 encoded bytes
//
// If the requested information wasn't stored in the index, it will be computed
// on the fly if possible.
// Additionally, term vectors could be computed for documents not even existing
// in the index, but instead provided by the user.
//
// > warn
// > Start and end offsets assume UTF-16 encoding is being used. If you want to
// use these offsets in order to get the original text that produced this token,
// you should make sure that the string you are taking a sub-string of is also
// encoded using UTF-16.
//
// **Behaviour**
//
// The term and field statistics are not accurate.
// Deleted documents are not taken into account.
// The information is only retrieved for the shard the requested document
// resides in.
// The term and field statistics are therefore only useful as relative measures
// whereas the absolute numbers have no meaning in this context.
// By default, when requesting term vectors of artificial documents, a shard to
// get the statistics from is randomly selected.
// Use `routing` only to hit a particular shard.
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

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string
	id    string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Get term vector information.
//
// Get information and statistics about terms in the fields of a particular
// document.
//
// You can retrieve term vectors for documents stored in the index or for
// artificial documents passed in the body of the request.
// You can specify the fields you are interested in through the `fields`
// parameter or by adding the fields to the request body.
// For example:
//
// ```
// GET /my-index-000001/_termvectors/1?fields=message
// ```
//
// Fields can be specified using wildcards, similar to the multi match query.
//
// Term vectors are real-time by default, not near real-time.
// This can be changed by setting `realtime` parameter to `false`.
//
// You can request three types of values: _term information_, _term statistics_,
// and _field statistics_.
// By default, all term information and field statistics are returned for all
// fields but term statistics are excluded.
//
// **Term information**
//
// * term frequency in the field (always returned)
// * term positions (`positions: true`)
// * start and end offsets (`offsets: true`)
// * term payloads (`payloads: true`), as base64 encoded bytes
//
// If the requested information wasn't stored in the index, it will be computed
// on the fly if possible.
// Additionally, term vectors could be computed for documents not even existing
// in the index, but instead provided by the user.
//
// > warn
// > Start and end offsets assume UTF-16 encoding is being used. If you want to
// use these offsets in order to get the original text that produced this token,
// you should make sure that the string you are taking a sub-string of is also
// encoded using UTF-16.
//
// **Behaviour**
//
// The term and field statistics are not accurate.
// Deleted documents are not taken into account.
// The information is only retrieved for the shard the requested document
// resides in.
// The term and field statistics are therefore only useful as relative measures
// whereas the absolute numbers have no meaning in this context.
// By default, when requesting term vectors of artificial documents, a shard to
// get the statistics from is randomly selected.
// Use `routing` only to hit a particular shard.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
func New(tp elastictransport.Interface) *Termvectors {
	r := &Termvectors{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Termvectors: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_termvectors")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r Termvectors) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "termvectors")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "termvectors")
		if reader := instrument.RecordRequestBody(ctx, "termvectors", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "termvectors")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Termvectors query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a termvectors.Response
func (r Termvectors) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "termvectors")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the Termvectors headers map.
func (r *Termvectors) Header(key, value string) *Termvectors {
	r.headers.Set(key, value)

	return r
}

// Index The name of the index that contains the document.
// API Name: index
func (r *Termvectors) _index(index string) *Termvectors {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Id A unique identifier for the document.
// API Name: id
func (r *Termvectors) Id(id string) *Termvectors {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Preference The node or shard the operation should be performed on.
// It is random by default.
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

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Termvectors) ErrorTrace(errortrace bool) *Termvectors {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Termvectors) FilterPath(filterpaths ...string) *Termvectors {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *Termvectors) Human(human bool) *Termvectors {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Termvectors) Pretty(pretty bool) *Termvectors {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Doc An artificial document (a document not present in the index) for which you
// want to retrieve term vectors.
// API name: doc
//
// doc should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *Termvectors) Doc(doc any) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
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

// FieldStatistics If `true`, the response includes:
//
// * The document count (how many documents contain this field).
// * The sum of document frequencies (the sum of document frequencies for all
// terms in this field).
// * The sum of total term frequencies (the sum of total term frequencies of
// each term in this field).
// API name: field_statistics
func (r *Termvectors) FieldStatistics(fieldstatistics bool) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.FieldStatistics = &fieldstatistics

	return r
}

// Fields A list of fields to include in the statistics.
// It is used as the default list unless a specific field list is provided in
// the `completion_fields` or `fielddata_fields` parameters.
// API name: fields
func (r *Termvectors) Fields(fields ...string) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Fields = fields

	return r
}

// Filter Filter terms based on their tf-idf scores.
// This could be useful in order find out a good characteristic vector of a
// document.
// This feature works in a similar manner to the second phase of the More Like
// This Query.
// API name: filter
func (r *Termvectors) Filter(filter *types.TermVectorsFilter) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Filter = filter

	return r
}

// Offsets If `true`, the response includes term offsets.
// API name: offsets
func (r *Termvectors) Offsets(offsets bool) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Offsets = &offsets

	return r
}

// Payloads If `true`, the response includes term payloads.
// API name: payloads
func (r *Termvectors) Payloads(payloads bool) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Payloads = &payloads

	return r
}

// PerFieldAnalyzer Override the default per-field analyzer.
// This is useful in order to generate term vectors in any fashion, especially
// when using artificial documents.
// When providing an analyzer for a field that already stores term vectors, the
// term vectors will be regenerated.
// API name: per_field_analyzer
func (r *Termvectors) PerFieldAnalyzer(perfieldanalyzer map[string]string) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PerFieldAnalyzer = perfieldanalyzer

	return r
}

// Positions If `true`, the response includes term positions.
// API name: positions
func (r *Termvectors) Positions(positions bool) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Positions = &positions

	return r
}

// Routing A custom value that is used to route operations to a specific shard.
// API name: routing
func (r *Termvectors) Routing(routing string) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Routing = &routing

	return r
}

// TermStatistics If `true`, the response includes:
//
// * The total term frequency (how often a term occurs in all documents).
// * The document frequency (the number of documents containing the current
// term).
//
// By default these values are not returned since term statistics can have a
// serious performance impact.
// API name: term_statistics
func (r *Termvectors) TermStatistics(termstatistics bool) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.TermStatistics = &termstatistics

	return r
}

// Version If `true`, returns the document version as part of a hit.
// API name: version
func (r *Termvectors) Version(versionnumber int64) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Version = &versionnumber

	return r
}

// VersionType The version type.
// API name: version_type
func (r *Termvectors) VersionType(versiontype versiontype.VersionType) *Termvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.VersionType = &versiontype

	return r
}

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

// Create or update a synonym rule.
// Create or update a synonym rule in a synonym set.
//
// If any of the synonym rules included is invalid, the API returns an error.
//
// When you update a synonym rule, all analyzers using the synonyms set will be
// reloaded automatically to reflect the new rule.
package putsynonymrule

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
)

const (
	setidMask = iota + 1

	ruleidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutSynonymRule struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	setid  string
	ruleid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutSynonymRule type alias for index.
type NewPutSynonymRule func(setid, ruleid string) *PutSynonymRule

// NewPutSynonymRuleFunc returns a new instance of PutSynonymRule with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutSynonymRuleFunc(tp elastictransport.Interface) NewPutSynonymRule {
	return func(setid, ruleid string) *PutSynonymRule {
		n := New(tp)

		n._setid(setid)

		n._ruleid(ruleid)

		return n
	}
}

// Create or update a synonym rule.
// Create or update a synonym rule in a synonym set.
//
// If any of the synonym rules included is invalid, the API returns an error.
//
// When you update a synonym rule, all analyzers using the synonyms set will be
// reloaded automatically to reflect the new rule.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonym-rule.html
func New(tp elastictransport.Interface) *PutSynonymRule {
	r := &PutSynonymRule{
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
func (r *PutSynonymRule) Raw(raw io.Reader) *PutSynonymRule {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutSynonymRule) Request(req *Request) *PutSynonymRule {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutSynonymRule) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutSynonymRule: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == setidMask|ruleidMask:
		path.WriteString("/")
		path.WriteString("_synonyms")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "setid", r.setid)
		}
		path.WriteString(r.setid)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "ruleid", r.ruleid)
		}
		path.WriteString(r.ruleid)

		method = http.MethodPut
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
func (r PutSynonymRule) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "synonyms.put_synonym_rule")
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
		instrument.BeforeRequest(req, "synonyms.put_synonym_rule")
		if reader := instrument.RecordRequestBody(ctx, "synonyms.put_synonym_rule", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "synonyms.put_synonym_rule")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutSynonymRule query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putsynonymrule.Response
func (r PutSynonymRule) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "synonyms.put_synonym_rule")
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

// Header set a key, value pair in the PutSynonymRule headers map.
func (r *PutSynonymRule) Header(key, value string) *PutSynonymRule {
	r.headers.Set(key, value)

	return r
}

// SetId The ID of the synonym set.
// API Name: setid
func (r *PutSynonymRule) _setid(setid string) *PutSynonymRule {
	r.paramSet |= setidMask
	r.setid = setid

	return r
}

// RuleId The ID of the synonym rule to be updated or created.
// API Name: ruleid
func (r *PutSynonymRule) _ruleid(ruleid string) *PutSynonymRule {
	r.paramSet |= ruleidMask
	r.ruleid = ruleid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutSynonymRule) ErrorTrace(errortrace bool) *PutSynonymRule {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutSynonymRule) FilterPath(filterpaths ...string) *PutSynonymRule {
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
func (r *PutSynonymRule) Human(human bool) *PutSynonymRule {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutSynonymRule) Pretty(pretty bool) *PutSynonymRule {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Synonyms The synonym rule information definition, which must be in Solr format.
// API name: synonyms
func (r *PutSynonymRule) Synonyms(synonymstring string) *PutSynonymRule {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Synonyms = synonymstring

	return r
}

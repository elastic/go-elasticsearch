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
// https://github.com/elastic/elasticsearch-specification/tree/07bf82537a186562d8699685e3704ea338b268ef

// Retrieves a synonym rule from a synonym set
package getsynonymrule

import (
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

type GetSynonymRule struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	setid  string
	ruleid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetSynonymRule type alias for index.
type NewGetSynonymRule func(setid, ruleid string) *GetSynonymRule

// NewGetSynonymRuleFunc returns a new instance of GetSynonymRule with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetSynonymRuleFunc(tp elastictransport.Interface) NewGetSynonymRule {
	return func(setid, ruleid string) *GetSynonymRule {
		n := New(tp)

		n._setid(setid)

		n._ruleid(ruleid)

		return n
	}
}

// Retrieves a synonym rule from a synonym set
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonym-rule.html
func New(tp elastictransport.Interface) *GetSynonymRule {
	r := &GetSynonymRule{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetSynonymRule) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

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

		method = http.MethodGet
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
func (r GetSynonymRule) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "synonyms.get_synonym_rule")
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
		instrument.BeforeRequest(req, "synonyms.get_synonym_rule")
		if reader := instrument.RecordRequestBody(ctx, "synonyms.get_synonym_rule", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "synonyms.get_synonym_rule")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetSynonymRule query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getsynonymrule.Response
func (r GetSynonymRule) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "synonyms.get_synonym_rule")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetSynonymRule) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "synonyms.get_synonym_rule")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the GetSynonymRule query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetSynonymRule headers map.
func (r *GetSynonymRule) Header(key, value string) *GetSynonymRule {
	r.headers.Set(key, value)

	return r
}

// SetId The id of the synonym set to retrieve the synonym rule from
// API Name: setid
func (r *GetSynonymRule) _setid(setid string) *GetSynonymRule {
	r.paramSet |= setidMask
	r.setid = setid

	return r
}

// RuleId The id of the synonym rule to retrieve
// API Name: ruleid
func (r *GetSynonymRule) _ruleid(ruleid string) *GetSynonymRule {
	r.paramSet |= ruleidMask
	r.ruleid = ruleid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetSynonymRule) ErrorTrace(errortrace bool) *GetSynonymRule {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetSynonymRule) FilterPath(filterpaths ...string) *GetSynonymRule {
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
func (r *GetSynonymRule) Human(human bool) *GetSynonymRule {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetSynonymRule) Pretty(pretty bool) *GetSynonymRule {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

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

// Create or update a query rule.
// Create or update a query rule within a query ruleset.
//
// IMPORTANT: Due to limitations within pinned queries, you can only pin
// documents using ids or docs, but cannot use both in single rule.
// It is advised to use one or the other in query rulesets, to avoid errors.
// Additionally, pinned queries have a maximum limit of 100 pinned hits.
// If multiple matching rules pin more than 100 documents, only the first 100
// documents are pinned in the order they are specified in the ruleset.
package putrule

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/queryruletype"
)

const (
	rulesetidMask = iota + 1

	ruleidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutRule struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	rulesetid string
	ruleid    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutRule type alias for index.
type NewPutRule func(rulesetid, ruleid string) *PutRule

// NewPutRuleFunc returns a new instance of PutRule with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutRuleFunc(tp elastictransport.Interface) NewPutRule {
	return func(rulesetid, ruleid string) *PutRule {
		n := New(tp)

		n._rulesetid(rulesetid)

		n._ruleid(ruleid)

		return n
	}
}

// Create or update a query rule.
// Create or update a query rule within a query ruleset.
//
// IMPORTANT: Due to limitations within pinned queries, you can only pin
// documents using ids or docs, but cannot use both in single rule.
// It is advised to use one or the other in query rulesets, to avoid errors.
// Additionally, pinned queries have a maximum limit of 100 pinned hits.
// If multiple matching rules pin more than 100 documents, only the first 100
// documents are pinned in the order they are specified in the ruleset.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-rule.html
func New(tp elastictransport.Interface) *PutRule {
	r := &PutRule{
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
func (r *PutRule) Raw(raw io.Reader) *PutRule {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutRule) Request(req *Request) *PutRule {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutRule) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutRule: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == rulesetidMask|ruleidMask:
		path.WriteString("/")
		path.WriteString("_query_rules")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "rulesetid", r.rulesetid)
		}
		path.WriteString(r.rulesetid)
		path.WriteString("/")
		path.WriteString("_rule")
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
func (r PutRule) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "query_rules.put_rule")
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
		instrument.BeforeRequest(req, "query_rules.put_rule")
		if reader := instrument.RecordRequestBody(ctx, "query_rules.put_rule", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "query_rules.put_rule")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutRule query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putrule.Response
func (r PutRule) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "query_rules.put_rule")
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

// Header set a key, value pair in the PutRule headers map.
func (r *PutRule) Header(key, value string) *PutRule {
	r.headers.Set(key, value)

	return r
}

// RulesetId The unique identifier of the query ruleset containing the rule to be created
// or updated.
// API Name: rulesetid
func (r *PutRule) _rulesetid(rulesetid string) *PutRule {
	r.paramSet |= rulesetidMask
	r.rulesetid = rulesetid

	return r
}

// RuleId The unique identifier of the query rule within the specified ruleset to be
// created or updated.
// API Name: ruleid
func (r *PutRule) _ruleid(ruleid string) *PutRule {
	r.paramSet |= ruleidMask
	r.ruleid = ruleid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutRule) ErrorTrace(errortrace bool) *PutRule {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutRule) FilterPath(filterpaths ...string) *PutRule {
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
func (r *PutRule) Human(human bool) *PutRule {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutRule) Pretty(pretty bool) *PutRule {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Actions The actions to take when the rule is matched.
// The format of this action depends on the rule type.
// API name: actions
func (r *PutRule) Actions(actions *types.QueryRuleActions) *PutRule {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Actions = *actions

	return r
}

// Criteria The criteria that must be met for the rule to be applied.
// If multiple criteria are specified for a rule, all criteria must be met for
// the rule to be applied.
// API name: criteria
func (r *PutRule) Criteria(criteria ...types.QueryRuleCriteria) *PutRule {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Criteria = criteria

	return r
}

// API name: priority
func (r *PutRule) Priority(priority int) *PutRule {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Priority = &priority

	return r
}

// Type The type of rule.
// API name: type
func (r *PutRule) Type(type_ queryruletype.QueryRuleType) *PutRule {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Type = type_

	return r
}

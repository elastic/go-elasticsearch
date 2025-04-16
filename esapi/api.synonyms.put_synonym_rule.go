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
//
// Code generated from specification version 8.19.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newSynonymsPutSynonymRuleFunc(t Transport) SynonymsPutSynonymRule {
	return func(body io.Reader, rule_id string, set_id string, o ...func(*SynonymsPutSynonymRuleRequest)) (*Response, error) {
		var r = SynonymsPutSynonymRuleRequest{Body: body, RuleID: rule_id, SetID: set_id}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.Instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SynonymsPutSynonymRule creates or updates a synonym rule in a synonym set
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/put-synonym-rule.html.
type SynonymsPutSynonymRule func(body io.Reader, rule_id string, set_id string, o ...func(*SynonymsPutSynonymRuleRequest)) (*Response, error)

// SynonymsPutSynonymRuleRequest configures the Synonyms Put Synonym Rule API request.
type SynonymsPutSynonymRuleRequest struct {
	Body io.Reader

	RuleID string
	SetID  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SynonymsPutSynonymRuleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "synonyms.put_synonym_rule")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_synonyms") + 1 + len(r.SetID) + 1 + len(r.RuleID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_synonyms")
	path.WriteString("/")
	path.WriteString(r.SetID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "set_id", r.SetID)
	}
	path.WriteString("/")
	path.WriteString(r.RuleID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "rule_id", r.RuleID)
	}

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "synonyms.put_synonym_rule")
		if reader := instrument.RecordRequestBody(ctx, "synonyms.put_synonym_rule", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "synonyms.put_synonym_rule")
	}
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f SynonymsPutSynonymRule) WithContext(v context.Context) func(*SynonymsPutSynonymRuleRequest) {
	return func(r *SynonymsPutSynonymRuleRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SynonymsPutSynonymRule) WithPretty() func(*SynonymsPutSynonymRuleRequest) {
	return func(r *SynonymsPutSynonymRuleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SynonymsPutSynonymRule) WithHuman() func(*SynonymsPutSynonymRuleRequest) {
	return func(r *SynonymsPutSynonymRuleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SynonymsPutSynonymRule) WithErrorTrace() func(*SynonymsPutSynonymRuleRequest) {
	return func(r *SynonymsPutSynonymRuleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SynonymsPutSynonymRule) WithFilterPath(v ...string) func(*SynonymsPutSynonymRuleRequest) {
	return func(r *SynonymsPutSynonymRuleRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SynonymsPutSynonymRule) WithHeader(h map[string]string) func(*SynonymsPutSynonymRuleRequest) {
	return func(r *SynonymsPutSynonymRuleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SynonymsPutSynonymRule) WithOpaqueID(s string) func(*SynonymsPutSynonymRuleRequest) {
	return func(r *SynonymsPutSynonymRuleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}

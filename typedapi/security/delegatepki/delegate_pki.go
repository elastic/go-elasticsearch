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

// Delegate PKI authentication.
//
// This API implements the exchange of an X509Certificate chain for an
// Elasticsearch access token.
// The certificate chain is validated, according to RFC 5280, by sequentially
// considering the trust configuration of every installed PKI realm that has
// `delegation.enabled` set to `true`.
// A successfully trusted client certificate is also subject to the validation
// of the subject distinguished name according to thw `username_pattern` of the
// respective realm.
//
// This API is called by smart and trusted proxies, such as Kibana, which
// terminate the user's TLS session but still want to authenticate the user by
// using a PKI realm—-​as if the user connected directly to Elasticsearch.
//
// IMPORTANT: The association between the subject public key in the target
// certificate and the corresponding private key is not validated.
// This is part of the TLS authentication process and it is delegated to the
// proxy that calls this API.
// The proxy is trusted to have performed the TLS authentication and this API
// translates that authentication into an Elasticsearch access token.
package delegatepki

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DelegatePki struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewDelegatePki type alias for index.
type NewDelegatePki func() *DelegatePki

// NewDelegatePkiFunc returns a new instance of DelegatePki with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDelegatePkiFunc(tp elastictransport.Interface) NewDelegatePki {
	return func() *DelegatePki {
		n := New(tp)

		return n
	}
}

// Delegate PKI authentication.
//
// This API implements the exchange of an X509Certificate chain for an
// Elasticsearch access token.
// The certificate chain is validated, according to RFC 5280, by sequentially
// considering the trust configuration of every installed PKI realm that has
// `delegation.enabled` set to `true`.
// A successfully trusted client certificate is also subject to the validation
// of the subject distinguished name according to thw `username_pattern` of the
// respective realm.
//
// This API is called by smart and trusted proxies, such as Kibana, which
// terminate the user's TLS session but still want to authenticate the user by
// using a PKI realm—-​as if the user connected directly to Elasticsearch.
//
// IMPORTANT: The association between the subject public key in the target
// certificate and the corresponding private key is not validated.
// This is part of the TLS authentication process and it is delegated to the
// proxy that calls this API.
// The proxy is trusted to have performed the TLS authentication and this API
// translates that authentication into an Elasticsearch access token.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delegate-pki-authentication.html
func New(tp elastictransport.Interface) *DelegatePki {
	r := &DelegatePki{
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
func (r *DelegatePki) Raw(raw io.Reader) *DelegatePki {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *DelegatePki) Request(req *Request) *DelegatePki {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DelegatePki) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for DelegatePki: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("delegate_pki")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r DelegatePki) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.delegate_pki")
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
		instrument.BeforeRequest(req, "security.delegate_pki")
		if reader := instrument.RecordRequestBody(ctx, "security.delegate_pki", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.delegate_pki")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the DelegatePki query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a delegatepki.Response
func (r DelegatePki) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.delegate_pki")
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

// Header set a key, value pair in the DelegatePki headers map.
func (r *DelegatePki) Header(key, value string) *DelegatePki {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *DelegatePki) ErrorTrace(errortrace bool) *DelegatePki {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *DelegatePki) FilterPath(filterpaths ...string) *DelegatePki {
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
func (r *DelegatePki) Human(human bool) *DelegatePki {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *DelegatePki) Pretty(pretty bool) *DelegatePki {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// X509CertificateChain The X509Certificate chain, which is represented as an ordered string array.
// Each string in the array is a base64-encoded (Section 4 of RFC4648 - not
// base64url-encoded) of the certificate's DER encoding.
//
// The first element is the target certificate that contains the subject
// distinguished name that is requesting access.
// This may be followed by additional certificates; each subsequent certificate
// is used to certify the previous one.
// API name: x509_certificate_chain
func (r *DelegatePki) X509CertificateChain(x509certificatechains ...string) *DelegatePki {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.X509CertificateChain = x509certificatechains

	return r
}

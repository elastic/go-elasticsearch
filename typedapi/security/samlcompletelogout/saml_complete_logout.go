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

// Logout of SAML completely.
//
// Verifies the logout response sent from the SAML IdP.
//
// NOTE: This API is intended for use by custom web applications other than
// Kibana.
// If you are using Kibana, refer to the documentation for configuring SAML
// single-sign-on on the Elastic Stack.
//
// The SAML IdP may send a logout response back to the SP after handling the
// SP-initiated SAML Single Logout.
// This API verifies the response by ensuring the content is relevant and
// validating its signature.
// An empty response is returned if the verification process is successful.
// The response can be sent by the IdP with either the HTTP-Redirect or the
// HTTP-Post binding.
// The caller of this API must prepare the request accordingly so that this API
// can handle either of them.
package samlcompletelogout

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
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SamlCompleteLogout struct {
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

// NewSamlCompleteLogout type alias for index.
type NewSamlCompleteLogout func() *SamlCompleteLogout

// NewSamlCompleteLogoutFunc returns a new instance of SamlCompleteLogout with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSamlCompleteLogoutFunc(tp elastictransport.Interface) NewSamlCompleteLogout {
	return func() *SamlCompleteLogout {
		n := New(tp)

		return n
	}
}

// Logout of SAML completely.
//
// Verifies the logout response sent from the SAML IdP.
//
// NOTE: This API is intended for use by custom web applications other than
// Kibana.
// If you are using Kibana, refer to the documentation for configuring SAML
// single-sign-on on the Elastic Stack.
//
// The SAML IdP may send a logout response back to the SP after handling the
// SP-initiated SAML Single Logout.
// This API verifies the response by ensuring the content is relevant and
// validating its signature.
// An empty response is returned if the verification process is successful.
// The response can be sent by the IdP with either the HTTP-Redirect or the
// HTTP-Post binding.
// The caller of this API must prepare the request accordingly so that this API
// can handle either of them.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-complete-logout.html
func New(tp elastictransport.Interface) *SamlCompleteLogout {
	r := &SamlCompleteLogout{
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
func (r *SamlCompleteLogout) Raw(raw io.Reader) *SamlCompleteLogout {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SamlCompleteLogout) Request(req *Request) *SamlCompleteLogout {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SamlCompleteLogout) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for SamlCompleteLogout: %w", err)
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
		path.WriteString("saml")
		path.WriteString("/")
		path.WriteString("complete_logout")

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
func (r SamlCompleteLogout) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.saml_complete_logout")
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
		instrument.BeforeRequest(req, "security.saml_complete_logout")
		if reader := instrument.RecordRequestBody(ctx, "security.saml_complete_logout", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.saml_complete_logout")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the SamlCompleteLogout query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Header set a key, value pair in the SamlCompleteLogout headers map.
func (r *SamlCompleteLogout) Header(key, value string) *SamlCompleteLogout {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *SamlCompleteLogout) ErrorTrace(errortrace bool) *SamlCompleteLogout {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *SamlCompleteLogout) FilterPath(filterpaths ...string) *SamlCompleteLogout {
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
func (r *SamlCompleteLogout) Human(human bool) *SamlCompleteLogout {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *SamlCompleteLogout) Pretty(pretty bool) *SamlCompleteLogout {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Content If the SAML IdP sends the logout response with the HTTP-Post binding, this
// field must be set to the value of the SAMLResponse form parameter from the
// logout response.
// API name: content
func (r *SamlCompleteLogout) Content(content string) *SamlCompleteLogout {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Content = &content

	return r
}

// Ids A JSON array with all the valid SAML Request Ids that the caller of the API
// has for the current user.
// API name: ids
func (r *SamlCompleteLogout) Ids(ids ...string) *SamlCompleteLogout {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Ids = ids

	return r
}

// QueryString If the SAML IdP sends the logout response with the HTTP-Redirect binding,
// this field must be set to the query string of the redirect URI.
// API name: query_string
func (r *SamlCompleteLogout) QueryString(querystring string) *SamlCompleteLogout {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.QueryString = &querystring

	return r
}

// Realm The name of the SAML realm in Elasticsearch for which the configuration is
// used to verify the logout response.
// API name: realm
func (r *SamlCompleteLogout) Realm(realm string) *SamlCompleteLogout {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Realm = realm

	return r
}

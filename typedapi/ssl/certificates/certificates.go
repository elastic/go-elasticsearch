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

// Get SSL certificates.
//
// Get information about the X.509 certificates that are used to encrypt
// communications in the cluster.
// The API returns a list that includes certificates from all TLS contexts
// including:
//
// - Settings for transport and HTTP interfaces
// - TLS settings that are used within authentication realms
// - TLS settings for remote monitoring exporters
//
// The list includes certificates that are used for configuring trust, such as
// those configured in the `xpack.security.transport.ssl.truststore` and
// `xpack.security.transport.ssl.certificate_authorities` settings.
// It also includes certificates that are used for configuring server identity,
// such as `xpack.security.http.ssl.keystore` and
// `xpack.security.http.ssl.certificate settings`.
//
// The list does not include certificates that are sourced from the default SSL
// context of the Java Runtime Environment (JRE), even if those certificates are
// in use within Elasticsearch.
//
// NOTE: When a PKCS#11 token is configured as the truststore of the JRE, the
// API returns all the certificates that are included in the PKCS#11 token
// irrespective of whether these are used in the Elasticsearch TLS
// configuration.
//
// If Elasticsearch is configured to use a keystore or truststore, the API
// output includes all certificates in that store, even though some of the
// certificates might not be in active use within the cluster.
package certificates

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Certificates struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewCertificates type alias for index.
type NewCertificates func() *Certificates

// NewCertificatesFunc returns a new instance of Certificates with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCertificatesFunc(tp elastictransport.Interface) NewCertificates {
	return func() *Certificates {
		n := New(tp)

		return n
	}
}

// Get SSL certificates.
//
// Get information about the X.509 certificates that are used to encrypt
// communications in the cluster.
// The API returns a list that includes certificates from all TLS contexts
// including:
//
// - Settings for transport and HTTP interfaces
// - TLS settings that are used within authentication realms
// - TLS settings for remote monitoring exporters
//
// The list includes certificates that are used for configuring trust, such as
// those configured in the `xpack.security.transport.ssl.truststore` and
// `xpack.security.transport.ssl.certificate_authorities` settings.
// It also includes certificates that are used for configuring server identity,
// such as `xpack.security.http.ssl.keystore` and
// `xpack.security.http.ssl.certificate settings`.
//
// The list does not include certificates that are sourced from the default SSL
// context of the Java Runtime Environment (JRE), even if those certificates are
// in use within Elasticsearch.
//
// NOTE: When a PKCS#11 token is configured as the truststore of the JRE, the
// API returns all the certificates that are included in the PKCS#11 token
// irrespective of whether these are used in the Elasticsearch TLS
// configuration.
//
// If Elasticsearch is configured to use a keystore or truststore, the API
// output includes all certificates in that store, even though some of the
// certificates might not be in active use within the cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-ssl.html
func New(tp elastictransport.Interface) *Certificates {
	r := &Certificates{
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
func (r *Certificates) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ssl")
		path.WriteString("/")
		path.WriteString("certificates")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Certificates) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ssl.certificates")
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
		instrument.BeforeRequest(req, "ssl.certificates")
		if reader := instrument.RecordRequestBody(ctx, "ssl.certificates", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ssl.certificates")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Certificates query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a certificates.Response
func (r Certificates) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ssl.certificates")
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
		err = json.NewDecoder(res.Body).Decode(&response)
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
func (r Certificates) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ssl.certificates")
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
		err := fmt.Errorf("an error happened during the Certificates query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Certificates headers map.
func (r *Certificates) Header(key, value string) *Certificates {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Certificates) ErrorTrace(errortrace bool) *Certificates {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Certificates) FilterPath(filterpaths ...string) *Certificates {
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
func (r *Certificates) Human(human bool) *Certificates {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Certificates) Pretty(pretty bool) *Certificates {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

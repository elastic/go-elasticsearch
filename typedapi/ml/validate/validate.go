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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Validates an anomaly detection job.
package validate

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Validate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int
}

// NewValidate type alias for index.
type NewValidate func() *Validate

// NewValidateFunc returns a new instance of Validate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewValidateFunc(tp elastictransport.Interface) NewValidate {
	return func() *Validate {
		n := New(tp)

		return n
	}
}

// Validates an anomaly detection job.
//
// https://www.elastic.co/guide/en/machine-learning/current/ml-jobs.html
func New(tp elastictransport.Interface) *Validate {
	r := &Validate{
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
func (r *Validate) Raw(raw io.Reader) *Validate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Validate) Request(req *Request) *Validate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Validate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Validate: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")
		path.WriteString("_validate")

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
func (r Validate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Validate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a validate.Response
func (r Validate) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Validate headers map.
func (r *Validate) Header(key, value string) *Validate {
	r.headers.Set(key, value)

	return r
}

// API name: analysis_config
func (r *Validate) AnalysisConfig(analysisconfig *types.AnalysisConfig) *Validate {

	r.req.AnalysisConfig = analysisconfig

	return r
}

// API name: analysis_limits
func (r *Validate) AnalysisLimits(analysislimits *types.AnalysisLimits) *Validate {

	r.req.AnalysisLimits = analysislimits

	return r
}

// API name: data_description
func (r *Validate) DataDescription(datadescription *types.DataDescription) *Validate {

	r.req.DataDescription = datadescription

	return r
}

// API name: description
func (r *Validate) Description(description string) *Validate {

	r.req.Description = &description

	return r
}

// API name: job_id
func (r *Validate) JobId(id string) *Validate {
	r.req.JobId = &id

	return r
}

// API name: model_plot
func (r *Validate) ModelPlot(modelplot *types.ModelPlotConfig) *Validate {

	r.req.ModelPlot = modelplot

	return r
}

// API name: model_snapshot_id
func (r *Validate) ModelSnapshotId(id string) *Validate {
	r.req.ModelSnapshotId = &id

	return r
}

// API name: model_snapshot_retention_days
func (r *Validate) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *Validate {

	r.req.ModelSnapshotRetentionDays = &modelsnapshotretentiondays

	return r
}

// API name: results_index_name
func (r *Validate) ResultsIndexName(indexname string) *Validate {
	r.req.ResultsIndexName = &indexname

	return r
}

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

// Explains a data frame analytics config.
package explaindataframeanalytics

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

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExplainDataFrameAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	id string
}

// NewExplainDataFrameAnalytics type alias for index.
type NewExplainDataFrameAnalytics func() *ExplainDataFrameAnalytics

// NewExplainDataFrameAnalyticsFunc returns a new instance of ExplainDataFrameAnalytics with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExplainDataFrameAnalyticsFunc(tp elastictransport.Interface) NewExplainDataFrameAnalytics {
	return func() *ExplainDataFrameAnalytics {
		n := New(tp)

		return n
	}
}

// Explains a data frame analytics config.
//
// http://www.elastic.co/guide/en/elasticsearch/reference/current/explain-dfanalytics.html
func New(tp elastictransport.Interface) *ExplainDataFrameAnalytics {
	r := &ExplainDataFrameAnalytics{
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
func (r *ExplainDataFrameAnalytics) Raw(raw io.Reader) *ExplainDataFrameAnalytics {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ExplainDataFrameAnalytics) Request(req *Request) *ExplainDataFrameAnalytics {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ExplainDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ExplainDataFrameAnalytics: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")
		path.WriteString("_explain")

		method = http.MethodPost
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")

		path.WriteString(r.id)
		path.WriteString("/")
		path.WriteString("_explain")

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
func (r ExplainDataFrameAnalytics) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ExplainDataFrameAnalytics query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a explaindataframeanalytics.Response
func (r ExplainDataFrameAnalytics) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the ExplainDataFrameAnalytics headers map.
func (r *ExplainDataFrameAnalytics) Header(key, value string) *ExplainDataFrameAnalytics {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the data frame analytics job. This identifier can contain
// lowercase alphanumeric characters (a-z and 0-9), hyphens, and
// underscores. It must start and end with alphanumeric characters.
// API Name: id
func (r *ExplainDataFrameAnalytics) Id(id string) *ExplainDataFrameAnalytics {
	r.paramSet |= idMask
	r.id = id

	return r
}

// AllowLazyStart Specifies whether this job can start when there is insufficient machine
// learning node capacity for it to be immediately assigned to a node.
// API name: allow_lazy_start
func (r *ExplainDataFrameAnalytics) AllowLazyStart(allowlazystart bool) *ExplainDataFrameAnalytics {
	r.req.AllowLazyStart = &allowlazystart

	return r
}

// Analysis The analysis configuration, which contains the information necessary to
// perform one of the following types of analysis: classification, outlier
// detection, or regression.
// API name: analysis
func (r *ExplainDataFrameAnalytics) Analysis(analysis *types.DataframeAnalysisContainer) *ExplainDataFrameAnalytics {

	r.req.Analysis = analysis

	return r
}

// AnalyzedFields Specify includes and/or excludes patterns to select which fields will be
// included in the analysis. The patterns specified in excludes are applied
// last, therefore excludes takes precedence. In other words, if the same
// field is specified in both includes and excludes, then the field will not
// be included in the analysis.
// API name: analyzed_fields
func (r *ExplainDataFrameAnalytics) AnalyzedFields(analyzedfields *types.DataframeAnalysisAnalyzedFields) *ExplainDataFrameAnalytics {

	r.req.AnalyzedFields = analyzedfields

	return r
}

// Description A description of the job.
// API name: description
func (r *ExplainDataFrameAnalytics) Description(description string) *ExplainDataFrameAnalytics {

	r.req.Description = &description

	return r
}

// Dest The destination configuration, consisting of index and optionally
// results_field (ml by default).
// API name: dest
func (r *ExplainDataFrameAnalytics) Dest(dest *types.DataframeAnalyticsDestination) *ExplainDataFrameAnalytics {

	r.req.Dest = dest

	return r
}

// MaxNumThreads The maximum number of threads to be used by the analysis. Using more
// threads may decrease the time necessary to complete the analysis at the
// cost of using more CPU. Note that the process may use additional threads
// for operational functionality other than the analysis itself.
// API name: max_num_threads
func (r *ExplainDataFrameAnalytics) MaxNumThreads(maxnumthreads int) *ExplainDataFrameAnalytics {
	r.req.MaxNumThreads = &maxnumthreads

	return r
}

// ModelMemoryLimit The approximate maximum amount of memory resources that are permitted for
// analytical processing. If your `elasticsearch.yml` file contains an
// `xpack.ml.max_model_memory_limit` setting, an error occurs when you try to
// create data frame analytics jobs that have `model_memory_limit` values
// greater than that setting.
// API name: model_memory_limit
func (r *ExplainDataFrameAnalytics) ModelMemoryLimit(modelmemorylimit string) *ExplainDataFrameAnalytics {

	r.req.ModelMemoryLimit = &modelmemorylimit

	return r
}

// Source The configuration of how to source the analysis data. It requires an
// index. Optionally, query and _source may be specified.
// API name: source
func (r *ExplainDataFrameAnalytics) Source(source *types.DataframeAnalyticsSource) *ExplainDataFrameAnalytics {

	r.req.Source = source

	return r
}

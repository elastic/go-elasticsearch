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

// Instantiates a data frame analytics job.
package putdataframeanalytics

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

type PutDataFrameAnalytics struct {
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

// NewPutDataFrameAnalytics type alias for index.
type NewPutDataFrameAnalytics func(id string) *PutDataFrameAnalytics

// NewPutDataFrameAnalyticsFunc returns a new instance of PutDataFrameAnalytics with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDataFrameAnalyticsFunc(tp elastictransport.Interface) NewPutDataFrameAnalytics {
	return func(id string) *PutDataFrameAnalytics {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Instantiates a data frame analytics job.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-dfanalytics.html
func New(tp elastictransport.Interface) *PutDataFrameAnalytics {
	r := &PutDataFrameAnalytics{
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
func (r *PutDataFrameAnalytics) Raw(raw io.Reader) *PutDataFrameAnalytics {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDataFrameAnalytics) Request(req *Request) *PutDataFrameAnalytics {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutDataFrameAnalytics: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")

		path.WriteString(r.id)

		method = http.MethodPut
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
func (r PutDataFrameAnalytics) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutDataFrameAnalytics query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdataframeanalytics.Response
func (r PutDataFrameAnalytics) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutDataFrameAnalytics headers map.
func (r *PutDataFrameAnalytics) Header(key, value string) *PutDataFrameAnalytics {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the data frame analytics job. This identifier can contain
// lowercase alphanumeric characters (a-z and 0-9), hyphens, and
// underscores. It must start and end with alphanumeric characters.
// API Name: id
func (r *PutDataFrameAnalytics) _id(id string) *PutDataFrameAnalytics {
	r.paramSet |= idMask
	r.id = id

	return r
}

// AllowLazyStart Specifies whether this job can start when there is insufficient machine
// learning node capacity for it to be immediately assigned to a node. If
// set to `false` and a machine learning node with capacity to run the job
// cannot be immediately found, the API returns an error. If set to `true`,
// the API does not return an error; the job waits in the `starting` state
// until sufficient machine learning node capacity is available. This
// behavior is also affected by the cluster-wide
// `xpack.ml.max_lazy_ml_nodes` setting.
// API name: allow_lazy_start
func (r *PutDataFrameAnalytics) AllowLazyStart(allowlazystart bool) *PutDataFrameAnalytics {
	r.req.AllowLazyStart = &allowlazystart

	return r
}

// Analysis The analysis configuration, which contains the information necessary to
// perform one of the following types of analysis: classification, outlier
// detection, or regression.
// API name: analysis
func (r *PutDataFrameAnalytics) Analysis(analysis *types.DataframeAnalysisContainer) *PutDataFrameAnalytics {

	r.req.Analysis = *analysis

	return r
}

// AnalyzedFields Specifies `includes` and/or `excludes` patterns to select which fields
// will be included in the analysis. The patterns specified in `excludes`
// are applied last, therefore `excludes` takes precedence. In other words,
// if the same field is specified in both `includes` and `excludes`, then
// the field will not be included in the analysis. If `analyzed_fields` is
// not set, only the relevant fields will be included. For example, all the
// numeric fields for outlier detection.
// The supported fields vary for each type of analysis. Outlier detection
// requires numeric or `boolean` data to analyze. The algorithms don’t
// support missing values therefore fields that have data types other than
// numeric or boolean are ignored. Documents where included fields contain
// missing values, null values, or an array are also ignored. Therefore the
// `dest` index may contain documents that don’t have an outlier score.
// Regression supports fields that are numeric, `boolean`, `text`,
// `keyword`, and `ip` data types. It is also tolerant of missing values.
// Fields that are supported are included in the analysis, other fields are
// ignored. Documents where included fields contain an array with two or
// more values are also ignored. Documents in the `dest` index that don’t
// contain a results field are not included in the regression analysis.
// Classification supports fields that are numeric, `boolean`, `text`,
// `keyword`, and `ip` data types. It is also tolerant of missing values.
// Fields that are supported are included in the analysis, other fields are
// ignored. Documents where included fields contain an array with two or
// more values are also ignored. Documents in the `dest` index that don’t
// contain a results field are not included in the classification analysis.
// Classification analysis can be improved by mapping ordinal variable
// values to a single number. For example, in case of age ranges, you can
// model the values as `0-14 = 0`, `15-24 = 1`, `25-34 = 2`, and so on.
// API name: analyzed_fields
func (r *PutDataFrameAnalytics) AnalyzedFields(analyzedfields *types.DataframeAnalysisAnalyzedFields) *PutDataFrameAnalytics {

	r.req.AnalyzedFields = analyzedfields

	return r
}

// Description A description of the job.
// API name: description
func (r *PutDataFrameAnalytics) Description(description string) *PutDataFrameAnalytics {

	r.req.Description = &description

	return r
}

// Dest The destination configuration.
// API name: dest
func (r *PutDataFrameAnalytics) Dest(dest *types.DataframeAnalyticsDestination) *PutDataFrameAnalytics {

	r.req.Dest = *dest

	return r
}

// API name: headers
func (r *PutDataFrameAnalytics) Headers(httpheaders types.HttpHeaders) *PutDataFrameAnalytics {
	r.req.Headers = httpheaders

	return r
}

// MaxNumThreads The maximum number of threads to be used by the analysis. Using more
// threads may decrease the time necessary to complete the analysis at the
// cost of using more CPU. Note that the process may use additional threads
// for operational functionality other than the analysis itself.
// API name: max_num_threads
func (r *PutDataFrameAnalytics) MaxNumThreads(maxnumthreads int) *PutDataFrameAnalytics {
	r.req.MaxNumThreads = &maxnumthreads

	return r
}

// ModelMemoryLimit The approximate maximum amount of memory resources that are permitted for
// analytical processing. If your `elasticsearch.yml` file contains an
// `xpack.ml.max_model_memory_limit` setting, an error occurs when you try
// to create data frame analytics jobs that have `model_memory_limit` values
// greater than that setting.
// API name: model_memory_limit
func (r *PutDataFrameAnalytics) ModelMemoryLimit(modelmemorylimit string) *PutDataFrameAnalytics {

	r.req.ModelMemoryLimit = &modelmemorylimit

	return r
}

// Source The configuration of how to source the analysis data.
// API name: source
func (r *PutDataFrameAnalytics) Source(source *types.DataframeAnalyticsSource) *PutDataFrameAnalytics {

	r.req.Source = *source

	return r
}

// API name: version
func (r *PutDataFrameAnalytics) Version(versionstring string) *PutDataFrameAnalytics {
	r.req.Version = &versionstring

	return r
}

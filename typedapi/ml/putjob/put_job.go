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

// Instantiates an anomaly detection job.
package putjob

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
	jobidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	jobid string
}

// NewPutJob type alias for index.
type NewPutJob func(jobid string) *PutJob

// NewPutJobFunc returns a new instance of PutJob with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutJobFunc(tp elastictransport.Interface) NewPutJob {
	return func(jobid string) *PutJob {
		n := New(tp)

		n._jobid(jobid)

		return n
	}
}

// Instantiates an anomaly detection job.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-job.html
func New(tp elastictransport.Interface) *PutJob {
	r := &PutJob{
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
func (r *PutJob) Raw(raw io.Reader) *PutJob {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutJob) Request(req *Request) *PutJob {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutJob) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutJob: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		path.WriteString(r.jobid)

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
func (r PutJob) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutJob query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putjob.Response
func (r PutJob) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutJob headers map.
func (r *PutJob) Header(key, value string) *PutJob {
	r.headers.Set(key, value)

	return r
}

// JobId The identifier for the anomaly detection job. This identifier can contain
// lowercase alphanumeric characters (a-z and 0-9), hyphens, and underscores. It
// must start and end with alphanumeric characters.
// API Name: jobid
func (r *PutJob) _jobid(jobid string) *PutJob {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// AllowLazyOpen Advanced configuration option. Specifies whether this job can open when there
// is insufficient machine learning node capacity for it to be immediately
// assigned to a node. By default, if a machine learning node with capacity to
// run the job cannot immediately be found, the open anomaly detection jobs API
// returns an error. However, this is also subject to the cluster-wide
// `xpack.ml.max_lazy_ml_nodes` setting. If this option is set to true, the open
// anomaly detection jobs API does not return an error and the job waits in the
// opening state until sufficient machine learning node capacity is available.
// API name: allow_lazy_open
func (r *PutJob) AllowLazyOpen(allowlazyopen bool) *PutJob {
	r.req.AllowLazyOpen = &allowlazyopen

	return r
}

// AnalysisConfig Specifies how to analyze the data. After you create a job, you cannot change
// the analysis configuration; all the properties are informational.
// API name: analysis_config
func (r *PutJob) AnalysisConfig(analysisconfig *types.AnalysisConfig) *PutJob {

	r.req.AnalysisConfig = *analysisconfig

	return r
}

// AnalysisLimits Limits can be applied for the resources required to hold the mathematical
// models in memory. These limits are approximate and can be set per job. They
// do not control the memory used by other processes, for example the
// Elasticsearch Java processes.
// API name: analysis_limits
func (r *PutJob) AnalysisLimits(analysislimits *types.AnalysisLimits) *PutJob {

	r.req.AnalysisLimits = analysislimits

	return r
}

// BackgroundPersistInterval Advanced configuration option. The time between each periodic persistence of
// the model. The default value is a randomized value between 3 to 4 hours,
// which avoids all jobs persisting at exactly the same time. The smallest
// allowed value is 1 hour. For very large models (several GB), persistence
// could take 10-20 minutes, so do not set the `background_persist_interval`
// value too low.
// API name: background_persist_interval
func (r *PutJob) BackgroundPersistInterval(duration types.Duration) *PutJob {
	r.req.BackgroundPersistInterval = duration

	return r
}

// CustomSettings Advanced configuration option. Contains custom meta data about the job.
// API name: custom_settings
func (r *PutJob) CustomSettings(customsettings json.RawMessage) *PutJob {
	r.req.CustomSettings = customsettings

	return r
}

// DailyModelSnapshotRetentionAfterDays Advanced configuration option, which affects the automatic removal of old
// model snapshots for this job. It specifies a period of time (in days) after
// which only the first snapshot per day is retained. This period is relative to
// the timestamp of the most recent snapshot for this job. Valid values range
// from 0 to `model_snapshot_retention_days`.
// API name: daily_model_snapshot_retention_after_days
func (r *PutJob) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *PutJob {

	r.req.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays

	return r
}

// DataDescription Defines the format of the input data when you send data to the job by using
// the post data API. Note that when configure a datafeed, these properties are
// automatically set. When data is received via the post data API, it is not
// stored in Elasticsearch. Only the results for anomaly detection are retained.
// API name: data_description
func (r *PutJob) DataDescription(datadescription *types.DataDescription) *PutJob {

	r.req.DataDescription = *datadescription

	return r
}

// DatafeedConfig Defines a datafeed for the anomaly detection job. If Elasticsearch security
// features are enabled, your datafeed remembers which roles the user who
// created it had at the time of creation and runs the query using those same
// roles. If you provide secondary authorization headers, those credentials are
// used instead.
// API name: datafeed_config
func (r *PutJob) DatafeedConfig(datafeedconfig *types.DatafeedConfig) *PutJob {

	r.req.DatafeedConfig = datafeedconfig

	return r
}

// Description A description of the job.
// API name: description
func (r *PutJob) Description(description string) *PutJob {

	r.req.Description = &description

	return r
}

// Groups A list of job groups. A job can belong to no groups or many.
// API name: groups
func (r *PutJob) Groups(groups ...string) *PutJob {
	r.req.Groups = groups

	return r
}

// ModelPlotConfig This advanced configuration option stores model information along with the
// results. It provides a more detailed view into anomaly detection. If you
// enable model plot it can add considerable overhead to the performance of the
// system; it is not feasible for jobs with many entities. Model plot provides a
// simplified and indicative view of the model and its bounds. It does not
// display complex features such as multivariate correlations or multimodal
// data. As such, anomalies may occasionally be reported which cannot be seen in
// the model plot. Model plot config can be configured when the job is created
// or updated later. It must be disabled if performance issues are experienced.
// API name: model_plot_config
func (r *PutJob) ModelPlotConfig(modelplotconfig *types.ModelPlotConfig) *PutJob {

	r.req.ModelPlotConfig = modelplotconfig

	return r
}

// ModelSnapshotRetentionDays Advanced configuration option, which affects the automatic removal of old
// model snapshots for this job. It specifies the maximum period of time (in
// days) that snapshots are retained. This period is relative to the timestamp
// of the most recent snapshot for this job. By default, snapshots ten days
// older than the newest snapshot are deleted.
// API name: model_snapshot_retention_days
func (r *PutJob) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *PutJob {

	r.req.ModelSnapshotRetentionDays = &modelsnapshotretentiondays

	return r
}

// RenormalizationWindowDays Advanced configuration option. The period over which adjustments to the score
// are applied, as new data is seen. The default value is the longer of 30 days
// or 100 bucket spans.
// API name: renormalization_window_days
func (r *PutJob) RenormalizationWindowDays(renormalizationwindowdays int64) *PutJob {

	r.req.RenormalizationWindowDays = &renormalizationwindowdays

	return r
}

// ResultsIndexName A text string that affects the name of the machine learning results index. By
// default, the job generates an index named `.ml-anomalies-shared`.
// API name: results_index_name
func (r *PutJob) ResultsIndexName(indexname string) *PutJob {
	r.req.ResultsIndexName = &indexname

	return r
}

// ResultsRetentionDays Advanced configuration option. The period of time (in days) that results are
// retained. Age is calculated relative to the timestamp of the latest bucket
// result. If this property has a non-null value, once per day at 00:30 (server
// time), results that are the specified number of days older than the latest
// bucket result are deleted from Elasticsearch. The default value is null,
// which means all results are retained. Annotations generated by the system
// also count as results for retention purposes; they are deleted after the same
// number of days as results. Annotations added by users are retained forever.
// API name: results_retention_days
func (r *PutJob) ResultsRetentionDays(resultsretentiondays int64) *PutJob {

	r.req.ResultsRetentionDays = &resultsretentiondays

	return r
}

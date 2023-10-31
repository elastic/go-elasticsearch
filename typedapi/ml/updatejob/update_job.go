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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Updates certain properties of an anomaly detection job.
package updatejob

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

type UpdateJob struct {
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

// NewUpdateJob type alias for index.
type NewUpdateJob func(jobid string) *UpdateJob

// NewUpdateJobFunc returns a new instance of UpdateJob with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateJobFunc(tp elastictransport.Interface) NewUpdateJob {
	return func(jobid string) *UpdateJob {
		n := New(tp)

		n._jobid(jobid)

		return n
	}
}

// Updates certain properties of an anomaly detection job.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-job.html
func New(tp elastictransport.Interface) *UpdateJob {
	r := &UpdateJob{
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
func (r *UpdateJob) Raw(raw io.Reader) *UpdateJob {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateJob) Request(req *Request) *UpdateJob {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateJob) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for UpdateJob: %w", err)
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
		path.WriteString("/")
		path.WriteString("_update")

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
func (r UpdateJob) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the UpdateJob query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updatejob.Response
func (r UpdateJob) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the UpdateJob headers map.
func (r *UpdateJob) Header(key, value string) *UpdateJob {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the job.
// API Name: jobid
func (r *UpdateJob) _jobid(jobid string) *UpdateJob {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// AllowLazyOpen Advanced configuration option. Specifies whether this job can open when
// there is insufficient machine learning node capacity for it to be
// immediately assigned to a node. If `false` and a machine learning node
// with capacity to run the job cannot immediately be found, the open
// anomaly detection jobs API returns an error. However, this is also
// subject to the cluster-wide `xpack.ml.max_lazy_ml_nodes` setting. If this
// option is set to `true`, the open anomaly detection jobs API does not
// return an error and the job waits in the opening state until sufficient
// machine learning node capacity is available.
// API name: allow_lazy_open
func (r *UpdateJob) AllowLazyOpen(allowlazyopen bool) *UpdateJob {
	r.req.AllowLazyOpen = &allowlazyopen

	return r
}

// API name: analysis_limits
func (r *UpdateJob) AnalysisLimits(analysislimits *types.AnalysisMemoryLimit) *UpdateJob {

	r.req.AnalysisLimits = analysislimits

	return r
}

// BackgroundPersistInterval Advanced configuration option. The time between each periodic persistence
// of the model.
// The default value is a randomized value between 3 to 4 hours, which
// avoids all jobs persisting at exactly the same time. The smallest allowed
// value is 1 hour.
// For very large models (several GB), persistence could take 10-20 minutes,
// so do not set the value too low.
// If the job is open when you make the update, you must stop the datafeed,
// close the job, then reopen the job and restart the datafeed for the
// changes to take effect.
// API name: background_persist_interval
func (r *UpdateJob) BackgroundPersistInterval(duration types.Duration) *UpdateJob {
	r.req.BackgroundPersistInterval = duration

	return r
}

// API name: categorization_filters
func (r *UpdateJob) CategorizationFilters(categorizationfilters ...string) *UpdateJob {
	r.req.CategorizationFilters = categorizationfilters

	return r
}

// CustomSettings Advanced configuration option. Contains custom meta data about the job.
// For example, it can contain custom URL information as shown in Adding
// custom URLs to machine learning results.
// API name: custom_settings
func (r *UpdateJob) CustomSettings(customsettings map[string]json.RawMessage) *UpdateJob {

	r.req.CustomSettings = customsettings

	return r
}

// DailyModelSnapshotRetentionAfterDays Advanced configuration option, which affects the automatic removal of old
// model snapshots for this job. It specifies a period of time (in days)
// after which only the first snapshot per day is retained. This period is
// relative to the timestamp of the most recent snapshot for this job. Valid
// values range from 0 to `model_snapshot_retention_days`. For jobs created
// before version 7.8.0, the default value matches
// `model_snapshot_retention_days`.
// API name: daily_model_snapshot_retention_after_days
func (r *UpdateJob) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *UpdateJob {

	r.req.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays

	return r
}

// Description A description of the job.
// API name: description
func (r *UpdateJob) Description(description string) *UpdateJob {

	r.req.Description = &description

	return r
}

// Detectors An array of detector update objects.
// API name: detectors
func (r *UpdateJob) Detectors(detectors ...types.Detector) *UpdateJob {
	r.req.Detectors = detectors

	return r
}

// Groups A list of job groups. A job can belong to no groups or many.
// API name: groups
func (r *UpdateJob) Groups(groups ...string) *UpdateJob {
	r.req.Groups = groups

	return r
}

// API name: model_plot_config
func (r *UpdateJob) ModelPlotConfig(modelplotconfig *types.ModelPlotConfig) *UpdateJob {

	r.req.ModelPlotConfig = modelplotconfig

	return r
}

// API name: model_prune_window
func (r *UpdateJob) ModelPruneWindow(duration types.Duration) *UpdateJob {
	r.req.ModelPruneWindow = duration

	return r
}

// ModelSnapshotRetentionDays Advanced configuration option, which affects the automatic removal of old
// model snapshots for this job. It specifies the maximum period of time (in
// days) that snapshots are retained. This period is relative to the
// timestamp of the most recent snapshot for this job.
// API name: model_snapshot_retention_days
func (r *UpdateJob) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *UpdateJob {

	r.req.ModelSnapshotRetentionDays = &modelsnapshotretentiondays

	return r
}

// PerPartitionCategorization Settings related to how categorization interacts with partition fields.
// API name: per_partition_categorization
func (r *UpdateJob) PerPartitionCategorization(perpartitioncategorization *types.PerPartitionCategorization) *UpdateJob {

	r.req.PerPartitionCategorization = perpartitioncategorization

	return r
}

// RenormalizationWindowDays Advanced configuration option. The period over which adjustments to the
// score are applied, as new data is seen.
// API name: renormalization_window_days
func (r *UpdateJob) RenormalizationWindowDays(renormalizationwindowdays int64) *UpdateJob {

	r.req.RenormalizationWindowDays = &renormalizationwindowdays

	return r
}

// ResultsRetentionDays Advanced configuration option. The period of time (in days) that results
// are retained. Age is calculated relative to the timestamp of the latest
// bucket result. If this property has a non-null value, once per day at
// 00:30 (server time), results that are the specified number of days older
// than the latest bucket result are deleted from Elasticsearch. The default
// value is null, which means all results are retained.
// API name: results_retention_days
func (r *UpdateJob) ResultsRetentionDays(resultsretentiondays int64) *UpdateJob {

	r.req.ResultsRetentionDays = &resultsretentiondays

	return r
}

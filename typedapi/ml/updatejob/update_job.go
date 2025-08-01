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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Update an anomaly detection job.
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
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Update an anomaly detection job.
// Updates certain properties of an anomaly detection job.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-ml-update-job
func New(tp elastictransport.Interface) *UpdateJob {
	r := &UpdateJob{
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for UpdateJob: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "jobid", r.jobid)
		}
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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=9")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r UpdateJob) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.update_job")
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
		instrument.BeforeRequest(req, "ml.update_job")
		if reader := instrument.RecordRequestBody(ctx, "ml.update_job", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.update_job")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the UpdateJob query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updatejob.Response
func (r UpdateJob) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.update_job")
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

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *UpdateJob) ErrorTrace(errortrace bool) *UpdateJob {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *UpdateJob) FilterPath(filterpaths ...string) *UpdateJob {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *UpdateJob) Human(human bool) *UpdateJob {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *UpdateJob) Pretty(pretty bool) *UpdateJob {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Advanced configuration option. Specifies whether this job can open when
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
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AllowLazyOpen = &allowlazyopen

	return r
}

// API name: analysis_limits
func (r *UpdateJob) AnalysisLimits(analysislimits types.AnalysisMemoryLimitVariant) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AnalysisLimits = analysislimits.AnalysisMemoryLimitCaster()

	return r
}

// Advanced configuration option. The time between each periodic persistence
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
func (r *UpdateJob) BackgroundPersistInterval(duration types.DurationVariant) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.BackgroundPersistInterval = *duration.DurationCaster()

	return r
}

// API name: categorization_filters
func (r *UpdateJob) CategorizationFilters(categorizationfilters ...string) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range categorizationfilters {

		r.req.CategorizationFilters = append(r.req.CategorizationFilters, v)

	}
	return r
}

// Advanced configuration option. Contains custom meta data about the job.
// For example, it can contain custom URL information as shown in Adding
// custom URLs to machine learning results.
// API name: custom_settings
func (r *UpdateJob) CustomSettings(customsettings map[string]json.RawMessage) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CustomSettings = customsettings
	return r
}

func (r *UpdateJob) AddCustomSetting(key string, value json.RawMessage) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]json.RawMessage
	if r.req.CustomSettings == nil {
		r.req.CustomSettings = make(map[string]json.RawMessage)
	} else {
		tmp = r.req.CustomSettings
	}

	tmp[key] = value

	r.req.CustomSettings = tmp
	return r
}

// Advanced configuration option, which affects the automatic removal of old
// model snapshots for this job. It specifies a period of time (in days)
// after which only the first snapshot per day is retained. This period is
// relative to the timestamp of the most recent snapshot for this job. Valid
// values range from 0 to `model_snapshot_retention_days`. For jobs created
// before version 7.8.0, the default value matches
// `model_snapshot_retention_days`.
// API name: daily_model_snapshot_retention_after_days
func (r *UpdateJob) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays

	return r
}

// A description of the job.
// API name: description
func (r *UpdateJob) Description(description string) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = &description

	return r
}

// An array of detector update objects.
// API name: detectors
func (r *UpdateJob) Detectors(detectors ...types.DetectorUpdateVariant) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range detectors {

		r.req.Detectors = append(r.req.Detectors, *v.DetectorUpdateCaster())

	}
	return r
}

// A list of job groups. A job can belong to no groups or many.
// API name: groups
func (r *UpdateJob) Groups(groups ...string) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range groups {

		r.req.Groups = append(r.req.Groups, v)

	}
	return r
}

// API name: model_plot_config
func (r *UpdateJob) ModelPlotConfig(modelplotconfig types.ModelPlotConfigVariant) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ModelPlotConfig = modelplotconfig.ModelPlotConfigCaster()

	return r
}

// API name: model_prune_window
func (r *UpdateJob) ModelPruneWindow(duration types.DurationVariant) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ModelPruneWindow = *duration.DurationCaster()

	return r
}

// Advanced configuration option, which affects the automatic removal of old
// model snapshots for this job. It specifies the maximum period of time (in
// days) that snapshots are retained. This period is relative to the
// timestamp of the most recent snapshot for this job.
// API name: model_snapshot_retention_days
func (r *UpdateJob) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ModelSnapshotRetentionDays = &modelsnapshotretentiondays

	return r
}

// Settings related to how categorization interacts with partition fields.
// API name: per_partition_categorization
func (r *UpdateJob) PerPartitionCategorization(perpartitioncategorization types.PerPartitionCategorizationVariant) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PerPartitionCategorization = perpartitioncategorization.PerPartitionCategorizationCaster()

	return r
}

// Advanced configuration option. The period over which adjustments to the
// score are applied, as new data is seen.
// API name: renormalization_window_days
func (r *UpdateJob) RenormalizationWindowDays(renormalizationwindowdays int64) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RenormalizationWindowDays = &renormalizationwindowdays

	return r
}

// Advanced configuration option. The period of time (in days) that results
// are retained. Age is calculated relative to the timestamp of the latest
// bucket result. If this property has a non-null value, once per day at
// 00:30 (server time), results that are the specified number of days older
// than the latest bucket result are deleted from Elasticsearch. The default
// value is null, which means all results are retained.
// API name: results_retention_days
func (r *UpdateJob) ResultsRetentionDays(resultsretentiondays int64) *UpdateJob {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ResultsRetentionDays = &resultsretentiondays

	return r
}

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

// Upgrade all transforms.
//
// Transforms are compatible across minor versions and between supported major
// versions.
// However, over time, the format of transform configuration information may
// change.
// This API identifies transforms that have a legacy configuration format and
// upgrades them to the latest version.
// It also cleans up the internal data structures that store the transform state
// and checkpoints.
// The upgrade does not affect the source and destination indices.
// The upgrade also does not affect the roles that transforms use when
// Elasticsearch security features are enabled; the role used to read source
// data and write to the destination index remains unchanged.
//
// If a transform upgrade step fails, the upgrade stops and an error is returned
// about the underlying issue.
// Resolve the issue then re-run the process again.
// A summary is returned when the upgrade is finished.
//
// To ensure continuous transforms remain running during a major version upgrade
// of the cluster – for example, from 7.16 to 8.0 – it is recommended to upgrade
// transforms before upgrading the cluster.
// You may want to perform a recent cluster backup prior to the upgrade.
package upgradetransforms

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

type UpgradeTransforms struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewUpgradeTransforms type alias for index.
type NewUpgradeTransforms func() *UpgradeTransforms

// NewUpgradeTransformsFunc returns a new instance of UpgradeTransforms with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpgradeTransformsFunc(tp elastictransport.Interface) NewUpgradeTransforms {
	return func() *UpgradeTransforms {
		n := New(tp)

		return n
	}
}

// Upgrade all transforms.
//
// Transforms are compatible across minor versions and between supported major
// versions.
// However, over time, the format of transform configuration information may
// change.
// This API identifies transforms that have a legacy configuration format and
// upgrades them to the latest version.
// It also cleans up the internal data structures that store the transform state
// and checkpoints.
// The upgrade does not affect the source and destination indices.
// The upgrade also does not affect the roles that transforms use when
// Elasticsearch security features are enabled; the role used to read source
// data and write to the destination index remains unchanged.
//
// If a transform upgrade step fails, the upgrade stops and an error is returned
// about the underlying issue.
// Resolve the issue then re-run the process again.
// A summary is returned when the upgrade is finished.
//
// To ensure continuous transforms remain running during a major version upgrade
// of the cluster – for example, from 7.16 to 8.0 – it is recommended to upgrade
// transforms before upgrading the cluster.
// You may want to perform a recent cluster backup prior to the upgrade.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-transform-upgrade-transforms
func New(tp elastictransport.Interface) *UpgradeTransforms {
	r := &UpgradeTransforms{
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
func (r *UpgradeTransforms) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_transform")
		path.WriteString("/")
		path.WriteString("_upgrade")

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
func (r UpgradeTransforms) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "transform.upgrade_transforms")
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
		instrument.BeforeRequest(req, "transform.upgrade_transforms")
		if reader := instrument.RecordRequestBody(ctx, "transform.upgrade_transforms", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "transform.upgrade_transforms")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the UpgradeTransforms query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a upgradetransforms.Response
func (r UpgradeTransforms) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "transform.upgrade_transforms")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r UpgradeTransforms) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "transform.upgrade_transforms")
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
		err := fmt.Errorf("an error happened during the UpgradeTransforms query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the UpgradeTransforms headers map.
func (r *UpgradeTransforms) Header(key, value string) *UpgradeTransforms {
	r.headers.Set(key, value)

	return r
}

// DryRun When true, the request checks for updates but does not run them.
// API name: dry_run
func (r *UpgradeTransforms) DryRun(dryrun bool) *UpgradeTransforms {
	r.values.Set("dry_run", strconv.FormatBool(dryrun))

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and
// returns an error.
// API name: timeout
func (r *UpgradeTransforms) Timeout(duration string) *UpgradeTransforms {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *UpgradeTransforms) ErrorTrace(errortrace bool) *UpgradeTransforms {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *UpgradeTransforms) FilterPath(filterpaths ...string) *UpgradeTransforms {
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
func (r *UpgradeTransforms) Human(human bool) *UpgradeTransforms {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *UpgradeTransforms) Pretty(pretty bool) *UpgradeTransforms {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

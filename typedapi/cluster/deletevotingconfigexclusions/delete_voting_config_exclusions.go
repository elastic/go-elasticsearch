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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

// Clear cluster voting config exclusions.
// Remove master-eligible nodes from the voting configuration exclusion list.
package deletevotingconfigexclusions

import (
	"context"
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

type DeleteVotingConfigExclusions struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewDeleteVotingConfigExclusions type alias for index.
type NewDeleteVotingConfigExclusions func() *DeleteVotingConfigExclusions

// NewDeleteVotingConfigExclusionsFunc returns a new instance of DeleteVotingConfigExclusions with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteVotingConfigExclusionsFunc(tp elastictransport.Interface) NewDeleteVotingConfigExclusions {
	return func() *DeleteVotingConfigExclusions {
		n := New(tp)

		return n
	}
}

// Clear cluster voting config exclusions.
// Remove master-eligible nodes from the voting configuration exclusion list.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-cluster-post-voting-config-exclusions
func New(tp elastictransport.Interface) *DeleteVotingConfigExclusions {
	r := &DeleteVotingConfigExclusions{
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
func (r *DeleteVotingConfigExclusions) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("voting_config_exclusions")

		method = http.MethodDelete
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
func (r DeleteVotingConfigExclusions) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.delete_voting_config_exclusions")
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
		instrument.BeforeRequest(req, "cluster.delete_voting_config_exclusions")
		if reader := instrument.RecordRequestBody(ctx, "cluster.delete_voting_config_exclusions", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.delete_voting_config_exclusions")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the DeleteVotingConfigExclusions query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deletevotingconfigexclusions.Response
func (r DeleteVotingConfigExclusions) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r DeleteVotingConfigExclusions) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.delete_voting_config_exclusions")
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
		err := fmt.Errorf("an error happened during the DeleteVotingConfigExclusions query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the DeleteVotingConfigExclusions headers map.
func (r *DeleteVotingConfigExclusions) Header(key, value string) *DeleteVotingConfigExclusions {
	r.headers.Set(key, value)

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// API name: master_timeout
func (r *DeleteVotingConfigExclusions) MasterTimeout(duration string) *DeleteVotingConfigExclusions {
	r.values.Set("master_timeout", duration)

	return r
}

// WaitForRemoval Specifies whether to wait for all excluded nodes to be removed from the
// cluster before clearing the voting configuration exclusions list.
// Defaults to true, meaning that all excluded nodes must be removed from
// the cluster before this API takes any action. If set to false then the
// voting configuration exclusions list is cleared even if some excluded
// nodes are still in the cluster.
// API name: wait_for_removal
func (r *DeleteVotingConfigExclusions) WaitForRemoval(waitforremoval bool) *DeleteVotingConfigExclusions {
	r.values.Set("wait_for_removal", strconv.FormatBool(waitforremoval))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *DeleteVotingConfigExclusions) ErrorTrace(errortrace bool) *DeleteVotingConfigExclusions {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *DeleteVotingConfigExclusions) FilterPath(filterpaths ...string) *DeleteVotingConfigExclusions {
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
func (r *DeleteVotingConfigExclusions) Human(human bool) *DeleteVotingConfigExclusions {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *DeleteVotingConfigExclusions) Pretty(pretty bool) *DeleteVotingConfigExclusions {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

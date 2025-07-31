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

// Get snapshot information.
package get

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/snapshotsort"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Get struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	repository string
	snapshot   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGet type alias for index.
type NewGet func(repository, snapshot string) *Get

// NewGetFunc returns a new instance of Get with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetFunc(tp elastictransport.Interface) NewGet {
	return func(repository, snapshot string) *Get {
		n := New(tp)

		n._repository(repository)

		n._snapshot(snapshot)

		return n
	}
}

// Get snapshot information.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-snapshot-api.html
func New(tp elastictransport.Interface) *Get {
	r := &Get{
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
func (r *Get) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask|snapshotMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "repository", r.repository)
		}
		path.WriteString(r.repository)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "snapshot", r.snapshot)
		}
		path.WriteString(r.snapshot)

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
func (r Get) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "snapshot.get")
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
		instrument.BeforeRequest(req, "snapshot.get")
		if reader := instrument.RecordRequestBody(ctx, "snapshot.get", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.get")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Get query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a get.Response
func (r Get) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.get")
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
func (r Get) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.get")
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
		err := fmt.Errorf("an error happened during the Get query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Get headers map.
func (r *Get) Header(key, value string) *Get {
	r.headers.Set(key, value)

	return r
}

// Repository Comma-separated list of snapshot repository names used to limit the request.
// Wildcard (*) expressions are supported.
// API Name: repository
func (r *Get) _repository(repository string) *Get {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// Snapshot Comma-separated list of snapshot names to retrieve. Also accepts wildcards
// (*).
// - To get information about all snapshots in a registered repository, use a
// wildcard (*) or _all.
// - To get information about any snapshots that are currently running, use
// _current.
// API Name: snapshot
func (r *Get) _snapshot(snapshot string) *Get {
	r.paramSet |= snapshotMask
	r.snapshot = snapshot

	return r
}

// IgnoreUnavailable If false, the request returns an error for any snapshots that are
// unavailable.
// API name: ignore_unavailable
func (r *Get) IgnoreUnavailable(ignoreunavailable bool) *Get {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *Get) MasterTimeout(duration string) *Get {
	r.values.Set("master_timeout", duration)

	return r
}

// Verbose If true, returns additional information about each snapshot such as the
// version of Elasticsearch which took the snapshot, the start and end times of
// the snapshot, and the number of shards snapshotted.
// API name: verbose
func (r *Get) Verbose(verbose bool) *Get {
	r.values.Set("verbose", strconv.FormatBool(verbose))

	return r
}

// IndexDetails If true, returns additional information about each index in the snapshot
// comprising the number of shards in the index, the total size of the index in
// bytes, and the maximum number of segments per shard in the index. Defaults to
// false, meaning that this information is omitted.
// API name: index_details
func (r *Get) IndexDetails(indexdetails bool) *Get {
	r.values.Set("index_details", strconv.FormatBool(indexdetails))

	return r
}

// IndexNames If true, returns the name of each index in each snapshot.
// API name: index_names
func (r *Get) IndexNames(indexnames bool) *Get {
	r.values.Set("index_names", strconv.FormatBool(indexnames))

	return r
}

// IncludeRepository If true, returns the repository name in each snapshot.
// API name: include_repository
func (r *Get) IncludeRepository(includerepository bool) *Get {
	r.values.Set("include_repository", strconv.FormatBool(includerepository))

	return r
}

// Sort Allows setting a sort order for the result. Defaults to start_time, i.e.
// sorting by snapshot start time stamp.
// API name: sort
func (r *Get) Sort(sort snapshotsort.SnapshotSort) *Get {
	r.values.Set("sort", sort.String())

	return r
}

// Size Maximum number of snapshots to return. Defaults to 0 which means return all
// that match the request without limit.
// API name: size
func (r *Get) Size(size int) *Get {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// Order Sort order. Valid values are asc for ascending and desc for descending order.
// Defaults to asc, meaning ascending order.
// API name: order
func (r *Get) Order(order sortorder.SortOrder) *Get {
	r.values.Set("order", order.String())

	return r
}

// After Offset identifier to start pagination from as returned by the next field in
// the response body.
// API name: after
func (r *Get) After(after string) *Get {
	r.values.Set("after", after)

	return r
}

// Offset Numeric offset to start pagination from based on the snapshots matching this
// request. Using a non-zero value for this parameter is mutually exclusive with
// using the after parameter. Defaults to 0.
// API name: offset
func (r *Get) Offset(offset int) *Get {
	r.values.Set("offset", strconv.Itoa(offset))

	return r
}

// FromSortValue Value of the current sort column at which to start retrieval. Can either be a
// string snapshot- or repository name when sorting by snapshot or repository
// name, a millisecond time value or a number when sorting by index- or shard
// count.
// API name: from_sort_value
func (r *Get) FromSortValue(fromsortvalue string) *Get {
	r.values.Set("from_sort_value", fromsortvalue)

	return r
}

// SlmPolicyFilter Filter snapshots by a comma-separated list of SLM policy names that snapshots
// belong to. Also accepts wildcards (*) and combinations of wildcards followed
// by exclude patterns starting with -. To include snapshots not created by an
// SLM policy you can use the special pattern _none that will match all
// snapshots without an SLM policy.
// API name: slm_policy_filter
func (r *Get) SlmPolicyFilter(name string) *Get {
	r.values.Set("slm_policy_filter", name)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Get) ErrorTrace(errortrace bool) *Get {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Get) FilterPath(filterpaths ...string) *Get {
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
func (r *Get) Human(human bool) *Get {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Get) Pretty(pretty bool) *Get {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

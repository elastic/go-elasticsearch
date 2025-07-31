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

// Forget a follower.
// Remove the cross-cluster replication follower retention leases from the
// leader.
//
// A following index takes out retention leases on its leader index.
// These leases are used to increase the likelihood that the shards of the
// leader index retain the history of operations that the shards of the
// following index need to run replication.
// When a follower index is converted to a regular index by the unfollow API
// (either by directly calling the API or by index lifecycle management tasks),
// these leases are removed.
// However, removal of the leases can fail, for example when the remote cluster
// containing the leader index is unavailable.
// While the leases will eventually expire on their own, their extended
// existence can cause the leader index to hold more history than necessary and
// prevent index lifecycle management from performing some operations on the
// leader index.
// This API exists to enable manually removing the leases when the unfollow API
// is unable to do so.
//
// NOTE: This API does not stop replication by a following index. If you use
// this API with a follower index that is still actively following, the
// following index will add back retention leases on the leader.
// The only purpose of this API is to handle the case of failure to remove the
// following retention leases after the unfollow API is invoked.
package forgetfollower

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ForgetFollower struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewForgetFollower type alias for index.
type NewForgetFollower func(index string) *ForgetFollower

// NewForgetFollowerFunc returns a new instance of ForgetFollower with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewForgetFollowerFunc(tp elastictransport.Interface) NewForgetFollower {
	return func(index string) *ForgetFollower {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Forget a follower.
// Remove the cross-cluster replication follower retention leases from the
// leader.
//
// A following index takes out retention leases on its leader index.
// These leases are used to increase the likelihood that the shards of the
// leader index retain the history of operations that the shards of the
// following index need to run replication.
// When a follower index is converted to a regular index by the unfollow API
// (either by directly calling the API or by index lifecycle management tasks),
// these leases are removed.
// However, removal of the leases can fail, for example when the remote cluster
// containing the leader index is unavailable.
// While the leases will eventually expire on their own, their extended
// existence can cause the leader index to hold more history than necessary and
// prevent index lifecycle management from performing some operations on the
// leader index.
// This API exists to enable manually removing the leases when the unfollow API
// is unable to do so.
//
// NOTE: This API does not stop replication by a following index. If you use
// this API with a follower index that is still actively following, the
// following index will add back retention leases on the leader.
// The only purpose of this API is to handle the case of failure to remove the
// following retention leases after the unfollow API is invoked.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-forget-follower.html
func New(tp elastictransport.Interface) *ForgetFollower {
	r := &ForgetFollower{
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
func (r *ForgetFollower) Raw(raw io.Reader) *ForgetFollower {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ForgetFollower) Request(req *Request) *ForgetFollower {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ForgetFollower) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ForgetFollower: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_ccr")
		path.WriteString("/")
		path.WriteString("forget_follower")

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
func (r ForgetFollower) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ccr.forget_follower")
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
		instrument.BeforeRequest(req, "ccr.forget_follower")
		if reader := instrument.RecordRequestBody(ctx, "ccr.forget_follower", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ccr.forget_follower")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ForgetFollower query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a forgetfollower.Response
func (r ForgetFollower) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ccr.forget_follower")
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

// Header set a key, value pair in the ForgetFollower headers map.
func (r *ForgetFollower) Header(key, value string) *ForgetFollower {
	r.headers.Set(key, value)

	return r
}

// Index the name of the leader index for which specified follower retention leases
// should be removed
// API Name: index
func (r *ForgetFollower) _index(index string) *ForgetFollower {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *ForgetFollower) Timeout(duration string) *ForgetFollower {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ForgetFollower) ErrorTrace(errortrace bool) *ForgetFollower {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ForgetFollower) FilterPath(filterpaths ...string) *ForgetFollower {
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
func (r *ForgetFollower) Human(human bool) *ForgetFollower {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ForgetFollower) Pretty(pretty bool) *ForgetFollower {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: follower_cluster
func (r *ForgetFollower) FollowerCluster(followercluster string) *ForgetFollower {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.FollowerCluster = &followercluster

	return r
}

// API name: follower_index
func (r *ForgetFollower) FollowerIndex(indexname string) *ForgetFollower {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.FollowerIndex = &indexname

	return r
}

// API name: follower_index_uuid
func (r *ForgetFollower) FollowerIndexUuid(uuid string) *ForgetFollower {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.FollowerIndexUuid = &uuid

	return r
}

// API name: leader_remote_cluster
func (r *ForgetFollower) LeaderRemoteCluster(leaderremotecluster string) *ForgetFollower {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.LeaderRemoteCluster = &leaderremotecluster

	return r
}

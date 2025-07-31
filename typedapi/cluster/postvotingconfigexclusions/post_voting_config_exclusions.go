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

// Update voting configuration exclusions.
// Update the cluster voting config exclusions by node IDs or node names.
// By default, if there are more than three master-eligible nodes in the cluster
// and you remove fewer than half of the master-eligible nodes in the cluster at
// once, the voting configuration automatically shrinks.
// If you want to shrink the voting configuration to contain fewer than three
// nodes or to remove half or more of the master-eligible nodes in the cluster
// at once, use this API to remove departing nodes from the voting configuration
// manually.
// The API adds an entry for each specified node to the cluster’s voting
// configuration exclusions list.
// It then waits until the cluster has reconfigured its voting configuration to
// exclude the specified nodes.
//
// Clusters should have no voting configuration exclusions in normal operation.
// Once the excluded nodes have stopped, clear the voting configuration
// exclusions with `DELETE /_cluster/voting_config_exclusions`.
// This API waits for the nodes to be fully removed from the cluster before it
// returns.
// If your cluster has voting configuration exclusions for nodes that you no
// longer intend to remove, use `DELETE
// /_cluster/voting_config_exclusions?wait_for_removal=false` to clear the
// voting configuration exclusions without waiting for the nodes to leave the
// cluster.
//
// A response to `POST /_cluster/voting_config_exclusions` with an HTTP status
// code of 200 OK guarantees that the node has been removed from the voting
// configuration and will not be reinstated until the voting configuration
// exclusions are cleared by calling `DELETE
// /_cluster/voting_config_exclusions`.
// If the call to `POST /_cluster/voting_config_exclusions` fails or returns a
// response with an HTTP status code other than 200 OK then the node may not
// have been removed from the voting configuration.
// In that case, you may safely retry the call.
//
// NOTE: Voting exclusions are required only when you remove at least half of
// the master-eligible nodes from a cluster in a short time period.
// They are not required when removing master-ineligible nodes or when removing
// fewer than half of the master-eligible nodes.
package postvotingconfigexclusions

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

type PostVotingConfigExclusions struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPostVotingConfigExclusions type alias for index.
type NewPostVotingConfigExclusions func() *PostVotingConfigExclusions

// NewPostVotingConfigExclusionsFunc returns a new instance of PostVotingConfigExclusions with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPostVotingConfigExclusionsFunc(tp elastictransport.Interface) NewPostVotingConfigExclusions {
	return func() *PostVotingConfigExclusions {
		n := New(tp)

		return n
	}
}

// Update voting configuration exclusions.
// Update the cluster voting config exclusions by node IDs or node names.
// By default, if there are more than three master-eligible nodes in the cluster
// and you remove fewer than half of the master-eligible nodes in the cluster at
// once, the voting configuration automatically shrinks.
// If you want to shrink the voting configuration to contain fewer than three
// nodes or to remove half or more of the master-eligible nodes in the cluster
// at once, use this API to remove departing nodes from the voting configuration
// manually.
// The API adds an entry for each specified node to the cluster’s voting
// configuration exclusions list.
// It then waits until the cluster has reconfigured its voting configuration to
// exclude the specified nodes.
//
// Clusters should have no voting configuration exclusions in normal operation.
// Once the excluded nodes have stopped, clear the voting configuration
// exclusions with `DELETE /_cluster/voting_config_exclusions`.
// This API waits for the nodes to be fully removed from the cluster before it
// returns.
// If your cluster has voting configuration exclusions for nodes that you no
// longer intend to remove, use `DELETE
// /_cluster/voting_config_exclusions?wait_for_removal=false` to clear the
// voting configuration exclusions without waiting for the nodes to leave the
// cluster.
//
// A response to `POST /_cluster/voting_config_exclusions` with an HTTP status
// code of 200 OK guarantees that the node has been removed from the voting
// configuration and will not be reinstated until the voting configuration
// exclusions are cleared by calling `DELETE
// /_cluster/voting_config_exclusions`.
// If the call to `POST /_cluster/voting_config_exclusions` fails or returns a
// response with an HTTP status code other than 200 OK then the node may not
// have been removed from the voting configuration.
// In that case, you may safely retry the call.
//
// NOTE: Voting exclusions are required only when you remove at least half of
// the master-eligible nodes from a cluster in a short time period.
// They are not required when removing master-ineligible nodes or when removing
// fewer than half of the master-eligible nodes.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-cluster-post-voting-config-exclusions
func New(tp elastictransport.Interface) *PostVotingConfigExclusions {
	r := &PostVotingConfigExclusions{
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
func (r *PostVotingConfigExclusions) HttpRequest(ctx context.Context) (*http.Request, error) {
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PostVotingConfigExclusions) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.post_voting_config_exclusions")
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
		instrument.BeforeRequest(req, "cluster.post_voting_config_exclusions")
		if reader := instrument.RecordRequestBody(ctx, "cluster.post_voting_config_exclusions", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.post_voting_config_exclusions")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PostVotingConfigExclusions query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a postvotingconfigexclusions.Response
func (r PostVotingConfigExclusions) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r PostVotingConfigExclusions) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.post_voting_config_exclusions")
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
		err := fmt.Errorf("an error happened during the PostVotingConfigExclusions query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the PostVotingConfigExclusions headers map.
func (r *PostVotingConfigExclusions) Header(key, value string) *PostVotingConfigExclusions {
	r.headers.Set(key, value)

	return r
}

// NodeNames A comma-separated list of the names of the nodes to exclude from the
// voting configuration. If specified, you may not also specify node_ids.
// API name: node_names
func (r *PostVotingConfigExclusions) NodeNames(names ...string) *PostVotingConfigExclusions {
	r.values.Set("node_names", strings.Join(names, ","))

	return r
}

// NodeIds A comma-separated list of the persistent ids of the nodes to exclude
// from the voting configuration. If specified, you may not also specify
// node_names.
// API name: node_ids
func (r *PostVotingConfigExclusions) NodeIds(ids ...string) *PostVotingConfigExclusions {
	r.values.Set("node_ids", strings.Join(ids, ","))

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// API name: master_timeout
func (r *PostVotingConfigExclusions) MasterTimeout(duration string) *PostVotingConfigExclusions {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout When adding a voting configuration exclusion, the API waits for the
// specified nodes to be excluded from the voting configuration before
// returning. If the timeout expires before the appropriate condition
// is satisfied, the request fails and returns an error.
// API name: timeout
func (r *PostVotingConfigExclusions) Timeout(duration string) *PostVotingConfigExclusions {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PostVotingConfigExclusions) ErrorTrace(errortrace bool) *PostVotingConfigExclusions {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PostVotingConfigExclusions) FilterPath(filterpaths ...string) *PostVotingConfigExclusions {
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
func (r *PostVotingConfigExclusions) Human(human bool) *PostVotingConfigExclusions {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PostVotingConfigExclusions) Pretty(pretty bool) *PostVotingConfigExclusions {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

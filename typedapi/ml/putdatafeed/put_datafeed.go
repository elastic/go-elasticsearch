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

// Instantiates a datafeed.
package putdatafeed

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	datafeedid string
}

// NewPutDatafeed type alias for index.
type NewPutDatafeed func(datafeedid string) *PutDatafeed

// NewPutDatafeedFunc returns a new instance of PutDatafeed with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDatafeedFunc(tp elastictransport.Interface) NewPutDatafeed {
	return func(datafeedid string) *PutDatafeed {
		n := New(tp)

		n._datafeedid(datafeedid)

		return n
	}
}

// Instantiates a datafeed.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-datafeed.html
func New(tp elastictransport.Interface) *PutDatafeed {
	r := &PutDatafeed{
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
func (r *PutDatafeed) Raw(raw io.Reader) *PutDatafeed {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDatafeed) Request(req *Request) *PutDatafeed {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutDatafeed: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		path.WriteString(r.datafeedid)

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
func (r PutDatafeed) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutDatafeed query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdatafeed.Response
func (r PutDatafeed) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutDatafeed headers map.
func (r *PutDatafeed) Header(key, value string) *PutDatafeed {
	r.headers.Set(key, value)

	return r
}

// DatafeedId A numerical character string that uniquely identifies the datafeed.
// This identifier can contain lowercase alphanumeric characters (a-z and 0-9),
// hyphens, and underscores.
// It must start and end with alphanumeric characters.
// API Name: datafeedid
func (r *PutDatafeed) _datafeedid(datafeedid string) *PutDatafeed {
	r.paramSet |= datafeedidMask
	r.datafeedid = datafeedid

	return r
}

// AllowNoIndices If true, wildcard indices expressions that resolve into no concrete indices
// are ignored. This includes the `_all`
// string or when no indices are specified.
// API name: allow_no_indices
func (r *PutDatafeed) AllowNoIndices(allownoindices bool) *PutDatafeed {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams, this argument determines
// whether wildcard expressions match hidden data streams. Supports
// comma-separated values.
// API name: expand_wildcards
func (r *PutDatafeed) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutDatafeed {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled If true, concrete, expanded, or aliased indices are ignored when frozen.
// API name: ignore_throttled
func (r *PutDatafeed) IgnoreThrottled(ignorethrottled bool) *PutDatafeed {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable If true, unavailable indices (missing or closed) are ignored.
// API name: ignore_unavailable
func (r *PutDatafeed) IgnoreUnavailable(ignoreunavailable bool) *PutDatafeed {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Aggregations If set, the datafeed performs aggregation searches.
// Support for aggregations is limited and should be used only with low
// cardinality data.
// API name: aggregations
func (r *PutDatafeed) Aggregations(aggregations map[string]types.Aggregations) *PutDatafeed {

	r.req.Aggregations = aggregations

	return r
}

// ChunkingConfig Datafeeds might be required to search over long time periods, for several
// months or years.
// This search is split into time chunks in order to ensure the load on
// Elasticsearch is managed.
// Chunking configuration controls how the size of these time chunks are
// calculated;
// it is an advanced configuration option.
// API name: chunking_config
func (r *PutDatafeed) ChunkingConfig(chunkingconfig *types.ChunkingConfig) *PutDatafeed {

	r.req.ChunkingConfig = chunkingconfig

	return r
}

// DelayedDataCheckConfig Specifies whether the datafeed checks for missing data and the size of the
// window.
// The datafeed can optionally search over indices that have already been read
// in an effort to determine whether
// any data has subsequently been added to the index. If missing data is found,
// it is a good indication that the
// `query_delay` is set too low and the data is being indexed after the datafeed
// has passed that moment in time.
// This check runs only on real-time datafeeds.
// API name: delayed_data_check_config
func (r *PutDatafeed) DelayedDataCheckConfig(delayeddatacheckconfig *types.DelayedDataCheckConfig) *PutDatafeed {

	r.req.DelayedDataCheckConfig = delayeddatacheckconfig

	return r
}

// Frequency The interval at which scheduled queries are made while the datafeed runs in
// real time.
// The default value is either the bucket span for short bucket spans, or, for
// longer bucket spans, a sensible
// fraction of the bucket span. When `frequency` is shorter than the bucket
// span, interim results for the last
// (partial) bucket are written then eventually overwritten by the full bucket
// results. If the datafeed uses
// aggregations, this value must be divisible by the interval of the date
// histogram aggregation.
// API name: frequency
func (r *PutDatafeed) Frequency(duration types.Duration) *PutDatafeed {
	r.req.Frequency = duration

	return r
}

// API name: headers
func (r *PutDatafeed) Headers(httpheaders types.HttpHeaders) *PutDatafeed {
	r.req.Headers = httpheaders

	return r
}

// Indices An array of index names. Wildcards are supported. If any of the indices are
// in remote clusters, the machine
// learning nodes must have the `remote_cluster_client` role.
// API name: indices
func (r *PutDatafeed) Indices(indices ...string) *PutDatafeed {
	r.req.Indices = indices

	return r
}

// IndicesOptions Specifies index expansion options that are used during search
// API name: indices_options
func (r *PutDatafeed) IndicesOptions(indicesoptions *types.IndicesOptions) *PutDatafeed {

	r.req.IndicesOptions = indicesoptions

	return r
}

// JobId Identifier for the anomaly detection job.
// API name: job_id
func (r *PutDatafeed) JobId(id string) *PutDatafeed {
	r.req.JobId = &id

	return r
}

// MaxEmptySearches If a real-time datafeed has never seen any data (including during any initial
// training period), it automatically
// stops and closes the associated job after this many real-time searches return
// no documents. In other words,
// it stops after `frequency` times `max_empty_searches` of real-time operation.
// If not set, a datafeed with no
// end time that sees no data remains started until it is explicitly stopped. By
// default, it is not set.
// API name: max_empty_searches
func (r *PutDatafeed) MaxEmptySearches(maxemptysearches int) *PutDatafeed {
	r.req.MaxEmptySearches = &maxemptysearches

	return r
}

// Query The Elasticsearch query domain-specific language (DSL). This value
// corresponds to the query object in an
// Elasticsearch search POST body. All the options that are supported by
// Elasticsearch can be used, as this
// object is passed verbatim to Elasticsearch.
// API name: query
func (r *PutDatafeed) Query(query *types.Query) *PutDatafeed {

	r.req.Query = query

	return r
}

// QueryDelay The number of seconds behind real time that data is queried. For example, if
// data from 10:04 a.m. might
// not be searchable in Elasticsearch until 10:06 a.m., set this property to 120
// seconds. The default
// value is randomly selected between `60s` and `120s`. This randomness improves
// the query performance
// when there are multiple jobs running on the same node.
// API name: query_delay
func (r *PutDatafeed) QueryDelay(duration types.Duration) *PutDatafeed {
	r.req.QueryDelay = duration

	return r
}

// RuntimeMappings Specifies runtime fields for the datafeed search.
// API name: runtime_mappings
func (r *PutDatafeed) RuntimeMappings(runtimefields types.RuntimeFields) *PutDatafeed {
	r.req.RuntimeMappings = runtimefields

	return r
}

// ScriptFields Specifies scripts that evaluate custom expressions and returns script fields
// to the datafeed.
// The detector configuration objects in a job can contain functions that use
// these script fields.
// API name: script_fields
func (r *PutDatafeed) ScriptFields(scriptfields map[string]types.ScriptField) *PutDatafeed {

	r.req.ScriptFields = scriptfields

	return r
}

// ScrollSize The size parameter that is used in Elasticsearch searches when the datafeed
// does not use aggregations.
// The maximum value is the value of `index.max_result_window`, which is 10,000
// by default.
// API name: scroll_size
func (r *PutDatafeed) ScrollSize(scrollsize int) *PutDatafeed {
	r.req.ScrollSize = &scrollsize

	return r
}

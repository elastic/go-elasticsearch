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

// Retrieves anomaly detection job results for one or more influencers.
package getinfluencers

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
	jobidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetInfluencers struct {
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

// NewGetInfluencers type alias for index.
type NewGetInfluencers func(jobid string) *GetInfluencers

// NewGetInfluencersFunc returns a new instance of GetInfluencers with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetInfluencersFunc(tp elastictransport.Interface) NewGetInfluencers {
	return func(jobid string) *GetInfluencers {
		n := New(tp)

		n._jobid(jobid)

		return n
	}
}

// Retrieves anomaly detection job results for one or more influencers.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-influencer.html
func New(tp elastictransport.Interface) *GetInfluencers {
	r := &GetInfluencers{
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
func (r *GetInfluencers) Raw(raw io.Reader) *GetInfluencers {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GetInfluencers) Request(req *Request) *GetInfluencers {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetInfluencers) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for GetInfluencers: %w", err)
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
		path.WriteString("results")
		path.WriteString("/")
		path.WriteString("influencers")

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
func (r GetInfluencers) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetInfluencers query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getinfluencers.Response
func (r GetInfluencers) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the GetInfluencers headers map.
func (r *GetInfluencers) Header(key, value string) *GetInfluencers {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *GetInfluencers) _jobid(jobid string) *GetInfluencers {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// Desc If true, the results are sorted in descending order.
// API name: desc
func (r *GetInfluencers) Desc(desc bool) *GetInfluencers {
	r.values.Set("desc", strconv.FormatBool(desc))

	return r
}

// End Returns influencers with timestamps earlier than this time.
// The default value means it is unset and results are not limited to
// specific timestamps.
// API name: end
func (r *GetInfluencers) End(datetime string) *GetInfluencers {
	r.values.Set("end", datetime)

	return r
}

// ExcludeInterim If true, the output excludes interim results. By default, interim results
// are included.
// API name: exclude_interim
func (r *GetInfluencers) ExcludeInterim(excludeinterim bool) *GetInfluencers {
	r.values.Set("exclude_interim", strconv.FormatBool(excludeinterim))

	return r
}

// InfluencerScore Returns influencers with anomaly scores greater than or equal to this
// value.
// API name: influencer_score
func (r *GetInfluencers) InfluencerScore(influencerscore string) *GetInfluencers {
	r.values.Set("influencer_score", influencerscore)

	return r
}

// From Skips the specified number of influencers.
// API name: from
func (r *GetInfluencers) From(from int) *GetInfluencers {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// Size Specifies the maximum number of influencers to obtain.
// API name: size
func (r *GetInfluencers) Size(size int) *GetInfluencers {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// Sort Specifies the sort field for the requested influencers. By default, the
// influencers are sorted by the `influencer_score` value.
// API name: sort
func (r *GetInfluencers) Sort(field string) *GetInfluencers {
	r.values.Set("sort", field)

	return r
}

// Start Returns influencers with timestamps after this time. The default value
// means it is unset and results are not limited to specific timestamps.
// API name: start
func (r *GetInfluencers) Start(datetime string) *GetInfluencers {
	r.values.Set("start", datetime)

	return r
}

// Page Configures pagination.
// This parameter has the `from` and `size` properties.
// API name: page
func (r *GetInfluencers) Page(page *types.Page) *GetInfluencers {

	r.req.Page = page

	return r
}

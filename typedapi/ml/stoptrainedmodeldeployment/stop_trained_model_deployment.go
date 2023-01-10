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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Stop a trained model deployment.
package stoptrainedmodeldeployment

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StopTrainedModelDeployment struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	modelid string
}

// NewStopTrainedModelDeployment type alias for index.
type NewStopTrainedModelDeployment func(modelid string) *StopTrainedModelDeployment

// NewStopTrainedModelDeploymentFunc returns a new instance of StopTrainedModelDeployment with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStopTrainedModelDeploymentFunc(tp elastictransport.Interface) NewStopTrainedModelDeployment {
	return func(modelid string) *StopTrainedModelDeployment {
		n := New(tp)

		n.ModelId(modelid)

		return n
	}
}

// Stop a trained model deployment.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/stop-trained-model-deployment.html
func New(tp elastictransport.Interface) *StopTrainedModelDeployment {
	r := &StopTrainedModelDeployment{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StopTrainedModelDeployment) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		path.WriteString(r.modelid)
		path.WriteString("/")
		path.WriteString("deployment")
		path.WriteString("/")
		path.WriteString("_stop")

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

// Do runs the http.Request through the provided transport.
func (r StopTrainedModelDeployment) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the StopTrainedModelDeployment query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r StopTrainedModelDeployment) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the StopTrainedModelDeployment headers map.
func (r *StopTrainedModelDeployment) Header(key, value string) *StopTrainedModelDeployment {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model.
// API Name: modelid
func (r *StopTrainedModelDeployment) ModelId(v string) *StopTrainedModelDeployment {
	r.paramSet |= modelidMask
	r.modelid = v

	return r
}

// AllowNoMatch Specifies what to do when the request: contains wildcard expressions and
// there are no deployments that match;
// contains the  `_all` string or no identifiers and there are no matches; or
// contains wildcard expressions and
// there are only partial matches. By default, it returns an empty array when
// there are no matches and the subset of results when there are partial
// matches.
// If `false`, the request returns a 404 status code when there are no matches
// or only partial matches.
// API name: allow_no_match
func (r *StopTrainedModelDeployment) AllowNoMatch(b bool) *StopTrainedModelDeployment {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// Force Forcefully stops the deployment, even if it is used by ingest pipelines. You
// can't use these pipelines until you
// restart the model deployment.
// API name: force
func (r *StopTrainedModelDeployment) Force(b bool) *StopTrainedModelDeployment {
	r.values.Set("force", strconv.FormatBool(b))

	return r
}

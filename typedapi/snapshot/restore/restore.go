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

// Restores a snapshot.
package restore

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
	repositoryMask = iota + 1

	snapshotMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Restore struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	repository string
	snapshot   string
}

// NewRestore type alias for index.
type NewRestore func(repository, snapshot string) *Restore

// NewRestoreFunc returns a new instance of Restore with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRestoreFunc(tp elastictransport.Interface) NewRestore {
	return func(repository, snapshot string) *Restore {
		n := New(tp)

		n._repository(repository)

		n._snapshot(snapshot)

		return n
	}
}

// Restores a snapshot.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
func New(tp elastictransport.Interface) *Restore {
	r := &Restore{
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
func (r *Restore) Raw(raw io.Reader) *Restore {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Restore) Request(req *Request) *Restore {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Restore) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Restore: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask|snapshotMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		path.WriteString(r.repository)
		path.WriteString("/")

		path.WriteString(r.snapshot)
		path.WriteString("/")
		path.WriteString("_restore")

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
func (r Restore) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Restore query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a restore.Response
func (r Restore) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Restore headers map.
func (r *Restore) Header(key, value string) *Restore {
	r.headers.Set(key, value)

	return r
}

// Repository A repository name
// API Name: repository
func (r *Restore) _repository(repository string) *Restore {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// Snapshot A snapshot name
// API Name: snapshot
func (r *Restore) _snapshot(snapshot string) *Restore {
	r.paramSet |= snapshotMask
	r.snapshot = snapshot

	return r
}

// MasterTimeout Explicit operation timeout for connection to master node
// API name: master_timeout
func (r *Restore) MasterTimeout(duration string) *Restore {
	r.values.Set("master_timeout", duration)

	return r
}

// WaitForCompletion Should this request wait until the operation has completed before returning
// API name: wait_for_completion
func (r *Restore) WaitForCompletion(waitforcompletion bool) *Restore {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// API name: feature_states
func (r *Restore) FeatureStates(featurestates ...string) *Restore {
	r.req.FeatureStates = featurestates

	return r
}

// API name: ignore_index_settings
func (r *Restore) IgnoreIndexSettings(ignoreindexsettings ...string) *Restore {
	r.req.IgnoreIndexSettings = ignoreindexsettings

	return r
}

// API name: ignore_unavailable
func (r *Restore) IgnoreUnavailable(ignoreunavailable bool) *Restore {
	r.req.IgnoreUnavailable = &ignoreunavailable

	return r
}

// API name: include_aliases
func (r *Restore) IncludeAliases(includealiases bool) *Restore {
	r.req.IncludeAliases = &includealiases

	return r
}

// API name: include_global_state
func (r *Restore) IncludeGlobalState(includeglobalstate bool) *Restore {
	r.req.IncludeGlobalState = &includeglobalstate

	return r
}

// API name: index_settings
func (r *Restore) IndexSettings(indexsettings *types.IndexSettings) *Restore {

	r.req.IndexSettings = indexsettings

	return r
}

// API name: indices
func (r *Restore) Indices(indices ...string) *Restore {
	r.req.Indices = indices

	return r
}

// API name: partial
func (r *Restore) Partial(partial bool) *Restore {
	r.req.Partial = &partial

	return r
}

// API name: rename_pattern
func (r *Restore) RenamePattern(renamepattern string) *Restore {

	r.req.RenamePattern = &renamepattern

	return r
}

// API name: rename_replacement
func (r *Restore) RenameReplacement(renamereplacement string) *Restore {

	r.req.RenameReplacement = &renamereplacement

	return r
}

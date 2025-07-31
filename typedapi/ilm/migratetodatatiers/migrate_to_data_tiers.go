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

// Migrate to data tiers routing.
// Switch the indices, ILM policies, and legacy, composable, and component
// templates from using custom node attributes and attribute-based allocation
// filters to using data tiers.
// Optionally, delete one legacy index template.
// Using node roles enables ILM to automatically move the indices between data
// tiers.
//
// Migrating away from custom node attributes routing can be manually performed.
// This API provides an automated way of performing three out of the four manual
// steps listed in the migration guide:
//
// 1. Stop setting the custom hot attribute on new indices.
// 1. Remove custom allocation settings from existing ILM policies.
// 1. Replace custom allocation settings from existing indices with the
// corresponding tier preference.
//
// ILM must be stopped before performing the migration.
// Use the stop ILM and get ILM status APIs to wait until the reported operation
// mode is `STOPPED`.
package migratetodatatiers

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MigrateToDataTiers struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewMigrateToDataTiers type alias for index.
type NewMigrateToDataTiers func() *MigrateToDataTiers

// NewMigrateToDataTiersFunc returns a new instance of MigrateToDataTiers with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMigrateToDataTiersFunc(tp elastictransport.Interface) NewMigrateToDataTiers {
	return func() *MigrateToDataTiers {
		n := New(tp)

		return n
	}
}

// Migrate to data tiers routing.
// Switch the indices, ILM policies, and legacy, composable, and component
// templates from using custom node attributes and attribute-based allocation
// filters to using data tiers.
// Optionally, delete one legacy index template.
// Using node roles enables ILM to automatically move the indices between data
// tiers.
//
// Migrating away from custom node attributes routing can be manually performed.
// This API provides an automated way of performing three out of the four manual
// steps listed in the migration guide:
//
// 1. Stop setting the custom hot attribute on new indices.
// 1. Remove custom allocation settings from existing ILM policies.
// 1. Replace custom allocation settings from existing indices with the
// corresponding tier preference.
//
// ILM must be stopped before performing the migration.
// Use the stop ILM and get ILM status APIs to wait until the reported operation
// mode is `STOPPED`.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-migrate-to-data-tiers.html
func New(tp elastictransport.Interface) *MigrateToDataTiers {
	r := &MigrateToDataTiers{
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
func (r *MigrateToDataTiers) Raw(raw io.Reader) *MigrateToDataTiers {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *MigrateToDataTiers) Request(req *Request) *MigrateToDataTiers {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *MigrateToDataTiers) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for MigrateToDataTiers: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ilm")
		path.WriteString("/")
		path.WriteString("migrate_to_data_tiers")

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
func (r MigrateToDataTiers) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ilm.migrate_to_data_tiers")
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
		instrument.BeforeRequest(req, "ilm.migrate_to_data_tiers")
		if reader := instrument.RecordRequestBody(ctx, "ilm.migrate_to_data_tiers", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ilm.migrate_to_data_tiers")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the MigrateToDataTiers query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a migratetodatatiers.Response
func (r MigrateToDataTiers) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ilm.migrate_to_data_tiers")
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

// Header set a key, value pair in the MigrateToDataTiers headers map.
func (r *MigrateToDataTiers) Header(key, value string) *MigrateToDataTiers {
	r.headers.Set(key, value)

	return r
}

// DryRun If true, simulates the migration from node attributes based allocation
// filters to data tiers, but does not perform the migration.
// This provides a way to retrieve the indices and ILM policies that need to be
// migrated.
// API name: dry_run
func (r *MigrateToDataTiers) DryRun(dryrun bool) *MigrateToDataTiers {
	r.values.Set("dry_run", strconv.FormatBool(dryrun))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *MigrateToDataTiers) ErrorTrace(errortrace bool) *MigrateToDataTiers {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *MigrateToDataTiers) FilterPath(filterpaths ...string) *MigrateToDataTiers {
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
func (r *MigrateToDataTiers) Human(human bool) *MigrateToDataTiers {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *MigrateToDataTiers) Pretty(pretty bool) *MigrateToDataTiers {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: legacy_template_to_delete
func (r *MigrateToDataTiers) LegacyTemplateToDelete(legacytemplatetodelete string) *MigrateToDataTiers {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.LegacyTemplateToDelete = &legacytemplatetodelete

	return r
}

// API name: node_attribute
func (r *MigrateToDataTiers) NodeAttribute(nodeattribute string) *MigrateToDataTiers {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.NodeAttribute = &nodeattribute

	return r
}

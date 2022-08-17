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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package query

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package query
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/sql/query/QuerySqlRequest.ts#L28-L111
type Request struct {

	// Catalog Default catalog (cluster) for queries. If unspecified, the queries execute on
	// the data in the local cluster only.
	Catalog *string `json:"catalog,omitempty"`

	Columnar *bool `json:"columnar,omitempty"`

	Cursor *string `json:"cursor,omitempty"`

	// FetchSize The maximum number of rows (or entries) to return in one response
	FetchSize *int `json:"fetch_size,omitempty"`

	// FieldMultiValueLeniency Throw an exception when encountering multiple values for a field (default) or
	// be lenient and return the first value from the list (without any guarantees
	// of what that will be - typically the first in natural ascending order).
	FieldMultiValueLeniency *bool `json:"field_multi_value_leniency,omitempty"`

	// Filter Optional Elasticsearch query DSL for additional filtering.
	Filter *types.QueryContainer `json:"filter,omitempty"`

	// IndexUsingFrozen If true, the search can run on frozen indices. Defaults to false.
	IndexUsingFrozen *bool `json:"index_using_frozen,omitempty"`

	// KeepAlive Retention period for an async or saved synchronous search.
	KeepAlive *types.Duration `json:"keep_alive,omitempty"`

	// KeepOnCompletion If true, Elasticsearch stores synchronous searches if you also specify the
	// wait_for_completion_timeout parameter. If false, Elasticsearch only stores
	// async searches that don’t finish before the wait_for_completion_timeout.
	KeepOnCompletion *bool `json:"keep_on_completion,omitempty"`

	// PageTimeout The timeout before a pagination request fails.
	PageTimeout *types.Duration `json:"page_timeout,omitempty"`

	// Params Values for parameters in the query.
	Params map[string]interface{} `json:"params,omitempty"`

	// Query SQL query to execute
	Query *string `json:"query,omitempty"`

	// RequestTimeout The timeout before the request fails.
	RequestTimeout *types.Duration `json:"request_timeout,omitempty"`

	// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
	// precedence over mapped fields with the same name.
	RuntimeMappings *types.RuntimeFields `json:"runtime_mappings,omitempty"`

	// TimeZone Time-zone in ISO 8601 used for executing the query on the server. More
	// information available here.
	TimeZone *types.TimeZone `json:"time_zone,omitempty"`

	// WaitForCompletionTimeout Period to wait for complete results. Defaults to no timeout, meaning the
	// request waits for complete search results. If the search doesn’t finish
	// within this period, the search becomes async.
	WaitForCompletionTimeout *types.Duration `json:"wait_for_completion_timeout,omitempty"`
}

// RequestBuilder is the builder API for the query.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Params: make(map[string]interface{}, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Query request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Catalog(catalog string) *RequestBuilder {
	rb.v.Catalog = &catalog
	return rb
}

func (rb *RequestBuilder) Columnar(columnar bool) *RequestBuilder {
	rb.v.Columnar = &columnar
	return rb
}

func (rb *RequestBuilder) Cursor(cursor string) *RequestBuilder {
	rb.v.Cursor = &cursor
	return rb
}

func (rb *RequestBuilder) FetchSize(fetchsize int) *RequestBuilder {
	rb.v.FetchSize = &fetchsize
	return rb
}

func (rb *RequestBuilder) FieldMultiValueLeniency(fieldmultivalueleniency bool) *RequestBuilder {
	rb.v.FieldMultiValueLeniency = &fieldmultivalueleniency
	return rb
}

func (rb *RequestBuilder) Filter(filter *types.QueryContainerBuilder) *RequestBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *RequestBuilder) IndexUsingFrozen(indexusingfrozen bool) *RequestBuilder {
	rb.v.IndexUsingFrozen = &indexusingfrozen
	return rb
}

func (rb *RequestBuilder) KeepAlive(keepalive *types.DurationBuilder) *RequestBuilder {
	v := keepalive.Build()
	rb.v.KeepAlive = &v
	return rb
}

func (rb *RequestBuilder) KeepOnCompletion(keeponcompletion bool) *RequestBuilder {
	rb.v.KeepOnCompletion = &keeponcompletion
	return rb
}

func (rb *RequestBuilder) PageTimeout(pagetimeout *types.DurationBuilder) *RequestBuilder {
	v := pagetimeout.Build()
	rb.v.PageTimeout = &v
	return rb
}

func (rb *RequestBuilder) Params(value map[string]interface{}) *RequestBuilder {
	rb.v.Params = value
	return rb
}

func (rb *RequestBuilder) Query(query string) *RequestBuilder {
	rb.v.Query = &query
	return rb
}

func (rb *RequestBuilder) RequestTimeout(requesttimeout *types.DurationBuilder) *RequestBuilder {
	v := requesttimeout.Build()
	rb.v.RequestTimeout = &v
	return rb
}

func (rb *RequestBuilder) RuntimeMappings(runtimemappings *types.RuntimeFieldsBuilder) *RequestBuilder {
	v := runtimemappings.Build()
	rb.v.RuntimeMappings = &v
	return rb
}

func (rb *RequestBuilder) TimeZone(timezone types.TimeZone) *RequestBuilder {
	rb.v.TimeZone = &timezone
	return rb
}

func (rb *RequestBuilder) WaitForCompletionTimeout(waitforcompletiontimeout *types.DurationBuilder) *RequestBuilder {
	v := waitforcompletiontimeout.Build()
	rb.v.WaitForCompletionTimeout = &v
	return rb
}

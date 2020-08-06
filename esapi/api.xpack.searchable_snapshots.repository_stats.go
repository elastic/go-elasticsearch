// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 8.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

// NewSearchableSnapshotsRepositoryStats creates a new API client for the SearchableSnapshotsRepositoryStats endpoint
//
func NewSearchableSnapshotsRepositoryStats(t Transport) SearchableSnapshotsRepositoryStats {
	return func(repository string, o ...func(*SearchableSnapshotsRepositoryStatsRequest)) (*Response, error) {
		var r = SearchableSnapshotsRepositoryStatsRequest{Repository: repository}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SearchableSnapshotsRepositoryStats - Retrieve usage statistics about a snapshot repository.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/searchable-snapshots-repository-stats.html.
//
type SearchableSnapshotsRepositoryStats func(repository string, o ...func(*SearchableSnapshotsRepositoryStatsRequest)) (*Response, error)

// SearchableSnapshotsRepositoryStatsRequest configures the Searchable Snapshots Repository Stats API request.
//
type SearchableSnapshotsRepositoryStatsRequest struct {
	Repository string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SearchableSnapshotsRepositoryStatsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_snapshot") + 1 + len(r.Repository) + 1 + len("_stats"))
	path.WriteString("/")
	path.WriteString("_snapshot")
	path.WriteString("/")
	path.WriteString(r.Repository)
	path.WriteString("/")
	path.WriteString("_stats")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f SearchableSnapshotsRepositoryStats) WithContext(v context.Context) func(*SearchableSnapshotsRepositoryStatsRequest) {
	return func(r *SearchableSnapshotsRepositoryStatsRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SearchableSnapshotsRepositoryStats) WithPretty() func(*SearchableSnapshotsRepositoryStatsRequest) {
	return func(r *SearchableSnapshotsRepositoryStatsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SearchableSnapshotsRepositoryStats) WithHuman() func(*SearchableSnapshotsRepositoryStatsRequest) {
	return func(r *SearchableSnapshotsRepositoryStatsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SearchableSnapshotsRepositoryStats) WithErrorTrace() func(*SearchableSnapshotsRepositoryStatsRequest) {
	return func(r *SearchableSnapshotsRepositoryStatsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SearchableSnapshotsRepositoryStats) WithFilterPath(v ...string) func(*SearchableSnapshotsRepositoryStatsRequest) {
	return func(r *SearchableSnapshotsRepositoryStatsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f SearchableSnapshotsRepositoryStats) WithHeader(h map[string]string) func(*SearchableSnapshotsRepositoryStatsRequest) {
	return func(r *SearchableSnapshotsRepositoryStatsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f SearchableSnapshotsRepositoryStats) WithOpaqueID(s string) func(*SearchableSnapshotsRepositoryStatsRequest) {
	return func(r *SearchableSnapshotsRepositoryStatsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}

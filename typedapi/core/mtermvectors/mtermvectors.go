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


// Returns multiple termvectors in one request.
package mtermvectors

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Mtermvectors struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	index string
}

// NewMtermvectors type alias for index.
type NewMtermvectors func() *Mtermvectors

// NewMtermvectorsFunc returns a new instance of Mtermvectors with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMtermvectorsFunc(tp elastictransport.Interface) NewMtermvectors {
	return func() *Mtermvectors {
		n := New(tp)

		return n
	}
}

// Returns multiple termvectors in one request.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-multi-termvectors.html
func New(tp elastictransport.Interface) *Mtermvectors {
	r := &Mtermvectors{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Mtermvectors) Raw(raw json.RawMessage) *Mtermvectors {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Mtermvectors) Request(req *Request) *Mtermvectors {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Mtermvectors) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Mtermvectors: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_mtermvectors")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_mtermvectors")

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
func (r Mtermvectors) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Mtermvectors query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the Mtermvectors headers map.
func (r *Mtermvectors) Header(key, value string) *Mtermvectors {
	r.headers.Set(key, value)

	return r
}

// Index The index in which the document resides.
// API Name: index
func (r *Mtermvectors) Index(v string) *Mtermvectors {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Ids A comma-separated list of documents ids. You must define ids as parameter or
// set "ids" or "docs" in the request body
// API name: ids
func (r *Mtermvectors) Ids(value string) *Mtermvectors {
	r.values.Set("ids", value)

	return r
}

// Fields A comma-separated list of fields to return. Applies to all returned documents
// unless otherwise specified in body "params" or "docs".
// API name: fields
func (r *Mtermvectors) Fields(value string) *Mtermvectors {
	r.values.Set("fields", value)

	return r
}

// FieldStatistics Specifies if document count, sum of document frequencies and sum of total
// term frequencies should be returned. Applies to all returned documents unless
// otherwise specified in body "params" or "docs".
// API name: field_statistics
func (r *Mtermvectors) FieldStatistics(b bool) *Mtermvectors {
	r.values.Set("field_statistics", strconv.FormatBool(b))

	return r
}

// Offsets Specifies if term offsets should be returned. Applies to all returned
// documents unless otherwise specified in body "params" or "docs".
// API name: offsets
func (r *Mtermvectors) Offsets(b bool) *Mtermvectors {
	r.values.Set("offsets", strconv.FormatBool(b))

	return r
}

// Payloads Specifies if term payloads should be returned. Applies to all returned
// documents unless otherwise specified in body "params" or "docs".
// API name: payloads
func (r *Mtermvectors) Payloads(b bool) *Mtermvectors {
	r.values.Set("payloads", strconv.FormatBool(b))

	return r
}

// Positions Specifies if term positions should be returned. Applies to all returned
// documents unless otherwise specified in body "params" or "docs".
// API name: positions
func (r *Mtermvectors) Positions(b bool) *Mtermvectors {
	r.values.Set("positions", strconv.FormatBool(b))

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random) .Applies to all returned documents unless otherwise specified in body
// "params" or "docs".
// API name: preference
func (r *Mtermvectors) Preference(value string) *Mtermvectors {
	r.values.Set("preference", value)

	return r
}

// Realtime Specifies if requests are real-time as opposed to near-real-time (default:
// true).
// API name: realtime
func (r *Mtermvectors) Realtime(b bool) *Mtermvectors {
	r.values.Set("realtime", strconv.FormatBool(b))

	return r
}

// Routing Specific routing value. Applies to all returned documents unless otherwise
// specified in body "params" or "docs".
// API name: routing
func (r *Mtermvectors) Routing(value string) *Mtermvectors {
	r.values.Set("routing", value)

	return r
}

// TermStatistics Specifies if total term frequency and document frequency should be returned.
// Applies to all returned documents unless otherwise specified in body "params"
// or "docs".
// API name: term_statistics
func (r *Mtermvectors) TermStatistics(b bool) *Mtermvectors {
	r.values.Set("term_statistics", strconv.FormatBool(b))

	return r
}

// Version Explicit version number for concurrency control
// API name: version
func (r *Mtermvectors) Version(value string) *Mtermvectors {
	r.values.Set("version", value)

	return r
}

// VersionType Specific version type
// API name: version_type
func (r *Mtermvectors) VersionType(enum versiontype.VersionType) *Mtermvectors {
	r.values.Set("version_type", enum.String())

	return r
}

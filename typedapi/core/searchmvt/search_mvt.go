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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Searches a vector tile for geospatial values. Returns results as a binary
// Mapbox vector tile.
package searchmvt

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gridaggregationtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gridtype"
)

const (
	indexMask = iota + 1

	fieldMask

	zoomMask

	xMask

	yMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SearchMvt struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	index string
	field string
	zoom  string
	x     string
	y     string
}

// NewSearchMvt type alias for index.
type NewSearchMvt func(index, field, zoom, x, y string) *SearchMvt

// NewSearchMvtFunc returns a new instance of SearchMvt with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSearchMvtFunc(tp elastictransport.Interface) NewSearchMvt {
	return func(index, field, zoom, x, y string) *SearchMvt {
		n := New(tp)

		n.Index(index)

		n.Field(field)

		n.Zoom(zoom)

		n.X(x)

		n.Y(y)

		return n
	}
}

// Searches a vector tile for geospatial values. Returns results as a binary
// Mapbox vector tile.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-vector-tile-api.html
func New(tp elastictransport.Interface) *SearchMvt {
	r := &SearchMvt{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *SearchMvt) Raw(raw io.Reader) *SearchMvt {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SearchMvt) Request(req *Request) *SearchMvt {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SearchMvt) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for SearchMvt: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|fieldMask|zoomMask|xMask|yMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_mvt")
		path.WriteString("/")

		path.WriteString(r.field)
		path.WriteString("/")

		path.WriteString(r.zoom)
		path.WriteString("/")

		path.WriteString(r.x)
		path.WriteString("/")

		path.WriteString(r.y)

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
			req.Header.Set("Content-Type", "application/json")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.mapbox-vector-tile")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r SearchMvt) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the SearchMvt query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a searchmvt.Response
func (r SearchMvt) Do(ctx context.Context) (Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
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

// Header set a key, value pair in the SearchMvt headers map.
func (r *SearchMvt) Header(key, value string) *SearchMvt {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, or aliases to search
// API Name: index
func (r *SearchMvt) Index(v string) *SearchMvt {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Field Field containing geospatial data to return
// API Name: field
func (r *SearchMvt) Field(v string) *SearchMvt {
	r.paramSet |= fieldMask
	r.field = v

	return r
}

// Zoom Zoom level for the vector tile to search
// API Name: zoom
func (r *SearchMvt) Zoom(v string) *SearchMvt {
	r.paramSet |= zoomMask
	r.zoom = v

	return r
}

// X X coordinate for the vector tile to search
// API Name: x
func (r *SearchMvt) X(v string) *SearchMvt {
	r.paramSet |= xMask
	r.x = v

	return r
}

// Y Y coordinate for the vector tile to search
// API Name: y
func (r *SearchMvt) Y(v string) *SearchMvt {
	r.paramSet |= yMask
	r.y = v

	return r
}

// ExactBounds If false, the meta layer’s feature is the bounding box of the tile.
// If true, the meta layer’s feature is a bounding box resulting from a
// geo_bounds aggregation. The aggregation runs on <field> values that intersect
// the <zoom>/<x>/<y> tile with wrap_longitude set to false. The resulting
// bounding box may be larger than the vector tile.
// API name: exact_bounds
func (r *SearchMvt) ExactBounds(b bool) *SearchMvt {
	r.values.Set("exact_bounds", strconv.FormatBool(b))

	return r
}

// Extent Size, in pixels, of a side of the tile. Vector tiles are square with equal
// sides.
// API name: extent
func (r *SearchMvt) Extent(i int) *SearchMvt {
	r.values.Set("extent", strconv.Itoa(i))

	return r
}

// GridAgg Aggregation used to create a grid for `field`.
// API name: grid_agg
func (r *SearchMvt) GridAgg(enum gridaggregationtype.GridAggregationType) *SearchMvt {
	r.values.Set("grid_agg", enum.String())

	return r
}

// GridPrecision Additional zoom levels available through the aggs layer. For example, if
// <zoom> is 7
// and grid_precision is 8, you can zoom in up to level 15. Accepts 0-8. If 0,
// results
// don’t include the aggs layer.
// API name: grid_precision
func (r *SearchMvt) GridPrecision(i int) *SearchMvt {
	r.values.Set("grid_precision", strconv.Itoa(i))

	return r
}

// GridType Determines the geometry type for features in the aggs layer. In the aggs
// layer,
// each feature represents a geotile_grid cell. If 'grid' each feature is a
// Polygon
// of the cells bounding box. If 'point' each feature is a Point that is the
// centroid
// of the cell.
// API name: grid_type
func (r *SearchMvt) GridType(enum gridtype.GridType) *SearchMvt {
	r.values.Set("grid_type", enum.String())

	return r
}

// Size Maximum number of features to return in the hits layer. Accepts 0-10000.
// If 0, results don’t include the hits layer.
// API name: size
func (r *SearchMvt) Size(i int) *SearchMvt {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// WithLabels If `true`, the hits and aggs layers will contain additional point features
// representing
// suggested label positions for the original features.
// API name: with_labels
func (r *SearchMvt) WithLabels(b bool) *SearchMvt {
	r.values.Set("with_labels", strconv.FormatBool(b))

	return r
}

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

// Search a vector tile.
//
// Search a vector tile for geospatial values.
// Before using this API, you should be familiar with the Mapbox vector tile
// specification.
// The API returns results as a binary mapbox vector tile.
//
// Internally, Elasticsearch translates a vector tile search API request into a
// search containing:
//
// * A `geo_bounding_box` query on the `<field>`. The query uses the
// `<zoom>/<x>/<y>` tile as a bounding box.
// * A `geotile_grid` or `geohex_grid` aggregation on the `<field>`. The
// `grid_agg` parameter determines the aggregation type. The aggregation uses
// the `<zoom>/<x>/<y>` tile as a bounding box.
// * Optionally, a `geo_bounds` aggregation on the `<field>`. The search only
// includes this aggregation if the `exact_bounds` parameter is `true`.
// * If the optional parameter `with_labels` is `true`, the internal search will
// include a dynamic runtime field that calls the `getLabelPosition` function of
// the geometry doc value. This enables the generation of new point features
// containing suggested geometry labels, so that, for example, multi-polygons
// will have only one label.
//
// For example, Elasticsearch may translate a vector tile search API request
// with a `grid_agg` argument of `geotile` and an `exact_bounds` argument of
// `true` into the following search
//
// ```
// GET my-index/_search
//
//	{
//	  "size": 10000,
//	  "query": {
//	    "geo_bounding_box": {
//	      "my-geo-field": {
//	        "top_left": {
//	          "lat": -40.979898069620134,
//	          "lon": -45
//	        },
//	        "bottom_right": {
//	          "lat": -66.51326044311186,
//	          "lon": 0
//	        }
//	      }
//	    }
//	  },
//	  "aggregations": {
//	    "grid": {
//	      "geotile_grid": {
//	        "field": "my-geo-field",
//	        "precision": 11,
//	        "size": 65536,
//	        "bounds": {
//	          "top_left": {
//	            "lat": -40.979898069620134,
//	            "lon": -45
//	          },
//	          "bottom_right": {
//	            "lat": -66.51326044311186,
//	            "lon": 0
//	          }
//	        }
//	      }
//	    },
//	    "bounds": {
//	      "geo_bounds": {
//	        "field": "my-geo-field",
//	        "wrap_longitude": false
//	      }
//	    }
//	  }
//	}
//
// ```
//
// The API returns results as a binary Mapbox vector tile.
// Mapbox vector tiles are encoded as Google Protobufs (PBF). By default, the
// tile contains three layers:
//
// * A `hits` layer containing a feature for each `<field>` value matching the
// `geo_bounding_box` query.
// * An `aggs` layer containing a feature for each cell of the `geotile_grid` or
// `geohex_grid`. The layer only contains features for cells with matching data.
// * A meta layer containing:
//   - A feature containing a bounding box. By default, this is the bounding box
//
// of the tile.
//   - Value ranges for any sub-aggregations on the `geotile_grid` or
//
// `geohex_grid`.
//   - Metadata for the search.
//
// The API only returns features that can display at its zoom level.
// For example, if a polygon feature has no area at its zoom level, the API
// omits it.
// The API returns errors as UTF-8 encoded JSON.
//
// IMPORTANT: You can specify several options for this API as either a query
// parameter or request body parameter.
// If you specify both parameters, the query parameter takes precedence.
//
// **Grid precision for geotile**
//
// For a `grid_agg` of `geotile`, you can use cells in the `aggs` layer as tiles
// for lower zoom levels.
// `grid_precision` represents the additional zoom levels available through
// these cells. The final precision is computed by as follows: `<zoom> +
// grid_precision`.
// For example, if `<zoom>` is 7 and `grid_precision` is 8, then the
// `geotile_grid` aggregation will use a precision of 15.
// The maximum final precision is 29.
// The `grid_precision` also determines the number of cells for the grid as
// follows: `(2^grid_precision) x (2^grid_precision)`.
// For example, a value of 8 divides the tile into a grid of 256 x 256 cells.
// The `aggs` layer only contains features for cells with matching data.
//
// **Grid precision for geohex**
//
// For a `grid_agg` of `geohex`, Elasticsearch uses `<zoom>` and
// `grid_precision` to calculate a final precision as follows: `<zoom> +
// grid_precision`.
//
// This precision determines the H3 resolution of the hexagonal cells produced
// by the `geohex` aggregation.
// The following table maps the H3 resolution for each precision.
// For example, if `<zoom>` is 3 and `grid_precision` is 3, the precision is 6.
// At a precision of 6, hexagonal cells have an H3 resolution of 2.
// If `<zoom>` is 3 and `grid_precision` is 4, the precision is 7.
// At a precision of 7, hexagonal cells have an H3 resolution of 3.
//
// | Precision | Unique tile bins | H3 resolution | Unique hex bins |	Ratio |
// | --------- | ---------------- | ------------- | ----------------| ----- |
// | 1  | 4                  | 0  | 122             | 30.5           |
// | 2  | 16                 | 0  | 122             | 7.625          |
// | 3  | 64                 | 1  | 842             | 13.15625       |
// | 4  | 256                | 1  | 842             | 3.2890625      |
// | 5  | 1024               | 2  | 5882            | 5.744140625    |
// | 6  | 4096               | 2  | 5882            | 1.436035156    |
// | 7  | 16384              | 3  | 41162           | 2.512329102    |
// | 8  | 65536              | 3  | 41162           | 0.6280822754   |
// | 9  | 262144             | 4  | 288122          | 1.099098206    |
// | 10 | 1048576            | 4  | 288122          | 0.2747745514   |
// | 11 | 4194304            | 5  | 2016842         | 0.4808526039   |
// | 12 | 16777216           | 6  | 14117882        | 0.8414913416   |
// | 13 | 67108864           | 6  | 14117882        | 0.2103728354   |
// | 14 | 268435456          | 7  | 98825162        | 0.3681524172   |
// | 15 | 1073741824         | 8  | 691776122       | 0.644266719    |
// | 16 | 4294967296         | 8  | 691776122       | 0.1610666797   |
// | 17 | 17179869184        | 9  | 4842432842      | 0.2818666889   |
// | 18 | 68719476736        | 10 | 33897029882     | 0.4932667053   |
// | 19 | 274877906944       | 11 | 237279209162    | 0.8632167343   |
// | 20 | 1099511627776      | 11 | 237279209162    | 0.2158041836   |
// | 21 | 4398046511104      | 12 | 1660954464122   | 0.3776573213   |
// | 22 | 17592186044416     | 13 | 11626681248842  | 0.6609003122   |
// | 23 | 70368744177664     | 13 | 11626681248842  | 0.165225078    |
// | 24 | 281474976710656    | 14 | 81386768741882  | 0.2891438866   |
// | 25 | 1125899906842620   | 15 | 569707381193162 | 0.5060018015   |
// | 26 | 4503599627370500   | 15 | 569707381193162 | 0.1265004504   |
// | 27 | 18014398509482000  | 15 | 569707381193162 | 0.03162511259  |
// | 28 | 72057594037927900  | 15 | 569707381193162 | 0.007906278149 |
// | 29 | 288230376151712000 | 15 | 569707381193162 | 0.001976569537 |
//
// Hexagonal cells don't align perfectly on a vector tile.
// Some cells may intersect more than one vector tile.
// To compute the H3 resolution for each precision, Elasticsearch compares the
// average density of hexagonal bins at each resolution with the average density
// of tile bins at each zoom level.
// Elasticsearch uses the H3 resolution that is closest to the corresponding
// geotile density.
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

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string
	field string
	zoom  string
	x     string
	y     string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewSearchMvt type alias for index.
type NewSearchMvt func(index, field, zoom, x, y string) *SearchMvt

// NewSearchMvtFunc returns a new instance of SearchMvt with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSearchMvtFunc(tp elastictransport.Interface) NewSearchMvt {
	return func(index, field, zoom, x, y string) *SearchMvt {
		n := New(tp)

		n._index(index)

		n._field(field)

		n._zoom(zoom)

		n._x(x)

		n._y(y)

		return n
	}
}

// Search a vector tile.
//
// Search a vector tile for geospatial values.
// Before using this API, you should be familiar with the Mapbox vector tile
// specification.
// The API returns results as a binary mapbox vector tile.
//
// Internally, Elasticsearch translates a vector tile search API request into a
// search containing:
//
// * A `geo_bounding_box` query on the `<field>`. The query uses the
// `<zoom>/<x>/<y>` tile as a bounding box.
// * A `geotile_grid` or `geohex_grid` aggregation on the `<field>`. The
// `grid_agg` parameter determines the aggregation type. The aggregation uses
// the `<zoom>/<x>/<y>` tile as a bounding box.
// * Optionally, a `geo_bounds` aggregation on the `<field>`. The search only
// includes this aggregation if the `exact_bounds` parameter is `true`.
// * If the optional parameter `with_labels` is `true`, the internal search will
// include a dynamic runtime field that calls the `getLabelPosition` function of
// the geometry doc value. This enables the generation of new point features
// containing suggested geometry labels, so that, for example, multi-polygons
// will have only one label.
//
// For example, Elasticsearch may translate a vector tile search API request
// with a `grid_agg` argument of `geotile` and an `exact_bounds` argument of
// `true` into the following search
//
// ```
// GET my-index/_search
//
//	{
//	  "size": 10000,
//	  "query": {
//	    "geo_bounding_box": {
//	      "my-geo-field": {
//	        "top_left": {
//	          "lat": -40.979898069620134,
//	          "lon": -45
//	        },
//	        "bottom_right": {
//	          "lat": -66.51326044311186,
//	          "lon": 0
//	        }
//	      }
//	    }
//	  },
//	  "aggregations": {
//	    "grid": {
//	      "geotile_grid": {
//	        "field": "my-geo-field",
//	        "precision": 11,
//	        "size": 65536,
//	        "bounds": {
//	          "top_left": {
//	            "lat": -40.979898069620134,
//	            "lon": -45
//	          },
//	          "bottom_right": {
//	            "lat": -66.51326044311186,
//	            "lon": 0
//	          }
//	        }
//	      }
//	    },
//	    "bounds": {
//	      "geo_bounds": {
//	        "field": "my-geo-field",
//	        "wrap_longitude": false
//	      }
//	    }
//	  }
//	}
//
// ```
//
// The API returns results as a binary Mapbox vector tile.
// Mapbox vector tiles are encoded as Google Protobufs (PBF). By default, the
// tile contains three layers:
//
// * A `hits` layer containing a feature for each `<field>` value matching the
// `geo_bounding_box` query.
// * An `aggs` layer containing a feature for each cell of the `geotile_grid` or
// `geohex_grid`. The layer only contains features for cells with matching data.
// * A meta layer containing:
//   - A feature containing a bounding box. By default, this is the bounding box
//
// of the tile.
//   - Value ranges for any sub-aggregations on the `geotile_grid` or
//
// `geohex_grid`.
//   - Metadata for the search.
//
// The API only returns features that can display at its zoom level.
// For example, if a polygon feature has no area at its zoom level, the API
// omits it.
// The API returns errors as UTF-8 encoded JSON.
//
// IMPORTANT: You can specify several options for this API as either a query
// parameter or request body parameter.
// If you specify both parameters, the query parameter takes precedence.
//
// **Grid precision for geotile**
//
// For a `grid_agg` of `geotile`, you can use cells in the `aggs` layer as tiles
// for lower zoom levels.
// `grid_precision` represents the additional zoom levels available through
// these cells. The final precision is computed by as follows: `<zoom> +
// grid_precision`.
// For example, if `<zoom>` is 7 and `grid_precision` is 8, then the
// `geotile_grid` aggregation will use a precision of 15.
// The maximum final precision is 29.
// The `grid_precision` also determines the number of cells for the grid as
// follows: `(2^grid_precision) x (2^grid_precision)`.
// For example, a value of 8 divides the tile into a grid of 256 x 256 cells.
// The `aggs` layer only contains features for cells with matching data.
//
// **Grid precision for geohex**
//
// For a `grid_agg` of `geohex`, Elasticsearch uses `<zoom>` and
// `grid_precision` to calculate a final precision as follows: `<zoom> +
// grid_precision`.
//
// This precision determines the H3 resolution of the hexagonal cells produced
// by the `geohex` aggregation.
// The following table maps the H3 resolution for each precision.
// For example, if `<zoom>` is 3 and `grid_precision` is 3, the precision is 6.
// At a precision of 6, hexagonal cells have an H3 resolution of 2.
// If `<zoom>` is 3 and `grid_precision` is 4, the precision is 7.
// At a precision of 7, hexagonal cells have an H3 resolution of 3.
//
// | Precision | Unique tile bins | H3 resolution | Unique hex bins |	Ratio |
// | --------- | ---------------- | ------------- | ----------------| ----- |
// | 1  | 4                  | 0  | 122             | 30.5           |
// | 2  | 16                 | 0  | 122             | 7.625          |
// | 3  | 64                 | 1  | 842             | 13.15625       |
// | 4  | 256                | 1  | 842             | 3.2890625      |
// | 5  | 1024               | 2  | 5882            | 5.744140625    |
// | 6  | 4096               | 2  | 5882            | 1.436035156    |
// | 7  | 16384              | 3  | 41162           | 2.512329102    |
// | 8  | 65536              | 3  | 41162           | 0.6280822754   |
// | 9  | 262144             | 4  | 288122          | 1.099098206    |
// | 10 | 1048576            | 4  | 288122          | 0.2747745514   |
// | 11 | 4194304            | 5  | 2016842         | 0.4808526039   |
// | 12 | 16777216           | 6  | 14117882        | 0.8414913416   |
// | 13 | 67108864           | 6  | 14117882        | 0.2103728354   |
// | 14 | 268435456          | 7  | 98825162        | 0.3681524172   |
// | 15 | 1073741824         | 8  | 691776122       | 0.644266719    |
// | 16 | 4294967296         | 8  | 691776122       | 0.1610666797   |
// | 17 | 17179869184        | 9  | 4842432842      | 0.2818666889   |
// | 18 | 68719476736        | 10 | 33897029882     | 0.4932667053   |
// | 19 | 274877906944       | 11 | 237279209162    | 0.8632167343   |
// | 20 | 1099511627776      | 11 | 237279209162    | 0.2158041836   |
// | 21 | 4398046511104      | 12 | 1660954464122   | 0.3776573213   |
// | 22 | 17592186044416     | 13 | 11626681248842  | 0.6609003122   |
// | 23 | 70368744177664     | 13 | 11626681248842  | 0.165225078    |
// | 24 | 281474976710656    | 14 | 81386768741882  | 0.2891438866   |
// | 25 | 1125899906842620   | 15 | 569707381193162 | 0.5060018015   |
// | 26 | 4503599627370500   | 15 | 569707381193162 | 0.1265004504   |
// | 27 | 18014398509482000  | 15 | 569707381193162 | 0.03162511259  |
// | 28 | 72057594037927900  | 15 | 569707381193162 | 0.007906278149 |
// | 29 | 288230376151712000 | 15 | 569707381193162 | 0.001976569537 |
//
// Hexagonal cells don't align perfectly on a vector tile.
// Some cells may intersect more than one vector tile.
// To compute the H3 resolution for each precision, Elasticsearch compares the
// average density of hexagonal bins at each resolution with the average density
// of tile bins at each zoom level.
// Elasticsearch uses the H3 resolution that is closest to the corresponding
// geotile density.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
func New(tp elastictransport.Interface) *SearchMvt {
	r := &SearchMvt{
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
			return nil, fmt.Errorf("could not serialise request for SearchMvt: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|fieldMask|zoomMask|xMask|yMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_mvt")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "field", r.field)
		}
		path.WriteString(r.field)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "zoom", r.zoom)
		}
		path.WriteString(r.zoom)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "x", r.x)
		}
		path.WriteString(r.x)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "y", r.y)
		}
		path.WriteString(r.y)

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
func (r SearchMvt) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "search_mvt")
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
		instrument.BeforeRequest(req, "search_mvt")
		if reader := instrument.RecordRequestBody(ctx, "search_mvt", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "search_mvt")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the SearchMvt query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a searchmvt.Response
func (r SearchMvt) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "search_mvt")
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
		response, err = io.ReadAll(res.Body)
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

// Header set a key, value pair in the SearchMvt headers map.
func (r *SearchMvt) Header(key, value string) *SearchMvt {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, or aliases to search
// API Name: index
func (r *SearchMvt) _index(index string) *SearchMvt {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Field Field containing geospatial data to return
// API Name: field
func (r *SearchMvt) _field(field string) *SearchMvt {
	r.paramSet |= fieldMask
	r.field = field

	return r
}

// Zoom Zoom level for the vector tile to search
// API Name: zoom
func (r *SearchMvt) _zoom(zoom string) *SearchMvt {
	r.paramSet |= zoomMask
	r.zoom = zoom

	return r
}

// X X coordinate for the vector tile to search
// API Name: x
func (r *SearchMvt) _x(x string) *SearchMvt {
	r.paramSet |= xMask
	r.x = x

	return r
}

// Y Y coordinate for the vector tile to search
// API Name: y
func (r *SearchMvt) _y(y string) *SearchMvt {
	r.paramSet |= yMask
	r.y = y

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *SearchMvt) ErrorTrace(errortrace bool) *SearchMvt {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *SearchMvt) FilterPath(filterpaths ...string) *SearchMvt {
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
func (r *SearchMvt) Human(human bool) *SearchMvt {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *SearchMvt) Pretty(pretty bool) *SearchMvt {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Aggs Sub-aggregations for the geotile_grid.
//
// It supports the following aggregation types:
//
// - `avg`
// - `boxplot`
// - `cardinality`
// - `extended stats`
// - `max`
// - `median absolute deviation`
// - `min`
// - `percentile`
// - `percentile-rank`
// - `stats`
// - `sum`
// - `value count`
//
// The aggregation names can't start with `_mvt_`. The `_mvt_` prefix is
// reserved for internal aggregations.
// API name: aggs
func (r *SearchMvt) Aggs(aggs map[string]types.Aggregations) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Aggs = aggs

	return r
}

// Buffer The size, in pixels, of a clipping buffer outside the tile. This allows
// renderers
// to avoid outline artifacts from geometries that extend past the extent of the
// tile.
// API name: buffer
func (r *SearchMvt) Buffer(buffer int) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Buffer = &buffer

	return r
}

// ExactBounds If `false`, the meta layer's feature is the bounding box of the tile.
// If `true`, the meta layer's feature is a bounding box resulting from a
// `geo_bounds` aggregation. The aggregation runs on <field> values that
// intersect
// the `<zoom>/<x>/<y>` tile with `wrap_longitude` set to `false`. The resulting
// bounding box may be larger than the vector tile.
// API name: exact_bounds
func (r *SearchMvt) ExactBounds(exactbounds bool) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ExactBounds = &exactbounds

	return r
}

// Extent The size, in pixels, of a side of the tile. Vector tiles are square with
// equal sides.
// API name: extent
func (r *SearchMvt) Extent(extent int) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Extent = &extent

	return r
}

// Fields The fields to return in the `hits` layer.
// It supports wildcards (`*`).
// This parameter does not support fields with array values. Fields with array
// values may return inconsistent results.
// API name: fields
func (r *SearchMvt) Fields(fields ...string) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Fields = fields

	return r
}

// GridAgg The aggregation used to create a grid for the `field`.
// API name: grid_agg
func (r *SearchMvt) GridAgg(gridagg gridaggregationtype.GridAggregationType) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.GridAgg = &gridagg

	return r
}

// GridPrecision Additional zoom levels available through the aggs layer. For example, if
// `<zoom>` is `7`
// and `grid_precision` is `8`, you can zoom in up to level 15. Accepts 0-8. If
// 0, results
// don't include the aggs layer.
// API name: grid_precision
func (r *SearchMvt) GridPrecision(gridprecision int) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.GridPrecision = &gridprecision

	return r
}

// GridType Determines the geometry type for features in the aggs layer. In the aggs
// layer,
// each feature represents a `geotile_grid` cell. If `grid, each feature is a
// polygon
// of the cells bounding box. If `point`, each feature is a Point that is the
// centroid
// of the cell.
// API name: grid_type
func (r *SearchMvt) GridType(gridtype gridtype.GridType) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.GridType = &gridtype

	return r
}

// Query The query DSL used to filter documents for the search.
// API name: query
func (r *SearchMvt) Query(query *types.Query) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}

// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
// precedence over mapped fields with the same name.
// API name: runtime_mappings
func (r *SearchMvt) RuntimeMappings(runtimefields types.RuntimeFields) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RuntimeMappings = runtimefields

	return r
}

// Size The maximum number of features to return in the hits layer. Accepts 0-10000.
// If 0, results don't include the hits layer.
// API name: size
func (r *SearchMvt) Size(size int) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Size = &size

	return r
}

// Sort Sort the features in the hits layer. By default, the API calculates a
// bounding
// box for each feature. It sorts features based on this box's diagonal length,
// from longest to shortest.
// API name: sort
func (r *SearchMvt) Sort(sorts ...types.SortCombinations) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Sort = sorts

	return r
}

// TrackTotalHits The number of hits matching the query to count accurately. If `true`, the
// exact number
// of hits is returned at the cost of some performance. If `false`, the response
// does
// not include the total number of hits matching the query.
// API name: track_total_hits
func (r *SearchMvt) TrackTotalHits(trackhits types.TrackHits) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.TrackTotalHits = trackhits

	return r
}

// WithLabels If `true`, the hits and aggs layers will contain additional point features
// representing
// suggested label positions for the original features.
//
// * `Point` and `MultiPoint` features will have one of the points selected.
// * `Polygon` and `MultiPolygon` features will have a single point generated,
// either the centroid, if it is within the polygon, or another point within the
// polygon selected from the sorted triangle-tree.
// * `LineString` features will likewise provide a roughly central point
// selected from the triangle-tree.
// * The aggregation results will provide one central point for each aggregation
// bucket.
//
// All attributes from the original features will also be copied to the new
// label features.
// In addition, the new features will be distinguishable using the tag
// `_mvt_label_position`.
// API name: with_labels
func (r *SearchMvt) WithLabels(withlabels bool) *SearchMvt {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.WithLabels = &withlabels

	return r
}

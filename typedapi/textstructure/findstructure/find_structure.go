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

// Finds the structure of a text file. The text file must contain data that is
// suitable to be ingested into Elasticsearch.
package findstructure

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

type FindStructure struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int
}

// NewFindStructure type alias for index.
type NewFindStructure func() *FindStructure

// NewFindStructureFunc returns a new instance of FindStructure with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewFindStructureFunc(tp elastictransport.Interface) NewFindStructure {
	return func() *FindStructure {
		n := New(tp)

		return n
	}
}

// Finds the structure of a text file. The text file must contain data that is
// suitable to be ingested into Elasticsearch.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-structure.html
func New(tp elastictransport.Interface) *FindStructure {
	r := &FindStructure{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *FindStructure) Raw(raw io.Reader) *FindStructure {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *FindStructure) Request(req *Request) *FindStructure {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *FindStructure) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		for _, elem := range *r.req {
			data, err := json.Marshal(elem)
			if err != nil {
				return nil, err
			}
			r.buf.Write(data)
			r.buf.Write([]byte("\n"))
		}

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for FindStructure: %w", err)
		}

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_text_structure")
		path.WriteString("/")
		path.WriteString("find_structure")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+x-ndjson;compatible-with=8")
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
func (r FindStructure) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the FindStructure query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a findstructure.Response
func (r FindStructure) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the FindStructure headers map.
func (r *FindStructure) Header(key, value string) *FindStructure {
	r.headers.Set(key, value)

	return r
}

// Charset The text’s character set. It must be a character set that is supported by the
// JVM that Elasticsearch uses. For example, UTF-8, UTF-16LE, windows-1252, or
// EUC-JP. If this parameter is not specified, the structure finder chooses an
// appropriate character set.
// API name: charset
func (r *FindStructure) Charset(charset string) *FindStructure {
	r.values.Set("charset", charset)

	return r
}

// ColumnNames If you have set format to delimited, you can specify the column names in a
// comma-separated list. If this parameter is not specified, the structure
// finder uses the column names from the header row of the text. If the text
// does not have a header role, columns are named "column1", "column2",
// "column3", etc.
// API name: column_names
func (r *FindStructure) ColumnNames(columnnames string) *FindStructure {
	r.values.Set("column_names", columnnames)

	return r
}

// Delimiter If you have set format to delimited, you can specify the character used to
// delimit the values in each row. Only a single character is supported; the
// delimiter cannot have multiple characters. By default, the API considers the
// following possibilities: comma, tab, semi-colon, and pipe (|). In this
// default scenario, all rows must have the same number of fields for the
// delimited format to be detected. If you specify a delimiter, up to 10% of the
// rows can have a different number of columns than the first row.
// API name: delimiter
func (r *FindStructure) Delimiter(delimiter string) *FindStructure {
	r.values.Set("delimiter", delimiter)

	return r
}

// Explain If this parameter is set to true, the response includes a field named
// explanation, which is an array of strings that indicate how the structure
// finder produced its result.
// API name: explain
func (r *FindStructure) Explain(explain bool) *FindStructure {
	r.values.Set("explain", strconv.FormatBool(explain))

	return r
}

// Format The high level structure of the text. Valid values are ndjson, xml,
// delimited, and semi_structured_text. By default, the API chooses the format.
// In this default scenario, all rows must have the same number of fields for a
// delimited format to be detected. If the format is set to delimited and the
// delimiter is not set, however, the API tolerates up to 5% of rows that have a
// different number of columns than the first row.
// API name: format
func (r *FindStructure) Format(format string) *FindStructure {
	r.values.Set("format", format)

	return r
}

// GrokPattern If you have set format to semi_structured_text, you can specify a Grok
// pattern that is used to extract fields from every message in the text. The
// name of the timestamp field in the Grok pattern must match what is specified
// in the timestamp_field parameter. If that parameter is not specified, the
// name of the timestamp field in the Grok pattern must match "timestamp". If
// grok_pattern is not specified, the structure finder creates a Grok pattern.
// API name: grok_pattern
func (r *FindStructure) GrokPattern(grokpattern string) *FindStructure {
	r.values.Set("grok_pattern", grokpattern)

	return r
}

// HasHeaderRow If you have set format to delimited, you can use this parameter to indicate
// whether the column names are in the first row of the text. If this parameter
// is not specified, the structure finder guesses based on the similarity of the
// first row of the text to other rows.
// API name: has_header_row
func (r *FindStructure) HasHeaderRow(hasheaderrow bool) *FindStructure {
	r.values.Set("has_header_row", strconv.FormatBool(hasheaderrow))

	return r
}

// LineMergeSizeLimit The maximum number of characters in a message when lines are merged to form
// messages while analyzing semi-structured text. If you have extremely long
// messages you may need to increase this, but be aware that this may lead to
// very long processing times if the way to group lines into messages is
// misdetected.
// API name: line_merge_size_limit
func (r *FindStructure) LineMergeSizeLimit(linemergesizelimit string) *FindStructure {
	r.values.Set("line_merge_size_limit", linemergesizelimit)

	return r
}

// LinesToSample The number of lines to include in the structural analysis, starting from the
// beginning of the text. The minimum is 2; If the value of this parameter is
// greater than the number of lines in the text, the analysis proceeds (as long
// as there are at least two lines in the text) for all of the lines.
// API name: lines_to_sample
func (r *FindStructure) LinesToSample(linestosample string) *FindStructure {
	r.values.Set("lines_to_sample", linestosample)

	return r
}

// Quote If you have set format to delimited, you can specify the character used to
// quote the values in each row if they contain newlines or the delimiter
// character. Only a single character is supported. If this parameter is not
// specified, the default value is a double quote ("). If your delimited text
// format does not use quoting, a workaround is to set this argument to a
// character that does not appear anywhere in the sample.
// API name: quote
func (r *FindStructure) Quote(quote string) *FindStructure {
	r.values.Set("quote", quote)

	return r
}

// ShouldTrimFields If you have set format to delimited, you can specify whether values between
// delimiters should have whitespace trimmed from them. If this parameter is not
// specified and the delimiter is pipe (|), the default value is true.
// Otherwise, the default value is false.
// API name: should_trim_fields
func (r *FindStructure) ShouldTrimFields(shouldtrimfields bool) *FindStructure {
	r.values.Set("should_trim_fields", strconv.FormatBool(shouldtrimfields))

	return r
}

// Timeout Sets the maximum amount of time that the structure analysis make take. If the
// analysis is still running when the timeout expires then it will be aborted.
// API name: timeout
func (r *FindStructure) Timeout(duration string) *FindStructure {
	r.values.Set("timeout", duration)

	return r
}

// TimestampField Optional parameter to specify the timestamp field in the file
// API name: timestamp_field
func (r *FindStructure) TimestampField(field string) *FindStructure {
	r.values.Set("timestamp_field", field)

	return r
}

// TimestampFormat The Java time format of the timestamp field in the text.
// API name: timestamp_format
func (r *FindStructure) TimestampFormat(timestampformat string) *FindStructure {
	r.values.Set("timestamp_format", timestampformat)

	return r
}

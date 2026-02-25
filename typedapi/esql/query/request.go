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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package query

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package query
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/esql/query/QueryRequest.ts#L28-L135
type Request struct {
	// Columnar By default, ES|QL returns results as rows. For example, FROM returns each
	// individual document as one row. For the JSON, YAML, CBOR and smile formats,
	// ES|QL can return the results in a columnar fashion where one row represents
	// all the values of a certain column in the results.
	Columnar *bool `json:"columnar,omitempty"`
	// Filter Specify a Query DSL query in the filter parameter to filter the set of
	// documents that an ES|QL query runs on.
	Filter *types.Query `json:"filter,omitempty"`
	// IncludeCcsMetadata When set to `true` and performing a cross-cluster/cross-project query, the
	// response will include an extra `_clusters` object with information about the
	// clusters that participated in the search along with info such as shards
	// count.
	IncludeCcsMetadata *bool `json:"include_ccs_metadata,omitempty"`
	// IncludeExecutionMetadata When set to `true`, the response will include an extra `_clusters` object
	// with information about the clusters that participated in the search along
	// with info such as shards count. This is similar to `include_ccs_metadata`,
	// but it also returns metadata when the query is not CCS/CPS
	IncludeExecutionMetadata *bool `json:"include_execution_metadata,omitempty"`
	// Locale Returns results (especially dates) formatted per the conventions of the
	// locale.
	Locale *string `json:"locale,omitempty"`
	// Params To avoid any attempts of hacking or code injection, extract the values in a
	// separate list of parameters. Use question mark placeholders (?) in the query
	// string for each of the parameters.
	Params types.ESQLParams `json:"params,omitempty"`
	// Profile If provided and `true` the response will include an extra `profile` object
	// with information on how the query was executed. This information is for human
	// debugging and its format can change at any time but it can give some insight
	// into the performance of each part of the query.
	Profile *bool `json:"profile,omitempty"`
	// Query The ES|QL query API accepts an ES|QL query string in the query parameter,
	// runs it, and returns the results.
	Query string `json:"query"`
	// Tables Tables to use with the LOOKUP operation. The top level key is the table name
	// and the next level key is the column name.
	Tables map[string]map[string]types.TableValuesContainer `json:"tables,omitempty"`
	// TimeZone Sets the default timezone of the query.
	TimeZone *string `json:"time_zone,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Tables: make(map[string]map[string]types.TableValuesContainer, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Query request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "columnar":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Columnar", err)
				}
				s.Columnar = &value
			case bool:
				s.Columnar = &v
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "include_ccs_metadata":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeCcsMetadata", err)
				}
				s.IncludeCcsMetadata = &value
			case bool:
				s.IncludeCcsMetadata = &v
			}

		case "include_execution_metadata":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeExecutionMetadata", err)
				}
				s.IncludeExecutionMetadata = &value
			case bool:
				s.IncludeExecutionMetadata = &v
			}

		case "locale":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Locale", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Locale = &o

		case "params":
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "profile":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Profile", err)
				}
				s.Profile = &value
			case bool:
				s.Profile = &v
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = o

		case "tables":
			if s.Tables == nil {
				s.Tables = make(map[string]map[string]types.TableValuesContainer, 0)
			}
			if err := dec.Decode(&s.Tables); err != nil {
				return fmt.Errorf("%s | %w", "Tables", err)
			}

		case "time_zone":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TimeZone", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TimeZone = &o

		}
	}
	return nil
}

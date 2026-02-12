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

package types

// DataStreamOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/indices/_types/DataStreamOptions.ts#L25-L34
type DataStreamOptions struct {
	// FailureStore If defined, it specifies configuration for the failure store of this data
	// stream.
	FailureStore *DataStreamFailureStore `json:"failure_store,omitempty"`
}

// NewDataStreamOptions returns a DataStreamOptions.
func NewDataStreamOptions() *DataStreamOptions {
	r := &DataStreamOptions{}

	return r
}

type DataStreamOptionsVariant interface {
	DataStreamOptionsCaster() *DataStreamOptions
}

func (s *DataStreamOptions) DataStreamOptionsCaster() *DataStreamOptions {
	return s
}

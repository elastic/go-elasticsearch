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

package types

// Groupings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/rollup/_types/Groupings.ts#L24-L40
type Groupings struct {
	// DateHistogram A date histogram group aggregates a date field into time-based buckets.
	// This group is mandatory; you currently cannot roll up documents without a
	// timestamp and a `date_histogram` group.
	DateHistogram *DateHistogramGrouping `json:"date_histogram,omitempty"`
	// Histogram The histogram group aggregates one or more numeric fields into numeric
	// histogram intervals.
	Histogram *HistogramGrouping `json:"histogram,omitempty"`
	// Terms The terms group can be used on keyword or numeric fields to allow bucketing
	// via the terms aggregation at a later point.
	// The indexer enumerates and stores all values of a field for each time-period.
	// This can be potentially costly for high-cardinality groups such as IP
	// addresses, especially if the time-bucket is particularly sparse.
	Terms *TermsGrouping `json:"terms,omitempty"`
}

// NewGroupings returns a Groupings.
func NewGroupings() *Groupings {
	r := &Groupings{}

	return r
}

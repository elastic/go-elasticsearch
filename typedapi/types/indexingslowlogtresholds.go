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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

// IndexingSlowlogTresholds type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L610-L617
type IndexingSlowlogTresholds struct {
	// Index The indexing slow log, similar in functionality to the search slow log. The
	// log file name ends with `_index_indexing_slowlog.json`.
	// Log and the thresholds are configured in the same way as the search slowlog.
	Index *SlowlogTresholdLevels `json:"index,omitempty"`
}

// NewIndexingSlowlogTresholds returns a IndexingSlowlogTresholds.
func NewIndexingSlowlogTresholds() *IndexingSlowlogTresholds {
	r := &IndexingSlowlogTresholds{}

	return r
}

type IndexingSlowlogTresholdsVariant interface {
	IndexingSlowlogTresholdsCaster() *IndexingSlowlogTresholds
}

func (s *IndexingSlowlogTresholds) IndexingSlowlogTresholdsCaster() *IndexingSlowlogTresholds {
	return s
}

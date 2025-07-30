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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package types

// IndexSettingResults type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e585438d116b00ff34643179e6286e402c0bcaaf/specification/indices/put_data_stream_settings/IndicesPutDataStreamSettingsResponse.ts#L57-L69
type IndexSettingResults struct {
	// AppliedToDataStreamAndBackingIndices The list of settings that were applied to the data stream and to all of its
	// backing indices. These settings will
	// also be applied to the write index the next time the data stream is rolled
	// over.
	AppliedToDataStreamAndBackingIndices []string `json:"applied_to_data_stream_and_backing_indices"`
	// AppliedToDataStreamOnly The list of settings that were applied to the data stream but not to backing
	// indices. These will be applied to
	// the write index the next time the data stream is rolled over.
	AppliedToDataStreamOnly []string                  `json:"applied_to_data_stream_only"`
	Errors                  []DataStreamSettingsError `json:"errors,omitempty"`
}

// NewIndexSettingResults returns a IndexSettingResults.
func NewIndexSettingResults() *IndexSettingResults {
	r := &IndexSettingResults{}

	return r
}

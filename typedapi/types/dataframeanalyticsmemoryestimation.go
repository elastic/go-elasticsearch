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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// DataframeAnalyticsMemoryEstimation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L70-L75
type DataframeAnalyticsMemoryEstimation struct {
	// ExpectedMemoryWithDisk Estimated memory usage under the assumption that overflowing to disk is
	// allowed during data frame analytics. expected_memory_with_disk is usually
	// smaller than expected_memory_without_disk as using disk allows to limit the
	// main memory needed to perform data frame analytics.
	ExpectedMemoryWithDisk string `json:"expected_memory_with_disk"`
	// ExpectedMemoryWithoutDisk Estimated memory usage under the assumption that the whole data frame
	// analytics should happen in memory (i.e. without overflowing to disk).
	ExpectedMemoryWithoutDisk string `json:"expected_memory_without_disk"`
}

// DataframeAnalyticsMemoryEstimationBuilder holds DataframeAnalyticsMemoryEstimation struct and provides a builder API.
type DataframeAnalyticsMemoryEstimationBuilder struct {
	v *DataframeAnalyticsMemoryEstimation
}

// NewDataframeAnalyticsMemoryEstimation provides a builder for the DataframeAnalyticsMemoryEstimation struct.
func NewDataframeAnalyticsMemoryEstimationBuilder() *DataframeAnalyticsMemoryEstimationBuilder {
	r := DataframeAnalyticsMemoryEstimationBuilder{
		&DataframeAnalyticsMemoryEstimation{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsMemoryEstimation struct
func (rb *DataframeAnalyticsMemoryEstimationBuilder) Build() DataframeAnalyticsMemoryEstimation {
	return *rb.v
}

// ExpectedMemoryWithDisk Estimated memory usage under the assumption that overflowing to disk is
// allowed during data frame analytics. expected_memory_with_disk is usually
// smaller than expected_memory_without_disk as using disk allows to limit the
// main memory needed to perform data frame analytics.

func (rb *DataframeAnalyticsMemoryEstimationBuilder) ExpectedMemoryWithDisk(expectedmemorywithdisk string) *DataframeAnalyticsMemoryEstimationBuilder {
	rb.v.ExpectedMemoryWithDisk = expectedmemorywithdisk
	return rb
}

// ExpectedMemoryWithoutDisk Estimated memory usage under the assumption that the whole data frame
// analytics should happen in memory (i.e. without overflowing to disk).

func (rb *DataframeAnalyticsMemoryEstimationBuilder) ExpectedMemoryWithoutDisk(expectedmemorywithoutdisk string) *DataframeAnalyticsMemoryEstimationBuilder {
	rb.v.ExpectedMemoryWithoutDisk = expectedmemorywithoutdisk
	return rb
}

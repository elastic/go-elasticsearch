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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dataframeAnalyticsDestination struct {
	v *types.DataframeAnalyticsDestination
}

func NewDataframeAnalyticsDestination() *_dataframeAnalyticsDestination {

	return &_dataframeAnalyticsDestination{v: types.NewDataframeAnalyticsDestination()}

}

// Defines the destination index to store the results of the data frame
// analytics job.
func (s *_dataframeAnalyticsDestination) Index(indexname string) *_dataframeAnalyticsDestination {

	s.v.Index = indexname

	return s
}

// Defines the name of the field in which to store the results of the analysis.
// Defaults to `ml`.
func (s *_dataframeAnalyticsDestination) ResultsField(field string) *_dataframeAnalyticsDestination {

	s.v.ResultsField = &field

	return s
}

func (s *_dataframeAnalyticsDestination) DataframeAnalyticsDestinationCaster() *types.DataframeAnalyticsDestination {
	return s.v
}

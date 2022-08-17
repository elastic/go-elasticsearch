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

// Defaults type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/info/types.ts#L24-L27
type Defaults struct {
	AnomalyDetectors AnomalyDetectors `json:"anomaly_detectors"`
	Datafeeds        Datafeeds        `json:"datafeeds"`
}

// DefaultsBuilder holds Defaults struct and provides a builder API.
type DefaultsBuilder struct {
	v *Defaults
}

// NewDefaults provides a builder for the Defaults struct.
func NewDefaultsBuilder() *DefaultsBuilder {
	r := DefaultsBuilder{
		&Defaults{},
	}

	return &r
}

// Build finalize the chain and returns the Defaults struct
func (rb *DefaultsBuilder) Build() Defaults {
	return *rb.v
}

func (rb *DefaultsBuilder) AnomalyDetectors(anomalydetectors *AnomalyDetectorsBuilder) *DefaultsBuilder {
	v := anomalydetectors.Build()
	rb.v.AnomalyDetectors = v
	return rb
}

func (rb *DefaultsBuilder) Datafeeds(datafeeds *DatafeedsBuilder) *DefaultsBuilder {
	v := datafeeds.Build()
	rb.v.Datafeeds = v
	return rb
}

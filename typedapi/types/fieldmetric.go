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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/metric"
)

// FieldMetric type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/rollup/_types/Metric.ts#L30-L35
type FieldMetric struct {
	// Field The field to collect metrics for. This must be a numeric of some kind.
	Field Field `json:"field"`
	// Metrics An array of metrics to collect for the field. At least one metric must be
	// configured.
	Metrics []metric.Metric `json:"metrics"`
}

// FieldMetricBuilder holds FieldMetric struct and provides a builder API.
type FieldMetricBuilder struct {
	v *FieldMetric
}

// NewFieldMetric provides a builder for the FieldMetric struct.
func NewFieldMetricBuilder() *FieldMetricBuilder {
	r := FieldMetricBuilder{
		&FieldMetric{},
	}

	return &r
}

// Build finalize the chain and returns the FieldMetric struct
func (rb *FieldMetricBuilder) Build() FieldMetric {
	return *rb.v
}

// Field The field to collect metrics for. This must be a numeric of some kind.

func (rb *FieldMetricBuilder) Field(field Field) *FieldMetricBuilder {
	rb.v.Field = field
	return rb
}

// Metrics An array of metrics to collect for the field. At least one metric must be
// configured.

func (rb *FieldMetricBuilder) Metrics(metrics ...metric.Metric) *FieldMetricBuilder {
	rb.v.Metrics = metrics
	return rb
}

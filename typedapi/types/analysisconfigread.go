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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// AnalysisConfigRead type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Analysis.ts#L79-L148
type AnalysisConfigRead struct {
	// BucketSpan The size of the interval that the analysis is aggregated into, typically
	// between `5m` and `1h`.
	BucketSpan Duration `json:"bucket_span"`
	// CategorizationAnalyzer If `categorization_field_name` is specified, you can also define the analyzer
	// that is used to interpret the categorization field.
	// This property cannot be used at the same time as `categorization_filters`.
	// The categorization analyzer specifies how the `categorization_field` is
	// interpreted by the categorization process.
	CategorizationAnalyzer CategorizationAnalyzer `json:"categorization_analyzer,omitempty"`
	// CategorizationFieldName If this property is specified, the values of the specified field will be
	// categorized.
	// The resulting categories must be used in a detector by setting
	// `by_field_name`, `over_field_name`, or `partition_field_name` to the keyword
	// `mlcategory`.
	CategorizationFieldName *string `json:"categorization_field_name,omitempty"`
	// CategorizationFilters If `categorization_field_name` is specified, you can also define optional
	// filters.
	// This property expects an array of regular expressions.
	// The expressions are used to filter out matching sequences from the
	// categorization field values.
	CategorizationFilters []string `json:"categorization_filters,omitempty"`
	// Detectors An array of detector configuration objects.
	// Detector configuration objects specify which data fields a job analyzes.
	// They also specify which analytical functions are used.
	// You can specify multiple detectors for a job.
	Detectors []DetectorRead `json:"detectors"`
	// Influencers A comma separated list of influencer field names.
	// Typically these can be the by, over, or partition fields that are used in the
	// detector configuration.
	// You might also want to use a field name that is not specifically named in a
	// detector, but is available as part of the input data.
	// When you use multiple detectors, the use of influencers is recommended as it
	// aggregates results for each influencer entity.
	Influencers []string `json:"influencers"`
	// Latency The size of the window in which to expect data that is out of time order.
	// Defaults to no latency.
	// If you specify a non-zero value, it must be greater than or equal to one
	// second.
	Latency Duration `json:"latency,omitempty"`
	// ModelPruneWindow Advanced configuration option.
	// Affects the pruning of models that have not been updated for the given time
	// duration.
	// The value must be set to a multiple of the `bucket_span`.
	// If set too low, important information may be removed from the model.
	// Typically, set to `30d` or longer.
	// If not set, model pruning only occurs if the model memory status reaches the
	// soft limit or the hard limit.
	// For jobs created in 8.1 and later, the default value is the greater of `30d`
	// or 20 times `bucket_span`.
	ModelPruneWindow Duration `json:"model_prune_window,omitempty"`
	// MultivariateByFields This functionality is reserved for internal use.
	// It is not supported for use in customer environments and is not subject to
	// the support SLA of official GA features.
	// If set to `true`, the analysis will automatically find correlations between
	// metrics for a given by field value and report anomalies when those
	// correlations cease to hold.
	MultivariateByFields *bool `json:"multivariate_by_fields,omitempty"`
	// PerPartitionCategorization Settings related to how categorization interacts with partition fields.
	PerPartitionCategorization *PerPartitionCategorization `json:"per_partition_categorization,omitempty"`
	// SummaryCountFieldName If this property is specified, the data that is fed to the job is expected to
	// be pre-summarized.
	// This property value is the name of the field that contains the count of raw
	// data points that have been summarized.
	// The same `summary_count_field_name` applies to all detectors in the job.
	SummaryCountFieldName *string `json:"summary_count_field_name,omitempty"`
}

func (s *AnalysisConfigRead) UnmarshalJSON(data []byte) error {

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

		case "bucket_span":
			if err := dec.Decode(&s.BucketSpan); err != nil {
				return err
			}

		case "categorization_analyzer":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := NewCategorizationAnalyzerDefinition()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.CategorizationAnalyzer = *o

			default:
				if err := localDec.Decode(&s.CategorizationAnalyzer); err != nil {
					return err
				}
			}

		case "categorization_field_name":
			if err := dec.Decode(&s.CategorizationFieldName); err != nil {
				return err
			}

		case "categorization_filters":
			if err := dec.Decode(&s.CategorizationFilters); err != nil {
				return err
			}

		case "detectors":
			if err := dec.Decode(&s.Detectors); err != nil {
				return err
			}

		case "influencers":
			if err := dec.Decode(&s.Influencers); err != nil {
				return err
			}

		case "latency":
			if err := dec.Decode(&s.Latency); err != nil {
				return err
			}

		case "model_prune_window":
			if err := dec.Decode(&s.ModelPruneWindow); err != nil {
				return err
			}

		case "multivariate_by_fields":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.MultivariateByFields = &value
			case bool:
				s.MultivariateByFields = &v
			}

		case "per_partition_categorization":
			if err := dec.Decode(&s.PerPartitionCategorization); err != nil {
				return err
			}

		case "summary_count_field_name":
			if err := dec.Decode(&s.SummaryCountFieldName); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAnalysisConfigRead returns a AnalysisConfigRead.
func NewAnalysisConfigRead() *AnalysisConfigRead {
	r := &AnalysisConfigRead{}

	return r
}

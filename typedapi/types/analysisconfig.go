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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// AnalysisConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/Analysis.ts#L29-L77
type AnalysisConfig struct {
	// BucketSpan The size of the interval that the analysis is aggregated into, typically
	// between `5m` and `1h`. This value should be either a whole number of days or
	// equate to a
	// whole number of buckets in one day. If the anomaly detection job uses a
	// datafeed with aggregations, this value must also be divisible by the interval
	// of the date histogram aggregation.
	// * @server_default 5m
	BucketSpan Duration `json:"bucket_span"`
	// CategorizationAnalyzer If `categorization_field_name` is specified, you can also define the analyzer
	// that is used to interpret the categorization field. This property cannot be
	// used at the same time as `categorization_filters`. The categorization
	// analyzer specifies how the `categorization_field` is interpreted by the
	// categorization process. The `categorization_analyzer` field can be specified
	// either as a string or as an object. If it is a string, it must refer to a
	// built-in analyzer or one added by another plugin.
	CategorizationAnalyzer *CategorizationAnalyzer `json:"categorization_analyzer,omitempty"`
	// CategorizationFieldName If this property is specified, the values of the specified field will be
	// categorized. The resulting categories must be used in a detector by setting
	// `by_field_name`, `over_field_name`, or `partition_field_name` to the keyword
	// `mlcategory`.
	CategorizationFieldName *string `json:"categorization_field_name,omitempty"`
	// CategorizationFilters If `categorization_field_name` is specified, you can also define optional
	// filters. This property expects an array of regular expressions. The
	// expressions are used to filter out matching sequences from the categorization
	// field values. You can use this functionality to fine tune the categorization
	// by excluding sequences from consideration when categories are defined. For
	// example, you can exclude SQL statements that appear in your log files. This
	// property cannot be used at the same time as `categorization_analyzer`. If you
	// only want to define simple regular expression filters that are applied prior
	// to tokenization, setting this property is the easiest method. If you also
	// want to customize the tokenizer or post-tokenization filtering, use the
	// `categorization_analyzer` property instead and include the filters as
	// pattern_replace character filters. The effect is exactly the same.
	CategorizationFilters []string `json:"categorization_filters,omitempty"`
	// Detectors Detector configuration objects specify which data fields a job analyzes. They
	// also specify which analytical functions are used. You can specify multiple
	// detectors for a job. If the detectors array does not contain at least one
	// detector, no analysis can occur and an error is returned.
	Detectors []Detector `json:"detectors"`
	// Influencers A comma separated list of influencer field names. Typically these can be the
	// by, over, or partition fields that are used in the detector configuration.
	// You might also want to use a field name that is not specifically named in a
	// detector, but is available as part of the input data. When you use multiple
	// detectors, the use of influencers is recommended as it aggregates results for
	// each influencer entity.
	Influencers []string `json:"influencers,omitempty"`
	// Latency The size of the window in which to expect data that is out of time order. If
	// you specify a non-zero value, it must be greater than or equal to one second.
	// NOTE: Latency is applicable only when you send data by using the post data
	// API.
	Latency *Duration `json:"latency,omitempty"`
	// ModelPruneWindow Advanced configuration option. Affects the pruning of models that have not
	// been updated for the given time duration. The value must be set to a multiple
	// of the `bucket_span`. If set too low, important information may be removed
	// from the model. For jobs created in 8.1 and later, the default value is the
	// greater of `30d` or 20 times `bucket_span`.
	ModelPruneWindow *Duration `json:"model_prune_window,omitempty"`
	// MultivariateByFields This functionality is reserved for internal use. It is not supported for use
	// in customer environments and is not subject to the support SLA of official GA
	// features. If set to `true`, the analysis will automatically find correlations
	// between metrics for a given by field value and report anomalies when those
	// correlations cease to hold. For example, suppose CPU and memory usage on host
	// A is usually highly correlated with the same metrics on host B. Perhaps this
	// correlation occurs because they are running a load-balanced application. If
	// you enable this property, anomalies will be reported when, for example, CPU
	// usage on host A is high and the value of CPU usage on host B is low. That is
	// to say, you’ll see an anomaly when the CPU of host A is unusual given the CPU
	// of host B. To use the `multivariate_by_fields` property, you must also
	// specify `by_field_name` in your detector.
	MultivariateByFields *bool `json:"multivariate_by_fields,omitempty"`
	// PerPartitionCategorization Settings related to how categorization interacts with partition fields.
	PerPartitionCategorization *PerPartitionCategorization `json:"per_partition_categorization,omitempty"`
	// SummaryCountFieldName If this property is specified, the data that is fed to the job is expected to
	// be pre-summarized. This property value is the name of the field that contains
	// the count of raw data points that have been summarized. The same
	// `summary_count_field_name` applies to all detectors in the job. NOTE: The
	// `summary_count_field_name` property cannot be used with the `metric`
	// function.
	SummaryCountFieldName *string `json:"summary_count_field_name,omitempty"`
}

// NewAnalysisConfig returns a AnalysisConfig.
func NewAnalysisConfig() *AnalysisConfig {
	r := &AnalysisConfig{}

	return r
}

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

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// JobConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Job.ts#L182-L283
type JobConfig struct {
	// AllowLazyOpen Advanced configuration option. Specifies whether this job can open when there
	// is insufficient machine learning node capacity for it to be immediately
	// assigned to a node.
	AllowLazyOpen *bool `json:"allow_lazy_open,omitempty"`
	// AnalysisConfig The analysis configuration, which specifies how to analyze the data.
	// After you create a job, you cannot change the analysis configuration; all the
	// properties are informational.
	AnalysisConfig AnalysisConfig `json:"analysis_config"`
	// AnalysisLimits Limits can be applied for the resources required to hold the mathematical
	// models in memory.
	// These limits are approximate and can be set per job.
	// They do not control the memory used by other processes, for example the
	// Elasticsearch Java processes.
	AnalysisLimits *AnalysisLimits `json:"analysis_limits,omitempty"`
	// BackgroundPersistInterval Advanced configuration option.
	// The time between each periodic persistence of the model.
	// The default value is a randomized value between 3 to 4 hours, which avoids
	// all jobs persisting at exactly the same time.
	// The smallest allowed value is 1 hour.
	BackgroundPersistInterval Duration `json:"background_persist_interval,omitempty"`
	// CustomSettings Advanced configuration option.
	// Contains custom metadata about the job.
	CustomSettings json.RawMessage `json:"custom_settings,omitempty"`
	// DailyModelSnapshotRetentionAfterDays Advanced configuration option, which affects the automatic removal of old
	// model snapshots for this job.
	// It specifies a period of time (in days) after which only the first snapshot
	// per day is retained.
	// This period is relative to the timestamp of the most recent snapshot for this
	// job.
	DailyModelSnapshotRetentionAfterDays *int64 `json:"daily_model_snapshot_retention_after_days,omitempty"`
	// DataDescription The data description defines the format of the input data when you send data
	// to the job by using the post data API.
	// Note that when configure a datafeed, these properties are automatically set.
	DataDescription DataDescription `json:"data_description"`
	// DatafeedConfig The datafeed, which retrieves data from Elasticsearch for analysis by the
	// job.
	// You can associate only one datafeed with each anomaly detection job.
	DatafeedConfig *DatafeedConfig `json:"datafeed_config,omitempty"`
	// Description A description of the job.
	Description *string `json:"description,omitempty"`
	// Groups A list of job groups. A job can belong to no groups or many.
	Groups []string `json:"groups,omitempty"`
	// JobId Identifier for the anomaly detection job.
	// This identifier can contain lowercase alphanumeric characters (a-z and 0-9),
	// hyphens, and underscores.
	// It must start and end with alphanumeric characters.
	JobId *string `json:"job_id,omitempty"`
	// JobType Reserved for future use, currently set to `anomaly_detector`.
	JobType *string `json:"job_type,omitempty"`
	// ModelPlotConfig This advanced configuration option stores model information along with the
	// results.
	// It provides a more detailed view into anomaly detection.
	// Model plot provides a simplified and indicative view of the model and its
	// bounds.
	ModelPlotConfig *ModelPlotConfig `json:"model_plot_config,omitempty"`
	// ModelSnapshotRetentionDays Advanced configuration option, which affects the automatic removal of old
	// model snapshots for this job.
	// It specifies the maximum period of time (in days) that snapshots are
	// retained.
	// This period is relative to the timestamp of the most recent snapshot for this
	// job.
	// The default value is `10`, which means snapshots ten days older than the
	// newest snapshot are deleted.
	ModelSnapshotRetentionDays *int64 `json:"model_snapshot_retention_days,omitempty"`
	// RenormalizationWindowDays Advanced configuration option.
	// The period over which adjustments to the score are applied, as new data is
	// seen.
	// The default value is the longer of 30 days or 100 `bucket_spans`.
	RenormalizationWindowDays *int64 `json:"renormalization_window_days,omitempty"`
	// ResultsIndexName A text string that affects the name of the machine learning results index.
	// The default value is `shared`, which generates an index named
	// `.ml-anomalies-shared`.
	ResultsIndexName *string `json:"results_index_name,omitempty"`
	// ResultsRetentionDays Advanced configuration option.
	// The period of time (in days) that results are retained.
	// Age is calculated relative to the timestamp of the latest bucket result.
	// If this property has a non-null value, once per day at 00:30 (server time),
	// results that are the specified number of days older than the latest bucket
	// result are deleted from Elasticsearch.
	// The default value is null, which means all results are retained.
	// Annotations generated by the system also count as results for retention
	// purposes; they are deleted after the same number of days as results.
	// Annotations added by users are retained forever.
	ResultsRetentionDays *int64 `json:"results_retention_days,omitempty"`
}

func (s *JobConfig) UnmarshalJSON(data []byte) error {

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

		case "allow_lazy_open":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowLazyOpen = &value
			case bool:
				s.AllowLazyOpen = &v
			}

		case "analysis_config":
			if err := dec.Decode(&s.AnalysisConfig); err != nil {
				return err
			}

		case "analysis_limits":
			if err := dec.Decode(&s.AnalysisLimits); err != nil {
				return err
			}

		case "background_persist_interval":
			if err := dec.Decode(&s.BackgroundPersistInterval); err != nil {
				return err
			}

		case "custom_settings":
			if err := dec.Decode(&s.CustomSettings); err != nil {
				return err
			}

		case "daily_model_snapshot_retention_after_days":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DailyModelSnapshotRetentionAfterDays = &value
			case float64:
				f := int64(v)
				s.DailyModelSnapshotRetentionAfterDays = &f
			}

		case "data_description":
			if err := dec.Decode(&s.DataDescription); err != nil {
				return err
			}

		case "datafeed_config":
			if err := dec.Decode(&s.DatafeedConfig); err != nil {
				return err
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "groups":
			if err := dec.Decode(&s.Groups); err != nil {
				return err
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "job_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.JobType = &o

		case "model_plot_config":
			if err := dec.Decode(&s.ModelPlotConfig); err != nil {
				return err
			}

		case "model_snapshot_retention_days":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ModelSnapshotRetentionDays = &value
			case float64:
				f := int64(v)
				s.ModelSnapshotRetentionDays = &f
			}

		case "renormalization_window_days":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RenormalizationWindowDays = &value
			case float64:
				f := int64(v)
				s.RenormalizationWindowDays = &f
			}

		case "results_index_name":
			if err := dec.Decode(&s.ResultsIndexName); err != nil {
				return err
			}

		case "results_retention_days":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ResultsRetentionDays = &value
			case float64:
				f := int64(v)
				s.ResultsRetentionDays = &f
			}

		}
	}
	return nil
}

// NewJobConfig returns a JobConfig.
func NewJobConfig() *JobConfig {
	r := &JobConfig{}

	return r
}

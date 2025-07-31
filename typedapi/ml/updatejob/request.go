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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package updatejob

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package updatejob
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/update_job/MlUpdateJobRequest.ts#L33-L147
type Request struct {

	// AllowLazyOpen Advanced configuration option. Specifies whether this job can open when
	// there is insufficient machine learning node capacity for it to be
	// immediately assigned to a node. If `false` and a machine learning node
	// with capacity to run the job cannot immediately be found, the open
	// anomaly detection jobs API returns an error. However, this is also
	// subject to the cluster-wide `xpack.ml.max_lazy_ml_nodes` setting. If this
	// option is set to `true`, the open anomaly detection jobs API does not
	// return an error and the job waits in the opening state until sufficient
	// machine learning node capacity is available.
	AllowLazyOpen  *bool                      `json:"allow_lazy_open,omitempty"`
	AnalysisLimits *types.AnalysisMemoryLimit `json:"analysis_limits,omitempty"`
	// BackgroundPersistInterval Advanced configuration option. The time between each periodic persistence
	// of the model.
	// The default value is a randomized value between 3 to 4 hours, which
	// avoids all jobs persisting at exactly the same time. The smallest allowed
	// value is 1 hour.
	// For very large models (several GB), persistence could take 10-20 minutes,
	// so do not set the value too low.
	// If the job is open when you make the update, you must stop the datafeed,
	// close the job, then reopen the job and restart the datafeed for the
	// changes to take effect.
	BackgroundPersistInterval types.Duration `json:"background_persist_interval,omitempty"`
	CategorizationFilters     []string       `json:"categorization_filters,omitempty"`
	// CustomSettings Advanced configuration option. Contains custom meta data about the job.
	// For example, it can contain custom URL information as shown in Adding
	// custom URLs to machine learning results.
	CustomSettings map[string]json.RawMessage `json:"custom_settings,omitempty"`
	// DailyModelSnapshotRetentionAfterDays Advanced configuration option, which affects the automatic removal of old
	// model snapshots for this job. It specifies a period of time (in days)
	// after which only the first snapshot per day is retained. This period is
	// relative to the timestamp of the most recent snapshot for this job. Valid
	// values range from 0 to `model_snapshot_retention_days`. For jobs created
	// before version 7.8.0, the default value matches
	// `model_snapshot_retention_days`.
	DailyModelSnapshotRetentionAfterDays *int64 `json:"daily_model_snapshot_retention_after_days,omitempty"`
	// Description A description of the job.
	Description *string `json:"description,omitempty"`
	// Detectors An array of detector update objects.
	Detectors []types.DetectorUpdate `json:"detectors,omitempty"`
	// Groups A list of job groups. A job can belong to no groups or many.
	Groups           []string               `json:"groups,omitempty"`
	ModelPlotConfig  *types.ModelPlotConfig `json:"model_plot_config,omitempty"`
	ModelPruneWindow types.Duration         `json:"model_prune_window,omitempty"`
	// ModelSnapshotRetentionDays Advanced configuration option, which affects the automatic removal of old
	// model snapshots for this job. It specifies the maximum period of time (in
	// days) that snapshots are retained. This period is relative to the
	// timestamp of the most recent snapshot for this job.
	ModelSnapshotRetentionDays *int64 `json:"model_snapshot_retention_days,omitempty"`
	// PerPartitionCategorization Settings related to how categorization interacts with partition fields.
	PerPartitionCategorization *types.PerPartitionCategorization `json:"per_partition_categorization,omitempty"`
	// RenormalizationWindowDays Advanced configuration option. The period over which adjustments to the
	// score are applied, as new data is seen.
	RenormalizationWindowDays *int64 `json:"renormalization_window_days,omitempty"`
	// ResultsRetentionDays Advanced configuration option. The period of time (in days) that results
	// are retained. Age is calculated relative to the timestamp of the latest
	// bucket result. If this property has a non-null value, once per day at
	// 00:30 (server time), results that are the specified number of days older
	// than the latest bucket result are deleted from Elasticsearch. The default
	// value is null, which means all results are retained.
	ResultsRetentionDays *int64 `json:"results_retention_days,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		CustomSettings: make(map[string]json.RawMessage, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Updatejob request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowLazyOpen", err)
				}
				s.AllowLazyOpen = &value
			case bool:
				s.AllowLazyOpen = &v
			}

		case "analysis_limits":
			if err := dec.Decode(&s.AnalysisLimits); err != nil {
				return fmt.Errorf("%s | %w", "AnalysisLimits", err)
			}

		case "background_persist_interval":
			if err := dec.Decode(&s.BackgroundPersistInterval); err != nil {
				return fmt.Errorf("%s | %w", "BackgroundPersistInterval", err)
			}

		case "categorization_filters":
			if err := dec.Decode(&s.CategorizationFilters); err != nil {
				return fmt.Errorf("%s | %w", "CategorizationFilters", err)
			}

		case "custom_settings":
			if s.CustomSettings == nil {
				s.CustomSettings = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.CustomSettings); err != nil {
				return fmt.Errorf("%s | %w", "CustomSettings", err)
			}

		case "daily_model_snapshot_retention_after_days":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DailyModelSnapshotRetentionAfterDays", err)
				}
				s.DailyModelSnapshotRetentionAfterDays = &value
			case float64:
				f := int64(v)
				s.DailyModelSnapshotRetentionAfterDays = &f
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "detectors":
			if err := dec.Decode(&s.Detectors); err != nil {
				return fmt.Errorf("%s | %w", "Detectors", err)
			}

		case "groups":
			if err := dec.Decode(&s.Groups); err != nil {
				return fmt.Errorf("%s | %w", "Groups", err)
			}

		case "model_plot_config":
			if err := dec.Decode(&s.ModelPlotConfig); err != nil {
				return fmt.Errorf("%s | %w", "ModelPlotConfig", err)
			}

		case "model_prune_window":
			if err := dec.Decode(&s.ModelPruneWindow); err != nil {
				return fmt.Errorf("%s | %w", "ModelPruneWindow", err)
			}

		case "model_snapshot_retention_days":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ModelSnapshotRetentionDays", err)
				}
				s.ModelSnapshotRetentionDays = &value
			case float64:
				f := int64(v)
				s.ModelSnapshotRetentionDays = &f
			}

		case "per_partition_categorization":
			if err := dec.Decode(&s.PerPartitionCategorization); err != nil {
				return fmt.Errorf("%s | %w", "PerPartitionCategorization", err)
			}

		case "renormalization_window_days":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RenormalizationWindowDays", err)
				}
				s.RenormalizationWindowDays = &value
			case float64:
				f := int64(v)
				s.RenormalizationWindowDays = &f
			}

		case "results_retention_days":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ResultsRetentionDays", err)
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

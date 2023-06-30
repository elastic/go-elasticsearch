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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package putjob

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putjob
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/put_job/MlPutJobRequest.ts#L30-L111
type Request struct {

	// AllowLazyOpen Advanced configuration option. Specifies whether this job can open when there
	// is insufficient machine learning node capacity for it to be immediately
	// assigned to a node. By default, if a machine learning node with capacity to
	// run the job cannot immediately be found, the open anomaly detection jobs API
	// returns an error. However, this is also subject to the cluster-wide
	// `xpack.ml.max_lazy_ml_nodes` setting. If this option is set to true, the open
	// anomaly detection jobs API does not return an error and the job waits in the
	// opening state until sufficient machine learning node capacity is available.
	AllowLazyOpen *bool `json:"allow_lazy_open,omitempty"`
	// AnalysisConfig Specifies how to analyze the data. After you create a job, you cannot change
	// the analysis configuration; all the properties are informational.
	AnalysisConfig types.AnalysisConfig `json:"analysis_config"`
	// AnalysisLimits Limits can be applied for the resources required to hold the mathematical
	// models in memory. These limits are approximate and can be set per job. They
	// do not control the memory used by other processes, for example the
	// Elasticsearch Java processes.
	AnalysisLimits *types.AnalysisLimits `json:"analysis_limits,omitempty"`
	// BackgroundPersistInterval Advanced configuration option. The time between each periodic persistence of
	// the model. The default value is a randomized value between 3 to 4 hours,
	// which avoids all jobs persisting at exactly the same time. The smallest
	// allowed value is 1 hour. For very large models (several GB), persistence
	// could take 10-20 minutes, so do not set the `background_persist_interval`
	// value too low.
	BackgroundPersistInterval types.Duration `json:"background_persist_interval,omitempty"`
	// CustomSettings Advanced configuration option. Contains custom meta data about the job.
	CustomSettings json.RawMessage `json:"custom_settings,omitempty"`
	// DailyModelSnapshotRetentionAfterDays Advanced configuration option, which affects the automatic removal of old
	// model snapshots for this job. It specifies a period of time (in days) after
	// which only the first snapshot per day is retained. This period is relative to
	// the timestamp of the most recent snapshot for this job. Valid values range
	// from 0 to `model_snapshot_retention_days`.
	DailyModelSnapshotRetentionAfterDays *int64 `json:"daily_model_snapshot_retention_after_days,omitempty"`
	// DataDescription Defines the format of the input data when you send data to the job by using
	// the post data API. Note that when configure a datafeed, these properties are
	// automatically set. When data is received via the post data API, it is not
	// stored in Elasticsearch. Only the results for anomaly detection are retained.
	DataDescription types.DataDescription `json:"data_description"`
	// DatafeedConfig Defines a datafeed for the anomaly detection job. If Elasticsearch security
	// features are enabled, your datafeed remembers which roles the user who
	// created it had at the time of creation and runs the query using those same
	// roles. If you provide secondary authorization headers, those credentials are
	// used instead.
	DatafeedConfig *types.DatafeedConfig `json:"datafeed_config,omitempty"`
	// Description A description of the job.
	Description *string `json:"description,omitempty"`
	// Groups A list of job groups. A job can belong to no groups or many.
	Groups []string `json:"groups,omitempty"`
	// ModelPlotConfig This advanced configuration option stores model information along with the
	// results. It provides a more detailed view into anomaly detection. If you
	// enable model plot it can add considerable overhead to the performance of the
	// system; it is not feasible for jobs with many entities. Model plot provides a
	// simplified and indicative view of the model and its bounds. It does not
	// display complex features such as multivariate correlations or multimodal
	// data. As such, anomalies may occasionally be reported which cannot be seen in
	// the model plot. Model plot config can be configured when the job is created
	// or updated later. It must be disabled if performance issues are experienced.
	ModelPlotConfig *types.ModelPlotConfig `json:"model_plot_config,omitempty"`
	// ModelSnapshotRetentionDays Advanced configuration option, which affects the automatic removal of old
	// model snapshots for this job. It specifies the maximum period of time (in
	// days) that snapshots are retained. This period is relative to the timestamp
	// of the most recent snapshot for this job. By default, snapshots ten days
	// older than the newest snapshot are deleted.
	ModelSnapshotRetentionDays *int64 `json:"model_snapshot_retention_days,omitempty"`
	// RenormalizationWindowDays Advanced configuration option. The period over which adjustments to the score
	// are applied, as new data is seen. The default value is the longer of 30 days
	// or 100 bucket spans.
	RenormalizationWindowDays *int64 `json:"renormalization_window_days,omitempty"`
	// ResultsIndexName A text string that affects the name of the machine learning results index. By
	// default, the job generates an index named `.ml-anomalies-shared`.
	ResultsIndexName *string `json:"results_index_name,omitempty"`
	// ResultsRetentionDays Advanced configuration option. The period of time (in days) that results are
	// retained. Age is calculated relative to the timestamp of the latest bucket
	// result. If this property has a non-null value, once per day at 00:30 (server
	// time), results that are the specified number of days older than the latest
	// bucket result are deleted from Elasticsearch. The default value is null,
	// which means all results are retained. Annotations generated by the system
	// also count as results for retention purposes; they are deleted after the same
	// number of days as results. Annotations added by users are retained forever.
	ResultsRetentionDays *int64 `json:"results_retention_days,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putjob request: %w", err)
	}

	return &req, nil
}

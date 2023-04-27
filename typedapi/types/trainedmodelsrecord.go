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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// TrainedModelsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/cat/ml_trained_models/types.ts#L23-L111
type TrainedModelsRecord struct {
	// CreateTime The time the model was created
	CreateTime DateTime `json:"create_time,omitempty"`
	// CreatedBy who created the model
	CreatedBy *string `json:"created_by,omitempty"`
	// DataFrameAnalysis The analysis used by the data frame to build the model
	DataFrameAnalysis *string `json:"data_frame.analysis,omitempty"`
	// DataFrameCreateTime The time the data frame analytics config was created
	DataFrameCreateTime *string `json:"data_frame.create_time,omitempty"`
	// DataFrameId The data frame analytics config id that created the model (if still
	// available)
	DataFrameId *string `json:"data_frame.id,omitempty"`
	// DataFrameSourceIndex The source index used to train in the data frame analysis
	DataFrameSourceIndex *string `json:"data_frame.source_index,omitempty"`
	// Description The model description
	Description *string `json:"description,omitempty"`
	// HeapSize the estimated heap size to keep the model in memory
	HeapSize ByteSize `json:"heap_size,omitempty"`
	// Id the trained model id
	Id *string `json:"id,omitempty"`
	// IngestCount The total number of docs processed by the model
	IngestCount *string `json:"ingest.count,omitempty"`
	// IngestCurrent The total documents currently being handled by the model
	IngestCurrent *string `json:"ingest.current,omitempty"`
	// IngestFailed The total count of failed ingest attempts with this model
	IngestFailed *string `json:"ingest.failed,omitempty"`
	// IngestPipelines The number of pipelines referencing the model
	IngestPipelines *string `json:"ingest.pipelines,omitempty"`
	// IngestTime The total time spent processing docs with this model
	IngestTime *string `json:"ingest.time,omitempty"`
	// License The license level of the model
	License *string `json:"license,omitempty"`
	// Operations the estimated number of operations to use the model
	Operations *string `json:"operations,omitempty"`
	Type       *string `json:"type,omitempty"`
	// Version The version of Elasticsearch when the model was created
	Version *string `json:"version,omitempty"`
}

func (s *TrainedModelsRecord) UnmarshalJSON(data []byte) error {

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

		case "create_time", "ct":
			if err := dec.Decode(&s.CreateTime); err != nil {
				return err
			}

		case "created_by", "c", "createdBy":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.CreatedBy = &o

		case "data_frame.analysis", "dfa", "dataFrameAnalyticsAnalysis":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.DataFrameAnalysis = &o

		case "data_frame.create_time", "dft", "dataFrameAnalyticsTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.DataFrameCreateTime = &o

		case "data_frame.id", "dfid", "dataFrameAnalytics":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.DataFrameId = &o

		case "data_frame.source_index", "dfsi", "dataFrameAnalyticsSrcIndex":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.DataFrameSourceIndex = &o

		case "description", "d":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Description = &o

		case "heap_size", "hs", "modelHeapSize":
			if err := dec.Decode(&s.HeapSize); err != nil {
				return err
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "ingest.count", "ic", "ingestCount":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IngestCount = &o

		case "ingest.current", "icurr", "ingestCurrent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IngestCurrent = &o

		case "ingest.failed", "if", "ingestFailed":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IngestFailed = &o

		case "ingest.pipelines", "ip", "ingestPipelines":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IngestPipelines = &o

		case "ingest.time", "it", "ingestTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IngestTime = &o

		case "license", "l":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.License = &o

		case "operations", "o", "modelOperations":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Operations = &o

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Type = &o

		case "version", "v":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTrainedModelsRecord returns a TrainedModelsRecord.
func NewTrainedModelsRecord() *TrainedModelsRecord {
	r := &TrainedModelsRecord{}

	return r
}

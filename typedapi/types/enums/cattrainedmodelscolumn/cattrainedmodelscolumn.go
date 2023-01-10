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


// Package cattrainedmodelscolumn
package cattrainedmodelscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/_types/CatBase.ts#L561-L635
type CatTrainedModelsColumn struct {
	Name string
}

var (
	Createtime = CatTrainedModelsColumn{"create_time"}

	Createdby = CatTrainedModelsColumn{"created_by"}

	Dataframeanalyticsid = CatTrainedModelsColumn{"data_frame_analytics_id"}

	Description = CatTrainedModelsColumn{"description"}

	Heapsize = CatTrainedModelsColumn{"heap_size"}

	Id = CatTrainedModelsColumn{"id"}

	IngestCount = CatTrainedModelsColumn{"ingest.count"}

	IngestCurrent = CatTrainedModelsColumn{"ingest.current"}

	IngestFailed = CatTrainedModelsColumn{"ingest.failed"}

	IngestPipelines = CatTrainedModelsColumn{"ingest.pipelines"}

	IngestTime = CatTrainedModelsColumn{"ingest.time"}

	License = CatTrainedModelsColumn{"license"}

	Operations = CatTrainedModelsColumn{"operations"}

	Version = CatTrainedModelsColumn{"version"}
)

func (c CatTrainedModelsColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatTrainedModelsColumn) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "create_time":
		*c = Createtime
	case "created_by":
		*c = Createdby
	case "data_frame_analytics_id":
		*c = Dataframeanalyticsid
	case "description":
		*c = Description
	case "heap_size":
		*c = Heapsize
	case "id":
		*c = Id
	case "ingest.count":
		*c = IngestCount
	case "ingest.current":
		*c = IngestCurrent
	case "ingest.failed":
		*c = IngestFailed
	case "ingest.pipelines":
		*c = IngestPipelines
	case "ingest.time":
		*c = IngestTime
	case "license":
		*c = License
	case "operations":
		*c = Operations
	case "version":
		*c = Version
	default:
		*c = CatTrainedModelsColumn{string(text)}
	}

	return nil
}

func (c CatTrainedModelsColumn) String() string {
	return c.Name
}

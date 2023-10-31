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

// Package cattrainedmodelscolumn
package cattrainedmodelscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cat/_types/CatBase.ts#L561-L635
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

	Ingestcount = CatTrainedModelsColumn{"ingest.count"}

	Ingestcurrent = CatTrainedModelsColumn{"ingest.current"}

	Ingestfailed = CatTrainedModelsColumn{"ingest.failed"}

	Ingestpipelines = CatTrainedModelsColumn{"ingest.pipelines"}

	Ingesttime = CatTrainedModelsColumn{"ingest.time"}

	License = CatTrainedModelsColumn{"license"}

	Operations = CatTrainedModelsColumn{"operations"}

	Version = CatTrainedModelsColumn{"version"}
)

func (c CatTrainedModelsColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatTrainedModelsColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

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
		*c = Ingestcount
	case "ingest.current":
		*c = Ingestcurrent
	case "ingest.failed":
		*c = Ingestfailed
	case "ingest.pipelines":
		*c = Ingestpipelines
	case "ingest.time":
		*c = Ingesttime
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

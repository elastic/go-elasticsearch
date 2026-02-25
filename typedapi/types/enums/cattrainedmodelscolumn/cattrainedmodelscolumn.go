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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package cattrainedmodelscolumn
package cattrainedmodelscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L2548-L2622
type CatTrainedModelsColumn struct {
	Name string
}

var (

	// Createtime The time when the trained model was created.
	Createtime = CatTrainedModelsColumn{"create_time"}

	// Createdby Information on the creator of the trained model.
	Createdby = CatTrainedModelsColumn{"created_by"}

	// Dataframeanalyticsid Identifier for the data frame analytics job that created the model. Only
	// displayed if it is still available.
	Dataframeanalyticsid = CatTrainedModelsColumn{"data_frame_analytics_id"}

	// Description The description of the trained model.
	Description = CatTrainedModelsColumn{"description"}

	// Heapsize The estimated heap size to keep the trained model in memory.
	Heapsize = CatTrainedModelsColumn{"heap_size"}

	// Id Identifier for the trained model.
	Id = CatTrainedModelsColumn{"id"}

	// Ingestcount The total number of documents that are processed by the model.
	Ingestcount = CatTrainedModelsColumn{"ingest.count"}

	// Ingestcurrent The total number of document that are currently being handled by the trained
	// model.
	Ingestcurrent = CatTrainedModelsColumn{"ingest.current"}

	// Ingestfailed The total number of failed ingest attempts with the trained model.
	Ingestfailed = CatTrainedModelsColumn{"ingest.failed"}

	// Ingestpipelines The total number of ingest pipelines that are referencing the trained model.
	Ingestpipelines = CatTrainedModelsColumn{"ingest.pipelines"}

	// Ingesttime The total time that is spent processing documents with the trained model.
	Ingesttime = CatTrainedModelsColumn{"ingest.time"}

	// License The license level of the trained model.
	License = CatTrainedModelsColumn{"license"}

	// Operations The estimated number of operations to use the trained model. This number
	// helps measuring the computational complexity of the model.
	Operations = CatTrainedModelsColumn{"operations"}

	// Version The Elasticsearch version number in which the trained model was created.
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

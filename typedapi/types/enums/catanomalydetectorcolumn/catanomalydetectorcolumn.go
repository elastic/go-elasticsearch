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

// Package catanomalydetectorcolumn
package catanomalydetectorcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cat/_types/CatBase.ts#L32-L401
type CatAnomalyDetectorColumn struct {
	Name string
}

var (
	Assignmentexplanation = CatAnomalyDetectorColumn{"assignment_explanation"}

	Bucketscount = CatAnomalyDetectorColumn{"buckets.count"}

	Bucketstimeexpavg = CatAnomalyDetectorColumn{"buckets.time.exp_avg"}

	Bucketstimeexpavghour = CatAnomalyDetectorColumn{"buckets.time.exp_avg_hour"}

	Bucketstimemax = CatAnomalyDetectorColumn{"buckets.time.max"}

	Bucketstimemin = CatAnomalyDetectorColumn{"buckets.time.min"}

	Bucketstimetotal = CatAnomalyDetectorColumn{"buckets.time.total"}

	Databuckets = CatAnomalyDetectorColumn{"data.buckets"}

	Dataearliestrecord = CatAnomalyDetectorColumn{"data.earliest_record"}

	Dataemptybuckets = CatAnomalyDetectorColumn{"data.empty_buckets"}

	Datainputbytes = CatAnomalyDetectorColumn{"data.input_bytes"}

	Datainputfields = CatAnomalyDetectorColumn{"data.input_fields"}

	Datainputrecords = CatAnomalyDetectorColumn{"data.input_records"}

	Datainvaliddates = CatAnomalyDetectorColumn{"data.invalid_dates"}

	Datalast = CatAnomalyDetectorColumn{"data.last"}

	Datalastemptybucket = CatAnomalyDetectorColumn{"data.last_empty_bucket"}

	Datalastsparsebucket = CatAnomalyDetectorColumn{"data.last_sparse_bucket"}

	Datalatestrecord = CatAnomalyDetectorColumn{"data.latest_record"}

	Datamissingfields = CatAnomalyDetectorColumn{"data.missing_fields"}

	Dataoutofordertimestamps = CatAnomalyDetectorColumn{"data.out_of_order_timestamps"}

	Dataprocessedfields = CatAnomalyDetectorColumn{"data.processed_fields"}

	Dataprocessedrecords = CatAnomalyDetectorColumn{"data.processed_records"}

	Datasparsebuckets = CatAnomalyDetectorColumn{"data.sparse_buckets"}

	Forecastsmemoryavg = CatAnomalyDetectorColumn{"forecasts.memory.avg"}

	Forecastsmemorymax = CatAnomalyDetectorColumn{"forecasts.memory.max"}

	Forecastsmemorymin = CatAnomalyDetectorColumn{"forecasts.memory.min"}

	Forecastsmemorytotal = CatAnomalyDetectorColumn{"forecasts.memory.total"}

	Forecastsrecordsavg = CatAnomalyDetectorColumn{"forecasts.records.avg"}

	Forecastsrecordsmax = CatAnomalyDetectorColumn{"forecasts.records.max"}

	Forecastsrecordsmin = CatAnomalyDetectorColumn{"forecasts.records.min"}

	Forecastsrecordstotal = CatAnomalyDetectorColumn{"forecasts.records.total"}

	Forecaststimeavg = CatAnomalyDetectorColumn{"forecasts.time.avg"}

	Forecaststimemax = CatAnomalyDetectorColumn{"forecasts.time.max"}

	Forecaststimemin = CatAnomalyDetectorColumn{"forecasts.time.min"}

	Forecaststimetotal = CatAnomalyDetectorColumn{"forecasts.time.total"}

	Forecaststotal = CatAnomalyDetectorColumn{"forecasts.total"}

	Id = CatAnomalyDetectorColumn{"id"}

	Modelbucketallocationfailures = CatAnomalyDetectorColumn{"model.bucket_allocation_failures"}

	Modelbyfields = CatAnomalyDetectorColumn{"model.by_fields"}

	Modelbytes = CatAnomalyDetectorColumn{"model.bytes"}

	Modelbytesexceeded = CatAnomalyDetectorColumn{"model.bytes_exceeded"}

	Modelcategorizationstatus = CatAnomalyDetectorColumn{"model.categorization_status"}

	Modelcategorizeddoccount = CatAnomalyDetectorColumn{"model.categorized_doc_count"}

	Modeldeadcategorycount = CatAnomalyDetectorColumn{"model.dead_category_count"}

	Modelfailedcategorycount = CatAnomalyDetectorColumn{"model.failed_category_count"}

	Modelfrequentcategorycount = CatAnomalyDetectorColumn{"model.frequent_category_count"}

	Modellogtime = CatAnomalyDetectorColumn{"model.log_time"}

	Modelmemorylimit = CatAnomalyDetectorColumn{"model.memory_limit"}

	Modelmemorystatus = CatAnomalyDetectorColumn{"model.memory_status"}

	Modeloverfields = CatAnomalyDetectorColumn{"model.over_fields"}

	Modelpartitionfields = CatAnomalyDetectorColumn{"model.partition_fields"}

	Modelrarecategorycount = CatAnomalyDetectorColumn{"model.rare_category_count"}

	Modeltimestamp = CatAnomalyDetectorColumn{"model.timestamp"}

	Modeltotalcategorycount = CatAnomalyDetectorColumn{"model.total_category_count"}

	Nodeaddress = CatAnomalyDetectorColumn{"node.address"}

	Nodeephemeralid = CatAnomalyDetectorColumn{"node.ephemeral_id"}

	Nodeid = CatAnomalyDetectorColumn{"node.id"}

	Nodename = CatAnomalyDetectorColumn{"node.name"}

	Openedtime = CatAnomalyDetectorColumn{"opened_time"}

	State = CatAnomalyDetectorColumn{"state"}
)

func (c CatAnomalyDetectorColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatAnomalyDetectorColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "assignment_explanation":
		*c = Assignmentexplanation
	case "buckets.count":
		*c = Bucketscount
	case "buckets.time.exp_avg":
		*c = Bucketstimeexpavg
	case "buckets.time.exp_avg_hour":
		*c = Bucketstimeexpavghour
	case "buckets.time.max":
		*c = Bucketstimemax
	case "buckets.time.min":
		*c = Bucketstimemin
	case "buckets.time.total":
		*c = Bucketstimetotal
	case "data.buckets":
		*c = Databuckets
	case "data.earliest_record":
		*c = Dataearliestrecord
	case "data.empty_buckets":
		*c = Dataemptybuckets
	case "data.input_bytes":
		*c = Datainputbytes
	case "data.input_fields":
		*c = Datainputfields
	case "data.input_records":
		*c = Datainputrecords
	case "data.invalid_dates":
		*c = Datainvaliddates
	case "data.last":
		*c = Datalast
	case "data.last_empty_bucket":
		*c = Datalastemptybucket
	case "data.last_sparse_bucket":
		*c = Datalastsparsebucket
	case "data.latest_record":
		*c = Datalatestrecord
	case "data.missing_fields":
		*c = Datamissingfields
	case "data.out_of_order_timestamps":
		*c = Dataoutofordertimestamps
	case "data.processed_fields":
		*c = Dataprocessedfields
	case "data.processed_records":
		*c = Dataprocessedrecords
	case "data.sparse_buckets":
		*c = Datasparsebuckets
	case "forecasts.memory.avg":
		*c = Forecastsmemoryavg
	case "forecasts.memory.max":
		*c = Forecastsmemorymax
	case "forecasts.memory.min":
		*c = Forecastsmemorymin
	case "forecasts.memory.total":
		*c = Forecastsmemorytotal
	case "forecasts.records.avg":
		*c = Forecastsrecordsavg
	case "forecasts.records.max":
		*c = Forecastsrecordsmax
	case "forecasts.records.min":
		*c = Forecastsrecordsmin
	case "forecasts.records.total":
		*c = Forecastsrecordstotal
	case "forecasts.time.avg":
		*c = Forecaststimeavg
	case "forecasts.time.max":
		*c = Forecaststimemax
	case "forecasts.time.min":
		*c = Forecaststimemin
	case "forecasts.time.total":
		*c = Forecaststimetotal
	case "forecasts.total":
		*c = Forecaststotal
	case "id":
		*c = Id
	case "model.bucket_allocation_failures":
		*c = Modelbucketallocationfailures
	case "model.by_fields":
		*c = Modelbyfields
	case "model.bytes":
		*c = Modelbytes
	case "model.bytes_exceeded":
		*c = Modelbytesexceeded
	case "model.categorization_status":
		*c = Modelcategorizationstatus
	case "model.categorized_doc_count":
		*c = Modelcategorizeddoccount
	case "model.dead_category_count":
		*c = Modeldeadcategorycount
	case "model.failed_category_count":
		*c = Modelfailedcategorycount
	case "model.frequent_category_count":
		*c = Modelfrequentcategorycount
	case "model.log_time":
		*c = Modellogtime
	case "model.memory_limit":
		*c = Modelmemorylimit
	case "model.memory_status":
		*c = Modelmemorystatus
	case "model.over_fields":
		*c = Modeloverfields
	case "model.partition_fields":
		*c = Modelpartitionfields
	case "model.rare_category_count":
		*c = Modelrarecategorycount
	case "model.timestamp":
		*c = Modeltimestamp
	case "model.total_category_count":
		*c = Modeltotalcategorycount
	case "node.address":
		*c = Nodeaddress
	case "node.ephemeral_id":
		*c = Nodeephemeralid
	case "node.id":
		*c = Nodeid
	case "node.name":
		*c = Nodename
	case "opened_time":
		*c = Openedtime
	case "state":
		*c = State
	default:
		*c = CatAnomalyDetectorColumn{string(text)}
	}

	return nil
}

func (c CatAnomalyDetectorColumn) String() string {
	return c.Name
}

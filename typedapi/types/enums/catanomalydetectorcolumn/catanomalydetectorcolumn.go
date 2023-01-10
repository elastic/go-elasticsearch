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


// Package catanomalydetectorcolumn
package catanomalydetectorcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/_types/CatBase.ts#L32-L401
type CatAnomalyDetectorColumn struct {
	Name string
}

var (
	Assignmentexplanation = CatAnomalyDetectorColumn{"assignment_explanation"}

	BucketsCount = CatAnomalyDetectorColumn{"buckets.count"}

	BucketsTimeExpavg = CatAnomalyDetectorColumn{"buckets.time.exp_avg"}

	BucketsTimeExpavghour = CatAnomalyDetectorColumn{"buckets.time.exp_avg_hour"}

	BucketsTimeMax = CatAnomalyDetectorColumn{"buckets.time.max"}

	BucketsTimeMin = CatAnomalyDetectorColumn{"buckets.time.min"}

	BucketsTimeTotal = CatAnomalyDetectorColumn{"buckets.time.total"}

	DataBuckets = CatAnomalyDetectorColumn{"data.buckets"}

	DataEarliestrecord = CatAnomalyDetectorColumn{"data.earliest_record"}

	DataEmptybuckets = CatAnomalyDetectorColumn{"data.empty_buckets"}

	DataInputbytes = CatAnomalyDetectorColumn{"data.input_bytes"}

	DataInputfields = CatAnomalyDetectorColumn{"data.input_fields"}

	DataInputrecords = CatAnomalyDetectorColumn{"data.input_records"}

	DataInvaliddates = CatAnomalyDetectorColumn{"data.invalid_dates"}

	DataLast = CatAnomalyDetectorColumn{"data.last"}

	DataLastemptybucket = CatAnomalyDetectorColumn{"data.last_empty_bucket"}

	DataLastsparsebucket = CatAnomalyDetectorColumn{"data.last_sparse_bucket"}

	DataLatestrecord = CatAnomalyDetectorColumn{"data.latest_record"}

	DataMissingfields = CatAnomalyDetectorColumn{"data.missing_fields"}

	DataOutofordertimestamps = CatAnomalyDetectorColumn{"data.out_of_order_timestamps"}

	DataProcessedfields = CatAnomalyDetectorColumn{"data.processed_fields"}

	DataProcessedrecords = CatAnomalyDetectorColumn{"data.processed_records"}

	DataSparsebuckets = CatAnomalyDetectorColumn{"data.sparse_buckets"}

	ForecastsMemoryAvg = CatAnomalyDetectorColumn{"forecasts.memory.avg"}

	ForecastsMemoryMax = CatAnomalyDetectorColumn{"forecasts.memory.max"}

	ForecastsMemoryMin = CatAnomalyDetectorColumn{"forecasts.memory.min"}

	ForecastsMemoryTotal = CatAnomalyDetectorColumn{"forecasts.memory.total"}

	ForecastsRecordsAvg = CatAnomalyDetectorColumn{"forecasts.records.avg"}

	ForecastsRecordsMax = CatAnomalyDetectorColumn{"forecasts.records.max"}

	ForecastsRecordsMin = CatAnomalyDetectorColumn{"forecasts.records.min"}

	ForecastsRecordsTotal = CatAnomalyDetectorColumn{"forecasts.records.total"}

	ForecastsTimeAvg = CatAnomalyDetectorColumn{"forecasts.time.avg"}

	ForecastsTimeMax = CatAnomalyDetectorColumn{"forecasts.time.max"}

	ForecastsTimeMin = CatAnomalyDetectorColumn{"forecasts.time.min"}

	ForecastsTimeTotal = CatAnomalyDetectorColumn{"forecasts.time.total"}

	ForecastsTotal = CatAnomalyDetectorColumn{"forecasts.total"}

	Id = CatAnomalyDetectorColumn{"id"}

	ModelBucketallocationfailures = CatAnomalyDetectorColumn{"model.bucket_allocation_failures"}

	ModelByfields = CatAnomalyDetectorColumn{"model.by_fields"}

	ModelBytes = CatAnomalyDetectorColumn{"model.bytes"}

	ModelBytesexceeded = CatAnomalyDetectorColumn{"model.bytes_exceeded"}

	ModelCategorizationstatus = CatAnomalyDetectorColumn{"model.categorization_status"}

	ModelCategorizeddoccount = CatAnomalyDetectorColumn{"model.categorized_doc_count"}

	ModelDeadcategorycount = CatAnomalyDetectorColumn{"model.dead_category_count"}

	ModelFailedcategorycount = CatAnomalyDetectorColumn{"model.failed_category_count"}

	ModelFrequentcategorycount = CatAnomalyDetectorColumn{"model.frequent_category_count"}

	ModelLogtime = CatAnomalyDetectorColumn{"model.log_time"}

	ModelMemorylimit = CatAnomalyDetectorColumn{"model.memory_limit"}

	ModelMemorystatus = CatAnomalyDetectorColumn{"model.memory_status"}

	ModelOverfields = CatAnomalyDetectorColumn{"model.over_fields"}

	ModelPartitionfields = CatAnomalyDetectorColumn{"model.partition_fields"}

	ModelRarecategorycount = CatAnomalyDetectorColumn{"model.rare_category_count"}

	ModelTimestamp = CatAnomalyDetectorColumn{"model.timestamp"}

	ModelTotalcategorycount = CatAnomalyDetectorColumn{"model.total_category_count"}

	NodeAddress = CatAnomalyDetectorColumn{"node.address"}

	NodeEphemeralid = CatAnomalyDetectorColumn{"node.ephemeral_id"}

	NodeId = CatAnomalyDetectorColumn{"node.id"}

	NodeName = CatAnomalyDetectorColumn{"node.name"}

	Openedtime = CatAnomalyDetectorColumn{"opened_time"}

	State = CatAnomalyDetectorColumn{"state"}
)

func (c CatAnomalyDetectorColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatAnomalyDetectorColumn) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "assignment_explanation":
		*c = Assignmentexplanation
	case "buckets.count":
		*c = BucketsCount
	case "buckets.time.exp_avg":
		*c = BucketsTimeExpavg
	case "buckets.time.exp_avg_hour":
		*c = BucketsTimeExpavghour
	case "buckets.time.max":
		*c = BucketsTimeMax
	case "buckets.time.min":
		*c = BucketsTimeMin
	case "buckets.time.total":
		*c = BucketsTimeTotal
	case "data.buckets":
		*c = DataBuckets
	case "data.earliest_record":
		*c = DataEarliestrecord
	case "data.empty_buckets":
		*c = DataEmptybuckets
	case "data.input_bytes":
		*c = DataInputbytes
	case "data.input_fields":
		*c = DataInputfields
	case "data.input_records":
		*c = DataInputrecords
	case "data.invalid_dates":
		*c = DataInvaliddates
	case "data.last":
		*c = DataLast
	case "data.last_empty_bucket":
		*c = DataLastemptybucket
	case "data.last_sparse_bucket":
		*c = DataLastsparsebucket
	case "data.latest_record":
		*c = DataLatestrecord
	case "data.missing_fields":
		*c = DataMissingfields
	case "data.out_of_order_timestamps":
		*c = DataOutofordertimestamps
	case "data.processed_fields":
		*c = DataProcessedfields
	case "data.processed_records":
		*c = DataProcessedrecords
	case "data.sparse_buckets":
		*c = DataSparsebuckets
	case "forecasts.memory.avg":
		*c = ForecastsMemoryAvg
	case "forecasts.memory.max":
		*c = ForecastsMemoryMax
	case "forecasts.memory.min":
		*c = ForecastsMemoryMin
	case "forecasts.memory.total":
		*c = ForecastsMemoryTotal
	case "forecasts.records.avg":
		*c = ForecastsRecordsAvg
	case "forecasts.records.max":
		*c = ForecastsRecordsMax
	case "forecasts.records.min":
		*c = ForecastsRecordsMin
	case "forecasts.records.total":
		*c = ForecastsRecordsTotal
	case "forecasts.time.avg":
		*c = ForecastsTimeAvg
	case "forecasts.time.max":
		*c = ForecastsTimeMax
	case "forecasts.time.min":
		*c = ForecastsTimeMin
	case "forecasts.time.total":
		*c = ForecastsTimeTotal
	case "forecasts.total":
		*c = ForecastsTotal
	case "id":
		*c = Id
	case "model.bucket_allocation_failures":
		*c = ModelBucketallocationfailures
	case "model.by_fields":
		*c = ModelByfields
	case "model.bytes":
		*c = ModelBytes
	case "model.bytes_exceeded":
		*c = ModelBytesexceeded
	case "model.categorization_status":
		*c = ModelCategorizationstatus
	case "model.categorized_doc_count":
		*c = ModelCategorizeddoccount
	case "model.dead_category_count":
		*c = ModelDeadcategorycount
	case "model.failed_category_count":
		*c = ModelFailedcategorycount
	case "model.frequent_category_count":
		*c = ModelFrequentcategorycount
	case "model.log_time":
		*c = ModelLogtime
	case "model.memory_limit":
		*c = ModelMemorylimit
	case "model.memory_status":
		*c = ModelMemorystatus
	case "model.over_fields":
		*c = ModelOverfields
	case "model.partition_fields":
		*c = ModelPartitionfields
	case "model.rare_category_count":
		*c = ModelRarecategorycount
	case "model.timestamp":
		*c = ModelTimestamp
	case "model.total_category_count":
		*c = ModelTotalcategorycount
	case "node.address":
		*c = NodeAddress
	case "node.ephemeral_id":
		*c = NodeEphemeralid
	case "node.id":
		*c = NodeId
	case "node.name":
		*c = NodeName
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

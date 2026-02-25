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

// Package catanomalydetectorcolumn
package catanomalydetectorcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L32-L406
type CatAnomalyDetectorColumn struct {
	Name string
}

var (

	// Assignmentexplanation For open anomaly detection jobs only, contains messages relating to the
	// selection of a node to run the job.
	Assignmentexplanation = CatAnomalyDetectorColumn{"assignment_explanation"}

	// Bucketscount The number of bucket results produced by the job.
	Bucketscount = CatAnomalyDetectorColumn{"buckets.count"}

	// Bucketstimeexpavg Exponential moving average of all bucket processing times, in milliseconds.
	Bucketstimeexpavg = CatAnomalyDetectorColumn{"buckets.time.exp_avg"}

	// Bucketstimeexpavghour Exponentially-weighted moving average of bucket processing times calculated
	// in a 1 hour time window, in milliseconds.
	Bucketstimeexpavghour = CatAnomalyDetectorColumn{"buckets.time.exp_avg_hour"}

	// Bucketstimemax Maximum among all bucket processing times, in milliseconds.
	Bucketstimemax = CatAnomalyDetectorColumn{"buckets.time.max"}

	// Bucketstimemin Minimum among all bucket processing times, in milliseconds.
	Bucketstimemin = CatAnomalyDetectorColumn{"buckets.time.min"}

	// Bucketstimetotal Sum of all bucket processing times, in milliseconds.
	Bucketstimetotal = CatAnomalyDetectorColumn{"buckets.time.total"}

	// Databuckets The number of buckets processed.
	Databuckets = CatAnomalyDetectorColumn{"data.buckets"}

	// Dataearliestrecord The timestamp of the earliest chronologically input document.
	Dataearliestrecord = CatAnomalyDetectorColumn{"data.earliest_record"}

	// Dataemptybuckets The number of buckets which did not contain any data.
	Dataemptybuckets = CatAnomalyDetectorColumn{"data.empty_buckets"}

	// Datainputbytes The number of bytes of input data posted to the anomaly detection job.
	Datainputbytes = CatAnomalyDetectorColumn{"data.input_bytes"}

	// Datainputfields The total number of fields in input documents posted to the anomaly detection
	// job. This count includes fields that are not used in the analysis. However,
	// be aware that if you are using a datafeed, it extracts only the required
	// fields from the documents it retrieves before posting them to the job.
	Datainputfields = CatAnomalyDetectorColumn{"data.input_fields"}

	// Datainputrecords The number of input documents posted to the anomaly detection job.
	Datainputrecords = CatAnomalyDetectorColumn{"data.input_records"}

	// Datainvaliddates The number of input documents with either a missing date field or a date that
	// could not be parsed.
	Datainvaliddates = CatAnomalyDetectorColumn{"data.invalid_dates"}

	// Datalast The timestamp at which data was last analyzed, according to server time.
	Datalast = CatAnomalyDetectorColumn{"data.last"}

	// Datalastemptybucket The timestamp of the last bucket that did not contain any data.
	Datalastemptybucket = CatAnomalyDetectorColumn{"data.last_empty_bucket"}

	// Datalastsparsebucket The timestamp of the last bucket that was considered sparse.
	Datalastsparsebucket = CatAnomalyDetectorColumn{"data.last_sparse_bucket"}

	// Datalatestrecord The timestamp of the latest chronologically input document.
	Datalatestrecord = CatAnomalyDetectorColumn{"data.latest_record"}

	// Datamissingfields The number of input documents that are missing a field that the anomaly
	// detection job is configured to analyze. Input documents with missing fields
	// are still processed because it is possible that not all fields are missing.
	Datamissingfields = CatAnomalyDetectorColumn{"data.missing_fields"}

	// Dataoutofordertimestamps The number of input documents that have a timestamp chronologically preceding
	// the start of the current anomaly detection bucket offset by the latency
	// window. This information is applicable only when you provide data to the
	// anomaly detection job by using the post data API. These out of order
	// documents are discarded, since jobs require time series data to be in
	// ascending chronological order.
	Dataoutofordertimestamps = CatAnomalyDetectorColumn{"data.out_of_order_timestamps"}

	// Dataprocessedfields The total number of fields in all the documents that have been processed by
	// the anomaly detection job. Only fields that are specified in the detector
	// configuration object contribute to this count. The timestamp is not included
	// in this count.
	Dataprocessedfields = CatAnomalyDetectorColumn{"data.processed_fields"}

	// Dataprocessedrecords The number of input documents that have been processed by the anomaly
	// detection job. This value includes documents with missing fields, since they
	// are nonetheless analyzed. If you use datafeeds and have aggregations in your
	// search query, the processed record count is the number of aggregation results
	// processed, not the number of Elasticsearch documents.
	Dataprocessedrecords = CatAnomalyDetectorColumn{"data.processed_records"}

	// Datasparsebuckets The number of buckets that contained few data points compared to the expected
	// number of data points.
	Datasparsebuckets = CatAnomalyDetectorColumn{"data.sparse_buckets"}

	// Forecastsmemoryavg The average memory usage in bytes for forecasts related to the anomaly
	// detection job.
	Forecastsmemoryavg = CatAnomalyDetectorColumn{"forecasts.memory.avg"}

	// Forecastsmemorymax The maximum memory usage in bytes for forecasts related to the anomaly
	// detection job.
	Forecastsmemorymax = CatAnomalyDetectorColumn{"forecasts.memory.max"}

	// Forecastsmemorymin The minimum memory usage in bytes for forecasts related to the anomaly
	// detection job.
	Forecastsmemorymin = CatAnomalyDetectorColumn{"forecasts.memory.min"}

	// Forecastsmemorytotal The total memory usage in bytes for forecasts related to the anomaly
	// detection job.
	Forecastsmemorytotal = CatAnomalyDetectorColumn{"forecasts.memory.total"}

	// Forecastsrecordsavg The average number of `m`odel_forecast` documents written for forecasts
	// related to the anomaly detection job.
	Forecastsrecordsavg = CatAnomalyDetectorColumn{"forecasts.records.avg"}

	// Forecastsrecordsmax The maximum number of `model_forecast` documents written for forecasts
	// related to the anomaly detection job.
	Forecastsrecordsmax = CatAnomalyDetectorColumn{"forecasts.records.max"}

	// Forecastsrecordsmin The minimum number of `model_forecast` documents written for forecasts
	// related to the anomaly detection job.
	Forecastsrecordsmin = CatAnomalyDetectorColumn{"forecasts.records.min"}

	// Forecastsrecordstotal The total number of `model_forecast` documents written for forecasts related
	// to the anomaly detection job.
	Forecastsrecordstotal = CatAnomalyDetectorColumn{"forecasts.records.total"}

	// Forecaststimeavg The average runtime in milliseconds for forecasts related to the anomaly
	// detection job.
	Forecaststimeavg = CatAnomalyDetectorColumn{"forecasts.time.avg"}

	// Forecaststimemax The maximum runtime in milliseconds for forecasts related to the anomaly
	// detection job.
	Forecaststimemax = CatAnomalyDetectorColumn{"forecasts.time.max"}

	// Forecaststimemin The minimum runtime in milliseconds for forecasts related to the anomaly
	// detection job.
	Forecaststimemin = CatAnomalyDetectorColumn{"forecasts.time.min"}

	// Forecaststimetotal The total runtime in milliseconds for forecasts related to the anomaly
	// detection job.
	Forecaststimetotal = CatAnomalyDetectorColumn{"forecasts.time.total"}

	// Forecaststotal The number of individual forecasts currently available for the job.
	Forecaststotal = CatAnomalyDetectorColumn{"forecasts.total"}

	// Id Identifier for the anomaly detection job.
	Id = CatAnomalyDetectorColumn{"id"}

	// Modelbucketallocationfailures The number of buckets for which new entities in incoming data were not
	// processed due to insufficient model memory.
	Modelbucketallocationfailures = CatAnomalyDetectorColumn{"model.bucket_allocation_failures"}

	// Modelbyfields The number of by field values that were analyzed by the models. This value is
	// cumulative for all detectors in the job.
	Modelbyfields = CatAnomalyDetectorColumn{"model.by_fields"}

	// Modelbytes The number of bytes of memory used by the models. This is the maximum value
	// since the last time the model was persisted. If the job is closed, this value
	// indicates the latest size.
	Modelbytes = CatAnomalyDetectorColumn{"model.bytes"}

	// Modelbytesexceeded The number of bytes over the high limit for memory usage at the last
	// allocation failure.
	Modelbytesexceeded = CatAnomalyDetectorColumn{"model.bytes_exceeded"}

	// Modelcategorizationstatus The status of categorization for the job: `ok` or `warn`. If `ok`,
	// categorization is performing acceptably well (or not being used at all). If
	// `warn`, categorization is detecting a distribution of categories that
	// suggests the input data is inappropriate for categorization. Problems could
	// be that there is only one category, more than 90% of categories are rare, the
	// number of categories is greater than 50% of the number of categorized
	// documents, there are no frequently matched categories, or more than 50% of
	// categories are dead.
	Modelcategorizationstatus = CatAnomalyDetectorColumn{"model.categorization_status"}

	// Modelcategorizeddoccount The number of documents that have had a field categorized.
	Modelcategorizeddoccount = CatAnomalyDetectorColumn{"model.categorized_doc_count"}

	// Modeldeadcategorycount The number of categories created by categorization that will never be
	// assigned again because another category’s definition makes it a superset of
	// the dead category. Dead categories are a side effect of the way
	// categorization has no prior training.
	Modeldeadcategorycount = CatAnomalyDetectorColumn{"model.dead_category_count"}

	// Modelfailedcategorycount The number of times that categorization wanted to create a new category but
	// couldn’t because the job had hit its model memory limit. This count does
	// not track which specific categories failed to be created. Therefore, you
	// cannot use this value to determine the number of unique categories that were
	// missed.
	Modelfailedcategorycount = CatAnomalyDetectorColumn{"model.failed_category_count"}

	// Modelfrequentcategorycount The number of categories that match more than 1% of categorized documents.
	Modelfrequentcategorycount = CatAnomalyDetectorColumn{"model.frequent_category_count"}

	// Modellogtime The timestamp when the model stats were gathered, according to server time.
	Modellogtime = CatAnomalyDetectorColumn{"model.log_time"}

	// Modelmemorylimit The timestamp when the model stats were gathered, according to server time.
	Modelmemorylimit = CatAnomalyDetectorColumn{"model.memory_limit"}

	// Modelmemorystatus The status of the mathematical models: `ok`, `soft_limit`, or `hard_limit`.
	// If `ok`, the models stayed below the configured value. If `soft_limit`, the
	// models used more than 60% of the configured memory limit and older unused
	// models will be pruned to free up space. Additionally, in categorization jobs
	// no further category examples will be stored. If `hard_limit`, the models used
	// more space than the configured memory limit. As a result, not all incoming
	// data was processed.
	Modelmemorystatus = CatAnomalyDetectorColumn{"model.memory_status"}

	// Modeloverfields The number of over field values that were analyzed by the models. This value
	// is cumulative for all detectors in the job.
	Modeloverfields = CatAnomalyDetectorColumn{"model.over_fields"}

	// Modelpartitionfields The number of partition field values that were analyzed by the models. This
	// value is cumulative for all detectors in the job.
	Modelpartitionfields = CatAnomalyDetectorColumn{"model.partition_fields"}

	// Modelrarecategorycount The number of categories that match just one categorized document.
	Modelrarecategorycount = CatAnomalyDetectorColumn{"model.rare_category_count"}

	// Modeltimestamp The timestamp of the last record when the model stats were gathered.
	Modeltimestamp = CatAnomalyDetectorColumn{"model.timestamp"}

	// Modeltotalcategorycount The number of categories created by categorization.
	Modeltotalcategorycount = CatAnomalyDetectorColumn{"model.total_category_count"}

	// Nodeaddress The network address of the node that runs the job. This information is
	// available only for open jobs.
	Nodeaddress = CatAnomalyDetectorColumn{"node.address"}

	// Nodeephemeralid The ephemeral ID of the node that runs the job. This information is available
	// only for open jobs.
	Nodeephemeralid = CatAnomalyDetectorColumn{"node.ephemeral_id"}

	// Nodeid The unique identifier of the node that runs the job. This information is
	// available only for open jobs.
	Nodeid = CatAnomalyDetectorColumn{"node.id"}

	// Nodename The name of the node that runs the job. This information is available only
	// for open jobs.
	Nodename = CatAnomalyDetectorColumn{"node.name"}

	// Openedtime For open jobs only, the elapsed time for which the job has been open.
	Openedtime = CatAnomalyDetectorColumn{"opened_time"}

	// State The status of the anomaly detection job: `closed`, `closing`, `failed`,
	// `opened`, or `opening`. If `closed`, the job finished successfully with its
	// model state persisted. The job must be opened before it can accept further
	// data. If `closing`, the job close action is in progress and has not yet
	// completed. A closing job cannot accept further data. If `failed`, the job did
	// not finish successfully due to an error. This situation can occur due to
	// invalid input data, a fatal error occurring during the analysis, or an
	// external interaction such as the process being killed by the Linux out of
	// memory (OOM) killer. If the job had irrevocably failed, it must be force
	// closed and then deleted. If the datafeed can be corrected, the job can be
	// closed and then re-opened. If `opened`, the job is available to receive and
	// process data. If `opening`, the job open action is in progress and has not
	// yet completed.
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

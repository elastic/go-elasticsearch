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

// Package cattransformcolumn
package cattransformcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L2627-L2831
type CatTransformColumn struct {
	Name string
}

var (

	// Changeslastdetectiontime The timestamp when changes were last detected in the source indices.
	Changeslastdetectiontime = CatTransformColumn{"changes_last_detection_time"}

	// Checkpoint The sequence number for the checkpoint.
	Checkpoint = CatTransformColumn{"checkpoint"}

	// Checkpointdurationtimeexpavg Exponential moving average of the duration of the checkpoint, in
	// milliseconds.
	Checkpointdurationtimeexpavg = CatTransformColumn{"checkpoint_duration_time_exp_avg"}

	// Checkpointprogress The progress of the next checkpoint that is currently in progress.
	Checkpointprogress = CatTransformColumn{"checkpoint_progress"}

	// Createtime The time the transform was created.
	Createtime = CatTransformColumn{"create_time"}

	// Deletetime The amount of time spent deleting, in milliseconds.
	Deletetime = CatTransformColumn{"delete_time"}

	// Description The description of the transform.
	Description = CatTransformColumn{"description"}

	// Destindex The destination index for the transform. The mappings of the destination
	// index are deduced based on the source fields when possible. If alternate
	// mappings are required, use the Create index API prior to starting the
	// transform.
	Destindex = CatTransformColumn{"dest_index"}

	// Documentsdeleted The number of documents that have been deleted from the destination index due
	// to the retention policy for this transform.
	Documentsdeleted = CatTransformColumn{"documents_deleted"}

	// Documentsindexed The number of documents that have been indexed into the destination index for
	// the transform.
	Documentsindexed = CatTransformColumn{"documents_indexed"}

	// Docspersecond Specifies a limit on the number of input documents per second. This setting
	// throttles the transform by adding a wait time between search requests. The
	// default value is `null`, which disables throttling.
	Docspersecond = CatTransformColumn{"docs_per_second"}

	// Documentsprocessed The number of documents that have been processed from the source index of the
	// transform.
	Documentsprocessed = CatTransformColumn{"documents_processed"}

	// Frequency The interval between checks for changes in the source indices when the
	// transform is running continuously. Also determines the retry interval in the
	// event of transient failures while the transform is searching or indexing. The
	// minimum value is `1s` and the maximum is `1h`. The default value is `1m`.
	Frequency = CatTransformColumn{"frequency"}

	// Id Identifier for the transform.
	Id = CatTransformColumn{"id"}

	// Indexfailure The number of indexing failures.
	Indexfailure = CatTransformColumn{"index_failure"}

	// Indextime The amount of time spent indexing, in milliseconds.
	Indextime = CatTransformColumn{"index_time"}

	// Indextotal The number of index operations.
	Indextotal = CatTransformColumn{"index_total"}

	// Indexeddocumentsexpavg Exponential moving average of the number of new documents that have been
	// indexed.
	Indexeddocumentsexpavg = CatTransformColumn{"indexed_documents_exp_avg"}

	// Lastsearchtime The timestamp of the last search in the source indices. This field is only
	// shown if the transform is running.
	Lastsearchtime = CatTransformColumn{"last_search_time"}

	// Maxpagesearchsize Defines the initial page size to use for the composite aggregation for each
	// checkpoint. If circuit breaker exceptions occur, the page size is dynamically
	// adjusted to a lower value. The minimum value is `10` and the maximum is
	// `65,536`. The default value is `500`.
	Maxpagesearchsize = CatTransformColumn{"max_page_search_size"}

	// Pagesprocessed The number of search or bulk index operations processed. Documents are
	// processed in batches instead of individually.
	Pagesprocessed = CatTransformColumn{"pages_processed"}

	// Pipeline The unique identifier for an ingest pipeline.
	Pipeline = CatTransformColumn{"pipeline"}

	// Processeddocumentsexpavg Exponential moving average of the number of documents that have been
	// processed.
	Processeddocumentsexpavg = CatTransformColumn{"processed_documents_exp_avg"}

	// Processingtime The amount of time spent processing results, in milliseconds.
	Processingtime = CatTransformColumn{"processing_time"}

	// Reason If a transform has a `failed` state, this property provides details about the
	// reason for the failure.
	Reason = CatTransformColumn{"reason"}

	// Searchfailure The number of search failures.
	Searchfailure = CatTransformColumn{"search_failure"}

	// Searchtime The amount of time spent searching, in milliseconds.
	Searchtime = CatTransformColumn{"search_time"}

	// Searchtotal The number of search operations on the source index for the transform.
	Searchtotal = CatTransformColumn{"search_total"}

	// Sourceindex The source indices for the transform. It can be a single index, an index
	// pattern (for example, `"my-index-*"`), an array of indices (for example,
	// `["my-index-000001", "my-index-000002"]`), or an array of index patterns (for
	// example, `["my-index-*", "my-other-index-*"]`. For remote indices use the
	// syntax `"remote_name:index_name"`. If any indices are in remote clusters then
	// the master node and at least one transform node must have the
	// `remote_cluster_client` node role.
	Sourceindex = CatTransformColumn{"source_index"}

	// State The status of the transform, which can be one of the following values:
	//
	//   - `aborting`: The transform is aborting.
	//   - `failed`: The transform failed. For more information about the failure,
	//     check the reason field.
	//   - `indexing`: The transform is actively processing data and creating new
	//     documents.
	//   - `started`: The transform is running but not actively indexing data.
	//   - `stopped`: The transform is stopped.
	//   - `stopping`: The transform is stopping.
	State = CatTransformColumn{"state"}

	// Transformtype Indicates the type of transform: `batch` or `continuous`.
	Transformtype = CatTransformColumn{"transform_type"}

	// Triggercount The number of times the transform has been triggered by the scheduler. For
	// example, the scheduler triggers the transform indexer to check for updates or
	// ingest new data at an interval specified in the `frequency` property.
	Triggercount = CatTransformColumn{"trigger_count"}

	// Version The version of Elasticsearch that existed on the node when the transform was
	// created.
	Version = CatTransformColumn{"version"}
)

func (c CatTransformColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatTransformColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "changes_last_detection_time":
		*c = Changeslastdetectiontime
	case "checkpoint":
		*c = Checkpoint
	case "checkpoint_duration_time_exp_avg":
		*c = Checkpointdurationtimeexpavg
	case "checkpoint_progress":
		*c = Checkpointprogress
	case "create_time":
		*c = Createtime
	case "delete_time":
		*c = Deletetime
	case "description":
		*c = Description
	case "dest_index":
		*c = Destindex
	case "documents_deleted":
		*c = Documentsdeleted
	case "documents_indexed":
		*c = Documentsindexed
	case "docs_per_second":
		*c = Docspersecond
	case "documents_processed":
		*c = Documentsprocessed
	case "frequency":
		*c = Frequency
	case "id":
		*c = Id
	case "index_failure":
		*c = Indexfailure
	case "index_time":
		*c = Indextime
	case "index_total":
		*c = Indextotal
	case "indexed_documents_exp_avg":
		*c = Indexeddocumentsexpavg
	case "last_search_time":
		*c = Lastsearchtime
	case "max_page_search_size":
		*c = Maxpagesearchsize
	case "pages_processed":
		*c = Pagesprocessed
	case "pipeline":
		*c = Pipeline
	case "processed_documents_exp_avg":
		*c = Processeddocumentsexpavg
	case "processing_time":
		*c = Processingtime
	case "reason":
		*c = Reason
	case "search_failure":
		*c = Searchfailure
	case "search_time":
		*c = Searchtime
	case "search_total":
		*c = Searchtotal
	case "source_index":
		*c = Sourceindex
	case "state":
		*c = State
	case "transform_type":
		*c = Transformtype
	case "trigger_count":
		*c = Triggercount
	case "version":
		*c = Version
	default:
		*c = CatTransformColumn{string(text)}
	}

	return nil
}

func (c CatTransformColumn) String() string {
	return c.Name
}

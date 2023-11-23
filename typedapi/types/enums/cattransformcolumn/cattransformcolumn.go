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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package cattransformcolumn
package cattransformcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cat/_types/CatBase.ts#L640-L844
type CatTransformColumn struct {
	Name string
}

var (
	Changeslastdetectiontime = CatTransformColumn{"changes_last_detection_time"}

	Checkpoint = CatTransformColumn{"checkpoint"}

	Checkpointdurationtimeexpavg = CatTransformColumn{"checkpoint_duration_time_exp_avg"}

	Checkpointprogress = CatTransformColumn{"checkpoint_progress"}

	Createtime = CatTransformColumn{"create_time"}

	Deletetime = CatTransformColumn{"delete_time"}

	Description = CatTransformColumn{"description"}

	Destindex = CatTransformColumn{"dest_index"}

	Documentsdeleted = CatTransformColumn{"documents_deleted"}

	Documentsindexed = CatTransformColumn{"documents_indexed"}

	Docspersecond = CatTransformColumn{"docs_per_second"}

	Documentsprocessed = CatTransformColumn{"documents_processed"}

	Frequency = CatTransformColumn{"frequency"}

	Id = CatTransformColumn{"id"}

	Indexfailure = CatTransformColumn{"index_failure"}

	Indextime = CatTransformColumn{"index_time"}

	Indextotal = CatTransformColumn{"index_total"}

	Indexeddocumentsexpavg = CatTransformColumn{"indexed_documents_exp_avg"}

	Lastsearchtime = CatTransformColumn{"last_search_time"}

	Maxpagesearchsize = CatTransformColumn{"max_page_search_size"}

	Pagesprocessed = CatTransformColumn{"pages_processed"}

	Pipeline = CatTransformColumn{"pipeline"}

	Processeddocumentsexpavg = CatTransformColumn{"processed_documents_exp_avg"}

	Processingtime = CatTransformColumn{"processing_time"}

	Reason = CatTransformColumn{"reason"}

	Searchfailure = CatTransformColumn{"search_failure"}

	Searchtime = CatTransformColumn{"search_time"}

	Searchtotal = CatTransformColumn{"search_total"}

	Sourceindex = CatTransformColumn{"source_index"}

	State = CatTransformColumn{"state"}

	Transformtype = CatTransformColumn{"transform_type"}

	Triggercount = CatTransformColumn{"trigger_count"}

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

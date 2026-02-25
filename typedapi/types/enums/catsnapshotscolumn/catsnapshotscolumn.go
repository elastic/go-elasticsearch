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

// Package catsnapshotscolumn
package catsnapshotscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L1093-L1160
type CatSnapshotsColumn struct {
	Name string
}

var (

	// Id The ID of the snapshot, such as 'snap1'.
	Id = CatSnapshotsColumn{"id"}

	// Repository The name of the repository, such as 'repo1'.
	Repository = CatSnapshotsColumn{"repository"}

	// Status State of the snapshot process. Returned values are: 'FAILED': The snapshot
	// process failed. 'INCOMPATIBLE': The snapshot process is incompatible with the
	// current cluster version. 'IN_PROGRESS': The snapshot process started but has
	// not completed. 'PARTIAL': The snapshot process completed with a partial
	// success. 'SUCCESS': The snapshot process completed with a full success.
	Status = CatSnapshotsColumn{"status"}

	// Startepoch The [unix epoch time](https://en.wikipedia.org/wiki/Unix_time) at which the
	// snapshot process started.
	Startepoch = CatSnapshotsColumn{"start_epoch"}

	// Starttime 'HH:MM:SS' time at which the snapshot process started.
	Starttime = CatSnapshotsColumn{"start_time"}

	// Endepoch The [unix epoch time](https://en.wikipedia.org/wiki/Unix_time) at which the
	// snapshot process ended.
	Endepoch = CatSnapshotsColumn{"end_epoch"}

	// Endtime 'HH:MM:SS' time at which the snapshot process ended.
	Endtime = CatSnapshotsColumn{"end_time"}

	// Duration The time it took the snapshot process to complete in [time
	// units](https://www.elastic.co/docs/reference/elasticsearch/rest-apis/api-conventions#time-units).
	Duration = CatSnapshotsColumn{"duration"}

	// Indices The number of indices in the snapshot.
	Indices = CatSnapshotsColumn{"indices"}

	// Successfulshards The number of successful shards in the snapshot.
	Successfulshards = CatSnapshotsColumn{"successful_shards"}

	// Failedshards The number of failed shards in the snapshot.
	Failedshards = CatSnapshotsColumn{"failed_shards"}

	// Totalshards The total number of shards in the snapshot.
	Totalshards = CatSnapshotsColumn{"total_shards"}

	// Reason The reason for any snapshot failures.
	Reason = CatSnapshotsColumn{"reason"}
)

func (c CatSnapshotsColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatSnapshotsColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "id":
		*c = Id
	case "repository":
		*c = Repository
	case "status":
		*c = Status
	case "start_epoch":
		*c = Startepoch
	case "start_time":
		*c = Starttime
	case "end_epoch":
		*c = Endepoch
	case "end_time":
		*c = Endtime
	case "duration":
		*c = Duration
	case "indices":
		*c = Indices
	case "successful_shards":
		*c = Successfulshards
	case "failed_shards":
		*c = Failedshards
	case "total_shards":
		*c = Totalshards
	case "reason":
		*c = Reason
	default:
		*c = CatSnapshotsColumn{string(text)}
	}

	return nil
}

func (c CatSnapshotsColumn) String() string {
	return c.Name
}

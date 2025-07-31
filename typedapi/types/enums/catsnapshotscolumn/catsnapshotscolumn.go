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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package catsnapshotscolumn
package catsnapshotscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cat/_types/CatBase.ts#L1214-L1281
type CatSnapshotsColumn struct {
	Name string
}

var (
	Id = CatSnapshotsColumn{"id"}

	Repository = CatSnapshotsColumn{"repository"}

	Status = CatSnapshotsColumn{"status"}

	Startepoch = CatSnapshotsColumn{"start_epoch"}

	Starttime = CatSnapshotsColumn{"start_time"}

	Endepoch = CatSnapshotsColumn{"end_epoch"}

	Endtime = CatSnapshotsColumn{"end_time"}

	Duration = CatSnapshotsColumn{"duration"}

	Indices = CatSnapshotsColumn{"indices"}

	Successfulshards = CatSnapshotsColumn{"successful_shards"}

	Failedshards = CatSnapshotsColumn{"failed_shards"}

	Totalshards = CatSnapshotsColumn{"total_shards"}

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

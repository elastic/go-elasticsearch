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
// https://github.com/elastic/elasticsearch-specification/tree/6ee016a765be615b0205fc209d3d3c515044689d

// Package catrecoverycolumn
package catrecoverycolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/6ee016a765be615b0205fc209d3d3c515044689d/specification/cat/_types/CatBase.ts#L1014-L1145
type CatRecoveryColumn struct {
	Name string
}

var (

	// Index The name of the index.
	Index = CatRecoveryColumn{"index"}

	// Shard The name of the shard.
	Shard = CatRecoveryColumn{"shard"}

	// Time The recovery time elasped.
	Time = CatRecoveryColumn{"time"}

	// Type The type of recovery, from a peer or a snapshot.
	Type = CatRecoveryColumn{"type"}

	// Stage The stage of the recovery. Returned values are: `INIT`, `INDEX`: recovery of
	// lucene files, either reusing local ones are copying new ones, `VERIFY_INDEX`:
	// potentially running check index, `TRANSLOG`: starting up the engine,
	// replaying the translog, `FINALIZE`: performing final task after all translog
	// ops have been done, `DONE`
	Stage = CatRecoveryColumn{"stage"}

	// Sourcehost The host address the index is moving from.
	Sourcehost = CatRecoveryColumn{"source_host"}

	// Sourcenode The node name the index is moving from.
	Sourcenode = CatRecoveryColumn{"source_node"}

	// Targethost The host address the index is moving to.
	Targethost = CatRecoveryColumn{"target_host"}

	// Targetnode The node name the index is moving to.
	Targetnode = CatRecoveryColumn{"target_node"}

	// Repository The name of the repository being used. if not relevant 'n/a'.
	Repository = CatRecoveryColumn{"repository"}

	// Snapshot The name of the snapshot being used. if not relevant 'n/a'.
	Snapshot = CatRecoveryColumn{"snapshot"}

	// Files The total number of files to recover.
	Files = CatRecoveryColumn{"files"}

	// Filesrecovered The number of files currently recovered.
	Filesrecovered = CatRecoveryColumn{"files_recovered"}

	// Filespercent The percentage of files currently recovered.
	Filespercent = CatRecoveryColumn{"files_percent"}

	// Filestotal The total number of files.
	Filestotal = CatRecoveryColumn{"files_total"}

	// Bytes The total number of bytes to recover.
	Bytes = CatRecoveryColumn{"bytes"}

	// Bytesrecovered Total number of bytes currently recovered.
	Bytesrecovered = CatRecoveryColumn{"bytes_recovered"}

	// Bytespercent The percentage of bytes currently recovered.
	Bytespercent = CatRecoveryColumn{"bytes_percent"}

	// Bytestotal The total number of bytes.
	Bytestotal = CatRecoveryColumn{"bytes_total"}

	// Translogops The total number of translog ops to recover.
	Translogops = CatRecoveryColumn{"translog_ops"}

	// Translogopsrecovered The total number of translog ops currently recovered.
	Translogopsrecovered = CatRecoveryColumn{"translog_ops_recovered"}

	// Translogopspercent The percentage of translog ops currently recovered.
	Translogopspercent = CatRecoveryColumn{"translog_ops_percent"}

	// Starttime The start time of the recovery operation.
	Starttime = CatRecoveryColumn{"start_time"}

	// Starttimemillis The start time of the recovery operation in eopch milliseconds.
	Starttimemillis = CatRecoveryColumn{"start_time_millis"}

	// Stoptime The end time of the recovery operation. If ongoing '1970-01-01T00:00:00.000Z'
	Stoptime = CatRecoveryColumn{"stop_time"}

	// Stoptimemillis The end time of the recovery operation in eopch milliseconds. If ongoing '0'
	Stoptimemillis = CatRecoveryColumn{"stop_time_millis"}
)

func (c CatRecoveryColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatRecoveryColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "index":
		*c = Index
	case "shard":
		*c = Shard
	case "time":
		*c = Time
	case "type":
		*c = Type
	case "stage":
		*c = Stage
	case "source_host":
		*c = Sourcehost
	case "source_node":
		*c = Sourcenode
	case "target_host":
		*c = Targethost
	case "target_node":
		*c = Targetnode
	case "repository":
		*c = Repository
	case "snapshot":
		*c = Snapshot
	case "files":
		*c = Files
	case "files_recovered":
		*c = Filesrecovered
	case "files_percent":
		*c = Filespercent
	case "files_total":
		*c = Filestotal
	case "bytes":
		*c = Bytes
	case "bytes_recovered":
		*c = Bytesrecovered
	case "bytes_percent":
		*c = Bytespercent
	case "bytes_total":
		*c = Bytestotal
	case "translog_ops":
		*c = Translogops
	case "translog_ops_recovered":
		*c = Translogopsrecovered
	case "translog_ops_percent":
		*c = Translogopspercent
	case "start_time":
		*c = Starttime
	case "start_time_millis":
		*c = Starttimemillis
	case "stop_time":
		*c = Stoptime
	case "stop_time_millis":
		*c = Stoptimemillis
	default:
		*c = CatRecoveryColumn{string(text)}
	}

	return nil
}

func (c CatRecoveryColumn) String() string {
	return c.Name
}

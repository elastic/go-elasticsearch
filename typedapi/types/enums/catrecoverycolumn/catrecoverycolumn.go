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

// Package catrecoverycolumn
package catrecoverycolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cat/_types/CatBase.ts#L1014-L1145
type CatRecoveryColumn struct {
	Name string
}

var (
	Index = CatRecoveryColumn{"index"}

	Shard = CatRecoveryColumn{"shard"}

	Time = CatRecoveryColumn{"time"}

	Type = CatRecoveryColumn{"type"}

	Stage = CatRecoveryColumn{"stage"}

	Sourcehost = CatRecoveryColumn{"source_host"}

	Sourcenode = CatRecoveryColumn{"source_node"}

	Targethost = CatRecoveryColumn{"target_host"}

	Targetnode = CatRecoveryColumn{"target_node"}

	Repository = CatRecoveryColumn{"repository"}

	Snapshot = CatRecoveryColumn{"snapshot"}

	Files = CatRecoveryColumn{"files"}

	Filesrecovered = CatRecoveryColumn{"files_recovered"}

	Filespercent = CatRecoveryColumn{"files_percent"}

	Filestotal = CatRecoveryColumn{"files_total"}

	Bytes = CatRecoveryColumn{"bytes"}

	Bytesrecovered = CatRecoveryColumn{"bytes_recovered"}

	Bytespercent = CatRecoveryColumn{"bytes_percent"}

	Bytestotal = CatRecoveryColumn{"bytes_total"}

	Translogops = CatRecoveryColumn{"translog_ops"}

	Translogopsrecovered = CatRecoveryColumn{"translog_ops_recovered"}

	Translogopspercent = CatRecoveryColumn{"translog_ops_percent"}

	Starttime = CatRecoveryColumn{"start_time"}

	Starttimemillis = CatRecoveryColumn{"start_time_millis"}

	Stoptime = CatRecoveryColumn{"stop_time"}

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

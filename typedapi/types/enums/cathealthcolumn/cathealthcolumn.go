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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

// Package cathealthcolumn
package cathealthcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/cat/_types/CatBase.ts#L1402-L1479
type CatHealthColumn struct {
	Name string
}

var (
	Epoch = CatHealthColumn{"epoch"}

	Timestamp = CatHealthColumn{"timestamp"}

	Cluster = CatHealthColumn{"cluster"}

	Status = CatHealthColumn{"status"}

	Nodetotal = CatHealthColumn{"node.total"}

	Nodedata = CatHealthColumn{"node.data"}

	Shards = CatHealthColumn{"shards"}

	Pri = CatHealthColumn{"pri"}

	Relo = CatHealthColumn{"relo"}

	Init = CatHealthColumn{"init"}

	Unassign = CatHealthColumn{"unassign"}

	Unassignpri = CatHealthColumn{"unassign.pri"}

	Pendingtasks = CatHealthColumn{"pending_tasks"}

	Maxtaskwaittime = CatHealthColumn{"max_task_wait_time"}

	Activeshardspercent = CatHealthColumn{"active_shards_percent"}
)

func (c CatHealthColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatHealthColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "epoch":
		*c = Epoch
	case "timestamp":
		*c = Timestamp
	case "cluster":
		*c = Cluster
	case "status":
		*c = Status
	case "node.total":
		*c = Nodetotal
	case "node.data":
		*c = Nodedata
	case "shards":
		*c = Shards
	case "pri":
		*c = Pri
	case "relo":
		*c = Relo
	case "init":
		*c = Init
	case "unassign":
		*c = Unassign
	case "unassign.pri":
		*c = Unassignpri
	case "pending_tasks":
		*c = Pendingtasks
	case "max_task_wait_time":
		*c = Maxtaskwaittime
	case "active_shards_percent":
		*c = Activeshardspercent
	default:
		*c = CatHealthColumn{string(text)}
	}

	return nil
}

func (c CatHealthColumn) String() string {
	return c.Name
}

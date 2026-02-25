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

// Package catdatafeedcolumn
package catdatafeedcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L410-L478
type CatDatafeedColumn struct {
	Name string
}

var (

	// Ae For started datafeeds only, contains messages relating to the selection of a
	// node.
	Ae = CatDatafeedColumn{"ae"}

	// Bc The number of buckets processed.
	Bc = CatDatafeedColumn{"bc"}

	// Id A numerical character string that uniquely identifies the datafeed.
	Id = CatDatafeedColumn{"id"}

	// Na For started datafeeds only, the network address of the node where the
	// datafeed is started.
	Na = CatDatafeedColumn{"na"}

	// Ne For started datafeeds only, the ephemeral ID of the node where the datafeed
	// is started.
	Ne = CatDatafeedColumn{"ne"}

	// Ni For started datafeeds only, the unique identifier of the node where the
	// datafeed is started.
	Ni = CatDatafeedColumn{"ni"}

	// Nn For started datafeeds only, the name of the node where the datafeed is
	// started.
	Nn = CatDatafeedColumn{"nn"}

	// Sba The average search time per bucket, in milliseconds.
	Sba = CatDatafeedColumn{"sba"}

	// Sc The number of searches run by the datafeed.
	Sc = CatDatafeedColumn{"sc"}

	// Seah The exponential average search time per hour, in milliseconds.
	Seah = CatDatafeedColumn{"seah"}

	// St The total time the datafeed spent searching, in milliseconds.
	St = CatDatafeedColumn{"st"}

	// S The status of the datafeed: `starting`, `started`, `stopping`, or `stopped`.
	// If `starting`, the datafeed has been requested to start but has not yet
	// started. If `started`, the datafeed is actively receiving data. If
	// `stopping`, the datafeed has been requested to stop gracefully and is
	// completing its final action. If `stopped`, the datafeed is stopped and will
	// not receive data until it is re-started.
	S = CatDatafeedColumn{"s"}
)

func (c CatDatafeedColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatDatafeedColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "ae":
		*c = Ae
	case "bc":
		*c = Bc
	case "id":
		*c = Id
	case "na":
		*c = Na
	case "ne":
		*c = Ne
	case "ni":
		*c = Ni
	case "nn":
		*c = Nn
	case "sba":
		*c = Sba
	case "sc":
		*c = Sc
	case "seah":
		*c = Seah
	case "st":
		*c = St
	case "s":
		*c = S
	default:
		*c = CatDatafeedColumn{string(text)}
	}

	return nil
}

func (c CatDatafeedColumn) String() string {
	return c.Name
}

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


// Package month
package month

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/watcher/_types/Schedule.ts#L70-L83
type Month struct {
	Name string
}

var (
	January = Month{"january"}

	February = Month{"february"}

	March = Month{"march"}

	April = Month{"april"}

	May = Month{"may"}

	June = Month{"june"}

	July = Month{"july"}

	August = Month{"august"}

	September = Month{"september"}

	October = Month{"october"}

	November = Month{"november"}

	December = Month{"december"}
)

func (m Month) MarshalText() (text []byte, err error) {
	return []byte(m.String()), nil
}

func (m *Month) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "january":
		*m = January
	case "february":
		*m = February
	case "march":
		*m = March
	case "april":
		*m = April
	case "may":
		*m = May
	case "june":
		*m = June
	case "july":
		*m = July
	case "august":
		*m = August
	case "september":
		*m = September
	case "october":
		*m = October
	case "november":
		*m = November
	case "december":
		*m = December
	default:
		*m = Month{string(text)}
	}

	return nil
}

func (m Month) String() string {
	return m.Name
}

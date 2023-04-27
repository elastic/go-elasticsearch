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

package types

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/totalhitsrelation"
)

// TotalHits type.
//
// https://github.com/elastic/elasticsearch-specification/blob/18d160a8583deec1bbef274d2c0e563a0cd20e2f/specification/_global/search/_types/hits.ts#L94-L97
type TotalHits struct {
	Relation totalhitsrelation.TotalHitsRelation `json:"relation"`
	Value    int64                               `json:"value"`
}

// UnmarshalJSON implements Unmarshaler interface, it handles the shortcut for total hits.
func (t *TotalHits) UnmarshalJSON(data []byte) error {
	type stub TotalHits
	tmp := stub{}
	dec := json.NewDecoder(bytes.NewReader(data))
	if _, err := strconv.Atoi(string(data)); err == nil {
		err := dec.Decode(&t.Value)
		if err != nil {
			return err
		}
		t.Relation = totalhitsrelation.Eq
	} else {
		err := dec.Decode(&tmp)
		if err != nil {
			return err
		}
		*t = TotalHits(tmp)
	}

	return nil
}

// NewTotalHits returns a TotalHits.
func NewTotalHits() *TotalHits {
	r := &TotalHits{}

	return r
}

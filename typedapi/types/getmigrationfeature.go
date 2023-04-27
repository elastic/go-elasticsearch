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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/migrationstatus"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// GetMigrationFeature type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/migration/get_feature_upgrade_status/GetFeatureUpgradeStatusResponse.ts#L37-L42
type GetMigrationFeature struct {
	FeatureName         string                          `json:"feature_name"`
	Indices             []MigrationFeatureIndexInfo     `json:"indices"`
	MigrationStatus     migrationstatus.MigrationStatus `json:"migration_status"`
	MinimumIndexVersion string                          `json:"minimum_index_version"`
}

func (s *GetMigrationFeature) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "feature_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.FeatureName = o

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "migration_status":
			if err := dec.Decode(&s.MigrationStatus); err != nil {
				return err
			}

		case "minimum_index_version":
			if err := dec.Decode(&s.MinimumIndexVersion); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewGetMigrationFeature returns a GetMigrationFeature.
func NewGetMigrationFeature() *GetMigrationFeature {
	r := &GetMigrationFeature{}

	return r
}

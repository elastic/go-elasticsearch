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
// https://github.com/elastic/elasticsearch-specification/tree/f6a370d0fba975752c644fc730f7c45610e28f36

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/modeenum"
)

// MigrateReindex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/f6a370d0fba975752c644fc730f7c45610e28f36/specification/indices/migrate_reindex/MigrateReindexRequest.ts#L39-L48
type MigrateReindex struct {
	// Mode Reindex mode. Currently only 'upgrade' is supported.
	Mode modeenum.ModeEnum `json:"mode"`
	// Source The source index or data stream (only data streams are currently supported).
	Source SourceIndex `json:"source"`
}

// NewMigrateReindex returns a MigrateReindex.
func NewMigrateReindex() *MigrateReindex {
	r := &MigrateReindex{}

	return r
}

// true

type MigrateReindexVariant interface {
	MigrateReindexCaster() *MigrateReindex
}

func (s *MigrateReindex) MigrateReindexCaster() *MigrateReindex {
	return s
}

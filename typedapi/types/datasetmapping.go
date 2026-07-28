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
// https://github.com/elastic/elasticsearch-specification/tree/9fcf6a64c550d2e8090c8134867f200b56fd7fc7

package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamic"
)

// A user-declared mapping (the `mappings` block) attached to a dataset. It is
// entirely optional: a dataset with no declared mapping relies on inference.
//
// https://github.com/elastic/elasticsearch-specification/blob/9fcf6a64c550d2e8090c8134867f200b56fd7fc7/specification/esql/_types/types.ts#L61-L80
type DatasetMapping struct {
	// Dynamic The policy for columns that are not declared in `properties`. `true` (the
	// default) infers undeclared columns and overlays the declarations; `false`
	// makes the declaration the entire schema, so undeclared columns are not
	// queryable.
	Dynamic *dynamic.Dynamic `json:"dynamic,omitempty"`
	// Id_ The `_id` meta-field configuration, sourcing the document identity from a
	// column.
	Id_ *IdPath `json:"_id,omitempty"`
	// Properties The per-column declarations, keyed by logical column name.
	Properties map[string]DatasetFieldMapping `json:"properties,omitempty"`
}

// NewDatasetMapping returns a DatasetMapping.
func NewDatasetMapping() *DatasetMapping {
	r := &DatasetMapping{
		Properties: make(map[string]DatasetFieldMapping),
	}

	return r
}

type DatasetMappingVariant interface {
	DatasetMappingCaster() *DatasetMapping
}

func (s *DatasetMapping) DatasetMappingCaster() *DatasetMapping {
	return s
}

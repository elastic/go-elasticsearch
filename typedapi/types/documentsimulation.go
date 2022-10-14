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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// DocumentSimulation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/ingest/simulate/types.ts#L47-L60
type DocumentSimulation struct {
	DocumentSimulation map[string]string         `json:"-"`
	Id_                Id                        `json:"_id"`
	Index_             IndexName                 `json:"_index"`
	Ingest_            Ingest                    `json:"_ingest"`
	Routing_           *string                   `json:"_routing,omitempty"`
	Source_            map[string]interface{}    `json:"_source"`
	VersionType_       *versiontype.VersionType  `json:"_version_type,omitempty"`
	Version_           *StringifiedVersionNumber `json:"_version,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DocumentSimulation) MarshalJSON() ([]byte, error) {
	type opt DocumentSimulation
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.DocumentSimulation {
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DocumentSimulationBuilder holds DocumentSimulation struct and provides a builder API.
type DocumentSimulationBuilder struct {
	v *DocumentSimulation
}

// NewDocumentSimulation provides a builder for the DocumentSimulation struct.
func NewDocumentSimulationBuilder() *DocumentSimulationBuilder {
	r := DocumentSimulationBuilder{
		&DocumentSimulation{
			DocumentSimulation: make(map[string]string, 0),
			Source_:            make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DocumentSimulation struct
func (rb *DocumentSimulationBuilder) Build() DocumentSimulation {
	return *rb.v
}

func (rb *DocumentSimulationBuilder) DocumentSimulation(value map[string]string) *DocumentSimulationBuilder {
	rb.v.DocumentSimulation = value
	return rb
}

func (rb *DocumentSimulationBuilder) Id_(id_ Id) *DocumentSimulationBuilder {
	rb.v.Id_ = id_
	return rb
}

func (rb *DocumentSimulationBuilder) Index_(index_ IndexName) *DocumentSimulationBuilder {
	rb.v.Index_ = index_
	return rb
}

func (rb *DocumentSimulationBuilder) Ingest_(ingest_ *IngestBuilder) *DocumentSimulationBuilder {
	v := ingest_.Build()
	rb.v.Ingest_ = v
	return rb
}

func (rb *DocumentSimulationBuilder) Routing_(routing_ string) *DocumentSimulationBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

func (rb *DocumentSimulationBuilder) Source_(value map[string]interface{}) *DocumentSimulationBuilder {
	rb.v.Source_ = value
	return rb
}

func (rb *DocumentSimulationBuilder) VersionType_(versiontype_ versiontype.VersionType) *DocumentSimulationBuilder {
	rb.v.VersionType_ = &versiontype_
	return rb
}

func (rb *DocumentSimulationBuilder) Version_(version_ *StringifiedVersionNumberBuilder) *DocumentSimulationBuilder {
	v := version_.Build()
	rb.v.Version_ = &v
	return rb
}

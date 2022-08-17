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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// TrainedModelEntities type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L386-L392
type TrainedModelEntities struct {
	ClassName        string  `json:"class_name"`
	ClassProbability float64 `json:"class_probability"`
	EndPos           int     `json:"end_pos"`
	Entity           string  `json:"entity"`
	StartPos         int     `json:"start_pos"`
}

// TrainedModelEntitiesBuilder holds TrainedModelEntities struct and provides a builder API.
type TrainedModelEntitiesBuilder struct {
	v *TrainedModelEntities
}

// NewTrainedModelEntities provides a builder for the TrainedModelEntities struct.
func NewTrainedModelEntitiesBuilder() *TrainedModelEntitiesBuilder {
	r := TrainedModelEntitiesBuilder{
		&TrainedModelEntities{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelEntities struct
func (rb *TrainedModelEntitiesBuilder) Build() TrainedModelEntities {
	return *rb.v
}

func (rb *TrainedModelEntitiesBuilder) ClassName(classname string) *TrainedModelEntitiesBuilder {
	rb.v.ClassName = classname
	return rb
}

func (rb *TrainedModelEntitiesBuilder) ClassProbability(classprobability float64) *TrainedModelEntitiesBuilder {
	rb.v.ClassProbability = classprobability
	return rb
}

func (rb *TrainedModelEntitiesBuilder) EndPos(endpos int) *TrainedModelEntitiesBuilder {
	rb.v.EndPos = endpos
	return rb
}

func (rb *TrainedModelEntitiesBuilder) Entity(entity string) *TrainedModelEntitiesBuilder {
	rb.v.Entity = entity
	return rb
}

func (rb *TrainedModelEntitiesBuilder) StartPos(startpos int) *TrainedModelEntitiesBuilder {
	rb.v.StartPos = startpos
	return rb
}

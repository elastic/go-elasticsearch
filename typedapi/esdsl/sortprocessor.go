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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _sortProcessor struct {
	v *types.SortProcessor
}

// Sorts the elements of an array ascending or descending.
// Homogeneous arrays of numbers will be sorted numerically, while arrays of
// strings or heterogeneous arrays of strings + numbers will be sorted
// lexicographically.
// Throws an error when the field is not an array.
func NewSortProcessor() *_sortProcessor {

	return &_sortProcessor{v: types.NewSortProcessor()}

}

func (s *_sortProcessor) Description(description string) *_sortProcessor {

	s.v.Description = &description

	return s
}

func (s *_sortProcessor) Field(field string) *_sortProcessor {

	s.v.Field = field

	return s
}

func (s *_sortProcessor) If(if_ types.ScriptVariant) *_sortProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_sortProcessor) IgnoreFailure(ignorefailure bool) *_sortProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_sortProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_sortProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_sortProcessor) Order(order sortorder.SortOrder) *_sortProcessor {

	s.v.Order = &order
	return s
}

func (s *_sortProcessor) Tag(tag string) *_sortProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_sortProcessor) TargetField(field string) *_sortProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_sortProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Sort = s.v

	return container
}

func (s *_sortProcessor) SortProcessorCaster() *types.SortProcessor {
	return s.v
}

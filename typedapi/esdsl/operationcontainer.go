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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _operationContainer struct {
	v *types.OperationContainer
}

func NewOperationContainer() *_operationContainer {
	return &_operationContainer{v: types.NewOperationContainer()}
}

// AdditionalOperationContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_operationContainer) AdditionalOperationContainerProperty(key string, value json.RawMessage) *_operationContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalOperationContainerProperty = tmp
	return s
}

// Index the specified document if it does not already exist.
// The following line must contain the source data to be indexed.
func (s *_operationContainer) Create(create types.CreateOperationVariant) *_operationContainer {

	s.v.Create = create.CreateOperationCaster()

	return s
}

// Remove the specified document from the index.
func (s *_operationContainer) Delete(delete types.DeleteOperationVariant) *_operationContainer {

	s.v.Delete = delete.DeleteOperationCaster()

	return s
}

// Index the specified document.
// If the document exists, it replaces the document and increments the version.
// The following line must contain the source data to be indexed.
func (s *_operationContainer) Index(index types.IndexOperationVariant) *_operationContainer {

	s.v.Index = index.IndexOperationCaster()

	return s
}

// Perform a partial document update.
// The following line must contain the partial document and update options.
func (s *_operationContainer) Update(update types.UpdateOperationVariant) *_operationContainer {

	s.v.Update = update.UpdateOperationCaster()

	return s
}

func (s *_operationContainer) OperationContainerCaster() *types.OperationContainer {
	return s.v
}

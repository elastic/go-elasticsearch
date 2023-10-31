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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

// OperationContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/bulk/types.ts#L145-L167
type OperationContainer struct {
	// Create Indexes the specified document if it does not already exist.
	// The following line must contain the source data to be indexed.
	Create *CreateOperation `json:"create,omitempty"`
	// Delete Removes the specified document from the index.
	Delete *DeleteOperation `json:"delete,omitempty"`
	// Index Indexes the specified document.
	// If the document exists, replaces the document and increments the version.
	// The following line must contain the source data to be indexed.
	Index *IndexOperation `json:"index,omitempty"`
	// Update Performs a partial document update.
	// The following line must contain the partial document and update options.
	Update *UpdateOperation `json:"update,omitempty"`
}

// NewOperationContainer returns a OperationContainer.
func NewOperationContainer() *OperationContainer {
	r := &OperationContainer{}

	return r
}

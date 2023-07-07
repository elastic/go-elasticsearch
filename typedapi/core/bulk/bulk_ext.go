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

package bulk

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// CreateOp is a helper function to add a CreateOperation to the current bulk request.
// doc argument can be a []byte, json.RawMessage or a struct.
func (r *Bulk) CreateOp(op *types.CreateOperation, doc interface{}) error {
	operation := types.OperationContainer{Create: op}
	header, err := json.Marshal(operation)
	if err != nil {
		return err
	}

	r.buf.Write(header)
	r.buf.Write([]byte("\n"))

	switch v := doc.(type) {
	case []byte:
		r.buf.Write(v)
	case json.RawMessage:
		r.buf.Write(v)
	default:
		body, err := json.Marshal(doc)
		if err != nil {
			return err
		}
		r.buf.Write(body)
	}

	r.buf.Write([]byte("\n"))

	return nil
}

// IndexOp is a helper function to add an IndexOperation to the current bulk request.
// doc argument can be a []byte, json.RawMessage or a struct.
func (r *Bulk) IndexOp(op *types.IndexOperation, doc interface{}) error {
	operation := types.OperationContainer{Index: op}
	header, err := json.Marshal(operation)
	if err != nil {
		return err
	}

	r.buf.Write(header)
	r.buf.Write([]byte("\n"))

	switch v := doc.(type) {
	case []byte:
		r.buf.Write(v)
	case json.RawMessage:
		r.buf.Write(v)
	default:
		body, err := json.Marshal(doc)
		if err != nil {
			return err
		}
		r.buf.Write(body)
	}

	r.buf.Write([]byte("\n"))

	return nil
}

// UpdateOp is a helper function to add an UpdateOperation with and UpdateAction to the current bulk request.
func (r *Bulk) UpdateOp(op *types.UpdateOperation, update *types.UpdateAction) error {
	operation := types.OperationContainer{Update: op}
	header, err := json.Marshal(operation)
	if err != nil {
		return err
	}

	r.buf.Write(header)
	r.buf.Write([]byte("\n"))

	body, err := json.Marshal(update)
	if err != nil {
		return err
	}
	r.buf.Write(body)

	r.buf.Write([]byte("\n"))

	return nil
}

// DeleteOp is a helper function to add a DeleteOperation to the current bulk request.
func (r *Bulk) DeleteOp(op *types.DeleteOperation) error {
	operation := types.OperationContainer{Delete: op}
	header, err := json.Marshal(operation)
	if err != nil {
		return err
	}

	r.buf.Write(header)
	r.buf.Write([]byte("\n"))

	return nil
}

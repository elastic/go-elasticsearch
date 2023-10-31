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
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// CreateOp is a helper function to add a CreateOperation to the current bulk request.
// doc argument can be a []byte, json.RawMessage or a struct.
func (r *Bulk) CreateOp(op types.CreateOperation, doc interface{}) error {
	operation := types.OperationContainer{Create: &op}
	header, err := json.Marshal(operation)
	if err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.CreateOp: %w", err)
	}

	if _, err := r.buf.Write(header); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.CreateOp: %w", err)
	}
	if _, err := r.buf.Write([]byte("\n")); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.CreateOp: %w", err)
	}

	switch v := doc.(type) {
	case []byte:
		if json.Valid(v) {
			if _, err := r.buf.Write(v); err != nil {
				r.buf.Reset()
				return fmt.Errorf("bulk.CreateOp: %w", err)
			}
		} else {
			r.buf.Reset()
			return fmt.Errorf("bulk.CreateOp: invalid json")
		}
	case json.RawMessage:
		if json.Valid(v) {
			if _, err := r.buf.Write(v); err != nil {
				r.buf.Reset()
				return fmt.Errorf("bulk.CreateOp: %w", err)
			}
		} else {
			r.buf.Reset()
			return fmt.Errorf("bulk.CreateOp: invalid json")
		}
	default:
		body, err := json.Marshal(doc)
		if err != nil {
			r.buf.Reset()
			return fmt.Errorf("bulk.CreateOp: %w", err)
		}
		if _, err := r.buf.Write(body); err != nil {
			r.buf.Reset()
			return fmt.Errorf("bulk.CreateOp: %w", err)
		}
	}

	if _, err := r.buf.Write([]byte("\n")); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.CreateOp: %w", err)
	}

	return nil
}

// IndexOp is a helper function to add an IndexOperation to the current bulk request.
// doc argument can be a []byte, json.RawMessage or a struct.
func (r *Bulk) IndexOp(op types.IndexOperation, doc interface{}) error {
	operation := types.OperationContainer{Index: &op}
	header, err := json.Marshal(operation)
	if err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.IndexOp: %w", err)
	}

	if _, err := r.buf.Write(header); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.IndexOp: %w", err)
	}
	if _, err := r.buf.Write([]byte("\n")); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.IndexOp: %w", err)
	}

	switch v := doc.(type) {
	case []byte:
		if json.Valid(v) {
			if _, err := r.buf.Write(v); err != nil {
				r.buf.Reset()
				return fmt.Errorf("bulk.IndexOp: %w", err)
			}
		} else {
			r.buf.Reset()
			return fmt.Errorf("bulk.IndexOp: invalid json")
		}
	case json.RawMessage:
		if json.Valid(v) {
			if _, err := r.buf.Write(v); err != nil {
				r.buf.Reset()
				return fmt.Errorf("bulk.IndexOp: %w", err)
			}
		} else {
			r.buf.Reset()
			return fmt.Errorf("bulk.IndexOp: invalid json")
		}
	default:
		body, err := json.Marshal(doc)
		if err != nil {
			r.buf.Reset()
			return fmt.Errorf("bulk.IndexOp: %w", err)
		}
		if _, err := r.buf.Write(body); err != nil {
			r.buf.Reset()
			return fmt.Errorf("bulk.IndexOp: %w", err)
		}
	}

	if _, err := r.buf.Write([]byte("\n")); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.IndexOp: %w", err)
	}

	return nil
}

// UpdateOp is a helper function to add an UpdateOperation with and UpdateAction to the current bulk request.
// update is optional, if both doc and update.Doc are provided, update.Doc has precedence.
func (r *Bulk) UpdateOp(op types.UpdateOperation, doc interface{}, update *types.UpdateAction) error {
	operation := types.OperationContainer{Update: &op}
	header, err := json.Marshal(operation)
	if err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.UpdateOp: %w", err)
	}

	if _, err := r.buf.Write(header); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.UpdateOp: %w", err)
	}
	if _, err := r.buf.Write([]byte("\n")); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.UpdateOp: %w", err)
	}

	if update == nil {
		update = types.NewUpdateAction()
	}

	if len(update.Doc) == 0 {
		switch v := doc.(type) {
		case []byte:
			if json.Valid(v) {
				update.Doc = v
			} else {
				r.buf.Reset()
				return fmt.Errorf("bulk.UpdateOp: invalid json")
			}
		case json.RawMessage:
			if json.Valid(v) {
				update.Doc = v
			} else {
				r.buf.Reset()
				return fmt.Errorf("bulk.UpdateOp: invalid json")
			}
		default:
			//doc can be nil if passed in script
			if doc == nil {
				break
			}
			body, err := json.Marshal(doc)
			if err != nil {
				r.buf.Reset()
				return fmt.Errorf("bulk.UpdateOp: %w", err)
			}
			update.Doc = body
		}
	}

	body, err := json.Marshal(update)
	if err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.UpdateOp: %w", err)
	}
	if _, err := r.buf.Write(body); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.UpdateOp: %w", err)
	}

	if _, err := r.buf.Write([]byte("\n")); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.UpdateOp: %w", err)
	}

	return nil
}

// DeleteOp is a helper function to add a DeleteOperation to the current bulk request.
func (r *Bulk) DeleteOp(op types.DeleteOperation) error {
	operation := types.OperationContainer{Delete: &op}
	header, err := json.Marshal(operation)
	if err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.DeleteOp: %w", err)
	}

	if _, err := r.buf.Write(header); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.DeleteOp: %w", err)
	}
	if _, err := r.buf.Write([]byte("\n")); err != nil {
		r.buf.Reset()
		return fmt.Errorf("bulk.DeleteOp: %w", err)
	}

	return nil
}

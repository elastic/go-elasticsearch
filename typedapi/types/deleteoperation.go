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

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// DeleteOperation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/bulk/types.ts#L134-L134
type DeleteOperation struct {
	// Id_ The document ID.
	Id_           *string `json:"_id,omitempty"`
	IfPrimaryTerm *int64  `json:"if_primary_term,omitempty"`
	IfSeqNo       *int64  `json:"if_seq_no,omitempty"`
	// Index_ Name of the index or index alias to perform the action on.
	Index_ *string `json:"_index,omitempty"`
	// Routing Custom value used to route operations to a specific shard.
	Routing     *string                  `json:"routing,omitempty"`
	Version     *int64                   `json:"version,omitempty"`
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *DeleteOperation) UnmarshalJSON(data []byte) error {

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

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return err
			}

		case "if_primary_term":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IfPrimaryTerm = &value
			case float64:
				f := int64(v)
				s.IfPrimaryTerm = &f
			}

		case "if_seq_no":
			if err := dec.Decode(&s.IfSeqNo); err != nil {
				return err
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return err
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDeleteOperation returns a DeleteOperation.
func NewDeleteOperation() *DeleteOperation {
	r := &DeleteOperation{}

	return r
}

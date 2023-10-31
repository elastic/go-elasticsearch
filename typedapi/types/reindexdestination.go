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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// ReindexDestination type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/reindex/types.ts#L39-L64
type ReindexDestination struct {
	// Index The name of the data stream, index, or index alias you are copying to.
	Index string `json:"index"`
	// OpType Set to `create` to only index documents that do not already exist.
	// Important: To reindex to a data stream destination, this argument must be
	// `create`.
	OpType *optype.OpType `json:"op_type,omitempty"`
	// Pipeline The name of the pipeline to use.
	Pipeline *string `json:"pipeline,omitempty"`
	// Routing By default, a document's routing is preserved unless it’s changed by the
	// script.
	// Set to `discard` to set routing to `null`,  or `=value` to route using the
	// specified `value`.
	Routing *string `json:"routing,omitempty"`
	// VersionType The versioning to use for the indexing operation.
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *ReindexDestination) UnmarshalJSON(data []byte) error {

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

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "op_type":
			if err := dec.Decode(&s.OpType); err != nil {
				return err
			}

		case "pipeline":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pipeline = &o

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
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

// NewReindexDestination returns a ReindexDestination.
func NewReindexDestination() *ReindexDestination {
	r := &ReindexDestination{}

	return r
}

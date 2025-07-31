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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// ReindexDestination type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/reindex/types.ts#L39-L67
type ReindexDestination struct {
	// Index The name of the data stream, index, or index alias you are copying to.
	Index string `json:"index"`
	// OpType If it is `create`, the operation will only index documents that do not
	// already exist (also known as "put if absent").
	//
	// IMPORTANT: To reindex to a data stream destination, this argument must be
	// `create`.
	OpType *optype.OpType `json:"op_type,omitempty"`
	// Pipeline The name of the pipeline to use.
	Pipeline *string `json:"pipeline,omitempty"`
	// Routing By default, a document's routing is preserved unless it's changed by the
	// script.
	// If it is `keep`, the routing on the bulk request sent for each match is set
	// to the routing on the match.
	// If it is `discard`, the routing on the bulk request sent for each match is
	// set to `null`.
	// If it is `=value`, the routing on the bulk request sent for each match is set
	// to all value specified after the equals sign (`=`).
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
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "op_type":
			if err := dec.Decode(&s.OpType); err != nil {
				return fmt.Errorf("%s | %w", "OpType", err)
			}

		case "pipeline":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Pipeline", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pipeline = &o

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return fmt.Errorf("%s | %w", "VersionType", err)
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

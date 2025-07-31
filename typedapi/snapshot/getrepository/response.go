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

package getrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package getrepository
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/get_repository/SnapshotGetRepositoryResponse.ts#L23-L25

type Response map[string]types.Repository

// NewResponse returns a Response
func NewResponse() Response {
	r := make(Response, 0)
	return r
}

func (r Response) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))
	o := make(map[string]any, 0)
	dec.Decode(&o)
	dec = json.NewDecoder(bytes.NewReader(data))
	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		key := fmt.Sprintf("%s", t)
		if target, ok := o[key]; ok {
			if t, ok := target.(map[string]any)["type"]; ok {

				switch t {

				case "azure":
					oo := types.NewAzureRepository()
					err := dec.Decode(&oo)
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return err
					}
					r[key] = oo

				case "gcs":
					oo := types.NewGcsRepository()
					err := dec.Decode(&oo)
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return err
					}
					r[key] = oo

				case "s3":
					oo := types.NewS3Repository()
					err := dec.Decode(&oo)
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return err
					}
					r[key] = oo

				case "fs":
					oo := types.NewSharedFileSystemRepository()
					err := dec.Decode(&oo)
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return err
					}
					r[key] = oo

				case "url":
					oo := types.NewReadOnlyUrlRepository()
					err := dec.Decode(&oo)
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return err
					}
					r[key] = oo

				case "source":
					oo := types.NewSourceOnlyRepository()
					err := dec.Decode(&oo)
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						return err
					}
					r[key] = oo

				}
			}
		}

	}
	return nil
}

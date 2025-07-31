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
)

// BlobDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/repository_analyze/SnapshotAnalyzeRepositoryResponse.ts#L250-L284
type BlobDetails struct {
	// Name The name of the blob.
	Name string `json:"name"`
	// Overwritten Indicates whether the blob was overwritten while the read operations were
	// ongoing.
	//
	//	/**
	Overwritten bool `json:"overwritten"`
	ReadEarly   bool `json:"read_early"`
	// ReadEnd The position, in bytes, at which read operations completed.
	ReadEnd int64 `json:"read_end"`
	// ReadStart The position, in bytes, at which read operations started.
	ReadStart int64 `json:"read_start"`
	// Reads A description of every read operation performed on the blob.
	Reads ReadBlobDetails `json:"reads"`
	// Size The size of the blob.
	Size ByteSize `json:"size"`
	// SizeBytes The size of the blob in bytes.
	SizeBytes int64 `json:"size_bytes"`
}

func (s *BlobDetails) UnmarshalJSON(data []byte) error {

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

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = o

		case "overwritten":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Overwritten", err)
				}
				s.Overwritten = value
			case bool:
				s.Overwritten = v
			}

		case "read_early":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReadEarly", err)
				}
				s.ReadEarly = value
			case bool:
				s.ReadEarly = v
			}

		case "read_end":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReadEnd", err)
				}
				s.ReadEnd = value
			case float64:
				f := int64(v)
				s.ReadEnd = f
			}

		case "read_start":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReadStart", err)
				}
				s.ReadStart = value
			case float64:
				f := int64(v)
				s.ReadStart = f
			}

		case "reads":
			if err := dec.Decode(&s.Reads); err != nil {
				return fmt.Errorf("%s | %w", "Reads", err)
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return fmt.Errorf("%s | %w", "Size", err)
			}

		case "size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SizeBytes", err)
				}
				s.SizeBytes = value
			case float64:
				f := int64(v)
				s.SizeBytes = f
			}

		}
	}
	return nil
}

// NewBlobDetails returns a BlobDetails.
func NewBlobDetails() *BlobDetails {
	r := &BlobDetails{}

	return r
}

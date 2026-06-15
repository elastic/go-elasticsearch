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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// FileContent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/inference/_types/CommonTypes.ts#L205-L214
type FileContent struct {
	// FileData The base64 encoded file data
	FileData string `json:"file_data"`
	// Filename The name of the file
	Filename string `json:"filename"`
}

func (s *FileContent) UnmarshalJSON(data []byte) error {

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

		case "file_data":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FileData", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FileData = o

		case "filename":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Filename", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Filename = o

		}
	}
	return nil
}

// NewFileContent returns a FileContent.
func NewFileContent() *FileContent {
	r := &FileContent{}

	return r
}

type FileContentVariant interface {
	FileContentCaster() *FileContent
}

func (s *FileContent) FileContentCaster() *FileContent {
	return s
}

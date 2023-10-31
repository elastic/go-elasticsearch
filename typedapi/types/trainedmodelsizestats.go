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
)

// TrainedModelSizeStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/TrainedModel.ts#L126-L131
type TrainedModelSizeStats struct {
	// ModelSizeBytes The size of the model in bytes.
	ModelSizeBytes ByteSize `json:"model_size_bytes"`
	// RequiredNativeMemoryBytes The amount of memory required to load the model in bytes.
	RequiredNativeMemoryBytes int `json:"required_native_memory_bytes"`
}

func (s *TrainedModelSizeStats) UnmarshalJSON(data []byte) error {

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

		case "model_size_bytes":
			if err := dec.Decode(&s.ModelSizeBytes); err != nil {
				return err
			}

		case "required_native_memory_bytes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.RequiredNativeMemoryBytes = value
			case float64:
				f := int(v)
				s.RequiredNativeMemoryBytes = f
			}

		}
	}
	return nil
}

// NewTrainedModelSizeStats returns a TrainedModelSizeStats.
func NewTrainedModelSizeStats() *TrainedModelSizeStats {
	r := &TrainedModelSizeStats{}

	return r
}

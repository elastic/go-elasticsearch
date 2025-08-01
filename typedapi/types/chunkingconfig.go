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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/chunkingmode"
)

// ChunkingConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/Datafeed.ts#L251-L264
type ChunkingConfig struct {
	// Mode If the mode is `auto`, the chunk size is dynamically calculated;
	// this is the recommended value when the datafeed does not use aggregations.
	// If the mode is `manual`, chunking is applied according to the specified
	// `time_span`;
	// use this mode when the datafeed uses aggregations. If the mode is `off`, no
	// chunking is applied.
	Mode chunkingmode.ChunkingMode `json:"mode"`
	// TimeSpan The time span that each search will be querying. This setting is applicable
	// only when the `mode` is set to `manual`.
	TimeSpan Duration `json:"time_span,omitempty"`
}

func (s *ChunkingConfig) UnmarshalJSON(data []byte) error {

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

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}

		case "time_span":
			if err := dec.Decode(&s.TimeSpan); err != nil {
				return fmt.Errorf("%s | %w", "TimeSpan", err)
			}

		}
	}
	return nil
}

// NewChunkingConfig returns a ChunkingConfig.
func NewChunkingConfig() *ChunkingConfig {
	r := &ChunkingConfig{}

	return r
}

type ChunkingConfigVariant interface {
	ChunkingConfigCaster() *ChunkingConfig
}

func (s *ChunkingConfig) ChunkingConfigCaster() *ChunkingConfig {
	return s
}

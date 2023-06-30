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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// AnalysisLimits type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Analysis.ts#L104-L115
type AnalysisLimits struct {
	// CategorizationExamplesLimit The maximum number of examples stored per category in memory and in the
	// results data store. If you increase this value, more examples are available,
	// however it requires that you have more storage available. If you set this
	// value to 0, no examples are stored. NOTE: The `categorization_examples_limit`
	// applies only to analysis that uses categorization.
	CategorizationExamplesLimit *int64 `json:"categorization_examples_limit,omitempty"`
	// ModelMemoryLimit The approximate maximum amount of memory resources that are required for
	// analytical processing. Once this limit is approached, data pruning becomes
	// more aggressive. Upon exceeding this limit, new entities are not modeled. If
	// the `xpack.ml.max_model_memory_limit` setting has a value greater than 0 and
	// less than 1024mb, that value is used instead of the default. The default
	// value is relatively small to ensure that high resource usage is a conscious
	// decision. If you have jobs that are expected to analyze high cardinality
	// fields, you will likely need to use a higher value. If you specify a number
	// instead of a string, the units are assumed to be MiB. Specifying a string is
	// recommended for clarity. If you specify a byte size unit of `b` or `kb` and
	// the number does not equate to a discrete number of megabytes, it is rounded
	// down to the closest MiB. The minimum valid value is 1 MiB. If you specify a
	// value less than 1 MiB, an error occurs. If you specify a value for the
	// `xpack.ml.max_model_memory_limit` setting, an error occurs when you try to
	// create jobs that have `model_memory_limit` values greater than that setting
	// value.
	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`
}

func (s *AnalysisLimits) UnmarshalJSON(data []byte) error {

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

		case "categorization_examples_limit":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CategorizationExamplesLimit = &value
			case float64:
				f := int64(v)
				s.CategorizationExamplesLimit = &f
			}

		case "model_memory_limit":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelMemoryLimit = &o

		}
	}
	return nil
}

// NewAnalysisLimits returns a AnalysisLimits.
func NewAnalysisLimits() *AnalysisLimits {
	r := &AnalysisLimits{}

	return r
}

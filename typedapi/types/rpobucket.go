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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// RpoBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/data_recovery/get_rpo_series/types.ts#L23-L33
type RpoBucket struct {
	// End The exclusive end of the bucket. The last bucket is clipped to the end of the
	// range.
	End DateTime `json:"end"`
	// MaxRecoveryPointAgeMillis The greatest age, in milliseconds, that the newest recoverable state reached
	// during the bucket. It is `null` if no retained recovery point covers the
	// start of the bucket or the bucket starts in the future.
	MaxRecoveryPointAgeMillis *int64 `json:"max_recovery_point_age_millis,omitempty"`
	// Start The inclusive start of the bucket.
	Start DateTime `json:"start"`
}

func (s *RpoBucket) UnmarshalJSON(data []byte) error {

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

		case "end":
			if err := dec.Decode(&s.End); err != nil {
				return fmt.Errorf("%s | %w", "End", err)
			}

		case "max_recovery_point_age_millis":
			if err := dec.Decode(&s.MaxRecoveryPointAgeMillis); err != nil {
				return fmt.Errorf("%s | %w", "MaxRecoveryPointAgeMillis", err)
			}

		case "start":
			if err := dec.Decode(&s.Start); err != nil {
				return fmt.Errorf("%s | %w", "Start", err)
			}

		}
	}
	return nil
}

// NewRpoBucket returns a RpoBucket.
func NewRpoBucket() *RpoBucket {
	r := &RpoBucket{}

	return r
}

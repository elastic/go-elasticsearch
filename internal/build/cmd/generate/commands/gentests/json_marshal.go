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

package gentests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
)

// marshalForYAMLTestJSON marshals a value into JSON with deterministic map key ordering
// and float formatting that preserves "float-ness" even for whole numbers (eg 100.0),
// which matters for some APIs (eg painless integer vs float division).
//
// It intentionally differs from encoding/json in two ways:
// - map keys are sorted for stable output
// - float64 values that are whole numbers are rendered with a trailing ".0"
func marshalForYAMLTestJSON(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	if err := writeYAMLTestJSON(&b, v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func writeYAMLTestJSON(b *bytes.Buffer, v interface{}) error {
	switch x := v.(type) {
	case nil:
		b.WriteString("null")
		return nil
	case bool:
		if x {
			b.WriteString("true")
		} else {
			b.WriteString("false")
		}
		return nil
	case string:
		// Use encoding/json for correct escaping.
		enc, err := json.Marshal(x)
		if err != nil {
			return err
		}
		b.Write(enc)
		return nil
	case int:
		b.WriteString(strconv.Itoa(x))
		return nil
	case int64:
		b.WriteString(strconv.FormatInt(x, 10))
		return nil
	case float64:
		if math.IsNaN(x) || math.IsInf(x, 0) {
			return fmt.Errorf("unsupported float value: %v", x)
		}
		// Preserve float-ness: if it's an integer value, keep one decimal place.
		if math.Trunc(x) == x {
			b.WriteString(strconv.FormatFloat(x, 'f', 1, 64))
			return nil
		}
		b.WriteString(strconv.FormatFloat(x, 'f', -1, 64))
		return nil
	case []interface{}:
		b.WriteByte('[')
		for i, vv := range x {
			if i > 0 {
				b.WriteByte(',')
			}
			if err := writeYAMLTestJSON(b, vv); err != nil {
				return err
			}
		}
		b.WriteByte(']')
		return nil
	case map[string]interface{}:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		b.WriteByte('{')
		for i, k := range keys {
			if i > 0 {
				b.WriteByte(',')
			}
			kk, _ := json.Marshal(k)
			b.Write(kk)
			b.WriteByte(':')
			if err := writeYAMLTestJSON(b, x[k]); err != nil {
				return err
			}
		}
		b.WriteByte('}')
		return nil
	default:
		// Fall back to encoding/json for any remaining primitive-ish types that
		// might leak through (eg uint, json.RawMessage).
		enc, err := json.Marshal(x)
		if err != nil {
			return err
		}
		b.Write(enc)
		return nil
	}
}

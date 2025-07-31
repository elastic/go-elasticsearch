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
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fingerprintdigest"
)

// FingerprintProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ingest/_types/Processors.ts#L905-L933
type FingerprintProcessor struct {
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// Fields Array of fields to include in the fingerprint. For objects, the processor
	// hashes both the field key and value. For other fields, the processor hashes
	// only the field value.
	Fields []string `json:"fields"`
	// If Conditionally execute the processor.
	If *Script `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If true, the processor ignores any missing fields. If all fields are
	// missing, the processor silently exits without modifying the document.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// Method The hash method used to compute the fingerprint. Must be one of MD5, SHA-1,
	// SHA-256, SHA-512, or MurmurHash3.
	Method *fingerprintdigest.FingerprintDigest `json:"method,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Salt Salt value for the hash function.
	Salt *string `json:"salt,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TargetField Output field for the fingerprint.
	TargetField *string `json:"target_field,omitempty"`
}

func (s *FingerprintProcessor) UnmarshalJSON(data []byte) error {

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

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}

				s.Fields = append(s.Fields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Fields); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}
			}

		case "if":
			if err := dec.Decode(&s.If); err != nil {
				return fmt.Errorf("%s | %w", "If", err)
			}

		case "ignore_failure":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreFailure", err)
				}
				s.IgnoreFailure = &value
			case bool:
				s.IgnoreFailure = &v
			}

		case "ignore_missing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreMissing", err)
				}
				s.IgnoreMissing = &value
			case bool:
				s.IgnoreMissing = &v
			}

		case "method":
			if err := dec.Decode(&s.Method); err != nil {
				return fmt.Errorf("%s | %w", "Method", err)
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
			}

		case "salt":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Salt", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Salt = &o

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tag", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tag = &o

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return fmt.Errorf("%s | %w", "TargetField", err)
			}

		}
	}
	return nil
}

// NewFingerprintProcessor returns a FingerprintProcessor.
func NewFingerprintProcessor() *FingerprintProcessor {
	r := &FingerprintProcessor{}

	return r
}

type FingerprintProcessorVariant interface {
	FingerprintProcessorCaster() *FingerprintProcessor
}

func (s *FingerprintProcessor) FingerprintProcessorCaster() *FingerprintProcessor {
	return s
}

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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/translogdurability"
)

// Translog type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L368-L390
type Translog struct {
	// Durability Whether or not to `fsync` and commit the translog after every index, delete,
	// update, or bulk request.
	Durability *translogdurability.TranslogDurability `json:"durability,omitempty"`
	// FlushThresholdSize The translog stores all operations that are not yet safely persisted in
	// Lucene (i.e., are not
	// part of a Lucene commit point). Although these operations are available for
	// reads, they will need
	// to be replayed if the shard was stopped and had to be recovered. This setting
	// controls the
	// maximum total size of these operations, to prevent recoveries from taking too
	// long. Once the
	// maximum size has been reached a flush will happen, generating a new Lucene
	// commit point.
	FlushThresholdSize ByteSize           `json:"flush_threshold_size,omitempty"`
	Retention          *TranslogRetention `json:"retention,omitempty"`
	// SyncInterval How often the translog is fsynced to disk and committed, regardless of write
	// operations.
	// Values less than 100ms are not allowed.
	SyncInterval Duration `json:"sync_interval,omitempty"`
}

func (s *Translog) UnmarshalJSON(data []byte) error {

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

		case "durability":
			if err := dec.Decode(&s.Durability); err != nil {
				return fmt.Errorf("%s | %w", "Durability", err)
			}

		case "flush_threshold_size":
			if err := dec.Decode(&s.FlushThresholdSize); err != nil {
				return fmt.Errorf("%s | %w", "FlushThresholdSize", err)
			}

		case "retention":
			if err := dec.Decode(&s.Retention); err != nil {
				return fmt.Errorf("%s | %w", "Retention", err)
			}

		case "sync_interval":
			if err := dec.Decode(&s.SyncInterval); err != nil {
				return fmt.Errorf("%s | %w", "SyncInterval", err)
			}

		}
	}
	return nil
}

// NewTranslog returns a Translog.
func NewTranslog() *Translog {
	r := &Translog{}

	return r
}

type TranslogVariant interface {
	TranslogCaster() *Translog
}

func (s *Translog) TranslogCaster() *Translog {
	return s
}

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

// TransformSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/transform/get_transform/types.ts#L33-L62
type TransformSummary struct {
	// Authorization The security privileges that the transform uses to run its queries. If
	// Elastic Stack security features were disabled at the time of the most recent
	// update to the transform, this property is omitted.
	Authorization *TransformAuthorization `json:"authorization,omitempty"`
	// CreateTime The time the transform was created.
	CreateTime       *int64   `json:"create_time,omitempty"`
	CreateTimeString DateTime `json:"create_time_string,omitempty"`
	// Description Free text description of the transform.
	Description *string `json:"description,omitempty"`
	// Dest The destination for the transform.
	Dest      ReindexDestination `json:"dest"`
	Frequency Duration           `json:"frequency,omitempty"`
	Id        string             `json:"id"`
	Latest    *Latest            `json:"latest,omitempty"`
	Meta_     Metadata           `json:"_meta,omitempty"`
	// Pivot The pivot method transforms the data by aggregating and grouping it.
	Pivot           *Pivot                    `json:"pivot,omitempty"`
	RetentionPolicy *RetentionPolicyContainer `json:"retention_policy,omitempty"`
	// Settings Defines optional transform settings.
	Settings *Settings `json:"settings,omitempty"`
	// Source The source of the data for the transform.
	Source TransformSource `json:"source"`
	// Sync Defines the properties transforms require to run continuously.
	Sync *SyncContainer `json:"sync,omitempty"`
	// Version The version of Elasticsearch that existed on the node when the transform was
	// created.
	Version *string `json:"version,omitempty"`
}

func (s *TransformSummary) UnmarshalJSON(data []byte) error {

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

		case "authorization":
			if err := dec.Decode(&s.Authorization); err != nil {
				return fmt.Errorf("%s | %w", "Authorization", err)
			}

		case "create_time":
			if err := dec.Decode(&s.CreateTime); err != nil {
				return fmt.Errorf("%s | %w", "CreateTime", err)
			}

		case "create_time_string":
			if err := dec.Decode(&s.CreateTimeString); err != nil {
				return fmt.Errorf("%s | %w", "CreateTimeString", err)
			}

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

		case "dest":
			if err := dec.Decode(&s.Dest); err != nil {
				return fmt.Errorf("%s | %w", "Dest", err)
			}

		case "frequency":
			if err := dec.Decode(&s.Frequency); err != nil {
				return fmt.Errorf("%s | %w", "Frequency", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "latest":
			if err := dec.Decode(&s.Latest); err != nil {
				return fmt.Errorf("%s | %w", "Latest", err)
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "pivot":
			if err := dec.Decode(&s.Pivot); err != nil {
				return fmt.Errorf("%s | %w", "Pivot", err)
			}

		case "retention_policy":
			if err := dec.Decode(&s.RetentionPolicy); err != nil {
				return fmt.Errorf("%s | %w", "RetentionPolicy", err)
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "source":
			if err := dec.Decode(&s.Source); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}

		case "sync":
			if err := dec.Decode(&s.Sync); err != nil {
				return fmt.Errorf("%s | %w", "Sync", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewTransformSummary returns a TransformSummary.
func NewTransformSummary() *TransformSummary {
	r := &TransformSummary{}

	return r
}

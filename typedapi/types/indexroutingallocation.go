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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexroutingallocationoptions"
)

// IndexRoutingAllocation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/indices/_types/IndexRouting.ts#L28-L39
type IndexRoutingAllocation struct {
	Disk            *IndexRoutingAllocationDisk                                  `json:"disk,omitempty"`
	Enable          *indexroutingallocationoptions.IndexRoutingAllocationOptions `json:"enable,omitempty"`
	Include         *IndexRoutingAllocationInclude                               `json:"include,omitempty"`
	InitialRecovery *IndexRoutingAllocationInitialRecovery                       `json:"initial_recovery,omitempty"`
	// Require Assign the index to a node whose attribute has all of the given values. Each
	// key is a node attribute: one of the built-in attributes `_name`, `_host_ip`,
	// `_publish_ip`, `_ip`, `_host`, `_id` and `_tier`, or a custom node attribute.
	// Each value is a comma-separated string or an array of strings, and can use
	// wildcards.
	Require IndexRoutingAllocationRequire `json:"require,omitempty"`
}

func (s *IndexRoutingAllocation) UnmarshalJSON(data []byte) error {

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

		case "disk":
			if err := dec.Decode(&s.Disk); err != nil {
				return fmt.Errorf("%s | %w", "Disk", err)
			}

		case "enable":
			if err := dec.Decode(&s.Enable); err != nil {
				return fmt.Errorf("%s | %w", "Enable", err)
			}

		case "include":
			if err := dec.Decode(&s.Include); err != nil {
				return fmt.Errorf("%s | %w", "Include", err)
			}

		case "initial_recovery":
			if err := dec.Decode(&s.InitialRecovery); err != nil {
				return fmt.Errorf("%s | %w", "InitialRecovery", err)
			}

		case "require":
			if err := dec.Decode(&s.Require); err != nil {
				return fmt.Errorf("%s | %w", "Require", err)
			}

		}
	}
	return nil
}

// NewIndexRoutingAllocation returns a IndexRoutingAllocation.
func NewIndexRoutingAllocation() *IndexRoutingAllocation {
	r := &IndexRoutingAllocation{}

	return r
}

type IndexRoutingAllocationVariant interface {
	IndexRoutingAllocationCaster() *IndexRoutingAllocation
}

func (s *IndexRoutingAllocation) IndexRoutingAllocationCaster() *IndexRoutingAllocation {
	return s
}

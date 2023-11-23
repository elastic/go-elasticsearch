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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package puttransform

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package puttransform
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/transform/put_transform/PutTransformRequest.ts#L33-L122
type Request struct {

	// Description Free text description of the transform.
	Description *string `json:"description,omitempty"`
	// Dest The destination for the transform.
	Dest types.TransformDestination `json:"dest"`
	// Frequency The interval between checks for changes in the source indices when the
	// transform is running continuously. Also
	// determines the retry interval in the event of transient failures while the
	// transform is searching or indexing.
	// The minimum value is `1s` and the maximum is `1h`.
	Frequency types.Duration `json:"frequency,omitempty"`
	// Latest The latest method transforms the data by finding the latest document for each
	// unique key.
	Latest *types.Latest `json:"latest,omitempty"`
	// Meta_ Defines optional transform metadata.
	Meta_ types.Metadata `json:"_meta,omitempty"`
	// Pivot The pivot method transforms the data by aggregating and grouping it. These
	// objects define the group by fields
	// and the aggregation to reduce the data.
	Pivot *types.Pivot `json:"pivot,omitempty"`
	// RetentionPolicy Defines a retention policy for the transform. Data that meets the defined
	// criteria is deleted from the
	// destination index.
	RetentionPolicy *types.RetentionPolicyContainer `json:"retention_policy,omitempty"`
	// Settings Defines optional transform settings.
	Settings *types.Settings `json:"settings,omitempty"`
	// Source The source of the data for the transform.
	Source types.TransformSource `json:"source"`
	// Sync Defines the properties transforms require to run continuously.
	Sync *types.SyncContainer `json:"sync,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Puttransform request: %w", err)
	}

	return &req, nil
}

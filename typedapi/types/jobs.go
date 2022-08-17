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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"encoding/json"
)

// Jobs type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L356-L358
type Jobs struct {
	All_ *AllJobs       `json:"_all,omitempty"`
	Jobs map[string]Job `json:"-"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s Jobs) MarshalJSON() ([]byte, error) {
	type opt Jobs
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Jobs {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// JobsBuilder holds Jobs struct and provides a builder API.
type JobsBuilder struct {
	v *Jobs
}

// NewJobs provides a builder for the Jobs struct.
func NewJobsBuilder() *JobsBuilder {
	r := JobsBuilder{
		&Jobs{
			Jobs: make(map[string]Job, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Jobs struct
func (rb *JobsBuilder) Build() Jobs {
	return *rb.v
}

func (rb *JobsBuilder) All_(all_ *AllJobsBuilder) *JobsBuilder {
	v := all_.Build()
	rb.v.All_ = &v
	return rb
}

func (rb *JobsBuilder) Jobs(values map[string]*JobBuilder) *JobsBuilder {
	tmp := make(map[string]Job, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Jobs = tmp
	return rb
}

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

// IoStatDevice type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L292-L299
type IoStatDevice struct {
	DeviceName      *string `json:"device_name,omitempty"`
	Operations      *int64  `json:"operations,omitempty"`
	ReadKilobytes   *int64  `json:"read_kilobytes,omitempty"`
	ReadOperations  *int64  `json:"read_operations,omitempty"`
	WriteKilobytes  *int64  `json:"write_kilobytes,omitempty"`
	WriteOperations *int64  `json:"write_operations,omitempty"`
}

// IoStatDeviceBuilder holds IoStatDevice struct and provides a builder API.
type IoStatDeviceBuilder struct {
	v *IoStatDevice
}

// NewIoStatDevice provides a builder for the IoStatDevice struct.
func NewIoStatDeviceBuilder() *IoStatDeviceBuilder {
	r := IoStatDeviceBuilder{
		&IoStatDevice{},
	}

	return &r
}

// Build finalize the chain and returns the IoStatDevice struct
func (rb *IoStatDeviceBuilder) Build() IoStatDevice {
	return *rb.v
}

func (rb *IoStatDeviceBuilder) DeviceName(devicename string) *IoStatDeviceBuilder {
	rb.v.DeviceName = &devicename
	return rb
}

func (rb *IoStatDeviceBuilder) Operations(operations int64) *IoStatDeviceBuilder {
	rb.v.Operations = &operations
	return rb
}

func (rb *IoStatDeviceBuilder) ReadKilobytes(readkilobytes int64) *IoStatDeviceBuilder {
	rb.v.ReadKilobytes = &readkilobytes
	return rb
}

func (rb *IoStatDeviceBuilder) ReadOperations(readoperations int64) *IoStatDeviceBuilder {
	rb.v.ReadOperations = &readoperations
	return rb
}

func (rb *IoStatDeviceBuilder) WriteKilobytes(writekilobytes int64) *IoStatDeviceBuilder {
	rb.v.WriteKilobytes = &writekilobytes
	return rb
}

func (rb *IoStatDeviceBuilder) WriteOperations(writeoperations int64) *IoStatDeviceBuilder {
	rb.v.WriteOperations = &writeoperations
	return rb
}

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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

// FieldValue holds the union for the following types:
//
//	bool
//	float64
//	int64
//	nil
//	string
//	interface{}
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/common.ts#L25-L37
type FieldValue interface{}

// FieldValueBuilder holds FieldValue struct and provides a builder API.
type FieldValueBuilder struct {
	v FieldValue
}

// NewFieldValue provides a builder for the FieldValue struct.
func NewFieldValueBuilder() *FieldValueBuilder {
	return &FieldValueBuilder{}
}

// Build finalize the chain and returns the FieldValue struct
func (u *FieldValueBuilder) Build() FieldValue {
	return u.v
}

func (u *FieldValueBuilder) Bool(bool bool) *FieldValueBuilder {
	u.v = &bool
	return u
}

func (u *FieldValueBuilder) Float64(float64 float64) *FieldValueBuilder {
	u.v = &float64
	return u
}

func (u *FieldValueBuilder) Int64(int64 int64) *FieldValueBuilder {
	u.v = &int64
	return u
}

func (u *FieldValueBuilder) Nil() *FieldValueBuilder {
	u.v = nil
	return u
}

func (u *FieldValueBuilder) String(string string) *FieldValueBuilder {
	u.v = &string
	return u
}

func (u *FieldValueBuilder) UserDefinedValue(userdefinedvalue interface{}) *FieldValueBuilder {
	u.v = userdefinedvalue
	return u
}

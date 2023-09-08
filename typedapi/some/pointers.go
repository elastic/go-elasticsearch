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

// Package some provides helpers to allow users to user inline pointers
// on primitive types for the TypedAPI.
package some

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// String returns a pointer to a string
func String(value string) *string {
	return &value
}

// Bool returns a pointer to a boolean
func Bool(value bool) *bool {
	return &value
}

// Int returns a pointer to an integer
func Int(value int) *int {
	return &value
}

// Int8 returns a pointer to an int8
func Int8(value int8) *int8 {
	return &value
}

// Int16 returns a pointer to an int16
func Int16(value int16) *int16 {
	return &value
}

// Int32 returns a pointer to an int32
func Int32(value int32) *int32 {
	return &value
}

// Int64 returns a pointer to an int64
func Int64(value int64) *int64 {
	return &value
}

// Uint returns a pointer to a uint
func Uint(value uint) *uint {
	return &value
}

// Uint8 returns a pointer to a uint8
func Uint8(value uint8) *uint8 {
	return &value
}

// Uint16 returns a pointer to a uint16
func Uint16(value uint16) *uint16 {
	return &value
}

// Uint32 returns a pointer to a uint32
func Uint32(value uint32) *uint32 {
	return &value
}

// Uint64 returns a pointer to a uint64
func Uint64(value uint64) *uint64 {
	return &value
}

// Byte returns a pointer to a byte
func Byte(value byte) *byte {
	return &value
}

// Rune returns a pointer to a rune
func Rune(value rune) *rune {
	return &value
}

// Float32 returns a pointer to a float32
func Float32(value float32) *float32 {
	return &value
}

// Float64 returns a pointer to a float64
func Float64(value float64) *types.Float64 {
	f := types.Float64(value)
	return &f
}

// Complex64 returns a pointer to a complex64
func Complex64(value complex64) *complex64 {
	return &value
}

// Complex128 returns a pointer to a complex128
func Complex128(value complex128) *complex128 {
	return &value
}

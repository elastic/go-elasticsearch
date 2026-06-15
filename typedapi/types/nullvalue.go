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

package types

// NullValue is a custom type used to represent the concept of a null or missing value.
// It can be used as a placeholder for variables or fields that are not initialized,
// or to indicate that a specific piece of data is intentionally absent.
type NullValue struct{}

// MarshalJSON converts the NullValue to JSON format.
// It always returns a "null" value as per JSON standard for null values.
func (n NullValue) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

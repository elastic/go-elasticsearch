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

package utils

// MapKeys returns the map keys as a slice of strings.
func MapKeys(s interface{}) []string {
	if s, ok := s.(map[interface{}]interface{}); ok {
		keys := make([]string, 0, len(s))
		for k := range s {
			if k, ok := k.(string); ok {
				keys = append(keys, k)
			}
		}
		return keys
	}
	keys := make([]string, 0)
	return keys
}

// MapValues returns the map values as a slice of interfaces.
func MapValues(s interface{}) []interface{} {
	if s, ok := s.(map[interface{}]interface{}); ok {
		values := make([]interface{}, 0, len(s))
		for _, v := range s {
			values = append(values, v)
		}
		return values
	}
	values := make([]interface{}, 0)
	return values
}

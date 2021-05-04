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
//
func MapKeys(s interface{}) (keys []string) {
	if s, ok := s.(map[interface{}]interface{}); ok {
		for k := range s {
			if k, ok := k.(string); ok {
				keys = append(keys, k)
			}
		}
	}
	return keys
}

// MapValues returns the map values as a slice of interfaces.
//
func MapValues(s interface{}) (values []interface{}) {
	if s, ok := s.(map[interface{}]interface{}); ok {
		for _, v := range s {
			values = append(values, v)
		}
	}
	return values
}

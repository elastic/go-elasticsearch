// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

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

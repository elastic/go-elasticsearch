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

package gentests

import (
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// yamlNodeToInterface converts a yaml.v3 Node tree into the interface{} representation
// used by the rest of gentests.
//
// Critical behavior: mapping keys are always returned as strings using the raw scalar
// value, regardless of YAML 1.1 schema resolution. This avoids issues like unquoted
// `y:` being interpreted as boolean true by YAML 1.1.
func yamlNodeToInterface(n *yaml.Node) interface{} {
	return yamlNodeToInterfaceWithKeyMode(n, false)
}

func yamlNodeToInterfaceWithKeyMode(n *yaml.Node, isMapKey bool) interface{} {
	if n == nil {
		return nil
	}

	switch n.Kind {
	case yaml.DocumentNode:
		if len(n.Content) == 0 {
			return nil
		}
		return yamlNodeToInterfaceWithKeyMode(n.Content[0], isMapKey)

	case yaml.MappingNode:
		m := make(map[interface{}]interface{}, len(n.Content)/2)
		for i := 0; i+1 < len(n.Content); i += 2 {
			k := n.Content[i]
			v := n.Content[i+1]
			key := yamlNodeToInterfaceWithKeyMode(k, true) // force string key
			m[key] = yamlNodeToInterfaceWithKeyMode(v, false)
		}
		return m

	case yaml.SequenceNode:
		out := make([]interface{}, 0, len(n.Content))
		for _, c := range n.Content {
			out = append(out, yamlNodeToInterfaceWithKeyMode(c, false))
		}
		return out

	case yaml.ScalarNode:
		if isMapKey {
			return n.Value
		}

		switch n.Tag {
		case "!!null":
			return nil
		case "!!bool":
			// YAML tests use explicit true/false; parse conservatively.
			b, err := strconv.ParseBool(strings.TrimSpace(strings.ToLower(n.Value)))
			if err == nil {
				return b
			}
			// Fallback: treat unknown boolean-like scalars as strings.
			return n.Value
		case "!!int":
			i, err := strconv.ParseInt(strings.TrimSpace(n.Value), 0, 64)
			if err == nil {
				// Keep parity with previous yaml.v2 decode behavior (int when it fits).
				return int(i)
			}
			return n.Value
		case "!!float":
			f, err := strconv.ParseFloat(strings.TrimSpace(n.Value), 64)
			if err == nil {
				return f
			}
			return n.Value
		default:
			return n.Value
		}

	default:
		return nil
	}
}

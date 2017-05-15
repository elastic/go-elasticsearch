/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package action

import (
	"fmt"
	"text/template"

	"github.com/elastic/go-elasticsearch/generator/api"

	yaml "gopkg.in/yaml.v2"
)

type action interface {
	Resolve(methods map[string]*api.Method, templates *template.Template) error
	String() (string, error)
}

// Router allows to parse actions into the action interface.
type Router struct {
	action
}

// UnmarshalYAML instantiates a new Router by parsing its YAML spec.
func (r *Router) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var spec yaml.MapSlice
	err := unmarshal(&spec)
	if err != nil {
		return err
	}
	if len(spec) > 1 {
		return fmt.Errorf("action has multiple names: %#v", spec)
	}
	for _, item := range spec {
		*r = Router{}
		name, ok := item.Key.(string)
		if !ok {
			return fmt.Errorf("unexpected type for action name: %T (expected string)", item.Key)
		}
		switch name {
		case "skip":
			r.action, err = newSkip(unmarshal)
		case "do":
			r.action, err = newDo(unmarshal)
		case "set":
			r.action, err = newSet(unmarshal)
		case "is_true":
			fallthrough
		case "is_false":
			r.action, err = newBoolCompare(unmarshal)
		case "match":
			r.action, err = newMatch(unmarshal)
		case "lt":
			fallthrough
		case "gt":
			fallthrough
		case "lte":
			fallthrough
		case "gte":
			r.action, err = newIntCompare(unmarshal)
		case "length":
			r.action, err = newLength(unmarshal)
		default:
			return fmt.Errorf("unknown action: %s", name)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

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

package generator

import (
	"fmt"
	"text/template"
)

type boolCompare struct {
	spec     map[string]string
	key      string
	Operator string
	NilValue string
}

func newBoolCompare(unmarshal func(interface{}) error) (action, error) {
	b := &boolCompare{}
	if err := unmarshal(&b.spec); err != nil {
		return nil, err
	}
	if len(b.spec) > 1 {
		return nil, fmt.Errorf("found more than one operation in %#v", b.spec)
	}
	for name, key := range b.spec {
		switch name {
		case "is_true":
			b.Operator = "!="
		case "is_false":
			b.Operator = "=="
		default:
			return nil, fmt.Errorf("unexpected bool comparison operation: %s", name)
		}
		b.key = key
	}
	return b, nil
}

func (b *boolCompare) resolve(methods map[string]*method, templates *template.Template) error {
	// TODO: implement resolution and per-type nil value
	return nil
}

func (b *boolCompare) String() (string, error) {
	// TODO: implement
	return "", nil
}

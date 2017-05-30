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
	"bytes"
	"fmt"
	"text/template"

	"github.com/elastic/go-elasticsearch/generator/api"
)

type intCompare struct {
	spec     map[string]map[string]int
	Key      string
	Value    int
	Operator string
	template *template.Template
}

func newIntCompare(unmarshal func(interface{}) error) (action, error) {
	i := &intCompare{}
	if err := unmarshal(&i.spec); err != nil {
		return nil, err
	}
	if len(i.spec) > 1 {
		return nil, fmt.Errorf("found more than one operation in %#v", i.spec)
	}
	for name, op := range i.spec {
		switch name {
		case "lt":
			i.Operator = "<"
		case "gt":
			i.Operator = ">"
		case "lte":
			i.Operator = "<="
		case "gte":
			i.Operator = ">="
		default:
			return nil, fmt.Errorf("unexpected int comparison operation: %s", name)
		}
		if len(op) > 1 {
			return nil, fmt.Errorf("found more than one operation in %#v", i.spec)
		}
		for varName, value := range op {
			i.Key = varName
			i.Value = value
		}
	}
	return i, nil
}

func (i *intCompare) Resolve(testSpecFile string, methods map[string]*api.Method, templates *template.Template) error {
	i.template = templates.Lookup("int_compare.tmpl")
	if i.template == nil {
		return fmt.Errorf("unable to find template for int comparison")
	}
	return nil
}

func (i *intCompare) String() (string, error) {
	var writer bytes.Buffer
	if err := i.template.Execute(&writer, i); err != nil {
		return "", err
	}
	return writer.String(), nil
}

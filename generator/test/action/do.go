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

type do struct {
	testSpecFile string
	spec         map[string]map[string]map[string]interface{}
	Method       *api.Method
	template     *template.Template
}

func newDo(unmarshal func(interface{}) error) (action, error) {
	d := &do{}
	if err := unmarshal(&d.spec); err != nil {
		var spec map[string]string
		if err := unmarshal(&spec); err != nil {
			// TODO: implement catch et al here and in resolve
			return d, nil
		}
	}
	return d, nil
}

func (d *do) Resolve(testSpecFile string, methods map[string]*api.Method, templates *template.Template) error {
	spec := d.spec["do"]
	for methodName, args := range spec {
		if methodName == "catch" || methodName == "warnings" || methodName == "headers" {
			continue
		}
		m, ok := methods[methodName]
		if !ok {
			return fmt.Errorf("invalid method name %q in %#v", methodName, spec)
		}
		// TODO: implement variables in args
		methodCall, err := m.Call(args)
		if err != nil {
			return err
		}
		d.Method = methodCall
	}
	d.template = templates.Lookup("do.tmpl")
	if d.template == nil {
		return fmt.Errorf("unable to find template for do")
	}
	d.testSpecFile = testSpecFile
	return nil
}

func (d *do) String() (string, error) {
	var writer bytes.Buffer
	err := d.template.Execute(&writer, d)
	if err != nil {
		return "", fmt.Errorf("failed to render %s: %s", d.testSpecFile, err)
	}
	return writer.String(), err
}

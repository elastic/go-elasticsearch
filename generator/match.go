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
	"bytes"
	"fmt"
	"text/template"
)

type match struct {
	spec     map[string]interface{}
	template *template.Template
}

func newMatch(unmarshal func(interface{}) error) (action, error) {
	m := &match{}
	err := unmarshal(&m.spec)
	if err != nil {
		return nil, err
	}
	if len(m.spec) > 1 {
		return nil, fmt.Errorf("expected to have a single entry in match, got %#v", m.spec)
	}
	// TODO: implement regex
	return m, nil
}

func (m *match) resolve(methods map[string]*method, templates *template.Template) error {
	m.template = templates.Lookup("match.tmpl")
	if m.template == nil {
		return fmt.Errorf("unable to find template for match")
	}
	return nil
}

func (m *match) String() (string, error) {
	var writer bytes.Buffer
	if err := m.template.Execute(&writer, m.spec["match"]); err != nil {
		return "", err
	}
	return writer.String(), nil
}

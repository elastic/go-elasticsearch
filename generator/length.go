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

type length struct {
	spec     map[string]map[string]int
	Key      string
	Length   int
	template *template.Template
}

func newLength(unmarshal func(interface{}) error) (action, error) {
	l := &length{}
	if err := unmarshal(&l.spec); err != nil {
		return nil, err
	}
	if len(l.spec) > 1 {
		return nil, fmt.Errorf("found more than one operation in %#v", l.spec)
	}
	for _, item := range l.spec {
		for name, length := range item {
			l.Key = name
			l.Length = length
		}
	}
	return l, nil
}

func (l *length) resolve(methods map[string]*method, templates *template.Template) error {
	l.template = templates.Lookup("length.tmpl")
	if l.template == nil {
		return fmt.Errorf("unable to find template for length")
	}
	return nil
}

func (l *length) String() (string, error) {
	var writer bytes.Buffer
	if err := l.template.Execute(&writer, l); err != nil {
		return "", err
	}
	return writer.String(), nil
}

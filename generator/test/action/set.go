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

type set struct {
	spec     map[string]map[string]string
	Key      string
	VarName  string
	template *template.Template
}

func newSet(unmarshal func(interface{}) error) (action, error) {
	s := &set{}
	if err := unmarshal(&s.spec); err != nil {
		return nil, err
	}
	if len(s.spec) > 1 {
		return nil, fmt.Errorf("found more than one operation in %#v", s.spec)
	}
	for _, item := range s.spec {
		if len(item) > 1 {
			return nil, fmt.Errorf("found more than one field in %#v", s.spec)
		}
		for key, name := range item {
			s.Key = key
			s.VarName = name
		}
	}
	return s, nil
}

func (s *set) Resolve(methods map[string]*api.Method, templates *template.Template) error {
	s.template = templates.Lookup("set.tmpl")
	if s.template == nil {
		return fmt.Errorf("unable to find template for set")
	}
	return nil
}

func (s *set) String() (string, error) {
	var writer bytes.Buffer
	if err := s.template.Execute(&writer, s); err != nil {
		return "", err
	}
	return writer.String(), nil
}

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
	"text/template"

	"github.com/elastic/go-elasticsearch/generator/api"
)

type skip struct {
	Version  string      `yaml:"version"`
	Reason   string      `yaml:"reason"`
	Features interface{} `yaml:"features"`
}

func newSkip(unmarshal func(interface{}) error) (action, error) {
	// TODO: implement
	return &skip{}, nil
}

func (s *skip) Resolve(testSpecFile string, methods map[string]*api.Method, templates *template.Template) error {
	return nil
}

func (s *skip) String() (string, error) {
	return "", nil
}

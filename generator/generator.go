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

const (
	apiDir = "../../api"
)

// Generator is a code generator based on JSON specs and Go template
type Generator struct {
	spec map[string]interface{}
}

// New creates a new generator
func New(specFile string) (*Generator, error) {
	spec, err := unmarshalSpec(specFile)
	if err != nil {
		return nil, err
	}
	return &Generator{spec: spec}, nil
}

// Run runs the generator
func (g *Generator) Run() error {
	return executeTemplate(g.spec, apiDir)
}

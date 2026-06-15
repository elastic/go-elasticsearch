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
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

var skipTests map[string][]string

func init() {
	err := yaml.NewDecoder(strings.NewReader(skipTestsYAML)).Decode(&skipTests)
	if err != nil {
		panic(fmt.Sprintf("ERROR: %v", err))
	}
}

var skipFiles = []string{
	"entsearch/50_connector_updates.yml",
	"machine_learning/jobs_crud.yml",
	"nodes/10_basic.yml",
	"search_mvt/10_basic.yml",
}

// TODO: Comments into descriptions for `Skip()`
var skipTestsYAML = `
---
eql/10_basic.yml:
  - "eql basic tests"
security/140_user.yml:
  - "security.user"
`

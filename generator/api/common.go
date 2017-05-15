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

package api

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

const (
	// CommonParamsSpecFile is the file containing the JSON spec for the common parameters.
	CommonParamsSpecFile = "_common.json"
)

// NewCommonParams instantiates common parameters.
func NewCommonParams(specDir string, templates *template.Template) (map[string]*Param, error) {
	bytes, err := ioutil.ReadFile(filepath.Join(specDir, "api", CommonParamsSpecFile))
	if err != nil {
		return nil, err
	}
	var spec map[string]*json.RawMessage
	if err = json.Unmarshal(bytes, &spec); err != nil {
		return nil, err
	}
	var params map[string]*Param
	if err = json.Unmarshal(*spec["params"], &params); err != nil {
		return nil, err
	}
	params["ignore"] = &Param{
		SpecType:           "number_list",
		Description:        "ignores the specified HTTP status codes",
		optionTemplateName: "ignore.tmpl",
	}
	for name, p := range params {
		if err := p.resolve(name, templates); err != nil {
			return nil, err
		}
	}
	return params, nil
}

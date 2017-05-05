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

import "text/template"

type intCompare struct {
	Spec map[string]map[string]int
}

func newIntCompare(unmarshal func(interface{}) error) (action, error) {
	// TODO: implement
	return &intCompare{}, nil
}

func (i *intCompare) resolve(methods map[string]*method, templates *template.Template) error {
	// TODO: implement
	return nil
}

func (i *intCompare) String() (string, error) {
	// TODO: implement
	return "", nil
}

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

type lt struct {
	Spec map[string]*param
}

func newLt(unmarshal func(interface{}) error) (action, error) {
	// TODO: implement
	return &lt{}, nil
}

func (l *lt) resolve(methods map[string]*method, templates *template.Template) error {
	// TODO: implement
	return nil
}

func (l *lt) String() (string, error) {
	// TODO: implement
	return "", nil
}

type gt struct {
	Spec map[string]*param
}

func newGt(unmarshal func(interface{}) error) (action, error) {
	// TODO: implement
	return &gt{}, nil
}

func (g *gt) resolve(methods map[string]*method, templates *template.Template) error {
	// TODO: implement
	return nil
}

func (g *gt) String() (string, error) {
	// TODO: implement
	return "", nil
}

type lte struct {
	Spec map[string]*param
}

func newLte(unmarshal func(interface{}) error) (action, error) {
	// TODO: implement
	return &lte{}, nil
}

func (l *lte) resolve(methods map[string]*method, templates *template.Template) error {
	// TODO: implement
	return nil
}

func (l *lte) String() (string, error) {
	// TODO: implement
	return "", nil
}

type gte struct {
	Spec map[string]*param
}

func newGte(unmarshal func(interface{}) error) (action, error) {
	// TODO: implement
	return &gte{}, nil
}

func (g *gte) resolve(methods map[string]*method, templates *template.Template) error {
	// TODO: implement
	return nil
}

func (g *gte) String() (string, error) {
	// TODO: implement
	return "", nil
}

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
	"fmt"
)

type noTypeError struct {
	p *Param
}

func (n *noTypeError) Error() string {
	return fmt.Sprintf("the type for %s is not set", n.p.Name)
}

type invalidTypeError struct {
	p *Param
}

func (i *invalidTypeError) Error() string {
	return fmt.Sprintf("invalid type for %s: %T (expected %s/%s)", i.p.Name, i.p.Value, i.p.SpecType, i.p.Type)
}

type multipleEnumValuesError struct {
	p *Param
}

func (m *multipleEnumValuesError) Error() string {
	return fmt.Sprintf("multiple values for enum %q", m.p.Name)
}

type invalidDictValueError struct {
	p   *Param
	err error
}

func (i *invalidDictValueError) Error() string {
	return fmt.Sprintf("invalid dict value: %s", i.err)
}

type invalidBulkValueError struct {
	p   *Param
	err error
}

func (i *invalidBulkValueError) Error() string {
	return fmt.Sprintf("invalid bulk value: %s", i.err)
}

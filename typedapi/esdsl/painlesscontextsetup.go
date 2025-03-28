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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _painlessContextSetup struct {
	v *types.PainlessContextSetup
}

func NewPainlessContextSetup(document json.RawMessage) *_painlessContextSetup {

	tmp := &_painlessContextSetup{v: types.NewPainlessContextSetup()}

	tmp.Document(document)

	return tmp

}

// Document that's temporarily indexed in-memory and accessible from the script.
func (s *_painlessContextSetup) Document(document json.RawMessage) *_painlessContextSetup {

	s.v.Document = document

	return s
}

// Index containing a mapping that's compatible with the indexed document.
// You may specify a remote index by prefixing the index with the remote cluster
// alias.
// For example, `remote1:my_index` indicates that you want to run the painless
// script against the "my_index" index on the "remote1" cluster.
// This request will be forwarded to the "remote1" cluster if you have
// configured a connection to that remote cluster.
//
// NOTE: Wildcards are not accepted in the index expression for this endpoint.
// The expression `*:myindex` will return the error "No such remote cluster" and
// the expression `logs*` or `remote1:logs*` will return the error "index not
// found".
func (s *_painlessContextSetup) Index(indexname string) *_painlessContextSetup {

	s.v.Index = indexname

	return s
}

// Use this parameter to specify a query for computing a score.
func (s *_painlessContextSetup) Query(query types.QueryVariant) *_painlessContextSetup {

	s.v.Query = query.QueryCaster()

	return s
}

func (s *_painlessContextSetup) PainlessContextSetupCaster() *types.PainlessContextSetup {
	return s.v
}

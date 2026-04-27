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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fileContent struct {
	v *types.FileContent
}

func NewFileContent(filedata string, filename string) *_fileContent {

	tmp := &_fileContent{v: types.NewFileContent()}

	tmp.FileData(filedata)

	tmp.Filename(filename)

	return tmp

}

func (s *_fileContent) FileData(filedata string) *_fileContent {

	s.v.FileData = filedata

	return s
}

func (s *_fileContent) Filename(filename string) *_fileContent {

	s.v.Filename = filename

	return s
}

func (s *_fileContent) FileContentCaster() *types.FileContent {
	return s.v
}

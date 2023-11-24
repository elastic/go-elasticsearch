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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package storagetype
package storagetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/_types/IndexSettings.ts#L502-L532
type StorageType struct {
	Name string
}

var (
	Fs = StorageType{"fs"}

	Niofs = StorageType{"niofs"}

	Mmapfs = StorageType{"mmapfs"}

	Hybridfs = StorageType{"hybridfs"}
)

func (s StorageType) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *StorageType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "fs":
		*s = Fs
	case "niofs":
		*s = Niofs
	case "mmapfs":
		*s = Mmapfs
	case "hybridfs":
		*s = Hybridfs
	default:
		*s = StorageType{string(text)}
	}

	return nil
}

func (s StorageType) String() string {
	return s.Name
}

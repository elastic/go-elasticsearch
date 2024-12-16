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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

// Package fingerprintdigest
package fingerprintdigest

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/ingest/_types/Processors.ts#L896-L902
type FingerprintDigest struct {
	Name string
}

var (
	Md5 = FingerprintDigest{"MD5"}

	Sha1 = FingerprintDigest{"SHA-1"}

	Sha256 = FingerprintDigest{"SHA-256"}

	Sha512 = FingerprintDigest{"SHA-512"}

	MurmurHash3 = FingerprintDigest{"MurmurHash3"}
)

func (f FingerprintDigest) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FingerprintDigest) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "MD5":
		*f = Md5
	case "SHA-1":
		*f = Sha1
	case "SHA-256":
		*f = Sha256
	case "SHA-512":
		*f = Sha512
	case "MurmurHash3":
		*f = MurmurHash3
	default:
		*f = FingerprintDigest{string(text)}
	}

	return nil
}

func (f FingerprintDigest) String() string {
	return f.Name
}

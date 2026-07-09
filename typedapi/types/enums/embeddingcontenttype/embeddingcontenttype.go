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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

// Package embeddingcontenttype
package embeddingcontenttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/inference/_types/CommonTypes.ts#L666-L675
type EmbeddingContentType struct {
	Name string
}

var (
	Text = EmbeddingContentType{"text"}

	Image = EmbeddingContentType{"image"}

	Audio = EmbeddingContentType{"audio"}

	Video = EmbeddingContentType{"video"}

	Pdf = EmbeddingContentType{"pdf"}
)

func (e EmbeddingContentType) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *EmbeddingContentType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "text":
		*e = Text
	case "image":
		*e = Image
	case "audio":
		*e = Audio
	case "video":
		*e = Video
	case "pdf":
		*e = Pdf
	default:
		*e = EmbeddingContentType{string(text)}
	}

	return nil
}

func (e EmbeddingContentType) String() string {
	return e.Name
}

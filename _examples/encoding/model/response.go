// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

//go:generate easyjson -all -snake_case $GOFILE

package model

type ErrorResponse struct {
	Info *ErrorInfo `json:"error,omitempty"`
}

type ErrorInfo struct {
	RootCause []*ErrorInfo
	Type      string
	Reason    string
	Phase     string
}

type IndexResponse struct {
	Index   string `json:"_index"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Result  string
}

type SearchResponse struct {
	Took int64
	Hits struct {
		Total struct {
			Value int64
		}
		Hits []*SearchHit
	}
}

type SearchHit struct {
	Score   float64 `json:"_score"`
	Index   string  `json:"_index"`
	Type    string  `json:"_type"`
	Version int64   `json:"_version,omitempty"`

	Source Article `json:"_source"`
}

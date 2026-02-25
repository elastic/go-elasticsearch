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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package amazonsagemakerservicetype
package amazonsagemakerservicetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/inference/_types/CommonTypes.ts#L740-L742
type AmazonSageMakerServiceType struct {
	Name string
}

var (
	Amazonsagemaker = AmazonSageMakerServiceType{"amazon_sagemaker"}
)

func (a AmazonSageMakerServiceType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AmazonSageMakerServiceType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "amazon_sagemaker":
		*a = Amazonsagemaker
	default:
		*a = AmazonSageMakerServiceType{string(text)}
	}

	return nil
}

func (a AmazonSageMakerServiceType) String() string {
	return a.Name
}

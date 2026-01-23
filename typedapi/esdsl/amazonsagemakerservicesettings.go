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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakerapi"
)

type _amazonSageMakerServiceSettings struct {
	v *types.AmazonSageMakerServiceSettings
}

func NewAmazonSageMakerServiceSettings(accesskey string, api amazonsagemakerapi.AmazonSageMakerApi, endpointname string, region string, secretkey string) *_amazonSageMakerServiceSettings {

	tmp := &_amazonSageMakerServiceSettings{v: types.NewAmazonSageMakerServiceSettings()}

	tmp.AccessKey(accesskey)

	tmp.Api(api)

	tmp.EndpointName(endpointname)

	tmp.Region(region)

	tmp.SecretKey(secretkey)

	return tmp

}

func (s *_amazonSageMakerServiceSettings) AccessKey(accesskey string) *_amazonSageMakerServiceSettings {

	s.v.AccessKey = accesskey

	return s
}

func (s *_amazonSageMakerServiceSettings) Api(api amazonsagemakerapi.AmazonSageMakerApi) *_amazonSageMakerServiceSettings {

	s.v.Api = api
	return s
}

func (s *_amazonSageMakerServiceSettings) BatchSize(batchsize int) *_amazonSageMakerServiceSettings {

	s.v.BatchSize = &batchsize

	return s
}

func (s *_amazonSageMakerServiceSettings) Dimensions(dimensions int) *_amazonSageMakerServiceSettings {

	s.v.Dimensions = &dimensions

	return s
}

func (s *_amazonSageMakerServiceSettings) EndpointName(endpointname string) *_amazonSageMakerServiceSettings {

	s.v.EndpointName = endpointname

	return s
}

func (s *_amazonSageMakerServiceSettings) InferenceComponentName(inferencecomponentname string) *_amazonSageMakerServiceSettings {

	s.v.InferenceComponentName = &inferencecomponentname

	return s
}

func (s *_amazonSageMakerServiceSettings) Region(region string) *_amazonSageMakerServiceSettings {

	s.v.Region = region

	return s
}

func (s *_amazonSageMakerServiceSettings) SecretKey(secretkey string) *_amazonSageMakerServiceSettings {

	s.v.SecretKey = secretkey

	return s
}

func (s *_amazonSageMakerServiceSettings) TargetContainerHostname(targetcontainerhostname string) *_amazonSageMakerServiceSettings {

	s.v.TargetContainerHostname = &targetcontainerhostname

	return s
}

func (s *_amazonSageMakerServiceSettings) TargetModel(targetmodel string) *_amazonSageMakerServiceSettings {

	s.v.TargetModel = &targetmodel

	return s
}

func (s *_amazonSageMakerServiceSettings) AmazonSageMakerServiceSettingsCaster() *types.AmazonSageMakerServiceSettings {
	return s.v
}

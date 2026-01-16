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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectorfieldtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/displaytype"
)

type _connectorConfigProperties struct {
	v *types.ConnectorConfigProperties
}

func NewConnectorConfigProperties(display displaytype.DisplayType, label string, required bool, sensitive bool, value json.RawMessage) *_connectorConfigProperties {

	tmp := &_connectorConfigProperties{v: types.NewConnectorConfigProperties()}

	tmp.Display(display)

	tmp.Label(label)

	tmp.Required(required)

	tmp.Sensitive(sensitive)

	tmp.Value(value)

	return tmp

}

func (s *_connectorConfigProperties) Category(category string) *_connectorConfigProperties {

	s.v.Category = &category

	return s
}

func (s *_connectorConfigProperties) DefaultValue(scalarvalue types.ScalarValueVariant) *_connectorConfigProperties {

	s.v.DefaultValue = *scalarvalue.ScalarValueCaster()

	return s
}

func (s *_connectorConfigProperties) DependsOn(dependsons ...types.DependencyVariant) *_connectorConfigProperties {

	for _, v := range dependsons {

		s.v.DependsOn = append(s.v.DependsOn, *v.DependencyCaster())

	}
	return s
}

func (s *_connectorConfigProperties) Display(display displaytype.DisplayType) *_connectorConfigProperties {

	s.v.Display = display
	return s
}

func (s *_connectorConfigProperties) Label(label string) *_connectorConfigProperties {

	s.v.Label = label

	return s
}

func (s *_connectorConfigProperties) Options(options ...types.SelectOptionVariant) *_connectorConfigProperties {

	for _, v := range options {

		s.v.Options = append(s.v.Options, *v.SelectOptionCaster())

	}
	return s
}

func (s *_connectorConfigProperties) Order(order int) *_connectorConfigProperties {

	s.v.Order = &order

	return s
}

func (s *_connectorConfigProperties) Placeholder(placeholder string) *_connectorConfigProperties {

	s.v.Placeholder = &placeholder

	return s
}

func (s *_connectorConfigProperties) Required(required bool) *_connectorConfigProperties {

	s.v.Required = required

	return s
}

func (s *_connectorConfigProperties) Sensitive(sensitive bool) *_connectorConfigProperties {

	s.v.Sensitive = sensitive

	return s
}

func (s *_connectorConfigProperties) Tooltip(tooltip string) *_connectorConfigProperties {

	s.v.Tooltip = &tooltip

	return s
}

func (s *_connectorConfigProperties) Type(type_ connectorfieldtype.ConnectorFieldType) *_connectorConfigProperties {

	s.v.Type = &type_
	return s
}

func (s *_connectorConfigProperties) UiRestrictions(uirestrictions ...string) *_connectorConfigProperties {

	for _, v := range uirestrictions {

		s.v.UiRestrictions = append(s.v.UiRestrictions, v)

	}
	return s
}

func (s *_connectorConfigProperties) Validations(validations ...types.ValidationVariant) *_connectorConfigProperties {

	for _, v := range validations {

		s.v.Validations = append(s.v.Validations, *v.ValidationCaster())

	}
	return s
}

func (s *_connectorConfigProperties) Value(value json.RawMessage) *_connectorConfigProperties {

	s.v.Value = value

	return s
}

func (s *_connectorConfigProperties) ConnectorConfigPropertiesCaster() *types.ConnectorConfigProperties {
	return s.v
}

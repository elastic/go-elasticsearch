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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

// Diagnosis type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_global/health_report/types.ts#L48-L54
type Diagnosis struct {
	Action            string                     `json:"action"`
	AffectedResources DiagnosisAffectedResources `json:"affected_resources"`
	Cause             string                     `json:"cause"`
	HelpUrl           string                     `json:"help_url"`
	Id                string                     `json:"id"`
}

// NewDiagnosis returns a Diagnosis.
func NewDiagnosis() *Diagnosis {
	r := &Diagnosis{}

	return r
}

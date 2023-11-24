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

// Package clusterprivilege
package clusterprivilege

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/security/_types/Privileges.ts#L41-L80
type ClusterPrivilege struct {
	Name string
}

var (
	All = ClusterPrivilege{"all"}

	Canceltask = ClusterPrivilege{"cancel_task"}

	Createsnapshot = ClusterPrivilege{"create_snapshot"}

	Grantapikey = ClusterPrivilege{"grant_api_key"}

	Manage = ClusterPrivilege{"manage"}

	Manageapikey = ClusterPrivilege{"manage_api_key"}

	Manageccr = ClusterPrivilege{"manage_ccr"}

	Manageenrich = ClusterPrivilege{"manage_enrich"}

	Manageilm = ClusterPrivilege{"manage_ilm"}

	Manageindextemplates = ClusterPrivilege{"manage_index_templates"}

	Manageingestpipelines = ClusterPrivilege{"manage_ingest_pipelines"}

	Managelogstashpipelines = ClusterPrivilege{"manage_logstash_pipelines"}

	Manageml = ClusterPrivilege{"manage_ml"}

	Manageoidc = ClusterPrivilege{"manage_oidc"}

	Manageownapikey = ClusterPrivilege{"manage_own_api_key"}

	Managepipeline = ClusterPrivilege{"manage_pipeline"}

	Managerollup = ClusterPrivilege{"manage_rollup"}

	Managesaml = ClusterPrivilege{"manage_saml"}

	Managesecurity = ClusterPrivilege{"manage_security"}

	Manageserviceaccount = ClusterPrivilege{"manage_service_account"}

	Manageslm = ClusterPrivilege{"manage_slm"}

	Managetoken = ClusterPrivilege{"manage_token"}

	Managetransform = ClusterPrivilege{"manage_transform"}

	Manageuserprofile = ClusterPrivilege{"manage_user_profile"}

	Managewatcher = ClusterPrivilege{"manage_watcher"}

	Monitor = ClusterPrivilege{"monitor"}

	Monitorml = ClusterPrivilege{"monitor_ml"}

	Monitorrollup = ClusterPrivilege{"monitor_rollup"}

	Monitorsnapshot = ClusterPrivilege{"monitor_snapshot"}

	Monitortextstructure = ClusterPrivilege{"monitor_text_structure"}

	Monitortransform = ClusterPrivilege{"monitor_transform"}

	Monitorwatcher = ClusterPrivilege{"monitor_watcher"}

	Readccr = ClusterPrivilege{"read_ccr"}

	Readilm = ClusterPrivilege{"read_ilm"}

	Readpipeline = ClusterPrivilege{"read_pipeline"}

	Readslm = ClusterPrivilege{"read_slm"}

	Transportclient = ClusterPrivilege{"transport_client"}
)

func (c ClusterPrivilege) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ClusterPrivilege) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "all":
		*c = All
	case "cancel_task":
		*c = Canceltask
	case "create_snapshot":
		*c = Createsnapshot
	case "grant_api_key":
		*c = Grantapikey
	case "manage":
		*c = Manage
	case "manage_api_key":
		*c = Manageapikey
	case "manage_ccr":
		*c = Manageccr
	case "manage_enrich":
		*c = Manageenrich
	case "manage_ilm":
		*c = Manageilm
	case "manage_index_templates":
		*c = Manageindextemplates
	case "manage_ingest_pipelines":
		*c = Manageingestpipelines
	case "manage_logstash_pipelines":
		*c = Managelogstashpipelines
	case "manage_ml":
		*c = Manageml
	case "manage_oidc":
		*c = Manageoidc
	case "manage_own_api_key":
		*c = Manageownapikey
	case "manage_pipeline":
		*c = Managepipeline
	case "manage_rollup":
		*c = Managerollup
	case "manage_saml":
		*c = Managesaml
	case "manage_security":
		*c = Managesecurity
	case "manage_service_account":
		*c = Manageserviceaccount
	case "manage_slm":
		*c = Manageslm
	case "manage_token":
		*c = Managetoken
	case "manage_transform":
		*c = Managetransform
	case "manage_user_profile":
		*c = Manageuserprofile
	case "manage_watcher":
		*c = Managewatcher
	case "monitor":
		*c = Monitor
	case "monitor_ml":
		*c = Monitorml
	case "monitor_rollup":
		*c = Monitorrollup
	case "monitor_snapshot":
		*c = Monitorsnapshot
	case "monitor_text_structure":
		*c = Monitortextstructure
	case "monitor_transform":
		*c = Monitortransform
	case "monitor_watcher":
		*c = Monitorwatcher
	case "read_ccr":
		*c = Readccr
	case "read_ilm":
		*c = Readilm
	case "read_pipeline":
		*c = Readpipeline
	case "read_slm":
		*c = Readslm
	case "transport_client":
		*c = Transportclient
	default:
		*c = ClusterPrivilege{string(text)}
	}

	return nil
}

func (c ClusterPrivilege) String() string {
	return c.Name
}

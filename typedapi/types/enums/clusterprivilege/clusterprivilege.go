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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package clusterprivilege
package clusterprivilege

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/_types/Privileges.ts#L42-L199
type ClusterPrivilege struct {
	Name string
}

var (
	All = ClusterPrivilege{"all"}

	Canceltask = ClusterPrivilege{"cancel_task"}

	Createsnapshot = ClusterPrivilege{"create_snapshot"}

	Crossclusterreplication = ClusterPrivilege{"cross_cluster_replication"}

	Crossclustersearch = ClusterPrivilege{"cross_cluster_search"}

	Delegatepki = ClusterPrivilege{"delegate_pki"}

	Grantapikey = ClusterPrivilege{"grant_api_key"}

	Manage = ClusterPrivilege{"manage"}

	Manageapikey = ClusterPrivilege{"manage_api_key"}

	Manageautoscaling = ClusterPrivilege{"manage_autoscaling"}

	Managebehavioralanalytics = ClusterPrivilege{"manage_behavioral_analytics"}

	Manageccr = ClusterPrivilege{"manage_ccr"}

	Managedataframetransforms = ClusterPrivilege{"manage_data_frame_transforms"}

	Managedatastreamglobalretention = ClusterPrivilege{"manage_data_stream_global_retention"}

	Manageenrich = ClusterPrivilege{"manage_enrich"}

	Manageilm = ClusterPrivilege{"manage_ilm"}

	Manageindextemplates = ClusterPrivilege{"manage_index_templates"}

	Manageinference = ClusterPrivilege{"manage_inference"}

	Manageingestpipelines = ClusterPrivilege{"manage_ingest_pipelines"}

	Managelogstashpipelines = ClusterPrivilege{"manage_logstash_pipelines"}

	Manageml = ClusterPrivilege{"manage_ml"}

	Manageoidc = ClusterPrivilege{"manage_oidc"}

	Manageownapikey = ClusterPrivilege{"manage_own_api_key"}

	Managepipeline = ClusterPrivilege{"manage_pipeline"}

	Managerollup = ClusterPrivilege{"manage_rollup"}

	Managesaml = ClusterPrivilege{"manage_saml"}

	Managesearchapplication = ClusterPrivilege{"manage_search_application"}

	Managesearchqueryrules = ClusterPrivilege{"manage_search_query_rules"}

	Managesearchsynonyms = ClusterPrivilege{"manage_search_synonyms"}

	Managesecurity = ClusterPrivilege{"manage_security"}

	Manageserviceaccount = ClusterPrivilege{"manage_service_account"}

	Manageslm = ClusterPrivilege{"manage_slm"}

	Managetoken = ClusterPrivilege{"manage_token"}

	Managetransform = ClusterPrivilege{"manage_transform"}

	Manageuserprofile = ClusterPrivilege{"manage_user_profile"}

	Managewatcher = ClusterPrivilege{"manage_watcher"}

	Monitor = ClusterPrivilege{"monitor"}

	Monitordataframetransforms = ClusterPrivilege{"monitor_data_frame_transforms"}

	Monitordatastreamglobalretention = ClusterPrivilege{"monitor_data_stream_global_retention"}

	Monitorenrich = ClusterPrivilege{"monitor_enrich"}

	Monitorinference = ClusterPrivilege{"monitor_inference"}

	Monitorml = ClusterPrivilege{"monitor_ml"}

	Monitorrollup = ClusterPrivilege{"monitor_rollup"}

	Monitorsnapshot = ClusterPrivilege{"monitor_snapshot"}

	Monitorstats = ClusterPrivilege{"monitor_stats"}

	Monitortextstructure = ClusterPrivilege{"monitor_text_structure"}

	Monitortransform = ClusterPrivilege{"monitor_transform"}

	Monitorwatcher = ClusterPrivilege{"monitor_watcher"}

	None = ClusterPrivilege{"none"}

	Postbehavioralanalyticsevent = ClusterPrivilege{"post_behavioral_analytics_event"}

	Readccr = ClusterPrivilege{"read_ccr"}

	Readfleetsecrets = ClusterPrivilege{"read_fleet_secrets"}

	Readilm = ClusterPrivilege{"read_ilm"}

	Readpipeline = ClusterPrivilege{"read_pipeline"}

	Readsecurity = ClusterPrivilege{"read_security"}

	Readslm = ClusterPrivilege{"read_slm"}

	Transportclient = ClusterPrivilege{"transport_client"}

	Writeconnectorsecrets = ClusterPrivilege{"write_connector_secrets"}

	Writefleetsecrets = ClusterPrivilege{"write_fleet_secrets"}
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
	case "cross_cluster_replication":
		*c = Crossclusterreplication
	case "cross_cluster_search":
		*c = Crossclustersearch
	case "delegate_pki":
		*c = Delegatepki
	case "grant_api_key":
		*c = Grantapikey
	case "manage":
		*c = Manage
	case "manage_api_key":
		*c = Manageapikey
	case "manage_autoscaling":
		*c = Manageautoscaling
	case "manage_behavioral_analytics":
		*c = Managebehavioralanalytics
	case "manage_ccr":
		*c = Manageccr
	case "manage_data_frame_transforms":
		*c = Managedataframetransforms
	case "manage_data_stream_global_retention":
		*c = Managedatastreamglobalretention
	case "manage_enrich":
		*c = Manageenrich
	case "manage_ilm":
		*c = Manageilm
	case "manage_index_templates":
		*c = Manageindextemplates
	case "manage_inference":
		*c = Manageinference
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
	case "manage_search_application":
		*c = Managesearchapplication
	case "manage_search_query_rules":
		*c = Managesearchqueryrules
	case "manage_search_synonyms":
		*c = Managesearchsynonyms
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
	case "monitor_data_frame_transforms":
		*c = Monitordataframetransforms
	case "monitor_data_stream_global_retention":
		*c = Monitordatastreamglobalretention
	case "monitor_enrich":
		*c = Monitorenrich
	case "monitor_inference":
		*c = Monitorinference
	case "monitor_ml":
		*c = Monitorml
	case "monitor_rollup":
		*c = Monitorrollup
	case "monitor_snapshot":
		*c = Monitorsnapshot
	case "monitor_stats":
		*c = Monitorstats
	case "monitor_text_structure":
		*c = Monitortextstructure
	case "monitor_transform":
		*c = Monitortransform
	case "monitor_watcher":
		*c = Monitorwatcher
	case "none":
		*c = None
	case "post_behavioral_analytics_event":
		*c = Postbehavioralanalyticsevent
	case "read_ccr":
		*c = Readccr
	case "read_fleet_secrets":
		*c = Readfleetsecrets
	case "read_ilm":
		*c = Readilm
	case "read_pipeline":
		*c = Readpipeline
	case "read_security":
		*c = Readsecurity
	case "read_slm":
		*c = Readslm
	case "transport_client":
		*c = Transportclient
	case "write_connector_secrets":
		*c = Writeconnectorsecrets
	case "write_fleet_secrets":
		*c = Writefleetsecrets
	default:
		*c = ClusterPrivilege{string(text)}
	}

	return nil
}

func (c ClusterPrivilege) String() string {
	return c.Name
}

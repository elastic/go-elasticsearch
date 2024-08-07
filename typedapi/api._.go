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
// https://github.com/elastic/elasticsearch-specification/tree/19027dbdd366978ccae41842a040a636730e7c10

package typedapi

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	async_search_delete "github.com/elastic/go-elasticsearch/v8/typedapi/asyncsearch/delete"
	async_search_get "github.com/elastic/go-elasticsearch/v8/typedapi/asyncsearch/get"
	async_search_status "github.com/elastic/go-elasticsearch/v8/typedapi/asyncsearch/status"
	async_search_submit "github.com/elastic/go-elasticsearch/v8/typedapi/asyncsearch/submit"
	autoscaling_delete_autoscaling_policy "github.com/elastic/go-elasticsearch/v8/typedapi/autoscaling/deleteautoscalingpolicy"
	autoscaling_get_autoscaling_capacity "github.com/elastic/go-elasticsearch/v8/typedapi/autoscaling/getautoscalingcapacity"
	autoscaling_get_autoscaling_policy "github.com/elastic/go-elasticsearch/v8/typedapi/autoscaling/getautoscalingpolicy"
	autoscaling_put_autoscaling_policy "github.com/elastic/go-elasticsearch/v8/typedapi/autoscaling/putautoscalingpolicy"
	capabilities "github.com/elastic/go-elasticsearch/v8/typedapi/capabilities"
	cat_aliases "github.com/elastic/go-elasticsearch/v8/typedapi/cat/aliases"
	cat_allocation "github.com/elastic/go-elasticsearch/v8/typedapi/cat/allocation"
	cat_component_templates "github.com/elastic/go-elasticsearch/v8/typedapi/cat/componenttemplates"
	cat_count "github.com/elastic/go-elasticsearch/v8/typedapi/cat/count"
	cat_fielddata "github.com/elastic/go-elasticsearch/v8/typedapi/cat/fielddata"
	cat_health "github.com/elastic/go-elasticsearch/v8/typedapi/cat/health"
	cat_help "github.com/elastic/go-elasticsearch/v8/typedapi/cat/help"
	cat_indices "github.com/elastic/go-elasticsearch/v8/typedapi/cat/indices"
	cat_master "github.com/elastic/go-elasticsearch/v8/typedapi/cat/master"
	cat_ml_datafeeds "github.com/elastic/go-elasticsearch/v8/typedapi/cat/mldatafeeds"
	cat_ml_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/cat/mldataframeanalytics"
	cat_ml_jobs "github.com/elastic/go-elasticsearch/v8/typedapi/cat/mljobs"
	cat_ml_trained_models "github.com/elastic/go-elasticsearch/v8/typedapi/cat/mltrainedmodels"
	cat_nodeattrs "github.com/elastic/go-elasticsearch/v8/typedapi/cat/nodeattrs"
	cat_nodes "github.com/elastic/go-elasticsearch/v8/typedapi/cat/nodes"
	cat_pending_tasks "github.com/elastic/go-elasticsearch/v8/typedapi/cat/pendingtasks"
	cat_plugins "github.com/elastic/go-elasticsearch/v8/typedapi/cat/plugins"
	cat_recovery "github.com/elastic/go-elasticsearch/v8/typedapi/cat/recovery"
	cat_repositories "github.com/elastic/go-elasticsearch/v8/typedapi/cat/repositories"
	cat_segments "github.com/elastic/go-elasticsearch/v8/typedapi/cat/segments"
	cat_shards "github.com/elastic/go-elasticsearch/v8/typedapi/cat/shards"
	cat_snapshots "github.com/elastic/go-elasticsearch/v8/typedapi/cat/snapshots"
	cat_tasks "github.com/elastic/go-elasticsearch/v8/typedapi/cat/tasks"
	cat_templates "github.com/elastic/go-elasticsearch/v8/typedapi/cat/templates"
	cat_thread_pool "github.com/elastic/go-elasticsearch/v8/typedapi/cat/threadpool"
	cat_transforms "github.com/elastic/go-elasticsearch/v8/typedapi/cat/transforms"
	ccr_delete_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/deleteautofollowpattern"
	ccr_follow "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/follow"
	ccr_follow_info "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/followinfo"
	ccr_follow_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/followstats"
	ccr_forget_follower "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/forgetfollower"
	ccr_get_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/getautofollowpattern"
	ccr_pause_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/pauseautofollowpattern"
	ccr_pause_follow "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/pausefollow"
	ccr_put_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/putautofollowpattern"
	ccr_resume_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/resumeautofollowpattern"
	ccr_resume_follow "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/resumefollow"
	ccr_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/stats"
	ccr_unfollow "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/unfollow"
	cluster_allocation_explain "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/allocationexplain"
	cluster_delete_component_template "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/deletecomponenttemplate"
	cluster_delete_voting_config_exclusions "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/deletevotingconfigexclusions"
	cluster_exists_component_template "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/existscomponenttemplate"
	cluster_get_component_template "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/getcomponenttemplate"
	cluster_get_settings "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/getsettings"
	cluster_health "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/health"
	cluster_info "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/info"
	cluster_pending_tasks "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/pendingtasks"
	cluster_post_voting_config_exclusions "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/postvotingconfigexclusions"
	cluster_put_component_template "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/putcomponenttemplate"
	cluster_put_settings "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/putsettings"
	cluster_remote_info "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/remoteinfo"
	cluster_reroute "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/reroute"
	cluster_state "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/state"
	cluster_stats "github.com/elastic/go-elasticsearch/v8/typedapi/cluster/stats"
	connector_check_in "github.com/elastic/go-elasticsearch/v8/typedapi/connector/checkin"
	connector_delete "github.com/elastic/go-elasticsearch/v8/typedapi/connector/delete"
	connector_get "github.com/elastic/go-elasticsearch/v8/typedapi/connector/get"
	connector_last_sync "github.com/elastic/go-elasticsearch/v8/typedapi/connector/lastsync"
	connector_list "github.com/elastic/go-elasticsearch/v8/typedapi/connector/list"
	connector_post "github.com/elastic/go-elasticsearch/v8/typedapi/connector/post"
	connector_put "github.com/elastic/go-elasticsearch/v8/typedapi/connector/put"
	connector_secret_post "github.com/elastic/go-elasticsearch/v8/typedapi/connector/secretpost"
	connector_sync_job_cancel "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobcancel"
	connector_sync_job_delete "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobdelete"
	connector_sync_job_get "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobget"
	connector_sync_job_list "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjoblist"
	connector_sync_job_post "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobpost"
	connector_update_active_filtering "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateactivefiltering"
	connector_update_api_key_id "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateapikeyid"
	connector_update_configuration "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateconfiguration"
	connector_update_error "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateerror"
	connector_update_filtering "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatefiltering"
	connector_update_filtering_validation "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatefilteringvalidation"
	connector_update_index_name "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateindexname"
	connector_update_name "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatename"
	connector_update_native "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatenative"
	connector_update_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatepipeline"
	connector_update_scheduling "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatescheduling"
	connector_update_service_type "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateservicetype"
	connector_update_status "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatestatus"
	core_bulk "github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	core_clear_scroll "github.com/elastic/go-elasticsearch/v8/typedapi/core/clearscroll"
	core_close_point_in_time "github.com/elastic/go-elasticsearch/v8/typedapi/core/closepointintime"
	core_count "github.com/elastic/go-elasticsearch/v8/typedapi/core/count"
	core_create "github.com/elastic/go-elasticsearch/v8/typedapi/core/create"
	core_delete "github.com/elastic/go-elasticsearch/v8/typedapi/core/delete"
	core_delete_by_query "github.com/elastic/go-elasticsearch/v8/typedapi/core/deletebyquery"
	core_delete_by_query_rethrottle "github.com/elastic/go-elasticsearch/v8/typedapi/core/deletebyqueryrethrottle"
	core_delete_script "github.com/elastic/go-elasticsearch/v8/typedapi/core/deletescript"
	core_exists "github.com/elastic/go-elasticsearch/v8/typedapi/core/exists"
	core_exists_source "github.com/elastic/go-elasticsearch/v8/typedapi/core/existssource"
	core_explain "github.com/elastic/go-elasticsearch/v8/typedapi/core/explain"
	core_field_caps "github.com/elastic/go-elasticsearch/v8/typedapi/core/fieldcaps"
	core_get "github.com/elastic/go-elasticsearch/v8/typedapi/core/get"
	core_get_script "github.com/elastic/go-elasticsearch/v8/typedapi/core/getscript"
	core_get_script_context "github.com/elastic/go-elasticsearch/v8/typedapi/core/getscriptcontext"
	core_get_script_languages "github.com/elastic/go-elasticsearch/v8/typedapi/core/getscriptlanguages"
	core_get_source "github.com/elastic/go-elasticsearch/v8/typedapi/core/getsource"
	core_health_report "github.com/elastic/go-elasticsearch/v8/typedapi/core/healthreport"
	core_index "github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
	core_info "github.com/elastic/go-elasticsearch/v8/typedapi/core/info"
	core_knn_search "github.com/elastic/go-elasticsearch/v8/typedapi/core/knnsearch"
	core_mget "github.com/elastic/go-elasticsearch/v8/typedapi/core/mget"
	core_msearch "github.com/elastic/go-elasticsearch/v8/typedapi/core/msearch"
	core_msearch_template "github.com/elastic/go-elasticsearch/v8/typedapi/core/msearchtemplate"
	core_mtermvectors "github.com/elastic/go-elasticsearch/v8/typedapi/core/mtermvectors"
	core_open_point_in_time "github.com/elastic/go-elasticsearch/v8/typedapi/core/openpointintime"
	core_ping "github.com/elastic/go-elasticsearch/v8/typedapi/core/ping"
	core_put_script "github.com/elastic/go-elasticsearch/v8/typedapi/core/putscript"
	core_rank_eval "github.com/elastic/go-elasticsearch/v8/typedapi/core/rankeval"
	core_reindex "github.com/elastic/go-elasticsearch/v8/typedapi/core/reindex"
	core_reindex_rethrottle "github.com/elastic/go-elasticsearch/v8/typedapi/core/reindexrethrottle"
	core_render_search_template "github.com/elastic/go-elasticsearch/v8/typedapi/core/rendersearchtemplate"
	core_scripts_painless_execute "github.com/elastic/go-elasticsearch/v8/typedapi/core/scriptspainlessexecute"
	core_scroll "github.com/elastic/go-elasticsearch/v8/typedapi/core/scroll"
	core_search "github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	core_search_mvt "github.com/elastic/go-elasticsearch/v8/typedapi/core/searchmvt"
	core_search_shards "github.com/elastic/go-elasticsearch/v8/typedapi/core/searchshards"
	core_search_template "github.com/elastic/go-elasticsearch/v8/typedapi/core/searchtemplate"
	core_terms_enum "github.com/elastic/go-elasticsearch/v8/typedapi/core/termsenum"
	core_termvectors "github.com/elastic/go-elasticsearch/v8/typedapi/core/termvectors"
	core_update "github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	core_update_by_query "github.com/elastic/go-elasticsearch/v8/typedapi/core/updatebyquery"
	core_update_by_query_rethrottle "github.com/elastic/go-elasticsearch/v8/typedapi/core/updatebyqueryrethrottle"
	dangling_indices_delete_dangling_index "github.com/elastic/go-elasticsearch/v8/typedapi/danglingindices/deletedanglingindex"
	dangling_indices_import_dangling_index "github.com/elastic/go-elasticsearch/v8/typedapi/danglingindices/importdanglingindex"
	dangling_indices_list_dangling_indices "github.com/elastic/go-elasticsearch/v8/typedapi/danglingindices/listdanglingindices"
	enrich_delete_policy "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/deletepolicy"
	enrich_execute_policy "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/executepolicy"
	enrich_get_policy "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/getpolicy"
	enrich_put_policy "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/putpolicy"
	enrich_stats "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/stats"
	eql_delete "github.com/elastic/go-elasticsearch/v8/typedapi/eql/delete"
	eql_get "github.com/elastic/go-elasticsearch/v8/typedapi/eql/get"
	eql_get_status "github.com/elastic/go-elasticsearch/v8/typedapi/eql/getstatus"
	eql_search "github.com/elastic/go-elasticsearch/v8/typedapi/eql/search"
	esql_async_query "github.com/elastic/go-elasticsearch/v8/typedapi/esql/asyncquery"
	esql_query "github.com/elastic/go-elasticsearch/v8/typedapi/esql/query"
	features_get_features "github.com/elastic/go-elasticsearch/v8/typedapi/features/getfeatures"
	features_reset_features "github.com/elastic/go-elasticsearch/v8/typedapi/features/resetfeatures"
	fleet_global_checkpoints "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/globalcheckpoints"
	fleet_msearch "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/msearch"
	fleet_post_secret "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/postsecret"
	fleet_search "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/search"
	graph_explore "github.com/elastic/go-elasticsearch/v8/typedapi/graph/explore"
	ilm_delete_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/deletelifecycle"
	ilm_explain_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/explainlifecycle"
	ilm_get_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/getlifecycle"
	ilm_get_status "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/getstatus"
	ilm_migrate_to_data_tiers "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/migratetodatatiers"
	ilm_move_to_step "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/movetostep"
	ilm_put_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/putlifecycle"
	ilm_remove_policy "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/removepolicy"
	ilm_retry "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/retry"
	ilm_start "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/start"
	ilm_stop "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/stop"
	indices_add_block "github.com/elastic/go-elasticsearch/v8/typedapi/indices/addblock"
	indices_analyze "github.com/elastic/go-elasticsearch/v8/typedapi/indices/analyze"
	indices_clear_cache "github.com/elastic/go-elasticsearch/v8/typedapi/indices/clearcache"
	indices_clone "github.com/elastic/go-elasticsearch/v8/typedapi/indices/clone"
	indices_close "github.com/elastic/go-elasticsearch/v8/typedapi/indices/close"
	indices_create "github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	indices_create_data_stream "github.com/elastic/go-elasticsearch/v8/typedapi/indices/createdatastream"
	indices_data_streams_stats "github.com/elastic/go-elasticsearch/v8/typedapi/indices/datastreamsstats"
	indices_delete "github.com/elastic/go-elasticsearch/v8/typedapi/indices/delete"
	indices_delete_alias "github.com/elastic/go-elasticsearch/v8/typedapi/indices/deletealias"
	indices_delete_data_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/indices/deletedatalifecycle"
	indices_delete_data_stream "github.com/elastic/go-elasticsearch/v8/typedapi/indices/deletedatastream"
	indices_delete_index_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/deleteindextemplate"
	indices_delete_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/deletetemplate"
	indices_disk_usage "github.com/elastic/go-elasticsearch/v8/typedapi/indices/diskusage"
	indices_downsample "github.com/elastic/go-elasticsearch/v8/typedapi/indices/downsample"
	indices_exists "github.com/elastic/go-elasticsearch/v8/typedapi/indices/exists"
	indices_exists_alias "github.com/elastic/go-elasticsearch/v8/typedapi/indices/existsalias"
	indices_exists_index_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/existsindextemplate"
	indices_exists_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/existstemplate"
	indices_explain_data_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/indices/explaindatalifecycle"
	indices_field_usage_stats "github.com/elastic/go-elasticsearch/v8/typedapi/indices/fieldusagestats"
	indices_flush "github.com/elastic/go-elasticsearch/v8/typedapi/indices/flush"
	indices_forcemerge "github.com/elastic/go-elasticsearch/v8/typedapi/indices/forcemerge"
	indices_get "github.com/elastic/go-elasticsearch/v8/typedapi/indices/get"
	indices_get_alias "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getalias"
	indices_get_data_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getdatalifecycle"
	indices_get_data_stream "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getdatastream"
	indices_get_field_mapping "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getfieldmapping"
	indices_get_index_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getindextemplate"
	indices_get_mapping "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getmapping"
	indices_get_settings "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getsettings"
	indices_get_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/gettemplate"
	indices_migrate_to_data_stream "github.com/elastic/go-elasticsearch/v8/typedapi/indices/migratetodatastream"
	indices_modify_data_stream "github.com/elastic/go-elasticsearch/v8/typedapi/indices/modifydatastream"
	indices_open "github.com/elastic/go-elasticsearch/v8/typedapi/indices/open"
	indices_promote_data_stream "github.com/elastic/go-elasticsearch/v8/typedapi/indices/promotedatastream"
	indices_put_alias "github.com/elastic/go-elasticsearch/v8/typedapi/indices/putalias"
	indices_put_data_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/indices/putdatalifecycle"
	indices_put_index_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/putindextemplate"
	indices_put_mapping "github.com/elastic/go-elasticsearch/v8/typedapi/indices/putmapping"
	indices_put_settings "github.com/elastic/go-elasticsearch/v8/typedapi/indices/putsettings"
	indices_put_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/puttemplate"
	indices_recovery "github.com/elastic/go-elasticsearch/v8/typedapi/indices/recovery"
	indices_refresh "github.com/elastic/go-elasticsearch/v8/typedapi/indices/refresh"
	indices_reload_search_analyzers "github.com/elastic/go-elasticsearch/v8/typedapi/indices/reloadsearchanalyzers"
	indices_resolve_cluster "github.com/elastic/go-elasticsearch/v8/typedapi/indices/resolvecluster"
	indices_resolve_index "github.com/elastic/go-elasticsearch/v8/typedapi/indices/resolveindex"
	indices_rollover "github.com/elastic/go-elasticsearch/v8/typedapi/indices/rollover"
	indices_segments "github.com/elastic/go-elasticsearch/v8/typedapi/indices/segments"
	indices_shard_stores "github.com/elastic/go-elasticsearch/v8/typedapi/indices/shardstores"
	indices_shrink "github.com/elastic/go-elasticsearch/v8/typedapi/indices/shrink"
	indices_simulate_index_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/simulateindextemplate"
	indices_simulate_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/simulatetemplate"
	indices_split "github.com/elastic/go-elasticsearch/v8/typedapi/indices/split"
	indices_stats "github.com/elastic/go-elasticsearch/v8/typedapi/indices/stats"
	indices_unfreeze "github.com/elastic/go-elasticsearch/v8/typedapi/indices/unfreeze"
	indices_update_aliases "github.com/elastic/go-elasticsearch/v8/typedapi/indices/updatealiases"
	indices_validate_query "github.com/elastic/go-elasticsearch/v8/typedapi/indices/validatequery"
	inference_delete "github.com/elastic/go-elasticsearch/v8/typedapi/inference/delete"
	inference_get "github.com/elastic/go-elasticsearch/v8/typedapi/inference/get"
	inference_inference "github.com/elastic/go-elasticsearch/v8/typedapi/inference/inference"
	inference_put "github.com/elastic/go-elasticsearch/v8/typedapi/inference/put"
	ingest_delete_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/deletepipeline"
	ingest_geo_ip_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/geoipstats"
	ingest_get_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/getpipeline"
	ingest_processor_grok "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/processorgrok"
	ingest_put_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/putpipeline"
	ingest_simulate "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/simulate"
	license_delete "github.com/elastic/go-elasticsearch/v8/typedapi/license/delete"
	license_get "github.com/elastic/go-elasticsearch/v8/typedapi/license/get"
	license_get_basic_status "github.com/elastic/go-elasticsearch/v8/typedapi/license/getbasicstatus"
	license_get_trial_status "github.com/elastic/go-elasticsearch/v8/typedapi/license/gettrialstatus"
	license_post "github.com/elastic/go-elasticsearch/v8/typedapi/license/post"
	license_post_start_basic "github.com/elastic/go-elasticsearch/v8/typedapi/license/poststartbasic"
	license_post_start_trial "github.com/elastic/go-elasticsearch/v8/typedapi/license/poststarttrial"
	logstash_delete_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/logstash/deletepipeline"
	logstash_get_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/logstash/getpipeline"
	logstash_put_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/logstash/putpipeline"
	migration_deprecations "github.com/elastic/go-elasticsearch/v8/typedapi/migration/deprecations"
	migration_get_feature_upgrade_status "github.com/elastic/go-elasticsearch/v8/typedapi/migration/getfeatureupgradestatus"
	migration_post_feature_upgrade "github.com/elastic/go-elasticsearch/v8/typedapi/migration/postfeatureupgrade"
	ml_clear_trained_model_deployment_cache "github.com/elastic/go-elasticsearch/v8/typedapi/ml/cleartrainedmodeldeploymentcache"
	ml_close_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/closejob"
	ml_delete_calendar "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletecalendar"
	ml_delete_calendar_event "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletecalendarevent"
	ml_delete_calendar_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletecalendarjob"
	ml_delete_datafeed "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletedatafeed"
	ml_delete_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletedataframeanalytics"
	ml_delete_expired_data "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deleteexpireddata"
	ml_delete_filter "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletefilter"
	ml_delete_forecast "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deleteforecast"
	ml_delete_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletejob"
	ml_delete_model_snapshot "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletemodelsnapshot"
	ml_delete_trained_model "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletetrainedmodel"
	ml_delete_trained_model_alias "github.com/elastic/go-elasticsearch/v8/typedapi/ml/deletetrainedmodelalias"
	ml_estimate_model_memory "github.com/elastic/go-elasticsearch/v8/typedapi/ml/estimatemodelmemory"
	ml_evaluate_data_frame "github.com/elastic/go-elasticsearch/v8/typedapi/ml/evaluatedataframe"
	ml_explain_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/ml/explaindataframeanalytics"
	ml_flush_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/flushjob"
	ml_forecast "github.com/elastic/go-elasticsearch/v8/typedapi/ml/forecast"
	ml_get_buckets "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getbuckets"
	ml_get_calendar_events "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getcalendarevents"
	ml_get_calendars "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getcalendars"
	ml_get_categories "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getcategories"
	ml_get_datafeeds "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getdatafeeds"
	ml_get_datafeed_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getdatafeedstats"
	ml_get_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getdataframeanalytics"
	ml_get_data_frame_analytics_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getdataframeanalyticsstats"
	ml_get_filters "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getfilters"
	ml_get_influencers "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getinfluencers"
	ml_get_jobs "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getjobs"
	ml_get_job_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getjobstats"
	ml_get_memory_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getmemorystats"
	ml_get_model_snapshots "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getmodelsnapshots"
	ml_get_model_snapshot_upgrade_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getmodelsnapshotupgradestats"
	ml_get_overall_buckets "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getoverallbuckets"
	ml_get_records "github.com/elastic/go-elasticsearch/v8/typedapi/ml/getrecords"
	ml_get_trained_models "github.com/elastic/go-elasticsearch/v8/typedapi/ml/gettrainedmodels"
	ml_get_trained_models_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ml/gettrainedmodelsstats"
	ml_infer_trained_model "github.com/elastic/go-elasticsearch/v8/typedapi/ml/infertrainedmodel"
	ml_info "github.com/elastic/go-elasticsearch/v8/typedapi/ml/info"
	ml_open_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/openjob"
	ml_post_calendar_events "github.com/elastic/go-elasticsearch/v8/typedapi/ml/postcalendarevents"
	ml_post_data "github.com/elastic/go-elasticsearch/v8/typedapi/ml/postdata"
	ml_preview_datafeed "github.com/elastic/go-elasticsearch/v8/typedapi/ml/previewdatafeed"
	ml_preview_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/ml/previewdataframeanalytics"
	ml_put_calendar "github.com/elastic/go-elasticsearch/v8/typedapi/ml/putcalendar"
	ml_put_calendar_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/putcalendarjob"
	ml_put_datafeed "github.com/elastic/go-elasticsearch/v8/typedapi/ml/putdatafeed"
	ml_put_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/ml/putdataframeanalytics"
	ml_put_filter "github.com/elastic/go-elasticsearch/v8/typedapi/ml/putfilter"
	ml_put_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/putjob"
	ml_put_trained_model "github.com/elastic/go-elasticsearch/v8/typedapi/ml/puttrainedmodel"
	ml_put_trained_model_alias "github.com/elastic/go-elasticsearch/v8/typedapi/ml/puttrainedmodelalias"
	ml_put_trained_model_definition_part "github.com/elastic/go-elasticsearch/v8/typedapi/ml/puttrainedmodeldefinitionpart"
	ml_put_trained_model_vocabulary "github.com/elastic/go-elasticsearch/v8/typedapi/ml/puttrainedmodelvocabulary"
	ml_reset_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/resetjob"
	ml_revert_model_snapshot "github.com/elastic/go-elasticsearch/v8/typedapi/ml/revertmodelsnapshot"
	ml_set_upgrade_mode "github.com/elastic/go-elasticsearch/v8/typedapi/ml/setupgrademode"
	ml_start_datafeed "github.com/elastic/go-elasticsearch/v8/typedapi/ml/startdatafeed"
	ml_start_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/ml/startdataframeanalytics"
	ml_start_trained_model_deployment "github.com/elastic/go-elasticsearch/v8/typedapi/ml/starttrainedmodeldeployment"
	ml_stop_datafeed "github.com/elastic/go-elasticsearch/v8/typedapi/ml/stopdatafeed"
	ml_stop_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/ml/stopdataframeanalytics"
	ml_stop_trained_model_deployment "github.com/elastic/go-elasticsearch/v8/typedapi/ml/stoptrainedmodeldeployment"
	ml_update_datafeed "github.com/elastic/go-elasticsearch/v8/typedapi/ml/updatedatafeed"
	ml_update_data_frame_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/ml/updatedataframeanalytics"
	ml_update_filter "github.com/elastic/go-elasticsearch/v8/typedapi/ml/updatefilter"
	ml_update_job "github.com/elastic/go-elasticsearch/v8/typedapi/ml/updatejob"
	ml_update_model_snapshot "github.com/elastic/go-elasticsearch/v8/typedapi/ml/updatemodelsnapshot"
	ml_update_trained_model_deployment "github.com/elastic/go-elasticsearch/v8/typedapi/ml/updatetrainedmodeldeployment"
	ml_upgrade_job_snapshot "github.com/elastic/go-elasticsearch/v8/typedapi/ml/upgradejobsnapshot"
	ml_validate "github.com/elastic/go-elasticsearch/v8/typedapi/ml/validate"
	ml_validate_detector "github.com/elastic/go-elasticsearch/v8/typedapi/ml/validatedetector"
	monitoring_bulk "github.com/elastic/go-elasticsearch/v8/typedapi/monitoring/bulk"
	nodes_clear_repositories_metering_archive "github.com/elastic/go-elasticsearch/v8/typedapi/nodes/clearrepositoriesmeteringarchive"
	nodes_get_repositories_metering_info "github.com/elastic/go-elasticsearch/v8/typedapi/nodes/getrepositoriesmeteringinfo"
	nodes_hot_threads "github.com/elastic/go-elasticsearch/v8/typedapi/nodes/hotthreads"
	nodes_info "github.com/elastic/go-elasticsearch/v8/typedapi/nodes/info"
	nodes_reload_secure_settings "github.com/elastic/go-elasticsearch/v8/typedapi/nodes/reloadsecuresettings"
	nodes_stats "github.com/elastic/go-elasticsearch/v8/typedapi/nodes/stats"
	nodes_usage "github.com/elastic/go-elasticsearch/v8/typedapi/nodes/usage"
	profiling_flamegraph "github.com/elastic/go-elasticsearch/v8/typedapi/profiling/flamegraph"
	profiling_stacktraces "github.com/elastic/go-elasticsearch/v8/typedapi/profiling/stacktraces"
	profiling_status "github.com/elastic/go-elasticsearch/v8/typedapi/profiling/status"
	profiling_topn_functions "github.com/elastic/go-elasticsearch/v8/typedapi/profiling/topnfunctions"
	query_rules_delete_rule "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/deleterule"
	query_rules_delete_ruleset "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/deleteruleset"
	query_rules_get_rule "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/getrule"
	query_rules_get_ruleset "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/getruleset"
	query_rules_list_rulesets "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/listrulesets"
	query_rules_put_rule "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/putrule"
	query_rules_put_ruleset "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/putruleset"
	rollup_delete_job "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/deletejob"
	rollup_get_jobs "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/getjobs"
	rollup_get_rollup_caps "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/getrollupcaps"
	rollup_get_rollup_index_caps "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/getrollupindexcaps"
	rollup_put_job "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/putjob"
	rollup_rollup_search "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/rollupsearch"
	rollup_start_job "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/startjob"
	rollup_stop_job "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/stopjob"
	searchable_snapshots_cache_stats "github.com/elastic/go-elasticsearch/v8/typedapi/searchablesnapshots/cachestats"
	searchable_snapshots_clear_cache "github.com/elastic/go-elasticsearch/v8/typedapi/searchablesnapshots/clearcache"
	searchable_snapshots_mount "github.com/elastic/go-elasticsearch/v8/typedapi/searchablesnapshots/mount"
	searchable_snapshots_stats "github.com/elastic/go-elasticsearch/v8/typedapi/searchablesnapshots/stats"
	search_application_delete "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/delete"
	search_application_delete_behavioral_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/deletebehavioralanalytics"
	search_application_get "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/get"
	search_application_get_behavioral_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/getbehavioralanalytics"
	search_application_list "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/list"
	search_application_put "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/put"
	search_application_put_behavioral_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/putbehavioralanalytics"
	search_application_search "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/search"
	security_activate_user_profile "github.com/elastic/go-elasticsearch/v8/typedapi/security/activateuserprofile"
	security_authenticate "github.com/elastic/go-elasticsearch/v8/typedapi/security/authenticate"
	security_bulk_delete_role "github.com/elastic/go-elasticsearch/v8/typedapi/security/bulkdeleterole"
	security_bulk_put_role "github.com/elastic/go-elasticsearch/v8/typedapi/security/bulkputrole"
	security_bulk_update_api_keys "github.com/elastic/go-elasticsearch/v8/typedapi/security/bulkupdateapikeys"
	security_change_password "github.com/elastic/go-elasticsearch/v8/typedapi/security/changepassword"
	security_clear_api_key_cache "github.com/elastic/go-elasticsearch/v8/typedapi/security/clearapikeycache"
	security_clear_cached_privileges "github.com/elastic/go-elasticsearch/v8/typedapi/security/clearcachedprivileges"
	security_clear_cached_realms "github.com/elastic/go-elasticsearch/v8/typedapi/security/clearcachedrealms"
	security_clear_cached_roles "github.com/elastic/go-elasticsearch/v8/typedapi/security/clearcachedroles"
	security_clear_cached_service_tokens "github.com/elastic/go-elasticsearch/v8/typedapi/security/clearcachedservicetokens"
	security_create_api_key "github.com/elastic/go-elasticsearch/v8/typedapi/security/createapikey"
	security_create_cross_cluster_api_key "github.com/elastic/go-elasticsearch/v8/typedapi/security/createcrossclusterapikey"
	security_create_service_token "github.com/elastic/go-elasticsearch/v8/typedapi/security/createservicetoken"
	security_delete_privileges "github.com/elastic/go-elasticsearch/v8/typedapi/security/deleteprivileges"
	security_delete_role "github.com/elastic/go-elasticsearch/v8/typedapi/security/deleterole"
	security_delete_role_mapping "github.com/elastic/go-elasticsearch/v8/typedapi/security/deleterolemapping"
	security_delete_service_token "github.com/elastic/go-elasticsearch/v8/typedapi/security/deleteservicetoken"
	security_delete_user "github.com/elastic/go-elasticsearch/v8/typedapi/security/deleteuser"
	security_disable_user "github.com/elastic/go-elasticsearch/v8/typedapi/security/disableuser"
	security_disable_user_profile "github.com/elastic/go-elasticsearch/v8/typedapi/security/disableuserprofile"
	security_enable_user "github.com/elastic/go-elasticsearch/v8/typedapi/security/enableuser"
	security_enable_user_profile "github.com/elastic/go-elasticsearch/v8/typedapi/security/enableuserprofile"
	security_enroll_kibana "github.com/elastic/go-elasticsearch/v8/typedapi/security/enrollkibana"
	security_enroll_node "github.com/elastic/go-elasticsearch/v8/typedapi/security/enrollnode"
	security_get_api_key "github.com/elastic/go-elasticsearch/v8/typedapi/security/getapikey"
	security_get_builtin_privileges "github.com/elastic/go-elasticsearch/v8/typedapi/security/getbuiltinprivileges"
	security_get_privileges "github.com/elastic/go-elasticsearch/v8/typedapi/security/getprivileges"
	security_get_role "github.com/elastic/go-elasticsearch/v8/typedapi/security/getrole"
	security_get_role_mapping "github.com/elastic/go-elasticsearch/v8/typedapi/security/getrolemapping"
	security_get_service_accounts "github.com/elastic/go-elasticsearch/v8/typedapi/security/getserviceaccounts"
	security_get_service_credentials "github.com/elastic/go-elasticsearch/v8/typedapi/security/getservicecredentials"
	security_get_settings "github.com/elastic/go-elasticsearch/v8/typedapi/security/getsettings"
	security_get_token "github.com/elastic/go-elasticsearch/v8/typedapi/security/gettoken"
	security_get_user "github.com/elastic/go-elasticsearch/v8/typedapi/security/getuser"
	security_get_user_privileges "github.com/elastic/go-elasticsearch/v8/typedapi/security/getuserprivileges"
	security_get_user_profile "github.com/elastic/go-elasticsearch/v8/typedapi/security/getuserprofile"
	security_grant_api_key "github.com/elastic/go-elasticsearch/v8/typedapi/security/grantapikey"
	security_has_privileges "github.com/elastic/go-elasticsearch/v8/typedapi/security/hasprivileges"
	security_has_privileges_user_profile "github.com/elastic/go-elasticsearch/v8/typedapi/security/hasprivilegesuserprofile"
	security_invalidate_api_key "github.com/elastic/go-elasticsearch/v8/typedapi/security/invalidateapikey"
	security_invalidate_token "github.com/elastic/go-elasticsearch/v8/typedapi/security/invalidatetoken"
	security_oidc_authenticate "github.com/elastic/go-elasticsearch/v8/typedapi/security/oidcauthenticate"
	security_oidc_logout "github.com/elastic/go-elasticsearch/v8/typedapi/security/oidclogout"
	security_oidc_prepare_authentication "github.com/elastic/go-elasticsearch/v8/typedapi/security/oidcprepareauthentication"
	security_put_privileges "github.com/elastic/go-elasticsearch/v8/typedapi/security/putprivileges"
	security_put_role "github.com/elastic/go-elasticsearch/v8/typedapi/security/putrole"
	security_put_role_mapping "github.com/elastic/go-elasticsearch/v8/typedapi/security/putrolemapping"
	security_put_user "github.com/elastic/go-elasticsearch/v8/typedapi/security/putuser"
	security_query_api_keys "github.com/elastic/go-elasticsearch/v8/typedapi/security/queryapikeys"
	security_query_role "github.com/elastic/go-elasticsearch/v8/typedapi/security/queryrole"
	security_query_user "github.com/elastic/go-elasticsearch/v8/typedapi/security/queryuser"
	security_saml_authenticate "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlauthenticate"
	security_saml_complete_logout "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlcompletelogout"
	security_saml_invalidate "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlinvalidate"
	security_saml_logout "github.com/elastic/go-elasticsearch/v8/typedapi/security/samllogout"
	security_saml_prepare_authentication "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlprepareauthentication"
	security_saml_service_provider_metadata "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlserviceprovidermetadata"
	security_suggest_user_profiles "github.com/elastic/go-elasticsearch/v8/typedapi/security/suggestuserprofiles"
	security_update_api_key "github.com/elastic/go-elasticsearch/v8/typedapi/security/updateapikey"
	security_update_settings "github.com/elastic/go-elasticsearch/v8/typedapi/security/updatesettings"
	security_update_user_profile_data "github.com/elastic/go-elasticsearch/v8/typedapi/security/updateuserprofiledata"
	shutdown_delete_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/deletenode"
	shutdown_get_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/getnode"
	shutdown_put_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/putnode"
	slm_delete_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/slm/deletelifecycle"
	slm_execute_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/slm/executelifecycle"
	slm_execute_retention "github.com/elastic/go-elasticsearch/v8/typedapi/slm/executeretention"
	slm_get_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/slm/getlifecycle"
	slm_get_stats "github.com/elastic/go-elasticsearch/v8/typedapi/slm/getstats"
	slm_get_status "github.com/elastic/go-elasticsearch/v8/typedapi/slm/getstatus"
	slm_put_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/slm/putlifecycle"
	slm_start "github.com/elastic/go-elasticsearch/v8/typedapi/slm/start"
	slm_stop "github.com/elastic/go-elasticsearch/v8/typedapi/slm/stop"
	snapshot_cleanup_repository "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/cleanuprepository"
	snapshot_clone "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/clone"
	snapshot_create "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/create"
	snapshot_create_repository "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/createrepository"
	snapshot_delete "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/delete"
	snapshot_delete_repository "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/deleterepository"
	snapshot_get "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/get"
	snapshot_get_repository "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/getrepository"
	snapshot_restore "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/restore"
	snapshot_status "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/status"
	snapshot_verify_repository "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/verifyrepository"
	sql_clear_cursor "github.com/elastic/go-elasticsearch/v8/typedapi/sql/clearcursor"
	sql_delete_async "github.com/elastic/go-elasticsearch/v8/typedapi/sql/deleteasync"
	sql_get_async "github.com/elastic/go-elasticsearch/v8/typedapi/sql/getasync"
	sql_get_async_status "github.com/elastic/go-elasticsearch/v8/typedapi/sql/getasyncstatus"
	sql_query "github.com/elastic/go-elasticsearch/v8/typedapi/sql/query"
	sql_translate "github.com/elastic/go-elasticsearch/v8/typedapi/sql/translate"
	ssl_certificates "github.com/elastic/go-elasticsearch/v8/typedapi/ssl/certificates"
	synonyms_delete_synonym "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/deletesynonym"
	synonyms_delete_synonym_rule "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/deletesynonymrule"
	synonyms_get_synonym "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/getsynonym"
	synonyms_get_synonym_rule "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/getsynonymrule"
	synonyms_get_synonyms_sets "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/getsynonymssets"
	synonyms_put_synonym "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/putsynonym"
	synonyms_put_synonym_rule "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/putsynonymrule"
	tasks_cancel "github.com/elastic/go-elasticsearch/v8/typedapi/tasks/cancel"
	tasks_get "github.com/elastic/go-elasticsearch/v8/typedapi/tasks/get"
	tasks_list "github.com/elastic/go-elasticsearch/v8/typedapi/tasks/list"
	text_structure_find_field_structure "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/findfieldstructure"
	text_structure_find_message_structure "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/findmessagestructure"
	text_structure_find_structure "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/findstructure"
	text_structure_test_grok_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/testgrokpattern"
	transform_delete_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/deletetransform"
	transform_get_node_stats "github.com/elastic/go-elasticsearch/v8/typedapi/transform/getnodestats"
	transform_get_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/gettransform"
	transform_get_transform_stats "github.com/elastic/go-elasticsearch/v8/typedapi/transform/gettransformstats"
	transform_preview_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/previewtransform"
	transform_put_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/puttransform"
	transform_reset_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/resettransform"
	transform_schedule_now_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/schedulenowtransform"
	transform_start_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/starttransform"
	transform_stop_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/stoptransform"
	transform_update_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/updatetransform"
	transform_upgrade_transforms "github.com/elastic/go-elasticsearch/v8/typedapi/transform/upgradetransforms"
	watcher_ack_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/ackwatch"
	watcher_activate_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/activatewatch"
	watcher_deactivate_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/deactivatewatch"
	watcher_delete_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/deletewatch"
	watcher_execute_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/executewatch"
	watcher_get_settings "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/getsettings"
	watcher_get_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/getwatch"
	watcher_put_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/putwatch"
	watcher_query_watches "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/querywatches"
	watcher_start "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/start"
	watcher_stats "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/stats"
	watcher_stop "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/stop"
	watcher_update_settings "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/updatesettings"
	xpack_info "github.com/elastic/go-elasticsearch/v8/typedapi/xpack/info"
	xpack_usage "github.com/elastic/go-elasticsearch/v8/typedapi/xpack/usage"
)

type AsyncSearch struct {
	// Deletes an async search by identifier.
	// If the search is still running, the search request will be cancelled.
	// Otherwise, the saved search results are deleted.
	// If the Elasticsearch security features are enabled, the deletion of a
	// specific async search is restricted to: the authenticated user that submitted
	// the original search request; users that have the `cancel_task` cluster
	// privilege.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Delete async_search_delete.NewDelete
	// Retrieves the results of a previously submitted async search request given
	// its identifier.
	// If the Elasticsearch security features are enabled, access to the results of
	// a specific async search is restricted to the user or API key that submitted
	// it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Get async_search_get.NewGet
	// Get async search status
	// Retrieves the status of a previously submitted async search request given its
	// identifier, without retrieving search results.
	// If the Elasticsearch security features are enabled, use of this API is
	// restricted to the `monitoring_user` role.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Status async_search_status.NewStatus
	// Runs a search request asynchronously.
	// When the primary sort of the results is an indexed field, shards get sorted
	// based on minimum and maximum value that they hold for that field, hence
	// partial results become available following the sort criteria that was
	// requested.
	// Warning: Async search does not support scroll nor search requests that only
	// include the suggest section.
	// By default, Elasticsearch doesn’t allow you to store an async search response
	// larger than 10Mb and an attempt to do this results in an error.
	// The maximum allowed size for a stored async search response can be set by
	// changing the `search.max_async_search_response_size` cluster level setting.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Submit async_search_submit.NewSubmit
}

type Autoscaling struct {
	// Deletes an autoscaling policy. Designed for indirect use by ECE/ESS and ECK.
	// Direct use is not supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-delete-autoscaling-policy.html
	DeleteAutoscalingPolicy autoscaling_delete_autoscaling_policy.NewDeleteAutoscalingPolicy
	// Gets the current autoscaling capacity based on the configured autoscaling
	// policy. Designed for indirect use by ECE/ESS and ECK. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-get-autoscaling-capacity.html
	GetAutoscalingCapacity autoscaling_get_autoscaling_capacity.NewGetAutoscalingCapacity
	// Retrieves an autoscaling policy. Designed for indirect use by ECE/ESS and
	// ECK. Direct use is not supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-get-autoscaling-capacity.html
	GetAutoscalingPolicy autoscaling_get_autoscaling_policy.NewGetAutoscalingPolicy
	// Creates a new autoscaling policy. Designed for indirect use by ECE/ESS and
	// ECK. Direct use is not supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-put-autoscaling-policy.html
	PutAutoscalingPolicy autoscaling_put_autoscaling_policy.NewPutAutoscalingPolicy
}

type Capabilities struct {
	// Checks if the specified combination of method, API, parameters, and arbitrary
	// capabilities are supported
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/capabilities.html
	Capabilities capabilities.NewCapabilities
}

type Cat struct {
	// Retrieves the cluster’s index aliases, including filter and routing
	// information.
	// The API does not return data stream aliases.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or the Kibana console. They are not intended for use by applications.
	// For application consumption, use the aliases API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-alias.html
	Aliases cat_aliases.NewAliases
	// Provides a snapshot of the number of shards allocated to each data node and
	// their disk space.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-allocation.html
	Allocation cat_allocation.NewAllocation
	// Returns information about component templates in a cluster.
	// Component templates are building blocks for constructing index templates that
	// specify index mappings, settings, and aliases.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the get component template API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-component-templates.html
	ComponentTemplates cat_component_templates.NewComponentTemplates
	// Provides quick access to a document count for a data stream, an index, or an
	// entire cluster.
	// NOTE: The document count only includes live documents, not deleted documents
	// which have not yet been removed by the merge process.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the count API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-count.html
	Count cat_count.NewCount
	// Returns the amount of heap memory currently used by the field data cache on
	// every data node in the cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the nodes stats API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-fielddata.html
	Fielddata cat_fielddata.NewFielddata
	// Returns the health status of a cluster, similar to the cluster health API.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the cluster health API.
	// This API is often used to check malfunctioning clusters.
	// To help you track cluster health alongside log files and alerting systems,
	// the API returns timestamps in two formats:
	// `HH:MM:SS`, which is human-readable but includes no date information;
	// `Unix epoch time`, which is machine-sortable and includes date information.
	// The latter format is useful for cluster recoveries that take multiple days.
	// You can use the cat health API to verify cluster health across multiple
	// nodes.
	// You also can use the API to track the recovery of a large cluster over a
	// longer period of time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-health.html
	Health cat_health.NewHealth
	// Returns help for the Cat APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat.html
	Help cat_help.NewHelp
	// Returns high-level information about indices in a cluster, including backing
	// indices for data streams.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the get index API.
	// Use the cat indices API to get the following information for each index in a
	// cluster: shard count; document count; deleted document count; primary store
	// size; total store size of all shards, including shard replicas.
	// These metrics are retrieved directly from Lucene, which Elasticsearch uses
	// internally to power indexing and search. As a result, all document counts
	// include hidden nested documents.
	// To get an accurate count of Elasticsearch documents, use the cat count or
	// count APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html
	Indices cat_indices.NewIndices
	// Returns information about the master node, including the ID, bound IP
	// address, and name.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-master.html
	Master cat_master.NewMaster
	// Returns configuration and usage information about data frame analytics jobs.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get data frame analytics jobs statistics
	// API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-dfanalytics.html
	MlDataFrameAnalytics cat_ml_data_frame_analytics.NewMlDataFrameAnalytics
	// Returns configuration and usage information about datafeeds.
	// This API returns a maximum of 10,000 datafeeds.
	// If the Elasticsearch security features are enabled, you must have
	// `monitor_ml`, `monitor`, `manage_ml`, or `manage`
	// cluster privileges to use this API.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get datafeed statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-datafeeds.html
	MlDatafeeds cat_ml_datafeeds.NewMlDatafeeds
	// Returns configuration and usage information for anomaly detection jobs.
	// This API returns a maximum of 10,000 jobs.
	// If the Elasticsearch security features are enabled, you must have
	// `monitor_ml`,
	// `monitor`, `manage_ml`, or `manage` cluster privileges to use this API.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get anomaly detection job statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-anomaly-detectors.html
	MlJobs cat_ml_jobs.NewMlJobs
	// Returns configuration and usage information about inference trained models.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get trained models statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-trained-model.html
	MlTrainedModels cat_ml_trained_models.NewMlTrainedModels
	// Returns information about custom node attributes.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodeattrs.html
	Nodeattrs cat_nodeattrs.NewNodeattrs
	// Returns information about the nodes in a cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodes.html
	Nodes cat_nodes.NewNodes
	// Returns cluster-level changes that have not yet been executed.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the pending cluster tasks API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-pending-tasks.html
	PendingTasks cat_pending_tasks.NewPendingTasks
	// Returns a list of plugins running on each node of a cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-plugins.html
	Plugins cat_plugins.NewPlugins
	// Returns information about ongoing and completed shard recoveries.
	// Shard recovery is the process of initializing a shard copy, such as restoring
	// a primary shard from a snapshot or syncing a replica shard from a primary
	// shard. When a shard recovery completes, the recovered shard is available for
	// search and indexing.
	// For data streams, the API returns information about the stream’s backing
	// indices.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the index recovery API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-recovery.html
	Recovery cat_recovery.NewRecovery
	// Returns the snapshot repositories for a cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the get snapshot repository API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-repositories.html
	Repositories cat_repositories.NewRepositories
	// Returns low-level information about the Lucene segments in index shards.
	// For data streams, the API returns information about the backing indices.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the index segments API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-segments.html
	Segments cat_segments.NewSegments
	// Returns information about the shards in a cluster.
	// For data streams, the API returns information about the backing indices.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-shards.html
	Shards cat_shards.NewShards
	// Returns information about the snapshots stored in one or more repositories.
	// A snapshot is a backup of an index or running Elasticsearch cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the get snapshot API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-snapshots.html
	Snapshots cat_snapshots.NewSnapshots
	// Returns information about tasks currently executing in the cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the task management API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	Tasks cat_tasks.NewTasks
	// Returns information about index templates in a cluster.
	// You can use index templates to apply index settings and field mappings to new
	// indices at creation.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the get index template API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-templates.html
	Templates cat_templates.NewTemplates
	// Returns thread pool statistics for each node in a cluster.
	// Returned information includes all built-in thread pools and custom thread
	// pools.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-thread-pool.html
	ThreadPool cat_thread_pool.NewThreadPool
	// Returns configuration and usage information about transforms.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get transform statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-transforms.html
	Transforms cat_transforms.NewTransforms
}

type Ccr struct {
	// Deletes auto-follow patterns.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-delete-auto-follow-pattern.html
	DeleteAutoFollowPattern ccr_delete_auto_follow_pattern.NewDeleteAutoFollowPattern
	// Creates a new follower index configured to follow the referenced leader
	// index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-follow.html
	Follow ccr_follow.NewFollow
	// Retrieves information about all follower indices, including parameters and
	// status for each follower index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-follow-info.html
	FollowInfo ccr_follow_info.NewFollowInfo
	// Retrieves follower stats. return shard-level stats about the following tasks
	// associated with each shard for the specified indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-follow-stats.html
	FollowStats ccr_follow_stats.NewFollowStats
	// Removes the follower retention leases from the leader.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-forget-follower.html
	ForgetFollower ccr_forget_follower.NewForgetFollower
	// Gets configured auto-follow patterns. Returns the specified auto-follow
	// pattern collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-auto-follow-pattern.html
	GetAutoFollowPattern ccr_get_auto_follow_pattern.NewGetAutoFollowPattern
	// Pauses an auto-follow pattern
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-pause-auto-follow-pattern.html
	PauseAutoFollowPattern ccr_pause_auto_follow_pattern.NewPauseAutoFollowPattern
	// Pauses a follower index. The follower index will not fetch any additional
	// operations from the leader index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-pause-follow.html
	PauseFollow ccr_pause_follow.NewPauseFollow
	// Creates a new named collection of auto-follow patterns against a specified
	// remote cluster. Newly created indices on the remote cluster matching any of
	// the specified patterns will be automatically configured as follower indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-auto-follow-pattern.html
	PutAutoFollowPattern ccr_put_auto_follow_pattern.NewPutAutoFollowPattern
	// Resumes an auto-follow pattern that has been paused
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-resume-auto-follow-pattern.html
	ResumeAutoFollowPattern ccr_resume_auto_follow_pattern.NewResumeAutoFollowPattern
	// Resumes a follower index that has been paused
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-resume-follow.html
	ResumeFollow ccr_resume_follow.NewResumeFollow
	// Gets all stats related to cross-cluster replication.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-stats.html
	Stats ccr_stats.NewStats
	// Stops the following task associated with a follower index and removes index
	// metadata and settings associated with cross-cluster replication.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-unfollow.html
	Unfollow ccr_unfollow.NewUnfollow
}

type Cluster struct {
	// Provides explanations for shard allocations in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-allocation-explain.html
	AllocationExplain cluster_allocation_explain.NewAllocationExplain
	// Deletes component templates.
	// Component templates are building blocks for constructing index templates that
	// specify index mappings, settings, and aliases.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	DeleteComponentTemplate cluster_delete_component_template.NewDeleteComponentTemplate
	// Clears cluster voting config exclusions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/voting-config-exclusions.html
	DeleteVotingConfigExclusions cluster_delete_voting_config_exclusions.NewDeleteVotingConfigExclusions
	// Returns information about whether a particular component template exist
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	ExistsComponentTemplate cluster_exists_component_template.NewExistsComponentTemplate
	// Retrieves information about component templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	GetComponentTemplate cluster_get_component_template.NewGetComponentTemplate
	// Returns cluster-wide settings.
	// By default, it returns only settings that have been explicitly defined.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-get-settings.html
	GetSettings cluster_get_settings.NewGetSettings
	// The cluster health API returns a simple status on the health of the cluster.
	// You can also use the API to get the health status of only specified data
	// streams and indices. For data streams, the API retrieves the health status of
	// the stream’s backing indices.
	// The cluster health status is: green, yellow or red. On the shard level, a red
	// status indicates that the specific shard is not allocated in the cluster,
	// yellow means that the primary shard is allocated but replicas are not, and
	// green means that all shards are allocated. The index level status is
	// controlled by the worst shard status. The cluster status is controlled by the
	// worst index status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html
	Health cluster_health.NewHealth
	// Returns different information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-info.html
	Info cluster_info.NewInfo
	// Returns cluster-level changes (such as create index, update mapping, allocate
	// or fail shard) that have not yet been executed.
	// NOTE: This API returns a list of any pending updates to the cluster state.
	// These are distinct from the tasks reported by the Task Management API which
	// include periodic tasks and tasks initiated by the user, such as node stats,
	// search queries, or create index requests.
	// However, if a user-initiated task such as a create index command causes a
	// cluster state update, the activity of this task might be reported by both
	// task api and pending cluster tasks API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-pending.html
	PendingTasks cluster_pending_tasks.NewPendingTasks
	// Updates the cluster voting config exclusions by node ids or node names.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/voting-config-exclusions.html
	PostVotingConfigExclusions cluster_post_voting_config_exclusions.NewPostVotingConfigExclusions
	// Creates or updates a component template.
	// Component templates are building blocks for constructing index templates that
	// specify index mappings, settings, and aliases.
	//
	// An index template can be composed of multiple component templates.
	// To use a component template, specify it in an index template’s `composed_of`
	// list.
	// Component templates are only applied to new data streams and indices as part
	// of a matching index template.
	//
	// Settings and mappings specified directly in the index template or the create
	// index request override any settings or mappings specified in a component
	// template.
	//
	// Component templates are only used during index creation.
	// For data streams, this includes data stream creation and the creation of a
	// stream’s backing indices.
	// Changes to component templates do not affect existing indices, including a
	// stream’s backing indices.
	//
	// You can use C-style `/* *\/` block comments in component templates.
	// You can include comments anywhere in the request body except before the
	// opening curly bracket.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	PutComponentTemplate cluster_put_component_template.NewPutComponentTemplate
	// Updates the cluster settings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-update-settings.html
	PutSettings cluster_put_settings.NewPutSettings
	// The cluster remote info API allows you to retrieve all of the configured
	// remote cluster information. It returns connection and endpoint information
	// keyed by the configured remote cluster alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-remote-info.html
	RemoteInfo cluster_remote_info.NewRemoteInfo
	// Allows to manually change the allocation of individual shards in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-reroute.html
	Reroute cluster_reroute.NewReroute
	// Returns a comprehensive information about the state of the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-state.html
	State cluster_state.NewState
	// Returns cluster statistics.
	// It returns basic index metrics (shard numbers, store size, memory usage) and
	// information about the current nodes that form the cluster (number, roles, os,
	// jvm versions, memory usage, cpu and installed plugins).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-stats.html
	Stats cluster_stats.NewStats
}

type Connector struct {
	// Updates the last_seen field in the connector, and sets it to current
	// timestamp
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/check-in-connector-api.html
	CheckIn connector_check_in.NewCheckIn
	// Deletes a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-connector-api.html
	Delete connector_delete.NewDelete
	// Retrieves a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-connector-api.html
	Get connector_get.NewGet
	// Updates last sync stats in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-last-sync-api.html
	LastSync connector_last_sync.NewLastSync
	// Returns existing connectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-connector-api.html
	List connector_list.NewList
	// Creates a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-api.html
	Post connector_post.NewPost
	// Creates or updates a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-api.html
	Put connector_put.NewPut
	// Creates a secret for a Connector.
	//
	SecretPost connector_secret_post.NewSecretPost
	// Cancels a connector sync job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cancel-connector-sync-job-api.html
	SyncJobCancel connector_sync_job_cancel.NewSyncJobCancel
	// Deletes a connector sync job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-connector-sync-job-api.html
	SyncJobDelete connector_sync_job_delete.NewSyncJobDelete
	// Retrieves a connector sync job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-connector-sync-job-api.html
	SyncJobGet connector_sync_job_get.NewSyncJobGet
	// Lists connector sync jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-connector-sync-jobs-api.html
	SyncJobList connector_sync_job_list.NewSyncJobList
	// Creates a connector sync job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-sync-job-api.html
	SyncJobPost connector_sync_job_post.NewSyncJobPost
	// Activates the valid draft filtering for a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-api.html
	UpdateActiveFiltering connector_update_active_filtering.NewUpdateActiveFiltering
	// Updates the API key id in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-api-key-id-api.html
	UpdateApiKeyId connector_update_api_key_id.NewUpdateApiKeyId
	// Updates the configuration field in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-configuration-api.html
	UpdateConfiguration connector_update_configuration.NewUpdateConfiguration
	// Updates the filtering field in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-error-api.html
	UpdateError connector_update_error.NewUpdateError
	// Updates the filtering field in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-api.html
	UpdateFiltering connector_update_filtering.NewUpdateFiltering
	// Updates the draft filtering validation info for a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-validation-api.html
	UpdateFilteringValidation connector_update_filtering_validation.NewUpdateFilteringValidation
	// Updates the index_name in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-index-name-api.html
	UpdateIndexName connector_update_index_name.NewUpdateIndexName
	// Updates the name and description fields in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-name-description-api.html
	UpdateName connector_update_name.NewUpdateName
	// Updates the is_native flag in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-native-api.html
	UpdateNative connector_update_native.NewUpdateNative
	// Updates the pipeline field in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-pipeline-api.html
	UpdatePipeline connector_update_pipeline.NewUpdatePipeline
	// Updates the scheduling field in the connector document
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-scheduling-api.html
	UpdateScheduling connector_update_scheduling.NewUpdateScheduling
	// Updates the service type of the connector
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-service-type-api.html
	UpdateServiceType connector_update_service_type.NewUpdateServiceType
	// Updates the status of the connector
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-status-api.html
	UpdateStatus connector_update_status.NewUpdateStatus
}

type Core struct {
	// Performs multiple indexing or delete operations in a single API call.
	// This reduces overhead and can greatly increase indexing speed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Clears the search context and results for a scrolling search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Closes a point-in-time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Removes a JSON document from the specified index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Deletes documents that match the specified query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Changes the number of requests per second for a particular Delete By Query
	// operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	// Deletes a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Checks if a document in an index exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Checks if a document's `_source` is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Returns information about why a specific document matches (or doesn’t match)
	// a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// The field capabilities API returns the information about the capabilities of
	// fields among multiple indices.
	// The field capabilities API returns runtime fields like any other field. For
	// example, a runtime field with a type
	// of keyword is returned as any other field that belongs to the `keyword`
	// family.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Returns a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Retrieves a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScript core_get_script.NewGetScript
	// Returns all script contexts.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-contexts.html
	GetScriptContext core_get_script_context.NewGetScriptContext
	// Returns available script types, languages and contexts
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScriptLanguages core_get_script_languages.NewGetScriptLanguages
	// Returns the source of a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Returns the health of the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html
	HealthReport core_health_report.NewHealthReport
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Index core_index.NewIndex
	// Returns basic information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Info core_info.NewInfo
	// Performs a kNN search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	KnnSearch core_knn_search.NewKnnSearch
	// Allows to get multiple documents in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
	Mget core_mget.NewMget
	// Allows to execute several search operations in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	Msearch core_msearch.NewMsearch
	// Runs multiple templated searches with a single request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Returns multiple termvectors in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// A search request by default executes against the most recent visible data of
	// the target indices,
	// which is called point in time. Elasticsearch pit (point in time) is a
	// lightweight view into the
	// state of the data as it existed when initiated. In some cases, it’s preferred
	// to perform multiple
	// search requests using the same point in time. For example, if refreshes
	// happen between
	// `search_after` requests, then the results of those requests might not be
	// consistent as changes happening
	// between searches are only visible to the more recent point in time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Creates or updates a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Enables you to evaluate the quality of ranked search results over a set of
	// typical search queries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Allows to copy documents from one index to another, optionally filtering the
	// source
	// documents by a query, changing the destination index settings, or fetching
	// the
	// documents from a remote cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Copies documents from a source to a destination.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle
	// Renders a search template as a search request body.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Runs a script and returns a result.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Allows to retrieve a large numbers of results from a single search request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Returns search hits that match the query defined in the request.
	// You can provide search queries using the `q` query string parameter or the
	// request body.
	// If both are specified, only the query parameter is used.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Searches a vector tile for geospatial values. Returns results as a binary
	// Mapbox vector tile.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Returns information about the indices and shards that a search request would
	// be executed against.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-shards.html
	SearchShards core_search_shards.NewSearchShards
	// Runs a search with a search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
	SearchTemplate core_search_template.NewSearchTemplate
	// The terms enum API  can be used to discover terms in the index that begin
	// with the provided string. It is designed for low-latency look-ups used in
	// auto-complete scenarios.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
	TermsEnum core_terms_enum.NewTermsEnum
	// Returns information and statistics about terms in the fields of a particular
	// document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
	Termvectors core_termvectors.NewTermvectors
	// Updates a document with a script or partial document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
	Update core_update.NewUpdate
	// Updates documents that match the specified query.
	// If no query is specified, performs an update on every document in the data
	// stream or index without modifying the source, which is useful for picking up
	// mapping changes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQuery core_update_by_query.NewUpdateByQuery
	// Changes the number of requests per second for a particular Update By Query
	// operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

type DanglingIndices struct {
	// Deletes the specified dangling index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-gateway-dangling-indices.html
	DeleteDanglingIndex dangling_indices_delete_dangling_index.NewDeleteDanglingIndex
	// Imports the specified dangling index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-gateway-dangling-indices.html
	ImportDanglingIndex dangling_indices_import_dangling_index.NewImportDanglingIndex
	// Returns all dangling indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-gateway-dangling-indices.html
	ListDanglingIndices dangling_indices_list_dangling_indices.NewListDanglingIndices
}

type Enrich struct {
	// Deletes an existing enrich policy and its enrich index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-enrich-policy-api.html
	DeletePolicy enrich_delete_policy.NewDeletePolicy
	// Creates the enrich index for an existing enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/execute-enrich-policy-api.html
	ExecutePolicy enrich_execute_policy.NewExecutePolicy
	// Returns information about an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-enrich-policy-api.html
	GetPolicy enrich_get_policy.NewGetPolicy
	// Creates an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-enrich-policy-api.html
	PutPolicy enrich_put_policy.NewPutPolicy
	// Returns enrich coordinator statistics and information about enrich policies
	// that are currently executing.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-stats-api.html
	Stats enrich_stats.NewStats
}

type Eql struct {
	// Deletes an async EQL search or a stored synchronous EQL search.
	// The API also deletes results for the search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Delete eql_delete.NewDelete
	// Returns the current status and available results for an async EQL search or a
	// stored synchronous EQL search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-search-api.html
	Get eql_get.NewGet
	// Returns the current status for an async EQL search or a stored synchronous
	// EQL search without returning results.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-status-api.html
	GetStatus eql_get_status.NewGetStatus
	// Returns results matching a query expressed in Event Query Language (EQL)
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Search eql_search.NewSearch
}

type Esql struct {
	// Executes an ESQL request asynchronously
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-async-query-api.html
	AsyncQuery esql_async_query.NewAsyncQuery
	// Executes an ES|QL request
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-rest.html
	Query esql_query.NewQuery
}

type Features struct {
	// Gets a list of features which can be included in snapshots using the
	// feature_states field when creating a snapshot
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-features-api.html
	GetFeatures features_get_features.NewGetFeatures
	// Resets the internal state of features, usually by deleting system indices
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	ResetFeatures features_reset_features.NewResetFeatures
}

type Fleet struct {
	// Returns the current global checkpoints for an index. This API is design for
	// internal use by the fleet server project.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-global-checkpoints.html
	GlobalCheckpoints fleet_global_checkpoints.NewGlobalCheckpoints
	// Executes several [fleet
	// searches](https://www.elastic.co/guide/en/elasticsearch/reference/current/fleet-search.html)
	// with a single API request.
	// The API follows the same structure as the [multi
	// search](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html)
	// API. However, similar to the fleet search API, it
	// supports the wait_for_checkpoints parameter.
	//
	Msearch fleet_msearch.NewMsearch
	// Creates a secret stored by Fleet.
	//
	PostSecret fleet_post_secret.NewPostSecret
	// The purpose of the fleet search api is to provide a search api where the
	// search will only be executed
	// after provided checkpoint has been processed and is visible for searches
	// inside of Elasticsearch.
	//
	Search fleet_search.NewSearch
}

type Graph struct {
	// Extracts and summarizes information about the documents and terms in an
	// Elasticsearch data stream or index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/graph-explore-api.html
	Explore graph_explore.NewExplore
}

type Ilm struct {
	// Deletes the specified lifecycle policy definition. You cannot delete policies
	// that are currently in use. If the policy is being used to manage any indices,
	// the request fails and returns an error.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-delete-lifecycle.html
	DeleteLifecycle ilm_delete_lifecycle.NewDeleteLifecycle
	// Retrieves information about the index’s current lifecycle state, such as the
	// currently executing phase, action, and step. Shows when the index entered
	// each one, the definition of the running phase, and information about any
	// failures.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-explain-lifecycle.html
	ExplainLifecycle ilm_explain_lifecycle.NewExplainLifecycle
	// Retrieves a lifecycle policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-lifecycle.html
	GetLifecycle ilm_get_lifecycle.NewGetLifecycle
	// Retrieves the current index lifecycle management (ILM) status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-status.html
	GetStatus ilm_get_status.NewGetStatus
	// Switches the indices, ILM policies, and legacy, composable and component
	// templates from using custom node attributes and
	// attribute-based allocation filters to using data tiers, and optionally
	// deletes one legacy index template.+
	// Using node roles enables ILM to automatically move the indices between data
	// tiers.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-migrate-to-data-tiers.html
	MigrateToDataTiers ilm_migrate_to_data_tiers.NewMigrateToDataTiers
	// Manually moves an index into the specified step and executes that step.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-move-to-step.html
	MoveToStep ilm_move_to_step.NewMoveToStep
	// Creates a lifecycle policy. If the specified policy exists, the policy is
	// replaced and the policy version is incremented.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-put-lifecycle.html
	PutLifecycle ilm_put_lifecycle.NewPutLifecycle
	// Removes the assigned lifecycle policy and stops managing the specified index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-remove-policy.html
	RemovePolicy ilm_remove_policy.NewRemovePolicy
	// Retries executing the policy for an index that is in the ERROR step.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-retry-policy.html
	Retry ilm_retry.NewRetry
	// Start the index lifecycle management (ILM) plugin.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-start.html
	Start ilm_start.NewStart
	// Halts all lifecycle management operations and stops the index lifecycle
	// management (ILM) plugin
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-stop.html
	Stop ilm_stop.NewStop
}

type Indices struct {
	// Adds a block to an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules-blocks.html
	AddBlock indices_add_block.NewAddBlock
	// Performs analysis on a text string and returns the resulting tokens.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-analyze.html
	Analyze indices_analyze.NewAnalyze
	// Clears the caches of one or more indices.
	// For data streams, the API clears the caches of the stream’s backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-clearcache.html
	ClearCache indices_clear_cache.NewClearCache
	// Clones an existing index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-clone-index.html
	Clone indices_clone.NewClone
	// Closes an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-close.html
	Close indices_close.NewClose
	// Creates a new index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
	Create indices_create.NewCreate
	// Creates a data stream.
	// You must have a matching index template with data stream enabled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	CreateDataStream indices_create_data_stream.NewCreateDataStream
	// Retrieves statistics for one or more data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	DataStreamsStats indices_data_streams_stats.NewDataStreamsStats
	// Deletes one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-index.html
	Delete indices_delete.NewDelete
	// Removes a data stream or index from an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	DeleteAlias indices_delete_alias.NewDeleteAlias
	// Removes the data lifecycle from a data stream rendering it not managed by the
	// data stream lifecycle
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-delete-lifecycle.html
	DeleteDataLifecycle indices_delete_data_lifecycle.NewDeleteDataLifecycle
	// Deletes one or more data streams and their backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	DeleteDataStream indices_delete_data_stream.NewDeleteDataStream
	// Delete an index template.
	// The provided <index-template> may contain multiple template names separated
	// by a comma. If multiple template
	// names are specified then there is no wildcard support and the provided names
	// should match completely with
	// existing templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-template.html
	DeleteIndexTemplate indices_delete_index_template.NewDeleteIndexTemplate
	// Deletes a legacy index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-template-v1.html
	DeleteTemplate indices_delete_template.NewDeleteTemplate
	// Analyzes the disk usage of each field of an index or data stream.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-disk-usage.html
	DiskUsage indices_disk_usage.NewDiskUsage
	// Aggregates a time series (TSDS) index and stores pre-computed statistical
	// summaries (`min`, `max`, `sum`, `value_count` and `avg`) for each metric
	// field grouped by a configured time interval.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-downsample-data-stream.html
	Downsample indices_downsample.NewDownsample
	// Checks if a data stream, index, or alias exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-exists.html
	Exists indices_exists.NewExists
	// Checks if an alias exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	ExistsAlias indices_exists_alias.NewExistsAlias
	// Returns information about whether a particular index template exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-templates.html
	ExistsIndexTemplate indices_exists_index_template.NewExistsIndexTemplate
	// Check existence of index templates.
	// Returns information about whether a particular index template exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-template-exists-v1.html
	ExistsTemplate indices_exists_template.NewExistsTemplate
	// Retrieves information about the index's current data stream lifecycle, such
	// as any potential encountered error, time since creation etc.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-explain-lifecycle.html
	ExplainDataLifecycle indices_explain_data_lifecycle.NewExplainDataLifecycle
	// Returns field usage information for each shard and field of an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/field-usage-stats.html
	FieldUsageStats indices_field_usage_stats.NewFieldUsageStats
	// Flushes one or more data streams or indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-flush.html
	Flush indices_flush.NewFlush
	// Performs the force merge operation on one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-forcemerge.html
	Forcemerge indices_forcemerge.NewForcemerge
	// Returns information about one or more indices. For data streams, the API
	// returns information about the
	// stream’s backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-index.html
	Get indices_get.NewGet
	// Retrieves information for one or more aliases.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	GetAlias indices_get_alias.NewGetAlias
	// Retrieves the data stream lifecycle configuration of one or more data
	// streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-get-lifecycle.html
	GetDataLifecycle indices_get_data_lifecycle.NewGetDataLifecycle
	// Retrieves information about one or more data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	GetDataStream indices_get_data_stream.NewGetDataStream
	// Retrieves mapping definitions for one or more fields.
	// For data streams, the API retrieves field mappings for the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-field-mapping.html
	GetFieldMapping indices_get_field_mapping.NewGetFieldMapping
	// Get index templates.
	// Returns information about one or more index templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-template.html
	GetIndexTemplate indices_get_index_template.NewGetIndexTemplate
	// Retrieves mapping definitions for one or more indices.
	// For data streams, the API retrieves mappings for the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-mapping.html
	GetMapping indices_get_mapping.NewGetMapping
	// Returns setting information for one or more indices. For data streams,
	// returns setting information for the stream’s backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-settings.html
	GetSettings indices_get_settings.NewGetSettings
	// Get index templates.
	// Retrieves information about one or more index templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-template-v1.html
	GetTemplate indices_get_template.NewGetTemplate
	// Converts an index alias to a data stream.
	// You must have a matching index template that is data stream enabled.
	// The alias must meet the following criteria:
	// The alias must have a write index;
	// All indices for the alias must have a `@timestamp` field mapping of a `date`
	// or `date_nanos` field type;
	// The alias must not have any filters;
	// The alias must not use custom routing.
	// If successful, the request removes the alias and creates a data stream with
	// the same name.
	// The indices for the alias become hidden backing indices for the stream.
	// The write index for the alias becomes the write index for the stream.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	MigrateToDataStream indices_migrate_to_data_stream.NewMigrateToDataStream
	// Performs one or more data stream modification actions in a single atomic
	// operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	ModifyDataStream indices_modify_data_stream.NewModifyDataStream
	// Opens a closed index.
	// For data streams, the API opens any closed backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-open-close.html
	Open indices_open.NewOpen
	// Promotes a data stream from a replicated data stream managed by CCR to a
	// regular data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	PromoteDataStream indices_promote_data_stream.NewPromoteDataStream
	// Adds a data stream or index to an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	PutAlias indices_put_alias.NewPutAlias
	// Update the data lifecycle of the specified data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-put-lifecycle.html
	PutDataLifecycle indices_put_data_lifecycle.NewPutDataLifecycle
	// Create or update an index template.
	// Index templates define settings, mappings, and aliases that can be applied
	// automatically to new indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-template.html
	PutIndexTemplate indices_put_index_template.NewPutIndexTemplate
	// Adds new fields to an existing data stream or index.
	// You can also use this API to change the search settings of existing fields.
	// For data streams, these changes are applied to all backing indices by
	// default.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-mapping.html
	PutMapping indices_put_mapping.NewPutMapping
	// Changes a dynamic index setting in real time. For data streams, index setting
	// changes are applied to all backing indices by default.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-update-settings.html
	PutSettings indices_put_settings.NewPutSettings
	// Create or update an index template.
	// Index templates define settings, mappings, and aliases that can be applied
	// automatically to new indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates-v1.html
	PutTemplate indices_put_template.NewPutTemplate
	// Returns information about ongoing and completed shard recoveries for one or
	// more indices.
	// For data streams, the API returns information for the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-recovery.html
	Recovery indices_recovery.NewRecovery
	// A refresh makes recent operations performed on one or more indices available
	// for search.
	// For data streams, the API runs the refresh operation on the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-refresh.html
	Refresh indices_refresh.NewRefresh
	// Reloads an index's search analyzers and their resources.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-reload-analyzers.html
	ReloadSearchAnalyzers indices_reload_search_analyzers.NewReloadSearchAnalyzers
	// Resolves the specified index expressions to return information about each
	// cluster, including
	// the local cluster, if included.
	// Multiple patterns and remote clusters are supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-resolve-cluster-api.html
	ResolveCluster indices_resolve_cluster.NewResolveCluster
	// Resolves the specified name(s) and/or index patterns for indices, aliases,
	// and data streams.
	// Multiple patterns and remote clusters are supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-resolve-index-api.html
	ResolveIndex indices_resolve_index.NewResolveIndex
	// Creates a new index for a data stream or index alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-rollover-index.html
	Rollover indices_rollover.NewRollover
	// Returns low-level information about the Lucene segments in index shards.
	// For data streams, the API returns information about the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-segments.html
	Segments indices_segments.NewSegments
	// Retrieves store information about replica shards in one or more indices.
	// For data streams, the API retrieves store information for the stream’s
	// backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shards-stores.html
	ShardStores indices_shard_stores.NewShardStores
	// Shrinks an existing index into a new index with fewer primary shards.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shrink-index.html
	Shrink indices_shrink.NewShrink
	// Simulate an index.
	// Returns the index configuration that would be applied to the specified index
	// from an existing index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-simulate-index.html
	SimulateIndexTemplate indices_simulate_index_template.NewSimulateIndexTemplate
	// Simulate an index template.
	// Returns the index configuration that would be applied by a particular index
	// template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-simulate-template.html
	SimulateTemplate indices_simulate_template.NewSimulateTemplate
	// Splits an existing index into a new index with more primary shards.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-split-index.html
	Split indices_split.NewSplit
	// Returns statistics for one or more indices.
	// For data streams, the API retrieves statistics for the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-stats.html
	Stats indices_stats.NewStats
	// Unfreezes an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/unfreeze-index-api.html
	Unfreeze indices_unfreeze.NewUnfreeze
	// Adds a data stream or index to an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	UpdateAliases indices_update_aliases.NewUpdateAliases
	// Validates a potentially expensive query without executing it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-validate.html
	ValidateQuery indices_validate_query.NewValidateQuery
}

type Inference struct {
	// Delete an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-inference-api.html
	Delete inference_delete.NewDelete
	// Get an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-inference-api.html
	Get inference_get.NewGet
	// Perform inference on the service
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html
	Inference inference_inference.NewInference
	// Create an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-inference-api.html
	Put inference_put.NewPut
}

type Ingest struct {
	// Deletes one or more existing ingest pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-pipeline-api.html
	DeletePipeline ingest_delete_pipeline.NewDeletePipeline
	// Gets download statistics for GeoIP2 databases used with the geoip processor.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html
	GeoIpStats ingest_geo_ip_stats.NewGeoIpStats
	// Returns information about one or more ingest pipelines.
	// This API returns a local reference of the pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-pipeline-api.html
	GetPipeline ingest_get_pipeline.NewGetPipeline
	// Extracts structured fields out of a single text field within a document.
	// You choose which field to extract matched fields from, as well as the grok
	// pattern you expect will match.
	// A grok pattern is like a regular expression that supports aliased expressions
	// that can be reused.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html
	ProcessorGrok ingest_processor_grok.NewProcessorGrok
	// Creates or updates an ingest pipeline.
	// Changes made using this API take effect immediately.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html
	PutPipeline ingest_put_pipeline.NewPutPipeline
	// Executes an ingest pipeline against a set of provided documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-pipeline-api.html
	Simulate ingest_simulate.NewSimulate
}

type License struct {
	// Deletes licensing information for the cluster
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-license.html
	Delete license_delete.NewDelete
	// This API returns information about the type of license, when it was issued,
	// and when it expires, for example.
	// For more information about the different types of licenses, see
	// https://www.elastic.co/subscriptions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-license.html
	Get license_get.NewGet
	// Retrieves information about the status of the basic license.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-basic-status.html
	GetBasicStatus license_get_basic_status.NewGetBasicStatus
	// Retrieves information about the status of the trial license.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trial-status.html
	GetTrialStatus license_get_trial_status.NewGetTrialStatus
	// Updates the license for the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-license.html
	Post license_post.NewPost
	// The start basic API enables you to initiate an indefinite basic license,
	// which gives access to all the basic features. If the basic license does not
	// support all of the features that are available with your current license,
	// however, you are notified in the response. You must then re-submit the API
	// request with the acknowledge parameter set to true.
	// To check the status of your basic license, use the following API: [Get basic
	// status](https://www.elastic.co/guide/en/elasticsearch/reference/current/get-basic-status.html).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-basic.html
	PostStartBasic license_post_start_basic.NewPostStartBasic
	// The start trial API enables you to start a 30-day trial, which gives access
	// to all subscription features.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trial.html
	PostStartTrial license_post_start_trial.NewPostStartTrial
}

type Logstash struct {
	// Deletes a pipeline used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-delete-pipeline.html
	DeletePipeline logstash_delete_pipeline.NewDeletePipeline
	// Retrieves pipelines used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-get-pipeline.html
	GetPipeline logstash_get_pipeline.NewGetPipeline
	// Creates or updates a pipeline used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-put-pipeline.html
	PutPipeline logstash_put_pipeline.NewPutPipeline
}

type Migration struct {
	// Retrieves information about different cluster, node, and index level settings
	// that use deprecated features that will be removed or changed in the next
	// major version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migration-api-deprecation.html
	Deprecations migration_deprecations.NewDeprecations
	// Find out whether system features need to be upgraded or not
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migration-api-feature-upgrade.html
	GetFeatureUpgradeStatus migration_get_feature_upgrade_status.NewGetFeatureUpgradeStatus
	// Begin upgrades for system features
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migration-api-feature-upgrade.html
	PostFeatureUpgrade migration_post_feature_upgrade.NewPostFeatureUpgrade
}

type Ml struct {
	// Clears a trained model deployment cache on all nodes where the trained model
	// is assigned.
	// A trained model deployment may have an inference cache enabled.
	// As requests are handled by each allocated node, their responses may be cached
	// on that individual node.
	// Calling this API clears the caches without restarting the deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-trained-model-deployment-cache.html
	ClearTrainedModelDeploymentCache ml_clear_trained_model_deployment_cache.NewClearTrainedModelDeploymentCache
	// Close anomaly detection jobs
	// A job can be opened and closed multiple times throughout its lifecycle. A
	// closed job cannot receive data or perform analysis operations, but you can
	// still explore and navigate results.
	// When you close a job, it runs housekeeping tasks such as pruning the model
	// history, flushing buffers, calculating final results and persisting the model
	// snapshots. Depending upon the size of the job, it could take several minutes
	// to close and the equivalent time to re-open. After it is closed, the job has
	// a minimal overhead on the cluster except for maintaining its meta data.
	// Therefore it is a best practice to close jobs that are no longer required to
	// process data.
	// If you close an anomaly detection job whose datafeed is running, the request
	// first tries to stop the datafeed. This behavior is equivalent to calling stop
	// datafeed API with the same timeout and force parameters as the close job
	// request.
	// When a datafeed that has a specified end date stops, it automatically closes
	// its associated job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-close-job.html
	CloseJob ml_close_job.NewCloseJob
	// Removes all scheduled events from a calendar, then deletes it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar.html
	DeleteCalendar ml_delete_calendar.NewDeleteCalendar
	// Deletes scheduled events from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-event.html
	DeleteCalendarEvent ml_delete_calendar_event.NewDeleteCalendarEvent
	// Deletes anomaly detection jobs from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-job.html
	DeleteCalendarJob ml_delete_calendar_job.NewDeleteCalendarJob
	// Deletes a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-dfanalytics.html
	DeleteDataFrameAnalytics ml_delete_data_frame_analytics.NewDeleteDataFrameAnalytics
	// Deletes an existing datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-datafeed.html
	DeleteDatafeed ml_delete_datafeed.NewDeleteDatafeed
	// Deletes expired and unused machine learning data.
	// Deletes all job results, model snapshots and forecast data that have exceeded
	// their retention days period. Machine learning state documents that are not
	// associated with any job are also deleted.
	// You can limit the request to a single or set of anomaly detection jobs by
	// using a job identifier, a group name, a comma-separated list of jobs, or a
	// wildcard expression. You can delete expired data for all anomaly detection
	// jobs by using _all, by specifying * as the <job_id>, or by omitting the
	// <job_id>.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-expired-data.html
	DeleteExpiredData ml_delete_expired_data.NewDeleteExpiredData
	// Deletes a filter.
	// If an anomaly detection job references the filter, you cannot delete the
	// filter. You must update or delete the job before you can delete the filter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-filter.html
	DeleteFilter ml_delete_filter.NewDeleteFilter
	// Deletes forecasts from a machine learning job.
	// By default, forecasts are retained for 14 days. You can specify a
	// different retention period with the `expires_in` parameter in the forecast
	// jobs API. The delete forecast API enables you to delete one or more
	// forecasts before they expire.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-forecast.html
	DeleteForecast ml_delete_forecast.NewDeleteForecast
	// Deletes an anomaly detection job.
	//
	// All job configuration, model state and results are deleted.
	// It is not currently possible to delete multiple jobs using wildcards or a
	// comma separated list. If you delete a job that has a datafeed, the request
	// first tries to delete the datafeed. This behavior is equivalent to calling
	// the delete datafeed API with the same timeout and force parameters as the
	// delete job request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-job.html
	DeleteJob ml_delete_job.NewDeleteJob
	// Deletes an existing model snapshot.
	// You cannot delete the active model snapshot. To delete that snapshot, first
	// revert to a different one. To identify the active model snapshot, refer to
	// the `model_snapshot_id` in the results from the get jobs API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-snapshot.html
	DeleteModelSnapshot ml_delete_model_snapshot.NewDeleteModelSnapshot
	// Deletes an existing trained inference model that is currently not referenced
	// by an ingest pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models.html
	DeleteTrainedModel ml_delete_trained_model.NewDeleteTrainedModel
	// Deletes a trained model alias.
	// This API deletes an existing model alias that refers to a trained model. If
	// the model alias is missing or refers to a model other than the one identified
	// by the `model_id`, this API returns an error.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models-aliases.html
	DeleteTrainedModelAlias ml_delete_trained_model_alias.NewDeleteTrainedModelAlias
	// Makes an estimation of the memory usage for an anomaly detection job model.
	// It is based on analysis configuration details for the job and cardinality
	// estimates for the fields it references.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-apis.html
	EstimateModelMemory ml_estimate_model_memory.NewEstimateModelMemory
	// Evaluates the data frame analytics for an annotated index.
	// The API packages together commonly used evaluation metrics for various types
	// of machine learning features. This has been designed for use on indexes
	// created by data frame analytics. Evaluation requires both a ground truth
	// field and an analytics result field to be present.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/evaluate-dfanalytics.html
	EvaluateDataFrame ml_evaluate_data_frame.NewEvaluateDataFrame
	// Explains a data frame analytics config.
	// This API provides explanations for a data frame analytics config that either
	// exists already or one that has not been created yet. The following
	// explanations are provided:
	// * which fields are included or not in the analysis and why,
	// * how much memory is estimated to be required. The estimate can be used when
	// deciding the appropriate value for model_memory_limit setting later on.
	// If you have object fields or fields that are excluded via source filtering,
	// they are not included in the explanation.
	// http://www.elastic.co/guide/en/elasticsearch/reference/current/explain-dfanalytics.html
	ExplainDataFrameAnalytics ml_explain_data_frame_analytics.NewExplainDataFrameAnalytics
	// Forces any buffered data to be processed by the job.
	// The flush jobs API is only applicable when sending data for analysis using
	// the post data API. Depending on the content of the buffer, then it might
	// additionally calculate new results. Both flush and close operations are
	// similar, however the flush is more efficient if you are expecting to send
	// more data for analysis. When flushing, the job remains open and is available
	// to continue analyzing data. A close operation additionally prunes and
	// persists the model state to disk and the job must be opened again before
	// analyzing further data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-flush-job.html
	FlushJob ml_flush_job.NewFlushJob
	// Predicts the future behavior of a time series by using its historical
	// behavior.
	//
	// Forecasts are not supported for jobs that perform population analysis; an
	// error occurs if you try to create a forecast for a job that has an
	// `over_field_name` in its configuration.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-forecast.html
	Forecast ml_forecast.NewForecast
	// Retrieves anomaly detection job results for one or more buckets.
	// The API presents a chronological view of the records, grouped by bucket.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-bucket.html
	GetBuckets ml_get_buckets.NewGetBuckets
	// Retrieves information about the scheduled events in calendars.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar-event.html
	GetCalendarEvents ml_get_calendar_events.NewGetCalendarEvents
	// Retrieves configuration information for calendars.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar.html
	GetCalendars ml_get_calendars.NewGetCalendars
	// Retrieves anomaly detection job results for one or more categories.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-category.html
	GetCategories ml_get_categories.NewGetCategories
	// Retrieves configuration information for data frame analytics jobs.
	// You can get information for multiple data frame analytics jobs in a single
	// API request by using a comma-separated list of data frame analytics jobs or a
	// wildcard expression.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics.html
	GetDataFrameAnalytics ml_get_data_frame_analytics.NewGetDataFrameAnalytics
	// Retrieves usage information for data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics-stats.html
	GetDataFrameAnalyticsStats ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStats
	// Retrieves usage information for datafeeds.
	// You can get statistics for multiple datafeeds in a single API request by
	// using a comma-separated list of datafeeds or a wildcard expression. You can
	// get statistics for all datafeeds by using `_all`, by specifying `*` as the
	// `<feed_id>`, or by omitting the `<feed_id>`. If the datafeed is stopped, the
	// only information you receive is the `datafeed_id` and the `state`.
	// This API returns a maximum of 10,000 datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed-stats.html
	GetDatafeedStats ml_get_datafeed_stats.NewGetDatafeedStats
	// Retrieves configuration information for datafeeds.
	// You can get information for multiple datafeeds in a single API request by
	// using a comma-separated list of datafeeds or a wildcard expression. You can
	// get information for all datafeeds by using `_all`, by specifying `*` as the
	// `<feed_id>`, or by omitting the `<feed_id>`.
	// This API returns a maximum of 10,000 datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed.html
	GetDatafeeds ml_get_datafeeds.NewGetDatafeeds
	// Retrieves filters.
	// You can get a single filter or all filters.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-filter.html
	GetFilters ml_get_filters.NewGetFilters
	// Retrieves anomaly detection job results for one or more influencers.
	// Influencers are the entities that have contributed to, or are to blame for,
	// the anomalies. Influencer results are available only if an
	// `influencer_field_name` is specified in the job configuration.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-influencer.html
	GetInfluencers ml_get_influencers.NewGetInfluencers
	// Retrieves usage information for anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-stats.html
	GetJobStats ml_get_job_stats.NewGetJobStats
	// Retrieves configuration information for anomaly detection jobs.
	// You can get information for multiple anomaly detection jobs in a single API
	// request by using a group name, a comma-separated list of jobs, or a wildcard
	// expression. You can get information for all anomaly detection jobs by using
	// `_all`, by specifying `*` as the `<job_id>`, or by omitting the `<job_id>`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job.html
	GetJobs ml_get_jobs.NewGetJobs
	// Get information about how machine learning jobs and trained models are using
	// memory,
	// on each node, both within the JVM heap, and natively, outside of the JVM.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-memory.html
	GetMemoryStats ml_get_memory_stats.NewGetMemoryStats
	// Retrieves usage information for anomaly detection job model snapshot
	// upgrades.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-model-snapshot-upgrade-stats.html
	GetModelSnapshotUpgradeStats ml_get_model_snapshot_upgrade_stats.NewGetModelSnapshotUpgradeStats
	// Retrieves information about model snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-snapshot.html
	GetModelSnapshots ml_get_model_snapshots.NewGetModelSnapshots
	// Retrieves overall bucket results that summarize the bucket results of
	// multiple anomaly detection jobs.
	//
	// The `overall_score` is calculated by combining the scores of all the
	// buckets within the overall bucket span. First, the maximum
	// `anomaly_score` per anomaly detection job in the overall bucket is
	// calculated. Then the `top_n` of those scores are averaged to result in
	// the `overall_score`. This means that you can fine-tune the
	// `overall_score` so that it is more or less sensitive to the number of
	// jobs that detect an anomaly at the same time. For example, if you set
	// `top_n` to `1`, the `overall_score` is the maximum bucket score in the
	// overall bucket. Alternatively, if you set `top_n` to the number of jobs,
	// the `overall_score` is high only when all jobs detect anomalies in that
	// overall bucket. If you set the `bucket_span` parameter (to a value
	// greater than its default), the `overall_score` is the maximum
	// `overall_score` of the overall buckets that have a span equal to the
	// jobs' largest bucket span.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-overall-buckets.html
	GetOverallBuckets ml_get_overall_buckets.NewGetOverallBuckets
	// Retrieves anomaly records for an anomaly detection job.
	// Records contain the detailed analytical results. They describe the anomalous
	// activity that has been identified in the input data based on the detector
	// configuration.
	// There can be many anomaly records depending on the characteristics and size
	// of the input data. In practice, there are often too many to be able to
	// manually process them. The machine learning features therefore perform a
	// sophisticated aggregation of the anomaly records into buckets.
	// The number of record results depends on the number of anomalies found in each
	// bucket, which relates to the number of time series being modeled and the
	// number of detectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-record.html
	GetRecords ml_get_records.NewGetRecords
	// Retrieves configuration information for a trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models.html
	GetTrainedModels ml_get_trained_models.NewGetTrainedModels
	// Retrieves usage information for trained models. You can get usage information
	// for multiple trained
	// models in a single API request by using a comma-separated list of model IDs
	// or a wildcard expression.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models-stats.html
	GetTrainedModelsStats ml_get_trained_models_stats.NewGetTrainedModelsStats
	// Evaluates a trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-trained-model.html
	InferTrainedModel ml_infer_trained_model.NewInferTrainedModel
	// Returns defaults and limits used by machine learning.
	// This endpoint is designed to be used by a user interface that needs to fully
	// understand machine learning configurations where some options are not
	// specified, meaning that the defaults should be used. This endpoint may be
	// used to find out what those defaults are. It also provides information about
	// the maximum size of machine learning jobs that could run in the current
	// cluster configuration.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-info.html
	Info ml_info.NewInfo
	// Opens one or more anomaly detection jobs.
	// An anomaly detection job must be opened in order for it to be ready to
	// receive and analyze data. It can be opened and closed multiple times
	// throughout its lifecycle.
	// When you open a new job, it starts with an empty model.
	// When you open an existing job, the most recent model state is automatically
	// loaded. The job is ready to resume its analysis from where it left off, once
	// new data is received.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-open-job.html
	OpenJob ml_open_job.NewOpenJob
	// Adds scheduled events to a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-calendar-event.html
	PostCalendarEvents ml_post_calendar_events.NewPostCalendarEvents
	// Sends data to an anomaly detection job for analysis.
	//
	// IMPORTANT: For each job, data can be accepted from only a single connection
	// at a time.
	// It is not currently possible to post data to multiple jobs using wildcards or
	// a comma-separated list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-data.html
	PostData ml_post_data.NewPostData
	// Previews the extracted features used by a data frame analytics config.
	// http://www.elastic.co/guide/en/elasticsearch/reference/current/preview-dfanalytics.html
	PreviewDataFrameAnalytics ml_preview_data_frame_analytics.NewPreviewDataFrameAnalytics
	// Previews a datafeed.
	// This API returns the first "page" of search results from a datafeed.
	// You can preview an existing datafeed or provide configuration details for a
	// datafeed
	// and anomaly detection job in the API. The preview shows the structure of the
	// data
	// that will be passed to the anomaly detection engine.
	// IMPORTANT: When Elasticsearch security features are enabled, the preview uses
	// the credentials of the user that
	// called the API. However, when the datafeed starts it uses the roles of the
	// last user that created or updated the
	// datafeed. To get a preview that accurately reflects the behavior of the
	// datafeed, use the appropriate credentials.
	// You can also use secondary authorization headers to supply the credentials.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-preview-datafeed.html
	PreviewDatafeed ml_preview_datafeed.NewPreviewDatafeed
	// Creates a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar.html
	PutCalendar ml_put_calendar.NewPutCalendar
	// Adds an anomaly detection job to a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar-job.html
	PutCalendarJob ml_put_calendar_job.NewPutCalendarJob
	// Instantiates a data frame analytics job.
	// This API creates a data frame analytics job that performs an analysis on the
	// source indices and stores the outcome in a destination index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-dfanalytics.html
	PutDataFrameAnalytics ml_put_data_frame_analytics.NewPutDataFrameAnalytics
	// Instantiates a datafeed.
	// Datafeeds retrieve data from Elasticsearch for analysis by an anomaly
	// detection job.
	// You can associate only one datafeed with each anomaly detection job.
	// The datafeed contains a query that runs at a defined interval (`frequency`).
	// If you are concerned about delayed data, you can add a delay (`query_delay')
	// at each interval.
	// When Elasticsearch security features are enabled, your datafeed remembers
	// which roles the user who created it had
	// at the time of creation and runs the query using those same roles. If you
	// provide secondary authorization headers,
	// those credentials are used instead.
	// You must use Kibana, this API, or the create anomaly detection jobs API to
	// create a datafeed. Do not add a datafeed
	// directly to the `.ml-config` index. Do not give users `write` privileges on
	// the `.ml-config` index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-datafeed.html
	PutDatafeed ml_put_datafeed.NewPutDatafeed
	// Instantiates a filter.
	// A filter contains a list of strings. It can be used by one or more anomaly
	// detection jobs.
	// Specifically, filters are referenced in the `custom_rules` property of
	// detector configuration objects.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-filter.html
	PutFilter ml_put_filter.NewPutFilter
	// Instantiates an anomaly detection job. If you include a `datafeed_config`,
	// you must have read index privileges on the source index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-job.html
	PutJob ml_put_job.NewPutJob
	// Enables you to supply a trained model that is not created by data frame
	// analytics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models.html
	PutTrainedModel ml_put_trained_model.NewPutTrainedModel
	// Creates or updates a trained model alias. A trained model alias is a logical
	// name used to reference a single trained model.
	// You can use aliases instead of trained model identifiers to make it easier to
	// reference your models. For example, you can use aliases in inference
	// aggregations and processors.
	// An alias must be unique and refer to only a single trained model. However,
	// you can have multiple aliases for each trained model.
	// If you use this API to update an alias such that it references a different
	// trained model ID and the model uses a different type of data frame analytics,
	// an error occurs. For example, this situation occurs if you have a trained
	// model for regression analysis and a trained model for classification
	// analysis; you cannot reassign an alias from one type of trained model to
	// another.
	// If you use this API to update an alias and there are very few input fields in
	// common between the old and new trained models for the model alias, the API
	// returns a warning.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models-aliases.html
	PutTrainedModelAlias ml_put_trained_model_alias.NewPutTrainedModelAlias
	// Creates part of a trained model definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-definition-part.html
	PutTrainedModelDefinitionPart ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPart
	// Creates a trained model vocabulary.
	// This API is supported only for natural language processing (NLP) models.
	// The vocabulary is stored in the index as described in
	// `inference_config.*.vocabulary` of the trained model definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-vocabulary.html
	PutTrainedModelVocabulary ml_put_trained_model_vocabulary.NewPutTrainedModelVocabulary
	// Resets an anomaly detection job.
	// All model state and results are deleted. The job is ready to start over as if
	// it had just been created.
	// It is not currently possible to reset multiple jobs using wildcards or a
	// comma separated list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-reset-job.html
	ResetJob ml_reset_job.NewResetJob
	// Reverts to a specific snapshot.
	// The machine learning features react quickly to anomalous input, learning new
	// behaviors in data. Highly anomalous input increases the variance in the
	// models whilst the system learns whether this is a new step-change in behavior
	// or a one-off event. In the case where this anomalous input is known to be a
	// one-off, then it might be appropriate to reset the model state to a time
	// before this event. For example, you might consider reverting to a saved
	// snapshot after Black Friday or a critical system failure.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-revert-snapshot.html
	RevertModelSnapshot ml_revert_model_snapshot.NewRevertModelSnapshot
	// Sets a cluster wide upgrade_mode setting that prepares machine learning
	// indices for an upgrade.
	// When upgrading your cluster, in some circumstances you must restart your
	// nodes and reindex your machine learning indices. In those circumstances,
	// there must be no machine learning jobs running. You can close the machine
	// learning jobs, do the upgrade, then open all the jobs again. Alternatively,
	// you can use this API to temporarily halt tasks associated with the jobs and
	// datafeeds and prevent new jobs from opening. You can also use this API
	// during upgrades that do not require you to reindex your machine learning
	// indices, though stopping jobs is not a requirement in that case.
	// You can see the current value for the upgrade_mode setting by using the get
	// machine learning info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-set-upgrade-mode.html
	SetUpgradeMode ml_set_upgrade_mode.NewSetUpgradeMode
	// Starts a data frame analytics job.
	// A data frame analytics job can be started and stopped multiple times
	// throughout its lifecycle.
	// If the destination index does not exist, it is created automatically the
	// first time you start the data frame analytics job. The
	// `index.number_of_shards` and `index.number_of_replicas` settings for the
	// destination index are copied from the source index. If there are multiple
	// source indices, the destination index copies the highest setting values. The
	// mappings for the destination index are also copied from the source indices.
	// If there are any mapping conflicts, the job fails to start.
	// If the destination index exists, it is used as is. You can therefore set up
	// the destination index in advance with custom settings and mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-dfanalytics.html
	StartDataFrameAnalytics ml_start_data_frame_analytics.NewStartDataFrameAnalytics
	// Starts one or more datafeeds.
	//
	// A datafeed must be started in order to retrieve data from Elasticsearch. A
	// datafeed can be started and stopped
	// multiple times throughout its lifecycle.
	//
	// Before you can start a datafeed, the anomaly detection job must be open.
	// Otherwise, an error occurs.
	//
	// If you restart a stopped datafeed, it continues processing input data from
	// the next millisecond after it was stopped.
	// If new data was indexed for that exact millisecond between stopping and
	// starting, it will be ignored.
	//
	// When Elasticsearch security features are enabled, your datafeed remembers
	// which roles the last user to create or
	// update it had at the time of creation or update and runs the query using
	// those same roles. If you provided secondary
	// authorization headers when you created or updated the datafeed, those
	// credentials are used instead.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-start-datafeed.html
	StartDatafeed ml_start_datafeed.NewStartDatafeed
	// Starts a trained model deployment, which allocates the model to every machine
	// learning node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trained-model-deployment.html
	StartTrainedModelDeployment ml_start_trained_model_deployment.NewStartTrainedModelDeployment
	// Stops one or more data frame analytics jobs.
	// A data frame analytics job can be started and stopped multiple times
	// throughout its lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-dfanalytics.html
	StopDataFrameAnalytics ml_stop_data_frame_analytics.NewStopDataFrameAnalytics
	// Stops one or more datafeeds.
	// A datafeed that is stopped ceases to retrieve data from Elasticsearch. A
	// datafeed can be started and stopped
	// multiple times throughout its lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-stop-datafeed.html
	StopDatafeed ml_stop_datafeed.NewStopDatafeed
	// Stops a trained model deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-trained-model-deployment.html
	StopTrainedModelDeployment ml_stop_trained_model_deployment.NewStopTrainedModelDeployment
	// Updates an existing data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-dfanalytics.html
	UpdateDataFrameAnalytics ml_update_data_frame_analytics.NewUpdateDataFrameAnalytics
	// Updates the properties of a datafeed.
	// You must stop and start the datafeed for the changes to be applied.
	// When Elasticsearch security features are enabled, your datafeed remembers
	// which roles the user who updated it had at
	// the time of the update and runs the query using those same roles. If you
	// provide secondary authorization headers,
	// those credentials are used instead.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-datafeed.html
	UpdateDatafeed ml_update_datafeed.NewUpdateDatafeed
	// Updates the description of a filter, adds items, or removes items from the
	// list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-filter.html
	UpdateFilter ml_update_filter.NewUpdateFilter
	// Updates certain properties of an anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-job.html
	UpdateJob ml_update_job.NewUpdateJob
	// Updates certain properties of a snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-snapshot.html
	UpdateModelSnapshot ml_update_model_snapshot.NewUpdateModelSnapshot
	// Starts a trained model deployment, which allocates the model to every machine
	// learning node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-trained-model-deployment.html
	UpdateTrainedModelDeployment ml_update_trained_model_deployment.NewUpdateTrainedModelDeployment
	// Upgrades an anomaly detection model snapshot to the latest major version.
	// Over time, older snapshot formats are deprecated and removed. Anomaly
	// detection jobs support only snapshots that are from the current or previous
	// major version.
	// This API provides a means to upgrade a snapshot to the current major version.
	// This aids in preparing the cluster for an upgrade to the next major version.
	// Only one snapshot per anomaly detection job can be upgraded at a time and the
	// upgraded snapshot cannot be the current snapshot of the anomaly detection
	// job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-upgrade-job-model-snapshot.html
	UpgradeJobSnapshot ml_upgrade_job_snapshot.NewUpgradeJobSnapshot
	// Validates an anomaly detection job.
	// https://www.elastic.co/guide/en/machine-learning/current/ml-jobs.html
	Validate ml_validate.NewValidate
	// Validates an anomaly detection detector.
	// https://www.elastic.co/guide/en/machine-learning/current/ml-jobs.html
	ValidateDetector ml_validate_detector.NewValidateDetector
}

type Monitoring struct {
	// Used by the monitoring features to send monitoring data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/monitor-elasticsearch-cluster.html
	Bulk monitoring_bulk.NewBulk
}

type Nodes struct {
	// You can use this API to clear the archived repositories metering information
	// in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-repositories-metering-archive-api.html
	ClearRepositoriesMeteringArchive nodes_clear_repositories_metering_archive.NewClearRepositoriesMeteringArchive
	// You can use the cluster repositories metering API to retrieve repositories
	// metering information in a cluster.
	// This API exposes monotonically non-decreasing counters and it’s expected that
	// clients would durably store the
	// information needed to compute aggregations over a period of time.
	// Additionally, the information exposed by this
	// API is volatile, meaning that it won’t be present after node restarts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-repositories-metering-api.html
	GetRepositoriesMeteringInfo nodes_get_repositories_metering_info.NewGetRepositoriesMeteringInfo
	// This API yields a breakdown of the hot threads on each selected node in the
	// cluster.
	// The output is plain text with a breakdown of each node’s top hot threads.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-hot-threads.html
	HotThreads nodes_hot_threads.NewHotThreads
	// Returns cluster nodes information.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-info.html
	Info nodes_info.NewInfo
	// Reloads the keystore on nodes in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/secure-settings.html#reloadable-secure-settings
	ReloadSecureSettings nodes_reload_secure_settings.NewReloadSecureSettings
	// Returns cluster nodes statistics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-stats.html
	Stats nodes_stats.NewStats
	// Returns information on the usage of features.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-usage.html
	Usage nodes_usage.NewUsage
}

type Profiling struct {
	// Extracts a UI-optimized structure to render flamegraphs from Universal
	// Profiling.
	// https://www.elastic.co/guide/en/observability/current/universal-profiling.html
	Flamegraph profiling_flamegraph.NewFlamegraph
	// Extracts raw stacktrace information from Universal Profiling.
	// https://www.elastic.co/guide/en/observability/current/universal-profiling.html
	Stacktraces profiling_stacktraces.NewStacktraces
	// Returns basic information about the status of Universal Profiling.
	// https://www.elastic.co/guide/en/observability/current/universal-profiling.html
	Status profiling_status.NewStatus
	// Extracts a list of topN functions from Universal Profiling.
	// https://www.elastic.co/guide/en/observability/current/universal-profiling.html
	TopnFunctions profiling_topn_functions.NewTopnFunctions
}

type QueryRules struct {
	// Deletes a query rule within a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-rule.html
	DeleteRule query_rules_delete_rule.NewDeleteRule
	// Deletes a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-ruleset.html
	DeleteRuleset query_rules_delete_ruleset.NewDeleteRuleset
	// Returns the details about a query rule within a query ruleset
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-rule.html
	GetRule query_rules_get_rule.NewGetRule
	// Returns the details about a query ruleset
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-ruleset.html
	GetRuleset query_rules_get_ruleset.NewGetRuleset
	// Returns summarized information about existing query rulesets.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-query-rulesets.html
	ListRulesets query_rules_list_rulesets.NewListRulesets
	// Creates or updates a query rule within a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-rule.html
	PutRule query_rules_put_rule.NewPutRule
	// Creates or updates a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-ruleset.html
	PutRuleset query_rules_put_ruleset.NewPutRuleset
}

type Rollup struct {
	// Deletes an existing rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-delete-job.html
	DeleteJob rollup_delete_job.NewDeleteJob
	// Retrieves the configuration, stats, and status of rollup jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-job.html
	GetJobs rollup_get_jobs.NewGetJobs
	// Returns the capabilities of any rollup jobs that have been configured for a
	// specific index or index pattern.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-rollup-caps.html
	GetRollupCaps rollup_get_rollup_caps.NewGetRollupCaps
	// Returns the rollup capabilities of all jobs inside of a rollup index (for
	// example, the index where rollup data is stored).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-rollup-index-caps.html
	GetRollupIndexCaps rollup_get_rollup_index_caps.NewGetRollupIndexCaps
	// Creates a rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-put-job.html
	PutJob rollup_put_job.NewPutJob
	// Enables searching rolled-up data using the standard Query DSL.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-search.html
	RollupSearch rollup_rollup_search.NewRollupSearch
	// Starts an existing, stopped rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-start-job.html
	StartJob rollup_start_job.NewStartJob
	// Stops an existing, started rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-stop-job.html
	StopJob rollup_stop_job.NewStopJob
}

type SearchApplication struct {
	// Deletes a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-search-application.html
	Delete search_application_delete.NewDelete
	// Delete a behavioral analytics collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-analytics-collection.html
	DeleteBehavioralAnalytics search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalytics
	// Returns the details about a search application
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-search-application.html
	Get search_application_get.NewGet
	// Returns the existing behavioral analytics collections.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-analytics-collection.html
	GetBehavioralAnalytics search_application_get_behavioral_analytics.NewGetBehavioralAnalytics
	// Returns the existing search applications.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-search-applications.html
	List search_application_list.NewList
	// Creates or updates a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-search-application.html
	Put search_application_put.NewPut
	// Creates a behavioral analytics collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-analytics-collection.html
	PutBehavioralAnalytics search_application_put_behavioral_analytics.NewPutBehavioralAnalytics
	// Perform a search against a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-application-search.html
	Search search_application_search.NewSearch
}

type SearchableSnapshots struct {
	// Retrieve node-level cache statistics about searchable snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-apis.html
	CacheStats searchable_snapshots_cache_stats.NewCacheStats
	// Clear the cache of searchable snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-apis.html
	ClearCache searchable_snapshots_clear_cache.NewClearCache
	// Mount a snapshot as a searchable index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-api-mount-snapshot.html
	Mount searchable_snapshots_mount.NewMount
	// Retrieve shard-level statistics about searchable snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-apis.html
	Stats searchable_snapshots_stats.NewStats
}

type Security struct {
	// Creates or updates a user profile on behalf of another user.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-activate-user-profile.html
	ActivateUserProfile security_activate_user_profile.NewActivateUserProfile
	// Enables you to submit a request with a basic auth header to authenticate a
	// user and retrieve information about the authenticated user.
	// A successful call returns a JSON structure that shows user information such
	// as their username, the roles that are assigned to the user, any assigned
	// metadata, and information about the realms that authenticated and authorized
	// the user.
	// If the user cannot be authenticated, this API returns a 401 status code.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-authenticate.html
	Authenticate security_authenticate.NewAuthenticate
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The bulk delete roles API cannot delete roles that are defined in roles
	// files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-delete-role.html
	BulkDeleteRole security_bulk_delete_role.NewBulkDeleteRole
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The bulk create or update roles API cannot update roles that are defined in
	// roles files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-put-role.html
	BulkPutRole security_bulk_put_role.NewBulkPutRole
	// Updates the attributes of multiple existing API keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-update-api-keys.html
	BulkUpdateApiKeys security_bulk_update_api_keys.NewBulkUpdateApiKeys
	// Changes the passwords of users in the native realm and built-in users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-change-password.html
	ChangePassword security_change_password.NewChangePassword
	// Evicts a subset of all entries from the API key cache.
	// The cache is also automatically cleared on state changes of the security
	// index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-api-key-cache.html
	ClearApiKeyCache security_clear_api_key_cache.NewClearApiKeyCache
	// Evicts application privileges from the native application privileges cache.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-privilege-cache.html
	ClearCachedPrivileges security_clear_cached_privileges.NewClearCachedPrivileges
	// Evicts users from the user cache. Can completely clear the cache or evict
	// specific users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-cache.html
	ClearCachedRealms security_clear_cached_realms.NewClearCachedRealms
	// Evicts roles from the native role cache.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-role-cache.html
	ClearCachedRoles security_clear_cached_roles.NewClearCachedRoles
	// Evicts tokens from the service account token caches.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-service-token-caches.html
	ClearCachedServiceTokens security_clear_cached_service_tokens.NewClearCachedServiceTokens
	// Creates an API key for access without requiring basic authentication.
	// A successful request returns a JSON structure that contains the API key, its
	// unique id, and its name.
	// If applicable, it also returns expiration information for the API key in
	// milliseconds.
	// NOTE: By default, API keys never expire. You can specify expiration
	// information when you create the API keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
	CreateApiKey security_create_api_key.NewCreateApiKey
	// Creates a cross-cluster API key for API key based remote cluster access.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-cross-cluster-api-key.html
	CreateCrossClusterApiKey security_create_cross_cluster_api_key.NewCreateCrossClusterApiKey
	// Creates a service accounts token for access without requiring basic
	// authentication.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-service-token.html
	CreateServiceToken security_create_service_token.NewCreateServiceToken
	// Removes application privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-privilege.html
	DeletePrivileges security_delete_privileges.NewDeletePrivileges
	// Removes roles in the native realm.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-role.html
	DeleteRole security_delete_role.NewDeleteRole
	// Removes role mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-role-mapping.html
	DeleteRoleMapping security_delete_role_mapping.NewDeleteRoleMapping
	// Deletes a service account token.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-service-token.html
	DeleteServiceToken security_delete_service_token.NewDeleteServiceToken
	// Deletes users from the native realm.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-user.html
	DeleteUser security_delete_user.NewDeleteUser
	// Disables users in the native realm.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-disable-user.html
	DisableUser security_disable_user.NewDisableUser
	// Disables a user profile so it's not visible in user profile searches.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-disable-user-profile.html
	DisableUserProfile security_disable_user_profile.NewDisableUserProfile
	// Enables users in the native realm.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-enable-user.html
	EnableUser security_enable_user.NewEnableUser
	// Enables a user profile so it's visible in user profile searches.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-enable-user-profile.html
	EnableUserProfile security_enable_user_profile.NewEnableUserProfile
	// Enables a Kibana instance to configure itself for communication with a
	// secured Elasticsearch cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-kibana-enrollment.html
	EnrollKibana security_enroll_kibana.NewEnrollKibana
	// Allows a new node to join an existing cluster with security features enabled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-node-enrollment.html
	EnrollNode security_enroll_node.NewEnrollNode
	// Retrieves information for one or more API keys.
	// NOTE: If you have only the `manage_own_api_key` privilege, this API returns
	// only the API keys that you own.
	// If you have `read_security`, `manage_api_key` or greater privileges
	// (including `manage_security`), this API returns all API keys regardless of
	// ownership.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html
	GetApiKey security_get_api_key.NewGetApiKey
	// Retrieves the list of cluster privileges and index privileges that are
	// available in this version of Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-builtin-privileges.html
	GetBuiltinPrivileges security_get_builtin_privileges.NewGetBuiltinPrivileges
	// Retrieves application privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-privileges.html
	GetPrivileges security_get_privileges.NewGetPrivileges
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The get roles API cannot retrieve roles that are defined in roles files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role.html
	GetRole security_get_role.NewGetRole
	// Retrieves role mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role-mapping.html
	GetRoleMapping security_get_role_mapping.NewGetRoleMapping
	// This API returns a list of service accounts that match the provided path
	// parameter(s).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-service-accounts.html
	GetServiceAccounts security_get_service_accounts.NewGetServiceAccounts
	// Retrieves information of all service credentials for a service account.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-service-credentials.html
	GetServiceCredentials security_get_service_credentials.NewGetServiceCredentials
	// Retrieve settings for the security system indices
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-settings.html
	GetSettings security_get_settings.NewGetSettings
	// Creates a bearer token for access without requiring basic authentication.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-token.html
	GetToken security_get_token.NewGetToken
	// Retrieves information about users in the native realm and built-in users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user.html
	GetUser security_get_user.NewGetUser
	// Retrieves security privileges for the logged in user.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-privileges.html
	GetUserPrivileges security_get_user_privileges.NewGetUserPrivileges
	// Retrieves a user's profile using the unique profile ID.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-profile.html
	GetUserProfile security_get_user_profile.NewGetUserProfile
	// Creates an API key on behalf of another user.
	// This API is similar to Create API keys, however it creates the API key for a
	// user that is different than the user that runs the API.
	// The caller must have authentication credentials (either an access token, or a
	// username and password) for the user on whose behalf the API key will be
	// created.
	// It is not possible to use this API to create an API key without that user’s
	// credentials.
	// The user, for whom the authentication credentials is provided, can optionally
	// "run as" (impersonate) another user.
	// In this case, the API key will be created on behalf of the impersonated user.
	//
	// This API is intended be used by applications that need to create and manage
	// API keys for end users, but cannot guarantee that those users have permission
	// to create API keys on their own behalf.
	//
	// A successful grant API key API call returns a JSON structure that contains
	// the API key, its unique id, and its name.
	// If applicable, it also returns expiration information for the API key in
	// milliseconds.
	//
	// By default, API keys never expire. You can specify expiration information
	// when you create the API keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-grant-api-key.html
	GrantApiKey security_grant_api_key.NewGrantApiKey
	// Determines whether the specified user has a specified list of privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges.html
	HasPrivileges security_has_privileges.NewHasPrivileges
	// Determines whether the users associated with the specified profile IDs have
	// all the requested privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges-user-profile.html
	HasPrivilegesUserProfile security_has_privileges_user_profile.NewHasPrivilegesUserProfile
	// Invalidates one or more API keys.
	// The `manage_api_key` privilege allows deleting any API keys.
	// The `manage_own_api_key` only allows deleting API keys that are owned by the
	// user.
	// In addition, with the `manage_own_api_key` privilege, an invalidation request
	// must be issued in one of the three formats:
	// - Set the parameter `owner=true`.
	// - Or, set both `username` and `realm_name` to match the user’s identity.
	// - Or, if the request is issued by an API key, i.e. an API key invalidates
	// itself, specify its ID in the `ids` field.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-invalidate-api-key.html
	InvalidateApiKey security_invalidate_api_key.NewInvalidateApiKey
	// Invalidates one or more access tokens or refresh tokens.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-invalidate-token.html
	InvalidateToken security_invalidate_token.NewInvalidateToken
	// Exchanges an OpenID Connection authentication response message for an
	// Elasticsearch access token and refresh token pair
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-oidc-authenticate.html
	OidcAuthenticate security_oidc_authenticate.NewOidcAuthenticate
	// Invalidates a refresh token and access token that was generated from the
	// OpenID Connect Authenticate API
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-oidc-logout.html
	OidcLogout security_oidc_logout.NewOidcLogout
	// Creates an OAuth 2.0 authentication request as a URL string
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-oidc-prepare-authentication.html
	OidcPrepareAuthentication security_oidc_prepare_authentication.NewOidcPrepareAuthentication
	// Adds or updates application privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-privileges.html
	PutPrivileges security_put_privileges.NewPutPrivileges
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The create or update roles API cannot update roles that are defined in roles
	// files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role.html
	PutRole security_put_role.NewPutRole
	// Creates and updates role mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role-mapping.html
	PutRoleMapping security_put_role_mapping.NewPutRoleMapping
	// Adds and updates users in the native realm. These users are commonly referred
	// to as native users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html
	PutUser security_put_user.NewPutUser
	// Retrieves information for API keys in a paginated manner. You can optionally
	// filter the results with a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-api-key.html
	QueryApiKeys security_query_api_keys.NewQueryApiKeys
	// Retrieves roles in a paginated manner. You can optionally filter the results
	// with a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-role.html
	QueryRole security_query_role.NewQueryRole
	// Retrieves information for Users in a paginated manner. You can optionally
	// filter the results with a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-user.html
	QueryUser security_query_user.NewQueryUser
	// Submits a SAML Response message to Elasticsearch for consumption.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-authenticate.html
	SamlAuthenticate security_saml_authenticate.NewSamlAuthenticate
	// Verifies the logout response sent from the SAML IdP.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-complete-logout.html
	SamlCompleteLogout security_saml_complete_logout.NewSamlCompleteLogout
	// Submits a SAML LogoutRequest message to Elasticsearch for consumption.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-invalidate.html
	SamlInvalidate security_saml_invalidate.NewSamlInvalidate
	// Submits a request to invalidate an access token and refresh token.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-logout.html
	SamlLogout security_saml_logout.NewSamlLogout
	// Creates a SAML authentication request (<AuthnRequest>) as a URL string, based
	// on the configuration of the respective SAML realm in Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-prepare-authentication.html
	SamlPrepareAuthentication security_saml_prepare_authentication.NewSamlPrepareAuthentication
	// Generate SAML metadata for a SAML 2.0 Service Provider.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-sp-metadata.html
	SamlServiceProviderMetadata security_saml_service_provider_metadata.NewSamlServiceProviderMetadata
	// Get suggestions for user profiles that match specified search criteria.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-suggest-user-profile.html
	SuggestUserProfiles security_suggest_user_profiles.NewSuggestUserProfiles
	// Updates attributes of an existing API key.
	// Users can only update API keys that they created or that were granted to
	// them.
	// Use this API to update API keys created by the create API Key or grant API
	// Key APIs.
	// If you need to apply the same update to many API keys, you can use bulk
	// update API Keys to reduce overhead.
	// It’s not possible to update expired API keys, or API keys that have been
	// invalidated by invalidate API Key.
	// This API supports updates to an API key’s access scope and metadata.
	// The access scope of an API key is derived from the `role_descriptors` you
	// specify in the request, and a snapshot of the owner user’s permissions at the
	// time of the request.
	// The snapshot of the owner’s permissions is updated automatically on every
	// call.
	// If you don’t specify `role_descriptors` in the request, a call to this API
	// might still change the API key’s access scope.
	// This change can occur if the owner user’s permissions have changed since the
	// API key was created or last modified.
	// To update another user’s API key, use the `run_as` feature to submit a
	// request on behalf of another user.
	// IMPORTANT: It’s not possible to use an API key as the authentication
	// credential for this API.
	// To update an API key, the owner user’s credentials are required.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-api-key.html
	UpdateApiKey security_update_api_key.NewUpdateApiKey
	// Update settings for the security system index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-settings.html
	UpdateSettings security_update_settings.NewUpdateSettings
	// Updates specific data for the user profile that's associated with the
	// specified unique ID.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-user-profile-data.html
	UpdateUserProfileData security_update_user_profile_data.NewUpdateUserProfileData
}

type Shutdown struct {
	// Removes a node from the shutdown list. Designed for indirect use by ECE/ESS
	// and ECK. Direct use is not supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current
	DeleteNode shutdown_delete_node.NewDeleteNode
	// Retrieve status of a node or nodes that are currently marked as shutting
	// down. Designed for indirect use by ECE/ESS and ECK. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current
	GetNode shutdown_get_node.NewGetNode
	// Adds a node to be shut down. Designed for indirect use by ECE/ESS and ECK.
	// Direct use is not supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current
	PutNode shutdown_put_node.NewPutNode
}

type Slm struct {
	// Deletes an existing snapshot lifecycle policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-delete-policy.html
	DeleteLifecycle slm_delete_lifecycle.NewDeleteLifecycle
	// Immediately creates a snapshot according to the lifecycle policy, without
	// waiting for the scheduled time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-execute-lifecycle.html
	ExecuteLifecycle slm_execute_lifecycle.NewExecuteLifecycle
	// Deletes any snapshots that are expired according to the policy's retention
	// rules.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-execute-retention.html
	ExecuteRetention slm_execute_retention.NewExecuteRetention
	// Retrieves one or more snapshot lifecycle policy definitions and information
	// about the latest snapshot attempts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-get-policy.html
	GetLifecycle slm_get_lifecycle.NewGetLifecycle
	// Returns global and policy-level statistics about actions taken by snapshot
	// lifecycle management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-get-stats.html
	GetStats slm_get_stats.NewGetStats
	// Retrieves the status of snapshot lifecycle management (SLM).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-get-status.html
	GetStatus slm_get_status.NewGetStatus
	// Creates or updates a snapshot lifecycle policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-put-policy.html
	PutLifecycle slm_put_lifecycle.NewPutLifecycle
	// Turns on snapshot lifecycle management (SLM).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-start.html
	Start slm_start.NewStart
	// Turns off snapshot lifecycle management (SLM).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-stop.html
	Stop slm_stop.NewStop
}

type Snapshot struct {
	// Triggers the review of a snapshot repository’s contents and deletes any stale
	// data not referenced by existing snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clean-up-snapshot-repo-api.html
	CleanupRepository snapshot_cleanup_repository.NewCleanupRepository
	// Clones indices from one snapshot into another snapshot in the same
	// repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	Clone snapshot_clone.NewClone
	// Creates a snapshot in a repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	Create snapshot_create.NewCreate
	// Creates a repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	CreateRepository snapshot_create_repository.NewCreateRepository
	// Deletes one or more snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	Delete snapshot_delete.NewDelete
	// Deletes a repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	DeleteRepository snapshot_delete_repository.NewDeleteRepository
	// Returns information about a snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	Get snapshot_get.NewGet
	// Returns information about a repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	GetRepository snapshot_get_repository.NewGetRepository
	// Restores a snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	Restore snapshot_restore.NewRestore
	// Returns information about the status of a snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	Status snapshot_status.NewStatus
	// Verifies a repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	VerifyRepository snapshot_verify_repository.NewVerifyRepository
}

type Sql struct {
	// Clears the SQL cursor
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-sql-cursor-api.html
	ClearCursor sql_clear_cursor.NewClearCursor
	// Deletes an async SQL search or a stored synchronous SQL search. If the search
	// is still running, the API cancels it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-async-sql-search-api.html
	DeleteAsync sql_delete_async.NewDeleteAsync
	// Returns the current status and available results for an async SQL search or
	// stored synchronous SQL search
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-sql-search-api.html
	GetAsync sql_get_async.NewGetAsync
	// Returns the current status of an async SQL search or a stored synchronous SQL
	// search
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-sql-search-status-api.html
	GetAsyncStatus sql_get_async_status.NewGetAsyncStatus
	// Executes a SQL request
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-search-api.html
	Query sql_query.NewQuery
	// Translates SQL into Elasticsearch queries
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-translate-api.html
	Translate sql_translate.NewTranslate
}

type Ssl struct {
	// Retrieves information about the X.509 certificates used to encrypt
	// communications in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-ssl.html
	Certificates ssl_certificates.NewCertificates
}

type Synonyms struct {
	// Deletes a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-synonyms-set.html
	DeleteSynonym synonyms_delete_synonym.NewDeleteSynonym
	// Deletes a synonym rule in a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-synonym-rule.html
	DeleteSynonymRule synonyms_delete_synonym_rule.NewDeleteSynonymRule
	// Retrieves a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonyms-set.html
	GetSynonym synonyms_get_synonym.NewGetSynonym
	// Retrieves a synonym rule from a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonym-rule.html
	GetSynonymRule synonyms_get_synonym_rule.NewGetSynonymRule
	// Retrieves a summary of all defined synonym sets
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-synonyms-sets.html
	GetSynonymsSets synonyms_get_synonyms_sets.NewGetSynonymsSets
	// Creates or updates a synonym set.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonyms-set.html
	PutSynonym synonyms_put_synonym.NewPutSynonym
	// Creates or updates a synonym rule in a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonym-rule.html
	PutSynonymRule synonyms_put_synonym_rule.NewPutSynonymRule
}

type Tasks struct {
	// Cancels a task, if it can be cancelled through an API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	Cancel tasks_cancel.NewCancel
	// Returns information about a task.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	Get tasks_get.NewGet
	// The task management API returns information about tasks currently executing
	// on one or more nodes in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	List tasks_list.NewList
}

type TextStructure struct {
	// Finds the structure of a text field in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-field-structure.html
	FindFieldStructure text_structure_find_field_structure.NewFindFieldStructure
	// Finds the structure of a list of messages. The messages must contain data
	// that is suitable to be ingested into Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-message-structure.html
	FindMessageStructure text_structure_find_message_structure.NewFindMessageStructure
	// Finds the structure of a text file. The text file must contain data that is
	// suitable to be ingested into Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-structure.html
	FindStructure text_structure_find_structure.NewFindStructure
	// Tests a Grok pattern on some text.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/test-grok-pattern.html
	TestGrokPattern text_structure_test_grok_pattern.NewTestGrokPattern
}

type Transform struct {
	// Deletes a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-transform.html
	DeleteTransform transform_delete_transform.NewDeleteTransform
	// Retrieves transform usage information for transform nodes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-node-stats.html
	GetNodeStats transform_get_node_stats.NewGetNodeStats
	// Retrieves configuration information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform.html
	GetTransform transform_get_transform.NewGetTransform
	// Retrieves usage information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-stats.html
	GetTransformStats transform_get_transform_stats.NewGetTransformStats
	// Previews a transform.
	//
	// It returns a maximum of 100 results. The calculations are based on all the
	// current data in the source index. It also
	// generates a list of mappings and settings for the destination index. These
	// values are determined based on the field
	// types of the source index and the transform aggregations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-transform.html
	PreviewTransform transform_preview_transform.NewPreviewTransform
	// Creates a transform.
	//
	// A transform copies data from source indices, transforms it, and persists it
	// into an entity-centric destination index. You can also think of the
	// destination index as a two-dimensional tabular data structure (known as
	// a data frame). The ID for each document in the data frame is generated from a
	// hash of the entity, so there is a
	// unique row per entity.
	//
	// You must choose either the latest or pivot method for your transform; you
	// cannot use both in a single transform. If
	// you choose to use the pivot method for your transform, the entities are
	// defined by the set of `group_by` fields in
	// the pivot object. If you choose to use the latest method, the entities are
	// defined by the `unique_key` field values
	// in the latest object.
	//
	// You must have `create_index`, `index`, and `read` privileges on the
	// destination index and `read` and
	// `view_index_metadata` privileges on the source indices. When Elasticsearch
	// security features are enabled, the
	// transform remembers which roles the user that created it had at the time of
	// creation and uses those same roles. If
	// those roles do not have the required privileges on the source and destination
	// indices, the transform fails when it
	// attempts unauthorized operations.
	//
	// NOTE: You must use Kibana or this API to create a transform. Do not add a
	// transform directly into any
	// `.transform-internal*` indices using the Elasticsearch index API. If
	// Elasticsearch security features are enabled, do
	// not give users any privileges on `.transform-internal*` indices. If you used
	// transforms prior to 7.5, also do not
	// give users any privileges on `.data-frame-internal*` indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-transform.html
	PutTransform transform_put_transform.NewPutTransform
	// Resets a transform.
	// Before you can reset it, you must stop it; alternatively, use the `force`
	// query parameter.
	// If the destination index was created by the transform, it is deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-transform.html
	ResetTransform transform_reset_transform.NewResetTransform
	// Schedules now a transform.
	//
	// If you _schedule_now a transform, it will process the new data instantly,
	// without waiting for the configured frequency interval. After _schedule_now
	// API is called,
	// the transform will be processed again at now + frequency unless _schedule_now
	// API
	// is called again in the meantime.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/schedule-now-transform.html
	ScheduleNowTransform transform_schedule_now_transform.NewScheduleNowTransform
	// Starts a transform.
	//
	// When you start a transform, it creates the destination index if it does not
	// already exist. The `number_of_shards` is
	// set to `1` and the `auto_expand_replicas` is set to `0-1`. If it is a pivot
	// transform, it deduces the mapping
	// definitions for the destination index from the source indices and the
	// transform aggregations. If fields in the
	// destination index are derived from scripts (as in the case of
	// `scripted_metric` or `bucket_script` aggregations),
	// the transform uses dynamic mappings unless an index template exists. If it is
	// a latest transform, it does not deduce
	// mapping definitions; it uses dynamic mappings. To use explicit mappings,
	// create the destination index before you
	// start the transform. Alternatively, you can create an index template, though
	// it does not affect the deduced mappings
	// in a pivot transform.
	//
	// When the transform starts, a series of validations occur to ensure its
	// success. If you deferred validation when you
	// created the transform, they occur when you start the transform—​with the
	// exception of privilege checks. When
	// Elasticsearch security features are enabled, the transform remembers which
	// roles the user that created it had at the
	// time of creation and uses those same roles. If those roles do not have the
	// required privileges on the source and
	// destination indices, the transform fails when it attempts unauthorized
	// operations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-transform.html
	StartTransform transform_start_transform.NewStartTransform
	// Stops one or more transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-transform.html
	StopTransform transform_stop_transform.NewStopTransform
	// Updates certain properties of a transform.
	//
	// All updated properties except `description` do not take effect until after
	// the transform starts the next checkpoint,
	// thus there is data consistency in each checkpoint. To use this API, you must
	// have `read` and `view_index_metadata`
	// privileges for the source indices. You must also have `index` and `read`
	// privileges for the destination index. When
	// Elasticsearch security features are enabled, the transform remembers which
	// roles the user who updated it had at the
	// time of update and runs with those privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-transform.html
	UpdateTransform transform_update_transform.NewUpdateTransform
	// Upgrades all transforms.
	// This API identifies transforms that have a legacy configuration format and
	// upgrades them to the latest version. It
	// also cleans up the internal data structures that store the transform state
	// and checkpoints. The upgrade does not
	// affect the source and destination indices. The upgrade also does not affect
	// the roles that transforms use when
	// Elasticsearch security features are enabled; the role used to read source
	// data and write to the destination index
	// remains unchanged.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/upgrade-transforms.html
	UpgradeTransforms transform_upgrade_transforms.NewUpgradeTransforms
}

type Watcher struct {
	// Acknowledges a watch, manually throttling the execution of the watch's
	// actions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-ack-watch.html
	AckWatch watcher_ack_watch.NewAckWatch
	// Activates a currently inactive watch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-activate-watch.html
	ActivateWatch watcher_activate_watch.NewActivateWatch
	// Deactivates a currently active watch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-deactivate-watch.html
	DeactivateWatch watcher_deactivate_watch.NewDeactivateWatch
	// Removes a watch from Watcher.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-delete-watch.html
	DeleteWatch watcher_delete_watch.NewDeleteWatch
	// This API can be used to force execution of the watch outside of its
	// triggering logic or to simulate the watch execution for debugging purposes.
	// For testing and debugging purposes, you also have fine-grained control on how
	// the watch runs. You can execute the watch without executing all of its
	// actions or alternatively by simulating them. You can also force execution by
	// ignoring the watch condition and control whether a watch record would be
	// written to the watch history after execution.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-execute-watch.html
	ExecuteWatch watcher_execute_watch.NewExecuteWatch
	// Retrieve settings for the watcher system index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-get-settings.html
	GetSettings watcher_get_settings.NewGetSettings
	// Retrieves a watch by its ID.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-get-watch.html
	GetWatch watcher_get_watch.NewGetWatch
	// Creates a new watch, or updates an existing one.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-put-watch.html
	PutWatch watcher_put_watch.NewPutWatch
	// Retrieves stored watches.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-query-watches.html
	QueryWatches watcher_query_watches.NewQueryWatches
	// Starts Watcher if it is not already running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-start.html
	Start watcher_start.NewStart
	// Retrieves the current Watcher metrics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-stats.html
	Stats watcher_stats.NewStats
	// Stops Watcher if it is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-stop.html
	Stop watcher_stop.NewStop
	// Update settings for the watcher system index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-update-settings.html
	UpdateSettings watcher_update_settings.NewUpdateSettings
}

type Xpack struct {
	// Provides general information about the installed X-Pack features.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html
	Info xpack_info.NewInfo
	// This API provides information about which features are currently enabled and
	// available under the current license and some usage statistics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/usage-api.html
	Usage xpack_usage.NewUsage
}

type API struct {
	AsyncSearch         AsyncSearch
	Autoscaling         Autoscaling
	Capabilities        Capabilities
	Cat                 Cat
	Ccr                 Ccr
	Cluster             Cluster
	Connector           Connector
	Core                Core
	DanglingIndices     DanglingIndices
	Enrich              Enrich
	Eql                 Eql
	Esql                Esql
	Features            Features
	Fleet               Fleet
	Graph               Graph
	Ilm                 Ilm
	Indices             Indices
	Inference           Inference
	Ingest              Ingest
	License             License
	Logstash            Logstash
	Migration           Migration
	Ml                  Ml
	Monitoring          Monitoring
	Nodes               Nodes
	Profiling           Profiling
	QueryRules          QueryRules
	Rollup              Rollup
	SearchApplication   SearchApplication
	SearchableSnapshots SearchableSnapshots
	Security            Security
	Shutdown            Shutdown
	Slm                 Slm
	Snapshot            Snapshot
	Sql                 Sql
	Ssl                 Ssl
	Synonyms            Synonyms
	Tasks               Tasks
	TextStructure       TextStructure
	Transform           Transform
	Watcher             Watcher
	Xpack               Xpack

	// Performs multiple indexing or delete operations in a single API call.
	// This reduces overhead and can greatly increase indexing speed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Clears the search context and results for a scrolling search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Closes a point-in-time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Removes a JSON document from the specified index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Deletes documents that match the specified query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Changes the number of requests per second for a particular Delete By Query
	// operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	// Deletes a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Checks if a document in an index exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Checks if a document's `_source` is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Returns information about why a specific document matches (or doesn’t match)
	// a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// The field capabilities API returns the information about the capabilities of
	// fields among multiple indices.
	// The field capabilities API returns runtime fields like any other field. For
	// example, a runtime field with a type
	// of keyword is returned as any other field that belongs to the `keyword`
	// family.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Returns a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Retrieves a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScript core_get_script.NewGetScript
	// Returns all script contexts.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-contexts.html
	GetScriptContext core_get_script_context.NewGetScriptContext
	// Returns available script types, languages and contexts
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScriptLanguages core_get_script_languages.NewGetScriptLanguages
	// Returns the source of a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Returns the health of the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html
	HealthReport core_health_report.NewHealthReport
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Index core_index.NewIndex
	// Returns basic information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Info core_info.NewInfo
	// Performs a kNN search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	KnnSearch core_knn_search.NewKnnSearch
	// Allows to get multiple documents in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
	Mget core_mget.NewMget
	// Allows to execute several search operations in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	Msearch core_msearch.NewMsearch
	// Runs multiple templated searches with a single request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Returns multiple termvectors in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// A search request by default executes against the most recent visible data of
	// the target indices,
	// which is called point in time. Elasticsearch pit (point in time) is a
	// lightweight view into the
	// state of the data as it existed when initiated. In some cases, it’s preferred
	// to perform multiple
	// search requests using the same point in time. For example, if refreshes
	// happen between
	// `search_after` requests, then the results of those requests might not be
	// consistent as changes happening
	// between searches are only visible to the more recent point in time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Creates or updates a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Enables you to evaluate the quality of ranked search results over a set of
	// typical search queries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Allows to copy documents from one index to another, optionally filtering the
	// source
	// documents by a query, changing the destination index settings, or fetching
	// the
	// documents from a remote cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Copies documents from a source to a destination.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle
	// Renders a search template as a search request body.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Runs a script and returns a result.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Allows to retrieve a large numbers of results from a single search request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Returns search hits that match the query defined in the request.
	// You can provide search queries using the `q` query string parameter or the
	// request body.
	// If both are specified, only the query parameter is used.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Searches a vector tile for geospatial values. Returns results as a binary
	// Mapbox vector tile.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Returns information about the indices and shards that a search request would
	// be executed against.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-shards.html
	SearchShards core_search_shards.NewSearchShards
	// Runs a search with a search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
	SearchTemplate core_search_template.NewSearchTemplate
	// The terms enum API  can be used to discover terms in the index that begin
	// with the provided string. It is designed for low-latency look-ups used in
	// auto-complete scenarios.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
	TermsEnum core_terms_enum.NewTermsEnum
	// Returns information and statistics about terms in the fields of a particular
	// document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
	Termvectors core_termvectors.NewTermvectors
	// Updates a document with a script or partial document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
	Update core_update.NewUpdate
	// Updates documents that match the specified query.
	// If no query is specified, performs an update on every document in the data
	// stream or index without modifying the source, which is useful for picking up
	// mapping changes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQuery core_update_by_query.NewUpdateByQuery
	// Changes the number of requests per second for a particular Update By Query
	// operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

func New(tp elastictransport.Interface) *API {
	return &API{
		// AsyncSearch
		AsyncSearch: AsyncSearch{
			Delete: async_search_delete.NewDeleteFunc(tp),
			Get:    async_search_get.NewGetFunc(tp),
			Status: async_search_status.NewStatusFunc(tp),
			Submit: async_search_submit.NewSubmitFunc(tp),
		},

		// Autoscaling
		Autoscaling: Autoscaling{
			DeleteAutoscalingPolicy: autoscaling_delete_autoscaling_policy.NewDeleteAutoscalingPolicyFunc(tp),
			GetAutoscalingCapacity:  autoscaling_get_autoscaling_capacity.NewGetAutoscalingCapacityFunc(tp),
			GetAutoscalingPolicy:    autoscaling_get_autoscaling_policy.NewGetAutoscalingPolicyFunc(tp),
			PutAutoscalingPolicy:    autoscaling_put_autoscaling_policy.NewPutAutoscalingPolicyFunc(tp),
		},

		// Capabilities
		Capabilities: Capabilities{
			Capabilities: capabilities.NewCapabilitiesFunc(tp),
		},

		// Cat
		Cat: Cat{
			Aliases:              cat_aliases.NewAliasesFunc(tp),
			Allocation:           cat_allocation.NewAllocationFunc(tp),
			ComponentTemplates:   cat_component_templates.NewComponentTemplatesFunc(tp),
			Count:                cat_count.NewCountFunc(tp),
			Fielddata:            cat_fielddata.NewFielddataFunc(tp),
			Health:               cat_health.NewHealthFunc(tp),
			Help:                 cat_help.NewHelpFunc(tp),
			Indices:              cat_indices.NewIndicesFunc(tp),
			Master:               cat_master.NewMasterFunc(tp),
			MlDataFrameAnalytics: cat_ml_data_frame_analytics.NewMlDataFrameAnalyticsFunc(tp),
			MlDatafeeds:          cat_ml_datafeeds.NewMlDatafeedsFunc(tp),
			MlJobs:               cat_ml_jobs.NewMlJobsFunc(tp),
			MlTrainedModels:      cat_ml_trained_models.NewMlTrainedModelsFunc(tp),
			Nodeattrs:            cat_nodeattrs.NewNodeattrsFunc(tp),
			Nodes:                cat_nodes.NewNodesFunc(tp),
			PendingTasks:         cat_pending_tasks.NewPendingTasksFunc(tp),
			Plugins:              cat_plugins.NewPluginsFunc(tp),
			Recovery:             cat_recovery.NewRecoveryFunc(tp),
			Repositories:         cat_repositories.NewRepositoriesFunc(tp),
			Segments:             cat_segments.NewSegmentsFunc(tp),
			Shards:               cat_shards.NewShardsFunc(tp),
			Snapshots:            cat_snapshots.NewSnapshotsFunc(tp),
			Tasks:                cat_tasks.NewTasksFunc(tp),
			Templates:            cat_templates.NewTemplatesFunc(tp),
			ThreadPool:           cat_thread_pool.NewThreadPoolFunc(tp),
			Transforms:           cat_transforms.NewTransformsFunc(tp),
		},

		// Ccr
		Ccr: Ccr{
			DeleteAutoFollowPattern: ccr_delete_auto_follow_pattern.NewDeleteAutoFollowPatternFunc(tp),
			Follow:                  ccr_follow.NewFollowFunc(tp),
			FollowInfo:              ccr_follow_info.NewFollowInfoFunc(tp),
			FollowStats:             ccr_follow_stats.NewFollowStatsFunc(tp),
			ForgetFollower:          ccr_forget_follower.NewForgetFollowerFunc(tp),
			GetAutoFollowPattern:    ccr_get_auto_follow_pattern.NewGetAutoFollowPatternFunc(tp),
			PauseAutoFollowPattern:  ccr_pause_auto_follow_pattern.NewPauseAutoFollowPatternFunc(tp),
			PauseFollow:             ccr_pause_follow.NewPauseFollowFunc(tp),
			PutAutoFollowPattern:    ccr_put_auto_follow_pattern.NewPutAutoFollowPatternFunc(tp),
			ResumeAutoFollowPattern: ccr_resume_auto_follow_pattern.NewResumeAutoFollowPatternFunc(tp),
			ResumeFollow:            ccr_resume_follow.NewResumeFollowFunc(tp),
			Stats:                   ccr_stats.NewStatsFunc(tp),
			Unfollow:                ccr_unfollow.NewUnfollowFunc(tp),
		},

		// Cluster
		Cluster: Cluster{
			AllocationExplain:            cluster_allocation_explain.NewAllocationExplainFunc(tp),
			DeleteComponentTemplate:      cluster_delete_component_template.NewDeleteComponentTemplateFunc(tp),
			DeleteVotingConfigExclusions: cluster_delete_voting_config_exclusions.NewDeleteVotingConfigExclusionsFunc(tp),
			ExistsComponentTemplate:      cluster_exists_component_template.NewExistsComponentTemplateFunc(tp),
			GetComponentTemplate:         cluster_get_component_template.NewGetComponentTemplateFunc(tp),
			GetSettings:                  cluster_get_settings.NewGetSettingsFunc(tp),
			Health:                       cluster_health.NewHealthFunc(tp),
			Info:                         cluster_info.NewInfoFunc(tp),
			PendingTasks:                 cluster_pending_tasks.NewPendingTasksFunc(tp),
			PostVotingConfigExclusions:   cluster_post_voting_config_exclusions.NewPostVotingConfigExclusionsFunc(tp),
			PutComponentTemplate:         cluster_put_component_template.NewPutComponentTemplateFunc(tp),
			PutSettings:                  cluster_put_settings.NewPutSettingsFunc(tp),
			RemoteInfo:                   cluster_remote_info.NewRemoteInfoFunc(tp),
			Reroute:                      cluster_reroute.NewRerouteFunc(tp),
			State:                        cluster_state.NewStateFunc(tp),
			Stats:                        cluster_stats.NewStatsFunc(tp),
		},

		// Connector
		Connector: Connector{
			CheckIn:                   connector_check_in.NewCheckInFunc(tp),
			Delete:                    connector_delete.NewDeleteFunc(tp),
			Get:                       connector_get.NewGetFunc(tp),
			LastSync:                  connector_last_sync.NewLastSyncFunc(tp),
			List:                      connector_list.NewListFunc(tp),
			Post:                      connector_post.NewPostFunc(tp),
			Put:                       connector_put.NewPutFunc(tp),
			SecretPost:                connector_secret_post.NewSecretPostFunc(tp),
			SyncJobCancel:             connector_sync_job_cancel.NewSyncJobCancelFunc(tp),
			SyncJobDelete:             connector_sync_job_delete.NewSyncJobDeleteFunc(tp),
			SyncJobGet:                connector_sync_job_get.NewSyncJobGetFunc(tp),
			SyncJobList:               connector_sync_job_list.NewSyncJobListFunc(tp),
			SyncJobPost:               connector_sync_job_post.NewSyncJobPostFunc(tp),
			UpdateActiveFiltering:     connector_update_active_filtering.NewUpdateActiveFilteringFunc(tp),
			UpdateApiKeyId:            connector_update_api_key_id.NewUpdateApiKeyIdFunc(tp),
			UpdateConfiguration:       connector_update_configuration.NewUpdateConfigurationFunc(tp),
			UpdateError:               connector_update_error.NewUpdateErrorFunc(tp),
			UpdateFiltering:           connector_update_filtering.NewUpdateFilteringFunc(tp),
			UpdateFilteringValidation: connector_update_filtering_validation.NewUpdateFilteringValidationFunc(tp),
			UpdateIndexName:           connector_update_index_name.NewUpdateIndexNameFunc(tp),
			UpdateName:                connector_update_name.NewUpdateNameFunc(tp),
			UpdateNative:              connector_update_native.NewUpdateNativeFunc(tp),
			UpdatePipeline:            connector_update_pipeline.NewUpdatePipelineFunc(tp),
			UpdateScheduling:          connector_update_scheduling.NewUpdateSchedulingFunc(tp),
			UpdateServiceType:         connector_update_service_type.NewUpdateServiceTypeFunc(tp),
			UpdateStatus:              connector_update_status.NewUpdateStatusFunc(tp),
		},

		// Core
		Core: Core{
			Bulk:                    core_bulk.NewBulkFunc(tp),
			ClearScroll:             core_clear_scroll.NewClearScrollFunc(tp),
			ClosePointInTime:        core_close_point_in_time.NewClosePointInTimeFunc(tp),
			Count:                   core_count.NewCountFunc(tp),
			Create:                  core_create.NewCreateFunc(tp),
			Delete:                  core_delete.NewDeleteFunc(tp),
			DeleteByQuery:           core_delete_by_query.NewDeleteByQueryFunc(tp),
			DeleteByQueryRethrottle: core_delete_by_query_rethrottle.NewDeleteByQueryRethrottleFunc(tp),
			DeleteScript:            core_delete_script.NewDeleteScriptFunc(tp),
			Exists:                  core_exists.NewExistsFunc(tp),
			ExistsSource:            core_exists_source.NewExistsSourceFunc(tp),
			Explain:                 core_explain.NewExplainFunc(tp),
			FieldCaps:               core_field_caps.NewFieldCapsFunc(tp),
			Get:                     core_get.NewGetFunc(tp),
			GetScript:               core_get_script.NewGetScriptFunc(tp),
			GetScriptContext:        core_get_script_context.NewGetScriptContextFunc(tp),
			GetScriptLanguages:      core_get_script_languages.NewGetScriptLanguagesFunc(tp),
			GetSource:               core_get_source.NewGetSourceFunc(tp),
			HealthReport:            core_health_report.NewHealthReportFunc(tp),
			Index:                   core_index.NewIndexFunc(tp),
			Info:                    core_info.NewInfoFunc(tp),
			KnnSearch:               core_knn_search.NewKnnSearchFunc(tp),
			Mget:                    core_mget.NewMgetFunc(tp),
			Msearch:                 core_msearch.NewMsearchFunc(tp),
			MsearchTemplate:         core_msearch_template.NewMsearchTemplateFunc(tp),
			Mtermvectors:            core_mtermvectors.NewMtermvectorsFunc(tp),
			OpenPointInTime:         core_open_point_in_time.NewOpenPointInTimeFunc(tp),
			Ping:                    core_ping.NewPingFunc(tp),
			PutScript:               core_put_script.NewPutScriptFunc(tp),
			RankEval:                core_rank_eval.NewRankEvalFunc(tp),
			Reindex:                 core_reindex.NewReindexFunc(tp),
			ReindexRethrottle:       core_reindex_rethrottle.NewReindexRethrottleFunc(tp),
			RenderSearchTemplate:    core_render_search_template.NewRenderSearchTemplateFunc(tp),
			ScriptsPainlessExecute:  core_scripts_painless_execute.NewScriptsPainlessExecuteFunc(tp),
			Scroll:                  core_scroll.NewScrollFunc(tp),
			Search:                  core_search.NewSearchFunc(tp),
			SearchMvt:               core_search_mvt.NewSearchMvtFunc(tp),
			SearchShards:            core_search_shards.NewSearchShardsFunc(tp),
			SearchTemplate:          core_search_template.NewSearchTemplateFunc(tp),
			TermsEnum:               core_terms_enum.NewTermsEnumFunc(tp),
			Termvectors:             core_termvectors.NewTermvectorsFunc(tp),
			Update:                  core_update.NewUpdateFunc(tp),
			UpdateByQuery:           core_update_by_query.NewUpdateByQueryFunc(tp),
			UpdateByQueryRethrottle: core_update_by_query_rethrottle.NewUpdateByQueryRethrottleFunc(tp),
		},

		// DanglingIndices
		DanglingIndices: DanglingIndices{
			DeleteDanglingIndex: dangling_indices_delete_dangling_index.NewDeleteDanglingIndexFunc(tp),
			ImportDanglingIndex: dangling_indices_import_dangling_index.NewImportDanglingIndexFunc(tp),
			ListDanglingIndices: dangling_indices_list_dangling_indices.NewListDanglingIndicesFunc(tp),
		},

		// Enrich
		Enrich: Enrich{
			DeletePolicy:  enrich_delete_policy.NewDeletePolicyFunc(tp),
			ExecutePolicy: enrich_execute_policy.NewExecutePolicyFunc(tp),
			GetPolicy:     enrich_get_policy.NewGetPolicyFunc(tp),
			PutPolicy:     enrich_put_policy.NewPutPolicyFunc(tp),
			Stats:         enrich_stats.NewStatsFunc(tp),
		},

		// Eql
		Eql: Eql{
			Delete:    eql_delete.NewDeleteFunc(tp),
			Get:       eql_get.NewGetFunc(tp),
			GetStatus: eql_get_status.NewGetStatusFunc(tp),
			Search:    eql_search.NewSearchFunc(tp),
		},

		// Esql
		Esql: Esql{
			AsyncQuery: esql_async_query.NewAsyncQueryFunc(tp),
			Query:      esql_query.NewQueryFunc(tp),
		},

		// Features
		Features: Features{
			GetFeatures:   features_get_features.NewGetFeaturesFunc(tp),
			ResetFeatures: features_reset_features.NewResetFeaturesFunc(tp),
		},

		// Fleet
		Fleet: Fleet{
			GlobalCheckpoints: fleet_global_checkpoints.NewGlobalCheckpointsFunc(tp),
			Msearch:           fleet_msearch.NewMsearchFunc(tp),
			PostSecret:        fleet_post_secret.NewPostSecretFunc(tp),
			Search:            fleet_search.NewSearchFunc(tp),
		},

		// Graph
		Graph: Graph{
			Explore: graph_explore.NewExploreFunc(tp),
		},

		// Ilm
		Ilm: Ilm{
			DeleteLifecycle:    ilm_delete_lifecycle.NewDeleteLifecycleFunc(tp),
			ExplainLifecycle:   ilm_explain_lifecycle.NewExplainLifecycleFunc(tp),
			GetLifecycle:       ilm_get_lifecycle.NewGetLifecycleFunc(tp),
			GetStatus:          ilm_get_status.NewGetStatusFunc(tp),
			MigrateToDataTiers: ilm_migrate_to_data_tiers.NewMigrateToDataTiersFunc(tp),
			MoveToStep:         ilm_move_to_step.NewMoveToStepFunc(tp),
			PutLifecycle:       ilm_put_lifecycle.NewPutLifecycleFunc(tp),
			RemovePolicy:       ilm_remove_policy.NewRemovePolicyFunc(tp),
			Retry:              ilm_retry.NewRetryFunc(tp),
			Start:              ilm_start.NewStartFunc(tp),
			Stop:               ilm_stop.NewStopFunc(tp),
		},

		// Indices
		Indices: Indices{
			AddBlock:              indices_add_block.NewAddBlockFunc(tp),
			Analyze:               indices_analyze.NewAnalyzeFunc(tp),
			ClearCache:            indices_clear_cache.NewClearCacheFunc(tp),
			Clone:                 indices_clone.NewCloneFunc(tp),
			Close:                 indices_close.NewCloseFunc(tp),
			Create:                indices_create.NewCreateFunc(tp),
			CreateDataStream:      indices_create_data_stream.NewCreateDataStreamFunc(tp),
			DataStreamsStats:      indices_data_streams_stats.NewDataStreamsStatsFunc(tp),
			Delete:                indices_delete.NewDeleteFunc(tp),
			DeleteAlias:           indices_delete_alias.NewDeleteAliasFunc(tp),
			DeleteDataLifecycle:   indices_delete_data_lifecycle.NewDeleteDataLifecycleFunc(tp),
			DeleteDataStream:      indices_delete_data_stream.NewDeleteDataStreamFunc(tp),
			DeleteIndexTemplate:   indices_delete_index_template.NewDeleteIndexTemplateFunc(tp),
			DeleteTemplate:        indices_delete_template.NewDeleteTemplateFunc(tp),
			DiskUsage:             indices_disk_usage.NewDiskUsageFunc(tp),
			Downsample:            indices_downsample.NewDownsampleFunc(tp),
			Exists:                indices_exists.NewExistsFunc(tp),
			ExistsAlias:           indices_exists_alias.NewExistsAliasFunc(tp),
			ExistsIndexTemplate:   indices_exists_index_template.NewExistsIndexTemplateFunc(tp),
			ExistsTemplate:        indices_exists_template.NewExistsTemplateFunc(tp),
			ExplainDataLifecycle:  indices_explain_data_lifecycle.NewExplainDataLifecycleFunc(tp),
			FieldUsageStats:       indices_field_usage_stats.NewFieldUsageStatsFunc(tp),
			Flush:                 indices_flush.NewFlushFunc(tp),
			Forcemerge:            indices_forcemerge.NewForcemergeFunc(tp),
			Get:                   indices_get.NewGetFunc(tp),
			GetAlias:              indices_get_alias.NewGetAliasFunc(tp),
			GetDataLifecycle:      indices_get_data_lifecycle.NewGetDataLifecycleFunc(tp),
			GetDataStream:         indices_get_data_stream.NewGetDataStreamFunc(tp),
			GetFieldMapping:       indices_get_field_mapping.NewGetFieldMappingFunc(tp),
			GetIndexTemplate:      indices_get_index_template.NewGetIndexTemplateFunc(tp),
			GetMapping:            indices_get_mapping.NewGetMappingFunc(tp),
			GetSettings:           indices_get_settings.NewGetSettingsFunc(tp),
			GetTemplate:           indices_get_template.NewGetTemplateFunc(tp),
			MigrateToDataStream:   indices_migrate_to_data_stream.NewMigrateToDataStreamFunc(tp),
			ModifyDataStream:      indices_modify_data_stream.NewModifyDataStreamFunc(tp),
			Open:                  indices_open.NewOpenFunc(tp),
			PromoteDataStream:     indices_promote_data_stream.NewPromoteDataStreamFunc(tp),
			PutAlias:              indices_put_alias.NewPutAliasFunc(tp),
			PutDataLifecycle:      indices_put_data_lifecycle.NewPutDataLifecycleFunc(tp),
			PutIndexTemplate:      indices_put_index_template.NewPutIndexTemplateFunc(tp),
			PutMapping:            indices_put_mapping.NewPutMappingFunc(tp),
			PutSettings:           indices_put_settings.NewPutSettingsFunc(tp),
			PutTemplate:           indices_put_template.NewPutTemplateFunc(tp),
			Recovery:              indices_recovery.NewRecoveryFunc(tp),
			Refresh:               indices_refresh.NewRefreshFunc(tp),
			ReloadSearchAnalyzers: indices_reload_search_analyzers.NewReloadSearchAnalyzersFunc(tp),
			ResolveCluster:        indices_resolve_cluster.NewResolveClusterFunc(tp),
			ResolveIndex:          indices_resolve_index.NewResolveIndexFunc(tp),
			Rollover:              indices_rollover.NewRolloverFunc(tp),
			Segments:              indices_segments.NewSegmentsFunc(tp),
			ShardStores:           indices_shard_stores.NewShardStoresFunc(tp),
			Shrink:                indices_shrink.NewShrinkFunc(tp),
			SimulateIndexTemplate: indices_simulate_index_template.NewSimulateIndexTemplateFunc(tp),
			SimulateTemplate:      indices_simulate_template.NewSimulateTemplateFunc(tp),
			Split:                 indices_split.NewSplitFunc(tp),
			Stats:                 indices_stats.NewStatsFunc(tp),
			Unfreeze:              indices_unfreeze.NewUnfreezeFunc(tp),
			UpdateAliases:         indices_update_aliases.NewUpdateAliasesFunc(tp),
			ValidateQuery:         indices_validate_query.NewValidateQueryFunc(tp),
		},

		// Inference
		Inference: Inference{
			Delete:    inference_delete.NewDeleteFunc(tp),
			Get:       inference_get.NewGetFunc(tp),
			Inference: inference_inference.NewInferenceFunc(tp),
			Put:       inference_put.NewPutFunc(tp),
		},

		// Ingest
		Ingest: Ingest{
			DeletePipeline: ingest_delete_pipeline.NewDeletePipelineFunc(tp),
			GeoIpStats:     ingest_geo_ip_stats.NewGeoIpStatsFunc(tp),
			GetPipeline:    ingest_get_pipeline.NewGetPipelineFunc(tp),
			ProcessorGrok:  ingest_processor_grok.NewProcessorGrokFunc(tp),
			PutPipeline:    ingest_put_pipeline.NewPutPipelineFunc(tp),
			Simulate:       ingest_simulate.NewSimulateFunc(tp),
		},

		// License
		License: License{
			Delete:         license_delete.NewDeleteFunc(tp),
			Get:            license_get.NewGetFunc(tp),
			GetBasicStatus: license_get_basic_status.NewGetBasicStatusFunc(tp),
			GetTrialStatus: license_get_trial_status.NewGetTrialStatusFunc(tp),
			Post:           license_post.NewPostFunc(tp),
			PostStartBasic: license_post_start_basic.NewPostStartBasicFunc(tp),
			PostStartTrial: license_post_start_trial.NewPostStartTrialFunc(tp),
		},

		// Logstash
		Logstash: Logstash{
			DeletePipeline: logstash_delete_pipeline.NewDeletePipelineFunc(tp),
			GetPipeline:    logstash_get_pipeline.NewGetPipelineFunc(tp),
			PutPipeline:    logstash_put_pipeline.NewPutPipelineFunc(tp),
		},

		// Migration
		Migration: Migration{
			Deprecations:            migration_deprecations.NewDeprecationsFunc(tp),
			GetFeatureUpgradeStatus: migration_get_feature_upgrade_status.NewGetFeatureUpgradeStatusFunc(tp),
			PostFeatureUpgrade:      migration_post_feature_upgrade.NewPostFeatureUpgradeFunc(tp),
		},

		// Ml
		Ml: Ml{
			ClearTrainedModelDeploymentCache: ml_clear_trained_model_deployment_cache.NewClearTrainedModelDeploymentCacheFunc(tp),
			CloseJob:                         ml_close_job.NewCloseJobFunc(tp),
			DeleteCalendar:                   ml_delete_calendar.NewDeleteCalendarFunc(tp),
			DeleteCalendarEvent:              ml_delete_calendar_event.NewDeleteCalendarEventFunc(tp),
			DeleteCalendarJob:                ml_delete_calendar_job.NewDeleteCalendarJobFunc(tp),
			DeleteDataFrameAnalytics:         ml_delete_data_frame_analytics.NewDeleteDataFrameAnalyticsFunc(tp),
			DeleteDatafeed:                   ml_delete_datafeed.NewDeleteDatafeedFunc(tp),
			DeleteExpiredData:                ml_delete_expired_data.NewDeleteExpiredDataFunc(tp),
			DeleteFilter:                     ml_delete_filter.NewDeleteFilterFunc(tp),
			DeleteForecast:                   ml_delete_forecast.NewDeleteForecastFunc(tp),
			DeleteJob:                        ml_delete_job.NewDeleteJobFunc(tp),
			DeleteModelSnapshot:              ml_delete_model_snapshot.NewDeleteModelSnapshotFunc(tp),
			DeleteTrainedModel:               ml_delete_trained_model.NewDeleteTrainedModelFunc(tp),
			DeleteTrainedModelAlias:          ml_delete_trained_model_alias.NewDeleteTrainedModelAliasFunc(tp),
			EstimateModelMemory:              ml_estimate_model_memory.NewEstimateModelMemoryFunc(tp),
			EvaluateDataFrame:                ml_evaluate_data_frame.NewEvaluateDataFrameFunc(tp),
			ExplainDataFrameAnalytics:        ml_explain_data_frame_analytics.NewExplainDataFrameAnalyticsFunc(tp),
			FlushJob:                         ml_flush_job.NewFlushJobFunc(tp),
			Forecast:                         ml_forecast.NewForecastFunc(tp),
			GetBuckets:                       ml_get_buckets.NewGetBucketsFunc(tp),
			GetCalendarEvents:                ml_get_calendar_events.NewGetCalendarEventsFunc(tp),
			GetCalendars:                     ml_get_calendars.NewGetCalendarsFunc(tp),
			GetCategories:                    ml_get_categories.NewGetCategoriesFunc(tp),
			GetDataFrameAnalytics:            ml_get_data_frame_analytics.NewGetDataFrameAnalyticsFunc(tp),
			GetDataFrameAnalyticsStats:       ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStatsFunc(tp),
			GetDatafeedStats:                 ml_get_datafeed_stats.NewGetDatafeedStatsFunc(tp),
			GetDatafeeds:                     ml_get_datafeeds.NewGetDatafeedsFunc(tp),
			GetFilters:                       ml_get_filters.NewGetFiltersFunc(tp),
			GetInfluencers:                   ml_get_influencers.NewGetInfluencersFunc(tp),
			GetJobStats:                      ml_get_job_stats.NewGetJobStatsFunc(tp),
			GetJobs:                          ml_get_jobs.NewGetJobsFunc(tp),
			GetMemoryStats:                   ml_get_memory_stats.NewGetMemoryStatsFunc(tp),
			GetModelSnapshotUpgradeStats:     ml_get_model_snapshot_upgrade_stats.NewGetModelSnapshotUpgradeStatsFunc(tp),
			GetModelSnapshots:                ml_get_model_snapshots.NewGetModelSnapshotsFunc(tp),
			GetOverallBuckets:                ml_get_overall_buckets.NewGetOverallBucketsFunc(tp),
			GetRecords:                       ml_get_records.NewGetRecordsFunc(tp),
			GetTrainedModels:                 ml_get_trained_models.NewGetTrainedModelsFunc(tp),
			GetTrainedModelsStats:            ml_get_trained_models_stats.NewGetTrainedModelsStatsFunc(tp),
			InferTrainedModel:                ml_infer_trained_model.NewInferTrainedModelFunc(tp),
			Info:                             ml_info.NewInfoFunc(tp),
			OpenJob:                          ml_open_job.NewOpenJobFunc(tp),
			PostCalendarEvents:               ml_post_calendar_events.NewPostCalendarEventsFunc(tp),
			PostData:                         ml_post_data.NewPostDataFunc(tp),
			PreviewDataFrameAnalytics:        ml_preview_data_frame_analytics.NewPreviewDataFrameAnalyticsFunc(tp),
			PreviewDatafeed:                  ml_preview_datafeed.NewPreviewDatafeedFunc(tp),
			PutCalendar:                      ml_put_calendar.NewPutCalendarFunc(tp),
			PutCalendarJob:                   ml_put_calendar_job.NewPutCalendarJobFunc(tp),
			PutDataFrameAnalytics:            ml_put_data_frame_analytics.NewPutDataFrameAnalyticsFunc(tp),
			PutDatafeed:                      ml_put_datafeed.NewPutDatafeedFunc(tp),
			PutFilter:                        ml_put_filter.NewPutFilterFunc(tp),
			PutJob:                           ml_put_job.NewPutJobFunc(tp),
			PutTrainedModel:                  ml_put_trained_model.NewPutTrainedModelFunc(tp),
			PutTrainedModelAlias:             ml_put_trained_model_alias.NewPutTrainedModelAliasFunc(tp),
			PutTrainedModelDefinitionPart:    ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPartFunc(tp),
			PutTrainedModelVocabulary:        ml_put_trained_model_vocabulary.NewPutTrainedModelVocabularyFunc(tp),
			ResetJob:                         ml_reset_job.NewResetJobFunc(tp),
			RevertModelSnapshot:              ml_revert_model_snapshot.NewRevertModelSnapshotFunc(tp),
			SetUpgradeMode:                   ml_set_upgrade_mode.NewSetUpgradeModeFunc(tp),
			StartDataFrameAnalytics:          ml_start_data_frame_analytics.NewStartDataFrameAnalyticsFunc(tp),
			StartDatafeed:                    ml_start_datafeed.NewStartDatafeedFunc(tp),
			StartTrainedModelDeployment:      ml_start_trained_model_deployment.NewStartTrainedModelDeploymentFunc(tp),
			StopDataFrameAnalytics:           ml_stop_data_frame_analytics.NewStopDataFrameAnalyticsFunc(tp),
			StopDatafeed:                     ml_stop_datafeed.NewStopDatafeedFunc(tp),
			StopTrainedModelDeployment:       ml_stop_trained_model_deployment.NewStopTrainedModelDeploymentFunc(tp),
			UpdateDataFrameAnalytics:         ml_update_data_frame_analytics.NewUpdateDataFrameAnalyticsFunc(tp),
			UpdateDatafeed:                   ml_update_datafeed.NewUpdateDatafeedFunc(tp),
			UpdateFilter:                     ml_update_filter.NewUpdateFilterFunc(tp),
			UpdateJob:                        ml_update_job.NewUpdateJobFunc(tp),
			UpdateModelSnapshot:              ml_update_model_snapshot.NewUpdateModelSnapshotFunc(tp),
			UpdateTrainedModelDeployment:     ml_update_trained_model_deployment.NewUpdateTrainedModelDeploymentFunc(tp),
			UpgradeJobSnapshot:               ml_upgrade_job_snapshot.NewUpgradeJobSnapshotFunc(tp),
			Validate:                         ml_validate.NewValidateFunc(tp),
			ValidateDetector:                 ml_validate_detector.NewValidateDetectorFunc(tp),
		},

		// Monitoring
		Monitoring: Monitoring{
			Bulk: monitoring_bulk.NewBulkFunc(tp),
		},

		// Nodes
		Nodes: Nodes{
			ClearRepositoriesMeteringArchive: nodes_clear_repositories_metering_archive.NewClearRepositoriesMeteringArchiveFunc(tp),
			GetRepositoriesMeteringInfo:      nodes_get_repositories_metering_info.NewGetRepositoriesMeteringInfoFunc(tp),
			HotThreads:                       nodes_hot_threads.NewHotThreadsFunc(tp),
			Info:                             nodes_info.NewInfoFunc(tp),
			ReloadSecureSettings:             nodes_reload_secure_settings.NewReloadSecureSettingsFunc(tp),
			Stats:                            nodes_stats.NewStatsFunc(tp),
			Usage:                            nodes_usage.NewUsageFunc(tp),
		},

		// Profiling
		Profiling: Profiling{
			Flamegraph:    profiling_flamegraph.NewFlamegraphFunc(tp),
			Stacktraces:   profiling_stacktraces.NewStacktracesFunc(tp),
			Status:        profiling_status.NewStatusFunc(tp),
			TopnFunctions: profiling_topn_functions.NewTopnFunctionsFunc(tp),
		},

		// QueryRules
		QueryRules: QueryRules{
			DeleteRule:    query_rules_delete_rule.NewDeleteRuleFunc(tp),
			DeleteRuleset: query_rules_delete_ruleset.NewDeleteRulesetFunc(tp),
			GetRule:       query_rules_get_rule.NewGetRuleFunc(tp),
			GetRuleset:    query_rules_get_ruleset.NewGetRulesetFunc(tp),
			ListRulesets:  query_rules_list_rulesets.NewListRulesetsFunc(tp),
			PutRule:       query_rules_put_rule.NewPutRuleFunc(tp),
			PutRuleset:    query_rules_put_ruleset.NewPutRulesetFunc(tp),
		},

		// Rollup
		Rollup: Rollup{
			DeleteJob:          rollup_delete_job.NewDeleteJobFunc(tp),
			GetJobs:            rollup_get_jobs.NewGetJobsFunc(tp),
			GetRollupCaps:      rollup_get_rollup_caps.NewGetRollupCapsFunc(tp),
			GetRollupIndexCaps: rollup_get_rollup_index_caps.NewGetRollupIndexCapsFunc(tp),
			PutJob:             rollup_put_job.NewPutJobFunc(tp),
			RollupSearch:       rollup_rollup_search.NewRollupSearchFunc(tp),
			StartJob:           rollup_start_job.NewStartJobFunc(tp),
			StopJob:            rollup_stop_job.NewStopJobFunc(tp),
		},

		// SearchApplication
		SearchApplication: SearchApplication{
			Delete:                    search_application_delete.NewDeleteFunc(tp),
			DeleteBehavioralAnalytics: search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalyticsFunc(tp),
			Get:                       search_application_get.NewGetFunc(tp),
			GetBehavioralAnalytics:    search_application_get_behavioral_analytics.NewGetBehavioralAnalyticsFunc(tp),
			List:                      search_application_list.NewListFunc(tp),
			Put:                       search_application_put.NewPutFunc(tp),
			PutBehavioralAnalytics:    search_application_put_behavioral_analytics.NewPutBehavioralAnalyticsFunc(tp),
			Search:                    search_application_search.NewSearchFunc(tp),
		},

		// SearchableSnapshots
		SearchableSnapshots: SearchableSnapshots{
			CacheStats: searchable_snapshots_cache_stats.NewCacheStatsFunc(tp),
			ClearCache: searchable_snapshots_clear_cache.NewClearCacheFunc(tp),
			Mount:      searchable_snapshots_mount.NewMountFunc(tp),
			Stats:      searchable_snapshots_stats.NewStatsFunc(tp),
		},

		// Security
		Security: Security{
			ActivateUserProfile:         security_activate_user_profile.NewActivateUserProfileFunc(tp),
			Authenticate:                security_authenticate.NewAuthenticateFunc(tp),
			BulkDeleteRole:              security_bulk_delete_role.NewBulkDeleteRoleFunc(tp),
			BulkPutRole:                 security_bulk_put_role.NewBulkPutRoleFunc(tp),
			BulkUpdateApiKeys:           security_bulk_update_api_keys.NewBulkUpdateApiKeysFunc(tp),
			ChangePassword:              security_change_password.NewChangePasswordFunc(tp),
			ClearApiKeyCache:            security_clear_api_key_cache.NewClearApiKeyCacheFunc(tp),
			ClearCachedPrivileges:       security_clear_cached_privileges.NewClearCachedPrivilegesFunc(tp),
			ClearCachedRealms:           security_clear_cached_realms.NewClearCachedRealmsFunc(tp),
			ClearCachedRoles:            security_clear_cached_roles.NewClearCachedRolesFunc(tp),
			ClearCachedServiceTokens:    security_clear_cached_service_tokens.NewClearCachedServiceTokensFunc(tp),
			CreateApiKey:                security_create_api_key.NewCreateApiKeyFunc(tp),
			CreateCrossClusterApiKey:    security_create_cross_cluster_api_key.NewCreateCrossClusterApiKeyFunc(tp),
			CreateServiceToken:          security_create_service_token.NewCreateServiceTokenFunc(tp),
			DeletePrivileges:            security_delete_privileges.NewDeletePrivilegesFunc(tp),
			DeleteRole:                  security_delete_role.NewDeleteRoleFunc(tp),
			DeleteRoleMapping:           security_delete_role_mapping.NewDeleteRoleMappingFunc(tp),
			DeleteServiceToken:          security_delete_service_token.NewDeleteServiceTokenFunc(tp),
			DeleteUser:                  security_delete_user.NewDeleteUserFunc(tp),
			DisableUser:                 security_disable_user.NewDisableUserFunc(tp),
			DisableUserProfile:          security_disable_user_profile.NewDisableUserProfileFunc(tp),
			EnableUser:                  security_enable_user.NewEnableUserFunc(tp),
			EnableUserProfile:           security_enable_user_profile.NewEnableUserProfileFunc(tp),
			EnrollKibana:                security_enroll_kibana.NewEnrollKibanaFunc(tp),
			EnrollNode:                  security_enroll_node.NewEnrollNodeFunc(tp),
			GetApiKey:                   security_get_api_key.NewGetApiKeyFunc(tp),
			GetBuiltinPrivileges:        security_get_builtin_privileges.NewGetBuiltinPrivilegesFunc(tp),
			GetPrivileges:               security_get_privileges.NewGetPrivilegesFunc(tp),
			GetRole:                     security_get_role.NewGetRoleFunc(tp),
			GetRoleMapping:              security_get_role_mapping.NewGetRoleMappingFunc(tp),
			GetServiceAccounts:          security_get_service_accounts.NewGetServiceAccountsFunc(tp),
			GetServiceCredentials:       security_get_service_credentials.NewGetServiceCredentialsFunc(tp),
			GetSettings:                 security_get_settings.NewGetSettingsFunc(tp),
			GetToken:                    security_get_token.NewGetTokenFunc(tp),
			GetUser:                     security_get_user.NewGetUserFunc(tp),
			GetUserPrivileges:           security_get_user_privileges.NewGetUserPrivilegesFunc(tp),
			GetUserProfile:              security_get_user_profile.NewGetUserProfileFunc(tp),
			GrantApiKey:                 security_grant_api_key.NewGrantApiKeyFunc(tp),
			HasPrivileges:               security_has_privileges.NewHasPrivilegesFunc(tp),
			HasPrivilegesUserProfile:    security_has_privileges_user_profile.NewHasPrivilegesUserProfileFunc(tp),
			InvalidateApiKey:            security_invalidate_api_key.NewInvalidateApiKeyFunc(tp),
			InvalidateToken:             security_invalidate_token.NewInvalidateTokenFunc(tp),
			OidcAuthenticate:            security_oidc_authenticate.NewOidcAuthenticateFunc(tp),
			OidcLogout:                  security_oidc_logout.NewOidcLogoutFunc(tp),
			OidcPrepareAuthentication:   security_oidc_prepare_authentication.NewOidcPrepareAuthenticationFunc(tp),
			PutPrivileges:               security_put_privileges.NewPutPrivilegesFunc(tp),
			PutRole:                     security_put_role.NewPutRoleFunc(tp),
			PutRoleMapping:              security_put_role_mapping.NewPutRoleMappingFunc(tp),
			PutUser:                     security_put_user.NewPutUserFunc(tp),
			QueryApiKeys:                security_query_api_keys.NewQueryApiKeysFunc(tp),
			QueryRole:                   security_query_role.NewQueryRoleFunc(tp),
			QueryUser:                   security_query_user.NewQueryUserFunc(tp),
			SamlAuthenticate:            security_saml_authenticate.NewSamlAuthenticateFunc(tp),
			SamlCompleteLogout:          security_saml_complete_logout.NewSamlCompleteLogoutFunc(tp),
			SamlInvalidate:              security_saml_invalidate.NewSamlInvalidateFunc(tp),
			SamlLogout:                  security_saml_logout.NewSamlLogoutFunc(tp),
			SamlPrepareAuthentication:   security_saml_prepare_authentication.NewSamlPrepareAuthenticationFunc(tp),
			SamlServiceProviderMetadata: security_saml_service_provider_metadata.NewSamlServiceProviderMetadataFunc(tp),
			SuggestUserProfiles:         security_suggest_user_profiles.NewSuggestUserProfilesFunc(tp),
			UpdateApiKey:                security_update_api_key.NewUpdateApiKeyFunc(tp),
			UpdateSettings:              security_update_settings.NewUpdateSettingsFunc(tp),
			UpdateUserProfileData:       security_update_user_profile_data.NewUpdateUserProfileDataFunc(tp),
		},

		// Shutdown
		Shutdown: Shutdown{
			DeleteNode: shutdown_delete_node.NewDeleteNodeFunc(tp),
			GetNode:    shutdown_get_node.NewGetNodeFunc(tp),
			PutNode:    shutdown_put_node.NewPutNodeFunc(tp),
		},

		// Slm
		Slm: Slm{
			DeleteLifecycle:  slm_delete_lifecycle.NewDeleteLifecycleFunc(tp),
			ExecuteLifecycle: slm_execute_lifecycle.NewExecuteLifecycleFunc(tp),
			ExecuteRetention: slm_execute_retention.NewExecuteRetentionFunc(tp),
			GetLifecycle:     slm_get_lifecycle.NewGetLifecycleFunc(tp),
			GetStats:         slm_get_stats.NewGetStatsFunc(tp),
			GetStatus:        slm_get_status.NewGetStatusFunc(tp),
			PutLifecycle:     slm_put_lifecycle.NewPutLifecycleFunc(tp),
			Start:            slm_start.NewStartFunc(tp),
			Stop:             slm_stop.NewStopFunc(tp),
		},

		// Snapshot
		Snapshot: Snapshot{
			CleanupRepository: snapshot_cleanup_repository.NewCleanupRepositoryFunc(tp),
			Clone:             snapshot_clone.NewCloneFunc(tp),
			Create:            snapshot_create.NewCreateFunc(tp),
			CreateRepository:  snapshot_create_repository.NewCreateRepositoryFunc(tp),
			Delete:            snapshot_delete.NewDeleteFunc(tp),
			DeleteRepository:  snapshot_delete_repository.NewDeleteRepositoryFunc(tp),
			Get:               snapshot_get.NewGetFunc(tp),
			GetRepository:     snapshot_get_repository.NewGetRepositoryFunc(tp),
			Restore:           snapshot_restore.NewRestoreFunc(tp),
			Status:            snapshot_status.NewStatusFunc(tp),
			VerifyRepository:  snapshot_verify_repository.NewVerifyRepositoryFunc(tp),
		},

		// Sql
		Sql: Sql{
			ClearCursor:    sql_clear_cursor.NewClearCursorFunc(tp),
			DeleteAsync:    sql_delete_async.NewDeleteAsyncFunc(tp),
			GetAsync:       sql_get_async.NewGetAsyncFunc(tp),
			GetAsyncStatus: sql_get_async_status.NewGetAsyncStatusFunc(tp),
			Query:          sql_query.NewQueryFunc(tp),
			Translate:      sql_translate.NewTranslateFunc(tp),
		},

		// Ssl
		Ssl: Ssl{
			Certificates: ssl_certificates.NewCertificatesFunc(tp),
		},

		// Synonyms
		Synonyms: Synonyms{
			DeleteSynonym:     synonyms_delete_synonym.NewDeleteSynonymFunc(tp),
			DeleteSynonymRule: synonyms_delete_synonym_rule.NewDeleteSynonymRuleFunc(tp),
			GetSynonym:        synonyms_get_synonym.NewGetSynonymFunc(tp),
			GetSynonymRule:    synonyms_get_synonym_rule.NewGetSynonymRuleFunc(tp),
			GetSynonymsSets:   synonyms_get_synonyms_sets.NewGetSynonymsSetsFunc(tp),
			PutSynonym:        synonyms_put_synonym.NewPutSynonymFunc(tp),
			PutSynonymRule:    synonyms_put_synonym_rule.NewPutSynonymRuleFunc(tp),
		},

		// Tasks
		Tasks: Tasks{
			Cancel: tasks_cancel.NewCancelFunc(tp),
			Get:    tasks_get.NewGetFunc(tp),
			List:   tasks_list.NewListFunc(tp),
		},

		// TextStructure
		TextStructure: TextStructure{
			FindFieldStructure:   text_structure_find_field_structure.NewFindFieldStructureFunc(tp),
			FindMessageStructure: text_structure_find_message_structure.NewFindMessageStructureFunc(tp),
			FindStructure:        text_structure_find_structure.NewFindStructureFunc(tp),
			TestGrokPattern:      text_structure_test_grok_pattern.NewTestGrokPatternFunc(tp),
		},

		// Transform
		Transform: Transform{
			DeleteTransform:      transform_delete_transform.NewDeleteTransformFunc(tp),
			GetNodeStats:         transform_get_node_stats.NewGetNodeStatsFunc(tp),
			GetTransform:         transform_get_transform.NewGetTransformFunc(tp),
			GetTransformStats:    transform_get_transform_stats.NewGetTransformStatsFunc(tp),
			PreviewTransform:     transform_preview_transform.NewPreviewTransformFunc(tp),
			PutTransform:         transform_put_transform.NewPutTransformFunc(tp),
			ResetTransform:       transform_reset_transform.NewResetTransformFunc(tp),
			ScheduleNowTransform: transform_schedule_now_transform.NewScheduleNowTransformFunc(tp),
			StartTransform:       transform_start_transform.NewStartTransformFunc(tp),
			StopTransform:        transform_stop_transform.NewStopTransformFunc(tp),
			UpdateTransform:      transform_update_transform.NewUpdateTransformFunc(tp),
			UpgradeTransforms:    transform_upgrade_transforms.NewUpgradeTransformsFunc(tp),
		},

		// Watcher
		Watcher: Watcher{
			AckWatch:        watcher_ack_watch.NewAckWatchFunc(tp),
			ActivateWatch:   watcher_activate_watch.NewActivateWatchFunc(tp),
			DeactivateWatch: watcher_deactivate_watch.NewDeactivateWatchFunc(tp),
			DeleteWatch:     watcher_delete_watch.NewDeleteWatchFunc(tp),
			ExecuteWatch:    watcher_execute_watch.NewExecuteWatchFunc(tp),
			GetSettings:     watcher_get_settings.NewGetSettingsFunc(tp),
			GetWatch:        watcher_get_watch.NewGetWatchFunc(tp),
			PutWatch:        watcher_put_watch.NewPutWatchFunc(tp),
			QueryWatches:    watcher_query_watches.NewQueryWatchesFunc(tp),
			Start:           watcher_start.NewStartFunc(tp),
			Stats:           watcher_stats.NewStatsFunc(tp),
			Stop:            watcher_stop.NewStopFunc(tp),
			UpdateSettings:  watcher_update_settings.NewUpdateSettingsFunc(tp),
		},

		// Xpack
		Xpack: Xpack{
			Info:  xpack_info.NewInfoFunc(tp),
			Usage: xpack_usage.NewUsageFunc(tp),
		},

		Bulk:                    core_bulk.NewBulkFunc(tp),
		ClearScroll:             core_clear_scroll.NewClearScrollFunc(tp),
		ClosePointInTime:        core_close_point_in_time.NewClosePointInTimeFunc(tp),
		Count:                   core_count.NewCountFunc(tp),
		Create:                  core_create.NewCreateFunc(tp),
		Delete:                  core_delete.NewDeleteFunc(tp),
		DeleteByQuery:           core_delete_by_query.NewDeleteByQueryFunc(tp),
		DeleteByQueryRethrottle: core_delete_by_query_rethrottle.NewDeleteByQueryRethrottleFunc(tp),
		DeleteScript:            core_delete_script.NewDeleteScriptFunc(tp),
		Exists:                  core_exists.NewExistsFunc(tp),
		ExistsSource:            core_exists_source.NewExistsSourceFunc(tp),
		Explain:                 core_explain.NewExplainFunc(tp),
		FieldCaps:               core_field_caps.NewFieldCapsFunc(tp),
		Get:                     core_get.NewGetFunc(tp),
		GetScript:               core_get_script.NewGetScriptFunc(tp),
		GetScriptContext:        core_get_script_context.NewGetScriptContextFunc(tp),
		GetScriptLanguages:      core_get_script_languages.NewGetScriptLanguagesFunc(tp),
		GetSource:               core_get_source.NewGetSourceFunc(tp),
		HealthReport:            core_health_report.NewHealthReportFunc(tp),
		Index:                   core_index.NewIndexFunc(tp),
		Info:                    core_info.NewInfoFunc(tp),
		KnnSearch:               core_knn_search.NewKnnSearchFunc(tp),
		Mget:                    core_mget.NewMgetFunc(tp),
		Msearch:                 core_msearch.NewMsearchFunc(tp),
		MsearchTemplate:         core_msearch_template.NewMsearchTemplateFunc(tp),
		Mtermvectors:            core_mtermvectors.NewMtermvectorsFunc(tp),
		OpenPointInTime:         core_open_point_in_time.NewOpenPointInTimeFunc(tp),
		Ping:                    core_ping.NewPingFunc(tp),
		PutScript:               core_put_script.NewPutScriptFunc(tp),
		RankEval:                core_rank_eval.NewRankEvalFunc(tp),
		Reindex:                 core_reindex.NewReindexFunc(tp),
		ReindexRethrottle:       core_reindex_rethrottle.NewReindexRethrottleFunc(tp),
		RenderSearchTemplate:    core_render_search_template.NewRenderSearchTemplateFunc(tp),
		ScriptsPainlessExecute:  core_scripts_painless_execute.NewScriptsPainlessExecuteFunc(tp),
		Scroll:                  core_scroll.NewScrollFunc(tp),
		Search:                  core_search.NewSearchFunc(tp),
		SearchMvt:               core_search_mvt.NewSearchMvtFunc(tp),
		SearchShards:            core_search_shards.NewSearchShardsFunc(tp),
		SearchTemplate:          core_search_template.NewSearchTemplateFunc(tp),
		TermsEnum:               core_terms_enum.NewTermsEnumFunc(tp),
		Termvectors:             core_termvectors.NewTermvectorsFunc(tp),
		Update:                  core_update.NewUpdateFunc(tp),
		UpdateByQuery:           core_update_by_query.NewUpdateByQueryFunc(tp),
		UpdateByQueryRethrottle: core_update_by_query_rethrottle.NewUpdateByQueryRethrottleFunc(tp),
	}
}

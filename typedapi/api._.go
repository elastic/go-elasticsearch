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
// https://github.com/elastic/elasticsearch-specification/tree/3a94b6715915b1e9311724a2614c643368eece90

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
	connector_sync_job_check_in "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobcheckin"
	connector_sync_job_claim "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobclaim"
	connector_sync_job_delete "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobdelete"
	connector_sync_job_error "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjoberror"
	connector_sync_job_get "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobget"
	connector_sync_job_list "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjoblist"
	connector_sync_job_post "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobpost"
	connector_sync_job_update_stats "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobupdatestats"
	connector_update_active_filtering "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateactivefiltering"
	connector_update_api_key_id "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateapikeyid"
	connector_update_configuration "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateconfiguration"
	connector_update_error "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateerror"
	connector_update_features "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatefeatures"
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
	esql_async_query_delete "github.com/elastic/go-elasticsearch/v8/typedapi/esql/asyncquerydelete"
	esql_async_query_get "github.com/elastic/go-elasticsearch/v8/typedapi/esql/asyncqueryget"
	esql_async_query_stop "github.com/elastic/go-elasticsearch/v8/typedapi/esql/asyncquerystop"
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
	indices_cancel_migrate_reindex "github.com/elastic/go-elasticsearch/v8/typedapi/indices/cancelmigratereindex"
	indices_clear_cache "github.com/elastic/go-elasticsearch/v8/typedapi/indices/clearcache"
	indices_clone "github.com/elastic/go-elasticsearch/v8/typedapi/indices/clone"
	indices_close "github.com/elastic/go-elasticsearch/v8/typedapi/indices/close"
	indices_create "github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	indices_create_data_stream "github.com/elastic/go-elasticsearch/v8/typedapi/indices/createdatastream"
	indices_create_from "github.com/elastic/go-elasticsearch/v8/typedapi/indices/createfrom"
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
	indices_get_data_lifecycle_stats "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getdatalifecyclestats"
	indices_get_data_stream "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getdatastream"
	indices_get_field_mapping "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getfieldmapping"
	indices_get_index_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getindextemplate"
	indices_get_mapping "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getmapping"
	indices_get_migrate_reindex_status "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getmigratereindexstatus"
	indices_get_settings "github.com/elastic/go-elasticsearch/v8/typedapi/indices/getsettings"
	indices_get_template "github.com/elastic/go-elasticsearch/v8/typedapi/indices/gettemplate"
	indices_migrate_reindex "github.com/elastic/go-elasticsearch/v8/typedapi/indices/migratereindex"
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
	inference_chat_completion_unified "github.com/elastic/go-elasticsearch/v8/typedapi/inference/chatcompletionunified"
	inference_completion "github.com/elastic/go-elasticsearch/v8/typedapi/inference/completion"
	inference_delete "github.com/elastic/go-elasticsearch/v8/typedapi/inference/delete"
	inference_get "github.com/elastic/go-elasticsearch/v8/typedapi/inference/get"
	inference_inference "github.com/elastic/go-elasticsearch/v8/typedapi/inference/inference"
	inference_put "github.com/elastic/go-elasticsearch/v8/typedapi/inference/put"
	inference_put_alibabacloud "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putalibabacloud"
	inference_put_amazonbedrock "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putamazonbedrock"
	inference_put_anthropic "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putanthropic"
	inference_put_azureaistudio "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putazureaistudio"
	inference_put_azureopenai "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putazureopenai"
	inference_put_cohere "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putcohere"
	inference_put_elasticsearch "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putelasticsearch"
	inference_put_elser "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putelser"
	inference_put_googleaistudio "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putgoogleaistudio"
	inference_put_googlevertexai "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putgooglevertexai"
	inference_put_hugging_face "github.com/elastic/go-elasticsearch/v8/typedapi/inference/puthuggingface"
	inference_put_jinaai "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putjinaai"
	inference_put_mistral "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putmistral"
	inference_put_openai "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putopenai"
	inference_put_voyageai "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putvoyageai"
	inference_put_watsonx "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putwatsonx"
	inference_rerank "github.com/elastic/go-elasticsearch/v8/typedapi/inference/rerank"
	inference_sparse_embedding "github.com/elastic/go-elasticsearch/v8/typedapi/inference/sparseembedding"
	inference_stream_completion "github.com/elastic/go-elasticsearch/v8/typedapi/inference/streamcompletion"
	inference_text_embedding "github.com/elastic/go-elasticsearch/v8/typedapi/inference/textembedding"
	inference_update "github.com/elastic/go-elasticsearch/v8/typedapi/inference/update"
	ingest_delete_geoip_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/deletegeoipdatabase"
	ingest_delete_ip_location_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/deleteiplocationdatabase"
	ingest_delete_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/deletepipeline"
	ingest_geo_ip_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/geoipstats"
	ingest_get_geoip_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/getgeoipdatabase"
	ingest_get_ip_location_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/getiplocationdatabase"
	ingest_get_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/getpipeline"
	ingest_processor_grok "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/processorgrok"
	ingest_put_geoip_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/putgeoipdatabase"
	ingest_put_ip_location_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/putiplocationdatabase"
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
	query_rules_test "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/test"
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
	search_application_post_behavioral_analytics_event "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/postbehavioralanalyticsevent"
	search_application_put "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/put"
	search_application_put_behavioral_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/putbehavioralanalytics"
	search_application_render_query "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/renderquery"
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
	security_delegate_pki "github.com/elastic/go-elasticsearch/v8/typedapi/security/delegatepki"
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
	security_update_cross_cluster_api_key "github.com/elastic/go-elasticsearch/v8/typedapi/security/updatecrossclusterapikey"
	security_update_settings "github.com/elastic/go-elasticsearch/v8/typedapi/security/updatesettings"
	security_update_user_profile_data "github.com/elastic/go-elasticsearch/v8/typedapi/security/updateuserprofiledata"
	shutdown_delete_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/deletenode"
	shutdown_get_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/getnode"
	shutdown_put_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/putnode"
	simulate_ingest "github.com/elastic/go-elasticsearch/v8/typedapi/simulate/ingest"
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
	snapshot_repository_analyze "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/repositoryanalyze"
	snapshot_repository_verify_integrity "github.com/elastic/go-elasticsearch/v8/typedapi/snapshot/repositoryverifyintegrity"
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
	// Delete an async search.
	//
	// If the asynchronous search is still running, it is cancelled.
	// Otherwise, the saved search results are deleted.
	// If the Elasticsearch security features are enabled, the deletion of a
	// specific async search is restricted to: the authenticated user that submitted
	// the original search request; users that have the `cancel_task` cluster
	// privilege.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Delete async_search_delete.NewDelete
	// Get async search results.
	//
	// Retrieve the results of a previously submitted asynchronous search request.
	// If the Elasticsearch security features are enabled, access to the results of
	// a specific async search is restricted to the user or API key that submitted
	// it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Get async_search_get.NewGet
	// Get the async search status.
	//
	// Get the status of a previously submitted async search request given its
	// identifier, without retrieving search results.
	// If the Elasticsearch security features are enabled, the access to the status
	// of a specific async search is restricted to:
	//
	// * The user or API key that submitted the original async search request.
	// * Users that have the `monitor` cluster privilege or greater privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Status async_search_status.NewStatus
	// Run an async search.
	//
	// When the primary sort of the results is an indexed field, shards get sorted
	// based on minimum and maximum value that they hold for that field. Partial
	// results become available following the sort criteria that was requested.
	//
	// Warning: Asynchronous search does not support scroll or search requests that
	// include only the suggest section.
	//
	// By default, Elasticsearch does not allow you to store an async search
	// response larger than 10Mb and an attempt to do this results in an error.
	// The maximum allowed size for a stored async search response can be set by
	// changing the `search.max_async_search_response_size` cluster level setting.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Submit async_search_submit.NewSubmit
}

type Autoscaling struct {
	// Delete an autoscaling policy.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-delete-autoscaling-policy.html
	DeleteAutoscalingPolicy autoscaling_delete_autoscaling_policy.NewDeleteAutoscalingPolicy
	// Get the autoscaling capacity.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	//
	// This API gets the current autoscaling capacity based on the configured
	// autoscaling policy.
	// It will return information to size the cluster appropriately to the current
	// workload.
	//
	// The `required_capacity` is calculated as the maximum of the
	// `required_capacity` result of all individual deciders that are enabled for
	// the policy.
	//
	// The operator should verify that the `current_nodes` match the operator’s
	// knowledge of the cluster to avoid making autoscaling decisions based on stale
	// or incomplete information.
	//
	// The response contains decider-specific information you can use to diagnose
	// how and why autoscaling determined a certain capacity was required.
	// This information is provided for diagnosis only.
	// Do not use this information to make autoscaling decisions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-get-autoscaling-capacity.html
	GetAutoscalingCapacity autoscaling_get_autoscaling_capacity.NewGetAutoscalingCapacity
	// Get an autoscaling policy.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-get-autoscaling-capacity.html
	GetAutoscalingPolicy autoscaling_get_autoscaling_policy.NewGetAutoscalingPolicy
	// Create or update an autoscaling policy.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-put-autoscaling-policy.html
	PutAutoscalingPolicy autoscaling_put_autoscaling_policy.NewPutAutoscalingPolicy
}

type Capabilities struct {
	// Checks if the specified combination of method, API, parameters, and arbitrary
	// capabilities are supported
	// https://github.com/elastic/elasticsearch/blob/main/rest-api-spec/src/yamlRestTest/resources/rest-api-spec/test/README.asciidoc#require-or-skip-api-capabilities
	Capabilities capabilities.NewCapabilities
}

type Cat struct {
	// Get aliases.
	//
	// Get the cluster's index aliases, including filter and routing information.
	// This API does not return data stream aliases.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the command
	// line or the Kibana console. They are not intended for use by applications.
	// For application consumption, use the aliases API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-alias.html
	Aliases cat_aliases.NewAliases
	// Get shard allocation information.
	//
	// Get a snapshot of the number of shards allocated to each data node and their
	// disk space.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-allocation.html
	Allocation cat_allocation.NewAllocation
	// Get component templates.
	//
	// Get information about component templates in a cluster.
	// Component templates are building blocks for constructing index templates that
	// specify index mappings, settings, and aliases.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the get component template API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-component-templates.html
	ComponentTemplates cat_component_templates.NewComponentTemplates
	// Get a document count.
	//
	// Get quick access to a document count for a data stream, an index, or an
	// entire cluster.
	// The document count only includes live documents, not deleted documents which
	// have not yet been removed by the merge process.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the count API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-count.html
	Count cat_count.NewCount
	// Get field data cache information.
	//
	// Get the amount of heap memory currently used by the field data cache on every
	// data node in the cluster.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the nodes stats API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-fielddata.html
	Fielddata cat_fielddata.NewFielddata
	// Get the cluster health status.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the command
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
	// Get CAT help.
	//
	// Get help for the CAT APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat.html
	Help cat_help.NewHelp
	// Get index information.
	//
	// Get high-level information about indices in a cluster, including backing
	// indices for data streams.
	//
	// Use this request to get the following information for each index in a
	// cluster:
	// - shard count
	// - document count
	// - deleted document count
	// - primary store size
	// - total store size of all shards, including shard replicas
	//
	// These metrics are retrieved directly from Lucene, which Elasticsearch uses
	// internally to power indexing and search. As a result, all document counts
	// include hidden nested documents.
	// To get an accurate count of Elasticsearch documents, use the cat count or
	// count APIs.
	//
	// CAT APIs are only intended for human consumption using the command line or
	// Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use an index endpoint.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html
	Indices cat_indices.NewIndices
	// Get master node information.
	//
	// Get information about the master node, including the ID, bound IP address,
	// and name.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-master.html
	Master cat_master.NewMaster
	// Get data frame analytics jobs.
	//
	// Get configuration and usage information about data frame analytics jobs.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get data frame analytics jobs statistics
	// API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-dfanalytics.html
	MlDataFrameAnalytics cat_ml_data_frame_analytics.NewMlDataFrameAnalytics
	// Get datafeeds.
	//
	// Get configuration and usage information about datafeeds.
	// This API returns a maximum of 10,000 datafeeds.
	// If the Elasticsearch security features are enabled, you must have
	// `monitor_ml`, `monitor`, `manage_ml`, or `manage`
	// cluster privileges to use this API.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get datafeed statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-datafeeds.html
	MlDatafeeds cat_ml_datafeeds.NewMlDatafeeds
	// Get anomaly detection jobs.
	//
	// Get configuration and usage information for anomaly detection jobs.
	// This API returns a maximum of 10,000 jobs.
	// If the Elasticsearch security features are enabled, you must have
	// `monitor_ml`,
	// `monitor`, `manage_ml`, or `manage` cluster privileges to use this API.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get anomaly detection job statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-anomaly-detectors.html
	MlJobs cat_ml_jobs.NewMlJobs
	// Get trained models.
	//
	// Get configuration and usage information about inference trained models.
	//
	// IMPORTANT: CAT APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get trained models statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-trained-model.html
	MlTrainedModels cat_ml_trained_models.NewMlTrainedModels
	// Get node attribute information.
	//
	// Get information about custom node attributes.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodeattrs.html
	Nodeattrs cat_nodeattrs.NewNodeattrs
	// Get node information.
	//
	// Get information about the nodes in a cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodes.html
	Nodes cat_nodes.NewNodes
	// Get pending task information.
	//
	// Get information about cluster-level changes that have not yet taken effect.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the pending cluster tasks API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-pending-tasks.html
	PendingTasks cat_pending_tasks.NewPendingTasks
	// Get plugin information.
	//
	// Get a list of plugins running on each node of a cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-plugins.html
	Plugins cat_plugins.NewPlugins
	// Get shard recovery information.
	//
	// Get information about ongoing and completed shard recoveries.
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
	// Get snapshot repository information.
	//
	// Get a list of snapshot repositories for a cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the get snapshot repository API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-repositories.html
	Repositories cat_repositories.NewRepositories
	// Get segment information.
	//
	// Get low-level information about the Lucene segments in index shards.
	// For data streams, the API returns information about the backing indices.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the index segments API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-segments.html
	Segments cat_segments.NewSegments
	// Get shard information.
	//
	// Get information about the shards in a cluster.
	// For data streams, the API returns information about the backing indices.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-shards.html
	Shards cat_shards.NewShards
	// Get snapshot information.
	//
	// Get information about the snapshots stored in one or more repositories.
	// A snapshot is a backup of an index or running Elasticsearch cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the get snapshot API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-snapshots.html
	Snapshots cat_snapshots.NewSnapshots
	// Get task information.
	//
	// Get information about tasks currently running in the cluster.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the task management API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-tasks.html
	Tasks cat_tasks.NewTasks
	// Get index template information.
	//
	// Get information about the index templates in a cluster.
	// You can use index templates to apply index settings and field mappings to new
	// indices at creation.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the get index template API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-templates.html
	Templates cat_templates.NewTemplates
	// Get thread pool statistics.
	//
	// Get thread pool statistics for each node in a cluster.
	// Returned information includes all built-in thread pools and custom thread
	// pools.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console. They are not intended for use by applications. For
	// application consumption, use the nodes info API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-thread-pool.html
	ThreadPool cat_thread_pool.NewThreadPool
	// Get transform information.
	//
	// Get configuration and usage information about transforms.
	//
	// CAT APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get transform statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-transforms.html
	Transforms cat_transforms.NewTransforms
}

type Ccr struct {
	// Delete auto-follow patterns.
	//
	// Delete a collection of cross-cluster replication auto-follow patterns.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-delete-auto-follow-pattern.html
	DeleteAutoFollowPattern ccr_delete_auto_follow_pattern.NewDeleteAutoFollowPattern
	// Create a follower.
	// Create a cross-cluster replication follower index that follows a specific
	// leader index.
	// When the API returns, the follower index exists and cross-cluster replication
	// starts replicating operations from the leader index to the follower index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-follow.html
	Follow ccr_follow.NewFollow
	// Get follower information.
	//
	// Get information about all cross-cluster replication follower indices.
	// For example, the results include follower index names, leader index names,
	// replication options, and whether the follower indices are active or paused.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-follow-info.html
	FollowInfo ccr_follow_info.NewFollowInfo
	// Get follower stats.
	//
	// Get cross-cluster replication follower stats.
	// The API returns shard-level stats about the "following tasks" associated with
	// each shard for the specified indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-follow-stats.html
	FollowStats ccr_follow_stats.NewFollowStats
	// Forget a follower.
	// Remove the cross-cluster replication follower retention leases from the
	// leader.
	//
	// A following index takes out retention leases on its leader index.
	// These leases are used to increase the likelihood that the shards of the
	// leader index retain the history of operations that the shards of the
	// following index need to run replication.
	// When a follower index is converted to a regular index by the unfollow API
	// (either by directly calling the API or by index lifecycle management tasks),
	// these leases are removed.
	// However, removal of the leases can fail, for example when the remote cluster
	// containing the leader index is unavailable.
	// While the leases will eventually expire on their own, their extended
	// existence can cause the leader index to hold more history than necessary and
	// prevent index lifecycle management from performing some operations on the
	// leader index.
	// This API exists to enable manually removing the leases when the unfollow API
	// is unable to do so.
	//
	// NOTE: This API does not stop replication by a following index. If you use
	// this API with a follower index that is still actively following, the
	// following index will add back retention leases on the leader.
	// The only purpose of this API is to handle the case of failure to remove the
	// following retention leases after the unfollow API is invoked.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-forget-follower.html
	ForgetFollower ccr_forget_follower.NewForgetFollower
	// Get auto-follow patterns.
	//
	// Get cross-cluster replication auto-follow patterns.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-auto-follow-pattern.html
	GetAutoFollowPattern ccr_get_auto_follow_pattern.NewGetAutoFollowPattern
	// Pause an auto-follow pattern.
	//
	// Pause a cross-cluster replication auto-follow pattern.
	// When the API returns, the auto-follow pattern is inactive.
	// New indices that are created on the remote cluster and match the auto-follow
	// patterns are ignored.
	//
	// You can resume auto-following with the resume auto-follow pattern API.
	// When it resumes, the auto-follow pattern is active again and automatically
	// configures follower indices for newly created indices on the remote cluster
	// that match its patterns.
	// Remote indices that were created while the pattern was paused will also be
	// followed, unless they have been deleted or closed in the interim.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-pause-auto-follow-pattern.html
	PauseAutoFollowPattern ccr_pause_auto_follow_pattern.NewPauseAutoFollowPattern
	// Pause a follower.
	//
	// Pause a cross-cluster replication follower index.
	// The follower index will not fetch any additional operations from the leader
	// index.
	// You can resume following with the resume follower API.
	// You can pause and resume a follower index to change the configuration of the
	// following task.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-pause-follow.html
	PauseFollow ccr_pause_follow.NewPauseFollow
	// Create or update auto-follow patterns.
	// Create a collection of cross-cluster replication auto-follow patterns for a
	// remote cluster.
	// Newly created indices on the remote cluster that match any of the patterns
	// are automatically configured as follower indices.
	// Indices on the remote cluster that were created before the auto-follow
	// pattern was created will not be auto-followed even if they match the pattern.
	//
	// This API can also be used to update auto-follow patterns.
	// NOTE: Follower indices that were configured automatically before updating an
	// auto-follow pattern will remain unchanged even if they do not match against
	// the new patterns.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-auto-follow-pattern.html
	PutAutoFollowPattern ccr_put_auto_follow_pattern.NewPutAutoFollowPattern
	// Resume an auto-follow pattern.
	//
	// Resume a cross-cluster replication auto-follow pattern that was paused.
	// The auto-follow pattern will resume configuring following indices for newly
	// created indices that match its patterns on the remote cluster.
	// Remote indices created while the pattern was paused will also be followed
	// unless they have been deleted or closed in the interim.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-resume-auto-follow-pattern.html
	ResumeAutoFollowPattern ccr_resume_auto_follow_pattern.NewResumeAutoFollowPattern
	// Resume a follower.
	// Resume a cross-cluster replication follower index that was paused.
	// The follower index could have been paused with the pause follower API.
	// Alternatively it could be paused due to replication that cannot be retried
	// due to failures during following tasks.
	// When this API returns, the follower index will resume fetching operations
	// from the leader index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-resume-follow.html
	ResumeFollow ccr_resume_follow.NewResumeFollow
	// Get cross-cluster replication stats.
	//
	// This API returns stats about auto-following and the same shard-level stats as
	// the get follower stats API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-stats.html
	Stats ccr_stats.NewStats
	// Unfollow an index.
	//
	// Convert a cross-cluster replication follower index to a regular index.
	// The API stops the following task associated with a follower index and removes
	// index metadata and settings associated with cross-cluster replication.
	// The follower index must be paused and closed before you call the unfollow
	// API.
	//
	// > info
	// > Currently cross-cluster replication does not support converting an existing
	// regular index to a follower index. Converting a follower index to a regular
	// index is an irreversible operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-unfollow.html
	Unfollow ccr_unfollow.NewUnfollow
}

type Cluster struct {
	// Explain the shard allocations.
	// Get explanations for shard allocations in the cluster.
	// For unassigned shards, it provides an explanation for why the shard is
	// unassigned.
	// For assigned shards, it provides an explanation for why the shard is
	// remaining on its current node and has not moved or rebalanced to another
	// node.
	// This API can be very useful when attempting to diagnose why a shard is
	// unassigned or why a shard continues to remain on its current node when you
	// might expect otherwise.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-allocation-explain.html
	AllocationExplain cluster_allocation_explain.NewAllocationExplain
	// Delete component templates.
	// Component templates are building blocks for constructing index templates that
	// specify index mappings, settings, and aliases.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	DeleteComponentTemplate cluster_delete_component_template.NewDeleteComponentTemplate
	// Clear cluster voting config exclusions.
	// Remove master-eligible nodes from the voting configuration exclusion list.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-cluster-post-voting-config-exclusions
	DeleteVotingConfigExclusions cluster_delete_voting_config_exclusions.NewDeleteVotingConfigExclusions
	// Check component templates.
	// Returns information about whether a particular component template exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	ExistsComponentTemplate cluster_exists_component_template.NewExistsComponentTemplate
	// Get component templates.
	// Get information about component templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	GetComponentTemplate cluster_get_component_template.NewGetComponentTemplate
	// Get cluster-wide settings.
	// By default, it returns only settings that have been explicitly defined.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-get-settings.html
	GetSettings cluster_get_settings.NewGetSettings
	// Get the cluster health status.
	// You can also use the API to get the health status of only specified data
	// streams and indices.
	// For data streams, the API retrieves the health status of the stream’s backing
	// indices.
	//
	// The cluster health status is: green, yellow or red.
	// On the shard level, a red status indicates that the specific shard is not
	// allocated in the cluster. Yellow means that the primary shard is allocated
	// but replicas are not. Green means that all shards are allocated.
	// The index level status is controlled by the worst shard status.
	//
	// One of the main benefits of the API is the ability to wait until the cluster
	// reaches a certain high watermark health level.
	// The cluster status is controlled by the worst index status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html
	Health cluster_health.NewHealth
	// Get cluster info.
	// Returns basic information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-info.html
	Info cluster_info.NewInfo
	// Get the pending cluster tasks.
	// Get information about cluster-level changes (such as create index, update
	// mapping, allocate or fail shard) that have not yet taken effect.
	//
	// NOTE: This API returns a list of any pending updates to the cluster state.
	// These are distinct from the tasks reported by the task management API which
	// include periodic tasks and tasks initiated by the user, such as node stats,
	// search queries, or create index requests.
	// However, if a user-initiated task such as a create index command causes a
	// cluster state update, the activity of this task might be reported by both
	// task api and pending cluster tasks API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-pending.html
	PendingTasks cluster_pending_tasks.NewPendingTasks
	// Update voting configuration exclusions.
	// Update the cluster voting config exclusions by node IDs or node names.
	// By default, if there are more than three master-eligible nodes in the cluster
	// and you remove fewer than half of the master-eligible nodes in the cluster at
	// once, the voting configuration automatically shrinks.
	// If you want to shrink the voting configuration to contain fewer than three
	// nodes or to remove half or more of the master-eligible nodes in the cluster
	// at once, use this API to remove departing nodes from the voting configuration
	// manually.
	// The API adds an entry for each specified node to the cluster’s voting
	// configuration exclusions list.
	// It then waits until the cluster has reconfigured its voting configuration to
	// exclude the specified nodes.
	//
	// Clusters should have no voting configuration exclusions in normal operation.
	// Once the excluded nodes have stopped, clear the voting configuration
	// exclusions with `DELETE /_cluster/voting_config_exclusions`.
	// This API waits for the nodes to be fully removed from the cluster before it
	// returns.
	// If your cluster has voting configuration exclusions for nodes that you no
	// longer intend to remove, use `DELETE
	// /_cluster/voting_config_exclusions?wait_for_removal=false` to clear the
	// voting configuration exclusions without waiting for the nodes to leave the
	// cluster.
	//
	// A response to `POST /_cluster/voting_config_exclusions` with an HTTP status
	// code of 200 OK guarantees that the node has been removed from the voting
	// configuration and will not be reinstated until the voting configuration
	// exclusions are cleared by calling `DELETE
	// /_cluster/voting_config_exclusions`.
	// If the call to `POST /_cluster/voting_config_exclusions` fails or returns a
	// response with an HTTP status code other than 200 OK then the node may not
	// have been removed from the voting configuration.
	// In that case, you may safely retry the call.
	//
	// NOTE: Voting exclusions are required only when you remove at least half of
	// the master-eligible nodes from a cluster in a short time period.
	// They are not required when removing master-ineligible nodes or when removing
	// fewer than half of the master-eligible nodes.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-cluster-post-voting-config-exclusions
	PostVotingConfigExclusions cluster_post_voting_config_exclusions.NewPostVotingConfigExclusions
	// Create or update a component template.
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
	//
	// **Applying component templates**
	//
	// You cannot directly apply a component template to a data stream or index.
	// To be applied, a component template must be included in an index template's
	// `composed_of` list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	PutComponentTemplate cluster_put_component_template.NewPutComponentTemplate
	// Update the cluster settings.
	// Configure and update dynamic settings on a running cluster.
	// You can also configure dynamic settings locally on an unstarted or shut down
	// node in `elasticsearch.yml`.
	//
	// Updates made with this API can be persistent, which apply across cluster
	// restarts, or transient, which reset after a cluster restart.
	// You can also reset transient or persistent settings by assigning them a null
	// value.
	//
	// If you configure the same setting using multiple methods, Elasticsearch
	// applies the settings in following order of precedence: 1) Transient setting;
	// 2) Persistent setting; 3) `elasticsearch.yml` setting; 4) Default setting
	// value.
	// For example, you can apply a transient setting to override a persistent
	// setting or `elasticsearch.yml` setting.
	// However, a change to an `elasticsearch.yml` setting will not override a
	// defined transient or persistent setting.
	//
	// TIP: In Elastic Cloud, use the user settings feature to configure all cluster
	// settings. This method automatically rejects unsafe settings that could break
	// your cluster.
	// If you run Elasticsearch on your own hardware, use this API to configure
	// dynamic cluster settings.
	// Only use `elasticsearch.yml` for static cluster settings and node settings.
	// The API doesn’t require a restart and ensures a setting’s value is the same
	// on all nodes.
	//
	// WARNING: Transient cluster settings are no longer recommended. Use persistent
	// cluster settings instead.
	// If a cluster becomes unstable, transient settings can clear unexpectedly,
	// resulting in a potentially undesired cluster configuration.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-update-settings.html
	PutSettings cluster_put_settings.NewPutSettings
	// Get remote cluster information.
	//
	// Get information about configured remote clusters.
	// The API returns connection and endpoint information keyed by the configured
	// remote cluster alias.
	//
	// > info
	// > This API returns information that reflects current state on the local
	// cluster.
	// > The `connected` field does not necessarily reflect whether a remote cluster
	// is down or unavailable, only whether there is currently an open connection to
	// it.
	// > Elasticsearch does not spontaneously try to reconnect to a disconnected
	// remote cluster.
	// > To trigger a reconnection, attempt a cross-cluster search, ES|QL
	// cross-cluster search, or try the `/_resolve/cluster` endpoint.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-remote-info.html
	RemoteInfo cluster_remote_info.NewRemoteInfo
	// Reroute the cluster.
	// Manually change the allocation of individual shards in the cluster.
	// For example, a shard can be moved from one node to another explicitly, an
	// allocation can be canceled, and an unassigned shard can be explicitly
	// allocated to a specific node.
	//
	// It is important to note that after processing any reroute commands
	// Elasticsearch will perform rebalancing as normal (respecting the values of
	// settings such as `cluster.routing.rebalance.enable`) in order to remain in a
	// balanced state.
	// For example, if the requested allocation includes moving a shard from node1
	// to node2 then this may cause a shard to be moved from node2 back to node1 to
	// even things out.
	//
	// The cluster can be set to disable allocations using the
	// `cluster.routing.allocation.enable` setting.
	// If allocations are disabled then the only allocations that will be performed
	// are explicit ones given using the reroute command, and consequent allocations
	// due to rebalancing.
	//
	// The cluster will attempt to allocate a shard a maximum of
	// `index.allocation.max_retries` times in a row (defaults to `5`), before
	// giving up and leaving the shard unallocated.
	// This scenario can be caused by structural problems such as having an analyzer
	// which refers to a stopwords file which doesn’t exist on all nodes.
	//
	// Once the problem has been corrected, allocation can be manually retried by
	// calling the reroute API with the `?retry_failed` URI query parameter, which
	// will attempt a single retry round for these shards.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-reroute.html
	Reroute cluster_reroute.NewReroute
	// Get the cluster state.
	// Get comprehensive information about the state of the cluster.
	//
	// The cluster state is an internal data structure which keeps track of a
	// variety of information needed by every node, including the identity and
	// attributes of the other nodes in the cluster; cluster-wide settings; index
	// metadata, including the mapping and settings for each index; the location and
	// status of every shard copy in the cluster.
	//
	// The elected master node ensures that every node in the cluster has a copy of
	// the same cluster state.
	// This API lets you retrieve a representation of this internal state for
	// debugging or diagnostic purposes.
	// You may need to consult the Elasticsearch source code to determine the
	// precise meaning of the response.
	//
	// By default the API will route requests to the elected master node since this
	// node is the authoritative source of cluster states.
	// You can also retrieve the cluster state held on the node handling the API
	// request by adding the `?local=true` query parameter.
	//
	// Elasticsearch may need to expend significant effort to compute a response to
	// this API in larger clusters, and the response may comprise a very large
	// quantity of data.
	// If you use this API repeatedly, your cluster may become unstable.
	//
	// WARNING: The response is a representation of an internal data structure.
	// Its format is not subject to the same compatibility guarantees as other more
	// stable APIs and may change from version to version.
	// Do not query this API using external monitoring tools.
	// Instead, obtain the information you require using other more stable cluster
	// APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-state.html
	State cluster_state.NewState
	// Get cluster statistics.
	// Get basic index metrics (shard numbers, store size, memory usage) and
	// information about the current nodes that form the cluster (number, roles, os,
	// jvm versions, memory usage, cpu and installed plugins).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-stats.html
	Stats cluster_stats.NewStats
}

type Connector struct {
	// Check in a connector.
	//
	// Update the `last_seen` field in the connector and set it to the current
	// timestamp.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/check-in-connector-api.html
	CheckIn connector_check_in.NewCheckIn
	// Delete a connector.
	//
	// Removes a connector and associated sync jobs.
	// This is a destructive action that is not recoverable.
	// NOTE: This action doesn’t delete any API keys, ingest pipelines, or data
	// indices associated with the connector.
	// These need to be removed manually.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-connector-api.html
	Delete connector_delete.NewDelete
	// Get a connector.
	//
	// Get the details about a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-connector-api.html
	Get connector_get.NewGet
	// Update the connector last sync stats.
	//
	// Update the fields related to the last sync of a connector.
	// This action is used for analytics and monitoring.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-last-sync-api.html
	LastSync connector_last_sync.NewLastSync
	// Get all connectors.
	//
	// Get information about all connectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-connector-api.html
	List connector_list.NewList
	// Create a connector.
	//
	// Connectors are Elasticsearch integrations that bring content from third-party
	// data sources, which can be deployed on Elastic Cloud or hosted on your own
	// infrastructure.
	// Elastic managed connectors (Native connectors) are a managed service on
	// Elastic Cloud.
	// Self-managed connectors (Connector clients) are self-managed on your
	// infrastructure.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-api.html
	Post connector_post.NewPost
	// Create or update a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-api.html
	Put connector_put.NewPut
	// Creates a secret for a Connector.
	//
	SecretPost connector_secret_post.NewSecretPost
	// Cancel a connector sync job.
	//
	// Cancel a connector sync job, which sets the status to cancelling and updates
	// `cancellation_requested_at` to the current time.
	// The connector service is then responsible for setting the status of connector
	// sync jobs to cancelled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cancel-connector-sync-job-api.html
	SyncJobCancel connector_sync_job_cancel.NewSyncJobCancel
	// Check in a connector sync job.
	// Check in a connector sync job and set the `last_seen` field to the current
	// time before updating it in the internal index.
	//
	// To sync data using self-managed connectors, you need to deploy the Elastic
	// connector service on your own infrastructure.
	// This service runs automatically on Elastic Cloud for Elastic managed
	// connectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/check-in-connector-sync-job-api.html
	SyncJobCheckIn connector_sync_job_check_in.NewSyncJobCheckIn
	// Claim a connector sync job.
	// This action updates the job status to `in_progress` and sets the `last_seen`
	// and `started_at` timestamps to the current time.
	// Additionally, it can set the `sync_cursor` property for the sync job.
	//
	// This API is not intended for direct connector management by users.
	// It supports the implementation of services that utilize the connector
	// protocol to communicate with Elasticsearch.
	//
	// To sync data using self-managed connectors, you need to deploy the Elastic
	// connector service on your own infrastructure.
	// This service runs automatically on Elastic Cloud for Elastic managed
	// connectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/claim-connector-sync-job-api.html
	SyncJobClaim connector_sync_job_claim.NewSyncJobClaim
	// Delete a connector sync job.
	//
	// Remove a connector sync job and its associated data.
	// This is a destructive action that is not recoverable.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-connector-sync-job-api.html
	SyncJobDelete connector_sync_job_delete.NewSyncJobDelete
	// Set a connector sync job error.
	// Set the `error` field for a connector sync job and set its `status` to
	// `error`.
	//
	// To sync data using self-managed connectors, you need to deploy the Elastic
	// connector service on your own infrastructure.
	// This service runs automatically on Elastic Cloud for Elastic managed
	// connectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/set-connector-sync-job-error-api.html
	SyncJobError connector_sync_job_error.NewSyncJobError
	// Get a connector sync job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-connector-sync-job-api.html
	SyncJobGet connector_sync_job_get.NewSyncJobGet
	// Get all connector sync jobs.
	//
	// Get information about all stored connector sync jobs listed by their creation
	// date in ascending order.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-connector-sync-jobs-api.html
	SyncJobList connector_sync_job_list.NewSyncJobList
	// Create a connector sync job.
	//
	// Create a connector sync job document in the internal index and initialize its
	// counters and timestamps with default values.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-sync-job-api.html
	SyncJobPost connector_sync_job_post.NewSyncJobPost
	// Set the connector sync job stats.
	// Stats include: `deleted_document_count`, `indexed_document_count`,
	// `indexed_document_volume`, and `total_document_count`.
	// You can also update `last_seen`.
	// This API is mainly used by the connector service for updating sync job
	// information.
	//
	// To sync data using self-managed connectors, you need to deploy the Elastic
	// connector service on your own infrastructure.
	// This service runs automatically on Elastic Cloud for Elastic managed
	// connectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/set-connector-sync-job-stats-api.html
	SyncJobUpdateStats connector_sync_job_update_stats.NewSyncJobUpdateStats
	// Activate the connector draft filter.
	//
	// Activates the valid draft filtering for a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-api.html
	UpdateActiveFiltering connector_update_active_filtering.NewUpdateActiveFiltering
	// Update the connector API key ID.
	//
	// Update the `api_key_id` and `api_key_secret_id` fields of a connector.
	// You can specify the ID of the API key used for authorization and the ID of
	// the connector secret where the API key is stored.
	// The connector secret ID is required only for Elastic managed (native)
	// connectors.
	// Self-managed connectors (connector clients) do not use this field.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-api-key-id-api.html
	UpdateApiKeyId connector_update_api_key_id.NewUpdateApiKeyId
	// Update the connector configuration.
	//
	// Update the configuration field in the connector document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-configuration-api.html
	UpdateConfiguration connector_update_configuration.NewUpdateConfiguration
	// Update the connector error field.
	//
	// Set the error field for the connector.
	// If the error provided in the request body is non-null, the connector’s status
	// is updated to error.
	// Otherwise, if the error is reset to null, the connector status is updated to
	// connected.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-error-api.html
	UpdateError connector_update_error.NewUpdateError
	// Update the connector features.
	// Update the connector features in the connector document.
	// This API can be used to control the following aspects of a connector:
	//
	// * document-level security
	// * incremental syncs
	// * advanced sync rules
	// * basic sync rules
	//
	// Normally, the running connector service automatically manages these features.
	// However, you can use this API to override the default behavior.
	//
	// To sync data using self-managed connectors, you need to deploy the Elastic
	// connector service on your own infrastructure.
	// This service runs automatically on Elastic Cloud for Elastic managed
	// connectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-features-api.html
	UpdateFeatures connector_update_features.NewUpdateFeatures
	// Update the connector filtering.
	//
	// Update the draft filtering configuration of a connector and marks the draft
	// validation state as edited.
	// The filtering draft is activated once validated by the running Elastic
	// connector service.
	// The filtering property is used to configure sync rules (both basic and
	// advanced) for a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-api.html
	UpdateFiltering connector_update_filtering.NewUpdateFiltering
	// Update the connector draft filtering validation.
	//
	// Update the draft filtering validation info for a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-validation-api.html
	UpdateFilteringValidation connector_update_filtering_validation.NewUpdateFilteringValidation
	// Update the connector index name.
	//
	// Update the `index_name` field of a connector, specifying the index where the
	// data ingested by the connector is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-index-name-api.html
	UpdateIndexName connector_update_index_name.NewUpdateIndexName
	// Update the connector name and description.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-name-description-api.html
	UpdateName connector_update_name.NewUpdateName
	// Update the connector is_native flag.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-native-api.html
	UpdateNative connector_update_native.NewUpdateNative
	// Update the connector pipeline.
	//
	// When you create a new connector, the configuration of an ingest pipeline is
	// populated with default settings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-pipeline-api.html
	UpdatePipeline connector_update_pipeline.NewUpdatePipeline
	// Update the connector scheduling.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-scheduling-api.html
	UpdateScheduling connector_update_scheduling.NewUpdateScheduling
	// Update the connector service type.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-service-type-api.html
	UpdateServiceType connector_update_service_type.NewUpdateServiceType
	// Update the connector status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-status-api.html
	UpdateStatus connector_update_status.NewUpdateStatus
}

type Core struct {
	// Bulk index or delete documents.
	// Perform multiple `index`, `create`, `delete`, and `update` actions in a
	// single request.
	// This reduces overhead and can greatly increase indexing speed.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or index alias:
	//
	// * To use the `create` action, you must have the `create_doc`, `create`,
	// `index`, or `write` index privilege. Data streams support only the `create`
	// action.
	// * To use the `index` action, you must have the `create`, `index`, or `write`
	// index privilege.
	// * To use the `delete` action, you must have the `delete` or `write` index
	// privilege.
	// * To use the `update` action, you must have the `index` or `write` index
	// privilege.
	// * To automatically create a data stream or index with a bulk API request, you
	// must have the `auto_configure`, `create_index`, or `manage` index privilege.
	// * To make the result of a bulk operation visible to search using the
	// `refresh` parameter, you must have the `maintenance` or `manage` index
	// privilege.
	//
	// Automatic data stream creation requires a matching index template with data
	// stream enabled.
	//
	// The actions are specified in the request body using a newline delimited JSON
	// (NDJSON) structure:
	//
	// ```
	// action_and_meta_data\n
	// optional_source\n
	// action_and_meta_data\n
	// optional_source\n
	// ....
	// action_and_meta_data\n
	// optional_source\n
	// ```
	//
	// The `index` and `create` actions expect a source on the next line and have
	// the same semantics as the `op_type` parameter in the standard index API.
	// A `create` action fails if a document with the same ID already exists in the
	// target
	// An `index` action adds or replaces a document as necessary.
	//
	// NOTE: Data streams support only the `create` action.
	// To update or delete a document in a data stream, you must target the backing
	// index containing the document.
	//
	// An `update` action expects that the partial doc, upsert, and script and its
	// options are specified on the next line.
	//
	// A `delete` action does not expect a source on the next line and has the same
	// semantics as the standard delete API.
	//
	// NOTE: The final line of data must end with a newline character (`\n`).
	// Each newline character may be preceded by a carriage return (`\r`).
	// When sending NDJSON data to the `_bulk` endpoint, use a `Content-Type` header
	// of `application/json` or `application/x-ndjson`.
	// Because this format uses literal newline characters (`\n`) as delimiters,
	// make sure that the JSON actions and sources are not pretty printed.
	//
	// If you provide a target in the request path, it is used for any actions that
	// don't explicitly specify an `_index` argument.
	//
	// A note on the format: the idea here is to make processing as fast as
	// possible.
	// As some of the actions are redirected to other shards on other nodes, only
	// `action_meta_data` is parsed on the receiving node side.
	//
	// Client libraries using this protocol should try and strive to do something
	// similar on the client side, and reduce buffering as much as possible.
	//
	// There is no "correct" number of actions to perform in a single bulk request.
	// Experiment with different settings to find the optimal size for your
	// particular workload.
	// Note that Elasticsearch limits the maximum size of a HTTP request to 100mb by
	// default so clients must ensure that no request exceeds this size.
	// It is not possible to index a single document that exceeds the size limit, so
	// you must pre-process any such documents into smaller pieces before sending
	// them to Elasticsearch.
	// For instance, split documents into pages or chapters before indexing them, or
	// store raw binary data in a system outside Elasticsearch and replace the raw
	// data with a link to the external system in the documents that you send to
	// Elasticsearch.
	//
	// **Client suppport for bulk requests**
	//
	// Some of the officially supported clients provide helpers to assist with bulk
	// requests and reindexing:
	//
	// * Go: Check out `esutil.BulkIndexer`
	// * Perl: Check out `Search::Elasticsearch::Client::5_0::Bulk` and
	// `Search::Elasticsearch::Client::5_0::Scroll`
	// * Python: Check out `elasticsearch.helpers.*`
	// * JavaScript: Check out `client.helpers.*`
	// * .NET: Check out `BulkAllObservable`
	// * PHP: Check out bulk indexing.
	//
	// **Submitting bulk requests with cURL**
	//
	// If you're providing text file input to `curl`, you must use the
	// `--data-binary` flag instead of plain `-d`.
	// The latter doesn't preserve newlines. For example:
	//
	// ```
	// $ cat requests
	// { "index" : { "_index" : "test", "_id" : "1" } }
	// { "field1" : "value1" }
	// $ curl -s -H "Content-Type: application/x-ndjson" -XPOST localhost:9200/_bulk
	// --data-binary "@requests"; echo
	// {"took":7, "errors": false,
	// "items":[{"index":{"_index":"test","_id":"1","_version":1,"result":"created","forced_refresh":false}}]}
	// ```
	//
	// **Optimistic concurrency control**
	//
	// Each `index` and `delete` action within a bulk API call may include the
	// `if_seq_no` and `if_primary_term` parameters in their respective action and
	// meta data lines.
	// The `if_seq_no` and `if_primary_term` parameters control how operations are
	// run, based on the last modification to existing documents. See Optimistic
	// concurrency control for more details.
	//
	// **Versioning**
	//
	// Each bulk item can include the version value using the `version` field.
	// It automatically follows the behavior of the index or delete operation based
	// on the `_version` mapping.
	// It also support the `version_type`.
	//
	// **Routing**
	//
	// Each bulk item can include the routing value using the `routing` field.
	// It automatically follows the behavior of the index or delete operation based
	// on the `_routing` mapping.
	//
	// NOTE: Data streams do not support custom routing unless they were created
	// with the `allow_custom_routing` setting enabled in the template.
	//
	// **Wait for active shards**
	//
	// When making bulk calls, you can set the `wait_for_active_shards` parameter to
	// require a minimum number of shard copies to be active before starting to
	// process the bulk request.
	//
	// **Refresh**
	//
	// Control when the changes made by this request are visible to search.
	//
	// NOTE: Only the shards that receive the bulk request will be affected by
	// refresh.
	// Imagine a `_bulk?refresh=wait_for` request with three documents in it that
	// happen to be routed to different shards in an index with five shards.
	// The request will only wait for those three shards to refresh.
	// The other two shards that make up the index do not participate in the `_bulk`
	// request at all.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Clear a scrolling search.
	// Clear the search context and results for a scrolling search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Close a point in time.
	// A point in time must be opened explicitly before being used in search
	// requests.
	// The `keep_alive` parameter tells Elasticsearch how long it should persist.
	// A point in time is automatically closed when the `keep_alive` period has
	// elapsed.
	// However, keeping points in time has a cost; close them as soon as they are no
	// longer required for search requests.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Count search results.
	// Get the number of documents matching a query.
	//
	// The query can be provided either by using a simple query string as a
	// parameter, or by defining Query DSL within the request body.
	// The query is optional. When no query is provided, the API uses `match_all` to
	// count all the documents.
	//
	// The count API supports multi-target syntax. You can run a single count API
	// search across multiple data streams and indices.
	//
	// The operation is broadcast across all shards.
	// For each shard ID group, a replica is chosen and the search is run against
	// it.
	// This means that replicas increase the scalability of the count.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Create a new document in the index.
	//
	// You can index a new JSON document with the `/<target>/_doc/` or
	// `/<target>/_create/<_id>` APIs
	// Using `_create` guarantees that the document is indexed only if it does not
	// already exist.
	// It returns a 409 response when a document with a same ID already exists in
	// the index.
	// To update an existing document, you must use the `/<target>/_doc/` API.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or index alias:
	//
	// * To add a document using the `PUT /<target>/_create/<_id>` or `POST
	// /<target>/_create/<_id>` request formats, you must have the `create_doc`,
	// `create`, `index`, or `write` index privilege.
	// * To automatically create a data stream or index with this API request, you
	// must have the `auto_configure`, `create_index`, or `manage` index privilege.
	//
	// Automatic data stream creation requires a matching index template with data
	// stream enabled.
	//
	// **Automatically create data streams and indices**
	//
	// If the request's target doesn't exist and matches an index template with a
	// `data_stream` definition, the index operation automatically creates the data
	// stream.
	//
	// If the target doesn't exist and doesn't match a data stream template, the
	// operation automatically creates the index and applies any matching index
	// templates.
	//
	// NOTE: Elasticsearch includes several built-in index templates. To avoid
	// naming collisions with these templates, refer to index pattern documentation.
	//
	// If no mapping exists, the index operation creates a dynamic mapping.
	// By default, new fields and objects are automatically added to the mapping if
	// needed.
	//
	// Automatic index creation is controlled by the `action.auto_create_index`
	// setting.
	// If it is `true`, any index can be created automatically.
	// You can modify this setting to explicitly allow or block automatic creation
	// of indices that match specified patterns or set it to `false` to turn off
	// automatic index creation entirely.
	// Specify a comma-separated list of patterns you want to allow or prefix each
	// pattern with `+` or `-` to indicate whether it should be allowed or blocked.
	// When a list is specified, the default behaviour is to disallow.
	//
	// NOTE: The `action.auto_create_index` setting affects the automatic creation
	// of indices only.
	// It does not affect the creation of data streams.
	//
	// **Routing**
	//
	// By default, shard placement — or routing — is controlled by using a hash of
	// the document's ID value.
	// For more explicit control, the value fed into the hash function used by the
	// router can be directly specified on a per-operation basis using the `routing`
	// parameter.
	//
	// When setting up explicit mapping, you can also use the `_routing` field to
	// direct the index operation to extract the routing value from the document
	// itself.
	// This does come at the (very minimal) cost of an additional document parsing
	// pass.
	// If the `_routing` mapping is defined and set to be required, the index
	// operation will fail if no routing value is provided or extracted.
	//
	// NOTE: Data streams do not support custom routing unless they were created
	// with the `allow_custom_routing` setting enabled in the template.
	//
	// **Distributed**
	//
	// The index operation is directed to the primary shard based on its route and
	// performed on the actual node containing this shard.
	// After the primary shard completes the operation, if needed, the update is
	// distributed to applicable replicas.
	//
	// **Active shards**
	//
	// To improve the resiliency of writes to the system, indexing operations can be
	// configured to wait for a certain number of active shard copies before
	// proceeding with the operation.
	// If the requisite number of active shard copies are not available, then the
	// write operation must wait and retry, until either the requisite shard copies
	// have started or a timeout occurs.
	// By default, write operations only wait for the primary shards to be active
	// before proceeding (that is to say `wait_for_active_shards` is `1`).
	// This default can be overridden in the index settings dynamically by setting
	// `index.write.wait_for_active_shards`.
	// To alter this behavior per operation, use the `wait_for_active_shards
	// request` parameter.
	//
	// Valid values are all or any positive integer up to the total number of
	// configured copies per shard in the index (which is `number_of_replicas`+1).
	// Specifying a negative value or a number greater than the number of shard
	// copies will throw an error.
	//
	// For example, suppose you have a cluster of three nodes, A, B, and C and you
	// create an index index with the number of replicas set to 3 (resulting in 4
	// shard copies, one more copy than there are nodes).
	// If you attempt an indexing operation, by default the operation will only
	// ensure the primary copy of each shard is available before proceeding.
	// This means that even if B and C went down and A hosted the primary shard
	// copies, the indexing operation would still proceed with only one copy of the
	// data.
	// If `wait_for_active_shards` is set on the request to `3` (and all three nodes
	// are up), the indexing operation will require 3 active shard copies before
	// proceeding.
	// This requirement should be met because there are 3 active nodes in the
	// cluster, each one holding a copy of the shard.
	// However, if you set `wait_for_active_shards` to `all` (or to `4`, which is
	// the same in this situation), the indexing operation will not proceed as you
	// do not have all 4 copies of each shard active in the index.
	// The operation will timeout unless a new node is brought up in the cluster to
	// host the fourth copy of the shard.
	//
	// It is important to note that this setting greatly reduces the chances of the
	// write operation not writing to the requisite number of shard copies, but it
	// does not completely eliminate the possibility, because this check occurs
	// before the write operation starts.
	// After the write operation is underway, it is still possible for replication
	// to fail on any number of shard copies but still succeed on the primary.
	// The `_shards` section of the API response reveals the number of shard copies
	// on which replication succeeded and failed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Delete a document.
	//
	// Remove a JSON document from the specified index.
	//
	// NOTE: You cannot send deletion requests directly to a data stream.
	// To delete a document in a data stream, you must target the backing index
	// containing the document.
	//
	// **Optimistic concurrency control**
	//
	// Delete operations can be made conditional and only be performed if the last
	// modification to the document was assigned the sequence number and primary
	// term specified by the `if_seq_no` and `if_primary_term` parameters.
	// If a mismatch is detected, the operation will result in a
	// `VersionConflictException` and a status code of `409`.
	//
	// **Versioning**
	//
	// Each document indexed is versioned.
	// When deleting a document, the version can be specified to make sure the
	// relevant document you are trying to delete is actually being deleted and it
	// has not changed in the meantime.
	// Every write operation run on a document, deletes included, causes its version
	// to be incremented.
	// The version number of a deleted document remains available for a short time
	// after deletion to allow for control of concurrent operations.
	// The length of time for which a deleted document's version remains available
	// is determined by the `index.gc_deletes` index setting.
	//
	// **Routing**
	//
	// If routing is used during indexing, the routing value also needs to be
	// specified to delete a document.
	//
	// If the `_routing` mapping is set to `required` and no routing value is
	// specified, the delete API throws a `RoutingMissingException` and rejects the
	// request.
	//
	// For example:
	//
	// ```
	// DELETE /my-index-000001/_doc/1?routing=shard-1
	// ```
	//
	// This request deletes the document with ID 1, but it is routed based on the
	// user.
	// The document is not deleted if the correct routing is not specified.
	//
	// **Distributed**
	//
	// The delete operation gets hashed into a specific shard ID.
	// It then gets redirected into the primary shard within that ID group and
	// replicated (if needed) to shard replicas within that ID group.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Delete documents.
	//
	// Deletes documents that match the specified query.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or alias:
	//
	// * `read`
	// * `delete` or `write`
	//
	// You can specify the query criteria in the request URI or the request body
	// using the same syntax as the search API.
	// When you submit a delete by query request, Elasticsearch gets a snapshot of
	// the data stream or index when it begins processing the request and deletes
	// matching documents using internal versioning.
	// If a document changes between the time that the snapshot is taken and the
	// delete operation is processed, it results in a version conflict and the
	// delete operation fails.
	//
	// NOTE: Documents with a version equal to 0 cannot be deleted using delete by
	// query because internal versioning does not support 0 as a valid version
	// number.
	//
	// While processing a delete by query request, Elasticsearch performs multiple
	// search requests sequentially to find all of the matching documents to delete.
	// A bulk delete request is performed for each batch of matching documents.
	// If a search or bulk request is rejected, the requests are retried up to 10
	// times, with exponential back off.
	// If the maximum retry limit is reached, processing halts and all failed
	// requests are returned in the response.
	// Any delete requests that completed successfully still stick, they are not
	// rolled back.
	//
	// You can opt to count version conflicts instead of halting and returning by
	// setting `conflicts` to `proceed`.
	// Note that if you opt to count version conflicts the operation could attempt
	// to delete more documents from the source than `max_docs` until it has
	// successfully deleted `max_docs documents`, or it has gone through every
	// document in the source query.
	//
	// **Throttling delete requests**
	//
	// To control the rate at which delete by query issues batches of delete
	// operations, you can set `requests_per_second` to any positive decimal number.
	// This pads each batch with a wait time to throttle the rate.
	// Set `requests_per_second` to `-1` to disable throttling.
	//
	// Throttling uses a wait time between batches so that the internal scroll
	// requests can be given a timeout that takes the request padding into account.
	// The padding time is the difference between the batch size divided by the
	// `requests_per_second` and the time spent writing.
	// By default the batch size is `1000`, so if `requests_per_second` is set to
	// `500`:
	//
	// ```
	// target_time = 1000 / 500 per second = 2 seconds
	// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
	// ```
	//
	// Since the batch is issued as a single `_bulk` request, large batch sizes
	// cause Elasticsearch to create many requests and wait before starting the next
	// set.
	// This is "bursty" instead of "smooth".
	//
	// **Slicing**
	//
	// Delete by query supports sliced scroll to parallelize the delete process.
	// This can improve efficiency and provide a convenient way to break the request
	// down into smaller parts.
	//
	// Setting `slices` to `auto` lets Elasticsearch choose the number of slices to
	// use.
	// This setting will use one slice per shard, up to a certain limit.
	// If there are multiple source data streams or indices, it will choose the
	// number of slices based on the index or backing index with the smallest number
	// of shards.
	// Adding slices to the delete by query operation creates sub-requests which
	// means it has some quirks:
	//
	// * You can see these requests in the tasks APIs. These sub-requests are
	// "child" tasks of the task for the request with slices.
	// * Fetching the status of the task for the request with slices only contains
	// the status of completed slices.
	// * These sub-requests are individually addressable for things like
	// cancellation and rethrottling.
	// * Rethrottling the request with `slices` will rethrottle the unfinished
	// sub-request proportionally.
	// * Canceling the request with `slices` will cancel each sub-request.
	// * Due to the nature of `slices` each sub-request won't get a perfectly even
	// portion of the documents. All documents will be addressed, but some slices
	// may be larger than others. Expect larger slices to have a more even
	// distribution.
	// * Parameters like `requests_per_second` and `max_docs` on a request with
	// `slices` are distributed proportionally to each sub-request. Combine that
	// with the earlier point about distribution being uneven and you should
	// conclude that using `max_docs` with `slices` might not result in exactly
	// `max_docs` documents being deleted.
	// * Each sub-request gets a slightly different snapshot of the source data
	// stream or index though these are all taken at approximately the same time.
	//
	// If you're slicing manually or otherwise tuning automatic slicing, keep in
	// mind that:
	//
	// * Query performance is most efficient when the number of slices is equal to
	// the number of shards in the index or backing index. If that number is large
	// (for example, 500), choose a lower number as too many `slices` hurts
	// performance. Setting `slices` higher than the number of shards generally does
	// not improve efficiency and adds overhead.
	// * Delete performance scales linearly across available resources with the
	// number of slices.
	//
	// Whether query or delete performance dominates the runtime depends on the
	// documents being reindexed and cluster resources.
	//
	// **Cancel a delete by query operation**
	//
	// Any delete by query can be canceled using the task cancel API. For example:
	//
	// ```
	// POST _tasks/r1A2WoRbTwKZ516z6NEs5A:36619/_cancel
	// ```
	//
	// The task ID can be found by using the get tasks API.
	//
	// Cancellation should happen quickly but might take a few seconds.
	// The get task status API will continue to list the delete by query task until
	// this task checks that it has been cancelled and terminates itself.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Throttle a delete by query operation.
	//
	// Change the number of requests per second for a particular delete by query
	// operation.
	// Rethrottling that speeds up the query takes effect immediately but
	// rethrotting that slows down the query takes effect after completing the
	// current batch to prevent scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html#docs-delete-by-query-rethrottle
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	// Delete a script or search template.
	// Deletes a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-stored-script-api.html
	DeleteScript core_delete_script.NewDeleteScript
	// Check a document.
	//
	// Verify that a document exists.
	// For example, check to see if a document with the `_id` 0 exists:
	//
	// ```
	// HEAD my-index-000001/_doc/0
	// ```
	//
	// If the document exists, the API returns a status code of `200 - OK`.
	// If the document doesn’t exist, the API returns `404 - Not Found`.
	//
	// **Versioning support**
	//
	// You can use the `version` parameter to check the document only if its current
	// version is equal to the specified one.
	//
	// Internally, Elasticsearch has marked the old document as deleted and added an
	// entirely new document.
	// The old version of the document doesn't disappear immediately, although you
	// won't be able to access it.
	// Elasticsearch cleans up deleted documents in the background as you continue
	// to index more data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Check for a document source.
	//
	// Check whether a document source exists in an index.
	// For example:
	//
	// ```
	// HEAD my-index-000001/_source/1
	// ```
	//
	// A document's source is not available if it is disabled in the mapping.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Explain a document match result.
	// Get information about why a specific document matches, or doesn't match, a
	// query.
	// It computes a score explanation for a query and a specific document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// Get the field capabilities.
	//
	// Get information about the capabilities of fields among multiple indices.
	//
	// For data streams, the API returns field capabilities among the stream’s
	// backing indices.
	// It returns runtime fields like any other field.
	// For example, a runtime field with a type of keyword is returned the same as
	// any other field that belongs to the `keyword` family.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Get a document by its ID.
	//
	// Get a document and its source or stored fields from an index.
	//
	// By default, this API is realtime and is not affected by the refresh rate of
	// the index (when data will become visible for search).
	// In the case where stored fields are requested with the `stored_fields`
	// parameter and the document has been updated but is not yet refreshed, the API
	// will have to parse and analyze the source to extract the stored fields.
	// To turn off realtime behavior, set the `realtime` parameter to false.
	//
	// **Source filtering**
	//
	// By default, the API returns the contents of the `_source` field unless you
	// have used the `stored_fields` parameter or the `_source` field is turned off.
	// You can turn off `_source` retrieval by using the `_source` parameter:
	//
	// ```
	// GET my-index-000001/_doc/0?_source=false
	// ```
	//
	// If you only need one or two fields from the `_source`, use the
	// `_source_includes` or `_source_excludes` parameters to include or filter out
	// particular fields.
	// This can be helpful with large documents where partial retrieval can save on
	// network overhead
	// Both parameters take a comma separated list of fields or wildcard
	// expressions.
	// For example:
	//
	// ```
	// GET my-index-000001/_doc/0?_source_includes=*.id&_source_excludes=entities
	// ```
	//
	// If you only want to specify includes, you can use a shorter notation:
	//
	// ```
	// GET my-index-000001/_doc/0?_source=*.id
	// ```
	//
	// **Routing**
	//
	// If routing is used during indexing, the routing value also needs to be
	// specified to retrieve a document.
	// For example:
	//
	// ```
	// GET my-index-000001/_doc/2?routing=user1
	// ```
	//
	// This request gets the document with ID 2, but it is routed based on the user.
	// The document is not fetched if the correct routing is not specified.
	//
	// **Distributed**
	//
	// The GET operation is hashed into a specific shard ID.
	// It is then redirected to one of the replicas within that shard ID and returns
	// the result.
	// The replicas are the primary shard and its replicas within that shard ID
	// group.
	// This means that the more replicas you have, the better your GET scaling will
	// be.
	//
	// **Versioning support**
	//
	// You can use the `version` parameter to retrieve the document only if its
	// current version is equal to the specified one.
	//
	// Internally, Elasticsearch has marked the old document as deleted and added an
	// entirely new document.
	// The old version of the document doesn't disappear immediately, although you
	// won't be able to access it.
	// Elasticsearch cleans up deleted documents in the background as you continue
	// to index more data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Get a script or search template.
	// Retrieves a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-stored-script-api.html
	GetScript core_get_script.NewGetScript
	// Get script contexts.
	//
	// Get a list of supported script contexts and their methods.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-script-contexts-api.html
	GetScriptContext core_get_script_context.NewGetScriptContext
	// Get script languages.
	//
	// Get a list of available script types, languages, and contexts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-script-languages-api.html
	GetScriptLanguages core_get_script_languages.NewGetScriptLanguages
	// Get a document's source.
	//
	// Get the source of a document.
	// For example:
	//
	// ```
	// GET my-index-000001/_source/1
	// ```
	//
	// You can use the source filtering parameters to control which parts of the
	// `_source` are returned:
	//
	// ```
	// GET
	// my-index-000001/_source/1/?_source_includes=*.id&_source_excludes=entities
	// ```
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Get the cluster health.
	// Get a report with the health status of an Elasticsearch cluster.
	// The report contains a list of indicators that compose Elasticsearch
	// functionality.
	//
	// Each indicator has a health status of: green, unknown, yellow or red.
	// The indicator will provide an explanation and metadata describing the reason
	// for its current health status.
	//
	// The cluster’s status is controlled by the worst indicator status.
	//
	// In the event that an indicator’s status is non-green, a list of impacts may
	// be present in the indicator result which detail the functionalities that are
	// negatively affected by the health issue.
	// Each impact carries with it a severity level, an area of the system that is
	// affected, and a simple description of the impact on the system.
	//
	// Some health indicators can determine the root cause of a health problem and
	// prescribe a set of steps that can be performed in order to improve the health
	// of the system.
	// The root cause and remediation steps are encapsulated in a diagnosis.
	// A diagnosis contains a cause detailing a root cause analysis, an action
	// containing a brief description of the steps to take to fix the problem, the
	// list of affected resources (if applicable), and a detailed step-by-step
	// troubleshooting guide to fix the diagnosed problem.
	//
	// NOTE: The health indicators perform root cause analysis of non-green health
	// statuses. This can be computationally expensive when called frequently.
	// When setting up automated polling of the API for health status, set verbose
	// to false to disable the more expensive analysis logic.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html
	HealthReport core_health_report.NewHealthReport
	// Create or update a document in an index.
	//
	// Add a JSON document to the specified data stream or index and make it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	//
	// NOTE: You cannot use this API to send update requests for existing documents
	// in a data stream.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or index alias:
	//
	// * To add or overwrite a document using the `PUT /<target>/_doc/<_id>` request
	// format, you must have the `create`, `index`, or `write` index privilege.
	// * To add a document using the `POST /<target>/_doc/` request format, you must
	// have the `create_doc`, `create`, `index`, or `write` index privilege.
	// * To automatically create a data stream or index with this API request, you
	// must have the `auto_configure`, `create_index`, or `manage` index privilege.
	//
	// Automatic data stream creation requires a matching index template with data
	// stream enabled.
	//
	// NOTE: Replica shards might not all be started when an indexing operation
	// returns successfully.
	// By default, only the primary is required. Set `wait_for_active_shards` to
	// change this default behavior.
	//
	// **Automatically create data streams and indices**
	//
	// If the request's target doesn't exist and matches an index template with a
	// `data_stream` definition, the index operation automatically creates the data
	// stream.
	//
	// If the target doesn't exist and doesn't match a data stream template, the
	// operation automatically creates the index and applies any matching index
	// templates.
	//
	// NOTE: Elasticsearch includes several built-in index templates. To avoid
	// naming collisions with these templates, refer to index pattern documentation.
	//
	// If no mapping exists, the index operation creates a dynamic mapping.
	// By default, new fields and objects are automatically added to the mapping if
	// needed.
	//
	// Automatic index creation is controlled by the `action.auto_create_index`
	// setting.
	// If it is `true`, any index can be created automatically.
	// You can modify this setting to explicitly allow or block automatic creation
	// of indices that match specified patterns or set it to `false` to turn off
	// automatic index creation entirely.
	// Specify a comma-separated list of patterns you want to allow or prefix each
	// pattern with `+` or `-` to indicate whether it should be allowed or blocked.
	// When a list is specified, the default behaviour is to disallow.
	//
	// NOTE: The `action.auto_create_index` setting affects the automatic creation
	// of indices only.
	// It does not affect the creation of data streams.
	//
	// **Optimistic concurrency control**
	//
	// Index operations can be made conditional and only be performed if the last
	// modification to the document was assigned the sequence number and primary
	// term specified by the `if_seq_no` and `if_primary_term` parameters.
	// If a mismatch is detected, the operation will result in a
	// `VersionConflictException` and a status code of `409`.
	//
	// **Routing**
	//
	// By default, shard placement — or routing — is controlled by using a hash of
	// the document's ID value.
	// For more explicit control, the value fed into the hash function used by the
	// router can be directly specified on a per-operation basis using the `routing`
	// parameter.
	//
	// When setting up explicit mapping, you can also use the `_routing` field to
	// direct the index operation to extract the routing value from the document
	// itself.
	// This does come at the (very minimal) cost of an additional document parsing
	// pass.
	// If the `_routing` mapping is defined and set to be required, the index
	// operation will fail if no routing value is provided or extracted.
	//
	// NOTE: Data streams do not support custom routing unless they were created
	// with the `allow_custom_routing` setting enabled in the template.
	//
	// **Distributed**
	//
	// The index operation is directed to the primary shard based on its route and
	// performed on the actual node containing this shard.
	// After the primary shard completes the operation, if needed, the update is
	// distributed to applicable replicas.
	//
	// **Active shards**
	//
	// To improve the resiliency of writes to the system, indexing operations can be
	// configured to wait for a certain number of active shard copies before
	// proceeding with the operation.
	// If the requisite number of active shard copies are not available, then the
	// write operation must wait and retry, until either the requisite shard copies
	// have started or a timeout occurs.
	// By default, write operations only wait for the primary shards to be active
	// before proceeding (that is to say `wait_for_active_shards` is `1`).
	// This default can be overridden in the index settings dynamically by setting
	// `index.write.wait_for_active_shards`.
	// To alter this behavior per operation, use the `wait_for_active_shards
	// request` parameter.
	//
	// Valid values are all or any positive integer up to the total number of
	// configured copies per shard in the index (which is `number_of_replicas`+1).
	// Specifying a negative value or a number greater than the number of shard
	// copies will throw an error.
	//
	// For example, suppose you have a cluster of three nodes, A, B, and C and you
	// create an index index with the number of replicas set to 3 (resulting in 4
	// shard copies, one more copy than there are nodes).
	// If you attempt an indexing operation, by default the operation will only
	// ensure the primary copy of each shard is available before proceeding.
	// This means that even if B and C went down and A hosted the primary shard
	// copies, the indexing operation would still proceed with only one copy of the
	// data.
	// If `wait_for_active_shards` is set on the request to `3` (and all three nodes
	// are up), the indexing operation will require 3 active shard copies before
	// proceeding.
	// This requirement should be met because there are 3 active nodes in the
	// cluster, each one holding a copy of the shard.
	// However, if you set `wait_for_active_shards` to `all` (or to `4`, which is
	// the same in this situation), the indexing operation will not proceed as you
	// do not have all 4 copies of each shard active in the index.
	// The operation will timeout unless a new node is brought up in the cluster to
	// host the fourth copy of the shard.
	//
	// It is important to note that this setting greatly reduces the chances of the
	// write operation not writing to the requisite number of shard copies, but it
	// does not completely eliminate the possibility, because this check occurs
	// before the write operation starts.
	// After the write operation is underway, it is still possible for replication
	// to fail on any number of shard copies but still succeed on the primary.
	// The `_shards` section of the API response reveals the number of shard copies
	// on which replication succeeded and failed.
	//
	// **No operation (noop) updates**
	//
	// When updating a document by using this API, a new version of the document is
	// always created even if the document hasn't changed.
	// If this isn't acceptable use the `_update` API with `detect_noop` set to
	// `true`.
	// The `detect_noop` option isn't available on this API because it doesn’t fetch
	// the old source and isn't able to compare it against the new source.
	//
	// There isn't a definitive rule for when noop updates aren't acceptable.
	// It's a combination of lots of factors like how frequently your data source
	// sends updates that are actually noops and how many queries per second
	// Elasticsearch runs on the shard receiving the updates.
	//
	// **Versioning**
	//
	// Each indexed document is given a version number.
	// By default, internal versioning is used that starts at 1 and increments with
	// each update, deletes included.
	// Optionally, the version number can be set to an external value (for example,
	// if maintained in a database).
	// To enable this functionality, `version_type` should be set to `external`.
	// The value provided must be a numeric, long value greater than or equal to 0,
	// and less than around `9.2e+18`.
	//
	// NOTE: Versioning is completely real time, and is not affected by the near
	// real time aspects of search operations.
	// If no version is provided, the operation runs without any version checks.
	//
	// When using the external version type, the system checks to see if the version
	// number passed to the index request is greater than the version of the
	// currently stored document.
	// If true, the document will be indexed and the new version number used.
	// If the value provided is less than or equal to the stored document's version
	// number, a version conflict will occur and the index operation will fail. For
	// example:
	//
	// ```
	// PUT my-index-000001/_doc/1?version=2&version_type=external
	// {
	//   "user": {
	//     "id": "elkbee"
	//   }
	// }
	//
	// In this example, the operation will succeed since the supplied version of 2
	// is higher than the current document version of 1.
	// If the document was already updated and its version was set to 2 or higher,
	// the indexing command will fail and result in a conflict (409 HTTP status
	// code).
	//
	// A nice side effect is that there is no need to maintain strict ordering of
	// async indexing operations run as a result of changes to a source database, as
	// long as version numbers from the source database are used.
	// Even the simple case of updating the Elasticsearch index using data from a
	// database is simplified if external versioning is used, as only the latest
	// version will be used if the index operations arrive out of order.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Index core_index.NewIndex
	// Get cluster info.
	// Get basic build, version, and cluster information.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rest-api-root.html
	Info core_info.NewInfo
	// Run a knn search.
	//
	// NOTE: The kNN search API has been replaced by the `knn` option in the search
	// API.
	//
	// Perform a k-nearest neighbor (kNN) search on a dense_vector field and return
	// the matching documents.
	// Given a query vector, the API finds the k closest vectors and returns those
	// documents as search hits.
	//
	// Elasticsearch uses the HNSW algorithm to support efficient kNN search.
	// Like most kNN algorithms, HNSW is an approximate method that sacrifices
	// result accuracy for improved search speed.
	// This means the results returned are not always the true k closest neighbors.
	//
	// The kNN search API supports restricting the search using a filter.
	// The search will return the top k documents that also match the filter query.
	//
	// A kNN search response has the exact same structure as a search API response.
	// However, certain sections have a meaning specific to kNN search:
	//
	// * The document `_score` is determined by the similarity between the query and
	// document vector.
	// * The `hits.total` object contains the total number of nearest neighbor
	// candidates considered, which is `num_candidates * num_shards`. The
	// `hits.total.relation` will always be `eq`, indicating an exact value.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/knn-search-api.html
	KnnSearch core_knn_search.NewKnnSearch
	// Get multiple documents.
	//
	// Get multiple JSON documents by ID from one or more indices.
	// If you specify an index in the request URI, you only need to specify the
	// document IDs in the request body.
	// To ensure fast responses, this multi get (mget) API responds with partial
	// results if one or more shards fail.
	//
	// **Filter source fields**
	//
	// By default, the `_source` field is returned for every document (if stored).
	// Use the `_source` and `_source_include` or `source_exclude` attributes to
	// filter what fields are returned for a particular document.
	// You can include the `_source`, `_source_includes`, and `_source_excludes`
	// query parameters in the request URI to specify the defaults to use when there
	// are no per-document instructions.
	//
	// **Get stored fields**
	//
	// Use the `stored_fields` attribute to specify the set of stored fields you
	// want to retrieve.
	// Any requested fields that are not stored are ignored.
	// You can include the `stored_fields` query parameter in the request URI to
	// specify the defaults to use when there are no per-document instructions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
	Mget core_mget.NewMget
	// Run multiple searches.
	//
	// The format of the request is similar to the bulk API format and makes use of
	// the newline delimited JSON (NDJSON) format.
	// The structure is as follows:
	//
	// ```
	// header\n
	// body\n
	// header\n
	// body\n
	// ```
	//
	// This structure is specifically optimized to reduce parsing if a specific
	// search ends up redirected to another node.
	//
	// IMPORTANT: The final line of data must end with a newline character `\n`.
	// Each newline character may be preceded by a carriage return `\r`.
	// When sending requests to this endpoint the `Content-Type` header should be
	// set to `application/x-ndjson`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	Msearch core_msearch.NewMsearch
	// Run multiple templated searches.
	//
	// Run multiple templated searches with a single request.
	// If you are providing a text file or text input to `curl`, use the
	// `--data-binary` flag instead of `-d` to preserve newlines.
	// For example:
	//
	// ```
	// $ cat requests
	// { "index": "my-index" }
	// { "id": "my-search-template", "params": { "query_string": "hello world",
	// "from": 0, "size": 10 }}
	// { "index": "my-other-index" }
	// { "id": "my-other-search-template", "params": { "query_type": "match_all" }}
	//
	// $ curl -H "Content-Type: application/x-ndjson" -XGET
	// localhost:9200/_msearch/template --data-binary "@requests"; echo
	// ```
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-search-template.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Get multiple term vectors.
	//
	// Get multiple term vectors with a single request.
	// You can specify existing documents by index and ID or provide artificial
	// documents in the body of the request.
	// You can specify the index in the request body or request URI.
	// The response contains a `docs` array with all the fetched termvectors.
	// Each element has the structure provided by the termvectors API.
	//
	// **Artificial documents**
	//
	// You can also use `mtermvectors` to generate term vectors for artificial
	// documents provided in the body of the request.
	// The mapping used is determined by the specified `_index`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// Open a point in time.
	//
	// A search request by default runs against the most recent visible data of the
	// target indices,
	// which is called point in time. Elasticsearch pit (point in time) is a
	// lightweight view into the
	// state of the data as it existed when initiated. In some cases, it’s preferred
	// to perform multiple
	// search requests using the same point in time. For example, if refreshes
	// happen between
	// `search_after` requests, then the results of those requests might not be
	// consistent as changes happening
	// between searches are only visible to the more recent point in time.
	//
	// A point in time must be opened explicitly before being used in search
	// requests.
	//
	// A subsequent search request with the `pit` parameter must not specify
	// `index`, `routing`, or `preference` values as these parameters are copied
	// from the point in time.
	//
	// Just like regular searches, you can use `from` and `size` to page through
	// point in time search results, up to the first 10,000 hits.
	// If you want to retrieve more hits, use PIT with `search_after`.
	//
	// IMPORTANT: The open point in time request and each subsequent search request
	// can return different identifiers; always use the most recently received ID
	// for the next search request.
	//
	// When a PIT that contains shard failures is used in a search request, the
	// missing are always reported in the search response as a
	// `NoShardAvailableActionException` exception.
	// To get rid of these exceptions, a new PIT needs to be created so that shards
	// missing from the previous PIT can be handled, assuming they become available
	// in the meantime.
	//
	// **Keeping point in time alive**
	//
	// The `keep_alive` parameter, which is passed to a open point in time request
	// and search request, extends the time to live of the corresponding point in
	// time.
	// The value does not need to be long enough to process all data — it just needs
	// to be long enough for the next request.
	//
	// Normally, the background merge process optimizes the index by merging
	// together smaller segments to create new, bigger segments.
	// Once the smaller segments are no longer needed they are deleted.
	// However, open point-in-times prevent the old segments from being deleted
	// since they are still in use.
	//
	// TIP: Keeping older segments alive means that more disk space and file handles
	// are needed.
	// Ensure that you have configured your nodes to have ample free file handles.
	//
	// Additionally, if a segment contains deleted or updated documents then the
	// point in time must keep track of whether each document in the segment was
	// live at the time of the initial search request.
	// Ensure that your nodes have sufficient heap space if you have many open
	// point-in-times on an index that is subject to ongoing deletes or updates.
	// Note that a point-in-time doesn't prevent its associated indices from being
	// deleted.
	// You can check how many point-in-times (that is, search contexts) are open
	// with the nodes stats API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Ping the cluster.
	// Get information about whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster.html
	Ping core_ping.NewPing
	// Create or update a script or search template.
	// Creates or updates a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-stored-script-api.html
	PutScript core_put_script.NewPutScript
	// Evaluate ranked search results.
	//
	// Evaluate the quality of ranked search results over a set of typical search
	// queries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Reindex documents.
	//
	// Copy documents from a source to a destination.
	// You can copy all documents to the destination index or reindex a subset of
	// the documents.
	// The source can be any existing index, alias, or data stream.
	// The destination must differ from the source.
	// For example, you cannot reindex a data stream into itself.
	//
	// IMPORTANT: Reindex requires `_source` to be enabled for all documents in the
	// source.
	// The destination should be configured as wanted before calling the reindex
	// API.
	// Reindex does not copy the settings from the source or its associated
	// template.
	// Mappings, shard counts, and replicas, for example, must be configured ahead
	// of time.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following security privileges:
	//
	// * The `read` index privilege for the source data stream, index, or alias.
	// * The `write` index privilege for the destination data stream, index, or
	// index alias.
	// * To automatically create a data stream or index with a reindex API request,
	// you must have the `auto_configure`, `create_index`, or `manage` index
	// privilege for the destination data stream, index, or alias.
	// * If reindexing from a remote cluster, the `source.remote.user` must have the
	// `monitor` cluster privilege and the `read` index privilege for the source
	// data stream, index, or alias.
	//
	// If reindexing from a remote cluster, you must explicitly allow the remote
	// host in the `reindex.remote.whitelist` setting.
	// Automatic data stream creation requires a matching index template with data
	// stream enabled.
	//
	// The `dest` element can be configured like the index API to control optimistic
	// concurrency control.
	// Omitting `version_type` or setting it to `internal` causes Elasticsearch to
	// blindly dump documents into the destination, overwriting any that happen to
	// have the same ID.
	//
	// Setting `version_type` to `external` causes Elasticsearch to preserve the
	// `version` from the source, create any documents that are missing, and update
	// any documents that have an older version in the destination than they do in
	// the source.
	//
	// Setting `op_type` to `create` causes the reindex API to create only missing
	// documents in the destination.
	// All existing documents will cause a version conflict.
	//
	// IMPORTANT: Because data streams are append-only, any reindex request to a
	// destination data stream must have an `op_type` of `create`.
	// A reindex can only add new documents to a destination data stream.
	// It cannot update existing documents in a destination data stream.
	//
	// By default, version conflicts abort the reindex process.
	// To continue reindexing if there are conflicts, set the `conflicts` request
	// body property to `proceed`.
	// In this case, the response includes a count of the version conflicts that
	// were encountered.
	// Note that the handling of other error types is unaffected by the `conflicts`
	// property.
	// Additionally, if you opt to count version conflicts, the operation could
	// attempt to reindex more documents from the source than `max_docs` until it
	// has successfully indexed `max_docs` documents into the target or it has gone
	// through every document in the source query.
	//
	// NOTE: The reindex API makes no effort to handle ID collisions.
	// The last document written will "win" but the order isn't usually predictable
	// so it is not a good idea to rely on this behavior.
	// Instead, make sure that IDs are unique by using a script.
	//
	// **Running reindex asynchronously**
	//
	// If the request contains `wait_for_completion=false`, Elasticsearch performs
	// some preflight checks, launches the request, and returns a task you can use
	// to cancel or get the status of the task.
	// Elasticsearch creates a record of this task as a document at
	// `_tasks/<task_id>`.
	//
	// **Reindex from multiple sources**
	//
	// If you have many sources to reindex it is generally better to reindex them
	// one at a time rather than using a glob pattern to pick up multiple sources.
	// That way you can resume the process if there are any errors by removing the
	// partially completed source and starting over.
	// It also makes parallelizing the process fairly simple: split the list of
	// sources to reindex and run each list in parallel.
	//
	// For example, you can use a bash script like this:
	//
	// ```
	// for index in i1 i2 i3 i4 i5; do
	//   curl -HContent-Type:application/json -XPOST localhost:9200/_reindex?pretty
	// -d'{
	//     "source": {
	//       "index": "'$index'"
	//     },
	//     "dest": {
	//       "index": "'$index'-reindexed"
	//     }
	//   }'
	// done
	// ```
	//
	// **Throttling**
	//
	// Set `requests_per_second` to any positive decimal number (`1.4`, `6`, `1000`,
	// for example) to throttle the rate at which reindex issues batches of index
	// operations.
	// Requests are throttled by padding each batch with a wait time.
	// To turn off throttling, set `requests_per_second` to `-1`.
	//
	// The throttling is done by waiting between batches so that the scroll that
	// reindex uses internally can be given a timeout that takes into account the
	// padding.
	// The padding time is the difference between the batch size divided by the
	// `requests_per_second` and the time spent writing.
	// By default the batch size is `1000`, so if `requests_per_second` is set to
	// `500`:
	//
	// ```
	// target_time = 1000 / 500 per second = 2 seconds
	// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
	// ```
	//
	// Since the batch is issued as a single bulk request, large batch sizes cause
	// Elasticsearch to create many requests and then wait for a while before
	// starting the next set.
	// This is "bursty" instead of "smooth".
	//
	// **Slicing**
	//
	// Reindex supports sliced scroll to parallelize the reindexing process.
	// This parallelization can improve efficiency and provide a convenient way to
	// break the request down into smaller parts.
	//
	// NOTE: Reindexing from remote clusters does not support manual or automatic
	// slicing.
	//
	// You can slice a reindex request manually by providing a slice ID and total
	// number of slices to each request.
	// You can also let reindex automatically parallelize by using sliced scroll to
	// slice on `_id`.
	// The `slices` parameter specifies the number of slices to use.
	//
	// Adding `slices` to the reindex request just automates the manual process,
	// creating sub-requests which means it has some quirks:
	//
	// * You can see these requests in the tasks API. These sub-requests are "child"
	// tasks of the task for the request with slices.
	// * Fetching the status of the task for the request with `slices` only contains
	// the status of completed slices.
	// * These sub-requests are individually addressable for things like
	// cancellation and rethrottling.
	// * Rethrottling the request with `slices` will rethrottle the unfinished
	// sub-request proportionally.
	// * Canceling the request with `slices` will cancel each sub-request.
	// * Due to the nature of `slices`, each sub-request won't get a perfectly even
	// portion of the documents. All documents will be addressed, but some slices
	// may be larger than others. Expect larger slices to have a more even
	// distribution.
	// * Parameters like `requests_per_second` and `max_docs` on a request with
	// `slices` are distributed proportionally to each sub-request. Combine that
	// with the previous point about distribution being uneven and you should
	// conclude that using `max_docs` with `slices` might not result in exactly
	// `max_docs` documents being reindexed.
	// * Each sub-request gets a slightly different snapshot of the source, though
	// these are all taken at approximately the same time.
	//
	// If slicing automatically, setting `slices` to `auto` will choose a reasonable
	// number for most indices.
	// If slicing manually or otherwise tuning automatic slicing, use the following
	// guidelines.
	//
	// Query performance is most efficient when the number of slices is equal to the
	// number of shards in the index.
	// If that number is large (for example, `500`), choose a lower number as too
	// many slices will hurt performance.
	// Setting slices higher than the number of shards generally does not improve
	// efficiency and adds overhead.
	//
	// Indexing performance scales linearly across available resources with the
	// number of slices.
	//
	// Whether query or indexing performance dominates the runtime depends on the
	// documents being reindexed and cluster resources.
	//
	// **Modify documents during reindexing**
	//
	// Like `_update_by_query`, reindex operations support a script that modifies
	// the document.
	// Unlike `_update_by_query`, the script is allowed to modify the document's
	// metadata.
	//
	// Just as in `_update_by_query`, you can set `ctx.op` to change the operation
	// that is run on the destination.
	// For example, set `ctx.op` to `noop` if your script decides that the document
	// doesn’t have to be indexed in the destination. This "no operation" will be
	// reported in the `noop` counter in the response body.
	// Set `ctx.op` to `delete` if your script decides that the document must be
	// deleted from the destination.
	// The deletion will be reported in the `deleted` counter in the response body.
	// Setting `ctx.op` to anything else will return an error, as will setting any
	// other field in `ctx`.
	//
	// Think of the possibilities! Just be careful; you are able to change:
	//
	// * `_id`
	// * `_index`
	// * `_version`
	// * `_routing`
	//
	// Setting `_version` to `null` or clearing it from the `ctx` map is just like
	// not sending the version in an indexing request.
	// It will cause the document to be overwritten in the destination regardless of
	// the version on the target or the version type you use in the reindex API.
	//
	// **Reindex from remote**
	//
	// Reindex supports reindexing from a remote Elasticsearch cluster.
	// The `host` parameter must contain a scheme, host, port, and optional path.
	// The `username` and `password` parameters are optional and when they are
	// present the reindex operation will connect to the remote Elasticsearch node
	// using basic authentication.
	// Be sure to use HTTPS when using basic authentication or the password will be
	// sent in plain text.
	// There are a range of settings available to configure the behavior of the
	// HTTPS connection.
	//
	// When using Elastic Cloud, it is also possible to authenticate against the
	// remote cluster through the use of a valid API key.
	// Remote hosts must be explicitly allowed with the `reindex.remote.whitelist`
	// setting.
	// It can be set to a comma delimited list of allowed remote host and port
	// combinations.
	// Scheme is ignored; only the host and port are used.
	// For example:
	//
	// ```
	// reindex.remote.whitelist: [otherhost:9200, another:9200, 127.0.10.*:9200,
	// localhost:*"]
	// ```
	//
	// The list of allowed hosts must be configured on any nodes that will
	// coordinate the reindex.
	// This feature should work with remote clusters of any version of
	// Elasticsearch.
	// This should enable you to upgrade from any version of Elasticsearch to the
	// current version by reindexing from a cluster of the old version.
	//
	// WARNING: Elasticsearch does not support forward compatibility across major
	// versions.
	// For example, you cannot reindex from a 7.x cluster into a 6.x cluster.
	//
	// To enable queries sent to older versions of Elasticsearch, the `query`
	// parameter is sent directly to the remote host without validation or
	// modification.
	//
	// NOTE: Reindexing from remote clusters does not support manual or automatic
	// slicing.
	//
	// Reindexing from a remote server uses an on-heap buffer that defaults to a
	// maximum size of 100mb.
	// If the remote index includes very large documents you'll need to use a
	// smaller batch size.
	// It is also possible to set the socket read timeout on the remote connection
	// with the `socket_timeout` field and the connection timeout with the
	// `connect_timeout` field.
	// Both default to 30 seconds.
	//
	// **Configuring SSL parameters**
	//
	// Reindex from remote supports configurable SSL settings.
	// These must be specified in the `elasticsearch.yml` file, with the exception
	// of the secure settings, which you add in the Elasticsearch keystore.
	// It is not possible to configure SSL in the body of the reindex request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Throttle a reindex operation.
	//
	// Change the number of requests per second for a particular reindex operation.
	// For example:
	//
	// ```
	// POST _reindex/r1A2WoRbTwKZ516z6NEs5A:36619/_rethrottle?requests_per_second=-1
	// ```
	//
	// Rethrottling that speeds up the query takes effect immediately.
	// Rethrottling that slows down the query will take effect after completing the
	// current batch.
	// This behavior prevents scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle
	// Render a search template.
	//
	// Render a search template as a search request body.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Run a script.
	//
	// Runs a script and returns a result.
	// Use this API to build and test scripts, such as when defining a script for a
	// runtime field.
	// This API requires very few dependencies and is especially useful if you don't
	// have permissions to write documents on a cluster.
	//
	// The API uses several _contexts_, which control how scripts are run, what
	// variables are available at runtime, and what the return type is.
	//
	// Each context requires a script, but additional parameters depend on the
	// context you're using for that script.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Run a scrolling search.
	//
	// IMPORTANT: The scroll API is no longer recommend for deep pagination. If you
	// need to preserve the index state while paging through more than 10,000 hits,
	// use the `search_after` parameter with a point in time (PIT).
	//
	// The scroll API gets large sets of results from a single scrolling search
	// request.
	// To get the necessary scroll ID, submit a search API request that includes an
	// argument for the `scroll` query parameter.
	// The `scroll` parameter indicates how long Elasticsearch should retain the
	// search context for the request.
	// The search response returns a scroll ID in the `_scroll_id` response body
	// parameter.
	// You can then use the scroll ID with the scroll API to retrieve the next batch
	// of results for the request.
	// If the Elasticsearch security features are enabled, the access to the results
	// of a specific scroll ID is restricted to the user or API key that submitted
	// the search.
	//
	// You can also use the scroll API to specify a new scroll parameter that
	// extends or shortens the retention period for the search context.
	//
	// IMPORTANT: Results from a scrolling search reflect the state of the index at
	// the time of the initial search request. Subsequent indexing or document
	// changes only affect later search and scroll requests.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/scroll-api.html
	Scroll core_scroll.NewScroll
	// Run a search.
	//
	// Get search hits that match the query defined in the request.
	// You can provide search queries using the `q` query string parameter or the
	// request body.
	// If both are specified, only the query parameter is used.
	//
	// If the Elasticsearch security features are enabled, you must have the read
	// index privilege for the target data stream, index, or alias. For
	// cross-cluster search, refer to the documentation about configuring CCS
	// privileges.
	// To search a point in time (PIT) for an alias, you must have the `read` index
	// privilege for the alias's data streams or indices.
	//
	// **Search slicing**
	//
	// When paging through a large number of documents, it can be helpful to split
	// the search into multiple slices to consume them independently with the
	// `slice` and `pit` properties.
	// By default the splitting is done first on the shards, then locally on each
	// shard.
	// The local splitting partitions the shard into contiguous ranges based on
	// Lucene document IDs.
	//
	// For instance if the number of shards is equal to 2 and you request 4 slices,
	// the slices 0 and 2 are assigned to the first shard and the slices 1 and 3 are
	// assigned to the second shard.
	//
	// IMPORTANT: The same point-in-time ID should be used for all slices.
	// If different PIT IDs are used, slices can overlap and miss documents.
	// This situation can occur because the splitting criterion is based on Lucene
	// document IDs, which are not stable across changes to the index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Search a vector tile.
	//
	// Search a vector tile for geospatial values.
	// Before using this API, you should be familiar with the Mapbox vector tile
	// specification.
	// The API returns results as a binary mapbox vector tile.
	//
	// Internally, Elasticsearch translates a vector tile search API request into a
	// search containing:
	//
	// * A `geo_bounding_box` query on the `<field>`. The query uses the
	// `<zoom>/<x>/<y>` tile as a bounding box.
	// * A `geotile_grid` or `geohex_grid` aggregation on the `<field>`. The
	// `grid_agg` parameter determines the aggregation type. The aggregation uses
	// the `<zoom>/<x>/<y>` tile as a bounding box.
	// * Optionally, a `geo_bounds` aggregation on the `<field>`. The search only
	// includes this aggregation if the `exact_bounds` parameter is `true`.
	// * If the optional parameter `with_labels` is `true`, the internal search will
	// include a dynamic runtime field that calls the `getLabelPosition` function of
	// the geometry doc value. This enables the generation of new point features
	// containing suggested geometry labels, so that, for example, multi-polygons
	// will have only one label.
	//
	// For example, Elasticsearch may translate a vector tile search API request
	// with a `grid_agg` argument of `geotile` and an `exact_bounds` argument of
	// `true` into the following search
	//
	// ```
	// GET my-index/_search
	// {
	//   "size": 10000,
	//   "query": {
	//     "geo_bounding_box": {
	//       "my-geo-field": {
	//         "top_left": {
	//           "lat": -40.979898069620134,
	//           "lon": -45
	//         },
	//         "bottom_right": {
	//           "lat": -66.51326044311186,
	//           "lon": 0
	//         }
	//       }
	//     }
	//   },
	//   "aggregations": {
	//     "grid": {
	//       "geotile_grid": {
	//         "field": "my-geo-field",
	//         "precision": 11,
	//         "size": 65536,
	//         "bounds": {
	//           "top_left": {
	//             "lat": -40.979898069620134,
	//             "lon": -45
	//           },
	//           "bottom_right": {
	//             "lat": -66.51326044311186,
	//             "lon": 0
	//           }
	//         }
	//       }
	//     },
	//     "bounds": {
	//       "geo_bounds": {
	//         "field": "my-geo-field",
	//         "wrap_longitude": false
	//       }
	//     }
	//   }
	// }
	// ```
	//
	// The API returns results as a binary Mapbox vector tile.
	// Mapbox vector tiles are encoded as Google Protobufs (PBF). By default, the
	// tile contains three layers:
	//
	// * A `hits` layer containing a feature for each `<field>` value matching the
	// `geo_bounding_box` query.
	// * An `aggs` layer containing a feature for each cell of the `geotile_grid` or
	// `geohex_grid`. The layer only contains features for cells with matching data.
	// * A meta layer containing:
	//   * A feature containing a bounding box. By default, this is the bounding box
	// of the tile.
	//   * Value ranges for any sub-aggregations on the `geotile_grid` or
	// `geohex_grid`.
	//   * Metadata for the search.
	//
	// The API only returns features that can display at its zoom level.
	// For example, if a polygon feature has no area at its zoom level, the API
	// omits it.
	// The API returns errors as UTF-8 encoded JSON.
	//
	// IMPORTANT: You can specify several options for this API as either a query
	// parameter or request body parameter.
	// If you specify both parameters, the query parameter takes precedence.
	//
	// **Grid precision for geotile**
	//
	// For a `grid_agg` of `geotile`, you can use cells in the `aggs` layer as tiles
	// for lower zoom levels.
	// `grid_precision` represents the additional zoom levels available through
	// these cells. The final precision is computed by as follows: `<zoom> +
	// grid_precision`.
	// For example, if `<zoom>` is 7 and `grid_precision` is 8, then the
	// `geotile_grid` aggregation will use a precision of 15.
	// The maximum final precision is 29.
	// The `grid_precision` also determines the number of cells for the grid as
	// follows: `(2^grid_precision) x (2^grid_precision)`.
	// For example, a value of 8 divides the tile into a grid of 256 x 256 cells.
	// The `aggs` layer only contains features for cells with matching data.
	//
	// **Grid precision for geohex**
	//
	// For a `grid_agg` of `geohex`, Elasticsearch uses `<zoom>` and
	// `grid_precision` to calculate a final precision as follows: `<zoom> +
	// grid_precision`.
	//
	// This precision determines the H3 resolution of the hexagonal cells produced
	// by the `geohex` aggregation.
	// The following table maps the H3 resolution for each precision.
	// For example, if `<zoom>` is 3 and `grid_precision` is 3, the precision is 6.
	// At a precision of 6, hexagonal cells have an H3 resolution of 2.
	// If `<zoom>` is 3 and `grid_precision` is 4, the precision is 7.
	// At a precision of 7, hexagonal cells have an H3 resolution of 3.
	//
	// | Precision | Unique tile bins | H3 resolution | Unique hex bins |	Ratio |
	// | --------- | ---------------- | ------------- | ----------------| ----- |
	// | 1  | 4                  | 0  | 122             | 30.5           |
	// | 2  | 16                 | 0  | 122             | 7.625          |
	// | 3  | 64                 | 1  | 842             | 13.15625       |
	// | 4  | 256                | 1  | 842             | 3.2890625      |
	// | 5  | 1024               | 2  | 5882            | 5.744140625    |
	// | 6  | 4096               | 2  | 5882            | 1.436035156    |
	// | 7  | 16384              | 3  | 41162           | 2.512329102    |
	// | 8  | 65536              | 3  | 41162           | 0.6280822754   |
	// | 9  | 262144             | 4  | 288122          | 1.099098206    |
	// | 10 | 1048576            | 4  | 288122          | 0.2747745514   |
	// | 11 | 4194304            | 5  | 2016842         | 0.4808526039   |
	// | 12 | 16777216           | 6  | 14117882        | 0.8414913416   |
	// | 13 | 67108864           | 6  | 14117882        | 0.2103728354   |
	// | 14 | 268435456          | 7  | 98825162        | 0.3681524172   |
	// | 15 | 1073741824         | 8  | 691776122       | 0.644266719    |
	// | 16 | 4294967296         | 8  | 691776122       | 0.1610666797   |
	// | 17 | 17179869184        | 9  | 4842432842      | 0.2818666889   |
	// | 18 | 68719476736        | 10 | 33897029882     | 0.4932667053   |
	// | 19 | 274877906944       | 11 | 237279209162    | 0.8632167343   |
	// | 20 | 1099511627776      | 11 | 237279209162    | 0.2158041836   |
	// | 21 | 4398046511104      | 12 | 1660954464122   | 0.3776573213   |
	// | 22 | 17592186044416     | 13 | 11626681248842  | 0.6609003122   |
	// | 23 | 70368744177664     | 13 | 11626681248842  | 0.165225078    |
	// | 24 | 281474976710656    | 14 | 81386768741882  | 0.2891438866   |
	// | 25 | 1125899906842620   | 15 | 569707381193162 | 0.5060018015   |
	// | 26 | 4503599627370500   | 15 | 569707381193162 | 0.1265004504   |
	// | 27 | 18014398509482000  | 15 | 569707381193162 | 0.03162511259  |
	// | 28 | 72057594037927900  | 15 | 569707381193162 | 0.007906278149 |
	// | 29 | 288230376151712000 | 15 | 569707381193162 | 0.001976569537 |
	//
	// Hexagonal cells don't align perfectly on a vector tile.
	// Some cells may intersect more than one vector tile.
	// To compute the H3 resolution for each precision, Elasticsearch compares the
	// average density of hexagonal bins at each resolution with the average density
	// of tile bins at each zoom level.
	// Elasticsearch uses the H3 resolution that is closest to the corresponding
	// geotile density.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Get the search shards.
	//
	// Get the indices and shards that a search request would be run against.
	// This information can be useful for working out issues or planning
	// optimizations with routing and shard preferences.
	// When filtered aliases are used, the filter is returned as part of the
	// `indices` section.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// `view_index_metadata` or `manage` index privilege for the target data stream,
	// index, or alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-shards.html
	SearchShards core_search_shards.NewSearchShards
	// Run a search with a search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template-api.html
	SearchTemplate core_search_template.NewSearchTemplate
	// Get terms in an index.
	//
	// Discover terms that match a partial string in an index.
	// This API is designed for low-latency look-ups used in auto-complete
	// scenarios.
	//
	// > info
	// > The terms enum API may return terms from deleted documents. Deleted
	// documents are initially only marked as deleted. It is not until their
	// segments are merged that documents are actually deleted. Until that happens,
	// the terms enum API will return terms from these documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
	TermsEnum core_terms_enum.NewTermsEnum
	// Get term vector information.
	//
	// Get information and statistics about terms in the fields of a particular
	// document.
	//
	// You can retrieve term vectors for documents stored in the index or for
	// artificial documents passed in the body of the request.
	// You can specify the fields you are interested in through the `fields`
	// parameter or by adding the fields to the request body.
	// For example:
	//
	// ```
	// GET /my-index-000001/_termvectors/1?fields=message
	// ```
	//
	// Fields can be specified using wildcards, similar to the multi match query.
	//
	// Term vectors are real-time by default, not near real-time.
	// This can be changed by setting `realtime` parameter to `false`.
	//
	// You can request three types of values: _term information_, _term statistics_,
	// and _field statistics_.
	// By default, all term information and field statistics are returned for all
	// fields but term statistics are excluded.
	//
	// **Term information**
	//
	// * term frequency in the field (always returned)
	// * term positions (`positions: true`)
	// * start and end offsets (`offsets: true`)
	// * term payloads (`payloads: true`), as base64 encoded bytes
	//
	// If the requested information wasn't stored in the index, it will be computed
	// on the fly if possible.
	// Additionally, term vectors could be computed for documents not even existing
	// in the index, but instead provided by the user.
	//
	// > warn
	// > Start and end offsets assume UTF-16 encoding is being used. If you want to
	// use these offsets in order to get the original text that produced this token,
	// you should make sure that the string you are taking a sub-string of is also
	// encoded using UTF-16.
	//
	// **Behaviour**
	//
	// The term and field statistics are not accurate.
	// Deleted documents are not taken into account.
	// The information is only retrieved for the shard the requested document
	// resides in.
	// The term and field statistics are therefore only useful as relative measures
	// whereas the absolute numbers have no meaning in this context.
	// By default, when requesting term vectors of artificial documents, a shard to
	// get the statistics from is randomly selected.
	// Use `routing` only to hit a particular shard.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
	Termvectors core_termvectors.NewTermvectors
	// Update a document.
	//
	// Update a document by running a script or passing a partial document.
	//
	// If the Elasticsearch security features are enabled, you must have the `index`
	// or `write` index privilege for the target index or index alias.
	//
	// The script can update, delete, or skip modifying the document.
	// The API also supports passing a partial document, which is merged into the
	// existing document.
	// To fully replace an existing document, use the index API.
	// This operation:
	//
	// * Gets the document (collocated with the shard) from the index.
	// * Runs the specified script.
	// * Indexes the result.
	//
	// The document must still be reindexed, but using this API removes some network
	// roundtrips and reduces chances of version conflicts between the GET and the
	// index operation.
	//
	// The `_source` field must be enabled to use this API.
	// In addition to `_source`, you can access the following variables through the
	// `ctx` map: `_index`, `_type`, `_id`, `_version`, `_routing`, and `_now` (the
	// current timestamp).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
	Update core_update.NewUpdate
	// Update documents.
	// Updates documents that match the specified query.
	// If no query is specified, performs an update on every document in the data
	// stream or index without modifying the source, which is useful for picking up
	// mapping changes.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or alias:
	//
	// * `read`
	// * `index` or `write`
	//
	// You can specify the query criteria in the request URI or the request body
	// using the same syntax as the search API.
	//
	// When you submit an update by query request, Elasticsearch gets a snapshot of
	// the data stream or index when it begins processing the request and updates
	// matching documents using internal versioning.
	// When the versions match, the document is updated and the version number is
	// incremented.
	// If a document changes between the time that the snapshot is taken and the
	// update operation is processed, it results in a version conflict and the
	// operation fails.
	// You can opt to count version conflicts instead of halting and returning by
	// setting `conflicts` to `proceed`.
	// Note that if you opt to count version conflicts, the operation could attempt
	// to update more documents from the source than `max_docs` until it has
	// successfully updated `max_docs` documents or it has gone through every
	// document in the source query.
	//
	// NOTE: Documents with a version equal to 0 cannot be updated using update by
	// query because internal versioning does not support 0 as a valid version
	// number.
	//
	// While processing an update by query request, Elasticsearch performs multiple
	// search requests sequentially to find all of the matching documents.
	// A bulk update request is performed for each batch of matching documents.
	// Any query or update failures cause the update by query request to fail and
	// the failures are shown in the response.
	// Any update requests that completed successfully still stick, they are not
	// rolled back.
	//
	// **Throttling update requests**
	//
	// To control the rate at which update by query issues batches of update
	// operations, you can set `requests_per_second` to any positive decimal number.
	// This pads each batch with a wait time to throttle the rate.
	// Set `requests_per_second` to `-1` to turn off throttling.
	//
	// Throttling uses a wait time between batches so that the internal scroll
	// requests can be given a timeout that takes the request padding into account.
	// The padding time is the difference between the batch size divided by the
	// `requests_per_second` and the time spent writing.
	// By default the batch size is 1000, so if `requests_per_second` is set to
	// `500`:
	//
	// ```
	// target_time = 1000 / 500 per second = 2 seconds
	// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
	// ```
	//
	// Since the batch is issued as a single _bulk request, large batch sizes cause
	// Elasticsearch to create many requests and wait before starting the next set.
	// This is "bursty" instead of "smooth".
	//
	// **Slicing**
	//
	// Update by query supports sliced scroll to parallelize the update process.
	// This can improve efficiency and provide a convenient way to break the request
	// down into smaller parts.
	//
	// Setting `slices` to `auto` chooses a reasonable number for most data streams
	// and indices.
	// This setting will use one slice per shard, up to a certain limit.
	// If there are multiple source data streams or indices, it will choose the
	// number of slices based on the index or backing index with the smallest number
	// of shards.
	//
	// Adding `slices` to `_update_by_query` just automates the manual process of
	// creating sub-requests, which means it has some quirks:
	//
	// * You can see these requests in the tasks APIs. These sub-requests are
	// "child" tasks of the task for the request with slices.
	// * Fetching the status of the task for the request with `slices` only contains
	// the status of completed slices.
	// * These sub-requests are individually addressable for things like
	// cancellation and rethrottling.
	// * Rethrottling the request with `slices` will rethrottle the unfinished
	// sub-request proportionally.
	// * Canceling the request with slices will cancel each sub-request.
	// * Due to the nature of slices each sub-request won't get a perfectly even
	// portion of the documents. All documents will be addressed, but some slices
	// may be larger than others. Expect larger slices to have a more even
	// distribution.
	// * Parameters like `requests_per_second` and `max_docs` on a request with
	// slices are distributed proportionally to each sub-request. Combine that with
	// the point above about distribution being uneven and you should conclude that
	// using `max_docs` with `slices` might not result in exactly `max_docs`
	// documents being updated.
	// * Each sub-request gets a slightly different snapshot of the source data
	// stream or index though these are all taken at approximately the same time.
	//
	// If you're slicing manually or otherwise tuning automatic slicing, keep in
	// mind that:
	//
	// * Query performance is most efficient when the number of slices is equal to
	// the number of shards in the index or backing index. If that number is large
	// (for example, 500), choose a lower number as too many slices hurts
	// performance. Setting slices higher than the number of shards generally does
	// not improve efficiency and adds overhead.
	// * Update performance scales linearly across available resources with the
	// number of slices.
	//
	// Whether query or update performance dominates the runtime depends on the
	// documents being reindexed and cluster resources.
	//
	// **Update the document source**
	//
	// Update by query supports scripts to update the document source.
	// As with the update API, you can set `ctx.op` to change the operation that is
	// performed.
	//
	// Set `ctx.op = "noop"` if your script decides that it doesn't have to make any
	// changes.
	// The update by query operation skips updating the document and increments the
	// `noop` counter.
	//
	// Set `ctx.op = "delete"` if your script decides that the document should be
	// deleted.
	// The update by query operation deletes the document and increments the
	// `deleted` counter.
	//
	// Update by query supports only `index`, `noop`, and `delete`.
	// Setting `ctx.op` to anything else is an error.
	// Setting any other field in `ctx` is an error.
	// This API enables you to only modify the source of matching documents; you
	// cannot move them.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQuery core_update_by_query.NewUpdateByQuery
	// Throttle an update by query operation.
	//
	// Change the number of requests per second for a particular update by query
	// operation.
	// Rethrottling that speeds up the query takes effect immediately but
	// rethrotting that slows down the query takes effect after completing the
	// current batch to prevent scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html#docs-update-by-query-rethrottle
	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

type DanglingIndices struct {
	// Delete a dangling index.
	// If Elasticsearch encounters index data that is absent from the current
	// cluster state, those indices are considered to be dangling.
	// For example, this can happen if you delete more than
	// `cluster.indices.tombstones.size` indices while an Elasticsearch node is
	// offline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/dangling-index-delete.html
	DeleteDanglingIndex dangling_indices_delete_dangling_index.NewDeleteDanglingIndex
	// Import a dangling index.
	//
	// If Elasticsearch encounters index data that is absent from the current
	// cluster state, those indices are considered to be dangling.
	// For example, this can happen if you delete more than
	// `cluster.indices.tombstones.size` indices while an Elasticsearch node is
	// offline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/dangling-index-import.html
	ImportDanglingIndex dangling_indices_import_dangling_index.NewImportDanglingIndex
	// Get the dangling indices.
	//
	// If Elasticsearch encounters index data that is absent from the current
	// cluster state, those indices are considered to be dangling.
	// For example, this can happen if you delete more than
	// `cluster.indices.tombstones.size` indices while an Elasticsearch node is
	// offline.
	//
	// Use this API to list dangling indices, which you can then import or delete.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/dangling-indices-list.html
	ListDanglingIndices dangling_indices_list_dangling_indices.NewListDanglingIndices
}

type Enrich struct {
	// Delete an enrich policy.
	// Deletes an existing enrich policy and its enrich index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-enrich-policy-api.html
	DeletePolicy enrich_delete_policy.NewDeletePolicy
	// Run an enrich policy.
	// Create the enrich index for an existing enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/execute-enrich-policy-api.html
	ExecutePolicy enrich_execute_policy.NewExecutePolicy
	// Get an enrich policy.
	// Returns information about an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-enrich-policy-api.html
	GetPolicy enrich_get_policy.NewGetPolicy
	// Create an enrich policy.
	// Creates an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-enrich-policy-api.html
	PutPolicy enrich_put_policy.NewPutPolicy
	// Get enrich stats.
	// Returns enrich coordinator statistics and information about enrich policies
	// that are currently executing.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-stats-api.html
	Stats enrich_stats.NewStats
}

type Eql struct {
	// Delete an async EQL search.
	// Delete an async EQL search or a stored synchronous EQL search.
	// The API also deletes results for the search.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-eql-delete
	Delete eql_delete.NewDelete
	// Get async EQL search results.
	// Get the current status and available results for an async EQL search or a
	// stored synchronous EQL search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-search-api.html
	Get eql_get.NewGet
	// Get the async EQL status.
	// Get the current status for an async EQL search or a stored synchronous EQL
	// search without returning results.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-status-api.html
	GetStatus eql_get_status.NewGetStatus
	// Get EQL search results.
	// Returns search results for an Event Query Language (EQL) query.
	// EQL assumes each document in a data stream or index corresponds to an event.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Search eql_search.NewSearch
}

type Esql struct {
	// Run an async ES|QL query.
	// Asynchronously run an ES|QL (Elasticsearch query language) query, monitor its
	// progress, and retrieve results when they become available.
	//
	// The API accepts the same parameters and request body as the synchronous query
	// API, along with additional async related properties.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-async-query-api.html
	AsyncQuery esql_async_query.NewAsyncQuery
	// Delete an async ES|QL query.
	// If the query is still running, it is cancelled.
	// Otherwise, the stored results are deleted.
	//
	// If the Elasticsearch security features are enabled, only the following users
	// can use this API to delete a query:
	//
	// * The authenticated user that submitted the original query request
	// * Users with the `cancel_task` cluster privilege
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-async-query-delete-api.html
	AsyncQueryDelete esql_async_query_delete.NewAsyncQueryDelete
	// Get async ES|QL query results.
	// Get the current status and available results or stored results for an ES|QL
	// asynchronous query.
	// If the Elasticsearch security features are enabled, only the user who first
	// submitted the ES|QL query can retrieve the results using this API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-async-query-get-api.html
	AsyncQueryGet esql_async_query_get.NewAsyncQueryGet
	// Stop async ES|QL query.
	//
	// This API interrupts the query execution and returns the results so far.
	// If the Elasticsearch security features are enabled, only the user who first
	// submitted the ES|QL query can stop it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-async-query-stop-api.html
	AsyncQueryStop esql_async_query_stop.NewAsyncQueryStop
	// Run an ES|QL query.
	// Get search results for an ES|QL (Elasticsearch query language) query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-rest.html
	Query esql_query.NewQuery
}

type Features struct {
	// Get the features.
	// Get a list of features that can be included in snapshots using the
	// `feature_states` field when creating a snapshot.
	// You can use this API to determine which feature states to include when taking
	// a snapshot.
	// By default, all feature states are included in a snapshot if that snapshot
	// includes the global state, or none if it does not.
	//
	// A feature state includes one or more system indices necessary for a given
	// feature to function.
	// In order to ensure data integrity, all system indices that comprise a feature
	// state are snapshotted and restored together.
	//
	// The features listed by this API are a combination of built-in features and
	// features defined by plugins.
	// In order for a feature state to be listed in this API and recognized as a
	// valid feature state by the create snapshot API, the plugin that defines that
	// feature must be installed on the master node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-features-api.html
	GetFeatures features_get_features.NewGetFeatures
	// Reset the features.
	// Clear all of the state information stored in system indices by Elasticsearch
	// features, including the security and machine learning indices.
	//
	// WARNING: Intended for development and testing use only. Do not reset features
	// on a production cluster.
	//
	// Return a cluster to the same state as a new installation by resetting the
	// feature state for all Elasticsearch features.
	// This deletes all state information stored in system indices.
	//
	// The response code is HTTP 200 if the state is successfully reset for all
	// features.
	// It is HTTP 500 if the reset operation failed for any feature.
	//
	// Note that select features might provide a way to reset particular system
	// indices.
	// Using this API resets all features, both those that are built-in and
	// implemented as plugins.
	//
	// To list the features that will be affected, use the get features API.
	//
	// IMPORTANT: The features installed on the node you submit this request to are
	// the features that will be reset. Run on the master node if you have any
	// doubts about which plugins are installed on individual nodes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-features-api.html
	ResetFeatures features_reset_features.NewResetFeatures
}

type Fleet struct {
	// Get global checkpoints.
	//
	// Get the current global checkpoints for an index.
	// This API is designed for internal use by the Fleet server project.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-global-checkpoints.html
	GlobalCheckpoints fleet_global_checkpoints.NewGlobalCheckpoints
	// Executes several fleet searches with a single API request.
	//
	// The API follows the same structure as the multi search (`_msearch`) API.
	// However, similar to the fleet search API, it supports the
	// `wait_for_checkpoints` parameter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/fleet-multi-search.html
	Msearch fleet_msearch.NewMsearch
	// Creates a secret stored by Fleet.
	//
	PostSecret fleet_post_secret.NewPostSecret
	// The purpose of the fleet search api is to provide a search api where the
	// search will only be executed
	// after provided checkpoint has been processed and is visible for searches
	// inside of Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/fleet-search.html
	Search fleet_search.NewSearch
}

type Graph struct {
	// Explore graph analytics.
	// Extract and summarize information about the documents and terms in an
	// Elasticsearch data stream or index.
	// The easiest way to understand the behavior of this API is to use the Graph UI
	// to explore connections.
	// An initial request to the `_explore` API contains a seed query that
	// identifies the documents of interest and specifies the fields that define the
	// vertices and connections you want to include in the graph.
	// Subsequent requests enable you to spider out from one more vertices of
	// interest.
	// You can exclude vertices that have already been returned.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/graph-explore-api.html
	Explore graph_explore.NewExplore
}

type Ilm struct {
	// Delete a lifecycle policy.
	// You cannot delete policies that are currently in use. If the policy is being
	// used to manage any indices, the request fails and returns an error.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-delete-lifecycle.html
	DeleteLifecycle ilm_delete_lifecycle.NewDeleteLifecycle
	// Explain the lifecycle state.
	// Get the current lifecycle status for one or more indices.
	// For data streams, the API retrieves the current lifecycle status for the
	// stream's backing indices.
	//
	// The response indicates when the index entered each lifecycle state, provides
	// the definition of the running phase, and information about any failures.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-explain-lifecycle.html
	ExplainLifecycle ilm_explain_lifecycle.NewExplainLifecycle
	// Get lifecycle policies.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-lifecycle.html
	GetLifecycle ilm_get_lifecycle.NewGetLifecycle
	// Get the ILM status.
	//
	// Get the current index lifecycle management status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-status.html
	GetStatus ilm_get_status.NewGetStatus
	// Migrate to data tiers routing.
	// Switch the indices, ILM policies, and legacy, composable, and component
	// templates from using custom node attributes and attribute-based allocation
	// filters to using data tiers.
	// Optionally, delete one legacy index template.
	// Using node roles enables ILM to automatically move the indices between data
	// tiers.
	//
	// Migrating away from custom node attributes routing can be manually performed.
	// This API provides an automated way of performing three out of the four manual
	// steps listed in the migration guide:
	//
	// 1. Stop setting the custom hot attribute on new indices.
	// 1. Remove custom allocation settings from existing ILM policies.
	// 1. Replace custom allocation settings from existing indices with the
	// corresponding tier preference.
	//
	// ILM must be stopped before performing the migration.
	// Use the stop ILM and get ILM status APIs to wait until the reported operation
	// mode is `STOPPED`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-migrate-to-data-tiers.html
	MigrateToDataTiers ilm_migrate_to_data_tiers.NewMigrateToDataTiers
	// Move to a lifecycle step.
	// Manually move an index into a specific step in the lifecycle policy and run
	// that step.
	//
	// WARNING: This operation can result in the loss of data. Manually moving an
	// index into a specific step runs that step even if it has already been
	// performed. This is a potentially destructive action and this should be
	// considered an expert level API.
	//
	// You must specify both the current step and the step to be executed in the
	// body of the request.
	// The request will fail if the current step does not match the step currently
	// running for the index
	// This is to prevent the index from being moved from an unexpected step into
	// the next step.
	//
	// When specifying the target (`next_step`) to which the index will be moved,
	// either the name or both the action and name fields are optional.
	// If only the phase is specified, the index will move to the first step of the
	// first action in the target phase.
	// If the phase and action are specified, the index will move to the first step
	// of the specified action in the specified phase.
	// Only actions specified in the ILM policy are considered valid.
	// An index cannot move to a step that is not part of its policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-move-to-step.html
	MoveToStep ilm_move_to_step.NewMoveToStep
	// Create or update a lifecycle policy.
	// If the specified policy exists, it is replaced and the policy version is
	// incremented.
	//
	// NOTE: Only the latest version of the policy is stored, you cannot revert to
	// previous versions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-put-lifecycle.html
	PutLifecycle ilm_put_lifecycle.NewPutLifecycle
	// Remove policies from an index.
	// Remove the assigned lifecycle policies from an index or a data stream's
	// backing indices.
	// It also stops managing the indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-remove-policy.html
	RemovePolicy ilm_remove_policy.NewRemovePolicy
	// Retry a policy.
	// Retry running the lifecycle policy for an index that is in the ERROR step.
	// The API sets the policy back to the step where the error occurred and runs
	// the step.
	// Use the explain lifecycle state API to determine whether an index is in the
	// ERROR step.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-retry-policy.html
	Retry ilm_retry.NewRetry
	// Start the ILM plugin.
	// Start the index lifecycle management plugin if it is currently stopped.
	// ILM is started automatically when the cluster is formed.
	// Restarting ILM is necessary only when it has been stopped using the stop ILM
	// API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-start.html
	Start ilm_start.NewStart
	// Stop the ILM plugin.
	// Halt all lifecycle management operations and stop the index lifecycle
	// management plugin.
	// This is useful when you are performing maintenance on the cluster and need to
	// prevent ILM from performing any actions on your indices.
	//
	// The API returns as soon as the stop request has been acknowledged, but the
	// plugin might continue to run until in-progress operations complete and the
	// plugin can be safely stopped.
	// Use the get ILM status API to check whether ILM is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-stop.html
	Stop ilm_stop.NewStop
}

type Indices struct {
	// Add an index block.
	//
	// Add an index block to an index.
	// Index blocks limit the operations allowed on an index by blocking specific
	// operation types.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules-blocks.html#add-index-block
	AddBlock indices_add_block.NewAddBlock
	// Get tokens from text analysis.
	// The analyze API performs analysis on a text string and returns the resulting
	// tokens.
	//
	// Generating excessive amount of tokens may cause a node to run out of memory.
	// The `index.analyze.max_token_count` setting enables you to limit the number
	// of tokens that can be produced.
	// If more than this limit of tokens gets generated, an error occurs.
	// The `_analyze` endpoint without a specified index will always use `10000` as
	// its limit.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-analyze
	Analyze indices_analyze.NewAnalyze
	// Cancel a migration reindex operation.
	//
	// Cancel a migration reindex attempt for a data stream or index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migrate-data-stream.html
	CancelMigrateReindex indices_cancel_migrate_reindex.NewCancelMigrateReindex
	// Clear the cache.
	// Clear the cache of one or more indices.
	// For data streams, the API clears the caches of the stream's backing indices.
	//
	// By default, the clear cache API clears all caches.
	// To clear only specific caches, use the `fielddata`, `query`, or `request`
	// parameters.
	// To clear the cache only of specific fields, use the `fields` parameter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-clearcache.html
	ClearCache indices_clear_cache.NewClearCache
	// Clone an index.
	// Clone an existing index into a new index.
	// Each original primary shard is cloned into a new primary shard in the new
	// index.
	//
	// IMPORTANT: Elasticsearch does not apply index templates to the resulting
	// index.
	// The API also does not copy index metadata from the original index.
	// Index metadata includes aliases, index lifecycle management phase
	// definitions, and cross-cluster replication (CCR) follower information.
	// For example, if you clone a CCR follower index, the resulting clone will not
	// be a follower index.
	//
	// The clone API copies most index settings from the source index to the
	// resulting index, with the exception of `index.number_of_replicas` and
	// `index.auto_expand_replicas`.
	// To set the number of replicas in the resulting index, configure these
	// settings in the clone request.
	//
	// Cloning works as follows:
	//
	// * First, it creates a new target index with the same definition as the source
	// index.
	// * Then it hard-links segments from the source index into the target index. If
	// the file system does not support hard-linking, all segments are copied into
	// the new index, which is a much more time consuming process.
	// * Finally, it recovers the target index as though it were a closed index
	// which had just been re-opened.
	//
	// IMPORTANT: Indices can only be cloned if they meet the following
	// requirements:
	//
	// * The index must be marked as read-only and have a cluster health status of
	// green.
	// * The target index must not exist.
	// * The source index must have the same number of primary shards as the target
	// index.
	// * The node handling the clone process must have sufficient free disk space to
	// accommodate a second copy of the existing index.
	//
	// The current write index on a data stream cannot be cloned.
	// In order to clone the current write index, the data stream must first be
	// rolled over so that a new write index is created and then the previous write
	// index can be cloned.
	//
	// NOTE: Mappings cannot be specified in the `_clone` request. The mappings of
	// the source index will be used for the target index.
	//
	// **Monitor the cloning process**
	//
	// The cloning process can be monitored with the cat recovery API or the cluster
	// health API can be used to wait until all primary shards have been allocated
	// by setting the `wait_for_status` parameter to `yellow`.
	//
	// The `_clone` API returns as soon as the target index has been added to the
	// cluster state, before any shards have been allocated.
	// At this point, all shards are in the state unassigned.
	// If, for any reason, the target index can't be allocated, its primary shard
	// will remain unassigned until it can be allocated on that node.
	//
	// Once the primary shard is allocated, it moves to state initializing, and the
	// clone process begins.
	// When the clone operation completes, the shard will become active.
	// At that point, Elasticsearch will try to allocate any replicas and may decide
	// to relocate the primary shard to another node.
	//
	// **Wait for active shards**
	//
	// Because the clone operation creates a new index to clone the shards to, the
	// wait for active shards setting on index creation applies to the clone index
	// action as well.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-clone-index.html
	Clone indices_clone.NewClone
	// Close an index.
	// A closed index is blocked for read or write operations and does not allow all
	// operations that opened indices allow.
	// It is not possible to index documents or to search for documents in a closed
	// index.
	// Closed indices do not have to maintain internal data structures for indexing
	// or searching documents, which results in a smaller overhead on the cluster.
	//
	// When opening or closing an index, the master node is responsible for
	// restarting the index shards to reflect the new state of the index.
	// The shards will then go through the normal recovery process.
	// The data of opened and closed indices is automatically replicated by the
	// cluster to ensure that enough shard copies are safely kept around at all
	// times.
	//
	// You can open and close multiple indices.
	// An error is thrown if the request explicitly refers to a missing index.
	// This behaviour can be turned off using the `ignore_unavailable=true`
	// parameter.
	//
	// By default, you must explicitly name the indices you are opening or closing.
	// To open or close indices with `_all`, `*`, or other wildcard expressions,
	// change the` action.destructive_requires_name` setting to `false`. This
	// setting can also be changed with the cluster update settings API.
	//
	// Closed indices consume a significant amount of disk-space which can cause
	// problems in managed environments.
	// Closing indices can be turned off with the cluster settings API by setting
	// `cluster.indices.close.enable` to `false`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-close.html
	Close indices_close.NewClose
	// Create an index.
	// You can use the create index API to add a new index to an Elasticsearch
	// cluster.
	// When creating an index, you can specify the following:
	//
	// * Settings for the index.
	// * Mappings for fields in the index.
	// * Index aliases
	//
	// **Wait for active shards**
	//
	// By default, index creation will only return a response to the client when the
	// primary copies of each shard have been started, or the request times out.
	// The index creation response will indicate what happened.
	// For example, `acknowledged` indicates whether the index was successfully
	// created in the cluster, `while shards_acknowledged` indicates whether the
	// requisite number of shard copies were started for each shard in the index
	// before timing out.
	// Note that it is still possible for either `acknowledged` or
	// `shards_acknowledged` to be `false`, but for the index creation to be
	// successful.
	// These values simply indicate whether the operation completed before the
	// timeout.
	// If `acknowledged` is false, the request timed out before the cluster state
	// was updated with the newly created index, but it probably will be created
	// sometime soon.
	// If `shards_acknowledged` is false, then the request timed out before the
	// requisite number of shards were started (by default just the primaries), even
	// if the cluster state was successfully updated to reflect the newly created
	// index (that is to say, `acknowledged` is `true`).
	//
	// You can change the default of only waiting for the primary shards to start
	// through the index setting `index.write.wait_for_active_shards`.
	// Note that changing this setting will also affect the `wait_for_active_shards`
	// value on all subsequent write operations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
	Create indices_create.NewCreate
	// Create a data stream.
	//
	// You must have a matching index template with data stream enabled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-data-stream.html
	CreateDataStream indices_create_data_stream.NewCreateDataStream
	// Create an index from a source index.
	//
	// Copy the mappings and settings from the source index to a destination index
	// while allowing request settings and mappings to override the source values.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migrate-data-stream.html
	CreateFrom indices_create_from.NewCreateFrom
	// Get data stream stats.
	//
	// Get statistics for one or more data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-stream-stats-api.html
	DataStreamsStats indices_data_streams_stats.NewDataStreamsStats
	// Delete indices.
	// Deleting an index deletes its documents, shards, and metadata.
	// It does not delete related Kibana components, such as data views,
	// visualizations, or dashboards.
	//
	// You cannot delete the current write index of a data stream.
	// To delete the index, you must roll over the data stream so a new write index
	// is created.
	// You can then use the delete index API to delete the previous write index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-index.html
	Delete indices_delete.NewDelete
	// Delete an alias.
	// Removes a data stream or index from an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-alias.html
	DeleteAlias indices_delete_alias.NewDeleteAlias
	// Delete data stream lifecycles.
	// Removes the data stream lifecycle from a data stream, rendering it not
	// managed by the data stream lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-delete-lifecycle.html
	DeleteDataLifecycle indices_delete_data_lifecycle.NewDeleteDataLifecycle
	// Delete data streams.
	// Deletes one or more data streams and their backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-data-stream.html
	DeleteDataStream indices_delete_data_stream.NewDeleteDataStream
	// Delete an index template.
	// The provided <index-template> may contain multiple template names separated
	// by a comma. If multiple template
	// names are specified then there is no wildcard support and the provided names
	// should match completely with
	// existing templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-template.html
	DeleteIndexTemplate indices_delete_index_template.NewDeleteIndexTemplate
	// Delete a legacy index template.
	// IMPORTANT: This documentation is about legacy index templates, which are
	// deprecated and will be replaced by the composable templates introduced in
	// Elasticsearch 7.8.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-template-v1.html
	DeleteTemplate indices_delete_template.NewDeleteTemplate
	// Analyze the index disk usage.
	// Analyze the disk usage of each field of an index or data stream.
	// This API might not support indices created in previous Elasticsearch
	// versions.
	// The result of a small index can be inaccurate as some parts of an index might
	// not be analyzed by the API.
	//
	// NOTE: The total size of fields of the analyzed shards of the index in the
	// response is usually smaller than the index `store_size` value because some
	// small metadata files are ignored and some parts of data files might not be
	// scanned by the API.
	// Since stored fields are stored together in a compressed format, the sizes of
	// stored fields are also estimates and can be inaccurate.
	// The stored size of the `_id` field is likely underestimated while the
	// `_source` field is overestimated.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-disk-usage.html
	DiskUsage indices_disk_usage.NewDiskUsage
	// Downsample an index.
	// Aggregate a time series (TSDS) index and store pre-computed statistical
	// summaries (`min`, `max`, `sum`, `value_count` and `avg`) for each metric
	// field grouped by a configured time interval.
	// For example, a TSDS index that contains metrics sampled every 10 seconds can
	// be downsampled to an hourly index.
	// All documents within an hour interval are summarized and stored as a single
	// document in the downsample index.
	//
	// NOTE: Only indices in a time series data stream are supported.
	// Neither field nor document level security can be defined on the source index.
	// The source index must be read only (`index.blocks.write: true`).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-downsample-data-stream.html
	Downsample indices_downsample.NewDownsample
	// Check indices.
	// Check if one or more indices, index aliases, or data streams exist.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-exists.html
	Exists indices_exists.NewExists
	// Check aliases.
	//
	// Check if one or more data stream or index aliases exist.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-exists-alias
	ExistsAlias indices_exists_alias.NewExistsAlias
	// Check index templates.
	//
	// Check whether index templates exist.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-exists-index-template
	ExistsIndexTemplate indices_exists_index_template.NewExistsIndexTemplate
	// Check existence of index templates.
	// Get information about whether index templates exist.
	// Index templates define settings, mappings, and aliases that can be applied
	// automatically to new indices.
	//
	// IMPORTANT: This documentation is about legacy index templates, which are
	// deprecated and will be replaced by the composable templates introduced in
	// Elasticsearch 7.8.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-template-exists-v1.html
	ExistsTemplate indices_exists_template.NewExistsTemplate
	// Get the status for a data stream lifecycle.
	// Get information about an index or data stream's current data stream lifecycle
	// status, such as time since index creation, time since rollover, the lifecycle
	// configuration managing the index, or any errors encountered during lifecycle
	// execution.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-explain-lifecycle.html
	ExplainDataLifecycle indices_explain_data_lifecycle.NewExplainDataLifecycle
	// Get field usage stats.
	// Get field usage information for each shard and field of an index.
	// Field usage statistics are automatically captured when queries are running on
	// a cluster.
	// A shard-level search request that accesses a given field, even if multiple
	// times during that request, is counted as a single use.
	//
	// The response body reports the per-shard usage count of the data structures
	// that back the fields in the index.
	// A given request will increment each count by a maximum value of 1, even if
	// the request accesses the same field multiple times.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/field-usage-stats.html
	FieldUsageStats indices_field_usage_stats.NewFieldUsageStats
	// Flush data streams or indices.
	// Flushing a data stream or index is the process of making sure that any data
	// that is currently only stored in the transaction log is also permanently
	// stored in the Lucene index.
	// When restarting, Elasticsearch replays any unflushed operations from the
	// transaction log into the Lucene index to bring it back into the state that it
	// was in before the restart.
	// Elasticsearch automatically triggers flushes as needed, using heuristics that
	// trade off the size of the unflushed transaction log against the cost of
	// performing each flush.
	//
	// After each operation has been flushed it is permanently stored in the Lucene
	// index.
	// This may mean that there is no need to maintain an additional copy of it in
	// the transaction log.
	// The transaction log is made up of multiple files, called generations, and
	// Elasticsearch will delete any generation files when they are no longer
	// needed, freeing up disk space.
	//
	// It is also possible to trigger a flush on one or more indices using the flush
	// API, although it is rare for users to need to call this API directly.
	// If you call the flush API after indexing some documents then a successful
	// response indicates that Elasticsearch has flushed all the documents that were
	// indexed before the flush API was called.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-flush.html
	Flush indices_flush.NewFlush
	// Force a merge.
	// Perform the force merge operation on the shards of one or more indices.
	// For data streams, the API forces a merge on the shards of the stream's
	// backing indices.
	//
	// Merging reduces the number of segments in each shard by merging some of them
	// together and also frees up the space used by deleted documents.
	// Merging normally happens automatically, but sometimes it is useful to trigger
	// a merge manually.
	//
	// WARNING: We recommend force merging only a read-only index (meaning the index
	// is no longer receiving writes).
	// When documents are updated or deleted, the old version is not immediately
	// removed but instead soft-deleted and marked with a "tombstone".
	// These soft-deleted documents are automatically cleaned up during regular
	// segment merges.
	// But force merge can cause very large (greater than 5 GB) segments to be
	// produced, which are not eligible for regular merges.
	// So the number of soft-deleted documents can then grow rapidly, resulting in
	// higher disk usage and worse search performance.
	// If you regularly force merge an index receiving writes, this can also make
	// snapshots more expensive, since the new documents can't be backed up
	// incrementally.
	//
	// **Blocks during a force merge**
	//
	// Calls to this API block until the merge is complete (unless request contains
	// `wait_for_completion=false`).
	// If the client connection is lost before completion then the force merge
	// process will continue in the background.
	// Any new requests to force merge the same indices will also block until the
	// ongoing force merge is complete.
	//
	// **Running force merge asynchronously**
	//
	// If the request contains `wait_for_completion=false`, Elasticsearch performs
	// some preflight checks, launches the request, and returns a task you can use
	// to get the status of the task.
	// However, you can not cancel this task as the force merge task is not
	// cancelable.
	// Elasticsearch creates a record of this task as a document at
	// `_tasks/<task_id>`.
	// When you are done with a task, you should delete the task document so
	// Elasticsearch can reclaim the space.
	//
	// **Force merging multiple indices**
	//
	// You can force merge multiple indices with a single request by targeting:
	//
	// * One or more data streams that contain multiple backing indices
	// * Multiple indices
	// * One or more aliases
	// * All data streams and indices in a cluster
	//
	// Each targeted shard is force-merged separately using the force_merge
	// threadpool.
	// By default each node only has a single `force_merge` thread which means that
	// the shards on that node are force-merged one at a time.
	// If you expand the `force_merge` threadpool on a node then it will force merge
	// its shards in parallel
	//
	// Force merge makes the storage for the shard being merged temporarily
	// increase, as it may require free space up to triple its size in case
	// `max_num_segments parameter` is set to `1`, to rewrite all segments into a
	// new one.
	//
	// **Data streams and time-based indices**
	//
	// Force-merging is useful for managing a data stream's older backing indices
	// and other time-based indices, particularly after a rollover.
	// In these cases, each index only receives indexing traffic for a certain
	// period of time.
	// Once an index receive no more writes, its shards can be force-merged to a
	// single segment.
	// This can be a good idea because single-segment shards can sometimes use
	// simpler and more efficient data structures to perform searches.
	// For example:
	//
	// ```
	// POST /.ds-my-data-stream-2099.03.07-000001/_forcemerge?max_num_segments=1
	// ```
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-forcemerge.html
	Forcemerge indices_forcemerge.NewForcemerge
	// Get index information.
	// Get information about one or more indices. For data streams, the API returns
	// information about the
	// stream’s backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-index.html
	Get indices_get.NewGet
	// Get aliases.
	// Retrieves information for one or more data stream or index aliases.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-alias.html
	GetAlias indices_get_alias.NewGetAlias
	// Get data stream lifecycles.
	//
	// Get the data stream lifecycle configuration of one or more data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-get-lifecycle.html
	GetDataLifecycle indices_get_data_lifecycle.NewGetDataLifecycle
	// Get data stream lifecycle stats.
	// Get statistics about the data streams that are managed by a data stream
	// lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-get-lifecycle-stats.html
	GetDataLifecycleStats indices_get_data_lifecycle_stats.NewGetDataLifecycleStats
	// Get data streams.
	//
	// Get information about one or more data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-data-stream.html
	GetDataStream indices_get_data_stream.NewGetDataStream
	// Get mapping definitions.
	// Retrieves mapping definitions for one or more fields.
	// For data streams, the API retrieves field mappings for the stream’s backing
	// indices.
	//
	// This API is useful if you don't need a complete mapping or if an index
	// mapping contains a large number of fields.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-field-mapping.html
	GetFieldMapping indices_get_field_mapping.NewGetFieldMapping
	// Get index templates.
	// Get information about one or more index templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-template.html
	GetIndexTemplate indices_get_index_template.NewGetIndexTemplate
	// Get mapping definitions.
	// For data streams, the API retrieves mappings for the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-mapping.html
	GetMapping indices_get_mapping.NewGetMapping
	// Get the migration reindexing status.
	//
	// Get the status of a migration reindex attempt for a data stream or index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migrate-data-stream.html
	GetMigrateReindexStatus indices_get_migrate_reindex_status.NewGetMigrateReindexStatus
	// Get index settings.
	// Get setting information for one or more indices.
	// For data streams, it returns setting information for the stream's backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-settings.html
	GetSettings indices_get_settings.NewGetSettings
	// Get legacy index templates.
	// Get information about one or more index templates.
	//
	// IMPORTANT: This documentation is about legacy index templates, which are
	// deprecated and will be replaced by the composable templates introduced in
	// Elasticsearch 7.8.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-template-v1.html
	GetTemplate indices_get_template.NewGetTemplate
	// Reindex legacy backing indices.
	//
	// Reindex all legacy backing indices for a data stream.
	// This operation occurs in a persistent task.
	// The persistent task ID is returned immediately and the reindexing work is
	// completed in that task.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migrate-data-stream.html
	MigrateReindex indices_migrate_reindex.NewMigrateReindex
	// Convert an index alias to a data stream.
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
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-migrate-to-data-stream
	MigrateToDataStream indices_migrate_to_data_stream.NewMigrateToDataStream
	// Update data streams.
	// Performs one or more data stream modification actions in a single atomic
	// operation.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-modify-data-stream
	ModifyDataStream indices_modify_data_stream.NewModifyDataStream
	// Open a closed index.
	// For data streams, the API opens any closed backing indices.
	//
	// A closed index is blocked for read/write operations and does not allow all
	// operations that opened indices allow.
	// It is not possible to index documents or to search for documents in a closed
	// index.
	// This allows closed indices to not have to maintain internal data structures
	// for indexing or searching documents, resulting in a smaller overhead on the
	// cluster.
	//
	// When opening or closing an index, the master is responsible for restarting
	// the index shards to reflect the new state of the index.
	// The shards will then go through the normal recovery process.
	// The data of opened or closed indices is automatically replicated by the
	// cluster to ensure that enough shard copies are safely kept around at all
	// times.
	//
	// You can open and close multiple indices.
	// An error is thrown if the request explicitly refers to a missing index.
	// This behavior can be turned off by using the `ignore_unavailable=true`
	// parameter.
	//
	// By default, you must explicitly name the indices you are opening or closing.
	// To open or close indices with `_all`, `*`, or other wildcard expressions,
	// change the `action.destructive_requires_name` setting to `false`.
	// This setting can also be changed with the cluster update settings API.
	//
	// Closed indices consume a significant amount of disk-space which can cause
	// problems in managed environments.
	// Closing indices can be turned off with the cluster settings API by setting
	// `cluster.indices.close.enable` to `false`.
	//
	// Because opening or closing an index allocates its shards, the
	// `wait_for_active_shards` setting on index creation applies to the `_open` and
	// `_close` index actions as well.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-open-close.html
	Open indices_open.NewOpen
	// Promote a data stream.
	// Promote a data stream from a replicated data stream managed by cross-cluster
	// replication (CCR) to a regular data stream.
	//
	// With CCR auto following, a data stream from a remote cluster can be
	// replicated to the local cluster.
	// These data streams can't be rolled over in the local cluster.
	// These replicated data streams roll over only if the upstream data stream
	// rolls over.
	// In the event that the remote cluster is no longer available, the data stream
	// in the local cluster can be promoted to a regular data stream, which allows
	// these data streams to be rolled over in the local cluster.
	//
	// NOTE: When promoting a data stream, ensure the local cluster has a data
	// stream enabled index template that matches the data stream.
	// If this is missing, the data stream will not be able to roll over until a
	// matching index template is created.
	// This will affect the lifecycle management of the data stream and interfere
	// with the data stream size and retention.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-promote-data-stream
	PromoteDataStream indices_promote_data_stream.NewPromoteDataStream
	// Create or update an alias.
	// Adds a data stream or index to an alias.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-put-alias
	PutAlias indices_put_alias.NewPutAlias
	// Update data stream lifecycles.
	// Update the data stream lifecycle of the specified data streams.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-put-data-lifecycle
	PutDataLifecycle indices_put_data_lifecycle.NewPutDataLifecycle
	// Create or update an index template.
	// Index templates define settings, mappings, and aliases that can be applied
	// automatically to new indices.
	//
	// Elasticsearch applies templates to new indices based on an wildcard pattern
	// that matches the index name.
	// Index templates are applied during data stream or index creation.
	// For data streams, these settings and mappings are applied when the stream's
	// backing indices are created.
	// Settings and mappings specified in a create index API request override any
	// settings or mappings specified in an index template.
	// Changes to index templates do not affect existing indices, including the
	// existing backing indices of a data stream.
	//
	// You can use C-style `/* *\/` block comments in index templates.
	// You can include comments anywhere in the request body, except before the
	// opening curly bracket.
	//
	// **Multiple matching templates**
	//
	// If multiple index templates match the name of a new index or data stream, the
	// template with the highest priority is used.
	//
	// Multiple templates with overlapping index patterns at the same priority are
	// not allowed and an error will be thrown when attempting to create a template
	// matching an existing index template at identical priorities.
	//
	// **Composing aliases, mappings, and settings**
	//
	// When multiple component templates are specified in the `composed_of` field
	// for an index template, they are merged in the order specified, meaning that
	// later component templates override earlier component templates.
	// Any mappings, settings, or aliases from the parent index template are merged
	// in next.
	// Finally, any configuration on the index request itself is merged.
	// Mapping definitions are merged recursively, which means that later mapping
	// components can introduce new field mappings and update the mapping
	// configuration.
	// If a field mapping is already contained in an earlier component, its
	// definition will be completely overwritten by the later one.
	// This recursive merging strategy applies not only to field mappings, but also
	// root options like `dynamic_templates` and `meta`.
	// If an earlier component contains a `dynamic_templates` block, then by default
	// new `dynamic_templates` entries are appended onto the end.
	// If an entry already exists with the same key, then it is overwritten by the
	// new definition.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-put-index-template
	PutIndexTemplate indices_put_index_template.NewPutIndexTemplate
	// Update field mappings.
	// Add new fields to an existing data stream or index.
	// You can also use this API to change the search settings of existing fields
	// and add new properties to existing object fields.
	// For data streams, these changes are applied to all backing indices by
	// default.
	//
	// **Add multi-fields to an existing field**
	//
	// Multi-fields let you index the same field in different ways.
	// You can use this API to update the fields mapping parameter and enable
	// multi-fields for an existing field.
	// WARNING: If an index (or data stream) contains documents when you add a
	// multi-field, those documents will not have values for the new multi-field.
	// You can populate the new multi-field with the update by query API.
	//
	// **Change supported mapping parameters for an existing field**
	//
	// The documentation for each mapping parameter indicates whether you can update
	// it for an existing field using this API.
	// For example, you can use the update mapping API to update the `ignore_above`
	// parameter.
	//
	// **Change the mapping of an existing field**
	//
	// Except for supported mapping parameters, you can't change the mapping or
	// field type of an existing field.
	// Changing an existing field could invalidate data that's already indexed.
	//
	// If you need to change the mapping of a field in a data stream's backing
	// indices, refer to documentation about modifying data streams.
	// If you need to change the mapping of a field in other indices, create a new
	// index with the correct mapping and reindex your data into that index.
	//
	// **Rename a field**
	//
	// Renaming a field would invalidate data already indexed under the old field
	// name.
	// Instead, add an alias field to create an alternate field name.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-mapping.html
	PutMapping indices_put_mapping.NewPutMapping
	// Update index settings.
	// Changes dynamic index settings in real time.
	// For data streams, index setting changes are applied to all backing indices by
	// default.
	//
	// To revert a setting to the default value, use a null value.
	// The list of per-index settings that can be updated dynamically on live
	// indices can be found in index settings documentation.
	// To preserve existing settings from being updated, set the `preserve_existing`
	// parameter to `true`.
	//
	//  There are multiple valid ways to represent index settings in the request
	// body. You can specify only the setting, for example:
	//
	// ```
	// {
	//   "number_of_replicas": 1
	// }
	// ```
	//
	// Or you can use an `index` setting object:
	// ```
	// {
	//   "index": {
	//     "number_of_replicas": 1
	//   }
	// }
	// ```
	//
	// Or you can use dot annotation:
	// ```
	// {
	//   "index.number_of_replicas": 1
	// }
	// ```
	//
	// Or you can embed any of the aforementioned options in a `settings` object.
	// For example:
	//
	// ```
	// {
	//   "settings": {
	//     "index": {
	//       "number_of_replicas": 1
	//     }
	//   }
	// }
	// ```
	//
	// NOTE: You can only define new analyzers on closed indices.
	// To add an analyzer, you must close the index, define the analyzer, and reopen
	// the index.
	// You cannot close the write index of a data stream.
	// To update the analyzer for a data stream's write index and future backing
	// indices, update the analyzer in the index template used by the stream.
	// Then roll over the data stream to apply the new analyzer to the stream's
	// write index and future backing indices.
	// This affects searches and any new data added to the stream after the
	// rollover.
	// However, it does not affect the data stream's backing indices or their
	// existing data.
	// To change the analyzer for existing backing indices, you must create a new
	// data stream and reindex your data into it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-update-settings.html
	PutSettings indices_put_settings.NewPutSettings
	// Create or update a legacy index template.
	// Index templates define settings, mappings, and aliases that can be applied
	// automatically to new indices.
	// Elasticsearch applies templates to new indices based on an index pattern that
	// matches the index name.
	//
	// IMPORTANT: This documentation is about legacy index templates, which are
	// deprecated and will be replaced by the composable templates introduced in
	// Elasticsearch 7.8.
	//
	// Composable templates always take precedence over legacy templates.
	// If no composable template matches a new index, matching legacy templates are
	// applied according to their order.
	//
	// Index templates are only applied during index creation.
	// Changes to index templates do not affect existing indices.
	// Settings and mappings specified in create index API requests override any
	// settings or mappings specified in an index template.
	//
	// You can use C-style `/* *\/` block comments in index templates.
	// You can include comments anywhere in the request body, except before the
	// opening curly bracket.
	//
	// **Indices matching multiple templates**
	//
	// Multiple index templates can potentially match an index, in this case, both
	// the settings and mappings are merged into the final configuration of the
	// index.
	// The order of the merging can be controlled using the order parameter, with
	// lower order being applied first, and higher orders overriding them.
	// NOTE: Multiple matching templates with the same order value will result in a
	// non-deterministic merging order.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates-v1.html
	PutTemplate indices_put_template.NewPutTemplate
	// Get index recovery information.
	// Get information about ongoing and completed shard recoveries for one or more
	// indices.
	// For data streams, the API returns information for the stream's backing
	// indices.
	//
	// All recoveries, whether ongoing or complete, are kept in the cluster state
	// and may be reported on at any time.
	//
	// Shard recovery is the process of initializing a shard copy, such as restoring
	// a primary shard from a snapshot or creating a replica shard from a primary
	// shard.
	// When a shard recovery completes, the recovered shard is available for search
	// and indexing.
	//
	// Recovery automatically occurs during the following processes:
	//
	// * When creating an index for the first time.
	// * When a node rejoins the cluster and starts up any missing primary shard
	// copies using the data that it holds in its data path.
	// * Creation of new replica shard copies from the primary.
	// * Relocation of a shard copy to a different node in the same cluster.
	// * A snapshot restore operation.
	// * A clone, shrink, or split operation.
	//
	// You can determine the cause of a shard recovery using the recovery or cat
	// recovery APIs.
	//
	// The index recovery API reports information about completed recoveries only
	// for shard copies that currently exist in the cluster.
	// It only reports the last recovery for each shard copy and does not report
	// historical information about earlier recoveries, nor does it report
	// information about the recoveries of shard copies that no longer exist.
	// This means that if a shard copy completes a recovery and then Elasticsearch
	// relocates it onto a different node then the information about the original
	// recovery will not be shown in the recovery API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-recovery.html
	Recovery indices_recovery.NewRecovery
	// Refresh an index.
	// A refresh makes recent operations performed on one or more indices available
	// for search.
	// For data streams, the API runs the refresh operation on the stream’s backing
	// indices.
	//
	// By default, Elasticsearch periodically refreshes indices every second, but
	// only on indices that have received one search request or more in the last 30
	// seconds.
	// You can change this default interval with the `index.refresh_interval`
	// setting.
	//
	// Refresh requests are synchronous and do not return a response until the
	// refresh operation completes.
	//
	// Refreshes are resource-intensive.
	// To ensure good cluster performance, it's recommended to wait for
	// Elasticsearch's periodic refresh rather than performing an explicit refresh
	// when possible.
	//
	// If your application workflow indexes documents and then runs a search to
	// retrieve the indexed document, it's recommended to use the index API's
	// `refresh=wait_for` query parameter option.
	// This option ensures the indexing operation waits for a periodic refresh
	// before running the search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-refresh.html
	Refresh indices_refresh.NewRefresh
	// Reload search analyzers.
	// Reload an index's search analyzers and their resources.
	// For data streams, the API reloads search analyzers and resources for the
	// stream's backing indices.
	//
	// IMPORTANT: After reloading the search analyzers you should clear the request
	// cache to make sure it doesn't contain responses derived from the previous
	// versions of the analyzer.
	//
	// You can use the reload search analyzers API to pick up changes to synonym
	// files used in the `synonym_graph` or `synonym` token filter of a search
	// analyzer.
	// To be eligible, the token filter must have an `updateable` flag of `true` and
	// only be used in search analyzers.
	//
	// NOTE: This API does not perform a reload for each shard of an index.
	// Instead, it performs a reload for each node containing index shards.
	// As a result, the total shard count returned by the API can differ from the
	// number of index shards.
	// Because reloading affects every node with an index shard, it is important to
	// update the synonym file on every data node in the cluster--including nodes
	// that don't contain a shard replica--before using this API.
	// This ensures the synonym file is updated everywhere in the cluster in case
	// shards are relocated in the future.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-reload-analyzers.html
	ReloadSearchAnalyzers indices_reload_search_analyzers.NewReloadSearchAnalyzers
	// Resolve the cluster.
	//
	// Resolve the specified index expressions to return information about each
	// cluster, including the local "querying" cluster, if included.
	// If no index expression is provided, the API will return information about all
	// the remote clusters that are configured on the querying cluster.
	//
	// This endpoint is useful before doing a cross-cluster search in order to
	// determine which remote clusters should be included in a search.
	//
	// You use the same index expression with this endpoint as you would for
	// cross-cluster search.
	// Index and cluster exclusions are also supported with this endpoint.
	//
	// For each cluster in the index expression, information is returned about:
	//
	// * Whether the querying ("local") cluster is currently connected to each
	// remote cluster specified in the index expression. Note that this endpoint
	// actively attempts to contact the remote clusters, unlike the `remote/info`
	// endpoint.
	// * Whether each remote cluster is configured with `skip_unavailable` as `true`
	// or `false`.
	// * Whether there are any indices, aliases, or data streams on that cluster
	// that match the index expression.
	// * Whether the search is likely to have errors returned when you do the
	// cross-cluster search (including any authorization errors if you do not have
	// permission to query the index).
	// * Cluster version information, including the Elasticsearch server version.
	//
	// For example, `GET /_resolve/cluster/my-index-*,cluster*:my-index-*` returns
	// information about the local cluster and all remotely configured clusters that
	// start with the alias `cluster*`.
	// Each cluster returns information about whether it has any indices, aliases or
	// data streams that match `my-index-*`.
	//
	// ## Note on backwards compatibility
	// The ability to query without an index expression was added in version 8.18,
	// so when
	// querying remote clusters older than that, the local cluster will send the
	// index
	// expression `dummy*` to those remote clusters. Thus, if an errors occur, you
	// may see a reference
	// to that index expression even though you didn't request it. If it causes a
	// problem, you can
	// instead include an index expression like `*:*` to bypass the issue.
	//
	// ## Advantages of using this endpoint before a cross-cluster search
	//
	// You may want to exclude a cluster or index from a search when:
	//
	// * A remote cluster is not currently connected and is configured with
	// `skip_unavailable=false`. Running a cross-cluster search under those
	// conditions will cause the entire search to fail.
	// * A cluster has no matching indices, aliases or data streams for the index
	// expression (or your user does not have permissions to search them). For
	// example, suppose your index expression is `logs*,remote1:logs*` and the
	// remote1 cluster has no indices, aliases or data streams that match `logs*`.
	// In that case, that cluster will return no results from that cluster if you
	// include it in a cross-cluster search.
	// * The index expression (combined with any query parameters you specify) will
	// likely cause an exception to be thrown when you do the search. In these
	// cases, the "error" field in the `_resolve/cluster` response will be present.
	// (This is also where security/permission errors will be shown.)
	// * A remote cluster is an older version that does not support the feature you
	// want to use in your search.
	//
	// ## Test availability of remote clusters
	//
	// The `remote/info` endpoint is commonly used to test whether the "local"
	// cluster (the cluster being queried) is connected to its remote clusters, but
	// it does not necessarily reflect whether the remote cluster is available or
	// not.
	// The remote cluster may be available, while the local cluster is not currently
	// connected to it.
	//
	// You can use the `_resolve/cluster` API to attempt to reconnect to remote
	// clusters.
	// For example with `GET _resolve/cluster` or `GET _resolve/cluster/*:*`.
	// The `connected` field in the response will indicate whether it was
	// successful.
	// If a connection was (re-)established, this will also cause the `remote/info`
	// endpoint to now indicate a connected status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-resolve-cluster-api.html
	ResolveCluster indices_resolve_cluster.NewResolveCluster
	// Resolve indices.
	// Resolve the names and/or index patterns for indices, aliases, and data
	// streams.
	// Multiple patterns and remote clusters are supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-resolve-index-api.html
	ResolveIndex indices_resolve_index.NewResolveIndex
	// Roll over to a new index.
	// TIP: It is recommended to use the index lifecycle rollover action to automate
	// rollovers.
	//
	// The rollover API creates a new index for a data stream or index alias.
	// The API behavior depends on the rollover target.
	//
	// **Roll over a data stream**
	//
	// If you roll over a data stream, the API creates a new write index for the
	// stream.
	// The stream's previous write index becomes a regular backing index.
	// A rollover also increments the data stream's generation.
	//
	// **Roll over an index alias with a write index**
	//
	// TIP: Prior to Elasticsearch 7.9, you'd typically use an index alias with a
	// write index to manage time series data.
	// Data streams replace this functionality, require less maintenance, and
	// automatically integrate with data tiers.
	//
	// If an index alias points to multiple indices, one of the indices must be a
	// write index.
	// The rollover API creates a new write index for the alias with
	// `is_write_index` set to `true`.
	// The API also `sets is_write_index` to `false` for the previous write index.
	//
	// **Roll over an index alias with one index**
	//
	// If you roll over an index alias that points to only one index, the API
	// creates a new index for the alias and removes the original index from the
	// alias.
	//
	// NOTE: A rollover creates a new index and is subject to the
	// `wait_for_active_shards` setting.
	//
	// **Increment index names for an alias**
	//
	// When you roll over an index alias, you can specify a name for the new index.
	// If you don't specify a name and the current index ends with `-` and a number,
	// such as `my-index-000001` or `my-index-3`, the new index name increments that
	// number.
	// For example, if you roll over an alias with a current index of
	// `my-index-000001`, the rollover creates a new index named `my-index-000002`.
	// This number is always six characters and zero-padded, regardless of the
	// previous index's name.
	//
	// If you use an index alias for time series data, you can use date math in the
	// index name to track the rollover date.
	// For example, you can create an alias that points to an index named
	// `<my-index-{now/d}-000001>`.
	// If you create the index on May 6, 2099, the index's name is
	// `my-index-2099.05.06-000001`.
	// If you roll over the alias on May 7, 2099, the new index's name is
	// `my-index-2099.05.07-000002`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-rollover-index.html
	Rollover indices_rollover.NewRollover
	// Get index segments.
	// Get low-level information about the Lucene segments in index shards.
	// For data streams, the API returns information about the stream's backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-segments.html
	Segments indices_segments.NewSegments
	// Get index shard stores.
	// Get store information about replica shards in one or more indices.
	// For data streams, the API retrieves store information for the stream's
	// backing indices.
	//
	// The index shard stores API returns the following information:
	//
	// * The node on which each replica shard exists.
	// * The allocation ID for each replica shard.
	// * A unique ID for each replica shard.
	// * Any errors encountered while opening the shard index or from an earlier
	// failure.
	//
	// By default, the API returns store information only for primary shards that
	// are unassigned or have one or more unassigned replica shards.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shards-stores.html
	ShardStores indices_shard_stores.NewShardStores
	// Shrink an index.
	// Shrink an index into a new index with fewer primary shards.
	//
	// Before you can shrink an index:
	//
	// * The index must be read-only.
	// * A copy of every shard in the index must reside on the same node.
	// * The index must have a green health status.
	//
	// To make shard allocation easier, we recommend you also remove the index's
	// replica shards.
	// You can later re-add replica shards as part of the shrink operation.
	//
	// The requested number of primary shards in the target index must be a factor
	// of the number of shards in the source index.
	// For example an index with 8 primary shards can be shrunk into 4, 2 or 1
	// primary shards or an index with 15 primary shards can be shrunk into 5, 3 or
	// 1.
	// If the number of shards in the index is a prime number it can only be shrunk
	// into a single primary shard
	//  Before shrinking, a (primary or replica) copy of every shard in the index
	// must be present on the same node.
	//
	// The current write index on a data stream cannot be shrunk. In order to shrink
	// the current write index, the data stream must first be rolled over so that a
	// new write index is created and then the previous write index can be shrunk.
	//
	// A shrink operation:
	//
	// * Creates a new target index with the same definition as the source index,
	// but with a smaller number of primary shards.
	// * Hard-links segments from the source index into the target index. If the
	// file system does not support hard-linking, then all segments are copied into
	// the new index, which is a much more time consuming process. Also if using
	// multiple data paths, shards on different data paths require a full copy of
	// segment files if they are not on the same disk since hardlinks do not work
	// across disks.
	// * Recovers the target index as though it were a closed index which had just
	// been re-opened. Recovers shards to the
	// `.routing.allocation.initial_recovery._id` index setting.
	//
	// IMPORTANT: Indices can only be shrunk if they satisfy the following
	// requirements:
	//
	// * The target index must not exist.
	// * The source index must have more primary shards than the target index.
	// * The number of primary shards in the target index must be a factor of the
	// number of primary shards in the source index. The source index must have more
	// primary shards than the target index.
	// * The index must not contain more than 2,147,483,519 documents in total
	// across all shards that will be shrunk into a single shard on the target index
	// as this is the maximum number of docs that can fit into a single shard.
	// * The node handling the shrink process must have sufficient free disk space
	// to accommodate a second copy of the existing index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shrink-index.html
	Shrink indices_shrink.NewShrink
	// Simulate an index.
	// Get the index configuration that would be applied to the specified index from
	// an existing index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-simulate-index.html
	SimulateIndexTemplate indices_simulate_index_template.NewSimulateIndexTemplate
	// Simulate an index template.
	// Get the index configuration that would be applied by a particular index
	// template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-simulate-template.html
	SimulateTemplate indices_simulate_template.NewSimulateTemplate
	// Split an index.
	// Split an index into a new index with more primary shards.
	// * Before you can split an index:
	//
	// * The index must be read-only.
	// * The cluster health status must be green.
	//
	// You can do make an index read-only with the following request using the add
	// index block API:
	//
	// ```
	// PUT /my_source_index/_block/write
	// ```
	//
	// The current write index on a data stream cannot be split.
	// In order to split the current write index, the data stream must first be
	// rolled over so that a new write index is created and then the previous write
	// index can be split.
	//
	// The number of times the index can be split (and the number of shards that
	// each original shard can be split into) is determined by the
	// `index.number_of_routing_shards` setting.
	// The number of routing shards specifies the hashing space that is used
	// internally to distribute documents across shards with consistent hashing.
	// For instance, a 5 shard index with `number_of_routing_shards` set to 30 (5 x
	// 2 x 3) could be split by a factor of 2 or 3.
	//
	// A split operation:
	//
	// * Creates a new target index with the same definition as the source index,
	// but with a larger number of primary shards.
	// * Hard-links segments from the source index into the target index. If the
	// file system doesn't support hard-linking, all segments are copied into the
	// new index, which is a much more time consuming process.
	// * Hashes all documents again, after low level files are created, to delete
	// documents that belong to a different shard.
	// * Recovers the target index as though it were a closed index which had just
	// been re-opened.
	//
	// IMPORTANT: Indices can only be split if they satisfy the following
	// requirements:
	//
	// * The target index must not exist.
	// * The source index must have fewer primary shards than the target index.
	// * The number of primary shards in the target index must be a multiple of the
	// number of primary shards in the source index.
	// * The node handling the split process must have sufficient free disk space to
	// accommodate a second copy of the existing index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-split-index.html
	Split indices_split.NewSplit
	// Get index statistics.
	// For data streams, the API retrieves statistics for the stream's backing
	// indices.
	//
	// By default, the returned statistics are index-level with `primaries` and
	// `total` aggregations.
	// `primaries` are the values for only the primary shards.
	// `total` are the accumulated values for both primary and replica shards.
	//
	// To get shard-level statistics, set the `level` parameter to `shards`.
	//
	// NOTE: When moving to another node, the shard-level statistics for a shard are
	// cleared.
	// Although the shard is no longer part of the node, that node retains any
	// node-level statistics to which the shard contributed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-stats.html
	Stats indices_stats.NewStats
	// Unfreeze an index.
	// When a frozen index is unfrozen, the index goes through the normal recovery
	// process and becomes writeable again.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/unfreeze-index-api.html
	Unfreeze indices_unfreeze.NewUnfreeze
	// Create or update an alias.
	// Adds a data stream or index to an alias.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-update-aliases
	UpdateAliases indices_update_aliases.NewUpdateAliases
	// Validate a query.
	// Validates a query without running it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-validate.html
	ValidateQuery indices_validate_query.NewValidateQuery
}

type Inference struct {
	// Perform chat completion inference
	//
	// The chat completion inference API enables real-time responses for chat
	// completion tasks by delivering answers incrementally, reducing response times
	// during computation.
	// It only works with the `chat_completion` task type for `openai`, `elastic`
	// and `googlevertexai` inference services.
	//
	// NOTE: The `chat_completion` task type is only available within the _stream
	// API and only supports streaming.
	// The Chat completion inference API and the Stream inference API differ in
	// their response structure and capabilities.
	// The Chat completion inference API provides more comprehensive customization
	// options through more fields and function calling support.
	// If you use the `openai` service or the `elastic` service, use the Chat
	// completion inference API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/chat-completion-inference-api.html
	ChatCompletionUnified inference_chat_completion_unified.NewChatCompletionUnified
	// Perform completion inference on the service
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html
	Completion inference_completion.NewCompletion
	// Delete an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-inference-api.html
	Delete inference_delete.NewDelete
	// Get an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-inference-api.html
	Get inference_get.NewGet
	// Perform inference on the service.
	//
	// This API enables you to use machine learning models to perform specific tasks
	// on data that you provide as an input.
	// It returns a response with the results of the tasks.
	// The inference endpoint you use can perform one specific task that has been
	// defined when the endpoint was created with the create inference API.
	//
	// For details about using this API with a service, such as Amazon Bedrock,
	// Anthropic, or HuggingFace, refer to the service-specific documentation.
	//
	// > info
	// > The inference APIs enable you to use certain services, such as built-in
	// machine learning models (ELSER, E5), models uploaded through Eland, Cohere,
	// OpenAI, Azure, Google AI Studio, Google Vertex AI, Anthropic, Watsonx.ai, or
	// Hugging Face. For built-in models and models uploaded through Eland, the
	// inference APIs offer an alternative way to use and manage trained models.
	// However, if you do not plan to use the inference APIs to use these models or
	// if you want to use non-NLP models, use the machine learning trained model
	// APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html
	Inference inference_inference.NewInference
	// Create an inference endpoint.
	//
	// IMPORTANT: The inference APIs enable you to use certain services, such as
	// built-in machine learning models (ELSER, E5), models uploaded through Eland,
	// Cohere, OpenAI, Mistral, Azure OpenAI, Google AI Studio, Google Vertex AI,
	// Anthropic, Watsonx.ai, or Hugging Face.
	// For built-in models and models uploaded through Eland, the inference APIs
	// offer an alternative way to use and manage trained models.
	// However, if you do not plan to use the inference APIs to use these models or
	// if you want to use non-NLP models, use the machine learning trained model
	// APIs.
	//
	// The following integrations are available through the inference API. You can
	// find the available task types next to the integration name:
	// * AlibabaCloud AI Search (`completion`, `rerank`, `sparse_embedding`,
	// `text_embedding`)
	// * Amazon Bedrock (`completion`, `text_embedding`)
	// * Anthropic (`completion`)
	// * Azure AI Studio (`completion`, `text_embedding`)
	// * Azure OpenAI (`completion`, `text_embedding`)
	// * Cohere (`completion`, `rerank`, `text_embedding`)
	// * Elasticsearch (`rerank`, `sparse_embedding`, `text_embedding` - this
	// service is for built-in models and models uploaded through Eland)
	// * ELSER (`sparse_embedding`)
	// * Google AI Studio (`completion`, `text_embedding`)
	// * Google Vertex AI (`rerank`, `text_embedding`)
	// * Hugging Face (`text_embedding`)
	// * Mistral (`text_embedding`)
	// * OpenAI (`chat_completion`, `completion`, `text_embedding`)
	// * VoyageAI (`text_embedding`, `rerank`)
	// * Watsonx inference integration (`text_embedding`)
	// * JinaAI (`text_embedding`, `rerank`)
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-inference-api.html
	Put inference_put.NewPut
	// Create an AlibabaCloud AI Search inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `alibabacloud-ai-search` service.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-alibabacloud-ai-search.html
	PutAlibabacloud inference_put_alibabacloud.NewPutAlibabacloud
	// Create an Amazon Bedrock inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `amazonbedrock` service.
	//
	// >info
	// > You need to provide the access and secret keys only once, during the
	// inference model creation. The get inference API does not retrieve your access
	// or secret keys. After creating the inference model, you cannot change the
	// associated key pairs. If you want to use a different access and secret key
	// pair, delete the inference model and recreate it with the same name and the
	// updated keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-amazon-bedrock.html
	PutAmazonbedrock inference_put_amazonbedrock.NewPutAmazonbedrock
	// Create an Anthropic inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `anthropic` service.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-anthropic.html
	PutAnthropic inference_put_anthropic.NewPutAnthropic
	// Create an Azure AI studio inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `azureaistudio` service.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-azure-ai-studio.html
	PutAzureaistudio inference_put_azureaistudio.NewPutAzureaistudio
	// Create an Azure OpenAI inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `azureopenai` service.
	//
	// The list of chat completion models that you can choose from in your Azure
	// OpenAI deployment include:
	//
	// * [GPT-4 and GPT-4 Turbo
	// models](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#gpt-4-and-gpt-4-turbo-models)
	// *
	// [GPT-3.5](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#gpt-35)
	//
	// The list of embeddings models that you can choose from in your deployment can
	// be found in the [Azure models
	// documentation](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#embeddings).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-azure-openai.html
	PutAzureopenai inference_put_azureopenai.NewPutAzureopenai
	// Create a Cohere inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the `cohere`
	// service.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-cohere.html
	PutCohere inference_put_cohere.NewPutCohere
	// Create an Elasticsearch inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `elasticsearch` service.
	//
	// > info
	// > Your Elasticsearch deployment contains preconfigured ELSER and E5 inference
	// endpoints, you only need to create the enpoints using the API if you want to
	// customize the settings.
	//
	// If you use the ELSER or the E5 model through the `elasticsearch` service, the
	// API request will automatically download and deploy the model if it isn't
	// downloaded yet.
	//
	// > info
	// > You might see a 502 bad gateway error in the response when using the Kibana
	// Console. This error usually just reflects a timeout, while the model
	// downloads in the background. You can check the download progress in the
	// Machine Learning UI. If using the Python client, you can set the timeout
	// parameter to a higher value.
	//
	// After creating the endpoint, wait for the model deployment to complete before
	// using it.
	// To verify the deployment status, use the get trained model statistics API.
	// Look for `"state": "fully_allocated"` in the response and ensure that the
	// `"allocation_count"` matches the `"target_allocation_count"`.
	// Avoid creating multiple endpoints for the same model unless required, as each
	// endpoint consumes significant resources.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-elasticsearch.html
	PutElasticsearch inference_put_elasticsearch.NewPutElasticsearch
	// Create an ELSER inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the `elser`
	// service.
	// You can also deploy ELSER by using the Elasticsearch inference integration.
	//
	// > info
	// > Your Elasticsearch deployment contains a preconfigured ELSER inference
	// endpoint, you only need to create the enpoint using the API if you want to
	// customize the settings.
	//
	// The API request will automatically download and deploy the ELSER model if it
	// isn't already downloaded.
	//
	// > info
	// > You might see a 502 bad gateway error in the response when using the Kibana
	// Console. This error usually just reflects a timeout, while the model
	// downloads in the background. You can check the download progress in the
	// Machine Learning UI. If using the Python client, you can set the timeout
	// parameter to a higher value.
	//
	// After creating the endpoint, wait for the model deployment to complete before
	// using it.
	// To verify the deployment status, use the get trained model statistics API.
	// Look for `"state": "fully_allocated"` in the response and ensure that the
	// `"allocation_count"` matches the `"target_allocation_count"`.
	// Avoid creating multiple endpoints for the same model unless required, as each
	// endpoint consumes significant resources.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-elser.html
	PutElser inference_put_elser.NewPutElser
	// Create an Google AI Studio inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `googleaistudio` service.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-google-ai-studio.html
	PutGoogleaistudio inference_put_googleaistudio.NewPutGoogleaistudio
	// Create a Google Vertex AI inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `googlevertexai` service.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-google-vertex-ai.html
	PutGooglevertexai inference_put_googlevertexai.NewPutGooglevertexai
	// Create a Hugging Face inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `hugging_face` service.
	//
	// You must first create an inference endpoint on the Hugging Face endpoint page
	// to get an endpoint URL.
	// Select the model you want to use on the new endpoint creation page (for
	// example `intfloat/e5-small-v2`), then select the sentence embeddings task
	// under the advanced configuration section.
	// Create the endpoint and copy the URL after the endpoint initialization has
	// been finished.
	//
	// The following models are recommended for the Hugging Face service:
	//
	// * `all-MiniLM-L6-v2`
	// * `all-MiniLM-L12-v2`
	// * `all-mpnet-base-v2`
	// * `e5-base-v2`
	// * `e5-small-v2`
	// * `multilingual-e5-base`
	// * `multilingual-e5-small`
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-hugging-face.html
	PutHuggingFace inference_put_hugging_face.NewPutHuggingFace
	// Create an JinaAI inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the `jinaai`
	// service.
	//
	// To review the available `rerank` models, refer to <https://jina.ai/reranker>.
	// To review the available `text_embedding` models, refer to the
	// <https://jina.ai/embeddings/>.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-jinaai.html
	PutJinaai inference_put_jinaai.NewPutJinaai
	// Create a Mistral inference endpoint.
	//
	// Creates an inference endpoint to perform an inference task with the `mistral`
	// service.
	// https://www.elastic.co/guide/en/elasticsearch/reference/{brnach}/infer-service-mistral.html
	PutMistral inference_put_mistral.NewPutMistral
	// Create an OpenAI inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the `openai`
	// service or `openai` compatible APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-openai.html
	PutOpenai inference_put_openai.NewPutOpenai
	// Create a VoyageAI inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the `voyageai`
	// service.
	//
	// Avoid creating multiple endpoints for the same model unless required, as each
	// endpoint consumes significant resources.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-voyageai.html
	PutVoyageai inference_put_voyageai.NewPutVoyageai
	// Create a Watsonx inference endpoint.
	//
	// Create an inference endpoint to perform an inference task with the
	// `watsonxai` service.
	// You need an IBM Cloud Databases for Elasticsearch deployment to use the
	// `watsonxai` inference service.
	// You can provision one through the IBM catalog, the Cloud Databases CLI
	// plug-in, the Cloud Databases API, or Terraform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-watsonx-ai.html
	PutWatsonx inference_put_watsonx.NewPutWatsonx
	// Perform rereanking inference on the service
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html
	Rerank inference_rerank.NewRerank
	// Perform sparse embedding inference on the service
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html
	SparseEmbedding inference_sparse_embedding.NewSparseEmbedding
	// Perform streaming inference.
	// Get real-time responses for completion tasks by delivering answers
	// incrementally, reducing response times during computation.
	// This API works only with the completion task type.
	//
	// IMPORTANT: The inference APIs enable you to use certain services, such as
	// built-in machine learning models (ELSER, E5), models uploaded through Eland,
	// Cohere, OpenAI, Azure, Google AI Studio, Google Vertex AI, Anthropic,
	// Watsonx.ai, or Hugging Face. For built-in models and models uploaded through
	// Eland, the inference APIs offer an alternative way to use and manage trained
	// models. However, if you do not plan to use the inference APIs to use these
	// models or if you want to use non-NLP models, use the machine learning trained
	// model APIs.
	//
	// This API requires the `monitor_inference` cluster privilege (the built-in
	// `inference_admin` and `inference_user` roles grant this privilege). You must
	// use a client that supports streaming.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stream-inference-api.html
	StreamCompletion inference_stream_completion.NewStreamCompletion
	// Perform text embedding inference on the service
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html
	TextEmbedding inference_text_embedding.NewTextEmbedding
	// Update an inference endpoint.
	//
	// Modify `task_settings`, secrets (within `service_settings`), or
	// `num_allocations` for an inference endpoint, depending on the specific
	// endpoint service and `task_type`.
	//
	// IMPORTANT: The inference APIs enable you to use certain services, such as
	// built-in machine learning models (ELSER, E5), models uploaded through Eland,
	// Cohere, OpenAI, Azure, Google AI Studio, Google Vertex AI, Anthropic,
	// Watsonx.ai, or Hugging Face.
	// For built-in models and models uploaded through Eland, the inference APIs
	// offer an alternative way to use and manage trained models.
	// However, if you do not plan to use the inference APIs to use these models or
	// if you want to use non-NLP models, use the machine learning trained model
	// APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-inference-api.html
	Update inference_update.NewUpdate
}

type Ingest struct {
	// Delete GeoIP database configurations.
	//
	// Delete one or more IP geolocation database configurations.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-ingest-delete-geoip-database
	DeleteGeoipDatabase ingest_delete_geoip_database.NewDeleteGeoipDatabase
	// Delete IP geolocation database configurations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-ip-location-database-api.html
	DeleteIpLocationDatabase ingest_delete_ip_location_database.NewDeleteIpLocationDatabase
	// Delete pipelines.
	// Delete one or more ingest pipelines.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-pipeline-api.html
	DeletePipeline ingest_delete_pipeline.NewDeletePipeline
	// Get GeoIP statistics.
	// Get download statistics for GeoIP2 databases that are used with the GeoIP
	// processor.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html
	GeoIpStats ingest_geo_ip_stats.NewGeoIpStats
	// Get GeoIP database configurations.
	//
	// Get information about one or more IP geolocation database configurations.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-ingest-get-geoip-database
	GetGeoipDatabase ingest_get_geoip_database.NewGetGeoipDatabase
	// Get IP geolocation database configurations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ip-location-database-api.html
	GetIpLocationDatabase ingest_get_ip_location_database.NewGetIpLocationDatabase
	// Get pipelines.
	//
	// Get information about one or more ingest pipelines.
	// This API returns a local reference of the pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-pipeline-api.html
	GetPipeline ingest_get_pipeline.NewGetPipeline
	// Run a grok processor.
	// Extract structured fields out of a single text field within a document.
	// You must choose which field to extract matched fields from, as well as the
	// grok pattern you expect will match.
	// A grok pattern is like a regular expression that supports aliased expressions
	// that can be reused.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html
	ProcessorGrok ingest_processor_grok.NewProcessorGrok
	// Create or update a GeoIP database configuration.
	//
	// Refer to the create or update IP geolocation database configuration API.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-ingest-put-geoip-database
	PutGeoipDatabase ingest_put_geoip_database.NewPutGeoipDatabase
	// Create or update an IP geolocation database configuration.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-ip-location-database-api.html
	PutIpLocationDatabase ingest_put_ip_location_database.NewPutIpLocationDatabase
	// Create or update a pipeline.
	// Changes made using this API take effect immediately.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html
	PutPipeline ingest_put_pipeline.NewPutPipeline
	// Simulate a pipeline.
	//
	// Run an ingest pipeline against a set of provided documents.
	// You can either specify an existing pipeline to use with the provided
	// documents or supply a pipeline definition in the body of the request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-pipeline-api.html
	Simulate ingest_simulate.NewSimulate
}

type License struct {
	// Delete the license.
	//
	// When the license expires, your subscription level reverts to Basic.
	//
	// If the operator privileges feature is enabled, only operator users can use
	// this API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-license.html
	Delete license_delete.NewDelete
	// Get license information.
	//
	// Get information about your Elastic license including its type, its status,
	// when it was issued, and when it expires.
	//
	// >info
	// > If the master node is generating a new cluster state, the get license API
	// may return a `404 Not Found` response.
	// > If you receive an unexpected 404 response after cluster startup, wait a
	// short period and retry the request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-license.html
	Get license_get.NewGet
	// Get the basic license status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-basic-status.html
	GetBasicStatus license_get_basic_status.NewGetBasicStatus
	// Get the trial status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trial-status.html
	GetTrialStatus license_get_trial_status.NewGetTrialStatus
	// Update the license.
	//
	// You can update your license at runtime without shutting down your nodes.
	// License updates take effect immediately.
	// If the license you are installing does not support all of the features that
	// were available with your previous license, however, you are notified in the
	// response.
	// You must then re-submit the API request with the acknowledge parameter set to
	// true.
	//
	// NOTE: If Elasticsearch security features are enabled and you are installing a
	// gold or higher license, you must enable TLS on the transport networking layer
	// before you install the license.
	// If the operator privileges feature is enabled, only operator users can use
	// this API.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-license-post
	Post license_post.NewPost
	// Start a basic license.
	//
	// Start an indefinite basic license, which gives access to all the basic
	// features.
	//
	// NOTE: In order to start a basic license, you must not currently have a basic
	// license.
	//
	// If the basic license does not support all of the features that are available
	// with your current license, however, you are notified in the response.
	// You must then re-submit the API request with the `acknowledge` parameter set
	// to `true`.
	//
	// To check the status of your basic license, use the get basic license API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-basic.html
	PostStartBasic license_post_start_basic.NewPostStartBasic
	// Start a trial.
	// Start a 30-day trial, which gives access to all subscription features.
	//
	// NOTE: You are allowed to start a trial only if your cluster has not already
	// activated a trial for the current major product version.
	// For example, if you have already activated a trial for v8.0, you cannot start
	// a new trial until v9.0. You can, however, request an extended trial at
	// https://www.elastic.co/trialextension.
	//
	// To check the status of your trial, use the get trial status API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trial.html
	PostStartTrial license_post_start_trial.NewPostStartTrial
}

type Logstash struct {
	// Delete a Logstash pipeline.
	// Delete a pipeline that is used for Logstash Central Management.
	// If the request succeeds, you receive an empty response with an appropriate
	// status code.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-delete-pipeline.html
	DeletePipeline logstash_delete_pipeline.NewDeletePipeline
	// Get Logstash pipelines.
	// Get pipelines that are used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-get-pipeline.html
	GetPipeline logstash_get_pipeline.NewGetPipeline
	// Create or update a Logstash pipeline.
	//
	// Create a pipeline that is used for Logstash Central Management.
	// If the specified pipeline exists, it is replaced.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-put-pipeline.html
	PutPipeline logstash_put_pipeline.NewPutPipeline
}

type Migration struct {
	// Get deprecation information.
	// Get information about different cluster, node, and index level settings that
	// use deprecated features that will be removed or changed in the next major
	// version.
	//
	// TIP: This APIs is designed for indirect use by the Upgrade Assistant.
	// You are strongly recommended to use the Upgrade Assistant.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migration-api-deprecation.html
	Deprecations migration_deprecations.NewDeprecations
	// Get feature migration information.
	// Version upgrades sometimes require changes to how features store
	// configuration information and data in system indices.
	// Check which features need to be migrated and the status of any migrations
	// that are in progress.
	//
	// TIP: This API is designed for indirect use by the Upgrade Assistant.
	// You are strongly recommended to use the Upgrade Assistant.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/feature-migration-api.html
	GetFeatureUpgradeStatus migration_get_feature_upgrade_status.NewGetFeatureUpgradeStatus
	// Start the feature migration.
	// Version upgrades sometimes require changes to how features store
	// configuration information and data in system indices.
	// This API starts the automatic migration process.
	//
	// Some functionality might be temporarily unavailable during the migration
	// process.
	//
	// TIP: The API is designed for indirect use by the Upgrade Assistant. We
	// strongly recommend you use the Upgrade Assistant.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/feature-migration-api.html
	PostFeatureUpgrade migration_post_feature_upgrade.NewPostFeatureUpgrade
}

type Ml struct {
	// Clear trained model deployment cache.
	//
	// Cache will be cleared on all nodes where the trained model is assigned.
	// A trained model deployment may have an inference cache enabled.
	// As requests are handled by each allocated node, their responses may be cached
	// on that individual node.
	// Calling this API clears the caches without restarting the deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-trained-model-deployment-cache.html
	ClearTrainedModelDeploymentCache ml_clear_trained_model_deployment_cache.NewClearTrainedModelDeploymentCache
	// Close anomaly detection jobs.
	//
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
	// Delete a calendar.
	//
	// Remove all scheduled events from a calendar, then delete it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar.html
	DeleteCalendar ml_delete_calendar.NewDeleteCalendar
	// Delete events from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-event.html
	DeleteCalendarEvent ml_delete_calendar_event.NewDeleteCalendarEvent
	// Delete anomaly jobs from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-job.html
	DeleteCalendarJob ml_delete_calendar_job.NewDeleteCalendarJob
	// Delete a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-dfanalytics.html
	DeleteDataFrameAnalytics ml_delete_data_frame_analytics.NewDeleteDataFrameAnalytics
	// Delete a datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-datafeed.html
	DeleteDatafeed ml_delete_datafeed.NewDeleteDatafeed
	// Delete expired ML data.
	//
	// Delete all job results, model snapshots and forecast data that have exceeded
	// their retention days period. Machine learning state documents that are not
	// associated with any job are also deleted.
	// You can limit the request to a single or set of anomaly detection jobs by
	// using a job identifier, a group name, a comma-separated list of jobs, or a
	// wildcard expression. You can delete expired data for all anomaly detection
	// jobs by using `_all`, by specifying `*` as the `<job_id>`, or by omitting the
	// `<job_id>`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-expired-data.html
	DeleteExpiredData ml_delete_expired_data.NewDeleteExpiredData
	// Delete a filter.
	//
	// If an anomaly detection job references the filter, you cannot delete the
	// filter. You must update or delete the job before you can delete the filter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-filter.html
	DeleteFilter ml_delete_filter.NewDeleteFilter
	// Delete forecasts from a job.
	//
	// By default, forecasts are retained for 14 days. You can specify a
	// different retention period with the `expires_in` parameter in the forecast
	// jobs API. The delete forecast API enables you to delete one or more
	// forecasts before they expire.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-forecast.html
	DeleteForecast ml_delete_forecast.NewDeleteForecast
	// Delete an anomaly detection job.
	//
	// All job configuration, model state and results are deleted.
	// It is not currently possible to delete multiple jobs using wildcards or a
	// comma separated list. If you delete a job that has a datafeed, the request
	// first tries to delete the datafeed. This behavior is equivalent to calling
	// the delete datafeed API with the same timeout and force parameters as the
	// delete job request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-job.html
	DeleteJob ml_delete_job.NewDeleteJob
	// Delete a model snapshot.
	//
	// You cannot delete the active model snapshot. To delete that snapshot, first
	// revert to a different one. To identify the active model snapshot, refer to
	// the `model_snapshot_id` in the results from the get jobs API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-snapshot.html
	DeleteModelSnapshot ml_delete_model_snapshot.NewDeleteModelSnapshot
	// Delete an unreferenced trained model.
	//
	// The request deletes a trained inference model that is not referenced by an
	// ingest pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models.html
	DeleteTrainedModel ml_delete_trained_model.NewDeleteTrainedModel
	// Delete a trained model alias.
	//
	// This API deletes an existing model alias that refers to a trained model. If
	// the model alias is missing or refers to a model other than the one identified
	// by the `model_id`, this API returns an error.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models-aliases.html
	DeleteTrainedModelAlias ml_delete_trained_model_alias.NewDeleteTrainedModelAlias
	// Estimate job model memory usage.
	//
	// Make an estimation of the memory usage for an anomaly detection job model.
	// The estimate is based on analysis configuration details for the job and
	// cardinality
	// estimates for the fields it references.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-estimate-model-memory.html
	EstimateModelMemory ml_estimate_model_memory.NewEstimateModelMemory
	// Evaluate data frame analytics.
	//
	// The API packages together commonly used evaluation metrics for various types
	// of machine learning features. This has been designed for use on indexes
	// created by data frame analytics. Evaluation requires both a ground truth
	// field and an analytics result field to be present.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/evaluate-dfanalytics.html
	EvaluateDataFrame ml_evaluate_data_frame.NewEvaluateDataFrame
	// Explain data frame analytics config.
	//
	// This API provides explanations for a data frame analytics config that either
	// exists already or one that has not been created yet. The following
	// explanations are provided:
	// * which fields are included or not in the analysis and why,
	// * how much memory is estimated to be required. The estimate can be used when
	// deciding the appropriate value for model_memory_limit setting later on.
	// If you have object fields or fields that are excluded via source filtering,
	// they are not included in the explanation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/explain-dfanalytics.html
	ExplainDataFrameAnalytics ml_explain_data_frame_analytics.NewExplainDataFrameAnalytics
	// Force buffered data to be processed.
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
	// Predict future behavior of a time series.
	//
	// Forecasts are not supported for jobs that perform population analysis; an
	// error occurs if you try to create a forecast for a job that has an
	// `over_field_name` in its configuration. Forcasts predict future behavior
	// based on historical data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-forecast.html
	Forecast ml_forecast.NewForecast
	// Get anomaly detection job results for buckets.
	// The API presents a chronological view of the records, grouped by bucket.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-bucket.html
	GetBuckets ml_get_buckets.NewGetBuckets
	// Get info about events in calendars.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar-event.html
	GetCalendarEvents ml_get_calendar_events.NewGetCalendarEvents
	// Get calendar configuration info.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar.html
	GetCalendars ml_get_calendars.NewGetCalendars
	// Get anomaly detection job results for categories.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-category.html
	GetCategories ml_get_categories.NewGetCategories
	// Get data frame analytics job configuration info.
	// You can get information for multiple data frame analytics jobs in a single
	// API request by using a comma-separated list of data frame analytics jobs or a
	// wildcard expression.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics.html
	GetDataFrameAnalytics ml_get_data_frame_analytics.NewGetDataFrameAnalytics
	// Get data frame analytics job stats.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics-stats.html
	GetDataFrameAnalyticsStats ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStats
	// Get datafeed stats.
	// You can get statistics for multiple datafeeds in a single API request by
	// using a comma-separated list of datafeeds or a wildcard expression. You can
	// get statistics for all datafeeds by using `_all`, by specifying `*` as the
	// `<feed_id>`, or by omitting the `<feed_id>`. If the datafeed is stopped, the
	// only information you receive is the `datafeed_id` and the `state`.
	// This API returns a maximum of 10,000 datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed-stats.html
	GetDatafeedStats ml_get_datafeed_stats.NewGetDatafeedStats
	// Get datafeeds configuration info.
	// You can get information for multiple datafeeds in a single API request by
	// using a comma-separated list of datafeeds or a wildcard expression. You can
	// get information for all datafeeds by using `_all`, by specifying `*` as the
	// `<feed_id>`, or by omitting the `<feed_id>`.
	// This API returns a maximum of 10,000 datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed.html
	GetDatafeeds ml_get_datafeeds.NewGetDatafeeds
	// Get filters.
	// You can get a single filter or all filters.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-filter.html
	GetFilters ml_get_filters.NewGetFilters
	// Get anomaly detection job results for influencers.
	// Influencers are the entities that have contributed to, or are to blame for,
	// the anomalies. Influencer results are available only if an
	// `influencer_field_name` is specified in the job configuration.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-influencer.html
	GetInfluencers ml_get_influencers.NewGetInfluencers
	// Get anomaly detection job stats.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-stats.html
	GetJobStats ml_get_job_stats.NewGetJobStats
	// Get anomaly detection jobs configuration info.
	// You can get information for multiple anomaly detection jobs in a single API
	// request by using a group name, a comma-separated list of jobs, or a wildcard
	// expression. You can get information for all anomaly detection jobs by using
	// `_all`, by specifying `*` as the `<job_id>`, or by omitting the `<job_id>`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job.html
	GetJobs ml_get_jobs.NewGetJobs
	// Get machine learning memory usage info.
	// Get information about how machine learning jobs and trained models are using
	// memory,
	// on each node, both within the JVM heap, and natively, outside of the JVM.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-memory.html
	GetMemoryStats ml_get_memory_stats.NewGetMemoryStats
	// Get anomaly detection job model snapshot upgrade usage info.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-model-snapshot-upgrade-stats.html
	GetModelSnapshotUpgradeStats ml_get_model_snapshot_upgrade_stats.NewGetModelSnapshotUpgradeStats
	// Get model snapshots info.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-snapshot.html
	GetModelSnapshots ml_get_model_snapshots.NewGetModelSnapshots
	// Get overall bucket results.
	//
	// Retrievs overall bucket results that summarize the bucket results of
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
	// Get anomaly records for an anomaly detection job.
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
	// Get trained model configuration info.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models.html
	GetTrainedModels ml_get_trained_models.NewGetTrainedModels
	// Get trained models usage info.
	// You can get usage information for multiple trained
	// models in a single API request by using a comma-separated list of model IDs
	// or a wildcard expression.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models-stats.html
	GetTrainedModelsStats ml_get_trained_models_stats.NewGetTrainedModelsStats
	// Evaluate a trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-trained-model.html
	InferTrainedModel ml_infer_trained_model.NewInferTrainedModel
	// Get machine learning information.
	// Get defaults and limits used by machine learning.
	// This endpoint is designed to be used by a user interface that needs to fully
	// understand machine learning configurations where some options are not
	// specified, meaning that the defaults should be used. This endpoint may be
	// used to find out what those defaults are. It also provides information about
	// the maximum size of machine learning jobs that could run in the current
	// cluster configuration.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-info.html
	Info ml_info.NewInfo
	// Open anomaly detection jobs.
	//
	// An anomaly detection job must be opened to be ready to receive and analyze
	// data. It can be opened and closed multiple times throughout its lifecycle.
	// When you open a new job, it starts with an empty model.
	// When you open an existing job, the most recent model state is automatically
	// loaded. The job is ready to resume its analysis from where it left off, once
	// new data is received.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-open-job.html
	OpenJob ml_open_job.NewOpenJob
	// Add scheduled events to the calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-calendar-event.html
	PostCalendarEvents ml_post_calendar_events.NewPostCalendarEvents
	// Send data to an anomaly detection job for analysis.
	//
	// IMPORTANT: For each job, data can be accepted from only a single connection
	// at a time.
	// It is not currently possible to post data to multiple jobs using wildcards or
	// a comma-separated list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-data.html
	PostData ml_post_data.NewPostData
	// Preview features used by data frame analytics.
	// Preview the extracted features used by a data frame analytics config.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-dfanalytics.html
	PreviewDataFrameAnalytics ml_preview_data_frame_analytics.NewPreviewDataFrameAnalytics
	// Preview a datafeed.
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
	// Create a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar.html
	PutCalendar ml_put_calendar.NewPutCalendar
	// Add anomaly detection job to calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar-job.html
	PutCalendarJob ml_put_calendar_job.NewPutCalendarJob
	// Create a data frame analytics job.
	// This API creates a data frame analytics job that performs an analysis on the
	// source indices and stores the outcome in a destination index.
	// By default, the query used in the source configuration is `{"match_all":
	// {}}`.
	//
	// If the destination index does not exist, it is created automatically when you
	// start the job.
	//
	// If you supply only a subset of the regression or classification parameters,
	// hyperparameter optimization occurs. It determines a value for each of the
	// undefined parameters.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-dfanalytics.html
	PutDataFrameAnalytics ml_put_data_frame_analytics.NewPutDataFrameAnalytics
	// Create a datafeed.
	// Datafeeds retrieve data from Elasticsearch for analysis by an anomaly
	// detection job.
	// You can associate only one datafeed with each anomaly detection job.
	// The datafeed contains a query that runs at a defined interval (`frequency`).
	// If you are concerned about delayed data, you can add a delay (`query_delay')
	// at each interval.
	// By default, the datafeed uses the following query: `{"match_all": {"boost":
	// 1}}`.
	//
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
	// Create a filter.
	// A filter contains a list of strings. It can be used by one or more anomaly
	// detection jobs.
	// Specifically, filters are referenced in the `custom_rules` property of
	// detector configuration objects.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-filter.html
	PutFilter ml_put_filter.NewPutFilter
	// Create an anomaly detection job.
	//
	// If you include a `datafeed_config`, you must have read index privileges on
	// the source index.
	// If you include a `datafeed_config` but do not provide a query, the datafeed
	// uses `{"match_all": {"boost": 1}}`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-job.html
	PutJob ml_put_job.NewPutJob
	// Create a trained model.
	// Enable you to supply a trained model that is not created by data frame
	// analytics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models.html
	PutTrainedModel ml_put_trained_model.NewPutTrainedModel
	// Create or update a trained model alias.
	// A trained model alias is a logical name used to reference a single trained
	// model.
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
	// Create part of a trained model definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-definition-part.html
	PutTrainedModelDefinitionPart ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPart
	// Create a trained model vocabulary.
	// This API is supported only for natural language processing (NLP) models.
	// The vocabulary is stored in the index as described in
	// `inference_config.*.vocabulary` of the trained model definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-vocabulary.html
	PutTrainedModelVocabulary ml_put_trained_model_vocabulary.NewPutTrainedModelVocabulary
	// Reset an anomaly detection job.
	// All model state and results are deleted. The job is ready to start over as if
	// it had just been created.
	// It is not currently possible to reset multiple jobs using wildcards or a
	// comma separated list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-reset-job.html
	ResetJob ml_reset_job.NewResetJob
	// Revert to a snapshot.
	// The machine learning features react quickly to anomalous input, learning new
	// behaviors in data. Highly anomalous input increases the variance in the
	// models whilst the system learns whether this is a new step-change in behavior
	// or a one-off event. In the case where this anomalous input is known to be a
	// one-off, then it might be appropriate to reset the model state to a time
	// before this event. For example, you might consider reverting to a saved
	// snapshot after Black Friday or a critical system failure.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-revert-snapshot.html
	RevertModelSnapshot ml_revert_model_snapshot.NewRevertModelSnapshot
	// Set upgrade_mode for ML indices.
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
	// Start a data frame analytics job.
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
	// Start datafeeds.
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
	// Start a trained model deployment.
	// It allocates the model to every machine learning node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trained-model-deployment.html
	StartTrainedModelDeployment ml_start_trained_model_deployment.NewStartTrainedModelDeployment
	// Stop data frame analytics jobs.
	// A data frame analytics job can be started and stopped multiple times
	// throughout its lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-dfanalytics.html
	StopDataFrameAnalytics ml_stop_data_frame_analytics.NewStopDataFrameAnalytics
	// Stop datafeeds.
	// A datafeed that is stopped ceases to retrieve data from Elasticsearch. A
	// datafeed can be started and stopped
	// multiple times throughout its lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-stop-datafeed.html
	StopDatafeed ml_stop_datafeed.NewStopDatafeed
	// Stop a trained model deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-trained-model-deployment.html
	StopTrainedModelDeployment ml_stop_trained_model_deployment.NewStopTrainedModelDeployment
	// Update a data frame analytics job.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-ml-update-data-frame-analytics
	UpdateDataFrameAnalytics ml_update_data_frame_analytics.NewUpdateDataFrameAnalytics
	// Update a datafeed.
	// You must stop and start the datafeed for the changes to be applied.
	// When Elasticsearch security features are enabled, your datafeed remembers
	// which roles the user who updated it had at
	// the time of the update and runs the query using those same roles. If you
	// provide secondary authorization headers,
	// those credentials are used instead.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-datafeed.html
	UpdateDatafeed ml_update_datafeed.NewUpdateDatafeed
	// Update a filter.
	// Updates the description of a filter, adds items, or removes items from the
	// list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-filter.html
	UpdateFilter ml_update_filter.NewUpdateFilter
	// Update an anomaly detection job.
	// Updates certain properties of an anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-job.html
	UpdateJob ml_update_job.NewUpdateJob
	// Update a snapshot.
	// Updates certain properties of a snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-snapshot.html
	UpdateModelSnapshot ml_update_model_snapshot.NewUpdateModelSnapshot
	// Update a trained model deployment.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-ml-update-trained-model-deployment
	UpdateTrainedModelDeployment ml_update_trained_model_deployment.NewUpdateTrainedModelDeployment
	// Upgrade a snapshot.
	// Upgrade an anomaly detection model snapshot to the latest major version.
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
	// Validate an anomaly detection job.
	// https://www.elastic.co/guide/en/machine-learning/current/ml-jobs.html
	Validate ml_validate.NewValidate
	// Validate an anomaly detection job.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8
	ValidateDetector ml_validate_detector.NewValidateDetector
}

type Monitoring struct {
	// Send monitoring data.
	// This API is used by the monitoring features to send monitoring data.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8
	Bulk monitoring_bulk.NewBulk
}

type Nodes struct {
	// Clear the archived repositories metering.
	// Clear the archived repositories metering information in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-repositories-metering-archive-api.html
	ClearRepositoriesMeteringArchive nodes_clear_repositories_metering_archive.NewClearRepositoriesMeteringArchive
	// Get cluster repositories metering.
	// Get repositories metering information for a cluster.
	// This API exposes monotonically non-decreasing counters and it is expected
	// that clients would durably store the information needed to compute
	// aggregations over a period of time.
	// Additionally, the information exposed by this API is volatile, meaning that
	// it will not be present after node restarts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-repositories-metering-api.html
	GetRepositoriesMeteringInfo nodes_get_repositories_metering_info.NewGetRepositoriesMeteringInfo
	// Get the hot threads for nodes.
	// Get a breakdown of the hot threads on each selected node in the cluster.
	// The output is plain text with a breakdown of the top hot threads for each
	// node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-hot-threads.html
	HotThreads nodes_hot_threads.NewHotThreads
	// Get node information.
	//
	// By default, the API returns all attributes and core settings for cluster
	// nodes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-info.html
	Info nodes_info.NewInfo
	// Reload the keystore on nodes in the cluster.
	//
	// Secure settings are stored in an on-disk keystore. Certain of these settings
	// are reloadable.
	// That is, you can change them on disk and reload them without restarting any
	// nodes in the cluster.
	// When you have updated reloadable secure settings in your keystore, you can
	// use this API to reload those settings on each node.
	//
	// When the Elasticsearch keystore is password protected and not simply
	// obfuscated, you must provide the password for the keystore when you reload
	// the secure settings.
	// Reloading the settings for the whole cluster assumes that the keystores for
	// all nodes are protected with the same password; this method is allowed only
	// when inter-node communications are encrypted.
	// Alternatively, you can reload the secure settings on each node by locally
	// accessing the API and passing the node-specific Elasticsearch keystore
	// password.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-reload-secure-settings.html
	ReloadSecureSettings nodes_reload_secure_settings.NewReloadSecureSettings
	// Get node statistics.
	// Get statistics for nodes in a cluster.
	// By default, all stats are returned. You can limit the returned information by
	// using metrics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-stats.html
	Stats nodes_stats.NewStats
	// Get feature usage information.
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
	// Delete a query rule.
	// Delete a query rule within a query ruleset.
	// This is a destructive action that is only recoverable by re-adding the same
	// rule with the create or update query rule API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-rule.html
	DeleteRule query_rules_delete_rule.NewDeleteRule
	// Delete a query ruleset.
	// Remove a query ruleset and its associated data.
	// This is a destructive action that is not recoverable.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-ruleset.html
	DeleteRuleset query_rules_delete_ruleset.NewDeleteRuleset
	// Get a query rule.
	// Get details about a query rule within a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-rule.html
	GetRule query_rules_get_rule.NewGetRule
	// Get a query ruleset.
	// Get details about a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-ruleset.html
	GetRuleset query_rules_get_ruleset.NewGetRuleset
	// Get all query rulesets.
	// Get summarized information about the query rulesets.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-query-rulesets.html
	ListRulesets query_rules_list_rulesets.NewListRulesets
	// Create or update a query rule.
	// Create or update a query rule within a query ruleset.
	//
	// IMPORTANT: Due to limitations within pinned queries, you can only pin
	// documents using ids or docs, but cannot use both in single rule.
	// It is advised to use one or the other in query rulesets, to avoid errors.
	// Additionally, pinned queries have a maximum limit of 100 pinned hits.
	// If multiple matching rules pin more than 100 documents, only the first 100
	// documents are pinned in the order they are specified in the ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-rule.html
	PutRule query_rules_put_rule.NewPutRule
	// Create or update a query ruleset.
	// There is a limit of 100 rules per ruleset.
	// This limit can be increased by using the
	// `xpack.applications.rules.max_rules_per_ruleset` cluster setting.
	//
	// IMPORTANT: Due to limitations within pinned queries, you can only select
	// documents using `ids` or `docs`, but cannot use both in single rule.
	// It is advised to use one or the other in query rulesets, to avoid errors.
	// Additionally, pinned queries have a maximum limit of 100 pinned hits.
	// If multiple matching rules pin more than 100 documents, only the first 100
	// documents are pinned in the order they are specified in the ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-ruleset.html
	PutRuleset query_rules_put_ruleset.NewPutRuleset
	// Test a query ruleset.
	// Evaluate match criteria against a query ruleset to identify the rules that
	// would match that criteria.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/test-query-ruleset.html
	Test query_rules_test.NewTest
}

type Rollup struct {
	// Delete a rollup job.
	//
	// A job must be stopped before it can be deleted.
	// If you attempt to delete a started job, an error occurs.
	// Similarly, if you attempt to delete a nonexistent job, an exception occurs.
	//
	// IMPORTANT: When you delete a job, you remove only the process that is
	// actively monitoring and rolling up data.
	// The API does not delete any previously rolled up data.
	// This is by design; a user may wish to roll up a static data set.
	// Because the data set is static, after it has been fully rolled up there is no
	// need to keep the indexing rollup job around (as there will be no new data).
	// Thus the job can be deleted, leaving behind the rolled up data for analysis.
	// If you wish to also remove the rollup data and the rollup index contains the
	// data for only a single job, you can delete the whole rollup index.
	// If the rollup index stores data from several jobs, you must issue a
	// delete-by-query that targets the rollup job's identifier in the rollup index.
	// For example:
	//
	// ```
	// POST my_rollup_index/_delete_by_query
	// {
	//   "query": {
	//     "term": {
	//       "_rollup.id": "the_rollup_job_id"
	//     }
	//   }
	// }
	// ```
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-delete-job.html
	DeleteJob rollup_delete_job.NewDeleteJob
	// Get rollup job information.
	// Get the configuration, stats, and status of rollup jobs.
	//
	// NOTE: This API returns only active (both `STARTED` and `STOPPED`) jobs.
	// If a job was created, ran for a while, then was deleted, the API does not
	// return any details about it.
	// For details about a historical rollup job, the rollup capabilities API may be
	// more useful.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-job.html
	GetJobs rollup_get_jobs.NewGetJobs
	// Get the rollup job capabilities.
	// Get the capabilities of any rollup jobs that have been configured for a
	// specific index or index pattern.
	//
	// This API is useful because a rollup job is often configured to rollup only a
	// subset of fields from the source index.
	// Furthermore, only certain aggregations can be configured for various fields,
	// leading to a limited subset of functionality depending on that configuration.
	// This API enables you to inspect an index and determine:
	//
	// 1. Does this index have associated rollup data somewhere in the cluster?
	// 2. If yes to the first question, what fields were rolled up, what
	// aggregations can be performed, and where does the data live?
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-rollup-caps.html
	GetRollupCaps rollup_get_rollup_caps.NewGetRollupCaps
	// Get the rollup index capabilities.
	// Get the rollup capabilities of all jobs inside of a rollup index.
	// A single rollup index may store the data for multiple rollup jobs and may
	// have a variety of capabilities depending on those jobs. This API enables you
	// to determine:
	//
	// * What jobs are stored in an index (or indices specified via a pattern)?
	// * What target indices were rolled up, what fields were used in those rollups,
	// and what aggregations can be performed on each job?
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-rollup-index-caps.html
	GetRollupIndexCaps rollup_get_rollup_index_caps.NewGetRollupIndexCaps
	// Create a rollup job.
	//
	// WARNING: From 8.15.0, calling this API in a cluster with no rollup usage will
	// fail with a message about the deprecation and planned removal of rollup
	// features. A cluster needs to contain either a rollup job or a rollup index in
	// order for this API to be allowed to run.
	//
	// The rollup job configuration contains all the details about how the job
	// should run, when it indexes documents, and what future queries will be able
	// to run against the rollup index.
	//
	// There are three main sections to the job configuration: the logistical
	// details about the job (for example, the cron schedule), the fields that are
	// used for grouping, and what metrics to collect for each group.
	//
	// Jobs are created in a `STOPPED` state. You can start them with the start
	// rollup jobs API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-put-job.html
	PutJob rollup_put_job.NewPutJob
	// Search rolled-up data.
	// The rollup search endpoint is needed because, internally, rolled-up documents
	// utilize a different document structure than the original data.
	// It rewrites standard Query DSL into a format that matches the rollup
	// documents then takes the response and rewrites it back to what a client would
	// expect given the original query.
	//
	// The request body supports a subset of features from the regular search API.
	// The following functionality is not available:
	//
	// `size`: Because rollups work on pre-aggregated data, no search hits can be
	// returned and so size must be set to zero or omitted entirely.
	// `highlighter`, `suggestors`, `post_filter`, `profile`, `explain`: These are
	// similarly disallowed.
	//
	// **Searching both historical rollup and non-rollup data**
	//
	// The rollup search API has the capability to search across both "live"
	// non-rollup data and the aggregated rollup data.
	// This is done by simply adding the live indices to the URI. For example:
	//
	// ```
	// GET sensor-1,sensor_rollup/_rollup_search
	// {
	//   "size": 0,
	//   "aggregations": {
	//      "max_temperature": {
	//       "max": {
	//         "field": "temperature"
	//       }
	//     }
	//   }
	// }
	// ```
	//
	// The rollup search endpoint does two things when the search runs:
	//
	// * The original request is sent to the non-rollup index unaltered.
	// * A rewritten version of the original request is sent to the rollup index.
	//
	// When the two responses are received, the endpoint rewrites the rollup
	// response and merges the two together.
	// During the merging process, if there is any overlap in buckets between the
	// two responses, the buckets from the non-rollup index are used.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-search.html
	RollupSearch rollup_rollup_search.NewRollupSearch
	// Start rollup jobs.
	// If you try to start a job that does not exist, an exception occurs.
	// If you try to start a job that is already started, nothing happens.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-start-job.html
	StartJob rollup_start_job.NewStartJob
	// Stop rollup jobs.
	// If you try to stop a job that does not exist, an exception occurs.
	// If you try to stop a job that is already stopped, nothing happens.
	//
	// Since only a stopped job can be deleted, it can be useful to block the API
	// until the indexer has fully stopped.
	// This is accomplished with the `wait_for_completion` query parameter, and
	// optionally a timeout. For example:
	//
	// ```
	// POST _rollup/job/sensor/_stop?wait_for_completion=true&timeout=10s
	// ```
	// The parameter blocks the API call from returning until either the job has
	// moved to STOPPED or the specified time has elapsed.
	// If the specified time elapses without the job moving to STOPPED, a timeout
	// exception occurs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-stop-job.html
	StopJob rollup_stop_job.NewStopJob
}

type SearchApplication struct {
	// Delete a search application.
	//
	// Remove a search application and its associated alias. Indices attached to the
	// search application are not removed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-search-application.html
	Delete search_application_delete.NewDelete
	// Delete a behavioral analytics collection.
	// The associated data stream is also deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-analytics-collection.html
	DeleteBehavioralAnalytics search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalytics
	// Get search application details.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-search-application.html
	Get search_application_get.NewGet
	// Get behavioral analytics collections.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-analytics-collection.html
	GetBehavioralAnalytics search_application_get_behavioral_analytics.NewGetBehavioralAnalytics
	// Get search applications.
	// Get information about search applications.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-analytics-collection.html
	List search_application_list.NewList
	// Create a behavioral analytics collection event.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-analytics-collection-event.html
	PostBehavioralAnalyticsEvent search_application_post_behavioral_analytics_event.NewPostBehavioralAnalyticsEvent
	// Create or update a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-search-application.html
	Put search_application_put.NewPut
	// Create a behavioral analytics collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-analytics-collection.html
	PutBehavioralAnalytics search_application_put_behavioral_analytics.NewPutBehavioralAnalytics
	// Render a search application query.
	// Generate an Elasticsearch query using the specified query parameters and the
	// search template associated with the search application or a default template
	// if none is specified.
	// If a parameter used in the search template is not specified in `params`, the
	// parameter's default value will be used.
	// The API returns the specific Elasticsearch query that would be generated and
	// run by calling the search application search API.
	//
	// You must have `read` privileges on the backing alias of the search
	// application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-application-render-query.html
	RenderQuery search_application_render_query.NewRenderQuery
	// Run a search application search.
	// Generate and run an Elasticsearch query that uses the specified query
	// parameteter and the search template associated with the search application or
	// default template.
	// Unspecified template parameters are assigned their default values if
	// applicable.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-application-search.html
	Search search_application_search.NewSearch
}

type SearchableSnapshots struct {
	// Get cache statistics.
	// Get statistics about the shared cache for partially mounted indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-api-cache-stats.html
	CacheStats searchable_snapshots_cache_stats.NewCacheStats
	// Clear the cache.
	// Clear indices and data streams from the shared cache for partially mounted
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-api-clear-cache.html
	ClearCache searchable_snapshots_clear_cache.NewClearCache
	// Mount a snapshot.
	// Mount a snapshot as a searchable snapshot index.
	// Do not use this API for snapshots managed by index lifecycle management
	// (ILM).
	// Manually mounting ILM-managed snapshots can interfere with ILM processes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-api-mount-snapshot.html
	Mount searchable_snapshots_mount.NewMount
	// Get searchable snapshot statistics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-api-stats.html
	Stats searchable_snapshots_stats.NewStats
}

type Security struct {
	// Activate a user profile.
	//
	// Create or update a user profile on behalf of another user.
	//
	// NOTE: The user profile feature is designed only for use by Kibana and
	// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
	// Individual users and external applications should not call this API directly.
	// The calling application must have either an `access_token` or a combination
	// of `username` and `password` for the user that the profile document is
	// intended for.
	// Elastic reserves the right to change or remove this feature in future
	// releases without prior notice.
	//
	// This API creates or updates a profile document for end users with information
	// that is extracted from the user's authentication object including `username`,
	// `full_name,` `roles`, and the authentication realm.
	// For example, in the JWT `access_token` case, the profile user's `username` is
	// extracted from the JWT token claim pointed to by the `claims.principal`
	// setting of the JWT realm that authenticated the token.
	//
	// When updating a profile document, the API enables the document if it was
	// disabled.
	// Any updates do not change existing content for either the `labels` or `data`
	// fields.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-activate-user-profile.html
	ActivateUserProfile security_activate_user_profile.NewActivateUserProfile
	// Authenticate a user.
	//
	// Authenticates a user and returns information about the authenticated user.
	// Include the user information in a [basic auth
	// header](https://en.wikipedia.org/wiki/Basic_access_authentication).
	// A successful call returns a JSON structure that shows user information such
	// as their username, the roles that are assigned to the user, any assigned
	// metadata, and information about the realms that authenticated and authorized
	// the user.
	// If the user cannot be authenticated, this API returns a 401 status code.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-authenticate.html
	Authenticate security_authenticate.NewAuthenticate
	// Bulk delete roles.
	//
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The bulk delete roles API cannot delete roles that are defined in roles
	// files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-delete-role.html
	BulkDeleteRole security_bulk_delete_role.NewBulkDeleteRole
	// Bulk create or update roles.
	//
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The bulk create or update roles API cannot update roles that are defined in
	// roles files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-put-role.html
	BulkPutRole security_bulk_put_role.NewBulkPutRole
	// Bulk update API keys.
	// Update the attributes for multiple API keys.
	//
	// IMPORTANT: It is not possible to use an API key as the authentication
	// credential for this API. To update API keys, the owner user's credentials are
	// required.
	//
	// This API is similar to the update API key API but enables you to apply the
	// same update to multiple API keys in one API call. This operation can greatly
	// improve performance over making individual updates.
	//
	// It is not possible to update expired or invalidated API keys.
	//
	// This API supports updates to API key access scope, metadata and expiration.
	// The access scope of each API key is derived from the `role_descriptors` you
	// specify in the request and a snapshot of the owner user's permissions at the
	// time of the request.
	// The snapshot of the owner's permissions is updated automatically on every
	// call.
	//
	// IMPORTANT: If you don't specify `role_descriptors` in the request, a call to
	// this API might still change an API key's access scope. This change can occur
	// if the owner user's permissions have changed since the API key was created or
	// last modified.
	//
	// A successful request returns a JSON structure that contains the IDs of all
	// updated API keys, the IDs of API keys that already had the requested changes
	// and did not require an update, and error details for any failed update.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-update-api-keys.html
	BulkUpdateApiKeys security_bulk_update_api_keys.NewBulkUpdateApiKeys
	// Change passwords.
	//
	// Change the passwords of users in the native realm and built-in users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-change-password.html
	ChangePassword security_change_password.NewChangePassword
	// Clear the API key cache.
	//
	// Evict a subset of all entries from the API key cache.
	// The cache is also automatically cleared on state changes of the security
	// index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-api-key-cache.html
	ClearApiKeyCache security_clear_api_key_cache.NewClearApiKeyCache
	// Clear the privileges cache.
	//
	// Evict privileges from the native application privilege cache.
	// The cache is also automatically cleared for applications that have their
	// privileges updated.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-privilege-cache.html
	ClearCachedPrivileges security_clear_cached_privileges.NewClearCachedPrivileges
	// Clear the user cache.
	//
	// Evict users from the user cache.
	// You can completely clear the cache or evict specific users.
	//
	// User credentials are cached in memory on each node to avoid connecting to a
	// remote authentication service or hitting the disk for every incoming request.
	// There are realm settings that you can use to configure the user cache.
	// For more information, refer to the documentation about controlling the user
	// cache.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-cache.html
	ClearCachedRealms security_clear_cached_realms.NewClearCachedRealms
	// Clear the roles cache.
	//
	// Evict roles from the native role cache.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-role-cache.html
	ClearCachedRoles security_clear_cached_roles.NewClearCachedRoles
	// Clear service account token caches.
	//
	// Evict a subset of all entries from the service account token caches.
	// Two separate caches exist for service account tokens: one cache for tokens
	// backed by the `service_tokens` file, and another for tokens backed by the
	// `.security` index.
	// This API clears matching entries from both caches.
	//
	// The cache for service account tokens backed by the `.security` index is
	// cleared automatically on state changes of the security index.
	// The cache for tokens backed by the `service_tokens` file is cleared
	// automatically on file changes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-service-token-caches.html
	ClearCachedServiceTokens security_clear_cached_service_tokens.NewClearCachedServiceTokens
	// Create an API key.
	//
	// Create an API key for access without requiring basic authentication.
	//
	// IMPORTANT: If the credential that is used to authenticate this request is an
	// API key, the derived API key cannot have any privileges.
	// If you specify privileges, the API returns an error.
	//
	// A successful request returns a JSON structure that contains the API key, its
	// unique id, and its name.
	// If applicable, it also returns expiration information for the API key in
	// milliseconds.
	//
	// NOTE: By default, API keys never expire. You can specify expiration
	// information when you create the API keys.
	//
	// The API keys are created by the Elasticsearch API key service, which is
	// automatically enabled.
	// To configure or turn off the API key service, refer to API key service
	// setting documentation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
	CreateApiKey security_create_api_key.NewCreateApiKey
	// Create a cross-cluster API key.
	//
	// Create an API key of the `cross_cluster` type for the API key based remote
	// cluster access.
	// A `cross_cluster` API key cannot be used to authenticate through the REST
	// interface.
	//
	// IMPORTANT: To authenticate this request you must use a credential that is not
	// an API key. Even if you use an API key that has the required privilege, the
	// API returns an error.
	//
	// Cross-cluster API keys are created by the Elasticsearch API key service,
	// which is automatically enabled.
	//
	// NOTE: Unlike REST API keys, a cross-cluster API key does not capture
	// permissions of the authenticated user. The API key’s effective permission is
	// exactly as specified with the `access` property.
	//
	// A successful request returns a JSON structure that contains the API key, its
	// unique ID, and its name. If applicable, it also returns expiration
	// information for the API key in milliseconds.
	//
	// By default, API keys never expire. You can specify expiration information
	// when you create the API keys.
	//
	// Cross-cluster API keys can only be updated with the update cross-cluster API
	// key API.
	// Attempting to update them with the update REST API key API or the bulk update
	// REST API keys API will result in an error.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-cross-cluster-api-key.html
	CreateCrossClusterApiKey security_create_cross_cluster_api_key.NewCreateCrossClusterApiKey
	// Create a service account token.
	//
	// Create a service accounts token for access without requiring basic
	// authentication.
	//
	// NOTE: Service account tokens never expire.
	// You must actively delete them if they are no longer needed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-service-token.html
	CreateServiceToken security_create_service_token.NewCreateServiceToken
	// Delegate PKI authentication.
	//
	// This API implements the exchange of an X509Certificate chain for an
	// Elasticsearch access token.
	// The certificate chain is validated, according to RFC 5280, by sequentially
	// considering the trust configuration of every installed PKI realm that has
	// `delegation.enabled` set to `true`.
	// A successfully trusted client certificate is also subject to the validation
	// of the subject distinguished name according to thw `username_pattern` of the
	// respective realm.
	//
	// This API is called by smart and trusted proxies, such as Kibana, which
	// terminate the user's TLS session but still want to authenticate the user by
	// using a PKI realm—-​as if the user connected directly to Elasticsearch.
	//
	// IMPORTANT: The association between the subject public key in the target
	// certificate and the corresponding private key is not validated.
	// This is part of the TLS authentication process and it is delegated to the
	// proxy that calls this API.
	// The proxy is trusted to have performed the TLS authentication and this API
	// translates that authentication into an Elasticsearch access token.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delegate-pki-authentication.html
	DelegatePki security_delegate_pki.NewDelegatePki
	// Delete application privileges.
	//
	// To use this API, you must have one of the following privileges:
	//
	// * The `manage_security` cluster privilege (or a greater privilege such as
	// `all`).
	// * The "Manage Application Privileges" global privilege for the application
	// being referenced in the request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-privilege.html
	DeletePrivileges security_delete_privileges.NewDeletePrivileges
	// Delete roles.
	//
	// Delete roles in the native realm.
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The delete roles API cannot remove roles that are defined in roles files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-role.html
	DeleteRole security_delete_role.NewDeleteRole
	// Delete role mappings.
	//
	// Role mappings define which roles are assigned to each user.
	// The role mapping APIs are generally the preferred way to manage role mappings
	// rather than using role mapping files.
	// The delete role mappings API cannot remove role mappings that are defined in
	// role mapping files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-role-mapping.html
	DeleteRoleMapping security_delete_role_mapping.NewDeleteRoleMapping
	// Delete service account tokens.
	//
	// Delete service account tokens for a service in a specified namespace.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-service-token.html
	DeleteServiceToken security_delete_service_token.NewDeleteServiceToken
	// Delete users.
	//
	// Delete users from the native realm.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-delete-user.html
	DeleteUser security_delete_user.NewDeleteUser
	// Disable users.
	//
	// Disable users in the native realm.
	// By default, when you create users, they are enabled.
	// You can use this API to revoke a user's access to Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-disable-user.html
	DisableUser security_disable_user.NewDisableUser
	// Disable a user profile.
	//
	// Disable user profiles so that they are not visible in user profile searches.
	//
	// NOTE: The user profile feature is designed only for use by Kibana and
	// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
	// Individual users and external applications should not call this API directly.
	// Elastic reserves the right to change or remove this feature in future
	// releases without prior notice.
	//
	// When you activate a user profile, its automatically enabled and visible in
	// user profile searches. You can use the disable user profile API to disable a
	// user profile so it’s not visible in these searches.
	// To re-enable a disabled user profile, use the enable user profile API .
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-disable-user-profile.html
	DisableUserProfile security_disable_user_profile.NewDisableUserProfile
	// Enable users.
	//
	// Enable users in the native realm.
	// By default, when you create users, they are enabled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-enable-user.html
	EnableUser security_enable_user.NewEnableUser
	// Enable a user profile.
	//
	// Enable user profiles to make them visible in user profile searches.
	//
	// NOTE: The user profile feature is designed only for use by Kibana and
	// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
	// Individual users and external applications should not call this API directly.
	// Elastic reserves the right to change or remove this feature in future
	// releases without prior notice.
	//
	// When you activate a user profile, it's automatically enabled and visible in
	// user profile searches.
	// If you later disable the user profile, you can use the enable user profile
	// API to make the profile visible in these searches again.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-enable-user-profile.html
	EnableUserProfile security_enable_user_profile.NewEnableUserProfile
	// Enroll Kibana.
	//
	// Enable a Kibana instance to configure itself for communication with a secured
	// Elasticsearch cluster.
	//
	// NOTE: This API is currently intended for internal use only by Kibana.
	// Kibana uses this API internally to configure itself for communications with
	// an Elasticsearch cluster that already has security features enabled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-kibana-enrollment.html
	EnrollKibana security_enroll_kibana.NewEnrollKibana
	// Enroll a node.
	//
	// Enroll a new node to allow it to join an existing cluster with security
	// features enabled.
	//
	// The response contains all the necessary information for the joining node to
	// bootstrap discovery and security related settings so that it can successfully
	// join the cluster.
	// The response contains key and certificate material that allows the caller to
	// generate valid signed certificates for the HTTP layer of all nodes in the
	// cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-node-enrollment.html
	EnrollNode security_enroll_node.NewEnrollNode
	// Get API key information.
	//
	// Retrieves information for one or more API keys.
	// NOTE: If you have only the `manage_own_api_key` privilege, this API returns
	// only the API keys that you own.
	// If you have `read_security`, `manage_api_key` or greater privileges
	// (including `manage_security`), this API returns all API keys regardless of
	// ownership.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html
	GetApiKey security_get_api_key.NewGetApiKey
	// Get builtin privileges.
	//
	// Get the list of cluster privileges and index privileges that are available in
	// this version of Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-builtin-privileges.html
	GetBuiltinPrivileges security_get_builtin_privileges.NewGetBuiltinPrivileges
	// Get application privileges.
	//
	// To use this API, you must have one of the following privileges:
	//
	// * The `read_security` cluster privilege (or a greater privilege such as
	// `manage_security` or `all`).
	// * The "Manage Application Privileges" global privilege for the application
	// being referenced in the request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-privileges.html
	GetPrivileges security_get_privileges.NewGetPrivileges
	// Get roles.
	//
	// Get roles in the native realm.
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The get roles API cannot retrieve roles that are defined in roles files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role.html
	GetRole security_get_role.NewGetRole
	// Get role mappings.
	//
	// Role mappings define which roles are assigned to each user.
	// The role mapping APIs are generally the preferred way to manage role mappings
	// rather than using role mapping files.
	// The get role mappings API cannot retrieve role mappings that are defined in
	// role mapping files.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role-mapping.html
	GetRoleMapping security_get_role_mapping.NewGetRoleMapping
	// Get service accounts.
	//
	// Get a list of service accounts that match the provided path parameters.
	//
	// NOTE: Currently, only the `elastic/fleet-server` service account is
	// available.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-service-accounts.html
	GetServiceAccounts security_get_service_accounts.NewGetServiceAccounts
	// Get service account credentials.
	//
	// To use this API, you must have at least the `read_security` cluster privilege
	// (or a greater privilege such as `manage_service_account` or
	// `manage_security`).
	//
	// The response includes service account tokens that were created with the
	// create service account tokens API as well as file-backed tokens from all
	// nodes of the cluster.
	//
	// NOTE: For tokens backed by the `service_tokens` file, the API collects them
	// from all nodes of the cluster.
	// Tokens with the same name from different nodes are assumed to be the same
	// token and are only counted once towards the total number of service tokens.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-service-credentials.html
	GetServiceCredentials security_get_service_credentials.NewGetServiceCredentials
	// Get security index settings.
	//
	// Get the user-configurable settings for the security internal index
	// (`.security` and associated indices).
	// Only a subset of the index settings — those that are user-configurable—will
	// be shown.
	// This includes:
	//
	// * `index.auto_expand_replicas`
	// * `index.number_of_replicas`
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-settings.html
	GetSettings security_get_settings.NewGetSettings
	// Get a token.
	//
	// Create a bearer token for access without requiring basic authentication.
	// The tokens are created by the Elasticsearch Token Service, which is
	// automatically enabled when you configure TLS on the HTTP interface.
	// Alternatively, you can explicitly enable the
	// `xpack.security.authc.token.enabled` setting.
	// When you are running in production mode, a bootstrap check prevents you from
	// enabling the token service unless you also enable TLS on the HTTP interface.
	//
	// The get token API takes the same parameters as a typical OAuth 2.0 token API
	// except for the use of a JSON request body.
	//
	// A successful get token API call returns a JSON structure that contains the
	// access token, the amount of time (seconds) that the token expires in, the
	// type, and the scope if available.
	//
	// The tokens returned by the get token API have a finite period of time for
	// which they are valid and after that time period, they can no longer be used.
	// That time period is defined by the `xpack.security.authc.token.timeout`
	// setting.
	// If you want to invalidate a token immediately, you can do so by using the
	// invalidate token API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-token.html
	GetToken security_get_token.NewGetToken
	// Get users.
	//
	// Get information about users in the native realm and built-in users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user.html
	GetUser security_get_user.NewGetUser
	// Get user privileges.
	//
	// Get the security privileges for the logged in user.
	// All users can use this API, but only to determine their own privileges.
	// To check the privileges of other users, you must use the run as feature.
	// To check whether a user has a specific list of privileges, use the has
	// privileges API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-privileges.html
	GetUserPrivileges security_get_user_privileges.NewGetUserPrivileges
	// Get a user profile.
	//
	// Get a user's profile using the unique profile ID.
	//
	// NOTE: The user profile feature is designed only for use by Kibana and
	// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
	// Individual users and external applications should not call this API directly.
	// Elastic reserves the right to change or remove this feature in future
	// releases without prior notice.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-profile.html
	GetUserProfile security_get_user_profile.NewGetUserProfile
	// Grant an API key.
	//
	// Create an API key on behalf of another user.
	// This API is similar to the create API keys API, however it creates the API
	// key for a user that is different than the user that runs the API.
	// The caller must have authentication credentials for the user on whose behalf
	// the API key will be created.
	// It is not possible to use this API to create an API key without that user's
	// credentials.
	// The supported user authentication credential types are:
	//
	// * username and password
	// * Elasticsearch access tokens
	// * JWTs
	//
	// The user, for whom the authentication credentials is provided, can optionally
	// "run as" (impersonate) another user.
	// In this case, the API key will be created on behalf of the impersonated user.
	//
	// This API is intended be used by applications that need to create and manage
	// API keys for end users, but cannot guarantee that those users have permission
	// to create API keys on their own behalf.
	// The API keys are created by the Elasticsearch API key service, which is
	// automatically enabled.
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
	// Check user privileges.
	//
	// Determine whether the specified user has a specified list of privileges.
	// All users can use this API, but only to determine their own privileges.
	// To check the privileges of other users, you must use the run as feature.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges.html
	HasPrivileges security_has_privileges.NewHasPrivileges
	// Check user profile privileges.
	//
	// Determine whether the users associated with the specified user profile IDs
	// have all the requested privileges.
	//
	// NOTE: The user profile feature is designed only for use by Kibana and
	// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
	// Individual users and external applications should not call this API directly.
	// Elastic reserves the right to change or remove this feature in future
	// releases without prior notice.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges-user-profile.html
	HasPrivilegesUserProfile security_has_privileges_user_profile.NewHasPrivilegesUserProfile
	// Invalidate API keys.
	//
	// This API invalidates API keys created by the create API key or grant API key
	// APIs.
	// Invalidated API keys fail authentication, but they can still be viewed using
	// the get API key information and query API key information APIs, for at least
	// the configured retention period, until they are automatically deleted.
	//
	// To use this API, you must have at least the `manage_security`,
	// `manage_api_key`, or `manage_own_api_key` cluster privileges.
	// The `manage_security` privilege allows deleting any API key, including both
	// REST and cross cluster API keys.
	// The `manage_api_key` privilege allows deleting any REST API key, but not
	// cross cluster API keys.
	// The `manage_own_api_key` only allows deleting REST API keys that are owned by
	// the user.
	// In addition, with the `manage_own_api_key` privilege, an invalidation request
	// must be issued in one of the three formats:
	//
	// - Set the parameter `owner=true`.
	// - Or, set both `username` and `realm_name` to match the user's identity.
	// - Or, if the request is issued by an API key, that is to say an API key
	// invalidates itself, specify its ID in the `ids` field.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-invalidate-api-key.html
	InvalidateApiKey security_invalidate_api_key.NewInvalidateApiKey
	// Invalidate a token.
	//
	// The access tokens returned by the get token API have a finite period of time
	// for which they are valid.
	// After that time period, they can no longer be used.
	// The time period is defined by the `xpack.security.authc.token.timeout`
	// setting.
	//
	// The refresh tokens returned by the get token API are only valid for 24 hours.
	// They can also be used exactly once.
	// If you want to invalidate one or more access or refresh tokens immediately,
	// use this invalidate token API.
	//
	// NOTE: While all parameters are optional, at least one of them is required.
	// More specifically, either one of `token` or `refresh_token` parameters is
	// required.
	// If none of these two are specified, then `realm_name` and/or `username` need
	// to be specified.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-invalidate-token.html
	InvalidateToken security_invalidate_token.NewInvalidateToken
	// Authenticate OpenID Connect.
	//
	// Exchange an OpenID Connect authentication response message for an
	// Elasticsearch internal access token and refresh token that can be
	// subsequently used for authentication.
	//
	// Elasticsearch exposes all the necessary OpenID Connect related functionality
	// with the OpenID Connect APIs.
	// These APIs are used internally by Kibana in order to provide OpenID Connect
	// based authentication, but can also be used by other, custom web applications
	// or other clients.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-oidc-authenticate.html
	OidcAuthenticate security_oidc_authenticate.NewOidcAuthenticate
	// Logout of OpenID Connect.
	//
	// Invalidate an access token and a refresh token that were generated as a
	// response to the `/_security/oidc/authenticate` API.
	//
	// If the OpenID Connect authentication realm in Elasticsearch is accordingly
	// configured, the response to this call will contain a URI pointing to the end
	// session endpoint of the OpenID Connect Provider in order to perform single
	// logout.
	//
	// Elasticsearch exposes all the necessary OpenID Connect related functionality
	// with the OpenID Connect APIs.
	// These APIs are used internally by Kibana in order to provide OpenID Connect
	// based authentication, but can also be used by other, custom web applications
	// or other clients.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-oidc-logout.html
	OidcLogout security_oidc_logout.NewOidcLogout
	// Prepare OpenID connect authentication.
	//
	// Create an oAuth 2.0 authentication request as a URL string based on the
	// configuration of the OpenID Connect authentication realm in Elasticsearch.
	//
	// The response of this API is a URL pointing to the Authorization Endpoint of
	// the configured OpenID Connect Provider, which can be used to redirect the
	// browser of the user in order to continue the authentication process.
	//
	// Elasticsearch exposes all the necessary OpenID Connect related functionality
	// with the OpenID Connect APIs.
	// These APIs are used internally by Kibana in order to provide OpenID Connect
	// based authentication, but can also be used by other, custom web applications
	// or other clients.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-oidc-prepare-authentication.html
	OidcPrepareAuthentication security_oidc_prepare_authentication.NewOidcPrepareAuthentication
	// Create or update application privileges.
	//
	// To use this API, you must have one of the following privileges:
	//
	// * The `manage_security` cluster privilege (or a greater privilege such as
	// `all`).
	// * The "Manage Application Privileges" global privilege for the application
	// being referenced in the request.
	//
	// Application names are formed from a prefix, with an optional suffix that
	// conform to the following rules:
	//
	// * The prefix must begin with a lowercase ASCII letter.
	// * The prefix must contain only ASCII letters or digits.
	// * The prefix must be at least 3 characters long.
	// * If the suffix exists, it must begin with either a dash `-` or `_`.
	// * The suffix cannot contain any of the following characters: `\`, `/`, `*`,
	// `?`, `"`, `<`, `>`, `|`, `,`, `*`.
	// * No part of the name can contain whitespace.
	//
	// Privilege names must begin with a lowercase ASCII letter and must contain
	// only ASCII letters and digits along with the characters `_`, `-`, and `.`.
	//
	// Action names can contain any number of printable ASCII characters and must
	// contain at least one of the following characters: `/`, `*`, `:`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-privileges.html
	PutPrivileges security_put_privileges.NewPutPrivileges
	// Create or update roles.
	//
	// The role management APIs are generally the preferred way to manage roles in
	// the native realm, rather than using file-based role management.
	// The create or update roles API cannot update roles that are defined in roles
	// files.
	// File-based role management is not available in Elastic Serverless.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role.html
	PutRole security_put_role.NewPutRole
	// Create or update role mappings.
	//
	// Role mappings define which roles are assigned to each user.
	// Each mapping has rules that identify users and a list of roles that are
	// granted to those users.
	// The role mapping APIs are generally the preferred way to manage role mappings
	// rather than using role mapping files. The create or update role mappings API
	// cannot update role mappings that are defined in role mapping files.
	//
	// NOTE: This API does not create roles. Rather, it maps users to existing
	// roles.
	// Roles can be created by using the create or update roles API or roles files.
	//
	// **Role templates**
	//
	// The most common use for role mappings is to create a mapping from a known
	// value on the user to a fixed role name.
	// For example, all users in the `cn=admin,dc=example,dc=com` LDAP group should
	// be given the superuser role in Elasticsearch.
	// The `roles` field is used for this purpose.
	//
	// For more complex needs, it is possible to use Mustache templates to
	// dynamically determine the names of the roles that should be granted to the
	// user.
	// The `role_templates` field is used for this purpose.
	//
	// NOTE: To use role templates successfully, the relevant scripting feature must
	// be enabled.
	// Otherwise, all attempts to create a role mapping with role templates fail.
	//
	// All of the user fields that are available in the role mapping rules are also
	// available in the role templates.
	// Thus it is possible to assign a user to a role that reflects their username,
	// their groups, or the name of the realm to which they authenticated.
	//
	// By default a template is evaluated to produce a single string that is the
	// name of the role which should be assigned to the user.
	// If the format of the template is set to "json" then the template is expected
	// to produce a JSON string or an array of JSON strings for the role names.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role-mapping.html
	PutRoleMapping security_put_role_mapping.NewPutRoleMapping
	// Create or update users.
	//
	// Add and update users in the native realm.
	// A password is required for adding a new user but is optional when updating an
	// existing user.
	// To change a user's password without updating any other fields, use the change
	// password API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html
	PutUser security_put_user.NewPutUser
	// Find API keys with a query.
	//
	// Get a paginated list of API keys and their information.
	// You can optionally filter the results with a query.
	//
	// To use this API, you must have at least the `manage_own_api_key` or the
	// `read_security` cluster privileges.
	// If you have only the `manage_own_api_key` privilege, this API returns only
	// the API keys that you own.
	// If you have the `read_security`, `manage_api_key`, or greater privileges
	// (including `manage_security`), this API returns all API keys regardless of
	// ownership.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-api-key.html
	QueryApiKeys security_query_api_keys.NewQueryApiKeys
	// Find roles with a query.
	//
	// Get roles in a paginated manner.
	// The role management APIs are generally the preferred way to manage roles,
	// rather than using file-based role management.
	// The query roles API does not retrieve roles that are defined in roles files,
	// nor built-in ones.
	// You can optionally filter the results with a query.
	// Also, the results can be paginated and sorted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-role.html
	QueryRole security_query_role.NewQueryRole
	// Find users with a query.
	//
	// Get information for users in a paginated manner.
	// You can optionally filter the results with a query.
	//
	// NOTE: As opposed to the get user API, built-in users are excluded from the
	// result.
	// This API is only for native users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-user.html
	QueryUser security_query_user.NewQueryUser
	// Authenticate SAML.
	//
	// Submit a SAML response message to Elasticsearch for consumption.
	//
	// NOTE: This API is intended for use by custom web applications other than
	// Kibana.
	// If you are using Kibana, refer to the documentation for configuring SAML
	// single-sign-on on the Elastic Stack.
	//
	// The SAML message that is submitted can be:
	//
	// * A response to a SAML authentication request that was previously created
	// using the SAML prepare authentication API.
	// * An unsolicited SAML message in the case of an IdP-initiated single sign-on
	// (SSO) flow.
	//
	// In either case, the SAML message needs to be a base64 encoded XML document
	// with a root element of `<Response>`.
	//
	// After successful validation, Elasticsearch responds with an Elasticsearch
	// internal access token and refresh token that can be subsequently used for
	// authentication.
	// This API endpoint essentially exchanges SAML responses that indicate
	// successful authentication in the IdP for Elasticsearch access and refresh
	// tokens, which can be used for authentication against Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-authenticate.html
	SamlAuthenticate security_saml_authenticate.NewSamlAuthenticate
	// Logout of SAML completely.
	//
	// Verifies the logout response sent from the SAML IdP.
	//
	// NOTE: This API is intended for use by custom web applications other than
	// Kibana.
	// If you are using Kibana, refer to the documentation for configuring SAML
	// single-sign-on on the Elastic Stack.
	//
	// The SAML IdP may send a logout response back to the SP after handling the
	// SP-initiated SAML Single Logout.
	// This API verifies the response by ensuring the content is relevant and
	// validating its signature.
	// An empty response is returned if the verification process is successful.
	// The response can be sent by the IdP with either the HTTP-Redirect or the
	// HTTP-Post binding.
	// The caller of this API must prepare the request accordingly so that this API
	// can handle either of them.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-complete-logout.html
	SamlCompleteLogout security_saml_complete_logout.NewSamlCompleteLogout
	// Invalidate SAML.
	//
	// Submit a SAML LogoutRequest message to Elasticsearch for consumption.
	//
	// NOTE: This API is intended for use by custom web applications other than
	// Kibana.
	// If you are using Kibana, refer to the documentation for configuring SAML
	// single-sign-on on the Elastic Stack.
	//
	// The logout request comes from the SAML IdP during an IdP initiated Single
	// Logout.
	// The custom web application can use this API to have Elasticsearch process the
	// `LogoutRequest`.
	// After successful validation of the request, Elasticsearch invalidates the
	// access token and refresh token that corresponds to that specific SAML
	// principal and provides a URL that contains a SAML LogoutResponse message.
	// Thus the user can be redirected back to their IdP.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-invalidate.html
	SamlInvalidate security_saml_invalidate.NewSamlInvalidate
	// Logout of SAML.
	//
	// Submits a request to invalidate an access token and refresh token.
	//
	// NOTE: This API is intended for use by custom web applications other than
	// Kibana.
	// If you are using Kibana, refer to the documentation for configuring SAML
	// single-sign-on on the Elastic Stack.
	//
	// This API invalidates the tokens that were generated for a user by the SAML
	// authenticate API.
	// If the SAML realm in Elasticsearch is configured accordingly and the SAML IdP
	// supports this, the Elasticsearch response contains a URL to redirect the user
	// to the IdP that contains a SAML logout request (starting an SP-initiated SAML
	// Single Logout).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-logout.html
	SamlLogout security_saml_logout.NewSamlLogout
	// Prepare SAML authentication.
	//
	// Create a SAML authentication request (`<AuthnRequest>`) as a URL string based
	// on the configuration of the respective SAML realm in Elasticsearch.
	//
	// NOTE: This API is intended for use by custom web applications other than
	// Kibana.
	// If you are using Kibana, refer to the documentation for configuring SAML
	// single-sign-on on the Elastic Stack.
	//
	// This API returns a URL pointing to the SAML Identity Provider.
	// You can use the URL to redirect the browser of the user in order to continue
	// the authentication process.
	// The URL includes a single parameter named `SAMLRequest`, which contains a
	// SAML Authentication request that is deflated and Base64 encoded.
	// If the configuration dictates that SAML authentication requests should be
	// signed, the URL has two extra parameters named `SigAlg` and `Signature`.
	// These parameters contain the algorithm used for the signature and the
	// signature value itself.
	// It also returns a random string that uniquely identifies this SAML
	// Authentication request.
	// The caller of this API needs to store this identifier as it needs to be used
	// in a following step of the authentication process.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-prepare-authentication.html
	SamlPrepareAuthentication security_saml_prepare_authentication.NewSamlPrepareAuthentication
	// Create SAML service provider metadata.
	//
	// Generate SAML metadata for a SAML 2.0 Service Provider.
	//
	// The SAML 2.0 specification provides a mechanism for Service Providers to
	// describe their capabilities and configuration using a metadata file.
	// This API generates Service Provider metadata based on the configuration of a
	// SAML realm in Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-sp-metadata.html
	SamlServiceProviderMetadata security_saml_service_provider_metadata.NewSamlServiceProviderMetadata
	// Suggest a user profile.
	//
	// Get suggestions for user profiles that match specified search criteria.
	//
	// NOTE: The user profile feature is designed only for use by Kibana and
	// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
	// Individual users and external applications should not call this API directly.
	// Elastic reserves the right to change or remove this feature in future
	// releases without prior notice.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-suggest-user-profile.html
	SuggestUserProfiles security_suggest_user_profiles.NewSuggestUserProfiles
	// Update an API key.
	//
	// Update attributes of an existing API key.
	// This API supports updates to an API key's access scope, expiration, and
	// metadata.
	//
	// To use this API, you must have at least the `manage_own_api_key` cluster
	// privilege.
	// Users can only update API keys that they created or that were granted to
	// them.
	// To update another user’s API key, use the `run_as` feature to submit a
	// request on behalf of another user.
	//
	// IMPORTANT: It's not possible to use an API key as the authentication
	// credential for this API. The owner user’s credentials are required.
	//
	// Use this API to update API keys created by the create API key or grant API
	// Key APIs.
	// If you need to apply the same update to many API keys, you can use the bulk
	// update API keys API to reduce overhead.
	// It's not possible to update expired API keys or API keys that have been
	// invalidated by the invalidate API key API.
	//
	// The access scope of an API key is derived from the `role_descriptors` you
	// specify in the request and a snapshot of the owner user's permissions at the
	// time of the request.
	// The snapshot of the owner's permissions is updated automatically on every
	// call.
	//
	// IMPORTANT: If you don't specify `role_descriptors` in the request, a call to
	// this API might still change the API key's access scope.
	// This change can occur if the owner user's permissions have changed since the
	// API key was created or last modified.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-api-key.html
	UpdateApiKey security_update_api_key.NewUpdateApiKey
	// Update a cross-cluster API key.
	//
	// Update the attributes of an existing cross-cluster API key, which is used for
	// API key based remote cluster access.
	//
	// To use this API, you must have at least the `manage_security` cluster
	// privilege.
	// Users can only update API keys that they created.
	// To update another user's API key, use the `run_as` feature to submit a
	// request on behalf of another user.
	//
	// IMPORTANT: It's not possible to use an API key as the authentication
	// credential for this API.
	// To update an API key, the owner user's credentials are required.
	//
	// It's not possible to update expired API keys, or API keys that have been
	// invalidated by the invalidate API key API.
	//
	// This API supports updates to an API key's access scope, metadata, and
	// expiration.
	// The owner user's information, such as the `username` and `realm`, is also
	// updated automatically on every call.
	//
	// NOTE: This API cannot update REST API keys, which should be updated by either
	// the update API key or bulk update API keys API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-cross-cluster-api-key.html
	UpdateCrossClusterApiKey security_update_cross_cluster_api_key.NewUpdateCrossClusterApiKey
	// Update security index settings.
	//
	// Update the user-configurable settings for the security internal index
	// (`.security` and associated indices). Only a subset of settings are allowed
	// to be modified. This includes `index.auto_expand_replicas` and
	// `index.number_of_replicas`.
	//
	// NOTE: If `index.auto_expand_replicas` is set, `index.number_of_replicas` will
	// be ignored during updates.
	//
	// If a specific index is not in use on the system and settings are provided for
	// it, the request will be rejected.
	// This API does not yet support configuring the settings for indices before
	// they are in use.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-settings.html
	UpdateSettings security_update_settings.NewUpdateSettings
	// Update user profile data.
	//
	// Update specific data for the user profile that is associated with a unique
	// ID.
	//
	// NOTE: The user profile feature is designed only for use by Kibana and
	// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
	// Individual users and external applications should not call this API directly.
	// Elastic reserves the right to change or remove this feature in future
	// releases without prior notice.
	//
	// To use this API, you must have one of the following privileges:
	//
	// * The `manage_user_profile` cluster privilege.
	// * The `update_profile_data` global privilege for the namespaces that are
	// referenced in the request.
	//
	// This API updates the `labels` and `data` fields of an existing user profile
	// document with JSON objects.
	// New keys and their values are added to the profile document and conflicting
	// keys are replaced by data that's included in the request.
	//
	// For both labels and data, content is namespaced by the top-level fields.
	// The `update_profile_data` global privilege grants privileges for updating
	// only the allowed namespaces.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-user-profile-data.html
	UpdateUserProfileData security_update_user_profile_data.NewUpdateUserProfileData
}

type Shutdown struct {
	// Cancel node shutdown preparations.
	// Remove a node from the shutdown list so it can resume normal operations.
	// You must explicitly clear the shutdown request when a node rejoins the
	// cluster or when a node has permanently left the cluster.
	// Shutdown requests are never removed automatically by Elasticsearch.
	//
	// NOTE: This feature is designed for indirect use by Elastic Cloud, Elastic
	// Cloud Enterprise, and Elastic Cloud on Kubernetes.
	// Direct use is not supported.
	//
	// If the operator privileges feature is enabled, you must be an operator to use
	// this API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-shutdown.html
	DeleteNode shutdown_delete_node.NewDeleteNode
	// Get the shutdown status.
	//
	// Get information about nodes that are ready to be shut down, have shut down
	// preparations still in progress, or have stalled.
	// The API returns status information for each part of the shut down process.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	//
	// If the operator privileges feature is enabled, you must be an operator to use
	// this API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-shutdown.html
	GetNode shutdown_get_node.NewGetNode
	// Prepare a node to be shut down.
	//
	// NOTE: This feature is designed for indirect use by Elastic Cloud, Elastic
	// Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	//
	// If you specify a node that is offline, it will be prepared for shut down when
	// it rejoins the cluster.
	//
	// If the operator privileges feature is enabled, you must be an operator to use
	// this API.
	//
	// The API migrates ongoing tasks and index shards to other nodes as needed to
	// prepare a node to be restarted or shut down and removed from the cluster.
	// This ensures that Elasticsearch can be stopped safely with minimal disruption
	// to the cluster.
	//
	// You must specify the type of shutdown: `restart`, `remove`, or `replace`.
	// If a node is already being prepared for shutdown, you can use this API to
	// change the shutdown type.
	//
	// IMPORTANT: This API does NOT terminate the Elasticsearch process.
	// Monitor the node shutdown status to determine when it is safe to stop
	// Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-shutdown.html
	PutNode shutdown_put_node.NewPutNode
}

type Simulate struct {
	// Simulate data ingestion.
	// Run ingest pipelines against a set of provided documents, optionally with
	// substitute pipeline definitions, to simulate ingesting data into an index.
	//
	// This API is meant to be used for troubleshooting or pipeline development, as
	// it does not actually index any data into Elasticsearch.
	//
	// The API runs the default and final pipeline for that index against a set of
	// documents provided in the body of the request.
	// If a pipeline contains a reroute processor, it follows that reroute processor
	// to the new index, running that index's pipelines as well the same way that a
	// non-simulated ingest would.
	// No data is indexed into Elasticsearch.
	// Instead, the transformed document is returned, along with the list of
	// pipelines that have been run and the name of the index where the document
	// would have been indexed if this were not a simulation.
	// The transformed document is validated against the mappings that would apply
	// to this index, and any validation error is reported in the result.
	//
	// This API differs from the simulate pipeline API in that you specify a single
	// pipeline for that API, and it runs only that one pipeline.
	// The simulate pipeline API is more useful for developing a single pipeline,
	// while the simulate ingest API is more useful for troubleshooting the
	// interaction of the various pipelines that get applied when ingesting into an
	// index.
	//
	// By default, the pipeline definitions that are currently in the system are
	// used.
	// However, you can supply substitute pipeline definitions in the body of the
	// request.
	// These will be used in place of the pipeline definitions that are already in
	// the system. This can be used to replace existing pipeline definitions or to
	// create new ones. The pipeline substitutions are used only within this
	// request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-ingest-api.html
	Ingest simulate_ingest.NewIngest
}

type Slm struct {
	// Delete a policy.
	// Delete a snapshot lifecycle policy definition.
	// This operation prevents any future snapshots from being taken but does not
	// cancel in-progress snapshots or remove previously-taken snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-delete-policy.html
	DeleteLifecycle slm_delete_lifecycle.NewDeleteLifecycle
	// Run a policy.
	// Immediately create a snapshot according to the snapshot lifecycle policy
	// without waiting for the scheduled time.
	// The snapshot policy is normally applied according to its schedule, but you
	// might want to manually run a policy before performing an upgrade or other
	// maintenance.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-execute-lifecycle.html
	ExecuteLifecycle slm_execute_lifecycle.NewExecuteLifecycle
	// Run a retention policy.
	// Manually apply the retention policy to force immediate removal of snapshots
	// that are expired according to the snapshot lifecycle policy retention rules.
	// The retention policy is normally applied according to its schedule.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-execute-retention.html
	ExecuteRetention slm_execute_retention.NewExecuteRetention
	// Get policy information.
	// Get snapshot lifecycle policy definitions and information about the latest
	// snapshot attempts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-get-policy.html
	GetLifecycle slm_get_lifecycle.NewGetLifecycle
	// Get snapshot lifecycle management statistics.
	// Get global and policy-level statistics about actions taken by snapshot
	// lifecycle management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-get-stats.html
	GetStats slm_get_stats.NewGetStats
	// Get the snapshot lifecycle management status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-get-status.html
	GetStatus slm_get_status.NewGetStatus
	// Create or update a policy.
	// Create or update a snapshot lifecycle policy.
	// If the policy already exists, this request increments the policy version.
	// Only the latest version of a policy is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-put-policy.html
	PutLifecycle slm_put_lifecycle.NewPutLifecycle
	// Start snapshot lifecycle management.
	// Snapshot lifecycle management (SLM) starts automatically when a cluster is
	// formed.
	// Manually starting SLM is necessary only if it has been stopped using the stop
	// SLM API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-start.html
	Start slm_start.NewStart
	// Stop snapshot lifecycle management.
	// Stop all snapshot lifecycle management (SLM) operations and the SLM plugin.
	// This API is useful when you are performing maintenance on a cluster and need
	// to prevent SLM from performing any actions on your data streams or indices.
	// Stopping SLM does not stop any snapshots that are in progress.
	// You can manually trigger snapshots with the run snapshot lifecycle policy API
	// even if SLM is stopped.
	//
	// The API returns a response as soon as the request is acknowledged, but the
	// plugin might continue to run until in-progress operations complete and it can
	// be safely stopped.
	// Use the get snapshot lifecycle management status API to see if SLM is
	// running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-stop.html
	Stop slm_stop.NewStop
}

type Snapshot struct {
	// Clean up the snapshot repository.
	// Trigger the review of the contents of a snapshot repository and delete any
	// stale data not referenced by existing snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clean-up-snapshot-repo-api.html
	CleanupRepository snapshot_cleanup_repository.NewCleanupRepository
	// Clone a snapshot.
	// Clone part of all of a snapshot into another snapshot in the same repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clone-snapshot-api.html
	Clone snapshot_clone.NewClone
	// Create a snapshot.
	// Take a snapshot of a cluster or of data streams and indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-snapshot-api.html
	Create snapshot_create.NewCreate
	// Create or update a snapshot repository.
	// IMPORTANT: If you are migrating searchable snapshots, the repository name
	// must be identical in the source and destination clusters.
	// To register a snapshot repository, the cluster's global metadata must be
	// writeable.
	// Ensure there are no cluster blocks (for example, `cluster.blocks.read_only`
	// and `clsuter.blocks.read_only_allow_delete` settings) that prevent write
	// access.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	CreateRepository snapshot_create_repository.NewCreateRepository
	// Delete snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-snapshot-api.html
	Delete snapshot_delete.NewDelete
	// Delete snapshot repositories.
	// When a repository is unregistered, Elasticsearch removes only the reference
	// to the location where the repository is storing the snapshots.
	// The snapshots themselves are left untouched and in place.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-snapshot-repo-api.html
	DeleteRepository snapshot_delete_repository.NewDeleteRepository
	// Get snapshot information.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-snapshot-api.html
	Get snapshot_get.NewGet
	// Get snapshot repository information.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-snapshot-repo-api.html
	GetRepository snapshot_get_repository.NewGetRepository
	// Analyze a snapshot repository.
	// Analyze the performance characteristics and any incorrect behaviour found in
	// a repository.
	//
	// The response exposes implementation details of the analysis which may change
	// from version to version.
	// The response body format is therefore not considered stable and may be
	// different in newer versions.
	//
	// There are a large number of third-party storage systems available, not all of
	// which are suitable for use as a snapshot repository by Elasticsearch.
	// Some storage systems behave incorrectly, or perform poorly, especially when
	// accessed concurrently by multiple clients as the nodes of an Elasticsearch
	// cluster do. This API performs a collection of read and write operations on
	// your repository which are designed to detect incorrect behaviour and to
	// measure the performance characteristics of your storage system.
	//
	// The default values for the parameters are deliberately low to reduce the
	// impact of running an analysis inadvertently and to provide a sensible
	// starting point for your investigations.
	// Run your first analysis with the default parameter values to check for simple
	// problems.
	// If successful, run a sequence of increasingly large analyses until you
	// encounter a failure or you reach a `blob_count` of at least `2000`, a
	// `max_blob_size` of at least `2gb`, a `max_total_data_size` of at least `1tb`,
	// and a `register_operation_count` of at least `100`.
	// Always specify a generous timeout, possibly `1h` or longer, to allow time for
	// each analysis to run to completion.
	// Perform the analyses using a multi-node cluster of a similar size to your
	// production cluster so that it can detect any problems that only arise when
	// the repository is accessed by many nodes at once.
	//
	// If the analysis fails, Elasticsearch detected that your repository behaved
	// unexpectedly.
	// This usually means you are using a third-party storage system with an
	// incorrect or incompatible implementation of the API it claims to support.
	// If so, this storage system is not suitable for use as a snapshot repository.
	// You will need to work with the supplier of your storage system to address the
	// incompatibilities that Elasticsearch detects.
	//
	// If the analysis is successful, the API returns details of the testing
	// process, optionally including how long each operation took.
	// You can use this information to determine the performance of your storage
	// system.
	// If any operation fails or returns an incorrect result, the API returns an
	// error.
	// If the API returns an error, it may not have removed all the data it wrote to
	// the repository.
	// The error will indicate the location of any leftover data and this path is
	// also recorded in the Elasticsearch logs.
	// You should verify that this location has been cleaned up correctly.
	// If there is still leftover data at the specified location, you should
	// manually remove it.
	//
	// If the connection from your client to Elasticsearch is closed while the
	// client is waiting for the result of the analysis, the test is cancelled.
	// Some clients are configured to close their connection if no response is
	// received within a certain timeout.
	// An analysis takes a long time to complete so you might need to relax any such
	// client-side timeouts.
	// On cancellation the analysis attempts to clean up the data it was writing,
	// but it may not be able to remove it all.
	// The path to the leftover data is recorded in the Elasticsearch logs.
	// You should verify that this location has been cleaned up correctly.
	// If there is still leftover data at the specified location, you should
	// manually remove it.
	//
	// If the analysis is successful then it detected no incorrect behaviour, but
	// this does not mean that correct behaviour is guaranteed.
	// The analysis attempts to detect common bugs but it does not offer 100%
	// coverage.
	// Additionally, it does not test the following:
	//
	// * Your repository must perform durable writes. Once a blob has been written
	// it must remain in place until it is deleted, even after a power loss or
	// similar disaster.
	// * Your repository must not suffer from silent data corruption. Once a blob
	// has been written, its contents must remain unchanged until it is deliberately
	// modified or deleted.
	// * Your repository must behave correctly even if connectivity from the cluster
	// is disrupted. Reads and writes may fail in this case, but they must not
	// return incorrect results.
	//
	// IMPORTANT: An analysis writes a substantial amount of data to your repository
	// and then reads it back again.
	// This consumes bandwidth on the network between the cluster and the
	// repository, and storage space and I/O bandwidth on the repository itself.
	// You must ensure this load does not affect other users of these systems.
	// Analyses respect the repository settings `max_snapshot_bytes_per_sec` and
	// `max_restore_bytes_per_sec` if available and the cluster setting
	// `indices.recovery.max_bytes_per_sec` which you can use to limit the bandwidth
	// they consume.
	//
	// NOTE: This API is intended for exploratory use by humans. You should expect
	// the request parameters and the response format to vary in future versions.
	//
	// NOTE: Different versions of Elasticsearch may perform different checks for
	// repository compatibility, with newer versions typically being stricter than
	// older ones.
	// A storage system that passes repository analysis with one version of
	// Elasticsearch may fail with a different version.
	// This indicates it behaves incorrectly in ways that the former version did not
	// detect.
	// You must work with the supplier of your storage system to address the
	// incompatibilities detected by the repository analysis API in any version of
	// Elasticsearch.
	//
	// NOTE: This API may not work correctly in a mixed-version cluster.
	//
	// *Implementation details*
	//
	// NOTE: This section of documentation describes how the repository analysis API
	// works in this version of Elasticsearch, but you should expect the
	// implementation to vary between versions. The request parameters and response
	// format depend on details of the implementation so may also be different in
	// newer versions.
	//
	// The analysis comprises a number of blob-level tasks, as set by the
	// `blob_count` parameter and a number of compare-and-exchange operations on
	// linearizable registers, as set by the `register_operation_count` parameter.
	// These tasks are distributed over the data and master-eligible nodes in the
	// cluster for execution.
	//
	// For most blob-level tasks, the executing node first writes a blob to the
	// repository and then instructs some of the other nodes in the cluster to
	// attempt to read the data it just wrote.
	// The size of the blob is chosen randomly, according to the `max_blob_size` and
	// `max_total_data_size` parameters.
	// If any of these reads fails then the repository does not implement the
	// necessary read-after-write semantics that Elasticsearch requires.
	//
	// For some blob-level tasks, the executing node will instruct some of its peers
	// to attempt to read the data before the writing process completes.
	// These reads are permitted to fail, but must not return partial data.
	// If any read returns partial data then the repository does not implement the
	// necessary atomicity semantics that Elasticsearch requires.
	//
	// For some blob-level tasks, the executing node will overwrite the blob while
	// its peers are reading it.
	// In this case the data read may come from either the original or the
	// overwritten blob, but the read operation must not return partial data or a
	// mix of data from the two blobs.
	// If any of these reads returns partial data or a mix of the two blobs then the
	// repository does not implement the necessary atomicity semantics that
	// Elasticsearch requires for overwrites.
	//
	// The executing node will use a variety of different methods to write the blob.
	// For instance, where applicable, it will use both single-part and multi-part
	// uploads.
	// Similarly, the reading nodes will use a variety of different methods to read
	// the data back again.
	// For instance they may read the entire blob from start to end or may read only
	// a subset of the data.
	//
	// For some blob-level tasks, the executing node will cancel the write before it
	// is complete.
	// In this case, it still instructs some of the other nodes in the cluster to
	// attempt to read the blob but all of these reads must fail to find the blob.
	//
	// Linearizable registers are special blobs that Elasticsearch manipulates using
	// an atomic compare-and-exchange operation.
	// This operation ensures correct and strongly-consistent behavior even when the
	// blob is accessed by multiple nodes at the same time.
	// The detailed implementation of the compare-and-exchange operation on
	// linearizable registers varies by repository type.
	// Repository analysis verifies that that uncontended compare-and-exchange
	// operations on a linearizable register blob always succeed.
	// Repository analysis also verifies that contended operations either succeed or
	// report the contention but do not return incorrect results.
	// If an operation fails due to contention, Elasticsearch retries the operation
	// until it succeeds.
	// Most of the compare-and-exchange operations performed by repository analysis
	// atomically increment a counter which is represented as an 8-byte blob.
	// Some operations also verify the behavior on small blobs with sizes other than
	// 8 bytes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/repo-analysis-api.html
	RepositoryAnalyze snapshot_repository_analyze.NewRepositoryAnalyze
	// Verify the repository integrity.
	// Verify the integrity of the contents of a snapshot repository.
	//
	// This API enables you to perform a comprehensive check of the contents of a
	// repository, looking for any anomalies in its data or metadata which might
	// prevent you from restoring snapshots from the repository or which might cause
	// future snapshot create or delete operations to fail.
	//
	// If you suspect the integrity of the contents of one of your snapshot
	// repositories, cease all write activity to this repository immediately, set
	// its `read_only` option to `true`, and use this API to verify its integrity.
	// Until you do so:
	//
	// * It may not be possible to restore some snapshots from this repository.
	// * Searchable snapshots may report errors when searched or may have unassigned
	// shards.
	// * Taking snapshots into this repository may fail or may appear to succeed but
	// have created a snapshot which cannot be restored.
	// * Deleting snapshots from this repository may fail or may appear to succeed
	// but leave the underlying data on disk.
	// * Continuing to write to the repository while it is in an invalid state may
	// causing additional damage to its contents.
	//
	// If the API finds any problems with the integrity of the contents of your
	// repository, Elasticsearch will not be able to repair the damage.
	// The only way to bring the repository back into a fully working state after
	// its contents have been damaged is by restoring its contents from a repository
	// backup which was taken before the damage occurred.
	// You must also identify what caused the damage and take action to prevent it
	// from happening again.
	//
	// If you cannot restore a repository backup, register a new repository and use
	// this for all future snapshot operations.
	// In some cases it may be possible to recover some of the contents of a damaged
	// repository, either by restoring as many of its snapshots as needed and taking
	// new snapshots of the restored data, or by using the reindex API to copy data
	// from any searchable snapshots mounted from the damaged repository.
	//
	// Avoid all operations which write to the repository while the verify
	// repository integrity API is running.
	// If something changes the repository contents while an integrity verification
	// is running then Elasticsearch may incorrectly report having detected some
	// anomalies in its contents due to the concurrent writes.
	// It may also incorrectly fail to report some anomalies that the concurrent
	// writes prevented it from detecting.
	//
	// NOTE: This API is intended for exploratory use by humans. You should expect
	// the request parameters and the response format to vary in future versions.
	//
	// NOTE: This API may not work correctly in a mixed-version cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/verify-repo-integrity-api.html
	RepositoryVerifyIntegrity snapshot_repository_verify_integrity.NewRepositoryVerifyIntegrity
	// Restore a snapshot.
	// Restore a snapshot of a cluster or data streams and indices.
	//
	// You can restore a snapshot only to a running cluster with an elected master
	// node.
	// The snapshot repository must be registered and available to the cluster.
	// The snapshot and cluster versions must be compatible.
	//
	// To restore a snapshot, the cluster's global metadata must be writable. Ensure
	// there are't any cluster blocks that prevent writes. The restore operation
	// ignores index blocks.
	//
	// Before you restore a data stream, ensure the cluster contains a matching
	// index template with data streams enabled. To check, use the index management
	// feature in Kibana or the get index template API:
	//
	// ```
	// GET
	// _index_template/*?filter_path=index_templates.name,index_templates.index_template.index_patterns,index_templates.index_template.data_stream
	// ```
	//
	// If no such template exists, you can create one or restore a cluster state
	// that contains one. Without a matching index template, a data stream can't
	// roll over or create backing indices.
	//
	// If your snapshot contains data from App Search or Workplace Search, you must
	// restore the Enterprise Search encryption key before you restore the snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/restore-snapshot-api.html
	Restore snapshot_restore.NewRestore
	// Get the snapshot status.
	// Get a detailed description of the current state for each shard participating
	// in the snapshot.
	// Note that this API should be used only to obtain detailed shard-level
	// information for ongoing snapshots.
	// If this detail is not needed or you want to obtain information about one or
	// more existing snapshots, use the get snapshot API.
	//
	// WARNING: Using the API to return the status of any snapshots other than
	// currently running snapshots can be expensive.
	// The API requires a read from the repository for each shard in each snapshot.
	// For example, if you have 100 snapshots with 1,000 shards each, an API request
	// that includes all snapshots will require 100,000 reads (100 snapshots x 1,000
	// shards).
	//
	// Depending on the latency of your storage, such requests can take an extremely
	// long time to return results.
	// These requests can also tax machine resources and, when using cloud storage,
	// incur high processing costs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-snapshot-status-api.html
	Status snapshot_status.NewStatus
	// Verify a snapshot repository.
	// Check for common misconfigurations in a snapshot repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/verify-snapshot-repo-api.html
	VerifyRepository snapshot_verify_repository.NewVerifyRepository
}

type Sql struct {
	// Clear an SQL search cursor.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-sql-cursor-api.html
	ClearCursor sql_clear_cursor.NewClearCursor
	// Delete an async SQL search.
	// Delete an async SQL search or a stored synchronous SQL search.
	// If the search is still running, the API cancels it.
	//
	// If the Elasticsearch security features are enabled, only the following users
	// can use this API to delete a search:
	//
	// * Users with the `cancel_task` cluster privilege.
	// * The user who first submitted the search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-async-sql-search-api.html
	DeleteAsync sql_delete_async.NewDeleteAsync
	// Get async SQL search results.
	// Get the current status and available results for an async SQL search or
	// stored synchronous SQL search.
	//
	// If the Elasticsearch security features are enabled, only the user who first
	// submitted the SQL search can retrieve the search using this API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-sql-search-api.html
	GetAsync sql_get_async.NewGetAsync
	// Get the async SQL search status.
	// Get the current status of an async SQL search or a stored synchronous SQL
	// search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-sql-search-status-api.html
	GetAsyncStatus sql_get_async_status.NewGetAsyncStatus
	// Get SQL search results.
	// Run an SQL request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-search-api.html
	Query sql_query.NewQuery
	// Translate SQL into Elasticsearch queries.
	// Translate an SQL search into a search API request containing Query DSL.
	// It accepts the same request body parameters as the SQL search API, excluding
	// `cursor`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-translate-api.html
	Translate sql_translate.NewTranslate
}

type Ssl struct {
	// Get SSL certificates.
	//
	// Get information about the X.509 certificates that are used to encrypt
	// communications in the cluster.
	// The API returns a list that includes certificates from all TLS contexts
	// including:
	//
	// - Settings for transport and HTTP interfaces
	// - TLS settings that are used within authentication realms
	// - TLS settings for remote monitoring exporters
	//
	// The list includes certificates that are used for configuring trust, such as
	// those configured in the `xpack.security.transport.ssl.truststore` and
	// `xpack.security.transport.ssl.certificate_authorities` settings.
	// It also includes certificates that are used for configuring server identity,
	// such as `xpack.security.http.ssl.keystore` and
	// `xpack.security.http.ssl.certificate settings`.
	//
	// The list does not include certificates that are sourced from the default SSL
	// context of the Java Runtime Environment (JRE), even if those certificates are
	// in use within Elasticsearch.
	//
	// NOTE: When a PKCS#11 token is configured as the truststore of the JRE, the
	// API returns all the certificates that are included in the PKCS#11 token
	// irrespective of whether these are used in the Elasticsearch TLS
	// configuration.
	//
	// If Elasticsearch is configured to use a keystore or truststore, the API
	// output includes all certificates in that store, even though some of the
	// certificates might not be in active use within the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-ssl.html
	Certificates ssl_certificates.NewCertificates
}

type Synonyms struct {
	// Delete a synonym set.
	//
	// You can only delete a synonyms set that is not in use by any index analyzer.
	//
	// Synonyms sets can be used in synonym graph token filters and synonym token
	// filters.
	// These synonym filters can be used as part of search analyzers.
	//
	// Analyzers need to be loaded when an index is restored (such as when a node
	// starts, or the index becomes open).
	// Even if the analyzer is not used on any field mapping, it still needs to be
	// loaded on the index recovery phase.
	//
	// If any analyzers cannot be loaded, the index becomes unavailable and the
	// cluster status becomes red or yellow as index shards are not available.
	// To prevent that, synonyms sets that are used in analyzers can't be deleted.
	// A delete request in this case will return a 400 response code.
	//
	// To remove a synonyms set, you must first remove all indices that contain
	// analyzers using it.
	// You can migrate an index by creating a new index that does not contain the
	// token filter with the synonyms set, and use the reindex API in order to copy
	// over the index data.
	// Once finished, you can delete the index.
	// When the synonyms set is not used in analyzers, you will be able to delete
	// it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-synonyms-set.html
	DeleteSynonym synonyms_delete_synonym.NewDeleteSynonym
	// Delete a synonym rule.
	// Delete a synonym rule from a synonym set.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-synonym-rule.html
	DeleteSynonymRule synonyms_delete_synonym_rule.NewDeleteSynonymRule
	// Get a synonym set.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonyms-set.html
	GetSynonym synonyms_get_synonym.NewGetSynonym
	// Get a synonym rule.
	// Get a synonym rule from a synonym set.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonym-rule.html
	GetSynonymRule synonyms_get_synonym_rule.NewGetSynonymRule
	// Get all synonym sets.
	// Get a summary of all defined synonym sets.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonyms-set.html
	GetSynonymsSets synonyms_get_synonyms_sets.NewGetSynonymsSets
	// Create or update a synonym set.
	// Synonyms sets are limited to a maximum of 10,000 synonym rules per set.
	// If you need to manage more synonym rules, you can create multiple synonym
	// sets.
	//
	// When an existing synonyms set is updated, the search analyzers that use the
	// synonyms set are reloaded automatically for all indices.
	// This is equivalent to invoking the reload search analyzers API for all
	// indices that use the synonyms set.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonyms-set.html
	PutSynonym synonyms_put_synonym.NewPutSynonym
	// Create or update a synonym rule.
	// Create or update a synonym rule in a synonym set.
	//
	// If any of the synonym rules included is invalid, the API returns an error.
	//
	// When you update a synonym rule, all analyzers using the synonyms set will be
	// reloaded automatically to reflect the new rule.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonym-rule.html
	PutSynonymRule synonyms_put_synonym_rule.NewPutSynonymRule
}

type Tasks struct {
	// Cancel a task.
	//
	// WARNING: The task management API is new and should still be considered a beta
	// feature.
	// The API may change in ways that are not backwards compatible.
	//
	// A task may continue to run for some time after it has been cancelled because
	// it may not be able to safely stop its current activity straight away.
	// It is also possible that Elasticsearch must complete its work on other tasks
	// before it can process the cancellation.
	// The get task information API will continue to list these cancelled tasks
	// until they complete.
	// The cancelled flag in the response indicates that the cancellation command
	// has been processed and the task will stop as soon as possible.
	//
	// To troubleshoot why a cancelled task does not complete promptly, use the get
	// task information API with the `?detailed` parameter to identify the other
	// tasks the system is running.
	// You can also use the node hot threads API to obtain detailed information
	// about the work the system is doing instead of completing the cancelled task.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	Cancel tasks_cancel.NewCancel
	// Get task information.
	// Get information about a task currently running in the cluster.
	//
	// WARNING: The task management API is new and should still be considered a beta
	// feature.
	// The API may change in ways that are not backwards compatible.
	//
	// If the task identifier is not found, a 404 response code indicates that there
	// are no resources that match the request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	Get tasks_get.NewGet
	// Get all tasks.
	// Get information about the tasks currently running on one or more nodes in the
	// cluster.
	//
	// WARNING: The task management API is new and should still be considered a beta
	// feature.
	// The API may change in ways that are not backwards compatible.
	//
	// **Identifying running tasks**
	//
	// The `X-Opaque-Id header`, when provided on the HTTP request header, is going
	// to be returned as a header in the response as well as in the headers field
	// for in the task information.
	// This enables you to track certain calls or associate certain tasks with the
	// client that started them.
	// For example:
	//
	// ```
	// curl -i -H "X-Opaque-Id: 123456"
	// "http://localhost:9200/_tasks?group_by=parents"
	// ```
	//
	// The API returns the following result:
	//
	// ```
	// HTTP/1.1 200 OK
	// X-Opaque-Id: 123456
	// content-type: application/json; charset=UTF-8
	// content-length: 831
	//
	// {
	//   "tasks" : {
	//     "u5lcZHqcQhu-rUoFaqDphA:45" : {
	//       "node" : "u5lcZHqcQhu-rUoFaqDphA",
	//       "id" : 45,
	//       "type" : "transport",
	//       "action" : "cluster:monitor/tasks/lists",
	//       "start_time_in_millis" : 1513823752749,
	//       "running_time_in_nanos" : 293139,
	//       "cancellable" : false,
	//       "headers" : {
	//         "X-Opaque-Id" : "123456"
	//       },
	//       "children" : [
	//         {
	//           "node" : "u5lcZHqcQhu-rUoFaqDphA",
	//           "id" : 46,
	//           "type" : "direct",
	//           "action" : "cluster:monitor/tasks/lists[n]",
	//           "start_time_in_millis" : 1513823752750,
	//           "running_time_in_nanos" : 92133,
	//           "cancellable" : false,
	//           "parent_task_id" : "u5lcZHqcQhu-rUoFaqDphA:45",
	//           "headers" : {
	//             "X-Opaque-Id" : "123456"
	//           }
	//         }
	//       ]
	//     }
	//   }
	//  }
	// ```
	// In this example, `X-Opaque-Id: 123456` is the ID as a part of the response
	// header.
	// The `X-Opaque-Id` in the task `headers` is the ID for the task that was
	// initiated by the REST request.
	// The `X-Opaque-Id` in the children `headers` is the child task of the task
	// that was initiated by the REST request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	List tasks_list.NewList
}

type TextStructure struct {
	// Find the structure of a text field.
	// Find the structure of a text field in an Elasticsearch index.
	//
	// This API provides a starting point for extracting further information from
	// log messages already ingested into Elasticsearch.
	// For example, if you have ingested data into a very simple index that has just
	// `@timestamp` and message fields, you can use this API to see what common
	// structure exists in the message field.
	//
	// The response from the API contains:
	//
	// * Sample messages.
	// * Statistics that reveal the most common values for all fields detected
	// within the text and basic numeric statistics for numeric fields.
	// * Information about the structure of the text, which is useful when you write
	// ingest configurations to index it or similarly formatted text.
	// * Appropriate mappings for an Elasticsearch index, which you could use to
	// ingest the text.
	//
	// All this information can be calculated by the structure finder with no
	// guidance.
	// However, you can optionally override some of the decisions about the text
	// structure by specifying one or more query parameters.
	//
	// If the structure finder produces unexpected results, specify the `explain`
	// query parameter and an explanation will appear in the response.
	// It helps determine why the returned structure was chosen.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-field-structure.html
	FindFieldStructure text_structure_find_field_structure.NewFindFieldStructure
	// Find the structure of text messages.
	// Find the structure of a list of text messages.
	// The messages must contain data that is suitable to be ingested into
	// Elasticsearch.
	//
	// This API provides a starting point for ingesting data into Elasticsearch in a
	// format that is suitable for subsequent use with other Elastic Stack
	// functionality.
	// Use this API rather than the find text structure API if your input text has
	// already been split up into separate messages by some other process.
	//
	// The response from the API contains:
	//
	// * Sample messages.
	// * Statistics that reveal the most common values for all fields detected
	// within the text and basic numeric statistics for numeric fields.
	// * Information about the structure of the text, which is useful when you write
	// ingest configurations to index it or similarly formatted text.
	// Appropriate mappings for an Elasticsearch index, which you could use to
	// ingest the text.
	//
	// All this information can be calculated by the structure finder with no
	// guidance.
	// However, you can optionally override some of the decisions about the text
	// structure by specifying one or more query parameters.
	//
	// If the structure finder produces unexpected results, specify the `explain`
	// query parameter and an explanation will appear in the response.
	// It helps determine why the returned structure was chosen.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-message-structure.html
	FindMessageStructure text_structure_find_message_structure.NewFindMessageStructure
	// Find the structure of a text file.
	// The text file must contain data that is suitable to be ingested into
	// Elasticsearch.
	//
	// This API provides a starting point for ingesting data into Elasticsearch in a
	// format that is suitable for subsequent use with other Elastic Stack
	// functionality.
	// Unlike other Elasticsearch endpoints, the data that is posted to this
	// endpoint does not need to be UTF-8 encoded and in JSON format.
	// It must, however, be text; binary text formats are not currently supported.
	// The size is limited to the Elasticsearch HTTP receive buffer size, which
	// defaults to 100 Mb.
	//
	// The response from the API contains:
	//
	// * A couple of messages from the beginning of the text.
	// * Statistics that reveal the most common values for all fields detected
	// within the text and basic numeric statistics for numeric fields.
	// * Information about the structure of the text, which is useful when you write
	// ingest configurations to index it or similarly formatted text.
	// * Appropriate mappings for an Elasticsearch index, which you could use to
	// ingest the text.
	//
	// All this information can be calculated by the structure finder with no
	// guidance.
	// However, you can optionally override some of the decisions about the text
	// structure by specifying one or more query parameters.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-structure.html
	FindStructure text_structure_find_structure.NewFindStructure
	// Test a Grok pattern.
	// Test a Grok pattern on one or more lines of text.
	// The API indicates whether the lines match the pattern together with the
	// offsets and lengths of the matched substrings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/test-grok-pattern.html
	TestGrokPattern text_structure_test_grok_pattern.NewTestGrokPattern
}

type Transform struct {
	// Delete a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-transform.html
	DeleteTransform transform_delete_transform.NewDeleteTransform
	// Retrieves transform usage information for transform nodes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-node-stats.html
	GetNodeStats transform_get_node_stats.NewGetNodeStats
	// Get transforms.
	// Get configuration information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform.html
	GetTransform transform_get_transform.NewGetTransform
	// Get transform stats.
	//
	// Get usage information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-stats.html
	GetTransformStats transform_get_transform_stats.NewGetTransformStats
	// Preview a transform.
	// Generates a preview of the results that you will get when you create a
	// transform with the same configuration.
	//
	// It returns a maximum of 100 results. The calculations are based on all the
	// current data in the source index. It also
	// generates a list of mappings and settings for the destination index. These
	// values are determined based on the field
	// types of the source index and the transform aggregations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-transform.html
	PreviewTransform transform_preview_transform.NewPreviewTransform
	// Create a transform.
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
	// Reset a transform.
	//
	// Before you can reset it, you must stop it; alternatively, use the `force`
	// query parameter.
	// If the destination index was created by the transform, it is deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-transform.html
	ResetTransform transform_reset_transform.NewResetTransform
	// Schedule a transform to start now.
	//
	// Instantly run a transform to process data.
	// If you run this API, the transform will process the new data instantly,
	// without waiting for the configured frequency interval. After the API is
	// called,
	// the transform will be processed again at `now + frequency` unless the API
	// is called again in the meantime.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/schedule-now-transform.html
	ScheduleNowTransform transform_schedule_now_transform.NewScheduleNowTransform
	// Start a transform.
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
	// Stop transforms.
	// Stops one or more transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-transform.html
	StopTransform transform_stop_transform.NewStopTransform
	// Update a transform.
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
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-transform-update-transform
	UpdateTransform transform_update_transform.NewUpdateTransform
	// Upgrade all transforms.
	//
	// Transforms are compatible across minor versions and between supported major
	// versions.
	// However, over time, the format of transform configuration information may
	// change.
	// This API identifies transforms that have a legacy configuration format and
	// upgrades them to the latest version.
	// It also cleans up the internal data structures that store the transform state
	// and checkpoints.
	// The upgrade does not affect the source and destination indices.
	// The upgrade also does not affect the roles that transforms use when
	// Elasticsearch security features are enabled; the role used to read source
	// data and write to the destination index remains unchanged.
	//
	// If a transform upgrade step fails, the upgrade stops and an error is returned
	// about the underlying issue.
	// Resolve the issue then re-run the process again.
	// A summary is returned when the upgrade is finished.
	//
	// To ensure continuous transforms remain running during a major version upgrade
	// of the cluster – for example, from 7.16 to 8.0 – it is recommended to upgrade
	// transforms before upgrading the cluster.
	// You may want to perform a recent cluster backup prior to the upgrade.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-transform-upgrade-transforms
	UpgradeTransforms transform_upgrade_transforms.NewUpgradeTransforms
}

type Watcher struct {
	// Acknowledge a watch.
	// Acknowledging a watch enables you to manually throttle the execution of the
	// watch's actions.
	//
	// The acknowledgement state of an action is stored in the
	// `status.actions.<id>.ack.state` structure.
	//
	// IMPORTANT: If the specified watch is currently being executed, this API will
	// return an error
	// The reason for this behavior is to prevent overwriting the watch status from
	// a watch execution.
	//
	// Acknowledging an action throttles further executions of that action until its
	// `ack.state` is reset to `awaits_successful_execution`.
	// This happens when the condition of the watch is not met (the condition
	// evaluates to false).
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-ack-watch
	AckWatch watcher_ack_watch.NewAckWatch
	// Activate a watch.
	// A watch can be either active or inactive.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-activate-watch
	ActivateWatch watcher_activate_watch.NewActivateWatch
	// Deactivate a watch.
	// A watch can be either active or inactive.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-deactivate-watch
	DeactivateWatch watcher_deactivate_watch.NewDeactivateWatch
	// Delete a watch.
	// When the watch is removed, the document representing the watch in the
	// `.watches` index is gone and it will never be run again.
	//
	// Deleting a watch does not delete any watch execution records related to this
	// watch from the watch history.
	//
	// IMPORTANT: Deleting a watch must be done by using only this API.
	// Do not delete the watch directly from the `.watches` index using the
	// Elasticsearch delete document API
	// When Elasticsearch security features are enabled, make sure no write
	// privileges are granted to anyone for the `.watches` index.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-delete-watch
	DeleteWatch watcher_delete_watch.NewDeleteWatch
	// Run a watch.
	// This API can be used to force execution of the watch outside of its
	// triggering logic or to simulate the watch execution for debugging purposes.
	//
	// For testing and debugging purposes, you also have fine-grained control on how
	// the watch runs.
	// You can run the watch without running all of its actions or alternatively by
	// simulating them.
	// You can also force execution by ignoring the watch condition and control
	// whether a watch record would be written to the watch history after it runs.
	//
	// You can use the run watch API to run watches that are not yet registered by
	// specifying the watch definition inline.
	// This serves as great tool for testing and debugging your watches prior to
	// adding them to Watcher.
	//
	// When Elasticsearch security features are enabled on your cluster, watches are
	// run with the privileges of the user that stored the watches.
	// If your user is allowed to read index `a`, but not index `b`, then the exact
	// same set of rules will apply during execution of a watch.
	//
	// When using the run watch API, the authorization data of the user that called
	// the API will be used as a base, instead of the information who stored the
	// watch.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-execute-watch
	ExecuteWatch watcher_execute_watch.NewExecuteWatch
	// Get Watcher index settings.
	// Get settings for the Watcher internal index (`.watches`).
	// Only a subset of settings are shown, for example `index.auto_expand_replicas`
	// and `index.number_of_replicas`.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-get-settings
	GetSettings watcher_get_settings.NewGetSettings
	// Get a watch.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-get-watch
	GetWatch watcher_get_watch.NewGetWatch
	// Create or update a watch.
	// When a watch is registered, a new document that represents the watch is added
	// to the `.watches` index and its trigger is immediately registered with the
	// relevant trigger engine.
	// Typically for the `schedule` trigger, the scheduler is the trigger engine.
	//
	// IMPORTANT: You must use Kibana or this API to create a watch.
	// Do not add a watch directly to the `.watches` index by using the
	// Elasticsearch index API.
	// If Elasticsearch security features are enabled, do not give users write
	// privileges on the `.watches` index.
	//
	// When you add a watch you can also define its initial active state by setting
	// the *active* parameter.
	//
	// When Elasticsearch security features are enabled, your watch can index or
	// search only on indices for which the user that stored the watch has
	// privileges.
	// If the user is able to read index `a`, but not index `b`, the same will apply
	// when the watch runs.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-put-watch
	PutWatch watcher_put_watch.NewPutWatch
	// Query watches.
	// Get all registered watches in a paginated manner and optionally filter
	// watches by a query.
	//
	// Note that only the `_id` and `metadata.*` fields are queryable or sortable.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-query-watches
	QueryWatches watcher_query_watches.NewQueryWatches
	// Start the watch service.
	// Start the Watcher service if it is not already running.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-start
	Start watcher_start.NewStart
	// Get Watcher statistics.
	// This API always returns basic metrics.
	// You retrieve more metrics by using the metric parameter.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-stats
	Stats watcher_stats.NewStats
	// Stop the watch service.
	// Stop the Watcher service if it is running.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-stop
	Stop watcher_stop.NewStop
	// Update Watcher index settings.
	// Update settings for the Watcher internal index (`.watches`).
	// Only a subset of settings can be modified.
	// This includes `index.auto_expand_replicas`, `index.number_of_replicas`,
	// `index.routing.allocation.exclude.*`,
	// `index.routing.allocation.include.*` and
	// `index.routing.allocation.require.*`.
	// Modification of `index.routing.allocation.include._tier_preference` is an
	// exception and is not allowed as the
	// Watcher shards must always be in the `data_content` tier.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-update-settings
	UpdateSettings watcher_update_settings.NewUpdateSettings
}

type Xpack struct {
	// Get information.
	// The information provided by the API includes:
	//
	// * Build information including the build number and timestamp.
	// * License information about the currently installed license.
	// * Feature information for the features that are currently enabled and
	// available under the current license.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html
	Info xpack_info.NewInfo
	// Get usage information.
	// Get information about the features that are currently enabled and available
	// under the current license.
	// The API also provides some usage statistics.
	// https://www.elastic.co/docs/api/doc/elasticsearch/v8/group/endpoint-xpack
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
	Simulate            Simulate
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

	// Bulk index or delete documents.
	// Perform multiple `index`, `create`, `delete`, and `update` actions in a
	// single request.
	// This reduces overhead and can greatly increase indexing speed.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or index alias:
	//
	// * To use the `create` action, you must have the `create_doc`, `create`,
	// `index`, or `write` index privilege. Data streams support only the `create`
	// action.
	// * To use the `index` action, you must have the `create`, `index`, or `write`
	// index privilege.
	// * To use the `delete` action, you must have the `delete` or `write` index
	// privilege.
	// * To use the `update` action, you must have the `index` or `write` index
	// privilege.
	// * To automatically create a data stream or index with a bulk API request, you
	// must have the `auto_configure`, `create_index`, or `manage` index privilege.
	// * To make the result of a bulk operation visible to search using the
	// `refresh` parameter, you must have the `maintenance` or `manage` index
	// privilege.
	//
	// Automatic data stream creation requires a matching index template with data
	// stream enabled.
	//
	// The actions are specified in the request body using a newline delimited JSON
	// (NDJSON) structure:
	//
	// ```
	// action_and_meta_data\n
	// optional_source\n
	// action_and_meta_data\n
	// optional_source\n
	// ....
	// action_and_meta_data\n
	// optional_source\n
	// ```
	//
	// The `index` and `create` actions expect a source on the next line and have
	// the same semantics as the `op_type` parameter in the standard index API.
	// A `create` action fails if a document with the same ID already exists in the
	// target
	// An `index` action adds or replaces a document as necessary.
	//
	// NOTE: Data streams support only the `create` action.
	// To update or delete a document in a data stream, you must target the backing
	// index containing the document.
	//
	// An `update` action expects that the partial doc, upsert, and script and its
	// options are specified on the next line.
	//
	// A `delete` action does not expect a source on the next line and has the same
	// semantics as the standard delete API.
	//
	// NOTE: The final line of data must end with a newline character (`\n`).
	// Each newline character may be preceded by a carriage return (`\r`).
	// When sending NDJSON data to the `_bulk` endpoint, use a `Content-Type` header
	// of `application/json` or `application/x-ndjson`.
	// Because this format uses literal newline characters (`\n`) as delimiters,
	// make sure that the JSON actions and sources are not pretty printed.
	//
	// If you provide a target in the request path, it is used for any actions that
	// don't explicitly specify an `_index` argument.
	//
	// A note on the format: the idea here is to make processing as fast as
	// possible.
	// As some of the actions are redirected to other shards on other nodes, only
	// `action_meta_data` is parsed on the receiving node side.
	//
	// Client libraries using this protocol should try and strive to do something
	// similar on the client side, and reduce buffering as much as possible.
	//
	// There is no "correct" number of actions to perform in a single bulk request.
	// Experiment with different settings to find the optimal size for your
	// particular workload.
	// Note that Elasticsearch limits the maximum size of a HTTP request to 100mb by
	// default so clients must ensure that no request exceeds this size.
	// It is not possible to index a single document that exceeds the size limit, so
	// you must pre-process any such documents into smaller pieces before sending
	// them to Elasticsearch.
	// For instance, split documents into pages or chapters before indexing them, or
	// store raw binary data in a system outside Elasticsearch and replace the raw
	// data with a link to the external system in the documents that you send to
	// Elasticsearch.
	//
	// **Client suppport for bulk requests**
	//
	// Some of the officially supported clients provide helpers to assist with bulk
	// requests and reindexing:
	//
	// * Go: Check out `esutil.BulkIndexer`
	// * Perl: Check out `Search::Elasticsearch::Client::5_0::Bulk` and
	// `Search::Elasticsearch::Client::5_0::Scroll`
	// * Python: Check out `elasticsearch.helpers.*`
	// * JavaScript: Check out `client.helpers.*`
	// * .NET: Check out `BulkAllObservable`
	// * PHP: Check out bulk indexing.
	//
	// **Submitting bulk requests with cURL**
	//
	// If you're providing text file input to `curl`, you must use the
	// `--data-binary` flag instead of plain `-d`.
	// The latter doesn't preserve newlines. For example:
	//
	// ```
	// $ cat requests
	// { "index" : { "_index" : "test", "_id" : "1" } }
	// { "field1" : "value1" }
	// $ curl -s -H "Content-Type: application/x-ndjson" -XPOST localhost:9200/_bulk
	// --data-binary "@requests"; echo
	// {"took":7, "errors": false,
	// "items":[{"index":{"_index":"test","_id":"1","_version":1,"result":"created","forced_refresh":false}}]}
	// ```
	//
	// **Optimistic concurrency control**
	//
	// Each `index` and `delete` action within a bulk API call may include the
	// `if_seq_no` and `if_primary_term` parameters in their respective action and
	// meta data lines.
	// The `if_seq_no` and `if_primary_term` parameters control how operations are
	// run, based on the last modification to existing documents. See Optimistic
	// concurrency control for more details.
	//
	// **Versioning**
	//
	// Each bulk item can include the version value using the `version` field.
	// It automatically follows the behavior of the index or delete operation based
	// on the `_version` mapping.
	// It also support the `version_type`.
	//
	// **Routing**
	//
	// Each bulk item can include the routing value using the `routing` field.
	// It automatically follows the behavior of the index or delete operation based
	// on the `_routing` mapping.
	//
	// NOTE: Data streams do not support custom routing unless they were created
	// with the `allow_custom_routing` setting enabled in the template.
	//
	// **Wait for active shards**
	//
	// When making bulk calls, you can set the `wait_for_active_shards` parameter to
	// require a minimum number of shard copies to be active before starting to
	// process the bulk request.
	//
	// **Refresh**
	//
	// Control when the changes made by this request are visible to search.
	//
	// NOTE: Only the shards that receive the bulk request will be affected by
	// refresh.
	// Imagine a `_bulk?refresh=wait_for` request with three documents in it that
	// happen to be routed to different shards in an index with five shards.
	// The request will only wait for those three shards to refresh.
	// The other two shards that make up the index do not participate in the `_bulk`
	// request at all.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Clear a scrolling search.
	// Clear the search context and results for a scrolling search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Close a point in time.
	// A point in time must be opened explicitly before being used in search
	// requests.
	// The `keep_alive` parameter tells Elasticsearch how long it should persist.
	// A point in time is automatically closed when the `keep_alive` period has
	// elapsed.
	// However, keeping points in time has a cost; close them as soon as they are no
	// longer required for search requests.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Count search results.
	// Get the number of documents matching a query.
	//
	// The query can be provided either by using a simple query string as a
	// parameter, or by defining Query DSL within the request body.
	// The query is optional. When no query is provided, the API uses `match_all` to
	// count all the documents.
	//
	// The count API supports multi-target syntax. You can run a single count API
	// search across multiple data streams and indices.
	//
	// The operation is broadcast across all shards.
	// For each shard ID group, a replica is chosen and the search is run against
	// it.
	// This means that replicas increase the scalability of the count.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Create a new document in the index.
	//
	// You can index a new JSON document with the `/<target>/_doc/` or
	// `/<target>/_create/<_id>` APIs
	// Using `_create` guarantees that the document is indexed only if it does not
	// already exist.
	// It returns a 409 response when a document with a same ID already exists in
	// the index.
	// To update an existing document, you must use the `/<target>/_doc/` API.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or index alias:
	//
	// * To add a document using the `PUT /<target>/_create/<_id>` or `POST
	// /<target>/_create/<_id>` request formats, you must have the `create_doc`,
	// `create`, `index`, or `write` index privilege.
	// * To automatically create a data stream or index with this API request, you
	// must have the `auto_configure`, `create_index`, or `manage` index privilege.
	//
	// Automatic data stream creation requires a matching index template with data
	// stream enabled.
	//
	// **Automatically create data streams and indices**
	//
	// If the request's target doesn't exist and matches an index template with a
	// `data_stream` definition, the index operation automatically creates the data
	// stream.
	//
	// If the target doesn't exist and doesn't match a data stream template, the
	// operation automatically creates the index and applies any matching index
	// templates.
	//
	// NOTE: Elasticsearch includes several built-in index templates. To avoid
	// naming collisions with these templates, refer to index pattern documentation.
	//
	// If no mapping exists, the index operation creates a dynamic mapping.
	// By default, new fields and objects are automatically added to the mapping if
	// needed.
	//
	// Automatic index creation is controlled by the `action.auto_create_index`
	// setting.
	// If it is `true`, any index can be created automatically.
	// You can modify this setting to explicitly allow or block automatic creation
	// of indices that match specified patterns or set it to `false` to turn off
	// automatic index creation entirely.
	// Specify a comma-separated list of patterns you want to allow or prefix each
	// pattern with `+` or `-` to indicate whether it should be allowed or blocked.
	// When a list is specified, the default behaviour is to disallow.
	//
	// NOTE: The `action.auto_create_index` setting affects the automatic creation
	// of indices only.
	// It does not affect the creation of data streams.
	//
	// **Routing**
	//
	// By default, shard placement — or routing — is controlled by using a hash of
	// the document's ID value.
	// For more explicit control, the value fed into the hash function used by the
	// router can be directly specified on a per-operation basis using the `routing`
	// parameter.
	//
	// When setting up explicit mapping, you can also use the `_routing` field to
	// direct the index operation to extract the routing value from the document
	// itself.
	// This does come at the (very minimal) cost of an additional document parsing
	// pass.
	// If the `_routing` mapping is defined and set to be required, the index
	// operation will fail if no routing value is provided or extracted.
	//
	// NOTE: Data streams do not support custom routing unless they were created
	// with the `allow_custom_routing` setting enabled in the template.
	//
	// **Distributed**
	//
	// The index operation is directed to the primary shard based on its route and
	// performed on the actual node containing this shard.
	// After the primary shard completes the operation, if needed, the update is
	// distributed to applicable replicas.
	//
	// **Active shards**
	//
	// To improve the resiliency of writes to the system, indexing operations can be
	// configured to wait for a certain number of active shard copies before
	// proceeding with the operation.
	// If the requisite number of active shard copies are not available, then the
	// write operation must wait and retry, until either the requisite shard copies
	// have started or a timeout occurs.
	// By default, write operations only wait for the primary shards to be active
	// before proceeding (that is to say `wait_for_active_shards` is `1`).
	// This default can be overridden in the index settings dynamically by setting
	// `index.write.wait_for_active_shards`.
	// To alter this behavior per operation, use the `wait_for_active_shards
	// request` parameter.
	//
	// Valid values are all or any positive integer up to the total number of
	// configured copies per shard in the index (which is `number_of_replicas`+1).
	// Specifying a negative value or a number greater than the number of shard
	// copies will throw an error.
	//
	// For example, suppose you have a cluster of three nodes, A, B, and C and you
	// create an index index with the number of replicas set to 3 (resulting in 4
	// shard copies, one more copy than there are nodes).
	// If you attempt an indexing operation, by default the operation will only
	// ensure the primary copy of each shard is available before proceeding.
	// This means that even if B and C went down and A hosted the primary shard
	// copies, the indexing operation would still proceed with only one copy of the
	// data.
	// If `wait_for_active_shards` is set on the request to `3` (and all three nodes
	// are up), the indexing operation will require 3 active shard copies before
	// proceeding.
	// This requirement should be met because there are 3 active nodes in the
	// cluster, each one holding a copy of the shard.
	// However, if you set `wait_for_active_shards` to `all` (or to `4`, which is
	// the same in this situation), the indexing operation will not proceed as you
	// do not have all 4 copies of each shard active in the index.
	// The operation will timeout unless a new node is brought up in the cluster to
	// host the fourth copy of the shard.
	//
	// It is important to note that this setting greatly reduces the chances of the
	// write operation not writing to the requisite number of shard copies, but it
	// does not completely eliminate the possibility, because this check occurs
	// before the write operation starts.
	// After the write operation is underway, it is still possible for replication
	// to fail on any number of shard copies but still succeed on the primary.
	// The `_shards` section of the API response reveals the number of shard copies
	// on which replication succeeded and failed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Delete a document.
	//
	// Remove a JSON document from the specified index.
	//
	// NOTE: You cannot send deletion requests directly to a data stream.
	// To delete a document in a data stream, you must target the backing index
	// containing the document.
	//
	// **Optimistic concurrency control**
	//
	// Delete operations can be made conditional and only be performed if the last
	// modification to the document was assigned the sequence number and primary
	// term specified by the `if_seq_no` and `if_primary_term` parameters.
	// If a mismatch is detected, the operation will result in a
	// `VersionConflictException` and a status code of `409`.
	//
	// **Versioning**
	//
	// Each document indexed is versioned.
	// When deleting a document, the version can be specified to make sure the
	// relevant document you are trying to delete is actually being deleted and it
	// has not changed in the meantime.
	// Every write operation run on a document, deletes included, causes its version
	// to be incremented.
	// The version number of a deleted document remains available for a short time
	// after deletion to allow for control of concurrent operations.
	// The length of time for which a deleted document's version remains available
	// is determined by the `index.gc_deletes` index setting.
	//
	// **Routing**
	//
	// If routing is used during indexing, the routing value also needs to be
	// specified to delete a document.
	//
	// If the `_routing` mapping is set to `required` and no routing value is
	// specified, the delete API throws a `RoutingMissingException` and rejects the
	// request.
	//
	// For example:
	//
	// ```
	// DELETE /my-index-000001/_doc/1?routing=shard-1
	// ```
	//
	// This request deletes the document with ID 1, but it is routed based on the
	// user.
	// The document is not deleted if the correct routing is not specified.
	//
	// **Distributed**
	//
	// The delete operation gets hashed into a specific shard ID.
	// It then gets redirected into the primary shard within that ID group and
	// replicated (if needed) to shard replicas within that ID group.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Delete documents.
	//
	// Deletes documents that match the specified query.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or alias:
	//
	// * `read`
	// * `delete` or `write`
	//
	// You can specify the query criteria in the request URI or the request body
	// using the same syntax as the search API.
	// When you submit a delete by query request, Elasticsearch gets a snapshot of
	// the data stream or index when it begins processing the request and deletes
	// matching documents using internal versioning.
	// If a document changes between the time that the snapshot is taken and the
	// delete operation is processed, it results in a version conflict and the
	// delete operation fails.
	//
	// NOTE: Documents with a version equal to 0 cannot be deleted using delete by
	// query because internal versioning does not support 0 as a valid version
	// number.
	//
	// While processing a delete by query request, Elasticsearch performs multiple
	// search requests sequentially to find all of the matching documents to delete.
	// A bulk delete request is performed for each batch of matching documents.
	// If a search or bulk request is rejected, the requests are retried up to 10
	// times, with exponential back off.
	// If the maximum retry limit is reached, processing halts and all failed
	// requests are returned in the response.
	// Any delete requests that completed successfully still stick, they are not
	// rolled back.
	//
	// You can opt to count version conflicts instead of halting and returning by
	// setting `conflicts` to `proceed`.
	// Note that if you opt to count version conflicts the operation could attempt
	// to delete more documents from the source than `max_docs` until it has
	// successfully deleted `max_docs documents`, or it has gone through every
	// document in the source query.
	//
	// **Throttling delete requests**
	//
	// To control the rate at which delete by query issues batches of delete
	// operations, you can set `requests_per_second` to any positive decimal number.
	// This pads each batch with a wait time to throttle the rate.
	// Set `requests_per_second` to `-1` to disable throttling.
	//
	// Throttling uses a wait time between batches so that the internal scroll
	// requests can be given a timeout that takes the request padding into account.
	// The padding time is the difference between the batch size divided by the
	// `requests_per_second` and the time spent writing.
	// By default the batch size is `1000`, so if `requests_per_second` is set to
	// `500`:
	//
	// ```
	// target_time = 1000 / 500 per second = 2 seconds
	// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
	// ```
	//
	// Since the batch is issued as a single `_bulk` request, large batch sizes
	// cause Elasticsearch to create many requests and wait before starting the next
	// set.
	// This is "bursty" instead of "smooth".
	//
	// **Slicing**
	//
	// Delete by query supports sliced scroll to parallelize the delete process.
	// This can improve efficiency and provide a convenient way to break the request
	// down into smaller parts.
	//
	// Setting `slices` to `auto` lets Elasticsearch choose the number of slices to
	// use.
	// This setting will use one slice per shard, up to a certain limit.
	// If there are multiple source data streams or indices, it will choose the
	// number of slices based on the index or backing index with the smallest number
	// of shards.
	// Adding slices to the delete by query operation creates sub-requests which
	// means it has some quirks:
	//
	// * You can see these requests in the tasks APIs. These sub-requests are
	// "child" tasks of the task for the request with slices.
	// * Fetching the status of the task for the request with slices only contains
	// the status of completed slices.
	// * These sub-requests are individually addressable for things like
	// cancellation and rethrottling.
	// * Rethrottling the request with `slices` will rethrottle the unfinished
	// sub-request proportionally.
	// * Canceling the request with `slices` will cancel each sub-request.
	// * Due to the nature of `slices` each sub-request won't get a perfectly even
	// portion of the documents. All documents will be addressed, but some slices
	// may be larger than others. Expect larger slices to have a more even
	// distribution.
	// * Parameters like `requests_per_second` and `max_docs` on a request with
	// `slices` are distributed proportionally to each sub-request. Combine that
	// with the earlier point about distribution being uneven and you should
	// conclude that using `max_docs` with `slices` might not result in exactly
	// `max_docs` documents being deleted.
	// * Each sub-request gets a slightly different snapshot of the source data
	// stream or index though these are all taken at approximately the same time.
	//
	// If you're slicing manually or otherwise tuning automatic slicing, keep in
	// mind that:
	//
	// * Query performance is most efficient when the number of slices is equal to
	// the number of shards in the index or backing index. If that number is large
	// (for example, 500), choose a lower number as too many `slices` hurts
	// performance. Setting `slices` higher than the number of shards generally does
	// not improve efficiency and adds overhead.
	// * Delete performance scales linearly across available resources with the
	// number of slices.
	//
	// Whether query or delete performance dominates the runtime depends on the
	// documents being reindexed and cluster resources.
	//
	// **Cancel a delete by query operation**
	//
	// Any delete by query can be canceled using the task cancel API. For example:
	//
	// ```
	// POST _tasks/r1A2WoRbTwKZ516z6NEs5A:36619/_cancel
	// ```
	//
	// The task ID can be found by using the get tasks API.
	//
	// Cancellation should happen quickly but might take a few seconds.
	// The get task status API will continue to list the delete by query task until
	// this task checks that it has been cancelled and terminates itself.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Throttle a delete by query operation.
	//
	// Change the number of requests per second for a particular delete by query
	// operation.
	// Rethrottling that speeds up the query takes effect immediately but
	// rethrotting that slows down the query takes effect after completing the
	// current batch to prevent scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html#docs-delete-by-query-rethrottle
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	// Delete a script or search template.
	// Deletes a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-stored-script-api.html
	DeleteScript core_delete_script.NewDeleteScript
	// Check a document.
	//
	// Verify that a document exists.
	// For example, check to see if a document with the `_id` 0 exists:
	//
	// ```
	// HEAD my-index-000001/_doc/0
	// ```
	//
	// If the document exists, the API returns a status code of `200 - OK`.
	// If the document doesn’t exist, the API returns `404 - Not Found`.
	//
	// **Versioning support**
	//
	// You can use the `version` parameter to check the document only if its current
	// version is equal to the specified one.
	//
	// Internally, Elasticsearch has marked the old document as deleted and added an
	// entirely new document.
	// The old version of the document doesn't disappear immediately, although you
	// won't be able to access it.
	// Elasticsearch cleans up deleted documents in the background as you continue
	// to index more data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Check for a document source.
	//
	// Check whether a document source exists in an index.
	// For example:
	//
	// ```
	// HEAD my-index-000001/_source/1
	// ```
	//
	// A document's source is not available if it is disabled in the mapping.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Explain a document match result.
	// Get information about why a specific document matches, or doesn't match, a
	// query.
	// It computes a score explanation for a query and a specific document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// Get the field capabilities.
	//
	// Get information about the capabilities of fields among multiple indices.
	//
	// For data streams, the API returns field capabilities among the stream’s
	// backing indices.
	// It returns runtime fields like any other field.
	// For example, a runtime field with a type of keyword is returned the same as
	// any other field that belongs to the `keyword` family.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Get a document by its ID.
	//
	// Get a document and its source or stored fields from an index.
	//
	// By default, this API is realtime and is not affected by the refresh rate of
	// the index (when data will become visible for search).
	// In the case where stored fields are requested with the `stored_fields`
	// parameter and the document has been updated but is not yet refreshed, the API
	// will have to parse and analyze the source to extract the stored fields.
	// To turn off realtime behavior, set the `realtime` parameter to false.
	//
	// **Source filtering**
	//
	// By default, the API returns the contents of the `_source` field unless you
	// have used the `stored_fields` parameter or the `_source` field is turned off.
	// You can turn off `_source` retrieval by using the `_source` parameter:
	//
	// ```
	// GET my-index-000001/_doc/0?_source=false
	// ```
	//
	// If you only need one or two fields from the `_source`, use the
	// `_source_includes` or `_source_excludes` parameters to include or filter out
	// particular fields.
	// This can be helpful with large documents where partial retrieval can save on
	// network overhead
	// Both parameters take a comma separated list of fields or wildcard
	// expressions.
	// For example:
	//
	// ```
	// GET my-index-000001/_doc/0?_source_includes=*.id&_source_excludes=entities
	// ```
	//
	// If you only want to specify includes, you can use a shorter notation:
	//
	// ```
	// GET my-index-000001/_doc/0?_source=*.id
	// ```
	//
	// **Routing**
	//
	// If routing is used during indexing, the routing value also needs to be
	// specified to retrieve a document.
	// For example:
	//
	// ```
	// GET my-index-000001/_doc/2?routing=user1
	// ```
	//
	// This request gets the document with ID 2, but it is routed based on the user.
	// The document is not fetched if the correct routing is not specified.
	//
	// **Distributed**
	//
	// The GET operation is hashed into a specific shard ID.
	// It is then redirected to one of the replicas within that shard ID and returns
	// the result.
	// The replicas are the primary shard and its replicas within that shard ID
	// group.
	// This means that the more replicas you have, the better your GET scaling will
	// be.
	//
	// **Versioning support**
	//
	// You can use the `version` parameter to retrieve the document only if its
	// current version is equal to the specified one.
	//
	// Internally, Elasticsearch has marked the old document as deleted and added an
	// entirely new document.
	// The old version of the document doesn't disappear immediately, although you
	// won't be able to access it.
	// Elasticsearch cleans up deleted documents in the background as you continue
	// to index more data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Get a script or search template.
	// Retrieves a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-stored-script-api.html
	GetScript core_get_script.NewGetScript
	// Get script contexts.
	//
	// Get a list of supported script contexts and their methods.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-script-contexts-api.html
	GetScriptContext core_get_script_context.NewGetScriptContext
	// Get script languages.
	//
	// Get a list of available script types, languages, and contexts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-script-languages-api.html
	GetScriptLanguages core_get_script_languages.NewGetScriptLanguages
	// Get a document's source.
	//
	// Get the source of a document.
	// For example:
	//
	// ```
	// GET my-index-000001/_source/1
	// ```
	//
	// You can use the source filtering parameters to control which parts of the
	// `_source` are returned:
	//
	// ```
	// GET
	// my-index-000001/_source/1/?_source_includes=*.id&_source_excludes=entities
	// ```
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Get the cluster health.
	// Get a report with the health status of an Elasticsearch cluster.
	// The report contains a list of indicators that compose Elasticsearch
	// functionality.
	//
	// Each indicator has a health status of: green, unknown, yellow or red.
	// The indicator will provide an explanation and metadata describing the reason
	// for its current health status.
	//
	// The cluster’s status is controlled by the worst indicator status.
	//
	// In the event that an indicator’s status is non-green, a list of impacts may
	// be present in the indicator result which detail the functionalities that are
	// negatively affected by the health issue.
	// Each impact carries with it a severity level, an area of the system that is
	// affected, and a simple description of the impact on the system.
	//
	// Some health indicators can determine the root cause of a health problem and
	// prescribe a set of steps that can be performed in order to improve the health
	// of the system.
	// The root cause and remediation steps are encapsulated in a diagnosis.
	// A diagnosis contains a cause detailing a root cause analysis, an action
	// containing a brief description of the steps to take to fix the problem, the
	// list of affected resources (if applicable), and a detailed step-by-step
	// troubleshooting guide to fix the diagnosed problem.
	//
	// NOTE: The health indicators perform root cause analysis of non-green health
	// statuses. This can be computationally expensive when called frequently.
	// When setting up automated polling of the API for health status, set verbose
	// to false to disable the more expensive analysis logic.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html
	HealthReport core_health_report.NewHealthReport
	// Create or update a document in an index.
	//
	// Add a JSON document to the specified data stream or index and make it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	//
	// NOTE: You cannot use this API to send update requests for existing documents
	// in a data stream.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or index alias:
	//
	// * To add or overwrite a document using the `PUT /<target>/_doc/<_id>` request
	// format, you must have the `create`, `index`, or `write` index privilege.
	// * To add a document using the `POST /<target>/_doc/` request format, you must
	// have the `create_doc`, `create`, `index`, or `write` index privilege.
	// * To automatically create a data stream or index with this API request, you
	// must have the `auto_configure`, `create_index`, or `manage` index privilege.
	//
	// Automatic data stream creation requires a matching index template with data
	// stream enabled.
	//
	// NOTE: Replica shards might not all be started when an indexing operation
	// returns successfully.
	// By default, only the primary is required. Set `wait_for_active_shards` to
	// change this default behavior.
	//
	// **Automatically create data streams and indices**
	//
	// If the request's target doesn't exist and matches an index template with a
	// `data_stream` definition, the index operation automatically creates the data
	// stream.
	//
	// If the target doesn't exist and doesn't match a data stream template, the
	// operation automatically creates the index and applies any matching index
	// templates.
	//
	// NOTE: Elasticsearch includes several built-in index templates. To avoid
	// naming collisions with these templates, refer to index pattern documentation.
	//
	// If no mapping exists, the index operation creates a dynamic mapping.
	// By default, new fields and objects are automatically added to the mapping if
	// needed.
	//
	// Automatic index creation is controlled by the `action.auto_create_index`
	// setting.
	// If it is `true`, any index can be created automatically.
	// You can modify this setting to explicitly allow or block automatic creation
	// of indices that match specified patterns or set it to `false` to turn off
	// automatic index creation entirely.
	// Specify a comma-separated list of patterns you want to allow or prefix each
	// pattern with `+` or `-` to indicate whether it should be allowed or blocked.
	// When a list is specified, the default behaviour is to disallow.
	//
	// NOTE: The `action.auto_create_index` setting affects the automatic creation
	// of indices only.
	// It does not affect the creation of data streams.
	//
	// **Optimistic concurrency control**
	//
	// Index operations can be made conditional and only be performed if the last
	// modification to the document was assigned the sequence number and primary
	// term specified by the `if_seq_no` and `if_primary_term` parameters.
	// If a mismatch is detected, the operation will result in a
	// `VersionConflictException` and a status code of `409`.
	//
	// **Routing**
	//
	// By default, shard placement — or routing — is controlled by using a hash of
	// the document's ID value.
	// For more explicit control, the value fed into the hash function used by the
	// router can be directly specified on a per-operation basis using the `routing`
	// parameter.
	//
	// When setting up explicit mapping, you can also use the `_routing` field to
	// direct the index operation to extract the routing value from the document
	// itself.
	// This does come at the (very minimal) cost of an additional document parsing
	// pass.
	// If the `_routing` mapping is defined and set to be required, the index
	// operation will fail if no routing value is provided or extracted.
	//
	// NOTE: Data streams do not support custom routing unless they were created
	// with the `allow_custom_routing` setting enabled in the template.
	//
	// **Distributed**
	//
	// The index operation is directed to the primary shard based on its route and
	// performed on the actual node containing this shard.
	// After the primary shard completes the operation, if needed, the update is
	// distributed to applicable replicas.
	//
	// **Active shards**
	//
	// To improve the resiliency of writes to the system, indexing operations can be
	// configured to wait for a certain number of active shard copies before
	// proceeding with the operation.
	// If the requisite number of active shard copies are not available, then the
	// write operation must wait and retry, until either the requisite shard copies
	// have started or a timeout occurs.
	// By default, write operations only wait for the primary shards to be active
	// before proceeding (that is to say `wait_for_active_shards` is `1`).
	// This default can be overridden in the index settings dynamically by setting
	// `index.write.wait_for_active_shards`.
	// To alter this behavior per operation, use the `wait_for_active_shards
	// request` parameter.
	//
	// Valid values are all or any positive integer up to the total number of
	// configured copies per shard in the index (which is `number_of_replicas`+1).
	// Specifying a negative value or a number greater than the number of shard
	// copies will throw an error.
	//
	// For example, suppose you have a cluster of three nodes, A, B, and C and you
	// create an index index with the number of replicas set to 3 (resulting in 4
	// shard copies, one more copy than there are nodes).
	// If you attempt an indexing operation, by default the operation will only
	// ensure the primary copy of each shard is available before proceeding.
	// This means that even if B and C went down and A hosted the primary shard
	// copies, the indexing operation would still proceed with only one copy of the
	// data.
	// If `wait_for_active_shards` is set on the request to `3` (and all three nodes
	// are up), the indexing operation will require 3 active shard copies before
	// proceeding.
	// This requirement should be met because there are 3 active nodes in the
	// cluster, each one holding a copy of the shard.
	// However, if you set `wait_for_active_shards` to `all` (or to `4`, which is
	// the same in this situation), the indexing operation will not proceed as you
	// do not have all 4 copies of each shard active in the index.
	// The operation will timeout unless a new node is brought up in the cluster to
	// host the fourth copy of the shard.
	//
	// It is important to note that this setting greatly reduces the chances of the
	// write operation not writing to the requisite number of shard copies, but it
	// does not completely eliminate the possibility, because this check occurs
	// before the write operation starts.
	// After the write operation is underway, it is still possible for replication
	// to fail on any number of shard copies but still succeed on the primary.
	// The `_shards` section of the API response reveals the number of shard copies
	// on which replication succeeded and failed.
	//
	// **No operation (noop) updates**
	//
	// When updating a document by using this API, a new version of the document is
	// always created even if the document hasn't changed.
	// If this isn't acceptable use the `_update` API with `detect_noop` set to
	// `true`.
	// The `detect_noop` option isn't available on this API because it doesn’t fetch
	// the old source and isn't able to compare it against the new source.
	//
	// There isn't a definitive rule for when noop updates aren't acceptable.
	// It's a combination of lots of factors like how frequently your data source
	// sends updates that are actually noops and how many queries per second
	// Elasticsearch runs on the shard receiving the updates.
	//
	// **Versioning**
	//
	// Each indexed document is given a version number.
	// By default, internal versioning is used that starts at 1 and increments with
	// each update, deletes included.
	// Optionally, the version number can be set to an external value (for example,
	// if maintained in a database).
	// To enable this functionality, `version_type` should be set to `external`.
	// The value provided must be a numeric, long value greater than or equal to 0,
	// and less than around `9.2e+18`.
	//
	// NOTE: Versioning is completely real time, and is not affected by the near
	// real time aspects of search operations.
	// If no version is provided, the operation runs without any version checks.
	//
	// When using the external version type, the system checks to see if the version
	// number passed to the index request is greater than the version of the
	// currently stored document.
	// If true, the document will be indexed and the new version number used.
	// If the value provided is less than or equal to the stored document's version
	// number, a version conflict will occur and the index operation will fail. For
	// example:
	//
	// ```
	// PUT my-index-000001/_doc/1?version=2&version_type=external
	// {
	//   "user": {
	//     "id": "elkbee"
	//   }
	// }
	//
	// In this example, the operation will succeed since the supplied version of 2
	// is higher than the current document version of 1.
	// If the document was already updated and its version was set to 2 or higher,
	// the indexing command will fail and result in a conflict (409 HTTP status
	// code).
	//
	// A nice side effect is that there is no need to maintain strict ordering of
	// async indexing operations run as a result of changes to a source database, as
	// long as version numbers from the source database are used.
	// Even the simple case of updating the Elasticsearch index using data from a
	// database is simplified if external versioning is used, as only the latest
	// version will be used if the index operations arrive out of order.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Index core_index.NewIndex
	// Get cluster info.
	// Get basic build, version, and cluster information.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rest-api-root.html
	Info core_info.NewInfo
	// Run a knn search.
	//
	// NOTE: The kNN search API has been replaced by the `knn` option in the search
	// API.
	//
	// Perform a k-nearest neighbor (kNN) search on a dense_vector field and return
	// the matching documents.
	// Given a query vector, the API finds the k closest vectors and returns those
	// documents as search hits.
	//
	// Elasticsearch uses the HNSW algorithm to support efficient kNN search.
	// Like most kNN algorithms, HNSW is an approximate method that sacrifices
	// result accuracy for improved search speed.
	// This means the results returned are not always the true k closest neighbors.
	//
	// The kNN search API supports restricting the search using a filter.
	// The search will return the top k documents that also match the filter query.
	//
	// A kNN search response has the exact same structure as a search API response.
	// However, certain sections have a meaning specific to kNN search:
	//
	// * The document `_score` is determined by the similarity between the query and
	// document vector.
	// * The `hits.total` object contains the total number of nearest neighbor
	// candidates considered, which is `num_candidates * num_shards`. The
	// `hits.total.relation` will always be `eq`, indicating an exact value.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/knn-search-api.html
	KnnSearch core_knn_search.NewKnnSearch
	// Get multiple documents.
	//
	// Get multiple JSON documents by ID from one or more indices.
	// If you specify an index in the request URI, you only need to specify the
	// document IDs in the request body.
	// To ensure fast responses, this multi get (mget) API responds with partial
	// results if one or more shards fail.
	//
	// **Filter source fields**
	//
	// By default, the `_source` field is returned for every document (if stored).
	// Use the `_source` and `_source_include` or `source_exclude` attributes to
	// filter what fields are returned for a particular document.
	// You can include the `_source`, `_source_includes`, and `_source_excludes`
	// query parameters in the request URI to specify the defaults to use when there
	// are no per-document instructions.
	//
	// **Get stored fields**
	//
	// Use the `stored_fields` attribute to specify the set of stored fields you
	// want to retrieve.
	// Any requested fields that are not stored are ignored.
	// You can include the `stored_fields` query parameter in the request URI to
	// specify the defaults to use when there are no per-document instructions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
	Mget core_mget.NewMget
	// Run multiple searches.
	//
	// The format of the request is similar to the bulk API format and makes use of
	// the newline delimited JSON (NDJSON) format.
	// The structure is as follows:
	//
	// ```
	// header\n
	// body\n
	// header\n
	// body\n
	// ```
	//
	// This structure is specifically optimized to reduce parsing if a specific
	// search ends up redirected to another node.
	//
	// IMPORTANT: The final line of data must end with a newline character `\n`.
	// Each newline character may be preceded by a carriage return `\r`.
	// When sending requests to this endpoint the `Content-Type` header should be
	// set to `application/x-ndjson`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	Msearch core_msearch.NewMsearch
	// Run multiple templated searches.
	//
	// Run multiple templated searches with a single request.
	// If you are providing a text file or text input to `curl`, use the
	// `--data-binary` flag instead of `-d` to preserve newlines.
	// For example:
	//
	// ```
	// $ cat requests
	// { "index": "my-index" }
	// { "id": "my-search-template", "params": { "query_string": "hello world",
	// "from": 0, "size": 10 }}
	// { "index": "my-other-index" }
	// { "id": "my-other-search-template", "params": { "query_type": "match_all" }}
	//
	// $ curl -H "Content-Type: application/x-ndjson" -XGET
	// localhost:9200/_msearch/template --data-binary "@requests"; echo
	// ```
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-search-template.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Get multiple term vectors.
	//
	// Get multiple term vectors with a single request.
	// You can specify existing documents by index and ID or provide artificial
	// documents in the body of the request.
	// You can specify the index in the request body or request URI.
	// The response contains a `docs` array with all the fetched termvectors.
	// Each element has the structure provided by the termvectors API.
	//
	// **Artificial documents**
	//
	// You can also use `mtermvectors` to generate term vectors for artificial
	// documents provided in the body of the request.
	// The mapping used is determined by the specified `_index`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// Open a point in time.
	//
	// A search request by default runs against the most recent visible data of the
	// target indices,
	// which is called point in time. Elasticsearch pit (point in time) is a
	// lightweight view into the
	// state of the data as it existed when initiated. In some cases, it’s preferred
	// to perform multiple
	// search requests using the same point in time. For example, if refreshes
	// happen between
	// `search_after` requests, then the results of those requests might not be
	// consistent as changes happening
	// between searches are only visible to the more recent point in time.
	//
	// A point in time must be opened explicitly before being used in search
	// requests.
	//
	// A subsequent search request with the `pit` parameter must not specify
	// `index`, `routing`, or `preference` values as these parameters are copied
	// from the point in time.
	//
	// Just like regular searches, you can use `from` and `size` to page through
	// point in time search results, up to the first 10,000 hits.
	// If you want to retrieve more hits, use PIT with `search_after`.
	//
	// IMPORTANT: The open point in time request and each subsequent search request
	// can return different identifiers; always use the most recently received ID
	// for the next search request.
	//
	// When a PIT that contains shard failures is used in a search request, the
	// missing are always reported in the search response as a
	// `NoShardAvailableActionException` exception.
	// To get rid of these exceptions, a new PIT needs to be created so that shards
	// missing from the previous PIT can be handled, assuming they become available
	// in the meantime.
	//
	// **Keeping point in time alive**
	//
	// The `keep_alive` parameter, which is passed to a open point in time request
	// and search request, extends the time to live of the corresponding point in
	// time.
	// The value does not need to be long enough to process all data — it just needs
	// to be long enough for the next request.
	//
	// Normally, the background merge process optimizes the index by merging
	// together smaller segments to create new, bigger segments.
	// Once the smaller segments are no longer needed they are deleted.
	// However, open point-in-times prevent the old segments from being deleted
	// since they are still in use.
	//
	// TIP: Keeping older segments alive means that more disk space and file handles
	// are needed.
	// Ensure that you have configured your nodes to have ample free file handles.
	//
	// Additionally, if a segment contains deleted or updated documents then the
	// point in time must keep track of whether each document in the segment was
	// live at the time of the initial search request.
	// Ensure that your nodes have sufficient heap space if you have many open
	// point-in-times on an index that is subject to ongoing deletes or updates.
	// Note that a point-in-time doesn't prevent its associated indices from being
	// deleted.
	// You can check how many point-in-times (that is, search contexts) are open
	// with the nodes stats API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Ping the cluster.
	// Get information about whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster.html
	Ping core_ping.NewPing
	// Create or update a script or search template.
	// Creates or updates a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-stored-script-api.html
	PutScript core_put_script.NewPutScript
	// Evaluate ranked search results.
	//
	// Evaluate the quality of ranked search results over a set of typical search
	// queries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Reindex documents.
	//
	// Copy documents from a source to a destination.
	// You can copy all documents to the destination index or reindex a subset of
	// the documents.
	// The source can be any existing index, alias, or data stream.
	// The destination must differ from the source.
	// For example, you cannot reindex a data stream into itself.
	//
	// IMPORTANT: Reindex requires `_source` to be enabled for all documents in the
	// source.
	// The destination should be configured as wanted before calling the reindex
	// API.
	// Reindex does not copy the settings from the source or its associated
	// template.
	// Mappings, shard counts, and replicas, for example, must be configured ahead
	// of time.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following security privileges:
	//
	// * The `read` index privilege for the source data stream, index, or alias.
	// * The `write` index privilege for the destination data stream, index, or
	// index alias.
	// * To automatically create a data stream or index with a reindex API request,
	// you must have the `auto_configure`, `create_index`, or `manage` index
	// privilege for the destination data stream, index, or alias.
	// * If reindexing from a remote cluster, the `source.remote.user` must have the
	// `monitor` cluster privilege and the `read` index privilege for the source
	// data stream, index, or alias.
	//
	// If reindexing from a remote cluster, you must explicitly allow the remote
	// host in the `reindex.remote.whitelist` setting.
	// Automatic data stream creation requires a matching index template with data
	// stream enabled.
	//
	// The `dest` element can be configured like the index API to control optimistic
	// concurrency control.
	// Omitting `version_type` or setting it to `internal` causes Elasticsearch to
	// blindly dump documents into the destination, overwriting any that happen to
	// have the same ID.
	//
	// Setting `version_type` to `external` causes Elasticsearch to preserve the
	// `version` from the source, create any documents that are missing, and update
	// any documents that have an older version in the destination than they do in
	// the source.
	//
	// Setting `op_type` to `create` causes the reindex API to create only missing
	// documents in the destination.
	// All existing documents will cause a version conflict.
	//
	// IMPORTANT: Because data streams are append-only, any reindex request to a
	// destination data stream must have an `op_type` of `create`.
	// A reindex can only add new documents to a destination data stream.
	// It cannot update existing documents in a destination data stream.
	//
	// By default, version conflicts abort the reindex process.
	// To continue reindexing if there are conflicts, set the `conflicts` request
	// body property to `proceed`.
	// In this case, the response includes a count of the version conflicts that
	// were encountered.
	// Note that the handling of other error types is unaffected by the `conflicts`
	// property.
	// Additionally, if you opt to count version conflicts, the operation could
	// attempt to reindex more documents from the source than `max_docs` until it
	// has successfully indexed `max_docs` documents into the target or it has gone
	// through every document in the source query.
	//
	// NOTE: The reindex API makes no effort to handle ID collisions.
	// The last document written will "win" but the order isn't usually predictable
	// so it is not a good idea to rely on this behavior.
	// Instead, make sure that IDs are unique by using a script.
	//
	// **Running reindex asynchronously**
	//
	// If the request contains `wait_for_completion=false`, Elasticsearch performs
	// some preflight checks, launches the request, and returns a task you can use
	// to cancel or get the status of the task.
	// Elasticsearch creates a record of this task as a document at
	// `_tasks/<task_id>`.
	//
	// **Reindex from multiple sources**
	//
	// If you have many sources to reindex it is generally better to reindex them
	// one at a time rather than using a glob pattern to pick up multiple sources.
	// That way you can resume the process if there are any errors by removing the
	// partially completed source and starting over.
	// It also makes parallelizing the process fairly simple: split the list of
	// sources to reindex and run each list in parallel.
	//
	// For example, you can use a bash script like this:
	//
	// ```
	// for index in i1 i2 i3 i4 i5; do
	//   curl -HContent-Type:application/json -XPOST localhost:9200/_reindex?pretty
	// -d'{
	//     "source": {
	//       "index": "'$index'"
	//     },
	//     "dest": {
	//       "index": "'$index'-reindexed"
	//     }
	//   }'
	// done
	// ```
	//
	// **Throttling**
	//
	// Set `requests_per_second` to any positive decimal number (`1.4`, `6`, `1000`,
	// for example) to throttle the rate at which reindex issues batches of index
	// operations.
	// Requests are throttled by padding each batch with a wait time.
	// To turn off throttling, set `requests_per_second` to `-1`.
	//
	// The throttling is done by waiting between batches so that the scroll that
	// reindex uses internally can be given a timeout that takes into account the
	// padding.
	// The padding time is the difference between the batch size divided by the
	// `requests_per_second` and the time spent writing.
	// By default the batch size is `1000`, so if `requests_per_second` is set to
	// `500`:
	//
	// ```
	// target_time = 1000 / 500 per second = 2 seconds
	// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
	// ```
	//
	// Since the batch is issued as a single bulk request, large batch sizes cause
	// Elasticsearch to create many requests and then wait for a while before
	// starting the next set.
	// This is "bursty" instead of "smooth".
	//
	// **Slicing**
	//
	// Reindex supports sliced scroll to parallelize the reindexing process.
	// This parallelization can improve efficiency and provide a convenient way to
	// break the request down into smaller parts.
	//
	// NOTE: Reindexing from remote clusters does not support manual or automatic
	// slicing.
	//
	// You can slice a reindex request manually by providing a slice ID and total
	// number of slices to each request.
	// You can also let reindex automatically parallelize by using sliced scroll to
	// slice on `_id`.
	// The `slices` parameter specifies the number of slices to use.
	//
	// Adding `slices` to the reindex request just automates the manual process,
	// creating sub-requests which means it has some quirks:
	//
	// * You can see these requests in the tasks API. These sub-requests are "child"
	// tasks of the task for the request with slices.
	// * Fetching the status of the task for the request with `slices` only contains
	// the status of completed slices.
	// * These sub-requests are individually addressable for things like
	// cancellation and rethrottling.
	// * Rethrottling the request with `slices` will rethrottle the unfinished
	// sub-request proportionally.
	// * Canceling the request with `slices` will cancel each sub-request.
	// * Due to the nature of `slices`, each sub-request won't get a perfectly even
	// portion of the documents. All documents will be addressed, but some slices
	// may be larger than others. Expect larger slices to have a more even
	// distribution.
	// * Parameters like `requests_per_second` and `max_docs` on a request with
	// `slices` are distributed proportionally to each sub-request. Combine that
	// with the previous point about distribution being uneven and you should
	// conclude that using `max_docs` with `slices` might not result in exactly
	// `max_docs` documents being reindexed.
	// * Each sub-request gets a slightly different snapshot of the source, though
	// these are all taken at approximately the same time.
	//
	// If slicing automatically, setting `slices` to `auto` will choose a reasonable
	// number for most indices.
	// If slicing manually or otherwise tuning automatic slicing, use the following
	// guidelines.
	//
	// Query performance is most efficient when the number of slices is equal to the
	// number of shards in the index.
	// If that number is large (for example, `500`), choose a lower number as too
	// many slices will hurt performance.
	// Setting slices higher than the number of shards generally does not improve
	// efficiency and adds overhead.
	//
	// Indexing performance scales linearly across available resources with the
	// number of slices.
	//
	// Whether query or indexing performance dominates the runtime depends on the
	// documents being reindexed and cluster resources.
	//
	// **Modify documents during reindexing**
	//
	// Like `_update_by_query`, reindex operations support a script that modifies
	// the document.
	// Unlike `_update_by_query`, the script is allowed to modify the document's
	// metadata.
	//
	// Just as in `_update_by_query`, you can set `ctx.op` to change the operation
	// that is run on the destination.
	// For example, set `ctx.op` to `noop` if your script decides that the document
	// doesn’t have to be indexed in the destination. This "no operation" will be
	// reported in the `noop` counter in the response body.
	// Set `ctx.op` to `delete` if your script decides that the document must be
	// deleted from the destination.
	// The deletion will be reported in the `deleted` counter in the response body.
	// Setting `ctx.op` to anything else will return an error, as will setting any
	// other field in `ctx`.
	//
	// Think of the possibilities! Just be careful; you are able to change:
	//
	// * `_id`
	// * `_index`
	// * `_version`
	// * `_routing`
	//
	// Setting `_version` to `null` or clearing it from the `ctx` map is just like
	// not sending the version in an indexing request.
	// It will cause the document to be overwritten in the destination regardless of
	// the version on the target or the version type you use in the reindex API.
	//
	// **Reindex from remote**
	//
	// Reindex supports reindexing from a remote Elasticsearch cluster.
	// The `host` parameter must contain a scheme, host, port, and optional path.
	// The `username` and `password` parameters are optional and when they are
	// present the reindex operation will connect to the remote Elasticsearch node
	// using basic authentication.
	// Be sure to use HTTPS when using basic authentication or the password will be
	// sent in plain text.
	// There are a range of settings available to configure the behavior of the
	// HTTPS connection.
	//
	// When using Elastic Cloud, it is also possible to authenticate against the
	// remote cluster through the use of a valid API key.
	// Remote hosts must be explicitly allowed with the `reindex.remote.whitelist`
	// setting.
	// It can be set to a comma delimited list of allowed remote host and port
	// combinations.
	// Scheme is ignored; only the host and port are used.
	// For example:
	//
	// ```
	// reindex.remote.whitelist: [otherhost:9200, another:9200, 127.0.10.*:9200,
	// localhost:*"]
	// ```
	//
	// The list of allowed hosts must be configured on any nodes that will
	// coordinate the reindex.
	// This feature should work with remote clusters of any version of
	// Elasticsearch.
	// This should enable you to upgrade from any version of Elasticsearch to the
	// current version by reindexing from a cluster of the old version.
	//
	// WARNING: Elasticsearch does not support forward compatibility across major
	// versions.
	// For example, you cannot reindex from a 7.x cluster into a 6.x cluster.
	//
	// To enable queries sent to older versions of Elasticsearch, the `query`
	// parameter is sent directly to the remote host without validation or
	// modification.
	//
	// NOTE: Reindexing from remote clusters does not support manual or automatic
	// slicing.
	//
	// Reindexing from a remote server uses an on-heap buffer that defaults to a
	// maximum size of 100mb.
	// If the remote index includes very large documents you'll need to use a
	// smaller batch size.
	// It is also possible to set the socket read timeout on the remote connection
	// with the `socket_timeout` field and the connection timeout with the
	// `connect_timeout` field.
	// Both default to 30 seconds.
	//
	// **Configuring SSL parameters**
	//
	// Reindex from remote supports configurable SSL settings.
	// These must be specified in the `elasticsearch.yml` file, with the exception
	// of the secure settings, which you add in the Elasticsearch keystore.
	// It is not possible to configure SSL in the body of the reindex request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Throttle a reindex operation.
	//
	// Change the number of requests per second for a particular reindex operation.
	// For example:
	//
	// ```
	// POST _reindex/r1A2WoRbTwKZ516z6NEs5A:36619/_rethrottle?requests_per_second=-1
	// ```
	//
	// Rethrottling that speeds up the query takes effect immediately.
	// Rethrottling that slows down the query will take effect after completing the
	// current batch.
	// This behavior prevents scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle
	// Render a search template.
	//
	// Render a search template as a search request body.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Run a script.
	//
	// Runs a script and returns a result.
	// Use this API to build and test scripts, such as when defining a script for a
	// runtime field.
	// This API requires very few dependencies and is especially useful if you don't
	// have permissions to write documents on a cluster.
	//
	// The API uses several _contexts_, which control how scripts are run, what
	// variables are available at runtime, and what the return type is.
	//
	// Each context requires a script, but additional parameters depend on the
	// context you're using for that script.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Run a scrolling search.
	//
	// IMPORTANT: The scroll API is no longer recommend for deep pagination. If you
	// need to preserve the index state while paging through more than 10,000 hits,
	// use the `search_after` parameter with a point in time (PIT).
	//
	// The scroll API gets large sets of results from a single scrolling search
	// request.
	// To get the necessary scroll ID, submit a search API request that includes an
	// argument for the `scroll` query parameter.
	// The `scroll` parameter indicates how long Elasticsearch should retain the
	// search context for the request.
	// The search response returns a scroll ID in the `_scroll_id` response body
	// parameter.
	// You can then use the scroll ID with the scroll API to retrieve the next batch
	// of results for the request.
	// If the Elasticsearch security features are enabled, the access to the results
	// of a specific scroll ID is restricted to the user or API key that submitted
	// the search.
	//
	// You can also use the scroll API to specify a new scroll parameter that
	// extends or shortens the retention period for the search context.
	//
	// IMPORTANT: Results from a scrolling search reflect the state of the index at
	// the time of the initial search request. Subsequent indexing or document
	// changes only affect later search and scroll requests.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/scroll-api.html
	Scroll core_scroll.NewScroll
	// Run a search.
	//
	// Get search hits that match the query defined in the request.
	// You can provide search queries using the `q` query string parameter or the
	// request body.
	// If both are specified, only the query parameter is used.
	//
	// If the Elasticsearch security features are enabled, you must have the read
	// index privilege for the target data stream, index, or alias. For
	// cross-cluster search, refer to the documentation about configuring CCS
	// privileges.
	// To search a point in time (PIT) for an alias, you must have the `read` index
	// privilege for the alias's data streams or indices.
	//
	// **Search slicing**
	//
	// When paging through a large number of documents, it can be helpful to split
	// the search into multiple slices to consume them independently with the
	// `slice` and `pit` properties.
	// By default the splitting is done first on the shards, then locally on each
	// shard.
	// The local splitting partitions the shard into contiguous ranges based on
	// Lucene document IDs.
	//
	// For instance if the number of shards is equal to 2 and you request 4 slices,
	// the slices 0 and 2 are assigned to the first shard and the slices 1 and 3 are
	// assigned to the second shard.
	//
	// IMPORTANT: The same point-in-time ID should be used for all slices.
	// If different PIT IDs are used, slices can overlap and miss documents.
	// This situation can occur because the splitting criterion is based on Lucene
	// document IDs, which are not stable across changes to the index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Search a vector tile.
	//
	// Search a vector tile for geospatial values.
	// Before using this API, you should be familiar with the Mapbox vector tile
	// specification.
	// The API returns results as a binary mapbox vector tile.
	//
	// Internally, Elasticsearch translates a vector tile search API request into a
	// search containing:
	//
	// * A `geo_bounding_box` query on the `<field>`. The query uses the
	// `<zoom>/<x>/<y>` tile as a bounding box.
	// * A `geotile_grid` or `geohex_grid` aggregation on the `<field>`. The
	// `grid_agg` parameter determines the aggregation type. The aggregation uses
	// the `<zoom>/<x>/<y>` tile as a bounding box.
	// * Optionally, a `geo_bounds` aggregation on the `<field>`. The search only
	// includes this aggregation if the `exact_bounds` parameter is `true`.
	// * If the optional parameter `with_labels` is `true`, the internal search will
	// include a dynamic runtime field that calls the `getLabelPosition` function of
	// the geometry doc value. This enables the generation of new point features
	// containing suggested geometry labels, so that, for example, multi-polygons
	// will have only one label.
	//
	// For example, Elasticsearch may translate a vector tile search API request
	// with a `grid_agg` argument of `geotile` and an `exact_bounds` argument of
	// `true` into the following search
	//
	// ```
	// GET my-index/_search
	// {
	//   "size": 10000,
	//   "query": {
	//     "geo_bounding_box": {
	//       "my-geo-field": {
	//         "top_left": {
	//           "lat": -40.979898069620134,
	//           "lon": -45
	//         },
	//         "bottom_right": {
	//           "lat": -66.51326044311186,
	//           "lon": 0
	//         }
	//       }
	//     }
	//   },
	//   "aggregations": {
	//     "grid": {
	//       "geotile_grid": {
	//         "field": "my-geo-field",
	//         "precision": 11,
	//         "size": 65536,
	//         "bounds": {
	//           "top_left": {
	//             "lat": -40.979898069620134,
	//             "lon": -45
	//           },
	//           "bottom_right": {
	//             "lat": -66.51326044311186,
	//             "lon": 0
	//           }
	//         }
	//       }
	//     },
	//     "bounds": {
	//       "geo_bounds": {
	//         "field": "my-geo-field",
	//         "wrap_longitude": false
	//       }
	//     }
	//   }
	// }
	// ```
	//
	// The API returns results as a binary Mapbox vector tile.
	// Mapbox vector tiles are encoded as Google Protobufs (PBF). By default, the
	// tile contains three layers:
	//
	// * A `hits` layer containing a feature for each `<field>` value matching the
	// `geo_bounding_box` query.
	// * An `aggs` layer containing a feature for each cell of the `geotile_grid` or
	// `geohex_grid`. The layer only contains features for cells with matching data.
	// * A meta layer containing:
	//   * A feature containing a bounding box. By default, this is the bounding box
	// of the tile.
	//   * Value ranges for any sub-aggregations on the `geotile_grid` or
	// `geohex_grid`.
	//   * Metadata for the search.
	//
	// The API only returns features that can display at its zoom level.
	// For example, if a polygon feature has no area at its zoom level, the API
	// omits it.
	// The API returns errors as UTF-8 encoded JSON.
	//
	// IMPORTANT: You can specify several options for this API as either a query
	// parameter or request body parameter.
	// If you specify both parameters, the query parameter takes precedence.
	//
	// **Grid precision for geotile**
	//
	// For a `grid_agg` of `geotile`, you can use cells in the `aggs` layer as tiles
	// for lower zoom levels.
	// `grid_precision` represents the additional zoom levels available through
	// these cells. The final precision is computed by as follows: `<zoom> +
	// grid_precision`.
	// For example, if `<zoom>` is 7 and `grid_precision` is 8, then the
	// `geotile_grid` aggregation will use a precision of 15.
	// The maximum final precision is 29.
	// The `grid_precision` also determines the number of cells for the grid as
	// follows: `(2^grid_precision) x (2^grid_precision)`.
	// For example, a value of 8 divides the tile into a grid of 256 x 256 cells.
	// The `aggs` layer only contains features for cells with matching data.
	//
	// **Grid precision for geohex**
	//
	// For a `grid_agg` of `geohex`, Elasticsearch uses `<zoom>` and
	// `grid_precision` to calculate a final precision as follows: `<zoom> +
	// grid_precision`.
	//
	// This precision determines the H3 resolution of the hexagonal cells produced
	// by the `geohex` aggregation.
	// The following table maps the H3 resolution for each precision.
	// For example, if `<zoom>` is 3 and `grid_precision` is 3, the precision is 6.
	// At a precision of 6, hexagonal cells have an H3 resolution of 2.
	// If `<zoom>` is 3 and `grid_precision` is 4, the precision is 7.
	// At a precision of 7, hexagonal cells have an H3 resolution of 3.
	//
	// | Precision | Unique tile bins | H3 resolution | Unique hex bins |	Ratio |
	// | --------- | ---------------- | ------------- | ----------------| ----- |
	// | 1  | 4                  | 0  | 122             | 30.5           |
	// | 2  | 16                 | 0  | 122             | 7.625          |
	// | 3  | 64                 | 1  | 842             | 13.15625       |
	// | 4  | 256                | 1  | 842             | 3.2890625      |
	// | 5  | 1024               | 2  | 5882            | 5.744140625    |
	// | 6  | 4096               | 2  | 5882            | 1.436035156    |
	// | 7  | 16384              | 3  | 41162           | 2.512329102    |
	// | 8  | 65536              | 3  | 41162           | 0.6280822754   |
	// | 9  | 262144             | 4  | 288122          | 1.099098206    |
	// | 10 | 1048576            | 4  | 288122          | 0.2747745514   |
	// | 11 | 4194304            | 5  | 2016842         | 0.4808526039   |
	// | 12 | 16777216           | 6  | 14117882        | 0.8414913416   |
	// | 13 | 67108864           | 6  | 14117882        | 0.2103728354   |
	// | 14 | 268435456          | 7  | 98825162        | 0.3681524172   |
	// | 15 | 1073741824         | 8  | 691776122       | 0.644266719    |
	// | 16 | 4294967296         | 8  | 691776122       | 0.1610666797   |
	// | 17 | 17179869184        | 9  | 4842432842      | 0.2818666889   |
	// | 18 | 68719476736        | 10 | 33897029882     | 0.4932667053   |
	// | 19 | 274877906944       | 11 | 237279209162    | 0.8632167343   |
	// | 20 | 1099511627776      | 11 | 237279209162    | 0.2158041836   |
	// | 21 | 4398046511104      | 12 | 1660954464122   | 0.3776573213   |
	// | 22 | 17592186044416     | 13 | 11626681248842  | 0.6609003122   |
	// | 23 | 70368744177664     | 13 | 11626681248842  | 0.165225078    |
	// | 24 | 281474976710656    | 14 | 81386768741882  | 0.2891438866   |
	// | 25 | 1125899906842620   | 15 | 569707381193162 | 0.5060018015   |
	// | 26 | 4503599627370500   | 15 | 569707381193162 | 0.1265004504   |
	// | 27 | 18014398509482000  | 15 | 569707381193162 | 0.03162511259  |
	// | 28 | 72057594037927900  | 15 | 569707381193162 | 0.007906278149 |
	// | 29 | 288230376151712000 | 15 | 569707381193162 | 0.001976569537 |
	//
	// Hexagonal cells don't align perfectly on a vector tile.
	// Some cells may intersect more than one vector tile.
	// To compute the H3 resolution for each precision, Elasticsearch compares the
	// average density of hexagonal bins at each resolution with the average density
	// of tile bins at each zoom level.
	// Elasticsearch uses the H3 resolution that is closest to the corresponding
	// geotile density.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Get the search shards.
	//
	// Get the indices and shards that a search request would be run against.
	// This information can be useful for working out issues or planning
	// optimizations with routing and shard preferences.
	// When filtered aliases are used, the filter is returned as part of the
	// `indices` section.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// `view_index_metadata` or `manage` index privilege for the target data stream,
	// index, or alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-shards.html
	SearchShards core_search_shards.NewSearchShards
	// Run a search with a search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template-api.html
	SearchTemplate core_search_template.NewSearchTemplate
	// Get terms in an index.
	//
	// Discover terms that match a partial string in an index.
	// This API is designed for low-latency look-ups used in auto-complete
	// scenarios.
	//
	// > info
	// > The terms enum API may return terms from deleted documents. Deleted
	// documents are initially only marked as deleted. It is not until their
	// segments are merged that documents are actually deleted. Until that happens,
	// the terms enum API will return terms from these documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
	TermsEnum core_terms_enum.NewTermsEnum
	// Get term vector information.
	//
	// Get information and statistics about terms in the fields of a particular
	// document.
	//
	// You can retrieve term vectors for documents stored in the index or for
	// artificial documents passed in the body of the request.
	// You can specify the fields you are interested in through the `fields`
	// parameter or by adding the fields to the request body.
	// For example:
	//
	// ```
	// GET /my-index-000001/_termvectors/1?fields=message
	// ```
	//
	// Fields can be specified using wildcards, similar to the multi match query.
	//
	// Term vectors are real-time by default, not near real-time.
	// This can be changed by setting `realtime` parameter to `false`.
	//
	// You can request three types of values: _term information_, _term statistics_,
	// and _field statistics_.
	// By default, all term information and field statistics are returned for all
	// fields but term statistics are excluded.
	//
	// **Term information**
	//
	// * term frequency in the field (always returned)
	// * term positions (`positions: true`)
	// * start and end offsets (`offsets: true`)
	// * term payloads (`payloads: true`), as base64 encoded bytes
	//
	// If the requested information wasn't stored in the index, it will be computed
	// on the fly if possible.
	// Additionally, term vectors could be computed for documents not even existing
	// in the index, but instead provided by the user.
	//
	// > warn
	// > Start and end offsets assume UTF-16 encoding is being used. If you want to
	// use these offsets in order to get the original text that produced this token,
	// you should make sure that the string you are taking a sub-string of is also
	// encoded using UTF-16.
	//
	// **Behaviour**
	//
	// The term and field statistics are not accurate.
	// Deleted documents are not taken into account.
	// The information is only retrieved for the shard the requested document
	// resides in.
	// The term and field statistics are therefore only useful as relative measures
	// whereas the absolute numbers have no meaning in this context.
	// By default, when requesting term vectors of artificial documents, a shard to
	// get the statistics from is randomly selected.
	// Use `routing` only to hit a particular shard.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
	Termvectors core_termvectors.NewTermvectors
	// Update a document.
	//
	// Update a document by running a script or passing a partial document.
	//
	// If the Elasticsearch security features are enabled, you must have the `index`
	// or `write` index privilege for the target index or index alias.
	//
	// The script can update, delete, or skip modifying the document.
	// The API also supports passing a partial document, which is merged into the
	// existing document.
	// To fully replace an existing document, use the index API.
	// This operation:
	//
	// * Gets the document (collocated with the shard) from the index.
	// * Runs the specified script.
	// * Indexes the result.
	//
	// The document must still be reindexed, but using this API removes some network
	// roundtrips and reduces chances of version conflicts between the GET and the
	// index operation.
	//
	// The `_source` field must be enabled to use this API.
	// In addition to `_source`, you can access the following variables through the
	// `ctx` map: `_index`, `_type`, `_id`, `_version`, `_routing`, and `_now` (the
	// current timestamp).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
	Update core_update.NewUpdate
	// Update documents.
	// Updates documents that match the specified query.
	// If no query is specified, performs an update on every document in the data
	// stream or index without modifying the source, which is useful for picking up
	// mapping changes.
	//
	// If the Elasticsearch security features are enabled, you must have the
	// following index privileges for the target data stream, index, or alias:
	//
	// * `read`
	// * `index` or `write`
	//
	// You can specify the query criteria in the request URI or the request body
	// using the same syntax as the search API.
	//
	// When you submit an update by query request, Elasticsearch gets a snapshot of
	// the data stream or index when it begins processing the request and updates
	// matching documents using internal versioning.
	// When the versions match, the document is updated and the version number is
	// incremented.
	// If a document changes between the time that the snapshot is taken and the
	// update operation is processed, it results in a version conflict and the
	// operation fails.
	// You can opt to count version conflicts instead of halting and returning by
	// setting `conflicts` to `proceed`.
	// Note that if you opt to count version conflicts, the operation could attempt
	// to update more documents from the source than `max_docs` until it has
	// successfully updated `max_docs` documents or it has gone through every
	// document in the source query.
	//
	// NOTE: Documents with a version equal to 0 cannot be updated using update by
	// query because internal versioning does not support 0 as a valid version
	// number.
	//
	// While processing an update by query request, Elasticsearch performs multiple
	// search requests sequentially to find all of the matching documents.
	// A bulk update request is performed for each batch of matching documents.
	// Any query or update failures cause the update by query request to fail and
	// the failures are shown in the response.
	// Any update requests that completed successfully still stick, they are not
	// rolled back.
	//
	// **Throttling update requests**
	//
	// To control the rate at which update by query issues batches of update
	// operations, you can set `requests_per_second` to any positive decimal number.
	// This pads each batch with a wait time to throttle the rate.
	// Set `requests_per_second` to `-1` to turn off throttling.
	//
	// Throttling uses a wait time between batches so that the internal scroll
	// requests can be given a timeout that takes the request padding into account.
	// The padding time is the difference between the batch size divided by the
	// `requests_per_second` and the time spent writing.
	// By default the batch size is 1000, so if `requests_per_second` is set to
	// `500`:
	//
	// ```
	// target_time = 1000 / 500 per second = 2 seconds
	// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
	// ```
	//
	// Since the batch is issued as a single _bulk request, large batch sizes cause
	// Elasticsearch to create many requests and wait before starting the next set.
	// This is "bursty" instead of "smooth".
	//
	// **Slicing**
	//
	// Update by query supports sliced scroll to parallelize the update process.
	// This can improve efficiency and provide a convenient way to break the request
	// down into smaller parts.
	//
	// Setting `slices` to `auto` chooses a reasonable number for most data streams
	// and indices.
	// This setting will use one slice per shard, up to a certain limit.
	// If there are multiple source data streams or indices, it will choose the
	// number of slices based on the index or backing index with the smallest number
	// of shards.
	//
	// Adding `slices` to `_update_by_query` just automates the manual process of
	// creating sub-requests, which means it has some quirks:
	//
	// * You can see these requests in the tasks APIs. These sub-requests are
	// "child" tasks of the task for the request with slices.
	// * Fetching the status of the task for the request with `slices` only contains
	// the status of completed slices.
	// * These sub-requests are individually addressable for things like
	// cancellation and rethrottling.
	// * Rethrottling the request with `slices` will rethrottle the unfinished
	// sub-request proportionally.
	// * Canceling the request with slices will cancel each sub-request.
	// * Due to the nature of slices each sub-request won't get a perfectly even
	// portion of the documents. All documents will be addressed, but some slices
	// may be larger than others. Expect larger slices to have a more even
	// distribution.
	// * Parameters like `requests_per_second` and `max_docs` on a request with
	// slices are distributed proportionally to each sub-request. Combine that with
	// the point above about distribution being uneven and you should conclude that
	// using `max_docs` with `slices` might not result in exactly `max_docs`
	// documents being updated.
	// * Each sub-request gets a slightly different snapshot of the source data
	// stream or index though these are all taken at approximately the same time.
	//
	// If you're slicing manually or otherwise tuning automatic slicing, keep in
	// mind that:
	//
	// * Query performance is most efficient when the number of slices is equal to
	// the number of shards in the index or backing index. If that number is large
	// (for example, 500), choose a lower number as too many slices hurts
	// performance. Setting slices higher than the number of shards generally does
	// not improve efficiency and adds overhead.
	// * Update performance scales linearly across available resources with the
	// number of slices.
	//
	// Whether query or update performance dominates the runtime depends on the
	// documents being reindexed and cluster resources.
	//
	// **Update the document source**
	//
	// Update by query supports scripts to update the document source.
	// As with the update API, you can set `ctx.op` to change the operation that is
	// performed.
	//
	// Set `ctx.op = "noop"` if your script decides that it doesn't have to make any
	// changes.
	// The update by query operation skips updating the document and increments the
	// `noop` counter.
	//
	// Set `ctx.op = "delete"` if your script decides that the document should be
	// deleted.
	// The update by query operation deletes the document and increments the
	// `deleted` counter.
	//
	// Update by query supports only `index`, `noop`, and `delete`.
	// Setting `ctx.op` to anything else is an error.
	// Setting any other field in `ctx` is an error.
	// This API enables you to only modify the source of matching documents; you
	// cannot move them.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQuery core_update_by_query.NewUpdateByQuery
	// Throttle an update by query operation.
	//
	// Change the number of requests per second for a particular update by query
	// operation.
	// Rethrottling that speeds up the query takes effect immediately but
	// rethrotting that slows down the query takes effect after completing the
	// current batch to prevent scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html#docs-update-by-query-rethrottle
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
			SyncJobCheckIn:            connector_sync_job_check_in.NewSyncJobCheckInFunc(tp),
			SyncJobClaim:              connector_sync_job_claim.NewSyncJobClaimFunc(tp),
			SyncJobDelete:             connector_sync_job_delete.NewSyncJobDeleteFunc(tp),
			SyncJobError:              connector_sync_job_error.NewSyncJobErrorFunc(tp),
			SyncJobGet:                connector_sync_job_get.NewSyncJobGetFunc(tp),
			SyncJobList:               connector_sync_job_list.NewSyncJobListFunc(tp),
			SyncJobPost:               connector_sync_job_post.NewSyncJobPostFunc(tp),
			SyncJobUpdateStats:        connector_sync_job_update_stats.NewSyncJobUpdateStatsFunc(tp),
			UpdateActiveFiltering:     connector_update_active_filtering.NewUpdateActiveFilteringFunc(tp),
			UpdateApiKeyId:            connector_update_api_key_id.NewUpdateApiKeyIdFunc(tp),
			UpdateConfiguration:       connector_update_configuration.NewUpdateConfigurationFunc(tp),
			UpdateError:               connector_update_error.NewUpdateErrorFunc(tp),
			UpdateFeatures:            connector_update_features.NewUpdateFeaturesFunc(tp),
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
			AsyncQuery:       esql_async_query.NewAsyncQueryFunc(tp),
			AsyncQueryDelete: esql_async_query_delete.NewAsyncQueryDeleteFunc(tp),
			AsyncQueryGet:    esql_async_query_get.NewAsyncQueryGetFunc(tp),
			AsyncQueryStop:   esql_async_query_stop.NewAsyncQueryStopFunc(tp),
			Query:            esql_query.NewQueryFunc(tp),
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
			AddBlock:                indices_add_block.NewAddBlockFunc(tp),
			Analyze:                 indices_analyze.NewAnalyzeFunc(tp),
			CancelMigrateReindex:    indices_cancel_migrate_reindex.NewCancelMigrateReindexFunc(tp),
			ClearCache:              indices_clear_cache.NewClearCacheFunc(tp),
			Clone:                   indices_clone.NewCloneFunc(tp),
			Close:                   indices_close.NewCloseFunc(tp),
			Create:                  indices_create.NewCreateFunc(tp),
			CreateDataStream:        indices_create_data_stream.NewCreateDataStreamFunc(tp),
			CreateFrom:              indices_create_from.NewCreateFromFunc(tp),
			DataStreamsStats:        indices_data_streams_stats.NewDataStreamsStatsFunc(tp),
			Delete:                  indices_delete.NewDeleteFunc(tp),
			DeleteAlias:             indices_delete_alias.NewDeleteAliasFunc(tp),
			DeleteDataLifecycle:     indices_delete_data_lifecycle.NewDeleteDataLifecycleFunc(tp),
			DeleteDataStream:        indices_delete_data_stream.NewDeleteDataStreamFunc(tp),
			DeleteIndexTemplate:     indices_delete_index_template.NewDeleteIndexTemplateFunc(tp),
			DeleteTemplate:          indices_delete_template.NewDeleteTemplateFunc(tp),
			DiskUsage:               indices_disk_usage.NewDiskUsageFunc(tp),
			Downsample:              indices_downsample.NewDownsampleFunc(tp),
			Exists:                  indices_exists.NewExistsFunc(tp),
			ExistsAlias:             indices_exists_alias.NewExistsAliasFunc(tp),
			ExistsIndexTemplate:     indices_exists_index_template.NewExistsIndexTemplateFunc(tp),
			ExistsTemplate:          indices_exists_template.NewExistsTemplateFunc(tp),
			ExplainDataLifecycle:    indices_explain_data_lifecycle.NewExplainDataLifecycleFunc(tp),
			FieldUsageStats:         indices_field_usage_stats.NewFieldUsageStatsFunc(tp),
			Flush:                   indices_flush.NewFlushFunc(tp),
			Forcemerge:              indices_forcemerge.NewForcemergeFunc(tp),
			Get:                     indices_get.NewGetFunc(tp),
			GetAlias:                indices_get_alias.NewGetAliasFunc(tp),
			GetDataLifecycle:        indices_get_data_lifecycle.NewGetDataLifecycleFunc(tp),
			GetDataLifecycleStats:   indices_get_data_lifecycle_stats.NewGetDataLifecycleStatsFunc(tp),
			GetDataStream:           indices_get_data_stream.NewGetDataStreamFunc(tp),
			GetFieldMapping:         indices_get_field_mapping.NewGetFieldMappingFunc(tp),
			GetIndexTemplate:        indices_get_index_template.NewGetIndexTemplateFunc(tp),
			GetMapping:              indices_get_mapping.NewGetMappingFunc(tp),
			GetMigrateReindexStatus: indices_get_migrate_reindex_status.NewGetMigrateReindexStatusFunc(tp),
			GetSettings:             indices_get_settings.NewGetSettingsFunc(tp),
			GetTemplate:             indices_get_template.NewGetTemplateFunc(tp),
			MigrateReindex:          indices_migrate_reindex.NewMigrateReindexFunc(tp),
			MigrateToDataStream:     indices_migrate_to_data_stream.NewMigrateToDataStreamFunc(tp),
			ModifyDataStream:        indices_modify_data_stream.NewModifyDataStreamFunc(tp),
			Open:                    indices_open.NewOpenFunc(tp),
			PromoteDataStream:       indices_promote_data_stream.NewPromoteDataStreamFunc(tp),
			PutAlias:                indices_put_alias.NewPutAliasFunc(tp),
			PutDataLifecycle:        indices_put_data_lifecycle.NewPutDataLifecycleFunc(tp),
			PutIndexTemplate:        indices_put_index_template.NewPutIndexTemplateFunc(tp),
			PutMapping:              indices_put_mapping.NewPutMappingFunc(tp),
			PutSettings:             indices_put_settings.NewPutSettingsFunc(tp),
			PutTemplate:             indices_put_template.NewPutTemplateFunc(tp),
			Recovery:                indices_recovery.NewRecoveryFunc(tp),
			Refresh:                 indices_refresh.NewRefreshFunc(tp),
			ReloadSearchAnalyzers:   indices_reload_search_analyzers.NewReloadSearchAnalyzersFunc(tp),
			ResolveCluster:          indices_resolve_cluster.NewResolveClusterFunc(tp),
			ResolveIndex:            indices_resolve_index.NewResolveIndexFunc(tp),
			Rollover:                indices_rollover.NewRolloverFunc(tp),
			Segments:                indices_segments.NewSegmentsFunc(tp),
			ShardStores:             indices_shard_stores.NewShardStoresFunc(tp),
			Shrink:                  indices_shrink.NewShrinkFunc(tp),
			SimulateIndexTemplate:   indices_simulate_index_template.NewSimulateIndexTemplateFunc(tp),
			SimulateTemplate:        indices_simulate_template.NewSimulateTemplateFunc(tp),
			Split:                   indices_split.NewSplitFunc(tp),
			Stats:                   indices_stats.NewStatsFunc(tp),
			Unfreeze:                indices_unfreeze.NewUnfreezeFunc(tp),
			UpdateAliases:           indices_update_aliases.NewUpdateAliasesFunc(tp),
			ValidateQuery:           indices_validate_query.NewValidateQueryFunc(tp),
		},

		// Inference
		Inference: Inference{
			ChatCompletionUnified: inference_chat_completion_unified.NewChatCompletionUnifiedFunc(tp),
			Completion:            inference_completion.NewCompletionFunc(tp),
			Delete:                inference_delete.NewDeleteFunc(tp),
			Get:                   inference_get.NewGetFunc(tp),
			Inference:             inference_inference.NewInferenceFunc(tp),
			Put:                   inference_put.NewPutFunc(tp),
			PutAlibabacloud:       inference_put_alibabacloud.NewPutAlibabacloudFunc(tp),
			PutAmazonbedrock:      inference_put_amazonbedrock.NewPutAmazonbedrockFunc(tp),
			PutAnthropic:          inference_put_anthropic.NewPutAnthropicFunc(tp),
			PutAzureaistudio:      inference_put_azureaistudio.NewPutAzureaistudioFunc(tp),
			PutAzureopenai:        inference_put_azureopenai.NewPutAzureopenaiFunc(tp),
			PutCohere:             inference_put_cohere.NewPutCohereFunc(tp),
			PutElasticsearch:      inference_put_elasticsearch.NewPutElasticsearchFunc(tp),
			PutElser:              inference_put_elser.NewPutElserFunc(tp),
			PutGoogleaistudio:     inference_put_googleaistudio.NewPutGoogleaistudioFunc(tp),
			PutGooglevertexai:     inference_put_googlevertexai.NewPutGooglevertexaiFunc(tp),
			PutHuggingFace:        inference_put_hugging_face.NewPutHuggingFaceFunc(tp),
			PutJinaai:             inference_put_jinaai.NewPutJinaaiFunc(tp),
			PutMistral:            inference_put_mistral.NewPutMistralFunc(tp),
			PutOpenai:             inference_put_openai.NewPutOpenaiFunc(tp),
			PutVoyageai:           inference_put_voyageai.NewPutVoyageaiFunc(tp),
			PutWatsonx:            inference_put_watsonx.NewPutWatsonxFunc(tp),
			Rerank:                inference_rerank.NewRerankFunc(tp),
			SparseEmbedding:       inference_sparse_embedding.NewSparseEmbeddingFunc(tp),
			StreamCompletion:      inference_stream_completion.NewStreamCompletionFunc(tp),
			TextEmbedding:         inference_text_embedding.NewTextEmbeddingFunc(tp),
			Update:                inference_update.NewUpdateFunc(tp),
		},

		// Ingest
		Ingest: Ingest{
			DeleteGeoipDatabase:      ingest_delete_geoip_database.NewDeleteGeoipDatabaseFunc(tp),
			DeleteIpLocationDatabase: ingest_delete_ip_location_database.NewDeleteIpLocationDatabaseFunc(tp),
			DeletePipeline:           ingest_delete_pipeline.NewDeletePipelineFunc(tp),
			GeoIpStats:               ingest_geo_ip_stats.NewGeoIpStatsFunc(tp),
			GetGeoipDatabase:         ingest_get_geoip_database.NewGetGeoipDatabaseFunc(tp),
			GetIpLocationDatabase:    ingest_get_ip_location_database.NewGetIpLocationDatabaseFunc(tp),
			GetPipeline:              ingest_get_pipeline.NewGetPipelineFunc(tp),
			ProcessorGrok:            ingest_processor_grok.NewProcessorGrokFunc(tp),
			PutGeoipDatabase:         ingest_put_geoip_database.NewPutGeoipDatabaseFunc(tp),
			PutIpLocationDatabase:    ingest_put_ip_location_database.NewPutIpLocationDatabaseFunc(tp),
			PutPipeline:              ingest_put_pipeline.NewPutPipelineFunc(tp),
			Simulate:                 ingest_simulate.NewSimulateFunc(tp),
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
			Test:          query_rules_test.NewTestFunc(tp),
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
			Delete:                       search_application_delete.NewDeleteFunc(tp),
			DeleteBehavioralAnalytics:    search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalyticsFunc(tp),
			Get:                          search_application_get.NewGetFunc(tp),
			GetBehavioralAnalytics:       search_application_get_behavioral_analytics.NewGetBehavioralAnalyticsFunc(tp),
			List:                         search_application_list.NewListFunc(tp),
			PostBehavioralAnalyticsEvent: search_application_post_behavioral_analytics_event.NewPostBehavioralAnalyticsEventFunc(tp),
			Put:                          search_application_put.NewPutFunc(tp),
			PutBehavioralAnalytics:       search_application_put_behavioral_analytics.NewPutBehavioralAnalyticsFunc(tp),
			RenderQuery:                  search_application_render_query.NewRenderQueryFunc(tp),
			Search:                       search_application_search.NewSearchFunc(tp),
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
			DelegatePki:                 security_delegate_pki.NewDelegatePkiFunc(tp),
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
			UpdateCrossClusterApiKey:    security_update_cross_cluster_api_key.NewUpdateCrossClusterApiKeyFunc(tp),
			UpdateSettings:              security_update_settings.NewUpdateSettingsFunc(tp),
			UpdateUserProfileData:       security_update_user_profile_data.NewUpdateUserProfileDataFunc(tp),
		},

		// Shutdown
		Shutdown: Shutdown{
			DeleteNode: shutdown_delete_node.NewDeleteNodeFunc(tp),
			GetNode:    shutdown_get_node.NewGetNodeFunc(tp),
			PutNode:    shutdown_put_node.NewPutNodeFunc(tp),
		},

		// Simulate
		Simulate: Simulate{
			Ingest: simulate_ingest.NewIngestFunc(tp),
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
			CleanupRepository:         snapshot_cleanup_repository.NewCleanupRepositoryFunc(tp),
			Clone:                     snapshot_clone.NewCloneFunc(tp),
			Create:                    snapshot_create.NewCreateFunc(tp),
			CreateRepository:          snapshot_create_repository.NewCreateRepositoryFunc(tp),
			Delete:                    snapshot_delete.NewDeleteFunc(tp),
			DeleteRepository:          snapshot_delete_repository.NewDeleteRepositoryFunc(tp),
			Get:                       snapshot_get.NewGetFunc(tp),
			GetRepository:             snapshot_get_repository.NewGetRepositoryFunc(tp),
			RepositoryAnalyze:         snapshot_repository_analyze.NewRepositoryAnalyzeFunc(tp),
			RepositoryVerifyIntegrity: snapshot_repository_verify_integrity.NewRepositoryVerifyIntegrityFunc(tp),
			Restore:                   snapshot_restore.NewRestoreFunc(tp),
			Status:                    snapshot_status.NewStatusFunc(tp),
			VerifyRepository:          snapshot_verify_repository.NewVerifyRepositoryFunc(tp),
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

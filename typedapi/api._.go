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
	inference_put_amazonsagemaker "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putamazonsagemaker"
	inference_put_anthropic "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putanthropic"
	inference_put_azureaistudio "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putazureaistudio"
	inference_put_azureopenai "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putazureopenai"
	inference_put_cohere "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putcohere"
	inference_put_custom "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putcustom"
	inference_put_deepseek "github.com/elastic/go-elasticsearch/v8/typedapi/inference/putdeepseek"
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
	streams_logs_disable "github.com/elastic/go-elasticsearch/v8/typedapi/streams/logsdisable"
	streams_logs_enable "github.com/elastic/go-elasticsearch/v8/typedapi/streams/logsenable"
	streams_status "github.com/elastic/go-elasticsearch/v8/typedapi/streams/status"
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
	Delete async_search_delete.NewDelete
	Get    async_search_get.NewGet
	Status async_search_status.NewStatus
	Submit async_search_submit.NewSubmit
}

type Autoscaling struct {
	DeleteAutoscalingPolicy autoscaling_delete_autoscaling_policy.NewDeleteAutoscalingPolicy
	GetAutoscalingCapacity  autoscaling_get_autoscaling_capacity.NewGetAutoscalingCapacity
	GetAutoscalingPolicy    autoscaling_get_autoscaling_policy.NewGetAutoscalingPolicy
	PutAutoscalingPolicy    autoscaling_put_autoscaling_policy.NewPutAutoscalingPolicy
}

type Capabilities struct {
	Capabilities capabilities.NewCapabilities
}

type Cat struct {
	Aliases              cat_aliases.NewAliases
	Allocation           cat_allocation.NewAllocation
	ComponentTemplates   cat_component_templates.NewComponentTemplates
	Count                cat_count.NewCount
	Fielddata            cat_fielddata.NewFielddata
	Health               cat_health.NewHealth
	Help                 cat_help.NewHelp
	Indices              cat_indices.NewIndices
	Master               cat_master.NewMaster
	MlDataFrameAnalytics cat_ml_data_frame_analytics.NewMlDataFrameAnalytics
	MlDatafeeds          cat_ml_datafeeds.NewMlDatafeeds
	MlJobs               cat_ml_jobs.NewMlJobs
	MlTrainedModels      cat_ml_trained_models.NewMlTrainedModels
	Nodeattrs            cat_nodeattrs.NewNodeattrs
	Nodes                cat_nodes.NewNodes
	PendingTasks         cat_pending_tasks.NewPendingTasks
	Plugins              cat_plugins.NewPlugins
	Recovery             cat_recovery.NewRecovery
	Repositories         cat_repositories.NewRepositories
	Segments             cat_segments.NewSegments
	Shards               cat_shards.NewShards
	Snapshots            cat_snapshots.NewSnapshots
	Tasks                cat_tasks.NewTasks
	Templates            cat_templates.NewTemplates
	ThreadPool           cat_thread_pool.NewThreadPool
	Transforms           cat_transforms.NewTransforms
}

type Ccr struct {
	DeleteAutoFollowPattern ccr_delete_auto_follow_pattern.NewDeleteAutoFollowPattern
	Follow                  ccr_follow.NewFollow
	FollowInfo              ccr_follow_info.NewFollowInfo
	FollowStats             ccr_follow_stats.NewFollowStats
	ForgetFollower          ccr_forget_follower.NewForgetFollower
	GetAutoFollowPattern    ccr_get_auto_follow_pattern.NewGetAutoFollowPattern
	PauseAutoFollowPattern  ccr_pause_auto_follow_pattern.NewPauseAutoFollowPattern
	PauseFollow             ccr_pause_follow.NewPauseFollow
	PutAutoFollowPattern    ccr_put_auto_follow_pattern.NewPutAutoFollowPattern
	ResumeAutoFollowPattern ccr_resume_auto_follow_pattern.NewResumeAutoFollowPattern
	ResumeFollow            ccr_resume_follow.NewResumeFollow
	Stats                   ccr_stats.NewStats
	Unfollow                ccr_unfollow.NewUnfollow
}

type Cluster struct {
	AllocationExplain            cluster_allocation_explain.NewAllocationExplain
	DeleteComponentTemplate      cluster_delete_component_template.NewDeleteComponentTemplate
	DeleteVotingConfigExclusions cluster_delete_voting_config_exclusions.NewDeleteVotingConfigExclusions
	ExistsComponentTemplate      cluster_exists_component_template.NewExistsComponentTemplate
	GetComponentTemplate         cluster_get_component_template.NewGetComponentTemplate
	GetSettings                  cluster_get_settings.NewGetSettings
	Health                       cluster_health.NewHealth
	Info                         cluster_info.NewInfo
	PendingTasks                 cluster_pending_tasks.NewPendingTasks
	PostVotingConfigExclusions   cluster_post_voting_config_exclusions.NewPostVotingConfigExclusions
	PutComponentTemplate         cluster_put_component_template.NewPutComponentTemplate
	PutSettings                  cluster_put_settings.NewPutSettings
	RemoteInfo                   cluster_remote_info.NewRemoteInfo
	Reroute                      cluster_reroute.NewReroute
	State                        cluster_state.NewState
	Stats                        cluster_stats.NewStats
}

type Connector struct {
	CheckIn                   connector_check_in.NewCheckIn
	Delete                    connector_delete.NewDelete
	Get                       connector_get.NewGet
	LastSync                  connector_last_sync.NewLastSync
	List                      connector_list.NewList
	Post                      connector_post.NewPost
	Put                       connector_put.NewPut
	SecretPost                connector_secret_post.NewSecretPost
	SyncJobCancel             connector_sync_job_cancel.NewSyncJobCancel
	SyncJobCheckIn            connector_sync_job_check_in.NewSyncJobCheckIn
	SyncJobClaim              connector_sync_job_claim.NewSyncJobClaim
	SyncJobDelete             connector_sync_job_delete.NewSyncJobDelete
	SyncJobError              connector_sync_job_error.NewSyncJobError
	SyncJobGet                connector_sync_job_get.NewSyncJobGet
	SyncJobList               connector_sync_job_list.NewSyncJobList
	SyncJobPost               connector_sync_job_post.NewSyncJobPost
	SyncJobUpdateStats        connector_sync_job_update_stats.NewSyncJobUpdateStats
	UpdateActiveFiltering     connector_update_active_filtering.NewUpdateActiveFiltering
	UpdateApiKeyId            connector_update_api_key_id.NewUpdateApiKeyId
	UpdateConfiguration       connector_update_configuration.NewUpdateConfiguration
	UpdateError               connector_update_error.NewUpdateError
	UpdateFeatures            connector_update_features.NewUpdateFeatures
	UpdateFiltering           connector_update_filtering.NewUpdateFiltering
	UpdateFilteringValidation connector_update_filtering_validation.NewUpdateFilteringValidation
	UpdateIndexName           connector_update_index_name.NewUpdateIndexName
	UpdateName                connector_update_name.NewUpdateName
	UpdateNative              connector_update_native.NewUpdateNative
	UpdatePipeline            connector_update_pipeline.NewUpdatePipeline
	UpdateScheduling          connector_update_scheduling.NewUpdateScheduling
	UpdateServiceType         connector_update_service_type.NewUpdateServiceType
	UpdateStatus              connector_update_status.NewUpdateStatus
}

type Core struct {
	Bulk                    core_bulk.NewBulk
	ClearScroll             core_clear_scroll.NewClearScroll
	ClosePointInTime        core_close_point_in_time.NewClosePointInTime
	Count                   core_count.NewCount
	Create                  core_create.NewCreate
	Delete                  core_delete.NewDelete
	DeleteByQuery           core_delete_by_query.NewDeleteByQuery
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	DeleteScript            core_delete_script.NewDeleteScript
	Exists                  core_exists.NewExists
	ExistsSource            core_exists_source.NewExistsSource
	Explain                 core_explain.NewExplain
	FieldCaps               core_field_caps.NewFieldCaps
	Get                     core_get.NewGet
	GetScript               core_get_script.NewGetScript
	GetScriptContext        core_get_script_context.NewGetScriptContext
	GetScriptLanguages      core_get_script_languages.NewGetScriptLanguages
	GetSource               core_get_source.NewGetSource
	HealthReport            core_health_report.NewHealthReport
	Index                   core_index.NewIndex
	Info                    core_info.NewInfo
	KnnSearch               core_knn_search.NewKnnSearch
	Mget                    core_mget.NewMget
	Msearch                 core_msearch.NewMsearch
	MsearchTemplate         core_msearch_template.NewMsearchTemplate
	Mtermvectors            core_mtermvectors.NewMtermvectors
	OpenPointInTime         core_open_point_in_time.NewOpenPointInTime
	Ping                    core_ping.NewPing
	PutScript               core_put_script.NewPutScript
	RankEval                core_rank_eval.NewRankEval
	Reindex                 core_reindex.NewReindex
	ReindexRethrottle       core_reindex_rethrottle.NewReindexRethrottle
	RenderSearchTemplate    core_render_search_template.NewRenderSearchTemplate
	ScriptsPainlessExecute  core_scripts_painless_execute.NewScriptsPainlessExecute
	Scroll                  core_scroll.NewScroll
	Search                  core_search.NewSearch
	SearchMvt               core_search_mvt.NewSearchMvt
	SearchShards            core_search_shards.NewSearchShards
	SearchTemplate          core_search_template.NewSearchTemplate
	TermsEnum               core_terms_enum.NewTermsEnum
	Termvectors             core_termvectors.NewTermvectors
	Update                  core_update.NewUpdate
	UpdateByQuery           core_update_by_query.NewUpdateByQuery
	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

type DanglingIndices struct {
	DeleteDanglingIndex dangling_indices_delete_dangling_index.NewDeleteDanglingIndex
	ImportDanglingIndex dangling_indices_import_dangling_index.NewImportDanglingIndex
	ListDanglingIndices dangling_indices_list_dangling_indices.NewListDanglingIndices
}

type Enrich struct {
	DeletePolicy  enrich_delete_policy.NewDeletePolicy
	ExecutePolicy enrich_execute_policy.NewExecutePolicy
	GetPolicy     enrich_get_policy.NewGetPolicy
	PutPolicy     enrich_put_policy.NewPutPolicy
	Stats         enrich_stats.NewStats
}

type Eql struct {
	Delete    eql_delete.NewDelete
	Get       eql_get.NewGet
	GetStatus eql_get_status.NewGetStatus
	Search    eql_search.NewSearch
}

type Esql struct {
	AsyncQuery       esql_async_query.NewAsyncQuery
	AsyncQueryDelete esql_async_query_delete.NewAsyncQueryDelete
	AsyncQueryGet    esql_async_query_get.NewAsyncQueryGet
	AsyncQueryStop   esql_async_query_stop.NewAsyncQueryStop
	Query            esql_query.NewQuery
}

type Features struct {
	GetFeatures   features_get_features.NewGetFeatures
	ResetFeatures features_reset_features.NewResetFeatures
}

type Fleet struct {
	GlobalCheckpoints fleet_global_checkpoints.NewGlobalCheckpoints
	Msearch           fleet_msearch.NewMsearch
	PostSecret        fleet_post_secret.NewPostSecret
	Search            fleet_search.NewSearch
}

type Graph struct {
	Explore graph_explore.NewExplore
}

type Ilm struct {
	DeleteLifecycle    ilm_delete_lifecycle.NewDeleteLifecycle
	ExplainLifecycle   ilm_explain_lifecycle.NewExplainLifecycle
	GetLifecycle       ilm_get_lifecycle.NewGetLifecycle
	GetStatus          ilm_get_status.NewGetStatus
	MigrateToDataTiers ilm_migrate_to_data_tiers.NewMigrateToDataTiers
	MoveToStep         ilm_move_to_step.NewMoveToStep
	PutLifecycle       ilm_put_lifecycle.NewPutLifecycle
	RemovePolicy       ilm_remove_policy.NewRemovePolicy
	Retry              ilm_retry.NewRetry
	Start              ilm_start.NewStart
	Stop               ilm_stop.NewStop
}

type Indices struct {
	AddBlock                indices_add_block.NewAddBlock
	Analyze                 indices_analyze.NewAnalyze
	CancelMigrateReindex    indices_cancel_migrate_reindex.NewCancelMigrateReindex
	ClearCache              indices_clear_cache.NewClearCache
	Clone                   indices_clone.NewClone
	Close                   indices_close.NewClose
	Create                  indices_create.NewCreate
	CreateDataStream        indices_create_data_stream.NewCreateDataStream
	CreateFrom              indices_create_from.NewCreateFrom
	DataStreamsStats        indices_data_streams_stats.NewDataStreamsStats
	Delete                  indices_delete.NewDelete
	DeleteAlias             indices_delete_alias.NewDeleteAlias
	DeleteDataLifecycle     indices_delete_data_lifecycle.NewDeleteDataLifecycle
	DeleteDataStream        indices_delete_data_stream.NewDeleteDataStream
	DeleteIndexTemplate     indices_delete_index_template.NewDeleteIndexTemplate
	DeleteTemplate          indices_delete_template.NewDeleteTemplate
	DiskUsage               indices_disk_usage.NewDiskUsage
	Downsample              indices_downsample.NewDownsample
	Exists                  indices_exists.NewExists
	ExistsAlias             indices_exists_alias.NewExistsAlias
	ExistsIndexTemplate     indices_exists_index_template.NewExistsIndexTemplate
	ExistsTemplate          indices_exists_template.NewExistsTemplate
	ExplainDataLifecycle    indices_explain_data_lifecycle.NewExplainDataLifecycle
	FieldUsageStats         indices_field_usage_stats.NewFieldUsageStats
	Flush                   indices_flush.NewFlush
	Forcemerge              indices_forcemerge.NewForcemerge
	Get                     indices_get.NewGet
	GetAlias                indices_get_alias.NewGetAlias
	GetDataLifecycle        indices_get_data_lifecycle.NewGetDataLifecycle
	GetDataLifecycleStats   indices_get_data_lifecycle_stats.NewGetDataLifecycleStats
	GetDataStream           indices_get_data_stream.NewGetDataStream
	GetFieldMapping         indices_get_field_mapping.NewGetFieldMapping
	GetIndexTemplate        indices_get_index_template.NewGetIndexTemplate
	GetMapping              indices_get_mapping.NewGetMapping
	GetMigrateReindexStatus indices_get_migrate_reindex_status.NewGetMigrateReindexStatus
	GetSettings             indices_get_settings.NewGetSettings
	GetTemplate             indices_get_template.NewGetTemplate
	MigrateReindex          indices_migrate_reindex.NewMigrateReindex
	MigrateToDataStream     indices_migrate_to_data_stream.NewMigrateToDataStream
	ModifyDataStream        indices_modify_data_stream.NewModifyDataStream
	Open                    indices_open.NewOpen
	PromoteDataStream       indices_promote_data_stream.NewPromoteDataStream
	PutAlias                indices_put_alias.NewPutAlias
	PutDataLifecycle        indices_put_data_lifecycle.NewPutDataLifecycle
	PutIndexTemplate        indices_put_index_template.NewPutIndexTemplate
	PutMapping              indices_put_mapping.NewPutMapping
	PutSettings             indices_put_settings.NewPutSettings
	PutTemplate             indices_put_template.NewPutTemplate
	Recovery                indices_recovery.NewRecovery
	Refresh                 indices_refresh.NewRefresh
	ReloadSearchAnalyzers   indices_reload_search_analyzers.NewReloadSearchAnalyzers
	ResolveCluster          indices_resolve_cluster.NewResolveCluster
	ResolveIndex            indices_resolve_index.NewResolveIndex
	Rollover                indices_rollover.NewRollover
	Segments                indices_segments.NewSegments
	ShardStores             indices_shard_stores.NewShardStores
	Shrink                  indices_shrink.NewShrink
	SimulateIndexTemplate   indices_simulate_index_template.NewSimulateIndexTemplate
	SimulateTemplate        indices_simulate_template.NewSimulateTemplate
	Split                   indices_split.NewSplit
	Stats                   indices_stats.NewStats
	Unfreeze                indices_unfreeze.NewUnfreeze
	UpdateAliases           indices_update_aliases.NewUpdateAliases
	ValidateQuery           indices_validate_query.NewValidateQuery
}

type Inference struct {
	ChatCompletionUnified inference_chat_completion_unified.NewChatCompletionUnified
	Completion            inference_completion.NewCompletion
	Delete                inference_delete.NewDelete
	Get                   inference_get.NewGet
	Inference             inference_inference.NewInference
	Put                   inference_put.NewPut
	PutAlibabacloud       inference_put_alibabacloud.NewPutAlibabacloud
	PutAmazonbedrock      inference_put_amazonbedrock.NewPutAmazonbedrock
	PutAmazonsagemaker    inference_put_amazonsagemaker.NewPutAmazonsagemaker
	PutAnthropic          inference_put_anthropic.NewPutAnthropic
	PutAzureaistudio      inference_put_azureaistudio.NewPutAzureaistudio
	PutAzureopenai        inference_put_azureopenai.NewPutAzureopenai
	PutCohere             inference_put_cohere.NewPutCohere
	PutCustom             inference_put_custom.NewPutCustom
	PutDeepseek           inference_put_deepseek.NewPutDeepseek
	PutElasticsearch      inference_put_elasticsearch.NewPutElasticsearch
	PutElser              inference_put_elser.NewPutElser
	PutGoogleaistudio     inference_put_googleaistudio.NewPutGoogleaistudio
	PutGooglevertexai     inference_put_googlevertexai.NewPutGooglevertexai
	PutHuggingFace        inference_put_hugging_face.NewPutHuggingFace
	PutJinaai             inference_put_jinaai.NewPutJinaai
	PutMistral            inference_put_mistral.NewPutMistral
	PutOpenai             inference_put_openai.NewPutOpenai
	PutVoyageai           inference_put_voyageai.NewPutVoyageai
	PutWatsonx            inference_put_watsonx.NewPutWatsonx
	Rerank                inference_rerank.NewRerank
	SparseEmbedding       inference_sparse_embedding.NewSparseEmbedding
	StreamCompletion      inference_stream_completion.NewStreamCompletion
	TextEmbedding         inference_text_embedding.NewTextEmbedding
	Update                inference_update.NewUpdate
}

type Ingest struct {
	DeleteGeoipDatabase      ingest_delete_geoip_database.NewDeleteGeoipDatabase
	DeleteIpLocationDatabase ingest_delete_ip_location_database.NewDeleteIpLocationDatabase
	DeletePipeline           ingest_delete_pipeline.NewDeletePipeline
	GeoIpStats               ingest_geo_ip_stats.NewGeoIpStats
	GetGeoipDatabase         ingest_get_geoip_database.NewGetGeoipDatabase
	GetIpLocationDatabase    ingest_get_ip_location_database.NewGetIpLocationDatabase
	GetPipeline              ingest_get_pipeline.NewGetPipeline
	ProcessorGrok            ingest_processor_grok.NewProcessorGrok
	PutGeoipDatabase         ingest_put_geoip_database.NewPutGeoipDatabase
	PutIpLocationDatabase    ingest_put_ip_location_database.NewPutIpLocationDatabase
	PutPipeline              ingest_put_pipeline.NewPutPipeline
	Simulate                 ingest_simulate.NewSimulate
}

type License struct {
	Delete         license_delete.NewDelete
	Get            license_get.NewGet
	GetBasicStatus license_get_basic_status.NewGetBasicStatus
	GetTrialStatus license_get_trial_status.NewGetTrialStatus
	Post           license_post.NewPost
	PostStartBasic license_post_start_basic.NewPostStartBasic
	PostStartTrial license_post_start_trial.NewPostStartTrial
}

type Logstash struct {
	DeletePipeline logstash_delete_pipeline.NewDeletePipeline
	GetPipeline    logstash_get_pipeline.NewGetPipeline
	PutPipeline    logstash_put_pipeline.NewPutPipeline
}

type Migration struct {
	Deprecations            migration_deprecations.NewDeprecations
	GetFeatureUpgradeStatus migration_get_feature_upgrade_status.NewGetFeatureUpgradeStatus
	PostFeatureUpgrade      migration_post_feature_upgrade.NewPostFeatureUpgrade
}

type Ml struct {
	ClearTrainedModelDeploymentCache ml_clear_trained_model_deployment_cache.NewClearTrainedModelDeploymentCache
	CloseJob                         ml_close_job.NewCloseJob
	DeleteCalendar                   ml_delete_calendar.NewDeleteCalendar
	DeleteCalendarEvent              ml_delete_calendar_event.NewDeleteCalendarEvent
	DeleteCalendarJob                ml_delete_calendar_job.NewDeleteCalendarJob
	DeleteDataFrameAnalytics         ml_delete_data_frame_analytics.NewDeleteDataFrameAnalytics
	DeleteDatafeed                   ml_delete_datafeed.NewDeleteDatafeed
	DeleteExpiredData                ml_delete_expired_data.NewDeleteExpiredData
	DeleteFilter                     ml_delete_filter.NewDeleteFilter
	DeleteForecast                   ml_delete_forecast.NewDeleteForecast
	DeleteJob                        ml_delete_job.NewDeleteJob
	DeleteModelSnapshot              ml_delete_model_snapshot.NewDeleteModelSnapshot
	DeleteTrainedModel               ml_delete_trained_model.NewDeleteTrainedModel
	DeleteTrainedModelAlias          ml_delete_trained_model_alias.NewDeleteTrainedModelAlias
	EstimateModelMemory              ml_estimate_model_memory.NewEstimateModelMemory
	EvaluateDataFrame                ml_evaluate_data_frame.NewEvaluateDataFrame
	ExplainDataFrameAnalytics        ml_explain_data_frame_analytics.NewExplainDataFrameAnalytics
	FlushJob                         ml_flush_job.NewFlushJob
	Forecast                         ml_forecast.NewForecast
	GetBuckets                       ml_get_buckets.NewGetBuckets
	GetCalendarEvents                ml_get_calendar_events.NewGetCalendarEvents
	GetCalendars                     ml_get_calendars.NewGetCalendars
	GetCategories                    ml_get_categories.NewGetCategories
	GetDataFrameAnalytics            ml_get_data_frame_analytics.NewGetDataFrameAnalytics
	GetDataFrameAnalyticsStats       ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStats
	GetDatafeedStats                 ml_get_datafeed_stats.NewGetDatafeedStats
	GetDatafeeds                     ml_get_datafeeds.NewGetDatafeeds
	GetFilters                       ml_get_filters.NewGetFilters
	GetInfluencers                   ml_get_influencers.NewGetInfluencers
	GetJobStats                      ml_get_job_stats.NewGetJobStats
	GetJobs                          ml_get_jobs.NewGetJobs
	GetMemoryStats                   ml_get_memory_stats.NewGetMemoryStats
	GetModelSnapshotUpgradeStats     ml_get_model_snapshot_upgrade_stats.NewGetModelSnapshotUpgradeStats
	GetModelSnapshots                ml_get_model_snapshots.NewGetModelSnapshots
	GetOverallBuckets                ml_get_overall_buckets.NewGetOverallBuckets
	GetRecords                       ml_get_records.NewGetRecords
	GetTrainedModels                 ml_get_trained_models.NewGetTrainedModels
	GetTrainedModelsStats            ml_get_trained_models_stats.NewGetTrainedModelsStats
	InferTrainedModel                ml_infer_trained_model.NewInferTrainedModel
	Info                             ml_info.NewInfo
	OpenJob                          ml_open_job.NewOpenJob
	PostCalendarEvents               ml_post_calendar_events.NewPostCalendarEvents
	PostData                         ml_post_data.NewPostData
	PreviewDataFrameAnalytics        ml_preview_data_frame_analytics.NewPreviewDataFrameAnalytics
	PreviewDatafeed                  ml_preview_datafeed.NewPreviewDatafeed
	PutCalendar                      ml_put_calendar.NewPutCalendar
	PutCalendarJob                   ml_put_calendar_job.NewPutCalendarJob
	PutDataFrameAnalytics            ml_put_data_frame_analytics.NewPutDataFrameAnalytics
	PutDatafeed                      ml_put_datafeed.NewPutDatafeed
	PutFilter                        ml_put_filter.NewPutFilter
	PutJob                           ml_put_job.NewPutJob
	PutTrainedModel                  ml_put_trained_model.NewPutTrainedModel
	PutTrainedModelAlias             ml_put_trained_model_alias.NewPutTrainedModelAlias
	PutTrainedModelDefinitionPart    ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPart
	PutTrainedModelVocabulary        ml_put_trained_model_vocabulary.NewPutTrainedModelVocabulary
	ResetJob                         ml_reset_job.NewResetJob
	RevertModelSnapshot              ml_revert_model_snapshot.NewRevertModelSnapshot
	SetUpgradeMode                   ml_set_upgrade_mode.NewSetUpgradeMode
	StartDataFrameAnalytics          ml_start_data_frame_analytics.NewStartDataFrameAnalytics
	StartDatafeed                    ml_start_datafeed.NewStartDatafeed
	StartTrainedModelDeployment      ml_start_trained_model_deployment.NewStartTrainedModelDeployment
	StopDataFrameAnalytics           ml_stop_data_frame_analytics.NewStopDataFrameAnalytics
	StopDatafeed                     ml_stop_datafeed.NewStopDatafeed
	StopTrainedModelDeployment       ml_stop_trained_model_deployment.NewStopTrainedModelDeployment
	UpdateDataFrameAnalytics         ml_update_data_frame_analytics.NewUpdateDataFrameAnalytics
	UpdateDatafeed                   ml_update_datafeed.NewUpdateDatafeed
	UpdateFilter                     ml_update_filter.NewUpdateFilter
	UpdateJob                        ml_update_job.NewUpdateJob
	UpdateModelSnapshot              ml_update_model_snapshot.NewUpdateModelSnapshot
	UpdateTrainedModelDeployment     ml_update_trained_model_deployment.NewUpdateTrainedModelDeployment
	UpgradeJobSnapshot               ml_upgrade_job_snapshot.NewUpgradeJobSnapshot
	Validate                         ml_validate.NewValidate
	ValidateDetector                 ml_validate_detector.NewValidateDetector
}

type Monitoring struct {
	Bulk monitoring_bulk.NewBulk
}

type Nodes struct {
	ClearRepositoriesMeteringArchive nodes_clear_repositories_metering_archive.NewClearRepositoriesMeteringArchive
	GetRepositoriesMeteringInfo      nodes_get_repositories_metering_info.NewGetRepositoriesMeteringInfo
	HotThreads                       nodes_hot_threads.NewHotThreads
	Info                             nodes_info.NewInfo
	ReloadSecureSettings             nodes_reload_secure_settings.NewReloadSecureSettings
	Stats                            nodes_stats.NewStats
	Usage                            nodes_usage.NewUsage
}

type Profiling struct {
	Flamegraph    profiling_flamegraph.NewFlamegraph
	Stacktraces   profiling_stacktraces.NewStacktraces
	Status        profiling_status.NewStatus
	TopnFunctions profiling_topn_functions.NewTopnFunctions
}

type QueryRules struct {
	DeleteRule    query_rules_delete_rule.NewDeleteRule
	DeleteRuleset query_rules_delete_ruleset.NewDeleteRuleset
	GetRule       query_rules_get_rule.NewGetRule
	GetRuleset    query_rules_get_ruleset.NewGetRuleset
	ListRulesets  query_rules_list_rulesets.NewListRulesets
	PutRule       query_rules_put_rule.NewPutRule
	PutRuleset    query_rules_put_ruleset.NewPutRuleset
	Test          query_rules_test.NewTest
}

type Rollup struct {
	DeleteJob          rollup_delete_job.NewDeleteJob
	GetJobs            rollup_get_jobs.NewGetJobs
	GetRollupCaps      rollup_get_rollup_caps.NewGetRollupCaps
	GetRollupIndexCaps rollup_get_rollup_index_caps.NewGetRollupIndexCaps
	PutJob             rollup_put_job.NewPutJob
	RollupSearch       rollup_rollup_search.NewRollupSearch
	StartJob           rollup_start_job.NewStartJob
	StopJob            rollup_stop_job.NewStopJob
}

type SearchApplication struct {
	Delete                       search_application_delete.NewDelete
	DeleteBehavioralAnalytics    search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalytics
	Get                          search_application_get.NewGet
	GetBehavioralAnalytics       search_application_get_behavioral_analytics.NewGetBehavioralAnalytics
	List                         search_application_list.NewList
	PostBehavioralAnalyticsEvent search_application_post_behavioral_analytics_event.NewPostBehavioralAnalyticsEvent
	Put                          search_application_put.NewPut
	PutBehavioralAnalytics       search_application_put_behavioral_analytics.NewPutBehavioralAnalytics
	RenderQuery                  search_application_render_query.NewRenderQuery
	Search                       search_application_search.NewSearch
}

type SearchableSnapshots struct {
	CacheStats searchable_snapshots_cache_stats.NewCacheStats
	ClearCache searchable_snapshots_clear_cache.NewClearCache
	Mount      searchable_snapshots_mount.NewMount
	Stats      searchable_snapshots_stats.NewStats
}

type Security struct {
	ActivateUserProfile         security_activate_user_profile.NewActivateUserProfile
	Authenticate                security_authenticate.NewAuthenticate
	BulkDeleteRole              security_bulk_delete_role.NewBulkDeleteRole
	BulkPutRole                 security_bulk_put_role.NewBulkPutRole
	BulkUpdateApiKeys           security_bulk_update_api_keys.NewBulkUpdateApiKeys
	ChangePassword              security_change_password.NewChangePassword
	ClearApiKeyCache            security_clear_api_key_cache.NewClearApiKeyCache
	ClearCachedPrivileges       security_clear_cached_privileges.NewClearCachedPrivileges
	ClearCachedRealms           security_clear_cached_realms.NewClearCachedRealms
	ClearCachedRoles            security_clear_cached_roles.NewClearCachedRoles
	ClearCachedServiceTokens    security_clear_cached_service_tokens.NewClearCachedServiceTokens
	CreateApiKey                security_create_api_key.NewCreateApiKey
	CreateCrossClusterApiKey    security_create_cross_cluster_api_key.NewCreateCrossClusterApiKey
	CreateServiceToken          security_create_service_token.NewCreateServiceToken
	DelegatePki                 security_delegate_pki.NewDelegatePki
	DeletePrivileges            security_delete_privileges.NewDeletePrivileges
	DeleteRole                  security_delete_role.NewDeleteRole
	DeleteRoleMapping           security_delete_role_mapping.NewDeleteRoleMapping
	DeleteServiceToken          security_delete_service_token.NewDeleteServiceToken
	DeleteUser                  security_delete_user.NewDeleteUser
	DisableUser                 security_disable_user.NewDisableUser
	DisableUserProfile          security_disable_user_profile.NewDisableUserProfile
	EnableUser                  security_enable_user.NewEnableUser
	EnableUserProfile           security_enable_user_profile.NewEnableUserProfile
	EnrollKibana                security_enroll_kibana.NewEnrollKibana
	EnrollNode                  security_enroll_node.NewEnrollNode
	GetApiKey                   security_get_api_key.NewGetApiKey
	GetBuiltinPrivileges        security_get_builtin_privileges.NewGetBuiltinPrivileges
	GetPrivileges               security_get_privileges.NewGetPrivileges
	GetRole                     security_get_role.NewGetRole
	GetRoleMapping              security_get_role_mapping.NewGetRoleMapping
	GetServiceAccounts          security_get_service_accounts.NewGetServiceAccounts
	GetServiceCredentials       security_get_service_credentials.NewGetServiceCredentials
	GetSettings                 security_get_settings.NewGetSettings
	GetToken                    security_get_token.NewGetToken
	GetUser                     security_get_user.NewGetUser
	GetUserPrivileges           security_get_user_privileges.NewGetUserPrivileges
	GetUserProfile              security_get_user_profile.NewGetUserProfile
	GrantApiKey                 security_grant_api_key.NewGrantApiKey
	HasPrivileges               security_has_privileges.NewHasPrivileges
	HasPrivilegesUserProfile    security_has_privileges_user_profile.NewHasPrivilegesUserProfile
	InvalidateApiKey            security_invalidate_api_key.NewInvalidateApiKey
	InvalidateToken             security_invalidate_token.NewInvalidateToken
	OidcAuthenticate            security_oidc_authenticate.NewOidcAuthenticate
	OidcLogout                  security_oidc_logout.NewOidcLogout
	OidcPrepareAuthentication   security_oidc_prepare_authentication.NewOidcPrepareAuthentication
	PutPrivileges               security_put_privileges.NewPutPrivileges
	PutRole                     security_put_role.NewPutRole
	PutRoleMapping              security_put_role_mapping.NewPutRoleMapping
	PutUser                     security_put_user.NewPutUser
	QueryApiKeys                security_query_api_keys.NewQueryApiKeys
	QueryRole                   security_query_role.NewQueryRole
	QueryUser                   security_query_user.NewQueryUser
	SamlAuthenticate            security_saml_authenticate.NewSamlAuthenticate
	SamlCompleteLogout          security_saml_complete_logout.NewSamlCompleteLogout
	SamlInvalidate              security_saml_invalidate.NewSamlInvalidate
	SamlLogout                  security_saml_logout.NewSamlLogout
	SamlPrepareAuthentication   security_saml_prepare_authentication.NewSamlPrepareAuthentication
	SamlServiceProviderMetadata security_saml_service_provider_metadata.NewSamlServiceProviderMetadata
	SuggestUserProfiles         security_suggest_user_profiles.NewSuggestUserProfiles
	UpdateApiKey                security_update_api_key.NewUpdateApiKey
	UpdateCrossClusterApiKey    security_update_cross_cluster_api_key.NewUpdateCrossClusterApiKey
	UpdateSettings              security_update_settings.NewUpdateSettings
	UpdateUserProfileData       security_update_user_profile_data.NewUpdateUserProfileData
}

type Shutdown struct {
	DeleteNode shutdown_delete_node.NewDeleteNode
	GetNode    shutdown_get_node.NewGetNode
	PutNode    shutdown_put_node.NewPutNode
}

type Simulate struct {
	Ingest simulate_ingest.NewIngest
}

type Slm struct {
	DeleteLifecycle  slm_delete_lifecycle.NewDeleteLifecycle
	ExecuteLifecycle slm_execute_lifecycle.NewExecuteLifecycle
	ExecuteRetention slm_execute_retention.NewExecuteRetention
	GetLifecycle     slm_get_lifecycle.NewGetLifecycle
	GetStats         slm_get_stats.NewGetStats
	GetStatus        slm_get_status.NewGetStatus
	PutLifecycle     slm_put_lifecycle.NewPutLifecycle
	Start            slm_start.NewStart
	Stop             slm_stop.NewStop
}

type Snapshot struct {
	CleanupRepository         snapshot_cleanup_repository.NewCleanupRepository
	Clone                     snapshot_clone.NewClone
	Create                    snapshot_create.NewCreate
	CreateRepository          snapshot_create_repository.NewCreateRepository
	Delete                    snapshot_delete.NewDelete
	DeleteRepository          snapshot_delete_repository.NewDeleteRepository
	Get                       snapshot_get.NewGet
	GetRepository             snapshot_get_repository.NewGetRepository
	RepositoryAnalyze         snapshot_repository_analyze.NewRepositoryAnalyze
	RepositoryVerifyIntegrity snapshot_repository_verify_integrity.NewRepositoryVerifyIntegrity
	Restore                   snapshot_restore.NewRestore
	Status                    snapshot_status.NewStatus
	VerifyRepository          snapshot_verify_repository.NewVerifyRepository
}

type Sql struct {
	ClearCursor    sql_clear_cursor.NewClearCursor
	DeleteAsync    sql_delete_async.NewDeleteAsync
	GetAsync       sql_get_async.NewGetAsync
	GetAsyncStatus sql_get_async_status.NewGetAsyncStatus
	Query          sql_query.NewQuery
	Translate      sql_translate.NewTranslate
}

type Ssl struct {
	Certificates ssl_certificates.NewCertificates
}

type Streams struct {
	LogsDisable streams_logs_disable.NewLogsDisable
	LogsEnable  streams_logs_enable.NewLogsEnable
	Status      streams_status.NewStatus
}

type Synonyms struct {
	DeleteSynonym     synonyms_delete_synonym.NewDeleteSynonym
	DeleteSynonymRule synonyms_delete_synonym_rule.NewDeleteSynonymRule
	GetSynonym        synonyms_get_synonym.NewGetSynonym
	GetSynonymRule    synonyms_get_synonym_rule.NewGetSynonymRule
	GetSynonymsSets   synonyms_get_synonyms_sets.NewGetSynonymsSets
	PutSynonym        synonyms_put_synonym.NewPutSynonym
	PutSynonymRule    synonyms_put_synonym_rule.NewPutSynonymRule
}

type Tasks struct {
	Cancel tasks_cancel.NewCancel
	Get    tasks_get.NewGet
	List   tasks_list.NewList
}

type TextStructure struct {
	FindFieldStructure   text_structure_find_field_structure.NewFindFieldStructure
	FindMessageStructure text_structure_find_message_structure.NewFindMessageStructure
	FindStructure        text_structure_find_structure.NewFindStructure
	TestGrokPattern      text_structure_test_grok_pattern.NewTestGrokPattern
}

type Transform struct {
	DeleteTransform      transform_delete_transform.NewDeleteTransform
	GetNodeStats         transform_get_node_stats.NewGetNodeStats
	GetTransform         transform_get_transform.NewGetTransform
	GetTransformStats    transform_get_transform_stats.NewGetTransformStats
	PreviewTransform     transform_preview_transform.NewPreviewTransform
	PutTransform         transform_put_transform.NewPutTransform
	ResetTransform       transform_reset_transform.NewResetTransform
	ScheduleNowTransform transform_schedule_now_transform.NewScheduleNowTransform
	StartTransform       transform_start_transform.NewStartTransform
	StopTransform        transform_stop_transform.NewStopTransform
	UpdateTransform      transform_update_transform.NewUpdateTransform
	UpgradeTransforms    transform_upgrade_transforms.NewUpgradeTransforms
}

type Watcher struct {
	AckWatch        watcher_ack_watch.NewAckWatch
	ActivateWatch   watcher_activate_watch.NewActivateWatch
	DeactivateWatch watcher_deactivate_watch.NewDeactivateWatch
	DeleteWatch     watcher_delete_watch.NewDeleteWatch
	ExecuteWatch    watcher_execute_watch.NewExecuteWatch
	GetSettings     watcher_get_settings.NewGetSettings
	GetWatch        watcher_get_watch.NewGetWatch
	PutWatch        watcher_put_watch.NewPutWatch
	QueryWatches    watcher_query_watches.NewQueryWatches
	Start           watcher_start.NewStart
	Stats           watcher_stats.NewStats
	Stop            watcher_stop.NewStop
	UpdateSettings  watcher_update_settings.NewUpdateSettings
}

type Xpack struct {
	Info  xpack_info.NewInfo
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
	Streams             Streams
	Synonyms            Synonyms
	Tasks               Tasks
	TextStructure       TextStructure
	Transform           Transform
	Watcher             Watcher
	Xpack               Xpack

	Bulk                    core_bulk.NewBulk
	ClearScroll             core_clear_scroll.NewClearScroll
	ClosePointInTime        core_close_point_in_time.NewClosePointInTime
	Count                   core_count.NewCount
	Create                  core_create.NewCreate
	Delete                  core_delete.NewDelete
	DeleteByQuery           core_delete_by_query.NewDeleteByQuery
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	DeleteScript            core_delete_script.NewDeleteScript
	Exists                  core_exists.NewExists
	ExistsSource            core_exists_source.NewExistsSource
	Explain                 core_explain.NewExplain
	FieldCaps               core_field_caps.NewFieldCaps
	Get                     core_get.NewGet
	GetScript               core_get_script.NewGetScript
	GetScriptContext        core_get_script_context.NewGetScriptContext
	GetScriptLanguages      core_get_script_languages.NewGetScriptLanguages
	GetSource               core_get_source.NewGetSource
	HealthReport            core_health_report.NewHealthReport
	Index                   core_index.NewIndex
	Info                    core_info.NewInfo
	KnnSearch               core_knn_search.NewKnnSearch
	Mget                    core_mget.NewMget
	Msearch                 core_msearch.NewMsearch
	MsearchTemplate         core_msearch_template.NewMsearchTemplate
	Mtermvectors            core_mtermvectors.NewMtermvectors
	OpenPointInTime         core_open_point_in_time.NewOpenPointInTime
	Ping                    core_ping.NewPing
	PutScript               core_put_script.NewPutScript
	RankEval                core_rank_eval.NewRankEval
	Reindex                 core_reindex.NewReindex
	ReindexRethrottle       core_reindex_rethrottle.NewReindexRethrottle
	RenderSearchTemplate    core_render_search_template.NewRenderSearchTemplate
	ScriptsPainlessExecute  core_scripts_painless_execute.NewScriptsPainlessExecute
	Scroll                  core_scroll.NewScroll
	Search                  core_search.NewSearch
	SearchMvt               core_search_mvt.NewSearchMvt
	SearchShards            core_search_shards.NewSearchShards
	SearchTemplate          core_search_template.NewSearchTemplate
	TermsEnum               core_terms_enum.NewTermsEnum
	Termvectors             core_termvectors.NewTermvectors
	Update                  core_update.NewUpdate
	UpdateByQuery           core_update_by_query.NewUpdateByQuery
	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

func New(tp elastictransport.Interface) *API {
	return &API{
		AsyncSearch: AsyncSearch{
			Delete: async_search_delete.NewDeleteFunc(tp),
			Get:    async_search_get.NewGetFunc(tp),
			Status: async_search_status.NewStatusFunc(tp),
			Submit: async_search_submit.NewSubmitFunc(tp),
		},

		Autoscaling: Autoscaling{
			DeleteAutoscalingPolicy: autoscaling_delete_autoscaling_policy.NewDeleteAutoscalingPolicyFunc(tp),
			GetAutoscalingCapacity:  autoscaling_get_autoscaling_capacity.NewGetAutoscalingCapacityFunc(tp),
			GetAutoscalingPolicy:    autoscaling_get_autoscaling_policy.NewGetAutoscalingPolicyFunc(tp),
			PutAutoscalingPolicy:    autoscaling_put_autoscaling_policy.NewPutAutoscalingPolicyFunc(tp),
		},

		Capabilities: Capabilities{
			Capabilities: capabilities.NewCapabilitiesFunc(tp),
		},

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

		DanglingIndices: DanglingIndices{
			DeleteDanglingIndex: dangling_indices_delete_dangling_index.NewDeleteDanglingIndexFunc(tp),
			ImportDanglingIndex: dangling_indices_import_dangling_index.NewImportDanglingIndexFunc(tp),
			ListDanglingIndices: dangling_indices_list_dangling_indices.NewListDanglingIndicesFunc(tp),
		},

		Enrich: Enrich{
			DeletePolicy:  enrich_delete_policy.NewDeletePolicyFunc(tp),
			ExecutePolicy: enrich_execute_policy.NewExecutePolicyFunc(tp),
			GetPolicy:     enrich_get_policy.NewGetPolicyFunc(tp),
			PutPolicy:     enrich_put_policy.NewPutPolicyFunc(tp),
			Stats:         enrich_stats.NewStatsFunc(tp),
		},

		Eql: Eql{
			Delete:    eql_delete.NewDeleteFunc(tp),
			Get:       eql_get.NewGetFunc(tp),
			GetStatus: eql_get_status.NewGetStatusFunc(tp),
			Search:    eql_search.NewSearchFunc(tp),
		},

		Esql: Esql{
			AsyncQuery:       esql_async_query.NewAsyncQueryFunc(tp),
			AsyncQueryDelete: esql_async_query_delete.NewAsyncQueryDeleteFunc(tp),
			AsyncQueryGet:    esql_async_query_get.NewAsyncQueryGetFunc(tp),
			AsyncQueryStop:   esql_async_query_stop.NewAsyncQueryStopFunc(tp),
			Query:            esql_query.NewQueryFunc(tp),
		},

		Features: Features{
			GetFeatures:   features_get_features.NewGetFeaturesFunc(tp),
			ResetFeatures: features_reset_features.NewResetFeaturesFunc(tp),
		},

		Fleet: Fleet{
			GlobalCheckpoints: fleet_global_checkpoints.NewGlobalCheckpointsFunc(tp),
			Msearch:           fleet_msearch.NewMsearchFunc(tp),
			PostSecret:        fleet_post_secret.NewPostSecretFunc(tp),
			Search:            fleet_search.NewSearchFunc(tp),
		},

		Graph: Graph{
			Explore: graph_explore.NewExploreFunc(tp),
		},

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

		Inference: Inference{
			ChatCompletionUnified: inference_chat_completion_unified.NewChatCompletionUnifiedFunc(tp),
			Completion:            inference_completion.NewCompletionFunc(tp),
			Delete:                inference_delete.NewDeleteFunc(tp),
			Get:                   inference_get.NewGetFunc(tp),
			Inference:             inference_inference.NewInferenceFunc(tp),
			Put:                   inference_put.NewPutFunc(tp),
			PutAlibabacloud:       inference_put_alibabacloud.NewPutAlibabacloudFunc(tp),
			PutAmazonbedrock:      inference_put_amazonbedrock.NewPutAmazonbedrockFunc(tp),
			PutAmazonsagemaker:    inference_put_amazonsagemaker.NewPutAmazonsagemakerFunc(tp),
			PutAnthropic:          inference_put_anthropic.NewPutAnthropicFunc(tp),
			PutAzureaistudio:      inference_put_azureaistudio.NewPutAzureaistudioFunc(tp),
			PutAzureopenai:        inference_put_azureopenai.NewPutAzureopenaiFunc(tp),
			PutCohere:             inference_put_cohere.NewPutCohereFunc(tp),
			PutCustom:             inference_put_custom.NewPutCustomFunc(tp),
			PutDeepseek:           inference_put_deepseek.NewPutDeepseekFunc(tp),
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

		License: License{
			Delete:         license_delete.NewDeleteFunc(tp),
			Get:            license_get.NewGetFunc(tp),
			GetBasicStatus: license_get_basic_status.NewGetBasicStatusFunc(tp),
			GetTrialStatus: license_get_trial_status.NewGetTrialStatusFunc(tp),
			Post:           license_post.NewPostFunc(tp),
			PostStartBasic: license_post_start_basic.NewPostStartBasicFunc(tp),
			PostStartTrial: license_post_start_trial.NewPostStartTrialFunc(tp),
		},

		Logstash: Logstash{
			DeletePipeline: logstash_delete_pipeline.NewDeletePipelineFunc(tp),
			GetPipeline:    logstash_get_pipeline.NewGetPipelineFunc(tp),
			PutPipeline:    logstash_put_pipeline.NewPutPipelineFunc(tp),
		},

		Migration: Migration{
			Deprecations:            migration_deprecations.NewDeprecationsFunc(tp),
			GetFeatureUpgradeStatus: migration_get_feature_upgrade_status.NewGetFeatureUpgradeStatusFunc(tp),
			PostFeatureUpgrade:      migration_post_feature_upgrade.NewPostFeatureUpgradeFunc(tp),
		},

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

		Monitoring: Monitoring{
			Bulk: monitoring_bulk.NewBulkFunc(tp),
		},

		Nodes: Nodes{
			ClearRepositoriesMeteringArchive: nodes_clear_repositories_metering_archive.NewClearRepositoriesMeteringArchiveFunc(tp),
			GetRepositoriesMeteringInfo:      nodes_get_repositories_metering_info.NewGetRepositoriesMeteringInfoFunc(tp),
			HotThreads:                       nodes_hot_threads.NewHotThreadsFunc(tp),
			Info:                             nodes_info.NewInfoFunc(tp),
			ReloadSecureSettings:             nodes_reload_secure_settings.NewReloadSecureSettingsFunc(tp),
			Stats:                            nodes_stats.NewStatsFunc(tp),
			Usage:                            nodes_usage.NewUsageFunc(tp),
		},

		Profiling: Profiling{
			Flamegraph:    profiling_flamegraph.NewFlamegraphFunc(tp),
			Stacktraces:   profiling_stacktraces.NewStacktracesFunc(tp),
			Status:        profiling_status.NewStatusFunc(tp),
			TopnFunctions: profiling_topn_functions.NewTopnFunctionsFunc(tp),
		},

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

		SearchableSnapshots: SearchableSnapshots{
			CacheStats: searchable_snapshots_cache_stats.NewCacheStatsFunc(tp),
			ClearCache: searchable_snapshots_clear_cache.NewClearCacheFunc(tp),
			Mount:      searchable_snapshots_mount.NewMountFunc(tp),
			Stats:      searchable_snapshots_stats.NewStatsFunc(tp),
		},

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

		Shutdown: Shutdown{
			DeleteNode: shutdown_delete_node.NewDeleteNodeFunc(tp),
			GetNode:    shutdown_get_node.NewGetNodeFunc(tp),
			PutNode:    shutdown_put_node.NewPutNodeFunc(tp),
		},

		Simulate: Simulate{
			Ingest: simulate_ingest.NewIngestFunc(tp),
		},

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

		Sql: Sql{
			ClearCursor:    sql_clear_cursor.NewClearCursorFunc(tp),
			DeleteAsync:    sql_delete_async.NewDeleteAsyncFunc(tp),
			GetAsync:       sql_get_async.NewGetAsyncFunc(tp),
			GetAsyncStatus: sql_get_async_status.NewGetAsyncStatusFunc(tp),
			Query:          sql_query.NewQueryFunc(tp),
			Translate:      sql_translate.NewTranslateFunc(tp),
		},

		Ssl: Ssl{
			Certificates: ssl_certificates.NewCertificatesFunc(tp),
		},

		Streams: Streams{
			LogsDisable: streams_logs_disable.NewLogsDisableFunc(tp),
			LogsEnable:  streams_logs_enable.NewLogsEnableFunc(tp),
			Status:      streams_status.NewStatusFunc(tp),
		},

		Synonyms: Synonyms{
			DeleteSynonym:     synonyms_delete_synonym.NewDeleteSynonymFunc(tp),
			DeleteSynonymRule: synonyms_delete_synonym_rule.NewDeleteSynonymRuleFunc(tp),
			GetSynonym:        synonyms_get_synonym.NewGetSynonymFunc(tp),
			GetSynonymRule:    synonyms_get_synonym_rule.NewGetSynonymRuleFunc(tp),
			GetSynonymsSets:   synonyms_get_synonyms_sets.NewGetSynonymsSetsFunc(tp),
			PutSynonym:        synonyms_put_synonym.NewPutSynonymFunc(tp),
			PutSynonymRule:    synonyms_put_synonym_rule.NewPutSynonymRuleFunc(tp),
		},

		Tasks: Tasks{
			Cancel: tasks_cancel.NewCancelFunc(tp),
			Get:    tasks_get.NewGetFunc(tp),
			List:   tasks_list.NewListFunc(tp),
		},

		TextStructure: TextStructure{
			FindFieldStructure:   text_structure_find_field_structure.NewFindFieldStructureFunc(tp),
			FindMessageStructure: text_structure_find_message_structure.NewFindMessageStructureFunc(tp),
			FindStructure:        text_structure_find_structure.NewFindStructureFunc(tp),
			TestGrokPattern:      text_structure_test_grok_pattern.NewTestGrokPatternFunc(tp),
		},

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

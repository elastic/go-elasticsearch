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
	features_get_features "github.com/elastic/go-elasticsearch/v8/typedapi/features/getfeatures"
	features_reset_features "github.com/elastic/go-elasticsearch/v8/typedapi/features/resetfeatures"
	fleet_global_checkpoints "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/globalcheckpoints"
	fleet_msearch "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/msearch"
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
	query_ruleset_delete "github.com/elastic/go-elasticsearch/v8/typedapi/queryruleset/delete"
	query_ruleset_get "github.com/elastic/go-elasticsearch/v8/typedapi/queryruleset/get"
	query_ruleset_list "github.com/elastic/go-elasticsearch/v8/typedapi/queryruleset/list"
	query_ruleset_put "github.com/elastic/go-elasticsearch/v8/typedapi/queryruleset/put"
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
	security_saml_authenticate "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlauthenticate"
	security_saml_complete_logout "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlcompletelogout"
	security_saml_invalidate "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlinvalidate"
	security_saml_logout "github.com/elastic/go-elasticsearch/v8/typedapi/security/samllogout"
	security_saml_prepare_authentication "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlprepareauthentication"
	security_saml_service_provider_metadata "github.com/elastic/go-elasticsearch/v8/typedapi/security/samlserviceprovidermetadata"
	security_suggest_user_profiles "github.com/elastic/go-elasticsearch/v8/typedapi/security/suggestuserprofiles"
	security_update_api_key "github.com/elastic/go-elasticsearch/v8/typedapi/security/updateapikey"
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
	text_structure_find_structure "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/findstructure"
	transform_delete_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/deletetransform"
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
	// Deletes an async search by ID. If the search is still running, the search
	// request will be cancelled. Otherwise, the saved search results are deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Delete async_search_delete.NewDelete
	// Retrieves the results of a previously submitted async search request given
	// its ID.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Get async_search_get.NewGet
	// Retrieves the status of a previously submitted async search request given its
	// ID.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Status async_search_status.NewStatus
	// Executes a search request asynchronously.
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

type Cat struct {
	// Shows information about currently configured aliases to indices including
	// filter and routing infos.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-alias.html
	Aliases cat_aliases.NewAliases
	// Provides a snapshot of how many shards are allocated to each data node and
	// how much disk space they are using.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-allocation.html
	Allocation cat_allocation.NewAllocation
	// Returns information about existing component_templates templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-component-templates.html
	ComponentTemplates cat_component_templates.NewComponentTemplates
	// Provides quick access to the document count of the entire cluster, or
	// individual indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-count.html
	Count cat_count.NewCount
	// Shows how much heap memory is currently being used by fielddata on every data
	// node in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-fielddata.html
	Fielddata cat_fielddata.NewFielddata
	// Returns a concise representation of the cluster health.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-health.html
	Health cat_health.NewHealth
	// Returns help for the Cat APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat.html
	Help cat_help.NewHelp
	// Returns information about indices: number of primaries and replicas, document
	// counts, disk size, ...
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html
	Indices cat_indices.NewIndices
	// Returns information about the master node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-master.html
	Master cat_master.NewMaster
	// Gets configuration and usage information about data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-dfanalytics.html
	MlDataFrameAnalytics cat_ml_data_frame_analytics.NewMlDataFrameAnalytics
	// Gets configuration and usage information about datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-datafeeds.html
	MlDatafeeds cat_ml_datafeeds.NewMlDatafeeds
	// Gets configuration and usage information about anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-anomaly-detectors.html
	MlJobs cat_ml_jobs.NewMlJobs
	// Gets configuration and usage information about inference trained models.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-trained-model.html
	MlTrainedModels cat_ml_trained_models.NewMlTrainedModels
	// Returns information about custom node attributes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodeattrs.html
	Nodeattrs cat_nodeattrs.NewNodeattrs
	// Returns basic statistics about performance of cluster nodes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodes.html
	Nodes cat_nodes.NewNodes
	// Returns a concise representation of the cluster pending tasks.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-pending-tasks.html
	PendingTasks cat_pending_tasks.NewPendingTasks
	// Returns information about installed plugins across nodes node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-plugins.html
	Plugins cat_plugins.NewPlugins
	// Returns information about index shard recoveries, both on-going completed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-recovery.html
	Recovery cat_recovery.NewRecovery
	// Returns information about snapshot repositories registered in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-repositories.html
	Repositories cat_repositories.NewRepositories
	// Provides low-level information about the segments in the shards of an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-segments.html
	Segments cat_segments.NewSegments
	// Provides a detailed view of shard allocation on nodes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-shards.html
	Shards cat_shards.NewShards
	// Returns all snapshots in a specific repository.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-snapshots.html
	Snapshots cat_snapshots.NewSnapshots
	// Returns information about the tasks currently executing on one or more nodes
	// in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	Tasks cat_tasks.NewTasks
	// Returns information about existing templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-templates.html
	Templates cat_templates.NewTemplates
	// Returns cluster-wide thread pool statistics per node.
	// By default the active, queue and rejected statistics are returned for all
	// thread pools.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-thread-pool.html
	ThreadPool cat_thread_pool.NewThreadPool
	// Gets configuration and usage information about transforms.
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
	// Deletes a component template
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	DeleteComponentTemplate cluster_delete_component_template.NewDeleteComponentTemplate
	// Clears cluster voting config exclusions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/voting-config-exclusions.html
	DeleteVotingConfigExclusions cluster_delete_voting_config_exclusions.NewDeleteVotingConfigExclusions
	// Returns information about whether a particular component template exist
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	ExistsComponentTemplate cluster_exists_component_template.NewExistsComponentTemplate
	// Returns one or more component templates
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	GetComponentTemplate cluster_get_component_template.NewGetComponentTemplate
	// Returns cluster settings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-get-settings.html
	GetSettings cluster_get_settings.NewGetSettings
	// Returns basic information about the health of the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html
	Health cluster_health.NewHealth
	// Returns different information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-info.html
	Info cluster_info.NewInfo
	// Returns a list of any cluster-level changes (e.g. create index, update
	// mapping,
	// allocate or fail shard) which have not yet been executed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-pending.html
	PendingTasks cluster_pending_tasks.NewPendingTasks
	// Updates the cluster voting config exclusions by node ids or node names.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/voting-config-exclusions.html
	PostVotingConfigExclusions cluster_post_voting_config_exclusions.NewPostVotingConfigExclusions
	// Creates or updates a component template
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	PutComponentTemplate cluster_put_component_template.NewPutComponentTemplate
	// Updates the cluster settings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-update-settings.html
	PutSettings cluster_put_settings.NewPutSettings
	// Returns the information about configured remote clusters.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-remote-info.html
	RemoteInfo cluster_remote_info.NewRemoteInfo
	// Allows to manually change the allocation of individual shards in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-reroute.html
	Reroute cluster_reroute.NewReroute
	// Returns a comprehensive information about the state of the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-state.html
	State cluster_state.NewState
	// Returns high-level overview of cluster statistics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-stats.html
	Stats cluster_stats.NewStats
}

type Core struct {
	// Allows to perform multiple index/update/delete operations in a single
	// request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Explicitly clears the search context for a scroll.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Close a point in time
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Creates a new document in the index.
	//
	// Returns a 409 response when a document with a same ID already exists in the
	// index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Removes a document from the index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Deletes documents matching the provided query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Changes the number of requests per second for a particular Delete By Query
	// operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	// Deletes a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Returns information about whether a document exists in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Returns information about whether a document source exists in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Returns information about why a specific matches (or doesn't match) a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// Returns the information about the capabilities of fields among multiple
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Returns a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Returns a script.
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
	// Creates or updates a document in an index.
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
	// Allows to execute several search template operations in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Returns multiple termvectors in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// Open a point in time that can be used in subsequent searches
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Creates or updates a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Allows to evaluate the quality of ranked search results over a set of typical
	// search queries
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Allows to copy documents from one index to another, optionally filtering the
	// source
	// documents by a query, changing the destination index settings, or fetching
	// the
	// documents from a remote cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Changes the number of requests per second for a particular Reindex operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle
	// Allows to use the Mustache language to pre-render a search definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Allows an arbitrary script to be executed and a result to be returned
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Allows to retrieve a large numbers of results from a single search request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Returns results matching a query.
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
	// Allows to use the Mustache language to pre-render a search definition.
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
	// Performs an update on every document in the index without changing the
	// source,
	// for example to pick up a mapping change.
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
	// Gets information about an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-enrich-policy-api.html
	GetPolicy enrich_get_policy.NewGetPolicy
	// Creates a new enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-enrich-policy-api.html
	PutPolicy enrich_put_policy.NewPutPolicy
	// Gets enrich coordinator statistics and information about enrich policies that
	// are currently executing.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-stats-api.html
	Stats enrich_stats.NewStats
}

type Eql struct {
	// Deletes an async EQL search by ID. If the search is still running, the search
	// request will be cancelled. Otherwise, the saved search results are deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Delete eql_delete.NewDelete
	// Returns async results from previously executed Event Query Language (EQL)
	// search
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-search-api.html
	Get eql_get.NewGet
	// Returns the status of a previously submitted async or stored Event Query
	// Language (EQL) search
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-status-api.html
	GetStatus eql_get_status.NewGetStatus
	// Returns results matching a query expressed in Event Query Language (EQL)
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Search eql_search.NewSearch
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
	// Multi Search API where the search will only be executed after specified
	// checkpoints are available due to a refresh. This API is designed for internal
	// use by the fleet server project.
	//
	Msearch fleet_msearch.NewMsearch
	// Search API where the search will only be executed after specified checkpoints
	// are available due to a refresh. This API is designed for internal use by the
	// fleet server project.
	//
	Search fleet_search.NewSearch
}

type Graph struct {
	// Explore extracted and summarized information about the documents and terms in
	// an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/graph-explore-api.html
	Explore graph_explore.NewExplore
}

type Ilm struct {
	// Deletes the specified lifecycle policy definition. A currently used policy
	// cannot be deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-delete-lifecycle.html
	DeleteLifecycle ilm_delete_lifecycle.NewDeleteLifecycle
	// Retrieves information about the index's current lifecycle state, such as the
	// currently executing phase, action, and step.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-explain-lifecycle.html
	ExplainLifecycle ilm_explain_lifecycle.NewExplainLifecycle
	// Returns the specified policy definition. Includes the policy version and last
	// modified date.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-lifecycle.html
	GetLifecycle ilm_get_lifecycle.NewGetLifecycle
	// Retrieves the current index lifecycle management (ILM) status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-status.html
	GetStatus ilm_get_status.NewGetStatus
	// Migrates the indices and ILM policies away from custom node attribute
	// allocation routing to data tiers routing
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-migrate-to-data-tiers.html
	MigrateToDataTiers ilm_migrate_to_data_tiers.NewMigrateToDataTiers
	// Manually moves an index into the specified step and executes that step.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-move-to-step.html
	MoveToStep ilm_move_to_step.NewMoveToStep
	// Creates a lifecycle policy
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
	// Performs the analysis process on a text and return the tokens breakdown of
	// the text.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-analyze.html
	Analyze indices_analyze.NewAnalyze
	// Clears all or specific caches for one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-clearcache.html
	ClearCache indices_clear_cache.NewClearCache
	// Clones an index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-clone-index.html
	Clone indices_clone.NewClone
	// Closes an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-close.html
	Close indices_close.NewClose
	// Creates an index with optional settings and mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
	Create indices_create.NewCreate
	// Creates a data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	CreateDataStream indices_create_data_stream.NewCreateDataStream
	// Provides statistics on operations happening in a data stream.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	DataStreamsStats indices_data_streams_stats.NewDataStreamsStats
	// Deletes an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-index.html
	Delete indices_delete.NewDelete
	// Deletes an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	DeleteAlias indices_delete_alias.NewDeleteAlias
	// Deletes the data stream lifecycle of the selected data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-delete-lifecycle.html
	DeleteDataLifecycle indices_delete_data_lifecycle.NewDeleteDataLifecycle
	// Deletes a data stream.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	DeleteDataStream indices_delete_data_stream.NewDeleteDataStream
	// Deletes an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	DeleteIndexTemplate indices_delete_index_template.NewDeleteIndexTemplate
	// Deletes an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	DeleteTemplate indices_delete_template.NewDeleteTemplate
	// Analyzes the disk usage of each field of an index or data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-disk-usage.html
	DiskUsage indices_disk_usage.NewDiskUsage
	// Downsample an index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-downsample-data-stream.html
	Downsample indices_downsample.NewDownsample
	// Returns information about whether a particular index exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-exists.html
	Exists indices_exists.NewExists
	// Returns information about whether a particular alias exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	ExistsAlias indices_exists_alias.NewExistsAlias
	// Returns information about whether a particular index template exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	ExistsIndexTemplate indices_exists_index_template.NewExistsIndexTemplate
	// Returns information about whether a particular index template exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	ExistsTemplate indices_exists_template.NewExistsTemplate
	// Retrieves information about the index's current data stream lifecycle, such
	// as any potential encountered error, time since creation etc.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-explain-lifecycle.html
	ExplainDataLifecycle indices_explain_data_lifecycle.NewExplainDataLifecycle
	// Returns the field usage stats for each field of an index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/field-usage-stats.html
	FieldUsageStats indices_field_usage_stats.NewFieldUsageStats
	// Performs the flush operation on one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-flush.html
	Flush indices_flush.NewFlush
	// Performs the force merge operation on one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-forcemerge.html
	Forcemerge indices_forcemerge.NewForcemerge
	// Returns information about one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-index.html
	Get indices_get.NewGet
	// Returns an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	GetAlias indices_get_alias.NewGetAlias
	// Returns the data stream lifecycle of the selected data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-get-lifecycle.html
	GetDataLifecycle indices_get_data_lifecycle.NewGetDataLifecycle
	// Returns data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	GetDataStream indices_get_data_stream.NewGetDataStream
	// Returns mapping for one or more fields.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-field-mapping.html
	GetFieldMapping indices_get_field_mapping.NewGetFieldMapping
	// Returns an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	GetIndexTemplate indices_get_index_template.NewGetIndexTemplate
	// Returns mappings for one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-mapping.html
	GetMapping indices_get_mapping.NewGetMapping
	// Returns settings for one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-settings.html
	GetSettings indices_get_settings.NewGetSettings
	// Returns an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	GetTemplate indices_get_template.NewGetTemplate
	// Migrates an alias to a data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	MigrateToDataStream indices_migrate_to_data_stream.NewMigrateToDataStream
	// Modifies a data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	ModifyDataStream indices_modify_data_stream.NewModifyDataStream
	// Opens an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-open-close.html
	Open indices_open.NewOpen
	// Promotes a data stream from a replicated data stream managed by CCR to a
	// regular data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	PromoteDataStream indices_promote_data_stream.NewPromoteDataStream
	// Creates or updates an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	PutAlias indices_put_alias.NewPutAlias
	// Updates the data stream lifecycle of the selected data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-put-lifecycle.html
	PutDataLifecycle indices_put_data_lifecycle.NewPutDataLifecycle
	// Creates or updates an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	PutIndexTemplate indices_put_index_template.NewPutIndexTemplate
	// Updates the index mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-mapping.html
	PutMapping indices_put_mapping.NewPutMapping
	// Updates the index settings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-update-settings.html
	PutSettings indices_put_settings.NewPutSettings
	// Creates or updates an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	PutTemplate indices_put_template.NewPutTemplate
	// Returns information about ongoing index shard recoveries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-recovery.html
	Recovery indices_recovery.NewRecovery
	// Performs the refresh operation in one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-refresh.html
	Refresh indices_refresh.NewRefresh
	// Reloads an index's search analyzers and their resources.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-reload-analyzers.html
	ReloadSearchAnalyzers indices_reload_search_analyzers.NewReloadSearchAnalyzers
	// Returns information about any matching indices, aliases, and data streams
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-resolve-index-api.html
	ResolveIndex indices_resolve_index.NewResolveIndex
	// Updates an alias to point to a new index when the existing index
	// is considered to be too large or too old.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-rollover-index.html
	Rollover indices_rollover.NewRollover
	// Provides low-level information about segments in a Lucene index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-segments.html
	Segments indices_segments.NewSegments
	// Provides store information for shard copies of indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shards-stores.html
	ShardStores indices_shard_stores.NewShardStores
	// Allow to shrink an existing index into a new index with fewer primary shards.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shrink-index.html
	Shrink indices_shrink.NewShrink
	// Simulate matching the given index name against the index templates in the
	// system
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	SimulateIndexTemplate indices_simulate_index_template.NewSimulateIndexTemplate
	// Simulate resolving the given template name or body
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	SimulateTemplate indices_simulate_template.NewSimulateTemplate
	// Allows you to split an existing index into a new index with more primary
	// shards.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-split-index.html
	Split indices_split.NewSplit
	// Provides statistics on operations happening in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-stats.html
	Stats indices_stats.NewStats
	// Unfreezes an index. When a frozen index is unfrozen, the index goes through
	// the normal recovery process and becomes writeable again.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/unfreeze-index-api.html
	Unfreeze indices_unfreeze.NewUnfreeze
	// Updates index aliases.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	UpdateAliases indices_update_aliases.NewUpdateAliases
	// Allows a user to validate a potentially expensive query without executing it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-validate.html
	ValidateQuery indices_validate_query.NewValidateQuery
}

type Ingest struct {
	// Deletes a pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-pipeline-api.html
	DeletePipeline ingest_delete_pipeline.NewDeletePipeline
	// Returns statistical information about geoip databases
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html
	GeoIpStats ingest_geo_ip_stats.NewGeoIpStats
	// Returns a pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-pipeline-api.html
	GetPipeline ingest_get_pipeline.NewGetPipeline
	// Returns a list of the built-in patterns.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html
	ProcessorGrok ingest_processor_grok.NewProcessorGrok
	// Creates or updates a pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html
	PutPipeline ingest_put_pipeline.NewPutPipeline
	// Allows to simulate a pipeline with example documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-pipeline-api.html
	Simulate ingest_simulate.NewSimulate
}

type License struct {
	// Deletes licensing information for the cluster
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-license.html
	Delete license_delete.NewDelete
	// Retrieves licensing information for the cluster
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
	// Starts an indefinite basic license.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-basic.html
	PostStartBasic license_post_start_basic.NewPostStartBasic
	// starts a limited time trial license.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trial.html
	PostStartTrial license_post_start_trial.NewPostStartTrial
}

type Logstash struct {
	// Deletes Logstash Pipelines used by Central Management
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-delete-pipeline.html
	DeletePipeline logstash_delete_pipeline.NewDeletePipeline
	// Retrieves Logstash Pipelines used by Central Management
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-get-pipeline.html
	GetPipeline logstash_get_pipeline.NewGetPipeline
	// Adds and updates Logstash Pipelines used for Central Management
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
	// Clear the cached results from a trained model deployment
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-trained-model-deployment-cache.html
	ClearTrainedModelDeploymentCache ml_clear_trained_model_deployment_cache.NewClearTrainedModelDeploymentCache
	// Closes one or more anomaly detection jobs. A job can be opened and closed
	// multiple times throughout its lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-close-job.html
	CloseJob ml_close_job.NewCloseJob
	// Deletes a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar.html
	DeleteCalendar ml_delete_calendar.NewDeleteCalendar
	// Deletes scheduled events from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-event.html
	DeleteCalendarEvent ml_delete_calendar_event.NewDeleteCalendarEvent
	// Deletes anomaly detection jobs from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-job.html
	DeleteCalendarJob ml_delete_calendar_job.NewDeleteCalendarJob
	// Deletes an existing data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-dfanalytics.html
	DeleteDataFrameAnalytics ml_delete_data_frame_analytics.NewDeleteDataFrameAnalytics
	// Deletes an existing datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-datafeed.html
	DeleteDatafeed ml_delete_datafeed.NewDeleteDatafeed
	// Deletes expired and unused machine learning data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-expired-data.html
	DeleteExpiredData ml_delete_expired_data.NewDeleteExpiredData
	// Deletes a filter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-filter.html
	DeleteFilter ml_delete_filter.NewDeleteFilter
	// Deletes forecasts from a machine learning job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-forecast.html
	DeleteForecast ml_delete_forecast.NewDeleteForecast
	// Deletes an existing anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-job.html
	DeleteJob ml_delete_job.NewDeleteJob
	// Deletes an existing model snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-snapshot.html
	DeleteModelSnapshot ml_delete_model_snapshot.NewDeleteModelSnapshot
	// Deletes an existing trained inference model that is currently not referenced
	// by an ingest pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models.html
	DeleteTrainedModel ml_delete_trained_model.NewDeleteTrainedModel
	// Deletes a model alias that refers to the trained model
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models-aliases.html
	DeleteTrainedModelAlias ml_delete_trained_model_alias.NewDeleteTrainedModelAlias
	// Estimates the model memory
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-apis.html
	EstimateModelMemory ml_estimate_model_memory.NewEstimateModelMemory
	// Evaluates the data frame analytics for an annotated index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/evaluate-dfanalytics.html
	EvaluateDataFrame ml_evaluate_data_frame.NewEvaluateDataFrame
	// Explains a data frame analytics config.
	// http://www.elastic.co/guide/en/elasticsearch/reference/current/explain-dfanalytics.html
	ExplainDataFrameAnalytics ml_explain_data_frame_analytics.NewExplainDataFrameAnalytics
	// Forces any buffered data to be processed by the job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-flush-job.html
	FlushJob ml_flush_job.NewFlushJob
	// Predicts the future behavior of a time series by using its historical
	// behavior.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-forecast.html
	Forecast ml_forecast.NewForecast
	// Retrieves anomaly detection job results for one or more buckets.
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
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics.html
	GetDataFrameAnalytics ml_get_data_frame_analytics.NewGetDataFrameAnalytics
	// Retrieves usage information for data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics-stats.html
	GetDataFrameAnalyticsStats ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStats
	// Retrieves usage information for datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed-stats.html
	GetDatafeedStats ml_get_datafeed_stats.NewGetDatafeedStats
	// Retrieves configuration information for datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed.html
	GetDatafeeds ml_get_datafeeds.NewGetDatafeeds
	// Retrieves filters.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-filter.html
	GetFilters ml_get_filters.NewGetFilters
	// Retrieves anomaly detection job results for one or more influencers.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-influencer.html
	GetInfluencers ml_get_influencers.NewGetInfluencers
	// Retrieves usage information for anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-stats.html
	GetJobStats ml_get_job_stats.NewGetJobStats
	// Retrieves configuration information for anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job.html
	GetJobs ml_get_jobs.NewGetJobs
	// Returns information on how ML is using memory.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-memory.html
	GetMemoryStats ml_get_memory_stats.NewGetMemoryStats
	// Gets stats for anomaly detection job model snapshot upgrades that are in
	// progress.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-model-snapshot-upgrade-stats.html
	GetModelSnapshotUpgradeStats ml_get_model_snapshot_upgrade_stats.NewGetModelSnapshotUpgradeStats
	// Retrieves information about model snapshots.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-snapshot.html
	GetModelSnapshots ml_get_model_snapshots.NewGetModelSnapshots
	// Retrieves overall bucket results that summarize the bucket results of
	// multiple anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-overall-buckets.html
	GetOverallBuckets ml_get_overall_buckets.NewGetOverallBuckets
	// Retrieves anomaly records for an anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-record.html
	GetRecords ml_get_records.NewGetRecords
	// Retrieves configuration information for a trained inference model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models.html
	GetTrainedModels ml_get_trained_models.NewGetTrainedModels
	// Retrieves usage information for trained inference models.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models-stats.html
	GetTrainedModelsStats ml_get_trained_models_stats.NewGetTrainedModelsStats
	// Evaluate a trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-trained-model.html
	InferTrainedModel ml_infer_trained_model.NewInferTrainedModel
	// Returns defaults and limits used by machine learning.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-info.html
	Info ml_info.NewInfo
	// Opens one or more anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-open-job.html
	OpenJob ml_open_job.NewOpenJob
	// Posts scheduled events in a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-calendar-event.html
	PostCalendarEvents ml_post_calendar_events.NewPostCalendarEvents
	// Sends data to an anomaly detection job for analysis.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-data.html
	PostData ml_post_data.NewPostData
	// Previews that will be analyzed given a data frame analytics config.
	// http://www.elastic.co/guide/en/elasticsearch/reference/current/preview-dfanalytics.html
	PreviewDataFrameAnalytics ml_preview_data_frame_analytics.NewPreviewDataFrameAnalytics
	// Previews a datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-preview-datafeed.html
	PreviewDatafeed ml_preview_datafeed.NewPreviewDatafeed
	// Instantiates a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar.html
	PutCalendar ml_put_calendar.NewPutCalendar
	// Adds an anomaly detection job to a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar-job.html
	PutCalendarJob ml_put_calendar_job.NewPutCalendarJob
	// Instantiates a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-dfanalytics.html
	PutDataFrameAnalytics ml_put_data_frame_analytics.NewPutDataFrameAnalytics
	// Instantiates a datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-datafeed.html
	PutDatafeed ml_put_datafeed.NewPutDatafeed
	// Instantiates a filter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-filter.html
	PutFilter ml_put_filter.NewPutFilter
	// Instantiates an anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-job.html
	PutJob ml_put_job.NewPutJob
	// Creates an inference trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models.html
	PutTrainedModel ml_put_trained_model.NewPutTrainedModel
	// Creates a new model alias (or reassigns an existing one) to refer to the
	// trained model
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models-aliases.html
	PutTrainedModelAlias ml_put_trained_model_alias.NewPutTrainedModelAlias
	// Creates part of a trained model definition
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-definition-part.html
	PutTrainedModelDefinitionPart ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPart
	// Creates a trained model vocabulary
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-vocabulary.html
	PutTrainedModelVocabulary ml_put_trained_model_vocabulary.NewPutTrainedModelVocabulary
	// Resets an existing anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-reset-job.html
	ResetJob ml_reset_job.NewResetJob
	// Reverts to a specific snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-revert-snapshot.html
	RevertModelSnapshot ml_revert_model_snapshot.NewRevertModelSnapshot
	// Sets a cluster wide upgrade_mode setting that prepares machine learning
	// indices for an upgrade.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-set-upgrade-mode.html
	SetUpgradeMode ml_set_upgrade_mode.NewSetUpgradeMode
	// Starts a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-dfanalytics.html
	StartDataFrameAnalytics ml_start_data_frame_analytics.NewStartDataFrameAnalytics
	// Starts one or more datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-start-datafeed.html
	StartDatafeed ml_start_datafeed.NewStartDatafeed
	// Start a trained model deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trained-model-deployment.html
	StartTrainedModelDeployment ml_start_trained_model_deployment.NewStartTrainedModelDeployment
	// Stops one or more data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-dfanalytics.html
	StopDataFrameAnalytics ml_stop_data_frame_analytics.NewStopDataFrameAnalytics
	// Stops one or more datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-stop-datafeed.html
	StopDatafeed ml_stop_datafeed.NewStopDatafeed
	// Stop a trained model deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-trained-model-deployment.html
	StopTrainedModelDeployment ml_stop_trained_model_deployment.NewStopTrainedModelDeployment
	// Updates certain properties of a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-dfanalytics.html
	UpdateDataFrameAnalytics ml_update_data_frame_analytics.NewUpdateDataFrameAnalytics
	// Updates certain properties of a datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-datafeed.html
	UpdateDatafeed ml_update_datafeed.NewUpdateDatafeed
	// Updates the description of a filter, adds items, or removes items.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-filter.html
	UpdateFilter ml_update_filter.NewUpdateFilter
	// Updates certain properties of an anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-job.html
	UpdateJob ml_update_job.NewUpdateJob
	// Updates certain properties of a snapshot.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-snapshot.html
	UpdateModelSnapshot ml_update_model_snapshot.NewUpdateModelSnapshot
	// Upgrades a given job snapshot to the current major version.
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
	// Removes the archived repositories metering information present in the
	// cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-repositories-metering-archive-api.html
	ClearRepositoriesMeteringArchive nodes_clear_repositories_metering_archive.NewClearRepositoriesMeteringArchive
	// Returns cluster repositories metering information.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-repositories-metering-api.html
	GetRepositoriesMeteringInfo nodes_get_repositories_metering_info.NewGetRepositoriesMeteringInfo
	// Returns information about hot threads on each node in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-hot-threads.html
	HotThreads nodes_hot_threads.NewHotThreads
	// Returns information about nodes in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-info.html
	Info nodes_info.NewInfo
	// Reloads secure settings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/secure-settings.html#reloadable-secure-settings
	ReloadSecureSettings nodes_reload_secure_settings.NewReloadSecureSettings
	// Returns statistical information about nodes in the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-stats.html
	Stats nodes_stats.NewStats
	// Returns low-level information about REST actions usage on nodes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-usage.html
	Usage nodes_usage.NewUsage
}

type QueryRuleset struct {
	// Deletes a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-ruleset.html
	Delete query_ruleset_delete.NewDelete
	// Returns the details about a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-ruleset.html
	Get query_ruleset_get.NewGet
	// Lists query rulesets.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-query-rulesets.html
	List query_ruleset_list.NewList
	// Creates or updates a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-ruleset.html
	Put query_ruleset_put.NewPut
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
	// Returns the rollup capabilities of all jobs inside of a rollup index (e.g.
	// the index where rollup data is stored).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-rollup-index-caps.html
	GetRollupIndexCaps rollup_get_rollup_index_caps.NewGetRollupIndexCaps
	// Creates a rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-put-job.html
	PutJob rollup_put_job.NewPutJob
	// Enables searching rolled-up data using the standard query DSL.
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
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-search-application.html
	Delete search_application_delete.NewDelete
	// Delete a behavioral analytics collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-analytics-collection.html
	DeleteBehavioralAnalytics search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalytics
	// Returns the details about a search application.
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
	// Perform a search against a search application
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
	// Creates or updates the user profile on behalf of another user.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-activate-user-profile.html
	ActivateUserProfile security_activate_user_profile.NewActivateUserProfile
	// Enables authentication as a user and retrieve information about the
	// authenticated user.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-authenticate.html
	Authenticate security_authenticate.NewAuthenticate
	// Updates the attributes of multiple existing API keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-update-api-keys.html
	BulkUpdateApiKeys security_bulk_update_api_keys.NewBulkUpdateApiKeys
	// Changes the passwords of users in the native realm and built-in users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-change-password.html
	ChangePassword security_change_password.NewChangePassword
	// Clear a subset or all entries from the API key cache.
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
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
	CreateApiKey security_create_api_key.NewCreateApiKey
	// Creates a cross-cluster API key for API key based remote cluster access.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-cross-cluster-api-key.html
	CreateCrossClusterApiKey security_create_cross_cluster_api_key.NewCreateCrossClusterApiKey
	// Creates a service account token for access without requiring basic
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
	// Allows a kibana instance to configure itself to communicate with a secured
	// elasticsearch cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-kibana-enrollment.html
	EnrollKibana security_enroll_kibana.NewEnrollKibana
	// Allows a new node to enroll to an existing cluster with security enabled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-node-enrollment.html
	EnrollNode security_enroll_node.NewEnrollNode
	// Retrieves information for one or more API keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html
	GetApiKey security_get_api_key.NewGetApiKey
	// Retrieves the list of cluster privileges and index privileges that are
	// available in this version of Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-builtin-privileges.html
	GetBuiltinPrivileges security_get_builtin_privileges.NewGetBuiltinPrivileges
	// Retrieves application privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-privileges.html
	GetPrivileges security_get_privileges.NewGetPrivileges
	// Retrieves roles in the native realm.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role.html
	GetRole security_get_role.NewGetRole
	// Retrieves role mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-role-mapping.html
	GetRoleMapping security_get_role_mapping.NewGetRoleMapping
	// Retrieves information about service accounts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-service-accounts.html
	GetServiceAccounts security_get_service_accounts.NewGetServiceAccounts
	// Retrieves information of all service credentials for a service account.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-service-credentials.html
	GetServiceCredentials security_get_service_credentials.NewGetServiceCredentials
	// Creates a bearer token for access without requiring basic authentication.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-token.html
	GetToken security_get_token.NewGetToken
	// Retrieves information about users in the native realm and built-in users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user.html
	GetUser security_get_user.NewGetUser
	// Retrieves security privileges for the logged in user.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-privileges.html
	GetUserPrivileges security_get_user_privileges.NewGetUserPrivileges
	// Retrieves user profiles for the given unique ID(s).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-profile.html
	GetUserProfile security_get_user_profile.NewGetUserProfile
	// Creates an API key on behalf of another user.
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
	// Adds and updates roles in the native realm.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role.html
	PutRole security_put_role.NewPutRole
	// Creates and updates role mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role-mapping.html
	PutRoleMapping security_put_role_mapping.NewPutRoleMapping
	// Adds and updates users in the native realm. These users are commonly referred
	// to as native users.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html
	PutUser security_put_user.NewPutUser
	// Retrieves information for API keys using a subset of query DSL
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-api-key.html
	QueryApiKeys security_query_api_keys.NewQueryApiKeys
	// Exchanges a SAML Response message for an Elasticsearch access token and
	// refresh token pair
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-authenticate.html
	SamlAuthenticate security_saml_authenticate.NewSamlAuthenticate
	// Verifies the logout response sent from the SAML IdP
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-complete-logout.html
	SamlCompleteLogout security_saml_complete_logout.NewSamlCompleteLogout
	// Consumes a SAML LogoutRequest
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-invalidate.html
	SamlInvalidate security_saml_invalidate.NewSamlInvalidate
	// Invalidates an access token and a refresh token that were generated via the
	// SAML Authenticate API
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-logout.html
	SamlLogout security_saml_logout.NewSamlLogout
	// Creates a SAML authentication request
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-prepare-authentication.html
	SamlPrepareAuthentication security_saml_prepare_authentication.NewSamlPrepareAuthentication
	// Generates SAML metadata for the Elastic stack SAML 2.0 Service Provider
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-sp-metadata.html
	SamlServiceProviderMetadata security_saml_service_provider_metadata.NewSamlServiceProviderMetadata
	// Get suggestions for user profiles that match specified search criteria.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-suggest-user-profile.html
	SuggestUserProfiles security_suggest_user_profiles.NewSuggestUserProfiles
	// Updates attributes of an existing API key.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-api-key.html
	UpdateApiKey security_update_api_key.NewUpdateApiKey
	// Update application specific data for the user profile of the given unique ID.
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
	// Removes stale data from repository.
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
	// Creates or updates a synonyms set
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
	// Returns a list of tasks.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	List tasks_list.NewList
}

type TextStructure struct {
	// Finds the structure of a text file. The text file must contain data that is
	// suitable to be ingested into Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-structure.html
	FindStructure text_structure_find_structure.NewFindStructure
}

type Transform struct {
	// Deletes an existing transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-transform.html
	DeleteTransform transform_delete_transform.NewDeleteTransform
	// Retrieves configuration information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform.html
	GetTransform transform_get_transform.NewGetTransform
	// Retrieves usage information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-stats.html
	GetTransformStats transform_get_transform_stats.NewGetTransformStats
	// Previews a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-transform.html
	PreviewTransform transform_preview_transform.NewPreviewTransform
	// Instantiates a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-transform.html
	PutTransform transform_put_transform.NewPutTransform
	// Resets an existing transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-transform.html
	ResetTransform transform_reset_transform.NewResetTransform
	// Schedules now a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/schedule-now-transform.html
	ScheduleNowTransform transform_schedule_now_transform.NewScheduleNowTransform
	// Starts one or more transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-transform.html
	StartTransform transform_start_transform.NewStartTransform
	// Stops one or more transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-transform.html
	StopTransform transform_stop_transform.NewStopTransform
	// Updates certain properties of a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-transform.html
	UpdateTransform transform_update_transform.NewUpdateTransform
	// Upgrades all transforms.
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
	// Forces the execution of a stored watch.
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
	// Retrieves information about the installed X-Pack features.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html
	Info xpack_info.NewInfo
	// Retrieves usage information about the installed X-Pack features.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/usage-api.html
	Usage xpack_usage.NewUsage
}

type API struct {
	AsyncSearch         AsyncSearch
	Autoscaling         Autoscaling
	Cat                 Cat
	Ccr                 Ccr
	Cluster             Cluster
	Core                Core
	DanglingIndices     DanglingIndices
	Enrich              Enrich
	Eql                 Eql
	Features            Features
	Fleet               Fleet
	Graph               Graph
	Ilm                 Ilm
	Indices             Indices
	Ingest              Ingest
	License             License
	Logstash            Logstash
	Migration           Migration
	Ml                  Ml
	Monitoring          Monitoring
	Nodes               Nodes
	QueryRuleset        QueryRuleset
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

	// Allows to perform multiple index/update/delete operations in a single
	// request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Explicitly clears the search context for a scroll.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Close a point in time
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Creates a new document in the index.
	//
	// Returns a 409 response when a document with a same ID already exists in the
	// index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Removes a document from the index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Deletes documents matching the provided query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Changes the number of requests per second for a particular Delete By Query
	// operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	// Deletes a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Returns information about whether a document exists in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Returns information about whether a document source exists in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Returns information about why a specific matches (or doesn't match) a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// Returns the information about the capabilities of fields among multiple
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Returns a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Returns a script.
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
	// Creates or updates a document in an index.
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
	// Allows to execute several search template operations in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Returns multiple termvectors in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// Open a point in time that can be used in subsequent searches
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Creates or updates a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Allows to evaluate the quality of ranked search results over a set of typical
	// search queries
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Allows to copy documents from one index to another, optionally filtering the
	// source
	// documents by a query, changing the destination index settings, or fetching
	// the
	// documents from a remote cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Changes the number of requests per second for a particular Reindex operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle
	// Allows to use the Mustache language to pre-render a search definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Allows an arbitrary script to be executed and a result to be returned
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Allows to retrieve a large numbers of results from a single search request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Returns results matching a query.
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
	// Allows to use the Mustache language to pre-render a search definition.
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
	// Performs an update on every document in the index without changing the
	// source,
	// for example to pick up a mapping change.
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

		// Features
		Features: Features{
			GetFeatures:   features_get_features.NewGetFeaturesFunc(tp),
			ResetFeatures: features_reset_features.NewResetFeaturesFunc(tp),
		},

		// Fleet
		Fleet: Fleet{
			GlobalCheckpoints: fleet_global_checkpoints.NewGlobalCheckpointsFunc(tp),
			Msearch:           fleet_msearch.NewMsearchFunc(tp),
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

		// QueryRuleset
		QueryRuleset: QueryRuleset{
			Delete: query_ruleset_delete.NewDeleteFunc(tp),
			Get:    query_ruleset_get.NewGetFunc(tp),
			List:   query_ruleset_list.NewListFunc(tp),
			Put:    query_ruleset_put.NewPutFunc(tp),
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
			SamlAuthenticate:            security_saml_authenticate.NewSamlAuthenticateFunc(tp),
			SamlCompleteLogout:          security_saml_complete_logout.NewSamlCompleteLogoutFunc(tp),
			SamlInvalidate:              security_saml_invalidate.NewSamlInvalidateFunc(tp),
			SamlLogout:                  security_saml_logout.NewSamlLogoutFunc(tp),
			SamlPrepareAuthentication:   security_saml_prepare_authentication.NewSamlPrepareAuthenticationFunc(tp),
			SamlServiceProviderMetadata: security_saml_service_provider_metadata.NewSamlServiceProviderMetadataFunc(tp),
			SuggestUserProfiles:         security_suggest_user_profiles.NewSuggestUserProfilesFunc(tp),
			UpdateApiKey:                security_update_api_key.NewUpdateApiKeyFunc(tp),
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
			FindStructure: text_structure_find_structure.NewFindStructureFunc(tp),
		},

		// Transform
		Transform: Transform{
			DeleteTransform:      transform_delete_transform.NewDeleteTransformFunc(tp),
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

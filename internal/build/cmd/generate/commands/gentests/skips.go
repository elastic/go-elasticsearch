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

package gentests

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

var skipTests map[string][]string

func init() {
	err := yaml.NewDecoder(strings.NewReader(skipTestsYAML)).Decode(&skipTests)
	if err != nil {
		panic(fmt.Sprintf("ERROR: %v", err))
	}
}

var skipFiles = []string{
	"update/85_fields_meta.yml",            // Uses non-existing API property
	"update/86_fields_meta_with_types.yml", // --||--

	"ml/jobs_get_result_buckets.yml",    // Passes string value to int variable
	"ml/jobs_get_result_categories.yml", // --||--
	"ml/set_upgrade_mode.yml",           // --||--

	"ml/evaluate_data_frame.yml", // Floats as map keys

	"search/320_disallow_queries.yml", // Tries to match key in an empty map (`transient:{}`)

	"watcher/stats/10_basic.yml", // Sets "emit_stacktraces" as string ("true"), not bool

	"search.highlight/20_fvh.yml", // bad backslash

	"indices.stats/50_disk_usage.yml",  // Needs a replacement mechanism implementation
	"indices.stats/60_field_usage.yml", // Needs a replacement mechanism implementation
	"eql/10_basic.yml",
	"field_caps/50_fieldtype_filter.yml", // Incompatible test, need handling for double escaping keys with dots
	"aggregations/variable_width_histogram.yml",
	"cluster.desired_nodes/10_basic.yml",  // incompatible $ stash replacement
	"api_key/12_grant.yml",                // incompatible $ stash replacement, need bearer token integration
	"aggregations/percentiles_bucket.yml", // incompatible maps
	"user_profile/10_basic.yml",
	"indices.stats/100_search_idle.yml", // incompatible maps of array
	"ml/3rd_party_deployment.yml",       // incompatible ml tests
	"dlm/10_usage.yml",                  // incompatible float expansion
	"api_key/60_admin_user.yml",
}

// TODO: Comments into descriptions for `Skip()`
var skipTestsYAML = `
---
# Cannot distinguish between missing value for refresh and an empty string
bulk/50_refresh.yml:
  - refresh=empty string immediately makes changes are visible in search
bulk/51_refresh_with_types.yml:
  - refresh=empty string immediately makes changes are visible in search
create/60_refresh.yml:
  - When refresh url parameter is an empty string that means "refresh immediately"
create/61_refresh_with_types.yml:
  - When refresh url parameter is an empty string that means "refresh immediately"
delete/50_refresh.yml:
  - When refresh url parameter is an empty string that means "refresh immediately"
delete/51_refresh_with_types.yml:
  - When refresh url parameter is an empty string that means "refresh immediately"
index/60_refresh.yml:
  - When refresh url parameter is an empty string that means "refresh immediately"
index/61_refresh_with_types.yml:
  - When refresh url parameter is an empty string that means "refresh immediately"
update/60_refresh.yml:
  - When refresh url parameter is an empty string that means "refresh immediately"
update/61_refresh_with_types.yml:
  - When refresh url parameter is an empty string that means "refresh immediately"

# catch: bad_request, Expected [status] to not be nil
indices.data_stream/10_basic.yml:
  - Create data stream with invalid name

# Stash in value
cluster.reroute/11_explain.yml:
nodes.info/30_settings.yml:
nodes.stats/20_response_filtering.yml:
nodes.stats/30_discovery.yml:
  - Discovery stats
nodes.discovery/30_discovery.yml:
  - Discovery stats

# wrongly assumed as []interface{} where these should be map[string]interface{}
data_stream/150_tsdb.yml:
eql/20_runtime_mappings.yml:
  - Execute EQL events query with search time keyword runtime field
  - Execute EQL events query with search time ip runtime field
spatial/60_geo_line.yml:
  - Test geo_line aggregation on geo points
  - Test empty buckets
spatial/100_geo_grid_ingest.yml:
  - Test geo_grid on geotile using default target format

# Arbitrary key
indices.shrink/10_basic.yml:
indices.shrink/20_source_mapping.yml:
indices.shrink/30_copy_settings.yml:
indices.split/30_copy_settings.yml:
nodes.info/10_basic.yml:
nodes.info/40_aggs.yml:
nodes.reload_secure_settings/10_basic.yml:
nodes.stats/50_indexing_pressure.yml:
nodes.stats/40_store_stats.yml:
nodes.stats/60_transport_stats.yml:

# Parsed response is YAML: value is map[interface {}]interface {}, not map[string]interface {}
cat.aliases/20_headers.yml:
  - Simple alias with yaml body through Accept header

# Incorrect int instead of float in match (aggregations.date_range.buckets.0.from: 1000000); TODO: PR
aggregations/range.yml:
  - Date range
  - Min and max long range bounds

# Mismatch in number parsing, 8623000 != 8.623e+06
aggregations/geo_distance.yml:
  - avg_bucket

# .key in map issue
aggregations/top_hits.yml:
  - explain

# No support for headers per request yet
tasks.list/10_basic.yml:
  - tasks_list headers

# Node Selector feature not implemented
cat.aliases/10_basic.yml:
  - "Help (pre 7.4.0)"
  - "Simple alias (pre 7.4.0)"
  - "Complex alias (pre 7.4.0)"
  - "Column headers (pre 7.4.0)"
  - "Alias against closed index (pre 7.4.0)"

# Checks for nil required arguments makes this test incompatible with the integration tests
indices.delete_alias/all_path_options.yml:
  - check delete with blank index and blank alias
indices.put_alias/all_path_options.yml:
  - put alias with blank index
  - put alias with missing name

indices.put_mapping/10_basic.yml:
  - "Put mappings with explicit _doc type bwc"

# Test fails with: [400 Bad Request] illegal_argument_exception, "template [test] has index patterns [test-*] matching patterns from existing index templates [test2,test] with patterns (test2 => [test-*],test => [test-*, test2-*]), use index templates (/_index_template) instead"
test/indices.put_template/10_basic.yml:

# Incompatible regex
cat.templates/10_basic.yml:
  - "Sort templates"
  - "Multiple template"
ml/trained_model_cat_apis.yml:
  - Test cat trained models

# Missing test setup
cluster.voting_config_exclusions/10_basic.yml:
  - "Add voting config exclusion by unknown node name"
indices.resolve_index/10_basic_resolve_index.yml:
  - "Resolve index with hidden and closed indices"

# Not relevant
search/issue4895.yml:
search/issue9606.yml:

# FIXME
bulk/80_cas.yml:
bulk/81_cas_with_types.yml:

# Incompatible variable replacement
tsdb/150_tsdb.yml:
  - 

# Incompatible to date with test runner
tsdb/80_index_resize.yml:
  - shrink
  - clone

# Number conversion needs to be addressed in test gen
tsdb/40_search.yml:
  - aggregate a metric

# Deliberate wrong type doesn't match Go types
cluster.desired_nodes/10_basic.yml:
  - Test version must be a number
  - Test update move to a new history id
  - Test delete desired nodes
  - Test using the same version with different definition is forbidden
  - Test unknown settings are allowed in future versions
  - Test some settings can be overridden
  - Test history_id must be present
  - Test update desired nodes is idempotent
  - Test going backwards within the same history is forbidden

# ----- X-Pack ----------------------------------------------------------------

# Floats "3.0" decoded as int "3" by gopkg.in/yaml.v2
analytics/top_metrics.yml:
runtime_fields/30_double.yml:
  - docvalue_fields
  - fetch fields
search.vectors/60_dense_vector_dynamic_mapping.yml:
  - Fields with float arrays below the threshold still map as float

# Stash in body
api_key/10_basic.yml:
  - Test invalidate api keys
api_key/11_invalidation.yml:
  - Test invalidate api key by username
rollup/put_job.yml:
  - Test put job with templates

# Changing password locks out tests
change_password/10_basic.yml:
  - Test user changing their own password
  - Test changing users password with prehashed password

change_password/12_custom_hash.yml:
  - Test changing users password with pre-hashed password

# Missing refreshes in the test
data_frame/transforms_start_stop.yml:
ml/index_layout.yml:
transform/transforms_start_stop.yml:
  - Verify start transform reuses destination index
transform/transforms_start_stop.yml:
  - Test get multiple transform stats
transform/transforms_stats.yml:
  - Test get multiple transform stats
  - Test get multiple transform stats where one does not have a task
transform/transforms_unattended.yml:
  - Test unattended put and start wildcard

# More QA tests than API tests
data_frame/transforms_stats.yml:
  - Test get multiple transform stats
  - Test get transform stats on missing transform
  - Test get multiple transform stats where one does not have a task
ml/jobs_crud.yml:
  - Test reopen job resets the finished time

# Invalid license makes subsequent tests fail
license/20_put_license.yml:

# Test tries to match on map from body, but Go keys are not sorted
ml/jobs_crud.yml:
  - Test job with rules
  - Test put job with model_memory_limit as number
  - Test put job with model_memory_limit as string and lazy open
  - Test update job

# Test gets stuck every time
ml/jobs_get_stats.yml:

# status_exception, Cannot process data because job [post-data-job] does not have a corresponding autodetect process
# resource_already_exists_exception, task with id {job-post-data-job} already exist
# status_exception, Cannot open job [start-stop-datafeed-job-foo-1] because it has already been opened
ml/post_data.yml:
  - Test flush with skip_time
  - Test POST data job api, flush, close and verify DataCounts doc
  - Test flush and close job WITHOUT sending any data
ml/start_stop_datafeed.yml:
  - Test stop given expression
transform/transforms_start_stop.yml:
  - Test start transform
  - Verify start transform reuses destination index

# Possible bad test setup, Cannot open job [start-stop-datafeed-job] because it has already been opened
# resource_already_exists_exception, task with id {job-start-stop-datafeed-job-foo-2} already exist
ml/start_stop_datafeed.yml:
  - Test start datafeed when persistent task allocation disabled

# Indexing step doesn't appear to work (getting total.hits=0)
monitoring/bulk/10_basic.yml:
  - Bulk indexing of monitoring data on closed indices should throw an export exception
# Indexing step doesn't appear to work (getting total.hits=0)
monitoring/bulk/20_privileges.yml:
  - Monitoring Bulk API

# Test tries to match on whole body, but map keys are unstable in Go
rollup/security_tests.yml:

# Test tries to match on map key, but map keys are unstable in Go
ml/data_frame_analytics_crud.yml:
  - Test put with description
  - Test put valid config with custom outlier detection

# Unsupported feature: allowed_warnings
ml/data_frame_analytics_crud.yml:
  - Test put classification given deprecated maximum_number_trees

# This test suite keeps failing too often: disable it altogether
ml/data_frame_analytics_crud.yml:

# 404s, panics, ... possible bad setup/teardown
ml/delete_model_snapshot.yml:
ml/get_datafeed_stats.yml:
ml/get_model_snapshots.yml:
eql/30_async_missing_events.yml:

# resource_already_exists_exception for model, need improved teardown for models
ml/semantic_search.yml:

# model is not deployed to any node
ml/text_expansion_search.yml:
ml/text_expansion_search_sparse_vector.yml:

# TEMPORARY: Missing 'body: { indices: "test_index" }' payload, TODO: PR
snapshot/10_basic.yml:
  - Create a source only snapshot and then restore it

# illegal_argument_exception: Provided password hash uses [NOOP] but the configured hashing algorithm is [BCRYPT]
users/10_basic.yml:
  - Test put user with password hash

# Slash in index name is not escaped (BUG)
security/authz/13_index_datemath.yml:
  - Test indexing documents with datemath, when permitted

# Possibly a cluster health color mismatch...
security/authz/14_cat_indices.yml:

# Test looks for "testnode.crt", but "ca.crt" is returned first
ssl/10_basic.yml:
  - Test get SSL certificates

# class org.elasticsearch.xpack.vectors.query.VectorScriptDocValues$DenseVectorScriptDocValues cannot be cast to class org.elasticsearch.xpack.vectors.query.VectorScriptDocValues$SparseVectorScriptDocValues ...
vectors/30_sparse_vector_basic.yml:
  - Dot Product
# java.lang.IllegalArgumentException: No field found for [my_dense_vector] in mapping
vectors/40_sparse_vector_special_cases.yml:
  - Vectors of different dimensions and data types
  - Dimensions can be sorted differently
  - Distance functions for documents missing vector field should return 0

# Cannot connect to Docker IP
watcher/execute_watch/60_http_input.yml:

# Test tries to match on "tagline", which requires "human=false", which doesn't work in the Go API.
# Also test does too much within a single test, so has to be disabled as whole, unfortunately.
xpack/15_basic.yml:

# Test uses "y" as a property name, which is parsed as 'true' in the Go YAML library;
# see https://yaml.org/type/bool.html
ml/explain_data_frame_analytics.yml:
  - Test empty data frame given body
  - Test non-empty data frame given body
runtime_fields/10_keyword.yml:
  - docvalue_fields
  - fetch fields
vector-tile/10_basic.yml:
vector-tile/20_aggregations.yml:

# Test uses a char as a property name, which is parsed as 'false' in the Go YAML library;
aggregations/derivative.yml:
  - in histogram
  - partially mapped
aggregations/histogram.yml:
  - histogram profiler
aggregations/moving_fn.yml:
  - in histogram

# Getting "no matching index template found for data stream [invalid-data-stream]"
data_stream/10_basic.yml:
  - Create data stream
  - Create data stream with invalid name
  - append-only writes to backing indices prohibited

# The matcher like 'indices.0.aliases.0' points to internal index
data_stream/80_resolve_index_data_streams.yml:
  - Resolve index with indices, aliases, and data streams
  - Resolve index with hidden and closed indices

# Zero matchers like '...shards.0.stores.0.allocation:primary' expect array, not map
data_stream/40_supported_apis.yml:
  - Verify shard stores api

# Failing with error 'Index [.security] is not on the current version. Security features relying on the index will not be available until the upgrade API is run on the index'
data_stream/40_supported_apis.yml:
  - Verify search shards api
  - Verify shard stores api
  - Verify rank eval with data streams
  - Verify get field mappings api
  - Open write index for data stream opens all backing indices
data_stream/90_reindex.yml:
  - Reindex from data stream into an index

# Needs further implementation of .key access to map variables
data_streams/10_data_stream_resolvability.yml:
  - Verify data stream resolvability in ILM remove policy API

# Error: constant 9223372036854775808 overflows int (https://play.golang.org/p/7pUdz-_Pdom)
unsigned_long/10_basic.yml:
unsigned_long/20_null_value.yml:
unsigned_long/30_multi_fields.yml:
unsigned_long/40_different_numeric.yml:
unsigned_long/50_script_values.yml:
unsigned_long/60_collapse.yml:

# Cannot compare float64 ending with .0 reliably due to inconsistent serialisation (https://github.com/golang/go/issues/26363)
search/330_fetch_fields.yml:
  - Test nested field inside object structure

search/350_point_in_time.yml:
  - msearch

nodes.stats/11_indices_metrics.yml:
  - Metric - http
  - Metric - blank for indices shards
  - Metric - _all for indices shards
  - indices shards total count test

data_stream/10_data_stream_resolvability.yml:
  - Verify data stream resolvability in ILM remove policy API

data_stream/60_get_backing_indices.yml:
  - Get backing indices for data stream

searchable_snapshots/10_usage.yml:
  - Tests searchable snapshots usage stats with full_copy and shared_cache indices

# Expects count 2 but returns only 1
service_accounts/10_basic.yml:
  - Test service account tokens

# Replace stash token in payload not yet implemented
api_key/20_query.yml:
  - Test query api key
api_key/30_update.yml:
  - Test bulk update api keys
  - Test bulk update api key without explicit field updates
  - Test bulk update api key with empty request fields
api_key/40_view_role_descriptors.yml:
  - Test API key role descriptors in Get and Query responses
api_key/50_cross_cluster.yml:
  - Test create a cross-cluster API key
authenticate/11_admin_user.yml:
  - Test authenticate with token

token/10_basic.yml:
  - Test invalidate user's tokens
  - Test invalidate realm's tokens

# Replacement with pattern matching fails
user_profile/40_has_privileges.yml:
  - Test profile has privileges api

# Bad type matching
aggregate-metrics/100_synthetic_source.yml:
  - constant_keyword
analytics/histogram.yml:
  - histogram with synthetic source

# incompatible storage
searchable_snapshots/20_synthetic_source.yml:
  - Tests searchable snapshots usage stats

# incompatible float format
aggregations/max_metric.yml:
  - Merging results with unmapped fields

# unsupported file upload
get/100_synthetic_source.yml:
  - indexed dense vectors
  - non-indexed dense vectors

indices.stats/70_write_load.yml:
  - Write load average is tracked at shard level

search/400_synthetic_source.yml:
  - stored keyword without sibling fields

health/10_usage.yml:
  - Usage stats on the health API
`

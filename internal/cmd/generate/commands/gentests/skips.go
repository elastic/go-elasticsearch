// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package gentests

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
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

	"watcher/stats/10_basic.yml", // Sets "emit_stacktraces" as string ("true"), not bool
}

// TODO: Comments into descriptions for `Skip()`
//
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

# Stash in value
cluster.reroute/11_explain.yml:
nodes.info/30_settings.yml:
nodes.stats/20_response_filtering.yml:
nodes.stats/30_discovery.yml:
  - Discovery stats
nodes.discovery/30_discovery.yml:
  - Discovery stats

# Arbitrary key
indices.shrink/10_basic.yml:
indices.shrink/20_source_mapping.yml:
indices.shrink/30_copy_settings.yml:
indices.split/30_copy_settings.yml:

# Parsed response is YAML: value is map[interface {}]interface {}, not map[string]interface {}
cat.aliases/20_headers.yml:
  - Simple alias with yaml body through Accept header

# Incorrect int instead of float in match (aggregations.date_range.buckets.0.from: 1000000); TODO: PR
search.aggregation/40_range.yml:
  - Date range

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

indices.put_mapping/10_basic.yml:
  - "Put mappings with explicit _doc type bwc"

# Not relevant
search/issue4895.yml:
search/issue9606.yml:

# FIXME
bulk/80_cas.yml:
bulk/81_cas_with_types.yml:

# ----- X-Pack ----------------------------------------------------------------

# Stash in body
api_key/10_basic.yml:
  - Test invalidate api key
api_key/11_invalidation.yml:
  - Test invalidate api key by username
rollup/put_job.yml:
  - Test put job with templates

# Changing password locks out tests
change_password/10_basic.yml:
  - Test user changing their own password

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
`

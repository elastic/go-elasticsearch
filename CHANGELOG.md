# Changelog

## [8.19.1](https://github.com/elastic/go-elasticsearch/compare/v8.19.0...v8.19.1) (2025-12-12)


### Features

* Add Close method to BaseClient ([#1056](https://github.com/elastic/go-elasticsearch/issues/1056)) ([#1076](https://github.com/elastic/go-elasticsearch/issues/1076)) ([e73cb92](https://github.com/elastic/go-elasticsearch/commit/e73cb926cff6fdd5f1be3bcc6d2f8889587dd9bb))
* Add queue size multiplier config to BulkIndexer ([#1029](https://github.com/elastic/go-elasticsearch/issues/1029)) ([#1055](https://github.com/elastic/go-elasticsearch/issues/1055)) ([cdaf2aa](https://github.com/elastic/go-elasticsearch/commit/cdaf2aa28586ea938230b165487f9489e2039944))
* Add support for interceptors in Elasticsearch client ([#1080](https://github.com/elastic/go-elasticsearch/issues/1080)) ([#1086](https://github.com/elastic/go-elasticsearch/issues/1086)) ([f4bda5f](https://github.com/elastic/go-elasticsearch/commit/f4bda5ff2682c18fee2dc66e4abc85e56b5222a5))


### Bug Fixes

* Notify items if an error occurs in bulk indexer ([#615](https://github.com/elastic/go-elasticsearch/issues/615)) ([#1057](https://github.com/elastic/go-elasticsearch/issues/1057)) ([00b0ac1](https://github.com/elastic/go-elasticsearch/commit/00b0ac19af728ea10bd61571c3dc4de71e6ce130))

## 8.19.0

## API

* Updated APIs to 8.19.0

## Typed API

* Update TypedAPI to latest [elasticsearch-specification 8.19](https://github.com/elastic/elasticsearch-specification/commit/470b4b9)

## 8.18.1

 * This patch release fixes the broken build found in 8.18.0. If you are using the `TypedClient`, you should update to this version.

## 8.18.0

* Update `elastictransport` to `8.7.0`.
* Thanks to @zaneli, the `TypedClient` can now be used in the `BulkIndexer`.

## New

* This release adds a `BaseClient` constructor with no attached APIs, allowing it to be used purely as a transport layer instead of a full-featured API client.

```go
baseClient, err := elasticsearch.NewBaseClient(elasticsearch.Config{
    Addresses: []string{
        "http://localhost:9200",
    },
})

if err != nil {
    log.Println(err)
    return
}

res, err := esapi.InfoRequest{
    Pretty:     false,
    Human:      false,
    ErrorTrace: false,
    FilterPath: nil,
    Header:     nil,
    Instrument: baseClient.InstrumentationEnabled(),
}.Do(context.Background(), baseClient)

if err != nil {
    log.Println(err)
    return
}
defer res.Body.Close()
if res.IsError() {
    log.Println("Error response:", res)
    return
}
var infoMap map[string]interface{}
if err := json.NewDecoder(res.Body).Decode(&infoMap); err != nil {
    log.Println("Error parsing response:", err)
    return
}
log.Printf("Elasticsearch version esapi: %s\n", infoMap["version"].(map[string]interface{})["number"])

typedRes, err := info.New(baseClient).Do(context.Background())
if err != nil {
    log.Println(err)
    return
}
log.Printf("Elasticsearch version typedapi: %s\n", typedRes.Version.Int)
```

## API

* Updated APIs to 8.18.0

## Typed API

* Update APIs to 8.18 ([cbfcc73](https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5))

## 8.17.1

* Update elastictransport to 8.6.1

Thanks to @AkisAya and @jmfrees for their contributions!

## 8.17.0

* Expose BulkIndexer total flushed bytes metric [#914](https://github.com/elastic/go-elasticsearch/pull/914) thanks to @aureleoules

## API

Updated APIs to 8.17.0

## Typed API

Update APIs to latest [elasticsearch-specification 8.17](https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64)

## 8.16.0

## API

* `InferenceStreamInference` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/stream-inference-api.html)
* `QueryRulesTest` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/test-query-ruleset.html)
* `Ingest.DeleteIPLocationDatabase` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-ip-location-database-api.html)
* `Ingest.GetIPLocationDatabase` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ip-location-database-api.html)
* `Ingest.PutIPLocationDatabase` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/put-ip-location-database-api.html)

## Typed API

Update APIs to latest [elasticsearch-specification 8.16](https://github.com/elastic/elasticsearch-specification/tree/4fcf747dfafc951e1dcf3077327e3dcee9107db3)

## 8.15.0

## API

* API is generated from the Elasticsearch 8.15.0 specification. 

## Typed API

Update APIs to latest [elasticsearch-specification 8.15](https://github.com/elastic/elasticsearch-specification/tree/19027dbdd366978ccae41842a040a636730e7c10)

## 8.14.0

## API

New APIs:

 * ConnectorUpdateActiveFiltering [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/update-connector-filtering-api.html)
 * ConnectorUpdateFilteringValidation [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/update-connector-filtering-api.html)
 * TextStructureFindFieldStructure [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/find-field-structure.html)
 * TextStructureFindMessageStructure [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/find-message-structure.html)

## Typed API

New APIs:

 * UpdateTrainedModelDeployment [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/update-trained-model-deployment.html)

## Transport
 * Fixed a deadlock in the connection pool https://github.com/elastic/elastic-transport-go/issues/20

## 8.13.1

## Typed API

Update APIs to latest [elasticsearch-specification 8.13](https://github.com/elastic/elasticsearch-specification/tree/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757)

## Fixes

This patch release brings a fix to the initialisation of the `Request` in endpoints which would prevent using the shortcuts for fields. 
Canonical`.Request()` method was unaffected.

* `Autoscaling.PutAutoscalingPolicy`
* `Indices.Downsample`
* `Indices.PutSettings`
* `Indices.SimulateTemplate`
* `Inference.PutModel`
* `Logstash.PutPipeline`
* `Ml.ValidateDetector`
* `SearchApplication.Put`

## 8.13.0

## API 

New APIS:

* `ConnectorSecretGet`
* `ConnectorSecretPost`
* `ConnectorSecretPut`
* `ConnectorSecretDelete`
* `ConnectorUpdateIndexName`
* `ConnectorUpdateNative`
* `ConnectorUpdateStatus`
* `ConnectorUpdateAPIKeyDocumentID`
* `ConnectorUpdateServiceDocumentType`


* `EsqlAsyncQuery` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/esql-async-query-api.html)
* `EsqlAsyncQueryGet` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/esql-async-query-get-api.html)
* `ProfilingFlamegraph` [Documentation](https://www.elastic.co/guide/en/observability/current/universal-profiling.html)
* `ProfilingStacktraces` [Documentation](https://www.elastic.co/guide/en/observability/current/universal-profiling.html)
* `TextStructureTestGrokPattern` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/test-grok-pattern.html)
* `Indices.ResolveCluster` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-resolve-cluster-api.html)
* `Security.QueryUser` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-user.html)

## Typed API

* `indices.ResolveCluster` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-resolve-cluster-api.html)
* `textstructure.TestGrokPattern` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/test-grok-pattern.html)

Thanks to @pakio, transport now has an optional pool based compression option. [link](https://github.com/elastic/elastic-transport-go/pull/19)

And to @tblyler for fixing a very subtle memory leak in the `BulkIndexer`. #797

## 8.12.1

* Fix: ticker memory leak in bulk indexer due to internal flush call resetting the ticker. #797
* Fix: Scroll now uses the body to pass the scroll_id. #785
* Add: generated UnmarshalJSON for Requests to allow injecting payloads using aliases.
* Fix: `put_synonym_rule` was not working due to a type issue in the [Elasticsearch API Specification](https://github.com/elastic/elasticsearch-specification/pull/2407).

## 8.12.0

## Client

### Golang version

The client now requires Golang version 1.20

### OpenTelemetry

The client now provides OpenTelemetry integration. This integration can be enabled in the config using the `elasticsearch.NewOpenTelemetryInstrumentation`.
Once set up, the provided `context` will be used to record spans with useful information about the request being made to the server.

More about what you can expect in the [Semantic Conventions for Elasticsearch](https://opentelemetry.io/docs/specs/semconv/database/elasticsearch/).

### BulkIndexer

`if_seq_no` & `if_primary_term` are now supported thanks to @benjyiw [#783](https://github.com/elastic/go-elasticsearch/pull/783)

## API

* `SimulateIngest`
* `ConnectorCheckIn`
* `ConnectorDelete`
* `ConnectorGet`
* `ConnectorLastSync`
* `ConnectorList`
* `ConnectorPost`
* `ConnectorPut`
* `ConnectorSyncJobCancel`
* `ConnectorSyncJobCheckIn`
* `ConnectorSyncJobDelete`
* `ConnectorSyncJobError`
* `ConnectorSyncJobGet`
* `ConnectorSyncJobList`
* `ConnectorSyncJobPost`
* `ConnectorSyncJobUpdateStats`
* `ConnectorUpdateConfiguration`
* `ConnectorUpdateError`
* `ConnectorUpdateFiltering`
* `ConnectorUpdateName`
* `ConnectorUpdatePipeline`
* `ConnectorUpdateScheduling`

## Typed API

* `Esql.Query` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-rest.html)
* `Fleet.PostSecret`
* `Inference` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/inference-apis.html)
  * `DeleteModel` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-inference-api.html)
  * `GetModel` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/get-inference-api.html)
  * `Inference` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html)
  * `PutModel` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/put-inference-api.html)
* `SearchApplication`
  * `GetSettings` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-settings.html)
  * `UpdateSettings` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-settings.html)

## 8.11.1

## Typed API

* Fix https://github.com/elastic/go-elasticsearch/issues/756 preventing from settings indices in `indices.PutSettings`

## 8.11.0

## API

**Experimental APIs**

* `EsqlQuery`            [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-query-api.html)
* `InferenceDeleteModel` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/delete-inference-api.html)
* `InferenceGetModel`    [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/get-inference-api.html)
* `InferenceInference`   [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/post-inference-api.html)
* `InferencePutModel`    [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/put-inference-api.html)

## Typed API

* Mandatory URL parameters are not exposed as functions anymore as they already exist in the constructor.

## New Compatibility Policy

Starting from version `8.12.0`, this library follow the Go language [policy](https://go.dev/doc/devel/release#policy). Each major Go release is supported until there are two newer major releases. For example, Go 1.5 was supported until the Go 1.7 release, and Go 1.6 was supported until the Go 1.8 release.

If you have any questions or concerns, please do not hesitate to reach out to us.

## 8.10.1

## Typed API

Update APIs to latest [elasticsearch-specification 8.10](https://github.com/elastic/elasticsearch-specification/commit/3b09f9d8e90178243f8a340a7bc324aab152c602)

## 8.10.0

## API
**Experimental APIs for internal use**
* `FleetDeleteSecret`
* `FleetGetSecret`
* `FleetPostSecret`

**Exprimental APIs**

`QueryRulesetList`

**Stable APIs**

`Security.GetSettings`
`Security.UpdateSettings`

## Typed API
**Exprimental APIs**

`QueryRuleset.List`

**Technical Preview**
* [QueryRuleSet](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-rules-apis.html)

**Beta**
* [Synonyms](https://www.elastic.co/guide/en/elasticsearch/reference/current/synonyms-apis.html)

## 8.9.0

## API
**New API**

* `Cluster.Info` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-info.html)

**Experimental APIs**

* `QueryRulesetGet` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/get-query-ruleset.html)
* `QueryRulesetDelete` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/delete-query-ruleset.html)
* `QueryRulesetPut` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/put-query-ruleset.html)
* `SearchApplicationRenderQuery` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/search-application-render-query.html)
* `Security.CreateCrossClusterAPIKey` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-cross-cluster-api-key.html)
* `Security.UpdateCrossClusterAPIKey` [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-cross-cluster-api-key.html)

## Typed API

* Propagated request fields towards the endpoint for ease of access, taking priority over same-name query string fields.
* Added a stub for Do methods on endpoints that only support a boolean response such as `core.exists`. 
* NDJSON endpoints support with custom serialization like `core.bulk`.
* Link to endpoints documentation in API index to better display and ease of use.

**fixes**

* Fixed a deserialization issue for `Property` & `Analyzer` #696

## 8.8.2

## Typed API

* Fixed deserialization for `Suggest` in search responses.
* Fixed double-quoted strings in deserialization for unions normalized as string. #684
* Fixed handling of `core.Get` response when the index did not exist. #678

## 8.7.0

## API

* `ML.DeleteJob`: Added `WithDeleteUserAnnotations`. Should annotations added by the user be deleted.
* `ML.ResetJob`: Added `WithDeleteUserAnnotations`. Should annotations added by the user be deleted.
* `ML.StartTrainedModelDeployment`: Added `WithPriority`. The deployment priority.
* `TransformGetTransformStats`: Added `WithTimeout`. Controls the time to wait for the stats.
* `TransformStartTransform`: Added `WithFrom`. Restricts the set of transformed entities to those changed after this time.

**New APIs**

`TransformScheduleNowTransform` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/8.7/schedule-now-transform.html).
`HealthReport` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/8.7/health-api.html).

## Typed API

* Inclusion of responses structures.

**Changes**

* `Do` method on endpoints now return a typed response, one per endpoint.
* `Perform` method added on endpoints, returns `http.Response` as did `Do`.
* Elasticsearch exceptions are now handled as `types.ElasticsearchError` with `.As` and `.Is` methods.
* `.Raw` now takes a reader as input.
* User defined values such as `_source` in `Hits` are now `json.RawMessage` to highlight they later deserializable nature.  

## 8.6.0

## API

* `ML.StartTrainedModelDeployment`: Added `WithPriority`

**New APIs**
* `ML.UpdateTrainedModelDeployment`: Updates certain properties of trained model deployment.


## Client
**BulkIndexer**

Improvements were made to the BulkIndexer memory usage to allow better handling under burst use cases. Thanks to @christos68k and @rockdaboot !

## 8.5.0

## API

* `ML.StartTrainedModelDeployment`: Description of `NumberOfAllocations` has been changed in "The total number of allocations this model is assigned across machine learning nodes".
* `Security.GetAPIKey`: Added `WithLimitedBy` boolean parameter. Flag to show the limited-by role descriptors of API Keys.
* `Security.GetUser`: Added `WithProfileUID` boolean parameter. Flag to retrieve profile uid (if exists) associated to the user.
* `Security.GetUserProfile`: Changed the description of uid parameter, a comma-separated list of unique identifier for user profiles.
* `Security.QueryAPIKeys`: Added `WithLimitedBy` boolean parameter. Flag to show the limited-by role descriptors of API Keys.
* `TextStructureFindStructure`: Added `EcsCompatibility` string parameter. Optional parameter to specify the compatibility mode with ECS Grok patterns - may be either 'v1' or 'disabled'.

**Promoted to stable**

* `ML.InferTrainedModel`
* `ML.PutTrainedModelDefinitionPart`
* `ML.PutTrainedModelVocabulary`
* `ML.StartTrainedModelDeployment`
* `ML.StopTrainedModelDeployment`
* `Security.activateUserProfile`
* `Security.DisableUserProfile`
* `Security.EnableUserProfile`
* `Security.GetUserProfile`
* `Security.HasPrivilegesUserProfile`
* `Security.SuggestUserProfiles`
* `Security.UpdateUserProfileData`

**New APIs**

* `ML.ClearTrainedModelDeploymentCache` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/clear-trained-model-deployment-cache.html).
* `Security.BulkUpdateAPIKeys` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-update-api-keys.html).
* `Indices.Downsample` (Experimental API) [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/xpack-rollup.html)

## Typed API

Following multiple feedbacks we decided to remove the builder API for the type tree.

In its place, work has started to further simplify the type tree by removing redundant type aliases. The API also now comes with a helper package named `some` that allows to call for inline pointers on primitive types.

In addition, a bug was fixed preventing the use of wildcards in index names, and enums are now extensible by default.

The Typed API remains in `alpha` stage while its development continues.

## 8.4.0

## API

* `get`, `mget` and `search` added `force_synthetic_source`: Should this request force synthetic _source? Use this to test if the mapping supports synthetic _source and to get a sense of the worst case performance. Fetches with this enabled will be slower the enabling synthetic source natively in the index.
* `ML.StartTrainedModelDeployment` added `cache_size`: A byte-size value for configuring the inference cache size. For example, 20mb.
* `Snapshot.Get` added `sort`, `size`, `order`, `from_sort_value`, `after`, `offset` and `slm_policy_filter`. More on these in the [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/get-snapshot-api.html).

**New API**

* `Security.UpdateAPIKey` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/security-api-update-api-key.html).

## Typed API

As highlighted in the release not for the 8.4.0-alpha.1, this release marks the beginning of the newly arrived `TypedClient`.

This new API is still in `alpha` stage and will be release alongside the existing `esapi`.

A few examples of standard use-cases can be found in the [TypedAPI section of the documentation](https://www.elastic.co/guide/en/elasticsearch/client/go-api/master/typedapi.html).

## 8.4.0-alpha.2

This second prerelease of the 8.4.0 updates the API for the client and fixes the serialization for types using [additional properties](https://github.com/elastic/elasticsearch-specification/blob/main/docs/behaviors.md#additionalproperties--additionalproperty).

## 8.4.0-alpha.1

This prerelease introduces a new typed API generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification). This generation from the common specification allows us to provide a complete API which uses an exhaustive hierarchy of types  reflecting the possibilities given by Elasticsearch.

This new API is the next iteration of the Go client for Elasticsearch, it now lives alongside the existing API, it is in `alpha` state and will gain features over time and releases.

## What's new

The `TypedClient` is built around a fluent builder for easier request creation and a collection of structures and helpers that mimics as closely as possible the Elasticsearch JSON API.

As a first example, here is a search request:
```go
cfg := elasticsearch.Config{
// Define your configuration
}
es, _ := elasticsearch.NewTypedClient(cfg)
res, err := es.Search().
Index("index_name").
Request(&search.Request{
Query: &types.QueryContainer{
Match: map[types.Field]types.MatchQuery{
"name": {Query: "Foo"},
},
},
},
).Do(context.Background())
```

The `Request` uses the structures found in the `typedapi/types` package which will lead you along the possibilities. A builder for each structure that allows easier access and declaration is also provided.

More on the specifics and a few examples of standard use-cases can be found in the [TypedAPI section of the documentation](https://www.elastic.co/guide/en/elasticsearch/client/go-api/master/typedapi.html).

## Limitations

While most of the endpoints are covered, a few points are still being worked on and will be part of future releases:

* NDJSON endpoints: `bulk`, `msearch`, `msearch_template`, `ML.post_data`, `find_structure`, to name a few.
* Response and Errors structures with deserialization.

## Transport & config

While being different, the new API uses all the existing layers that were built so far, `elastic-transport-go` remains the preferred transport and all your configuration and credentials applies, same as before.

## Feedback

Feedback is very welcome, play with it, use it, let us know what you think!

## 8.3.0

## API

* `ML.InferTrainedModelDeployment` renamed to `InferTrainedModel`
* `ML.PreviewDatafeed` has two new parameters, `start` and `end`. [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/8.3/ml-preview-datafeed.html)
* `ML.StartTrainedModelDeployment` has three new parameters, `number_of_allocations`, `threads_per_allocation` and `queue_capacity`. [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/start-trained-model-deployment.html)
* `Cluster.DeleteVotingConfigExclusions` has a new `master_timeout` parameter.
* `Cluster.PostVotingConfigExclusions` has a new `master_timeout` parameter.
* `Snapshot.Get` has a new `index_names` parameters (boolean). Whether to include the name of each index in the snapshot. Defaults to true.

**New APIs**

* `Security.HasPrivilegesUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges-user-profile.html)

## 8.2.0
## Client

* Fixed a serialisation error for `retry_on_conflict` in the BulkIndexer. Thanks to @lpflpf for the help!
* Fixed a concurrent map error in the BulkIndexer when custom headers are applied. Thanks to @chzhuo for the contribution!

## API

**New APIs**

* `Cat.ComponentTemplates`
* `ML.GetMemoryStats` [documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/get-ml-memory.html)

* `Security.activateUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-activate-user-profile.html)
* `Security.disableUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/security-api-disable-user-profile.html)
* `Security.enableUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/security-api-enable-user-profile.html)
* `Security.getUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-profile.html)
* `Security.suggestUserProfiles` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/security-api-suggest-user-profile.html)
* `Security.updateUserProfileData` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-user-profile-data.html)

## 8.1.0
## API

* API is generated from the Elasticsearch 8.1.0 specification.

**New parameters**

* `WithWaitForCompletion` for `Indices.Forcemerge`
* `WithFeatures` for `Indices.Get`
* `WithForce` for `ML.DeleteTrainedModel`

**New APIs**

* `OidcAuthenticate`, `OidcLogout` and `OidcPrepareAuthentication` [see documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api.html#security-openid-apis)
* `TransformResetTransform`

## 8.0.0
## Client

* The client now uses `elastic-transport-go` dependency which lives in its [own repository](https://github.com/elastic/elastic-transport-go/).
* With the knewly extracted transport, the `retryOnTimeout` has been replaced with a `retryOnError` callback. This allows to select more finely which error should be retried by the client.
* `BulkIndexerItem` `Body` field is now an `io.ReadSeeker` allowing reread without increasing memory consumption.
* `BulkIndexerItem` know correctly uses the `routing` property instead of the deprecated `_routing`.

## API

* API is generated from the Elasticsearch 8.0.0 specification.

# 8.8.2

## Typed API

 * Fixed deserialization for `Suggest` in search responses.
 * Fixed double-quoted strings in deserialization for unions normalized as string. #684
 * Fixed handling of `core.Get` response when the index did not exist. #678

# 8.8.0

## API

**New APIs**

* `Watcher.GetSettings` Documentation: https://www.elastic.co/guide/en/elasticsearch/reference/8.8/watcher-api-get-settings.html
* `Watcher.UpdateSettings` Documentation: https://www.elastic.co/guide/en/elasticsearch/reference/8.8/watcher-api-update-settings.html

**Experimental APIs**

* `ML.DeleteDataLifecycle` Documentation: https://www.elastic.co/guide/en/elasticsearch/reference/8.8/dlm-delete-lifecycle.html
* `ML.ExplainDataLifecycle` Documentation: https://www.elastic.co/guide/en/elasticsearch/reference/8.8/dlm-explain-lifecycle.html
* `ML.GetDataLifecycle` Documentation: https://www.elastic.co/guide/en/elasticsearch/reference/8.8/dlm-get-lifecycle.html
* `ML.PutDataLifecycle` Documentation: https://www.elastic.co/guide/en/elasticsearch/reference/8.8/dlm-put-lifecycle.html

* `SearchApplications` https://www.elastic.co/guide/en/elasticsearch/reference/8.8/search-application-apis.html

# 8.7.0

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

# 8.6.0

## API

* `ML.StartTrainedModelDeployment`: Added `WithPriority`

**New APIs**
* `ML.UpdateTrainedModelDeployment`: Updates certain properties of trained model deployment.


## Client
**BulkIndexer**

Improvements were made to the BulkIndexer memory usage to allow better handling under burst use cases. Thanks to @christos68k and @rockdaboot !

# 8.5.0

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

# 8.4.0

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

# 8.4.0-alpha.2

This second prerelease of the 8.4.0 updates the API for the client and fixes the serialization for types using [additional properties](https://github.com/elastic/elasticsearch-specification/blob/main/docs/behaviors.md#additionalproperties--additionalproperty).

# 8.4.0-alpha.1

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

# 8.3.0

## API

* `ML.InferTrainedModelDeployment` renamed to `InferTrainedModel`
* `ML.PreviewDatafeed` has two new parameters, `start` and `end`. [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/8.3/ml-preview-datafeed.html)
* `ML.StartTrainedModelDeployment` has three new parameters, `number_of_allocations`, `threads_per_allocation` and `queue_capacity`. [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/master/start-trained-model-deployment.html)
* `Cluster.DeleteVotingConfigExclusions` has a new `master_timeout` parameter.
* `Cluster.PostVotingConfigExclusions` has a new `master_timeout` parameter.
* `Snapshot.Get` has a new `index_names` parameters (boolean). Whether to include the name of each index in the snapshot. Defaults to true.

**New APIs**

* `Security.HasPrivilegesUserProfile` (Experimental API) [Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges-user-profile.html)

# 8.2.0
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

# 8.1.0
## API

* API is generated from the Elasticsearch 8.1.0 specification.

**New parameters**

* `WithWaitForCompletion` for `Indices.Forcemerge`
* `WithFeatures` for `Indices.Get`
* `WithForce` for `ML.DeleteTrainedModel`

**New APIs**

* `OidcAuthenticate`, `OidcLogout` and `OidcPrepareAuthentication` [see documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api.html#security-openid-apis)
* `TransformResetTransform`

# 8.0.0
## Client

* The client now uses `elastic-transport-go` dependency which lives in its [own repository](https://github.com/elastic/elastic-transport-go/).
* With the knewly extracted transport, the `retryOnTimeout` has been replaced with a `retryOnError` callback. This allows to select more finely which error should be retried by the client.
* `BulkIndexerItem` `Body` field is now an `io.ReadSeeker` allowing reread without increasing memory consumption.
* `BulkIndexerItem` know correctly uses the `routing` property instead of the deprecated `_routing`.

## API

* API is generated from the Elasticsearch 8.0.0 specification.





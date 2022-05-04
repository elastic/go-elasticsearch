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





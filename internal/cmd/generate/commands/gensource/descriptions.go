package gensource

var apiDescriptionsYAML = `
---
bulk:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html
  description: |-
    Allows to perform multiple index/update/delete operations in a single request.

cat.aliases:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-alias.html
  description: |-
    Shows information about currently configured aliases to indices including filter and routing infos.

cat.allocation:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-allocation.html
  description: |-
    Provides a snapshot of how many shards are allocated to each data node and how much disk space they are using.

cat.count:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-count.html
  description: |-
    Provides quick access to the document count of the entire cluster, or individual indices.

cat.fielddata:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-fielddata.html
  description: |-
    Shows how much heap memory is currently being used by fielddata on every data node in the cluster.

cat.health:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-health.html
  description: |-
    Returns a concise representation of the cluster health.

cat.help:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat.html
  description:
    Returns help for the Cat APIs.

cat.indices:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-indices.html
  description: |-
    Returns information about indices: number of primaries and replicas, document counts, disk size, ...

cat.master:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-master.html
  description: |-
    Returns information about the master node.

cat.nodeattrs:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-nodeattrs.html
  description: |-
    Returns information about custom node attributes.

cat.nodes:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-nodes.html
  description: |-
    Returns basic statistics about performance of cluster nodes.

cat.pending_tasks:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-pending-tasks.html
  description: |-
    Returns a concise representation of the cluster pending tasks.

cat.plugins:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-plugins.html
  description: |-
    Returns information about installed plugins across nodes node.

cat.recovery:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-recovery.html
  description: |-
    Returns information about index shard recoveries, both on-going completed.

cat.repositories:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-repositories.html
  description: |-
    Returns information about snapshot repositories registered in the cluster.

cat.segments:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-segments.html
  description: |-
    Provides low-level information about the segments in the shards of an index.

cat.shards:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-shards.html
  description: |-
    Provides a detailed view of shard allocation on nodes.

cat.snapshots:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-snapshots.html
  description: |-
    Returns all snapshots in a specific repository.

cat.tasks:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/tasks.html
  description: |-
    Returns information about the tasks currently executing on one or more nodes in the cluster.

cat.templates:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-templates.html
  description: |-
    Returns information about existing templates.

cat.thread_pool:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-thread-pool.html
  description: |-
    Returns cluster-wide thread pool statistics per node.
    By default the active, queue and rejected statistics are returned for all thread pools.

clear_scroll:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-request-scroll.html
  description: |-
    Explicitly clears the search context for a scroll.

cluster.allocation_explain:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-allocation-explain.html
  description: |-
    Provides explanations for shard allocations in the cluster.

cluster.get_settings:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-update-settings.html
  description: |-
    Returns cluster settings.

cluster.health:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-health.html
  description: |-
    Returns basic information about the health of the cluster.

cluster.pending_tasks:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-pending.html
  description: |-
    Returns a list of any cluster-level changes (e.g. create index, update mapping,
    allocate or fail shard) which have not yet been executed.

cluster.put_settings:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-update-settings.html
  description: |-
    Updates the cluster settings.

cluster.remote_info:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-remote-info.html
  description: |-
    Returns the information about configured remote clusters.

cluster.reroute:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-reroute.html
  description: |-
    Allows to manually change the allocation of individual shards in the cluster.

cluster.state:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-state.html
  description: |-
    Returns a comprehensive information about the state of the cluster.

cluster.stats:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-stats.html
  description: |-
    Returns high-level overview of cluster statistics.

count:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-count.html
  description: |-
    Returns number of documents matching a query.

create:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html
  description: |-
    Creates a new document in the index.

    Returns a 409 response when a document with a same ID already exists in the index.

delete:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-delete.html
  description: |-
    Removes a document from the index.

delete_by_query:
  # https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-delete-by-query.html
  description: |-
    Deletes documents matching the provided query.

delete_by_query_rethrottle:
  # https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
  description: |-
    Changes the number of requests per second for a particular Delete By Query operation.

delete_script:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-scripting.html
  description: |-
    Deletes a script.

exists:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html
  description: |-
    Returns information about whether a document exists in an index.

exists_source:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html
  description: |-
    Returns information about whether a document source exists in an index.

explain:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-explain.html
  description: |-
    Returns information about why a specific matches (or doesn't match) a query.

field_caps:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-field-caps.html
  description: |-
    Returns the information about the capabilities of fields among multiple indices.

get:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html
  description: |-
    Returns a document.

get_script:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-scripting.html
  description: |-
    Returns a script.

get_source:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html
  description: |-
    Returns the source of a document.

index:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html
  description: |-
    Creates or updates a document in an index.

indices.analyze:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-analyze.html
  description: |-
    Performs the analysis process on a text and return the tokens breakdown of the text.

indices.clear_cache:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-clearcache.html
  description: |-
    Clears all or specific caches for one or more indices.

indices.close:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-open-close.html
  description: |-
    Closes an index.

indices.create:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-create-index.html
  description: |-
    Creates an index with optional settings and mappings.

indices.delete:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-delete-index.html
  description: |-
    Deletes an index.

indices.delete_alias:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html
  description: |-
    Deletes an alias.

indices.delete_template:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-templates.html
  description: |-
    Deletes an index template.

indices.exists:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-exists.html
  description: |-
    Returns information about whether a particular index exists.

indices.exists_alias:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html
  description: |-
    Returns information about whether a particular alias exists.

indices.exists_template:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-templates.html
  description: |-
    Returns information about whether a particular index template exists.

indices.exists_type:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-types-exists.html
  description: |-
    Returns information about whether a particular document type exists. (DEPRECATED)

indices.flush:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-flush.html
  description: |-
    Performs the flush operation on one or more indices.

indices.flush_synced:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-synced-flush.html
  description: |-
    Performs a synced flush operation on one or more indices.

indices.forcemerge:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-forcemerge.html
  description: |-
    Performs the force merge operation on one or more indices.

indices.get:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-index.html
  description: |-
    Returns information about one or more indices.

indices.get_alias:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html
  description: |-
    Returns an alias.

indices.get_field_mapping:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-field-mapping.html
  description: |-
    Returns mapping for one or more fields.

indices.get_mapping:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-mapping.html
  description: |-
    Returns mappings for one or more indices.

indices.get_settings:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-settings.html
  description: |-
    Returns settings for one or more indices.

indices.get_template:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-templates.html
  description: |-
    Returns an index template.

indices.get_upgrade:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-upgrade.html
  description: |-
    The _upgrade API is no longer useful and will be removed.

indices.open:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-open-close.html
  description: |-
    Opens an index.

indices.put_alias:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html
  description: |-
    Creates or updates an alias.

indices.put_mapping:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-put-mapping.html
  description: |-
    Updates the index mappings.

indices.put_settings:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-update-settings.html
  description: |-
    Updates the index settings.

indices.put_template:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-templates.html
  description: |-
    Creates or updates an index template.

indices.recovery:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-recovery.html
  description: |-
    Returns information about ongoing index shard recoveries.

indices.refresh:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-refresh.html
  description: |-
    Performs the refresh operation in one or more indices.

indices.rollover:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-rollover-index.html
  description: |-
    Updates an alias to point to a new index when the existing index
    is considered to be too large or too old.

indices.segments:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-segments.html
  description: |-
    Provides low-level information about segments in a Lucene index.

indices.shard_stores:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-shards-stores.html
  description: |-
    Provides store information for shard copies of indices.

indices.shrink:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-shrink-index.html
  description: |-
    Allow to shrink an existing index into a new index with fewer primary shards.

indices.split:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-split-index.html
  description: |-
    Allows you to split an existing index into a new index with more primary shards.

indices.stats:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-stats.html
  description: |-
    Provides statistics on operations happening in an index.

indices.update_aliases:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html
  description: |-
    Updates index aliases.

indices.upgrade:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-upgrade.html
  description: |-
    The _upgrade API is no longer useful and will be removed.

indices.validate_query:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-validate.html
  description: |-
    Allows a user to validate a potentially expensive query without executing it.

info:
  # http://www.elastic.co/guide/
  description: |-
    Returns basic information about the cluster.

ingest.delete_pipeline:
  url: https://www.elastic.co/guide/en/elasticsearch/reference/master/delete-pipeline-api.html
  description: |-
    Deletes a pipeline.

ingest.get_pipeline:
  url: https://www.elastic.co/guide/en/elasticsearch/reference/master/get-pipeline-api.html
  description: |-
    Returns a pipeline.

ingest.processor_grok:
  # https://www.elastic.co/guide/en/elasticsearch/plugins/master/ingest.html
  description: |-
    Returns a list of the built-in patterns.

ingest.put_pipeline:
  url: https://www.elastic.co/guide/en/elasticsearch/reference/master/put-pipeline-api.html
  description: |-
    Creates or updates a pipeline.

ingest.simulate:
  url: https://www.elastic.co/guide/en/elasticsearch/reference/master/simulate-pipeline-api.html
  description: |-
    Allows to simulate a pipeline with example documents.

mget:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-multi-get.html
  description: |-
    Allows to get multiple documents in one request.

msearch:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-multi-search.html
  description: |-
    Allows to execute several search operations in one request.

msearch_template:
  # http://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
  description: |-
    Allows to execute several search template operations in one request.

mtermvectors:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-multi-termvectors.html
  description: |-
    Returns multiple termvectors in one request.

nodes.hot_threads:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-nodes-hot-threads.html
  description: |-
    Returns information about hot threads on each node in the cluster.

nodes.info:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-nodes-info.html
  description: |-
    Returns information about nodes in the cluster.

nodes.reload_secure_settings:
  # https://www.elastic.co/guide/en/elasticsearch/reference/master/secure-settings.html#reloadable-secure-settings
  description: |-
    Reloads secure settings.

nodes.stats:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-nodes-stats.html
  description: |-
    Returns statistical information about nodes in the cluster.

nodes.usage:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-nodes-usage.html
  description: |-
    Returns low-level information about REST actions usage on nodes.

ping:
  # http://www.elastic.co/guide/
  description:
    Returns whether the cluster is running.

put_script:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-scripting.html
  description: |-
    Creates or updates a script.

rank_eval:
  # https://www.elastic.co/guide/en/elasticsearch/reference/master/search-rank-eval.html
  description: |-
    Allows to evaluate the quality of ranked search results over a set of typical search queries

reindex:
  # https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-reindex.html
  description: |-
    Allows to copy documents from one index to another, optionally filtering the source
    documents by a query, changing the destination index settings, or fetching the
    documents from a remote cluster.

reindex_rethrottle:
  # https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-reindex.html
  description: |-
    Changes the number of requests per second for a particular Reindex operation.

render_search_template:
  # http://www.elasticsearch.org/guide/en/elasticsearch/reference/master/search-template.html
  description: |-
    Allows to use the Mustache language to pre-render a search definition.

scripts_painless_execute:
  # https://www.elastic.co/guide/en/elasticsearch/painless/master/painless-execute-api.html
  description: |-
    Allows an arbitrary script to be executed and a result to be returned

scripts_painless_context:
  # https://www.elastic.co/guide/en/elasticsearch/painless/master/painless-execute-api.html
  description: |-
    Allows to query context information.

scroll:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-request-scroll.html
  description: |-
    Allows to retrieve a large numbers of results from a single search request.

search:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-search.html
  description: |-
    Returns results matching a query.

search_shards:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/search-shards.html
  description: |-
    Returns information about the indices and shards that a search request would be executed against.

search_template:
  # http://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
  description: |-
    Allows to use the Mustache language to pre-render a search definition.

snapshot.create:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Creates a snapshot in a repository.

snapshot.create_repository:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Creates a repository.

snapshot.delete:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Deletes a snapshot.

snapshot.delete_repository:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Deletes a repository.

snapshot.get:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Returns information about a snapshot.

snapshot.get_repository:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Returns information about a repository.

snapshot.restore:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Restores a snapshot.

snapshot.status:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Returns information about the status of a snapshot.

snapshot.verify_repository:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
  description: |-
    Verifies a repository.

tasks.cancel:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/tasks.html
  description: |-
    Cancels a task, if it can be cancelled through an API.

tasks.get:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/tasks.html
  description: |-
    Returns information about a task.

tasks.list:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/tasks.html
  description: |-
    Returns a list of tasks.

termvectors:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-termvectors.html
  description: |-
    Returns information and statistics about terms in the fields of a particular document.

update:
  # http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-update.html
  description: |-
    Updates a document with a script or partial document.

update_by_query:
  # https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-update-by-query.html
  description: |-
    Performs an update on every document in the index without changing the source,
    for example to pick up a mapping change.

update_by_query_rethrottle:
  # https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
  description: |-
    Changes the number of requests per second for a particular Update By Query operation.
`

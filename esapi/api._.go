// Code generated from specification version 5.6.15 (fe7575a32e2|v5.6.15): DO NOT EDIT

package esapi

// API contains the Elasticsearch APIs
//
type API struct {
	Cat      *Cat
	Cluster  *Cluster
	Indices  *Indices
	Ingest   *Ingest
	Nodes    *Nodes
	Remote   *Remote
	Snapshot *Snapshot
	Tasks    *Tasks

	Bulk                 Bulk
	ClearScroll          ClearScroll
	Count                Count
	CountPercolate       CountPercolate
	Create               Create
	Delete               Delete
	DeleteByQuery        DeleteByQuery
	DeleteScript         DeleteScript
	DeleteTemplate       DeleteTemplate
	Exists               Exists
	ExistsSource         ExistsSource
	Explain              Explain
	FieldCaps            FieldCaps
	FieldStats           FieldStats
	Get                  Get
	GetScript            GetScript
	GetSource            GetSource
	GetTemplate          GetTemplate
	Index                Index
	Info                 Info
	Mget                 Mget
	Mpercolate           Mpercolate
	Msearch              Msearch
	MsearchTemplate      MsearchTemplate
	Mtermvectors         Mtermvectors
	Percolate            Percolate
	Ping                 Ping
	PutScript            PutScript
	PutTemplate          PutTemplate
	Reindex              Reindex
	ReindexRethrottle    ReindexRethrottle
	RenderSearchTemplate RenderSearchTemplate
	Scroll               Scroll
	Search               Search
	SearchShards         SearchShards
	SearchTemplate       SearchTemplate
	Suggest              Suggest
	Termvectors          Termvectors
	Update               Update
	UpdateByQuery        UpdateByQuery
}

// Cat contains the Cat APIs
type Cat struct {
	Aliases      CatAliases
	Allocation   CatAllocation
	Count        CatCount
	Fielddata    CatFielddata
	Health       CatHealth
	Help         CatHelp
	Indices      CatIndices
	Master       CatMaster
	Nodeattrs    CatNodeattrs
	Nodes        CatNodes
	PendingTasks CatPendingTasks
	Plugins      CatPlugins
	Recovery     CatRecovery
	Repositories CatRepositories
	Segments     CatSegments
	Shards       CatShards
	Snapshots    CatSnapshots
	Tasks        CatTasks
	Templates    CatTemplates
	ThreadPool   CatThreadPool
}

// Cluster contains the Cluster APIs
type Cluster struct {
	AllocationExplain ClusterAllocationExplain
	GetSettings       ClusterGetSettings
	Health            ClusterHealth
	PendingTasks      ClusterPendingTasks
	PutSettings       ClusterPutSettings
	RemoteInfo        ClusterRemoteInfo
	Reroute           ClusterReroute
	State             ClusterState
	Stats             ClusterStats
}

// Indices contains the Indices APIs
type Indices struct {
	Analyze         IndicesAnalyze
	ClearCache      IndicesClearCache
	Close           IndicesClose
	Create          IndicesCreate
	Delete          IndicesDelete
	DeleteAlias     IndicesDeleteAlias
	DeleteTemplate  IndicesDeleteTemplate
	Exists          IndicesExists
	ExistsAlias     IndicesExistsAlias
	ExistsTemplate  IndicesExistsTemplate
	ExistsType      IndicesExistsType
	Flush           IndicesFlush
	FlushSynced     IndicesFlushSynced
	Forcemerge      IndicesForcemerge
	Get             IndicesGet
	GetAlias        IndicesGetAlias
	GetFieldMapping IndicesGetFieldMapping
	GetMapping      IndicesGetMapping
	GetSettings     IndicesGetSettings
	GetTemplate     IndicesGetTemplate
	GetUpgrade      IndicesGetUpgrade
	Open            IndicesOpen
	PutAlias        IndicesPutAlias
	PutMapping      IndicesPutMapping
	PutSettings     IndicesPutSettings
	PutTemplate     IndicesPutTemplate
	Recovery        IndicesRecovery
	Refresh         IndicesRefresh
	Rollover        IndicesRollover
	Segments        IndicesSegments
	ShardStores     IndicesShardStores
	Shrink          IndicesShrink
	Stats           IndicesStats
	UpdateAliases   IndicesUpdateAliases
	Upgrade         IndicesUpgrade
	ValidateQuery   IndicesValidateQuery
}

// Ingest contains the Ingest APIs
type Ingest struct {
	DeletePipeline IngestDeletePipeline
	GetPipeline    IngestGetPipeline
	ProcessorGrok  IngestProcessorGrok
	PutPipeline    IngestPutPipeline
	Simulate       IngestSimulate
}

// Nodes contains the Nodes APIs
type Nodes struct {
	HotThreads NodesHotThreads
	Info       NodesInfo
	Stats      NodesStats
}

// Remote contains the Remote APIs
type Remote struct {
}

// Snapshot contains the Snapshot APIs
type Snapshot struct {
	Create           SnapshotCreate
	CreateRepository SnapshotCreateRepository
	Delete           SnapshotDelete
	DeleteRepository SnapshotDeleteRepository
	Get              SnapshotGet
	GetRepository    SnapshotGetRepository
	Restore          SnapshotRestore
	Status           SnapshotStatus
	VerifyRepository SnapshotVerifyRepository
}

// Tasks contains the Tasks APIs
type Tasks struct {
	Cancel TasksCancel
	Get    TasksGet
	List   TasksList
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
		Bulk:                 newBulkFunc(t),
		ClearScroll:          newClearScrollFunc(t),
		Count:                newCountFunc(t),
		CountPercolate:       newCountPercolateFunc(t),
		Create:               newCreateFunc(t),
		Delete:               newDeleteFunc(t),
		DeleteByQuery:        newDeleteByQueryFunc(t),
		DeleteScript:         newDeleteScriptFunc(t),
		DeleteTemplate:       newDeleteTemplateFunc(t),
		Exists:               newExistsFunc(t),
		ExistsSource:         newExistsSourceFunc(t),
		Explain:              newExplainFunc(t),
		FieldCaps:            newFieldCapsFunc(t),
		FieldStats:           newFieldStatsFunc(t),
		Get:                  newGetFunc(t),
		GetScript:            newGetScriptFunc(t),
		GetSource:            newGetSourceFunc(t),
		GetTemplate:          newGetTemplateFunc(t),
		Index:                newIndexFunc(t),
		Info:                 newInfoFunc(t),
		Mget:                 newMgetFunc(t),
		Mpercolate:           newMpercolateFunc(t),
		Msearch:              newMsearchFunc(t),
		MsearchTemplate:      newMsearchTemplateFunc(t),
		Mtermvectors:         newMtermvectorsFunc(t),
		Percolate:            newPercolateFunc(t),
		Ping:                 newPingFunc(t),
		PutScript:            newPutScriptFunc(t),
		PutTemplate:          newPutTemplateFunc(t),
		Reindex:              newReindexFunc(t),
		ReindexRethrottle:    newReindexRethrottleFunc(t),
		RenderSearchTemplate: newRenderSearchTemplateFunc(t),
		Scroll:               newScrollFunc(t),
		Search:               newSearchFunc(t),
		SearchShards:         newSearchShardsFunc(t),
		SearchTemplate:       newSearchTemplateFunc(t),
		Suggest:              newSuggestFunc(t),
		Termvectors:          newTermvectorsFunc(t),
		Update:               newUpdateFunc(t),
		UpdateByQuery:        newUpdateByQueryFunc(t),
		Cat: &Cat{
			Aliases:      newCatAliasesFunc(t),
			Allocation:   newCatAllocationFunc(t),
			Count:        newCatCountFunc(t),
			Fielddata:    newCatFielddataFunc(t),
			Health:       newCatHealthFunc(t),
			Help:         newCatHelpFunc(t),
			Indices:      newCatIndicesFunc(t),
			Master:       newCatMasterFunc(t),
			Nodeattrs:    newCatNodeattrsFunc(t),
			Nodes:        newCatNodesFunc(t),
			PendingTasks: newCatPendingTasksFunc(t),
			Plugins:      newCatPluginsFunc(t),
			Recovery:     newCatRecoveryFunc(t),
			Repositories: newCatRepositoriesFunc(t),
			Segments:     newCatSegmentsFunc(t),
			Shards:       newCatShardsFunc(t),
			Snapshots:    newCatSnapshotsFunc(t),
			Tasks:        newCatTasksFunc(t),
			Templates:    newCatTemplatesFunc(t),
			ThreadPool:   newCatThreadPoolFunc(t),
		},
		Cluster: &Cluster{
			AllocationExplain: newClusterAllocationExplainFunc(t),
			GetSettings:       newClusterGetSettingsFunc(t),
			Health:            newClusterHealthFunc(t),
			PendingTasks:      newClusterPendingTasksFunc(t),
			PutSettings:       newClusterPutSettingsFunc(t),
			RemoteInfo:        newClusterRemoteInfoFunc(t),
			Reroute:           newClusterRerouteFunc(t),
			State:             newClusterStateFunc(t),
			Stats:             newClusterStatsFunc(t),
		},
		Indices: &Indices{
			Analyze:         newIndicesAnalyzeFunc(t),
			ClearCache:      newIndicesClearCacheFunc(t),
			Close:           newIndicesCloseFunc(t),
			Create:          newIndicesCreateFunc(t),
			Delete:          newIndicesDeleteFunc(t),
			DeleteAlias:     newIndicesDeleteAliasFunc(t),
			DeleteTemplate:  newIndicesDeleteTemplateFunc(t),
			Exists:          newIndicesExistsFunc(t),
			ExistsAlias:     newIndicesExistsAliasFunc(t),
			ExistsTemplate:  newIndicesExistsTemplateFunc(t),
			ExistsType:      newIndicesExistsTypeFunc(t),
			Flush:           newIndicesFlushFunc(t),
			FlushSynced:     newIndicesFlushSyncedFunc(t),
			Forcemerge:      newIndicesForcemergeFunc(t),
			Get:             newIndicesGetFunc(t),
			GetAlias:        newIndicesGetAliasFunc(t),
			GetFieldMapping: newIndicesGetFieldMappingFunc(t),
			GetMapping:      newIndicesGetMappingFunc(t),
			GetSettings:     newIndicesGetSettingsFunc(t),
			GetTemplate:     newIndicesGetTemplateFunc(t),
			GetUpgrade:      newIndicesGetUpgradeFunc(t),
			Open:            newIndicesOpenFunc(t),
			PutAlias:        newIndicesPutAliasFunc(t),
			PutMapping:      newIndicesPutMappingFunc(t),
			PutSettings:     newIndicesPutSettingsFunc(t),
			PutTemplate:     newIndicesPutTemplateFunc(t),
			Recovery:        newIndicesRecoveryFunc(t),
			Refresh:         newIndicesRefreshFunc(t),
			Rollover:        newIndicesRolloverFunc(t),
			Segments:        newIndicesSegmentsFunc(t),
			ShardStores:     newIndicesShardStoresFunc(t),
			Shrink:          newIndicesShrinkFunc(t),
			Stats:           newIndicesStatsFunc(t),
			UpdateAliases:   newIndicesUpdateAliasesFunc(t),
			Upgrade:         newIndicesUpgradeFunc(t),
			ValidateQuery:   newIndicesValidateQueryFunc(t),
		},
		Ingest: &Ingest{
			DeletePipeline: newIngestDeletePipelineFunc(t),
			GetPipeline:    newIngestGetPipelineFunc(t),
			ProcessorGrok:  newIngestProcessorGrokFunc(t),
			PutPipeline:    newIngestPutPipelineFunc(t),
			Simulate:       newIngestSimulateFunc(t),
		},
		Nodes: &Nodes{
			HotThreads: newNodesHotThreadsFunc(t),
			Info:       newNodesInfoFunc(t),
			Stats:      newNodesStatsFunc(t),
		},
		Remote: &Remote{},
		Snapshot: &Snapshot{
			Create:           newSnapshotCreateFunc(t),
			CreateRepository: newSnapshotCreateRepositoryFunc(t),
			Delete:           newSnapshotDeleteFunc(t),
			DeleteRepository: newSnapshotDeleteRepositoryFunc(t),
			Get:              newSnapshotGetFunc(t),
			GetRepository:    newSnapshotGetRepositoryFunc(t),
			Restore:          newSnapshotRestoreFunc(t),
			Status:           newSnapshotStatusFunc(t),
			VerifyRepository: newSnapshotVerifyRepositoryFunc(t),
		},
		Tasks: &Tasks{
			Cancel: newTasksCancelFunc(t),
			Get:    newTasksGetFunc(t),
			List:   newTasksListFunc(t),
		},
	}
}

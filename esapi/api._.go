// Code generated from specification version 8.0.0 (0fd6a17aa17f73a9f981394b0a224c3afec57e90): DO NOT EDIT

package esapi

// API contains the Elasticsearch APIs
//
type API struct {
	Cat         *Cat
	Cluster     *Cluster
	Indices     *Indices
	Ingest      *Ingest
	Nodes       *Nodes
	Remote      *Remote
	Snapshot    *Snapshot
	Tasks       *Tasks
	AsyncSearch *AsyncSearch
	CCR         *CCR
	ILM         *ILM
	License     *License
	Migration   *Migration
	ML          *ML
	Monitoring  *Monitoring
	Rollup      *Rollup
	Security    *Security
	SQL         *SQL
	SSL         *SSL
	Watcher     *Watcher
	XPack       *XPack

	AutoscalingDeleteAutoscalingPolicy            AutoscalingDeleteAutoscalingPolicy
	AutoscalingGetAutoscalingDecision             AutoscalingGetAutoscalingDecision
	AutoscalingGetAutoscalingPolicy               AutoscalingGetAutoscalingPolicy
	AutoscalingPutAutoscalingPolicy               AutoscalingPutAutoscalingPolicy
	Bulk                                          Bulk
	ClearScroll                                   ClearScroll
	Count                                         Count
	Create                                        Create
	DanglingIndicesDeleteDanglingIndex            DanglingIndicesDeleteDanglingIndex
	DanglingIndicesImportDanglingIndex            DanglingIndicesImportDanglingIndex
	DanglingIndicesListDanglingIndices            DanglingIndicesListDanglingIndices
	DataFrameTransformDeprecatedDeleteTransform   DataFrameTransformDeprecatedDeleteTransform
	DataFrameTransformDeprecatedGetTransform      DataFrameTransformDeprecatedGetTransform
	DataFrameTransformDeprecatedGetTransformStats DataFrameTransformDeprecatedGetTransformStats
	DataFrameTransformDeprecatedPreviewTransform  DataFrameTransformDeprecatedPreviewTransform
	DataFrameTransformDeprecatedPutTransform      DataFrameTransformDeprecatedPutTransform
	DataFrameTransformDeprecatedStartTransform    DataFrameTransformDeprecatedStartTransform
	DataFrameTransformDeprecatedStopTransform     DataFrameTransformDeprecatedStopTransform
	DataFrameTransformDeprecatedUpdateTransform   DataFrameTransformDeprecatedUpdateTransform
	DeleteByQuery                                 DeleteByQuery
	DeleteByQueryRethrottle                       DeleteByQueryRethrottle
	Delete                                        Delete
	DeleteScript                                  DeleteScript
	EnrichDeletePolicy                            EnrichDeletePolicy
	EnrichExecutePolicy                           EnrichExecutePolicy
	EnrichGetPolicy                               EnrichGetPolicy
	EnrichPutPolicy                               EnrichPutPolicy
	EnrichStats                                   EnrichStats
	EqlDelete                                     EqlDelete
	EqlGet                                        EqlGet
	EqlSearch                                     EqlSearch
	Exists                                        Exists
	ExistsSource                                  ExistsSource
	Explain                                       Explain
	FieldCaps                                     FieldCaps
	Get                                           Get
	GetScriptContext                              GetScriptContext
	GetScriptLanguages                            GetScriptLanguages
	GetScript                                     GetScript
	GetSource                                     GetSource
	GraphExplore                                  GraphExplore
	Index                                         Index
	Info                                          Info
	Mget                                          Mget
	Msearch                                       Msearch
	MsearchTemplate                               MsearchTemplate
	Mtermvectors                                  Mtermvectors
	Ping                                          Ping
	PutScript                                     PutScript
	RankEval                                      RankEval
	Reindex                                       Reindex
	ReindexRethrottle                             ReindexRethrottle
	RenderSearchTemplate                          RenderSearchTemplate
	ScriptsPainlessExecute                        ScriptsPainlessExecute
	Scroll                                        Scroll
	Search                                        Search
	SearchShards                                  SearchShards
	SearchTemplate                                SearchTemplate
	SearchableSnapshotsClearCache                 SearchableSnapshotsClearCache
	SearchableSnapshotsMount                      SearchableSnapshotsMount
	SearchableSnapshotsRepositoryStats            SearchableSnapshotsRepositoryStats
	SearchableSnapshotsStats                      SearchableSnapshotsStats
	SlmDeleteLifecycle                            SlmDeleteLifecycle
	SlmExecuteLifecycle                           SlmExecuteLifecycle
	SlmExecuteRetention                           SlmExecuteRetention
	SlmGetLifecycle                               SlmGetLifecycle
	SlmGetStats                                   SlmGetStats
	SlmGetStatus                                  SlmGetStatus
	SlmPutLifecycle                               SlmPutLifecycle
	SlmStart                                      SlmStart
	SlmStop                                       SlmStop
	Termvectors                                   Termvectors
	TransformDeleteTransform                      TransformDeleteTransform
	TransformGetTransform                         TransformGetTransform
	TransformGetTransformStats                    TransformGetTransformStats
	TransformPreviewTransform                     TransformPreviewTransform
	TransformPutTransform                         TransformPutTransform
	TransformStartTransform                       TransformStartTransform
	TransformStopTransform                        TransformStopTransform
	TransformUpdateTransform                      TransformUpdateTransform
	UpdateByQuery                                 UpdateByQuery
	UpdateByQueryRethrottle                       UpdateByQueryRethrottle
	Update                                        Update
}

// Cat contains the Cat APIs
type Cat struct {
	Aliases              CatAliases
	Allocation           CatAllocation
	Count                CatCount
	Fielddata            CatFielddata
	Health               CatHealth
	Help                 CatHelp
	Indices              CatIndices
	MLDataFrameAnalytics CatMLDataFrameAnalytics
	MLDatafeeds          CatMLDatafeeds
	MLJobs               CatMLJobs
	MLTrainedModels      CatMLTrainedModels
	Master               CatMaster
	Nodeattrs            CatNodeattrs
	Nodes                CatNodes
	PendingTasks         CatPendingTasks
	Plugins              CatPlugins
	Recovery             CatRecovery
	Repositories         CatRepositories
	Segments             CatSegments
	Shards               CatShards
	Snapshots            CatSnapshots
	Tasks                CatTasks
	Templates            CatTemplates
	ThreadPool           CatThreadPool
	Transforms           CatTransforms
}

// NewCat creates a new API client for Cat APIs
//
func NewCat(t Transport) *Cat {
	return &Cat{
		Aliases:              newCatAliasesFunc(t),
		Allocation:           newCatAllocationFunc(t),
		Count:                newCatCountFunc(t),
		Fielddata:            newCatFielddataFunc(t),
		Health:               newCatHealthFunc(t),
		Help:                 newCatHelpFunc(t),
		Indices:              newCatIndicesFunc(t),
		MLDataFrameAnalytics: newCatMLDataFrameAnalyticsFunc(t),
		MLDatafeeds:          newCatMLDatafeedsFunc(t),
		MLJobs:               newCatMLJobsFunc(t),
		MLTrainedModels:      newCatMLTrainedModelsFunc(t),
		Master:               newCatMasterFunc(t),
		Nodeattrs:            newCatNodeattrsFunc(t),
		Nodes:                newCatNodesFunc(t),
		PendingTasks:         newCatPendingTasksFunc(t),
		Plugins:              newCatPluginsFunc(t),
		Recovery:             newCatRecoveryFunc(t),
		Repositories:         newCatRepositoriesFunc(t),
		Segments:             newCatSegmentsFunc(t),
		Shards:               newCatShardsFunc(t),
		Snapshots:            newCatSnapshotsFunc(t),
		Tasks:                newCatTasksFunc(t),
		Templates:            newCatTemplatesFunc(t),
		ThreadPool:           newCatThreadPoolFunc(t),
		Transforms:           newCatTransformsFunc(t),
	}
}

// Cluster contains the Cluster APIs
type Cluster struct {
	AllocationExplain            ClusterAllocationExplain
	DeleteComponentTemplate      ClusterDeleteComponentTemplate
	DeleteVotingConfigExclusions ClusterDeleteVotingConfigExclusions
	ExistsComponentTemplate      ClusterExistsComponentTemplate
	GetComponentTemplate         ClusterGetComponentTemplate
	GetSettings                  ClusterGetSettings
	Health                       ClusterHealth
	PendingTasks                 ClusterPendingTasks
	PostVotingConfigExclusions   ClusterPostVotingConfigExclusions
	PutComponentTemplate         ClusterPutComponentTemplate
	PutSettings                  ClusterPutSettings
	RemoteInfo                   ClusterRemoteInfo
	Reroute                      ClusterReroute
	State                        ClusterState
	Stats                        ClusterStats
}

// NewCluster creates a new API client for Cluster APIs
//
func NewCluster(t Transport) *Cluster {
	return &Cluster{
		AllocationExplain:            newClusterAllocationExplainFunc(t),
		DeleteComponentTemplate:      newClusterDeleteComponentTemplateFunc(t),
		DeleteVotingConfigExclusions: newClusterDeleteVotingConfigExclusionsFunc(t),
		ExistsComponentTemplate:      newClusterExistsComponentTemplateFunc(t),
		GetComponentTemplate:         newClusterGetComponentTemplateFunc(t),
		GetSettings:                  newClusterGetSettingsFunc(t),
		Health:                       newClusterHealthFunc(t),
		PendingTasks:                 newClusterPendingTasksFunc(t),
		PostVotingConfigExclusions:   newClusterPostVotingConfigExclusionsFunc(t),
		PutComponentTemplate:         newClusterPutComponentTemplateFunc(t),
		PutSettings:                  newClusterPutSettingsFunc(t),
		RemoteInfo:                   newClusterRemoteInfoFunc(t),
		Reroute:                      newClusterRerouteFunc(t),
		State:                        newClusterStateFunc(t),
		Stats:                        newClusterStatsFunc(t),
	}
}

// Indices contains the Indices APIs
type Indices struct {
	AddBlock              IndicesAddBlock
	Analyze               IndicesAnalyze
	ClearCache            IndicesClearCache
	Clone                 IndicesClone
	Close                 IndicesClose
	CreateDataStream      IndicesCreateDataStream
	Create                IndicesCreate
	DataStreamsStats      IndicesDataStreamsStats
	DeleteAlias           IndicesDeleteAlias
	DeleteDataStream      IndicesDeleteDataStream
	DeleteIndexTemplate   IndicesDeleteIndexTemplate
	Delete                IndicesDelete
	DeleteTemplate        IndicesDeleteTemplate
	ExistsAlias           IndicesExistsAlias
	ExistsDocumentType    IndicesExistsDocumentType
	ExistsIndexTemplate   IndicesExistsIndexTemplate
	Exists                IndicesExists
	ExistsTemplate        IndicesExistsTemplate
	Flush                 IndicesFlush
	Forcemerge            IndicesForcemerge
	Freeze                IndicesFreeze
	GetAlias              IndicesGetAlias
	GetDataStream         IndicesGetDataStream
	GetFieldMapping       IndicesGetFieldMapping
	GetIndexTemplate      IndicesGetIndexTemplate
	GetMapping            IndicesGetMapping
	Get                   IndicesGet
	GetSettings           IndicesGetSettings
	GetTemplate           IndicesGetTemplate
	GetUpgrade            IndicesGetUpgrade
	Open                  IndicesOpen
	PutAlias              IndicesPutAlias
	PutIndexTemplate      IndicesPutIndexTemplate
	PutMapping            IndicesPutMapping
	PutSettings           IndicesPutSettings
	PutTemplate           IndicesPutTemplate
	Recovery              IndicesRecovery
	Refresh               IndicesRefresh
	ReloadSearchAnalyzers IndicesReloadSearchAnalyzers
	ResolveIndex          IndicesResolveIndex
	Rollover              IndicesRollover
	Segments              IndicesSegments
	ShardStores           IndicesShardStores
	Shrink                IndicesShrink
	SimulateIndexTemplate IndicesSimulateIndexTemplate
	SimulateTemplate      IndicesSimulateTemplate
	Split                 IndicesSplit
	Stats                 IndicesStats
	Unfreeze              IndicesUnfreeze
	UpdateAliases         IndicesUpdateAliases
	Upgrade               IndicesUpgrade
	ValidateQuery         IndicesValidateQuery
}

// NewIndices creates a new API client for Indices APIs
//
func NewIndices(t Transport) *Indices {
	return &Indices{
		AddBlock:              newIndicesAddBlockFunc(t),
		Analyze:               newIndicesAnalyzeFunc(t),
		ClearCache:            newIndicesClearCacheFunc(t),
		Clone:                 newIndicesCloneFunc(t),
		Close:                 newIndicesCloseFunc(t),
		CreateDataStream:      newIndicesCreateDataStreamFunc(t),
		Create:                newIndicesCreateFunc(t),
		DataStreamsStats:      newIndicesDataStreamsStatsFunc(t),
		DeleteAlias:           newIndicesDeleteAliasFunc(t),
		DeleteDataStream:      newIndicesDeleteDataStreamFunc(t),
		DeleteIndexTemplate:   newIndicesDeleteIndexTemplateFunc(t),
		Delete:                newIndicesDeleteFunc(t),
		DeleteTemplate:        newIndicesDeleteTemplateFunc(t),
		ExistsAlias:           newIndicesExistsAliasFunc(t),
		ExistsDocumentType:    newIndicesExistsDocumentTypeFunc(t),
		ExistsIndexTemplate:   newIndicesExistsIndexTemplateFunc(t),
		Exists:                newIndicesExistsFunc(t),
		ExistsTemplate:        newIndicesExistsTemplateFunc(t),
		Flush:                 newIndicesFlushFunc(t),
		Forcemerge:            newIndicesForcemergeFunc(t),
		Freeze:                newIndicesFreezeFunc(t),
		GetAlias:              newIndicesGetAliasFunc(t),
		GetDataStream:         newIndicesGetDataStreamFunc(t),
		GetFieldMapping:       newIndicesGetFieldMappingFunc(t),
		GetIndexTemplate:      newIndicesGetIndexTemplateFunc(t),
		GetMapping:            newIndicesGetMappingFunc(t),
		Get:                   newIndicesGetFunc(t),
		GetSettings:           newIndicesGetSettingsFunc(t),
		GetTemplate:           newIndicesGetTemplateFunc(t),
		GetUpgrade:            newIndicesGetUpgradeFunc(t),
		Open:                  newIndicesOpenFunc(t),
		PutAlias:              newIndicesPutAliasFunc(t),
		PutIndexTemplate:      newIndicesPutIndexTemplateFunc(t),
		PutMapping:            newIndicesPutMappingFunc(t),
		PutSettings:           newIndicesPutSettingsFunc(t),
		PutTemplate:           newIndicesPutTemplateFunc(t),
		Recovery:              newIndicesRecoveryFunc(t),
		Refresh:               newIndicesRefreshFunc(t),
		ReloadSearchAnalyzers: newIndicesReloadSearchAnalyzersFunc(t),
		ResolveIndex:          newIndicesResolveIndexFunc(t),
		Rollover:              newIndicesRolloverFunc(t),
		Segments:              newIndicesSegmentsFunc(t),
		ShardStores:           newIndicesShardStoresFunc(t),
		Shrink:                newIndicesShrinkFunc(t),
		SimulateIndexTemplate: newIndicesSimulateIndexTemplateFunc(t),
		SimulateTemplate:      newIndicesSimulateTemplateFunc(t),
		Split:                 newIndicesSplitFunc(t),
		Stats:                 newIndicesStatsFunc(t),
		Unfreeze:              newIndicesUnfreezeFunc(t),
		UpdateAliases:         newIndicesUpdateAliasesFunc(t),
		Upgrade:               newIndicesUpgradeFunc(t),
		ValidateQuery:         newIndicesValidateQueryFunc(t),
	}
}

// Ingest contains the Ingest APIs
type Ingest struct {
	DeletePipeline IngestDeletePipeline
	GetPipeline    IngestGetPipeline
	ProcessorGrok  IngestProcessorGrok
	PutPipeline    IngestPutPipeline
	Simulate       IngestSimulate
}

// NewIngest creates a new API client for Ingest APIs
//
func NewIngest(t Transport) *Ingest {
	return &Ingest{
		DeletePipeline: newIngestDeletePipelineFunc(t),
		GetPipeline:    newIngestGetPipelineFunc(t),
		ProcessorGrok:  newIngestProcessorGrokFunc(t),
		PutPipeline:    newIngestPutPipelineFunc(t),
		Simulate:       newIngestSimulateFunc(t),
	}
}

// Nodes contains the Nodes APIs
type Nodes struct {
	HotThreads           NodesHotThreads
	Info                 NodesInfo
	ReloadSecureSettings NodesReloadSecureSettings
	Stats                NodesStats
	Usage                NodesUsage
}

// NewNodes creates a new API client for Nodes APIs
//
func NewNodes(t Transport) *Nodes {
	return &Nodes{
		HotThreads:           newNodesHotThreadsFunc(t),
		Info:                 newNodesInfoFunc(t),
		ReloadSecureSettings: newNodesReloadSecureSettingsFunc(t),
		Stats:                newNodesStatsFunc(t),
		Usage:                newNodesUsageFunc(t),
	}
}

// Remote contains the Remote APIs
type Remote struct {
}

// NewRemote creates a new API client for Remote APIs
//
func NewRemote(t Transport) *Remote {
	return &Remote{}
}

// Snapshot contains the Snapshot APIs
type Snapshot struct {
	CleanupRepository SnapshotCleanupRepository
	CreateRepository  SnapshotCreateRepository
	Create            SnapshotCreate
	DeleteRepository  SnapshotDeleteRepository
	Delete            SnapshotDelete
	GetRepository     SnapshotGetRepository
	Get               SnapshotGet
	Restore           SnapshotRestore
	Status            SnapshotStatus
	VerifyRepository  SnapshotVerifyRepository
}

// NewSnapshot creates a new API client for Snapshot APIs
//
func NewSnapshot(t Transport) *Snapshot {
	return &Snapshot{
		CleanupRepository: newSnapshotCleanupRepositoryFunc(t),
		CreateRepository:  newSnapshotCreateRepositoryFunc(t),
		Create:            newSnapshotCreateFunc(t),
		DeleteRepository:  newSnapshotDeleteRepositoryFunc(t),
		Delete:            newSnapshotDeleteFunc(t),
		GetRepository:     newSnapshotGetRepositoryFunc(t),
		Get:               newSnapshotGetFunc(t),
		Restore:           newSnapshotRestoreFunc(t),
		Status:            newSnapshotStatusFunc(t),
		VerifyRepository:  newSnapshotVerifyRepositoryFunc(t),
	}
}

// Tasks contains the Tasks APIs
type Tasks struct {
	Cancel TasksCancel
	Get    TasksGet
	List   TasksList
}

// NewTasks creates a new API client for Tasks APIs
//
func NewTasks(t Transport) *Tasks {
	return &Tasks{
		Cancel: newTasksCancelFunc(t),
		Get:    newTasksGetFunc(t),
		List:   newTasksListFunc(t),
	}
}

// AsyncSearch contains the AsyncSearch APIs
type AsyncSearch struct {
	Delete AsyncSearchDelete
	Get    AsyncSearchGet
	Submit AsyncSearchSubmit
}

// NewAsyncSearch creates a new API client for AsyncSearch APIs
//
func NewAsyncSearch(t Transport) *AsyncSearch {
	return &AsyncSearch{
		Delete: newAsyncSearchDeleteFunc(t),
		Get:    newAsyncSearchGetFunc(t),
		Submit: newAsyncSearchSubmitFunc(t),
	}
}

// CCR contains the CCR APIs
type CCR struct {
	DeleteAutoFollowPattern CCRDeleteAutoFollowPattern
	FollowInfo              CCRFollowInfo
	Follow                  CCRFollow
	FollowStats             CCRFollowStats
	ForgetFollower          CCRForgetFollower
	GetAutoFollowPattern    CCRGetAutoFollowPattern
	PauseAutoFollowPattern  CCRPauseAutoFollowPattern
	PauseFollow             CCRPauseFollow
	PutAutoFollowPattern    CCRPutAutoFollowPattern
	ResumeAutoFollowPattern CCRResumeAutoFollowPattern
	ResumeFollow            CCRResumeFollow
	Stats                   CCRStats
	Unfollow                CCRUnfollow
}

// NewCCR creates a new API client for CCR APIs
//
func NewCCR(t Transport) *CCR {
	return &CCR{
		DeleteAutoFollowPattern: newCCRDeleteAutoFollowPatternFunc(t),
		FollowInfo:              newCCRFollowInfoFunc(t),
		Follow:                  newCCRFollowFunc(t),
		FollowStats:             newCCRFollowStatsFunc(t),
		ForgetFollower:          newCCRForgetFollowerFunc(t),
		GetAutoFollowPattern:    newCCRGetAutoFollowPatternFunc(t),
		PauseAutoFollowPattern:  newCCRPauseAutoFollowPatternFunc(t),
		PauseFollow:             newCCRPauseFollowFunc(t),
		PutAutoFollowPattern:    newCCRPutAutoFollowPatternFunc(t),
		ResumeAutoFollowPattern: newCCRResumeAutoFollowPatternFunc(t),
		ResumeFollow:            newCCRResumeFollowFunc(t),
		Stats:                   newCCRStatsFunc(t),
		Unfollow:                newCCRUnfollowFunc(t),
	}
}

// ILM contains the ILM APIs
type ILM struct {
	DeleteLifecycle  ILMDeleteLifecycle
	ExplainLifecycle ILMExplainLifecycle
	GetLifecycle     ILMGetLifecycle
	GetStatus        ILMGetStatus
	MoveToStep       ILMMoveToStep
	PutLifecycle     ILMPutLifecycle
	RemovePolicy     ILMRemovePolicy
	Retry            ILMRetry
	Start            ILMStart
	Stop             ILMStop
}

// NewILM creates a new API client for ILM APIs
//
func NewILM(t Transport) *ILM {
	return &ILM{
		DeleteLifecycle:  newILMDeleteLifecycleFunc(t),
		ExplainLifecycle: newILMExplainLifecycleFunc(t),
		GetLifecycle:     newILMGetLifecycleFunc(t),
		GetStatus:        newILMGetStatusFunc(t),
		MoveToStep:       newILMMoveToStepFunc(t),
		PutLifecycle:     newILMPutLifecycleFunc(t),
		RemovePolicy:     newILMRemovePolicyFunc(t),
		Retry:            newILMRetryFunc(t),
		Start:            newILMStartFunc(t),
		Stop:             newILMStopFunc(t),
	}
}

// License contains the License APIs
type License struct {
	Delete         LicenseDelete
	GetBasicStatus LicenseGetBasicStatus
	Get            LicenseGet
	GetTrialStatus LicenseGetTrialStatus
	Post           LicensePost
	PostStartBasic LicensePostStartBasic
	PostStartTrial LicensePostStartTrial
}

// NewLicense creates a new API client for License APIs
//
func NewLicense(t Transport) *License {
	return &License{
		Delete:         newLicenseDeleteFunc(t),
		GetBasicStatus: newLicenseGetBasicStatusFunc(t),
		Get:            newLicenseGetFunc(t),
		GetTrialStatus: newLicenseGetTrialStatusFunc(t),
		Post:           newLicensePostFunc(t),
		PostStartBasic: newLicensePostStartBasicFunc(t),
		PostStartTrial: newLicensePostStartTrialFunc(t),
	}
}

// Migration contains the Migration APIs
type Migration struct {
	Deprecations MigrationDeprecations
}

// NewMigration creates a new API client for Migration APIs
//
func NewMigration(t Transport) *Migration {
	return &Migration{
		Deprecations: newMigrationDeprecationsFunc(t),
	}
}

// ML contains the ML APIs
type ML struct {
	CloseJob                   MLCloseJob
	DeleteCalendarEvent        MLDeleteCalendarEvent
	DeleteCalendarJob          MLDeleteCalendarJob
	DeleteCalendar             MLDeleteCalendar
	DeleteDataFrameAnalytics   MLDeleteDataFrameAnalytics
	DeleteDatafeed             MLDeleteDatafeed
	DeleteExpiredData          MLDeleteExpiredData
	DeleteFilter               MLDeleteFilter
	DeleteForecast             MLDeleteForecast
	DeleteJob                  MLDeleteJob
	DeleteModelSnapshot        MLDeleteModelSnapshot
	DeleteTrainedModel         MLDeleteTrainedModel
	EstimateModelMemory        MLEstimateModelMemory
	EvaluateDataFrame          MLEvaluateDataFrame
	ExplainDataFrameAnalytics  MLExplainDataFrameAnalytics
	FindFileStructure          MLFindFileStructure
	FlushJob                   MLFlushJob
	Forecast                   MLForecast
	GetBuckets                 MLGetBuckets
	GetCalendarEvents          MLGetCalendarEvents
	GetCalendars               MLGetCalendars
	GetCategories              MLGetCategories
	GetDataFrameAnalytics      MLGetDataFrameAnalytics
	GetDataFrameAnalyticsStats MLGetDataFrameAnalyticsStats
	GetDatafeedStats           MLGetDatafeedStats
	GetDatafeeds               MLGetDatafeeds
	GetFilters                 MLGetFilters
	GetInfluencers             MLGetInfluencers
	GetJobStats                MLGetJobStats
	GetJobs                    MLGetJobs
	GetModelSnapshots          MLGetModelSnapshots
	GetOverallBuckets          MLGetOverallBuckets
	GetRecords                 MLGetRecords
	GetTrainedModels           MLGetTrainedModels
	GetTrainedModelsStats      MLGetTrainedModelsStats
	Info                       MLInfo
	OpenJob                    MLOpenJob
	PostCalendarEvents         MLPostCalendarEvents
	PostData                   MLPostData
	PreviewDatafeed            MLPreviewDatafeed
	PutCalendarJob             MLPutCalendarJob
	PutCalendar                MLPutCalendar
	PutDataFrameAnalytics      MLPutDataFrameAnalytics
	PutDatafeed                MLPutDatafeed
	PutFilter                  MLPutFilter
	PutJob                     MLPutJob
	PutTrainedModel            MLPutTrainedModel
	RevertModelSnapshot        MLRevertModelSnapshot
	SetUpgradeMode             MLSetUpgradeMode
	StartDataFrameAnalytics    MLStartDataFrameAnalytics
	StartDatafeed              MLStartDatafeed
	StopDataFrameAnalytics     MLStopDataFrameAnalytics
	StopDatafeed               MLStopDatafeed
	UpdateDataFrameAnalytics   MLUpdateDataFrameAnalytics
	UpdateDatafeed             MLUpdateDatafeed
	UpdateFilter               MLUpdateFilter
	UpdateJob                  MLUpdateJob
	UpdateModelSnapshot        MLUpdateModelSnapshot
	ValidateDetector           MLValidateDetector
	Validate                   MLValidate
}

// NewML creates a new API client for ML APIs
//
func NewML(t Transport) *ML {
	return &ML{
		CloseJob:                   newMLCloseJobFunc(t),
		DeleteCalendarEvent:        newMLDeleteCalendarEventFunc(t),
		DeleteCalendarJob:          newMLDeleteCalendarJobFunc(t),
		DeleteCalendar:             newMLDeleteCalendarFunc(t),
		DeleteDataFrameAnalytics:   newMLDeleteDataFrameAnalyticsFunc(t),
		DeleteDatafeed:             newMLDeleteDatafeedFunc(t),
		DeleteExpiredData:          newMLDeleteExpiredDataFunc(t),
		DeleteFilter:               newMLDeleteFilterFunc(t),
		DeleteForecast:             newMLDeleteForecastFunc(t),
		DeleteJob:                  newMLDeleteJobFunc(t),
		DeleteModelSnapshot:        newMLDeleteModelSnapshotFunc(t),
		DeleteTrainedModel:         newMLDeleteTrainedModelFunc(t),
		EstimateModelMemory:        newMLEstimateModelMemoryFunc(t),
		EvaluateDataFrame:          newMLEvaluateDataFrameFunc(t),
		ExplainDataFrameAnalytics:  newMLExplainDataFrameAnalyticsFunc(t),
		FindFileStructure:          newMLFindFileStructureFunc(t),
		FlushJob:                   newMLFlushJobFunc(t),
		Forecast:                   newMLForecastFunc(t),
		GetBuckets:                 newMLGetBucketsFunc(t),
		GetCalendarEvents:          newMLGetCalendarEventsFunc(t),
		GetCalendars:               newMLGetCalendarsFunc(t),
		GetCategories:              newMLGetCategoriesFunc(t),
		GetDataFrameAnalytics:      newMLGetDataFrameAnalyticsFunc(t),
		GetDataFrameAnalyticsStats: newMLGetDataFrameAnalyticsStatsFunc(t),
		GetDatafeedStats:           newMLGetDatafeedStatsFunc(t),
		GetDatafeeds:               newMLGetDatafeedsFunc(t),
		GetFilters:                 newMLGetFiltersFunc(t),
		GetInfluencers:             newMLGetInfluencersFunc(t),
		GetJobStats:                newMLGetJobStatsFunc(t),
		GetJobs:                    newMLGetJobsFunc(t),
		GetModelSnapshots:          newMLGetModelSnapshotsFunc(t),
		GetOverallBuckets:          newMLGetOverallBucketsFunc(t),
		GetRecords:                 newMLGetRecordsFunc(t),
		GetTrainedModels:           newMLGetTrainedModelsFunc(t),
		GetTrainedModelsStats:      newMLGetTrainedModelsStatsFunc(t),
		Info:                       newMLInfoFunc(t),
		OpenJob:                    newMLOpenJobFunc(t),
		PostCalendarEvents:         newMLPostCalendarEventsFunc(t),
		PostData:                   newMLPostDataFunc(t),
		PreviewDatafeed:            newMLPreviewDatafeedFunc(t),
		PutCalendarJob:             newMLPutCalendarJobFunc(t),
		PutCalendar:                newMLPutCalendarFunc(t),
		PutDataFrameAnalytics:      newMLPutDataFrameAnalyticsFunc(t),
		PutDatafeed:                newMLPutDatafeedFunc(t),
		PutFilter:                  newMLPutFilterFunc(t),
		PutJob:                     newMLPutJobFunc(t),
		PutTrainedModel:            newMLPutTrainedModelFunc(t),
		RevertModelSnapshot:        newMLRevertModelSnapshotFunc(t),
		SetUpgradeMode:             newMLSetUpgradeModeFunc(t),
		StartDataFrameAnalytics:    newMLStartDataFrameAnalyticsFunc(t),
		StartDatafeed:              newMLStartDatafeedFunc(t),
		StopDataFrameAnalytics:     newMLStopDataFrameAnalyticsFunc(t),
		StopDatafeed:               newMLStopDatafeedFunc(t),
		UpdateDataFrameAnalytics:   newMLUpdateDataFrameAnalyticsFunc(t),
		UpdateDatafeed:             newMLUpdateDatafeedFunc(t),
		UpdateFilter:               newMLUpdateFilterFunc(t),
		UpdateJob:                  newMLUpdateJobFunc(t),
		UpdateModelSnapshot:        newMLUpdateModelSnapshotFunc(t),
		ValidateDetector:           newMLValidateDetectorFunc(t),
		Validate:                   newMLValidateFunc(t),
	}
}

// Monitoring contains the Monitoring APIs
type Monitoring struct {
	Bulk MonitoringBulk
}

// NewMonitoring creates a new API client for Monitoring APIs
//
func NewMonitoring(t Transport) *Monitoring {
	return &Monitoring{
		Bulk: newMonitoringBulkFunc(t),
	}
}

// Rollup contains the Rollup APIs
type Rollup struct {
	DeleteJob    RollupDeleteJob
	GetJobs      RollupGetJobs
	GetCaps      RollupGetRollupCaps
	GetIndexCaps RollupGetRollupIndexCaps
	PutJob       RollupPutJob
	Search       RollupRollupSearch
	StartJob     RollupStartJob
	StopJob      RollupStopJob
}

// NewRollup creates a new API client for Rollup APIs
//
func NewRollup(t Transport) *Rollup {
	return &Rollup{
		DeleteJob:    newRollupDeleteJobFunc(t),
		GetJobs:      newRollupGetJobsFunc(t),
		GetCaps:      newRollupGetRollupCapsFunc(t),
		GetIndexCaps: newRollupGetRollupIndexCapsFunc(t),
		PutJob:       newRollupPutJobFunc(t),
		Search:       newRollupRollupSearchFunc(t),
		StartJob:     newRollupStartJobFunc(t),
		StopJob:      newRollupStopJobFunc(t),
	}
}

// Security contains the Security APIs
type Security struct {
	Authenticate          SecurityAuthenticate
	ChangePassword        SecurityChangePassword
	ClearCachedPrivileges SecurityClearCachedPrivileges
	ClearCachedRealms     SecurityClearCachedRealms
	ClearCachedRoles      SecurityClearCachedRoles
	CreateAPIKey          SecurityCreateAPIKey
	DeletePrivileges      SecurityDeletePrivileges
	DeleteRoleMapping     SecurityDeleteRoleMapping
	DeleteRole            SecurityDeleteRole
	DeleteUser            SecurityDeleteUser
	DisableUser           SecurityDisableUser
	EnableUser            SecurityEnableUser
	GetAPIKey             SecurityGetAPIKey
	GetBuiltinPrivileges  SecurityGetBuiltinPrivileges
	GetPrivileges         SecurityGetPrivileges
	GetRoleMapping        SecurityGetRoleMapping
	GetRole               SecurityGetRole
	GetToken              SecurityGetToken
	GetUserPrivileges     SecurityGetUserPrivileges
	GetUser               SecurityGetUser
	HasPrivileges         SecurityHasPrivileges
	InvalidateAPIKey      SecurityInvalidateAPIKey
	InvalidateToken       SecurityInvalidateToken
	PutPrivileges         SecurityPutPrivileges
	PutRoleMapping        SecurityPutRoleMapping
	PutRole               SecurityPutRole
	PutUser               SecurityPutUser
}

// NewSecurity creates a new API client for Security APIs
//
func NewSecurity(t Transport) *Security {
	return &Security{
		Authenticate:          newSecurityAuthenticateFunc(t),
		ChangePassword:        newSecurityChangePasswordFunc(t),
		ClearCachedPrivileges: newSecurityClearCachedPrivilegesFunc(t),
		ClearCachedRealms:     newSecurityClearCachedRealmsFunc(t),
		ClearCachedRoles:      newSecurityClearCachedRolesFunc(t),
		CreateAPIKey:          newSecurityCreateAPIKeyFunc(t),
		DeletePrivileges:      newSecurityDeletePrivilegesFunc(t),
		DeleteRoleMapping:     newSecurityDeleteRoleMappingFunc(t),
		DeleteRole:            newSecurityDeleteRoleFunc(t),
		DeleteUser:            newSecurityDeleteUserFunc(t),
		DisableUser:           newSecurityDisableUserFunc(t),
		EnableUser:            newSecurityEnableUserFunc(t),
		GetAPIKey:             newSecurityGetAPIKeyFunc(t),
		GetBuiltinPrivileges:  newSecurityGetBuiltinPrivilegesFunc(t),
		GetPrivileges:         newSecurityGetPrivilegesFunc(t),
		GetRoleMapping:        newSecurityGetRoleMappingFunc(t),
		GetRole:               newSecurityGetRoleFunc(t),
		GetToken:              newSecurityGetTokenFunc(t),
		GetUserPrivileges:     newSecurityGetUserPrivilegesFunc(t),
		GetUser:               newSecurityGetUserFunc(t),
		HasPrivileges:         newSecurityHasPrivilegesFunc(t),
		InvalidateAPIKey:      newSecurityInvalidateAPIKeyFunc(t),
		InvalidateToken:       newSecurityInvalidateTokenFunc(t),
		PutPrivileges:         newSecurityPutPrivilegesFunc(t),
		PutRoleMapping:        newSecurityPutRoleMappingFunc(t),
		PutRole:               newSecurityPutRoleFunc(t),
		PutUser:               newSecurityPutUserFunc(t),
	}
}

// SQL contains the SQL APIs
type SQL struct {
	ClearCursor SQLClearCursor
	Query       SQLQuery
	Translate   SQLTranslate
}

// NewSQL creates a new API client for SQL APIs
//
func NewSQL(t Transport) *SQL {
	return &SQL{
		ClearCursor: newSQLClearCursorFunc(t),
		Query:       newSQLQueryFunc(t),
		Translate:   newSQLTranslateFunc(t),
	}
}

// SSL contains the SSL APIs
type SSL struct {
	Certificates SSLCertificates
}

// NewSSL creates a new API client for SSL APIs
//
func NewSSL(t Transport) *SSL {
	return &SSL{
		Certificates: newSSLCertificatesFunc(t),
	}
}

// Watcher contains the Watcher APIs
type Watcher struct {
	AckWatch        WatcherAckWatch
	ActivateWatch   WatcherActivateWatch
	DeactivateWatch WatcherDeactivateWatch
	DeleteWatch     WatcherDeleteWatch
	ExecuteWatch    WatcherExecuteWatch
	GetWatch        WatcherGetWatch
	PutWatch        WatcherPutWatch
	Start           WatcherStart
	Stats           WatcherStats
	Stop            WatcherStop
}

// NewWatcher creates a new API client for Watcher APIs
//
func NewWatcher(t Transport) *Watcher {
	return &Watcher{
		AckWatch:        newWatcherAckWatchFunc(t),
		ActivateWatch:   newWatcherActivateWatchFunc(t),
		DeactivateWatch: newWatcherDeactivateWatchFunc(t),
		DeleteWatch:     newWatcherDeleteWatchFunc(t),
		ExecuteWatch:    newWatcherExecuteWatchFunc(t),
		GetWatch:        newWatcherGetWatchFunc(t),
		PutWatch:        newWatcherPutWatchFunc(t),
		Start:           newWatcherStartFunc(t),
		Stats:           newWatcherStatsFunc(t),
		Stop:            newWatcherStopFunc(t),
	}
}

// XPack contains the XPack APIs
type XPack struct {
	Info  XPackInfo
	Usage XPackUsage
}

// NewXPack creates a new API client for XPack APIs
//
func NewXPack(t Transport) *XPack {
	return &XPack{
		Info:  newXPackInfoFunc(t),
		Usage: newXPackUsageFunc(t),
	}
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
		AutoscalingDeleteAutoscalingPolicy:            newAutoscalingDeleteAutoscalingPolicyFunc(t),
		AutoscalingGetAutoscalingDecision:             newAutoscalingGetAutoscalingDecisionFunc(t),
		AutoscalingGetAutoscalingPolicy:               newAutoscalingGetAutoscalingPolicyFunc(t),
		AutoscalingPutAutoscalingPolicy:               newAutoscalingPutAutoscalingPolicyFunc(t),
		Bulk:                                          newBulkFunc(t),
		ClearScroll:                                   newClearScrollFunc(t),
		Count:                                         newCountFunc(t),
		Create:                                        newCreateFunc(t),
		DanglingIndicesDeleteDanglingIndex:            newDanglingIndicesDeleteDanglingIndexFunc(t),
		DanglingIndicesImportDanglingIndex:            newDanglingIndicesImportDanglingIndexFunc(t),
		DanglingIndicesListDanglingIndices:            newDanglingIndicesListDanglingIndicesFunc(t),
		DataFrameTransformDeprecatedDeleteTransform:   newDataFrameTransformDeprecatedDeleteTransformFunc(t),
		DataFrameTransformDeprecatedGetTransform:      newDataFrameTransformDeprecatedGetTransformFunc(t),
		DataFrameTransformDeprecatedGetTransformStats: newDataFrameTransformDeprecatedGetTransformStatsFunc(t),
		DataFrameTransformDeprecatedPreviewTransform:  newDataFrameTransformDeprecatedPreviewTransformFunc(t),
		DataFrameTransformDeprecatedPutTransform:      newDataFrameTransformDeprecatedPutTransformFunc(t),
		DataFrameTransformDeprecatedStartTransform:    newDataFrameTransformDeprecatedStartTransformFunc(t),
		DataFrameTransformDeprecatedStopTransform:     newDataFrameTransformDeprecatedStopTransformFunc(t),
		DataFrameTransformDeprecatedUpdateTransform:   newDataFrameTransformDeprecatedUpdateTransformFunc(t),
		DeleteByQuery:                                 newDeleteByQueryFunc(t),
		DeleteByQueryRethrottle:                       newDeleteByQueryRethrottleFunc(t),
		Delete:                                        newDeleteFunc(t),
		DeleteScript:                                  newDeleteScriptFunc(t),
		EnrichDeletePolicy:                            newEnrichDeletePolicyFunc(t),
		EnrichExecutePolicy:                           newEnrichExecutePolicyFunc(t),
		EnrichGetPolicy:                               newEnrichGetPolicyFunc(t),
		EnrichPutPolicy:                               newEnrichPutPolicyFunc(t),
		EnrichStats:                                   newEnrichStatsFunc(t),
		EqlDelete:                                     newEqlDeleteFunc(t),
		EqlGet:                                        newEqlGetFunc(t),
		EqlSearch:                                     newEqlSearchFunc(t),
		Exists:                                        newExistsFunc(t),
		ExistsSource:                                  newExistsSourceFunc(t),
		Explain:                                       newExplainFunc(t),
		FieldCaps:                                     newFieldCapsFunc(t),
		Get:                                           newGetFunc(t),
		GetScriptContext:                              newGetScriptContextFunc(t),
		GetScriptLanguages:                            newGetScriptLanguagesFunc(t),
		GetScript:                                     newGetScriptFunc(t),
		GetSource:                                     newGetSourceFunc(t),
		GraphExplore:                                  newGraphExploreFunc(t),
		Index:                                         newIndexFunc(t),
		Info:                                          newInfoFunc(t),
		Mget:                                          newMgetFunc(t),
		Msearch:                                       newMsearchFunc(t),
		MsearchTemplate:                               newMsearchTemplateFunc(t),
		Mtermvectors:                                  newMtermvectorsFunc(t),
		Ping:                                          newPingFunc(t),
		PutScript:                                     newPutScriptFunc(t),
		RankEval:                                      newRankEvalFunc(t),
		Reindex:                                       newReindexFunc(t),
		ReindexRethrottle:                             newReindexRethrottleFunc(t),
		RenderSearchTemplate:                          newRenderSearchTemplateFunc(t),
		ScriptsPainlessExecute:                        newScriptsPainlessExecuteFunc(t),
		Scroll:                                        newScrollFunc(t),
		Search:                                        newSearchFunc(t),
		SearchShards:                                  newSearchShardsFunc(t),
		SearchTemplate:                                newSearchTemplateFunc(t),
		SearchableSnapshotsClearCache:                 newSearchableSnapshotsClearCacheFunc(t),
		SearchableSnapshotsMount:                      newSearchableSnapshotsMountFunc(t),
		SearchableSnapshotsRepositoryStats:            newSearchableSnapshotsRepositoryStatsFunc(t),
		SearchableSnapshotsStats:                      newSearchableSnapshotsStatsFunc(t),
		SlmDeleteLifecycle:                            newSlmDeleteLifecycleFunc(t),
		SlmExecuteLifecycle:                           newSlmExecuteLifecycleFunc(t),
		SlmExecuteRetention:                           newSlmExecuteRetentionFunc(t),
		SlmGetLifecycle:                               newSlmGetLifecycleFunc(t),
		SlmGetStats:                                   newSlmGetStatsFunc(t),
		SlmGetStatus:                                  newSlmGetStatusFunc(t),
		SlmPutLifecycle:                               newSlmPutLifecycleFunc(t),
		SlmStart:                                      newSlmStartFunc(t),
		SlmStop:                                       newSlmStopFunc(t),
		Termvectors:                                   newTermvectorsFunc(t),
		TransformDeleteTransform:                      newTransformDeleteTransformFunc(t),
		TransformGetTransform:                         newTransformGetTransformFunc(t),
		TransformGetTransformStats:                    newTransformGetTransformStatsFunc(t),
		TransformPreviewTransform:                     newTransformPreviewTransformFunc(t),
		TransformPutTransform:                         newTransformPutTransformFunc(t),
		TransformStartTransform:                       newTransformStartTransformFunc(t),
		TransformStopTransform:                        newTransformStopTransformFunc(t),
		TransformUpdateTransform:                      newTransformUpdateTransformFunc(t),
		UpdateByQuery:                                 newUpdateByQueryFunc(t),
		UpdateByQueryRethrottle:                       newUpdateByQueryRethrottleFunc(t),
		Update:                                        newUpdateFunc(t),
		Cat:                                           NewCat(t),
		Cluster:                                       NewCluster(t),
		Indices:                                       NewIndices(t),
		Ingest:                                        NewIngest(t),
		Nodes:                                         NewNodes(t),
		Remote:                                        NewRemote(t),
		Snapshot:                                      NewSnapshot(t),
		Tasks:                                         NewTasks(t),
		AsyncSearch:                                   NewAsyncSearch(t),
		CCR:                                           NewCCR(t),
		ILM:                                           NewILM(t),
		License:                                       NewLicense(t),
		Migration:                                     NewMigration(t),
		ML:                                            NewML(t),
		Monitoring:                                    NewMonitoring(t),
		Rollup:                                        NewRollup(t),
		Security:                                      NewSecurity(t),
		SQL:                                           NewSQL(t),
		SSL:                                           NewSSL(t),
		Watcher:                                       NewWatcher(t),
		XPack:                                         NewXPack(t),
	}
}

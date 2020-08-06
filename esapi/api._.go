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
		Aliases:              NewCatAliases(t),
		Allocation:           NewCatAllocation(t),
		Count:                NewCatCount(t),
		Fielddata:            NewCatFielddata(t),
		Health:               NewCatHealth(t),
		Help:                 NewCatHelp(t),
		Indices:              NewCatIndices(t),
		MLDataFrameAnalytics: NewCatMLDataFrameAnalytics(t),
		MLDatafeeds:          NewCatMLDatafeeds(t),
		MLJobs:               NewCatMLJobs(t),
		MLTrainedModels:      NewCatMLTrainedModels(t),
		Master:               NewCatMaster(t),
		Nodeattrs:            NewCatNodeattrs(t),
		Nodes:                NewCatNodes(t),
		PendingTasks:         NewCatPendingTasks(t),
		Plugins:              NewCatPlugins(t),
		Recovery:             NewCatRecovery(t),
		Repositories:         NewCatRepositories(t),
		Segments:             NewCatSegments(t),
		Shards:               NewCatShards(t),
		Snapshots:            NewCatSnapshots(t),
		Tasks:                NewCatTasks(t),
		Templates:            NewCatTemplates(t),
		ThreadPool:           NewCatThreadPool(t),
		Transforms:           NewCatTransforms(t),
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
		AllocationExplain:            NewClusterAllocationExplain(t),
		DeleteComponentTemplate:      NewClusterDeleteComponentTemplate(t),
		DeleteVotingConfigExclusions: NewClusterDeleteVotingConfigExclusions(t),
		ExistsComponentTemplate:      NewClusterExistsComponentTemplate(t),
		GetComponentTemplate:         NewClusterGetComponentTemplate(t),
		GetSettings:                  NewClusterGetSettings(t),
		Health:                       NewClusterHealth(t),
		PendingTasks:                 NewClusterPendingTasks(t),
		PostVotingConfigExclusions:   NewClusterPostVotingConfigExclusions(t),
		PutComponentTemplate:         NewClusterPutComponentTemplate(t),
		PutSettings:                  NewClusterPutSettings(t),
		RemoteInfo:                   NewClusterRemoteInfo(t),
		Reroute:                      NewClusterReroute(t),
		State:                        NewClusterState(t),
		Stats:                        NewClusterStats(t),
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
		AddBlock:              NewIndicesAddBlock(t),
		Analyze:               NewIndicesAnalyze(t),
		ClearCache:            NewIndicesClearCache(t),
		Clone:                 NewIndicesClone(t),
		Close:                 NewIndicesClose(t),
		CreateDataStream:      NewIndicesCreateDataStream(t),
		Create:                NewIndicesCreate(t),
		DataStreamsStats:      NewIndicesDataStreamsStats(t),
		DeleteAlias:           NewIndicesDeleteAlias(t),
		DeleteDataStream:      NewIndicesDeleteDataStream(t),
		DeleteIndexTemplate:   NewIndicesDeleteIndexTemplate(t),
		Delete:                NewIndicesDelete(t),
		DeleteTemplate:        NewIndicesDeleteTemplate(t),
		ExistsAlias:           NewIndicesExistsAlias(t),
		ExistsDocumentType:    NewIndicesExistsDocumentType(t),
		ExistsIndexTemplate:   NewIndicesExistsIndexTemplate(t),
		Exists:                NewIndicesExists(t),
		ExistsTemplate:        NewIndicesExistsTemplate(t),
		Flush:                 NewIndicesFlush(t),
		Forcemerge:            NewIndicesForcemerge(t),
		Freeze:                NewIndicesFreeze(t),
		GetAlias:              NewIndicesGetAlias(t),
		GetDataStream:         NewIndicesGetDataStream(t),
		GetFieldMapping:       NewIndicesGetFieldMapping(t),
		GetIndexTemplate:      NewIndicesGetIndexTemplate(t),
		GetMapping:            NewIndicesGetMapping(t),
		Get:                   NewIndicesGet(t),
		GetSettings:           NewIndicesGetSettings(t),
		GetTemplate:           NewIndicesGetTemplate(t),
		GetUpgrade:            NewIndicesGetUpgrade(t),
		Open:                  NewIndicesOpen(t),
		PutAlias:              NewIndicesPutAlias(t),
		PutIndexTemplate:      NewIndicesPutIndexTemplate(t),
		PutMapping:            NewIndicesPutMapping(t),
		PutSettings:           NewIndicesPutSettings(t),
		PutTemplate:           NewIndicesPutTemplate(t),
		Recovery:              NewIndicesRecovery(t),
		Refresh:               NewIndicesRefresh(t),
		ReloadSearchAnalyzers: NewIndicesReloadSearchAnalyzers(t),
		ResolveIndex:          NewIndicesResolveIndex(t),
		Rollover:              NewIndicesRollover(t),
		Segments:              NewIndicesSegments(t),
		ShardStores:           NewIndicesShardStores(t),
		Shrink:                NewIndicesShrink(t),
		SimulateIndexTemplate: NewIndicesSimulateIndexTemplate(t),
		SimulateTemplate:      NewIndicesSimulateTemplate(t),
		Split:                 NewIndicesSplit(t),
		Stats:                 NewIndicesStats(t),
		Unfreeze:              NewIndicesUnfreeze(t),
		UpdateAliases:         NewIndicesUpdateAliases(t),
		Upgrade:               NewIndicesUpgrade(t),
		ValidateQuery:         NewIndicesValidateQuery(t),
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
		DeletePipeline: NewIngestDeletePipeline(t),
		GetPipeline:    NewIngestGetPipeline(t),
		ProcessorGrok:  NewIngestProcessorGrok(t),
		PutPipeline:    NewIngestPutPipeline(t),
		Simulate:       NewIngestSimulate(t),
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
		HotThreads:           NewNodesHotThreads(t),
		Info:                 NewNodesInfo(t),
		ReloadSecureSettings: NewNodesReloadSecureSettings(t),
		Stats:                NewNodesStats(t),
		Usage:                NewNodesUsage(t),
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
		CleanupRepository: NewSnapshotCleanupRepository(t),
		CreateRepository:  NewSnapshotCreateRepository(t),
		Create:            NewSnapshotCreate(t),
		DeleteRepository:  NewSnapshotDeleteRepository(t),
		Delete:            NewSnapshotDelete(t),
		GetRepository:     NewSnapshotGetRepository(t),
		Get:               NewSnapshotGet(t),
		Restore:           NewSnapshotRestore(t),
		Status:            NewSnapshotStatus(t),
		VerifyRepository:  NewSnapshotVerifyRepository(t),
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
		Cancel: NewTasksCancel(t),
		Get:    NewTasksGet(t),
		List:   NewTasksList(t),
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
		Delete: NewAsyncSearchDelete(t),
		Get:    NewAsyncSearchGet(t),
		Submit: NewAsyncSearchSubmit(t),
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
		DeleteAutoFollowPattern: NewCCRDeleteAutoFollowPattern(t),
		FollowInfo:              NewCCRFollowInfo(t),
		Follow:                  NewCCRFollow(t),
		FollowStats:             NewCCRFollowStats(t),
		ForgetFollower:          NewCCRForgetFollower(t),
		GetAutoFollowPattern:    NewCCRGetAutoFollowPattern(t),
		PauseAutoFollowPattern:  NewCCRPauseAutoFollowPattern(t),
		PauseFollow:             NewCCRPauseFollow(t),
		PutAutoFollowPattern:    NewCCRPutAutoFollowPattern(t),
		ResumeAutoFollowPattern: NewCCRResumeAutoFollowPattern(t),
		ResumeFollow:            NewCCRResumeFollow(t),
		Stats:                   NewCCRStats(t),
		Unfollow:                NewCCRUnfollow(t),
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
		DeleteLifecycle:  NewILMDeleteLifecycle(t),
		ExplainLifecycle: NewILMExplainLifecycle(t),
		GetLifecycle:     NewILMGetLifecycle(t),
		GetStatus:        NewILMGetStatus(t),
		MoveToStep:       NewILMMoveToStep(t),
		PutLifecycle:     NewILMPutLifecycle(t),
		RemovePolicy:     NewILMRemovePolicy(t),
		Retry:            NewILMRetry(t),
		Start:            NewILMStart(t),
		Stop:             NewILMStop(t),
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
		Delete:         NewLicenseDelete(t),
		GetBasicStatus: NewLicenseGetBasicStatus(t),
		Get:            NewLicenseGet(t),
		GetTrialStatus: NewLicenseGetTrialStatus(t),
		Post:           NewLicensePost(t),
		PostStartBasic: NewLicensePostStartBasic(t),
		PostStartTrial: NewLicensePostStartTrial(t),
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
		Deprecations: NewMigrationDeprecations(t),
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
		CloseJob:                   NewMLCloseJob(t),
		DeleteCalendarEvent:        NewMLDeleteCalendarEvent(t),
		DeleteCalendarJob:          NewMLDeleteCalendarJob(t),
		DeleteCalendar:             NewMLDeleteCalendar(t),
		DeleteDataFrameAnalytics:   NewMLDeleteDataFrameAnalytics(t),
		DeleteDatafeed:             NewMLDeleteDatafeed(t),
		DeleteExpiredData:          NewMLDeleteExpiredData(t),
		DeleteFilter:               NewMLDeleteFilter(t),
		DeleteForecast:             NewMLDeleteForecast(t),
		DeleteJob:                  NewMLDeleteJob(t),
		DeleteModelSnapshot:        NewMLDeleteModelSnapshot(t),
		DeleteTrainedModel:         NewMLDeleteTrainedModel(t),
		EstimateModelMemory:        NewMLEstimateModelMemory(t),
		EvaluateDataFrame:          NewMLEvaluateDataFrame(t),
		ExplainDataFrameAnalytics:  NewMLExplainDataFrameAnalytics(t),
		FindFileStructure:          NewMLFindFileStructure(t),
		FlushJob:                   NewMLFlushJob(t),
		Forecast:                   NewMLForecast(t),
		GetBuckets:                 NewMLGetBuckets(t),
		GetCalendarEvents:          NewMLGetCalendarEvents(t),
		GetCalendars:               NewMLGetCalendars(t),
		GetCategories:              NewMLGetCategories(t),
		GetDataFrameAnalytics:      NewMLGetDataFrameAnalytics(t),
		GetDataFrameAnalyticsStats: NewMLGetDataFrameAnalyticsStats(t),
		GetDatafeedStats:           NewMLGetDatafeedStats(t),
		GetDatafeeds:               NewMLGetDatafeeds(t),
		GetFilters:                 NewMLGetFilters(t),
		GetInfluencers:             NewMLGetInfluencers(t),
		GetJobStats:                NewMLGetJobStats(t),
		GetJobs:                    NewMLGetJobs(t),
		GetModelSnapshots:          NewMLGetModelSnapshots(t),
		GetOverallBuckets:          NewMLGetOverallBuckets(t),
		GetRecords:                 NewMLGetRecords(t),
		GetTrainedModels:           NewMLGetTrainedModels(t),
		GetTrainedModelsStats:      NewMLGetTrainedModelsStats(t),
		Info:                       NewMLInfo(t),
		OpenJob:                    NewMLOpenJob(t),
		PostCalendarEvents:         NewMLPostCalendarEvents(t),
		PostData:                   NewMLPostData(t),
		PreviewDatafeed:            NewMLPreviewDatafeed(t),
		PutCalendarJob:             NewMLPutCalendarJob(t),
		PutCalendar:                NewMLPutCalendar(t),
		PutDataFrameAnalytics:      NewMLPutDataFrameAnalytics(t),
		PutDatafeed:                NewMLPutDatafeed(t),
		PutFilter:                  NewMLPutFilter(t),
		PutJob:                     NewMLPutJob(t),
		PutTrainedModel:            NewMLPutTrainedModel(t),
		RevertModelSnapshot:        NewMLRevertModelSnapshot(t),
		SetUpgradeMode:             NewMLSetUpgradeMode(t),
		StartDataFrameAnalytics:    NewMLStartDataFrameAnalytics(t),
		StartDatafeed:              NewMLStartDatafeed(t),
		StopDataFrameAnalytics:     NewMLStopDataFrameAnalytics(t),
		StopDatafeed:               NewMLStopDatafeed(t),
		UpdateDataFrameAnalytics:   NewMLUpdateDataFrameAnalytics(t),
		UpdateDatafeed:             NewMLUpdateDatafeed(t),
		UpdateFilter:               NewMLUpdateFilter(t),
		UpdateJob:                  NewMLUpdateJob(t),
		UpdateModelSnapshot:        NewMLUpdateModelSnapshot(t),
		ValidateDetector:           NewMLValidateDetector(t),
		Validate:                   NewMLValidate(t),
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
		Bulk: NewMonitoringBulk(t),
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
		DeleteJob:    NewRollupDeleteJob(t),
		GetJobs:      NewRollupGetJobs(t),
		GetCaps:      NewRollupGetRollupCaps(t),
		GetIndexCaps: NewRollupGetRollupIndexCaps(t),
		PutJob:       NewRollupPutJob(t),
		Search:       NewRollupRollupSearch(t),
		StartJob:     NewRollupStartJob(t),
		StopJob:      NewRollupStopJob(t),
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
		Authenticate:          NewSecurityAuthenticate(t),
		ChangePassword:        NewSecurityChangePassword(t),
		ClearCachedPrivileges: NewSecurityClearCachedPrivileges(t),
		ClearCachedRealms:     NewSecurityClearCachedRealms(t),
		ClearCachedRoles:      NewSecurityClearCachedRoles(t),
		CreateAPIKey:          NewSecurityCreateAPIKey(t),
		DeletePrivileges:      NewSecurityDeletePrivileges(t),
		DeleteRoleMapping:     NewSecurityDeleteRoleMapping(t),
		DeleteRole:            NewSecurityDeleteRole(t),
		DeleteUser:            NewSecurityDeleteUser(t),
		DisableUser:           NewSecurityDisableUser(t),
		EnableUser:            NewSecurityEnableUser(t),
		GetAPIKey:             NewSecurityGetAPIKey(t),
		GetBuiltinPrivileges:  NewSecurityGetBuiltinPrivileges(t),
		GetPrivileges:         NewSecurityGetPrivileges(t),
		GetRoleMapping:        NewSecurityGetRoleMapping(t),
		GetRole:               NewSecurityGetRole(t),
		GetToken:              NewSecurityGetToken(t),
		GetUserPrivileges:     NewSecurityGetUserPrivileges(t),
		GetUser:               NewSecurityGetUser(t),
		HasPrivileges:         NewSecurityHasPrivileges(t),
		InvalidateAPIKey:      NewSecurityInvalidateAPIKey(t),
		InvalidateToken:       NewSecurityInvalidateToken(t),
		PutPrivileges:         NewSecurityPutPrivileges(t),
		PutRoleMapping:        NewSecurityPutRoleMapping(t),
		PutRole:               NewSecurityPutRole(t),
		PutUser:               NewSecurityPutUser(t),
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
		ClearCursor: NewSQLClearCursor(t),
		Query:       NewSQLQuery(t),
		Translate:   NewSQLTranslate(t),
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
		Certificates: NewSSLCertificates(t),
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
		AckWatch:        NewWatcherAckWatch(t),
		ActivateWatch:   NewWatcherActivateWatch(t),
		DeactivateWatch: NewWatcherDeactivateWatch(t),
		DeleteWatch:     NewWatcherDeleteWatch(t),
		ExecuteWatch:    NewWatcherExecuteWatch(t),
		GetWatch:        NewWatcherGetWatch(t),
		PutWatch:        NewWatcherPutWatch(t),
		Start:           NewWatcherStart(t),
		Stats:           NewWatcherStats(t),
		Stop:            NewWatcherStop(t),
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
		Info:  NewXPackInfo(t),
		Usage: NewXPackUsage(t),
	}
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
		AutoscalingDeleteAutoscalingPolicy:            NewAutoscalingDeleteAutoscalingPolicy(t),
		AutoscalingGetAutoscalingDecision:             NewAutoscalingGetAutoscalingDecision(t),
		AutoscalingGetAutoscalingPolicy:               NewAutoscalingGetAutoscalingPolicy(t),
		AutoscalingPutAutoscalingPolicy:               NewAutoscalingPutAutoscalingPolicy(t),
		Bulk:                                          NewBulk(t),
		ClearScroll:                                   NewClearScroll(t),
		Count:                                         NewCount(t),
		Create:                                        NewCreate(t),
		DanglingIndicesDeleteDanglingIndex:            NewDanglingIndicesDeleteDanglingIndex(t),
		DanglingIndicesImportDanglingIndex:            NewDanglingIndicesImportDanglingIndex(t),
		DanglingIndicesListDanglingIndices:            NewDanglingIndicesListDanglingIndices(t),
		DataFrameTransformDeprecatedDeleteTransform:   NewDataFrameTransformDeprecatedDeleteTransform(t),
		DataFrameTransformDeprecatedGetTransform:      NewDataFrameTransformDeprecatedGetTransform(t),
		DataFrameTransformDeprecatedGetTransformStats: NewDataFrameTransformDeprecatedGetTransformStats(t),
		DataFrameTransformDeprecatedPreviewTransform:  NewDataFrameTransformDeprecatedPreviewTransform(t),
		DataFrameTransformDeprecatedPutTransform:      NewDataFrameTransformDeprecatedPutTransform(t),
		DataFrameTransformDeprecatedStartTransform:    NewDataFrameTransformDeprecatedStartTransform(t),
		DataFrameTransformDeprecatedStopTransform:     NewDataFrameTransformDeprecatedStopTransform(t),
		DataFrameTransformDeprecatedUpdateTransform:   NewDataFrameTransformDeprecatedUpdateTransform(t),
		DeleteByQuery:                                 NewDeleteByQuery(t),
		DeleteByQueryRethrottle:                       NewDeleteByQueryRethrottle(t),
		Delete:                                        NewDelete(t),
		DeleteScript:                                  NewDeleteScript(t),
		EnrichDeletePolicy:                            NewEnrichDeletePolicy(t),
		EnrichExecutePolicy:                           NewEnrichExecutePolicy(t),
		EnrichGetPolicy:                               NewEnrichGetPolicy(t),
		EnrichPutPolicy:                               NewEnrichPutPolicy(t),
		EnrichStats:                                   NewEnrichStats(t),
		EqlDelete:                                     NewEqlDelete(t),
		EqlGet:                                        NewEqlGet(t),
		EqlSearch:                                     NewEqlSearch(t),
		Exists:                                        NewExists(t),
		ExistsSource:                                  NewExistsSource(t),
		Explain:                                       NewExplain(t),
		FieldCaps:                                     NewFieldCaps(t),
		Get:                                           NewGet(t),
		GetScriptContext:                              NewGetScriptContext(t),
		GetScriptLanguages:                            NewGetScriptLanguages(t),
		GetScript:                                     NewGetScript(t),
		GetSource:                                     NewGetSource(t),
		GraphExplore:                                  NewGraphExplore(t),
		Index:                                         NewIndex(t),
		Info:                                          NewInfo(t),
		Mget:                                          NewMget(t),
		Msearch:                                       NewMsearch(t),
		MsearchTemplate:                               NewMsearchTemplate(t),
		Mtermvectors:                                  NewMtermvectors(t),
		Ping:                                          NewPing(t),
		PutScript:                                     NewPutScript(t),
		RankEval:                                      NewRankEval(t),
		Reindex:                                       NewReindex(t),
		ReindexRethrottle:                             NewReindexRethrottle(t),
		RenderSearchTemplate:                          NewRenderSearchTemplate(t),
		ScriptsPainlessExecute:                        NewScriptsPainlessExecute(t),
		Scroll:                                        NewScroll(t),
		Search:                                        NewSearch(t),
		SearchShards:                                  NewSearchShards(t),
		SearchTemplate:                                NewSearchTemplate(t),
		SearchableSnapshotsClearCache:                 NewSearchableSnapshotsClearCache(t),
		SearchableSnapshotsMount:                      NewSearchableSnapshotsMount(t),
		SearchableSnapshotsRepositoryStats:            NewSearchableSnapshotsRepositoryStats(t),
		SearchableSnapshotsStats:                      NewSearchableSnapshotsStats(t),
		SlmDeleteLifecycle:                            NewSlmDeleteLifecycle(t),
		SlmExecuteLifecycle:                           NewSlmExecuteLifecycle(t),
		SlmExecuteRetention:                           NewSlmExecuteRetention(t),
		SlmGetLifecycle:                               NewSlmGetLifecycle(t),
		SlmGetStats:                                   NewSlmGetStats(t),
		SlmGetStatus:                                  NewSlmGetStatus(t),
		SlmPutLifecycle:                               NewSlmPutLifecycle(t),
		SlmStart:                                      NewSlmStart(t),
		SlmStop:                                       NewSlmStop(t),
		Termvectors:                                   NewTermvectors(t),
		TransformDeleteTransform:                      NewTransformDeleteTransform(t),
		TransformGetTransform:                         NewTransformGetTransform(t),
		TransformGetTransformStats:                    NewTransformGetTransformStats(t),
		TransformPreviewTransform:                     NewTransformPreviewTransform(t),
		TransformPutTransform:                         NewTransformPutTransform(t),
		TransformStartTransform:                       NewTransformStartTransform(t),
		TransformStopTransform:                        NewTransformStopTransform(t),
		TransformUpdateTransform:                      NewTransformUpdateTransform(t),
		UpdateByQuery:                                 NewUpdateByQuery(t),
		UpdateByQueryRethrottle:                       NewUpdateByQueryRethrottle(t),
		Update:                                        NewUpdate(t),
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

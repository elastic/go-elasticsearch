// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"time"

	"github.com/elastic/go-elasticsearch/transport"
)

// ExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type ExpandWildcards int

const (
	// ExpandWildcardsOpen can be used to set ExpandWildcards to "open"
	ExpandWildcardsOpen = iota
	// ExpandWildcardsClosed can be used to set ExpandWildcards to "closed"
	ExpandWildcardsClosed = iota
	// ExpandWildcardsNone can be used to set ExpandWildcards to "none"
	ExpandWildcardsNone = iota
	// ExpandWildcardsAll can be used to set ExpandWildcards to "all"
	ExpandWildcardsAll = iota
)

// Level - specify the level of detail for returned information.
type Level int

const (
	// LevelCluster can be used to set Level to "cluster"
	LevelCluster = iota
	// LevelIndices can be used to set Level to "indices"
	LevelIndices = iota
	// LevelShards can be used to set Level to "shards"
	LevelShards = iota
)

// WaitForEvents - wait until all currently queued events with the given priority are processed.
type WaitForEvents int

const (
	// WaitForEventsImmediate can be used to set WaitForEvents to "immediate"
	WaitForEventsImmediate = iota
	// WaitForEventsUrgent can be used to set WaitForEvents to "urgent"
	WaitForEventsUrgent = iota
	// WaitForEventsHigh can be used to set WaitForEvents to "high"
	WaitForEventsHigh = iota
	// WaitForEventsNormal can be used to set WaitForEvents to "normal"
	WaitForEventsNormal = iota
	// WaitForEventsLow can be used to set WaitForEvents to "low"
	WaitForEventsLow = iota
	// WaitForEventsLanguid can be used to set WaitForEvents to "languid"
	WaitForEventsLanguid = iota
)

// WaitForStatus - wait until cluster is in a specific state.
type WaitForStatus int

const (
	// WaitForStatusGreen can be used to set WaitForStatus to "green"
	WaitForStatusGreen = iota
	// WaitForStatusYellow can be used to set WaitForStatus to "yellow"
	WaitForStatusYellow = iota
	// WaitForStatusRed can be used to set WaitForStatus to "red"
	WaitForStatusRed = iota
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option struct {
	name  string
	apply func(r *transport.Request)
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithAllowNoIndices(allowNoIndices bool) Option {
	return Option{
		name: "WithAllowNoIndices",
		apply: func(r *transport.Request) {
		},
	}
}

// WithBody - the settings to be updated. Can be either "transient" or "persistent" (survives cluster restart).
func WithBody(body map[string]interface{}) Option {
	return Option{
		name: "WithBody",
		apply: func(r *transport.Request) {
		},
	}
}

// WithDryRun - simulate the operation only and return the resulting state.
func WithDryRun(dryRun bool) Option {
	return Option{
		name: "WithDryRun",
		apply: func(r *transport.Request) {
		},
	}
}

// WithErrorTrace - include the stack trace of returned errors.
func WithErrorTrace(errorTrace bool) Option {
	return Option{
		name: "WithErrorTrace",
		apply: func(r *transport.Request) {
		},
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithExpandWildcards(expandWildcards ExpandWildcards) Option {
	return Option{
		name: "WithExpandWildcards",
		apply: func(r *transport.Request) {
		},
	}
}

// WithExplain - return an explanation of why the commands can or cannot be executed.
func WithExplain(explain bool) Option {
	return Option{
		name: "WithExplain",
		apply: func(r *transport.Request) {
		},
	}
}

// WithFilterPath - a comma-separated list of filters used to reduce the respone.
func WithFilterPath(filterPath []string) Option {
	return Option{
		name: "WithFilterPath",
		apply: func(r *transport.Request) {
		},
	}
}

// WithFlatSettings - return settings in flat format (default: false).
func WithFlatSettings(flatSettings bool) Option {
	return Option{
		name: "WithFlatSettings",
		apply: func(r *transport.Request) {
		},
	}
}

// WithHuman - return human readable values for statistics.
func WithHuman(human bool) Option {
	return Option{
		name: "WithHuman",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIgnore - ignores the specified HTTP status codes.
func WithIgnore(ignore []int) Option {
	return Option{
		name: "WithIgnore",
		apply: func(r *transport.Request) {
			for _, status := range ignore {
				r.IgnoredStatuses[status] = struct{}{}
			}
		},
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithIgnoreUnavailable(ignoreUnavailable bool) Option {
	return Option{
		name: "WithIgnoreUnavailable",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIncludeDefaults - whether to return all default clusters setting.
func WithIncludeDefaults(includeDefaults bool) Option {
	return Option{
		name: "WithIncludeDefaults",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIncludeDiskInfo - return information about disk usage and shard sizes (default: false).
func WithIncludeDiskInfo(includeDiskInfo bool) Option {
	return Option{
		name: "WithIncludeDiskInfo",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIncludeYesDecisions - return 'YES' decisions in explanation (default: false).
func WithIncludeYesDecisions(includeYesDecisions bool) Option {
	return Option{
		name: "WithIncludeYesDecisions",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIndex - limit the information returned to a specific index.
func WithIndex(index []string) Option {
	return Option{
		name: "WithIndex",
		apply: func(r *transport.Request) {
		},
	}
}

// WithLevel - specify the level of detail for returned information.
func WithLevel(level Level) Option {
	return Option{
		name: "WithLevel",
		apply: func(r *transport.Request) {
		},
	}
}

// WithLocal - return local information, do not retrieve the state from master node (default: false).
func WithLocal(local bool) Option {
	return Option{
		name: "WithLocal",
		apply: func(r *transport.Request) {
		},
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func WithMasterTimeout(masterTimeout time.Duration) Option {
	return Option{
		name: "WithMasterTimeout",
		apply: func(r *transport.Request) {
		},
	}
}

// WithMetric - limit the information returned to the specified metrics.
func WithMetric(metric []string) Option {
	return Option{
		name: "WithMetric",
		apply: func(r *transport.Request) {
		},
	}
}

// WithNodeID - a comma-separated list of node IDs or names to limit the returned information; use "_local" to return information from the node you're connecting to, leave empty to get information from all nodes.
func WithNodeID(nodeID []string) Option {
	return Option{
		name: "WithNodeID",
		apply: func(r *transport.Request) {
		},
	}
}

// WithPretty - pretty format the returned JSON response.
func WithPretty(pretty bool) Option {
	return Option{
		name: "WithPretty",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSourceParam(sourceParam string) Option {
	return Option{
		name: "WithSourceParam",
		apply: func(r *transport.Request) {
		},
	}
}

// WithTimeout - explicit operation timeout.
func WithTimeout(timeout time.Duration) Option {
	return Option{
		name: "WithTimeout",
		apply: func(r *transport.Request) {
		},
	}
}

// WithWaitForActiveShards - wait until the specified number of shards is active.
func WithWaitForActiveShards(waitForActiveShards string) Option {
	return Option{
		name: "WithWaitForActiveShards",
		apply: func(r *transport.Request) {
		},
	}
}

// WithWaitForEvents - wait until all currently queued events with the given priority are processed.
func WithWaitForEvents(waitForEvents WaitForEvents) Option {
	return Option{
		name: "WithWaitForEvents",
		apply: func(r *transport.Request) {
		},
	}
}

// WithWaitForNoRelocatingShards - whether to wait until there are no relocating shards in the cluster.
func WithWaitForNoRelocatingShards(waitForNoRelocatingShards bool) Option {
	return Option{
		name: "WithWaitForNoRelocatingShards",
		apply: func(r *transport.Request) {
		},
	}
}

// WithWaitForNodes - wait until the specified number of nodes is available.
func WithWaitForNodes(waitForNodes string) Option {
	return Option{
		name: "WithWaitForNodes",
		apply: func(r *transport.Request) {
		},
	}
}

// WithWaitForStatus - wait until cluster is in a specific state.
func WithWaitForStatus(waitForStatus WaitForStatus) Option {
	return Option{
		name: "WithWaitForStatus",
		apply: func(r *transport.Request) {
		},
	}
}

var (
	supportedOptions = map[string]map[string]struct{}{
		"GetSettings": map[string]struct{}{
			"WithFlatSettings":    struct{}{},
			"WithIncludeDefaults": struct{}{},
			"WithMasterTimeout":   struct{}{},
			"WithTimeout":         struct{}{},
			"WithErrorTrace":      struct{}{},
			"WithFilterPath":      struct{}{},
			"WithHuman":           struct{}{},
			"WithIgnore":          struct{}{},
			"WithPretty":          struct{}{},
			"WithSourceParam":     struct{}{},
		},
		"State": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithMetric":            struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithFlatSettings":      struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithLocal":             struct{}{},
			"WithMasterTimeout":     struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"Health": map[string]struct{}{
			"WithIndex":                     struct{}{},
			"WithLevel":                     struct{}{},
			"WithLocal":                     struct{}{},
			"WithMasterTimeout":             struct{}{},
			"WithTimeout":                   struct{}{},
			"WithWaitForActiveShards":       struct{}{},
			"WithWaitForEvents":             struct{}{},
			"WithWaitForNoRelocatingShards": struct{}{},
			"WithWaitForNodes":              struct{}{},
			"WithWaitForStatus":             struct{}{},
			"WithErrorTrace":                struct{}{},
			"WithFilterPath":                struct{}{},
			"WithHuman":                     struct{}{},
			"WithIgnore":                    struct{}{},
			"WithPretty":                    struct{}{},
			"WithSourceParam":               struct{}{},
		},
		"Stats": map[string]struct{}{
			"WithNodeID":       struct{}{},
			"WithFlatSettings": struct{}{},
			"WithTimeout":      struct{}{},
			"WithErrorTrace":   struct{}{},
			"WithFilterPath":   struct{}{},
			"WithHuman":        struct{}{},
			"WithIgnore":       struct{}{},
			"WithPretty":       struct{}{},
			"WithSourceParam":  struct{}{},
		},
		"PendingTasks": map[string]struct{}{
			"WithLocal":         struct{}{},
			"WithMasterTimeout": struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
		"AllocationExplain": map[string]struct{}{
			"WithIncludeDiskInfo":     struct{}{},
			"WithIncludeYesDecisions": struct{}{},
			"WithBody":                struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"Reroute": map[string]struct{}{
			"WithDryRun":        struct{}{},
			"WithExplain":       struct{}{},
			"WithMasterTimeout": struct{}{},
			"WithMetric":        struct{}{},
			"WithRetryFailed":   struct{}{},
			"WithTimeout":       struct{}{},
			"WithBody":          struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
		"PutSettings": map[string]struct{}{
			"WithFlatSettings":  struct{}{},
			"WithMasterTimeout": struct{}{},
			"WithTimeout":       struct{}{},
			"WithBody":          struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
	}
)

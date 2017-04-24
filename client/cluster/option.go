// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"net/http"
	"time"
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option func(r *http.Request)

// WithDryRun simulate the operation only and return the resulting state.
func WithDryRun(dryRun bool) Option {
	return func(r *http.Request) {
	}
}

// WithExplain return an explanation of why the commands can or cannot be executed.
func WithExplain(explain bool) Option {
	return func(r *http.Request) {
	}
}

// WithFlatSettings return settings in flat format (default: false).
func WithFlatSettings(flatSettings bool) Option {
	return func(r *http.Request) {
	}
}

// WithIncludeDefaults whether to return all default clusters setting.
func WithIncludeDefaults(includeDefaults bool) Option {
	return func(r *http.Request) {
	}
}

// WithIncludeDiskInfo return information about disk usage and shard sizes (default: false).
func WithIncludeDiskInfo(includeDiskInfo bool) Option {
	return func(r *http.Request) {
	}
}

// WithIncludeYesDecisions return 'YES' decisions in explanation (default: false).
func WithIncludeYesDecisions(includeYesDecisions bool) Option {
	return func(r *http.Request) {
	}
}

// WithIndex limit the information returned to a specific index.
func WithIndex(index []string) Option {
	return func(r *http.Request) {
	}
}

// WithLevel specify the level of detail for returned information.
func WithLevel(level struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithLocal return local information, do not retrieve the state from master node (default: false).
func WithLocal(local bool) Option {
	return func(r *http.Request) {
	}
}

// WithMasterTimeout explicit operation timeout for connection to master node.
func WithMasterTimeout(masterTimeout time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithMetric limit the information returned to the specified metrics. Defaults to all but metadata.
func WithMetric(metric []string) Option {
	return func(r *http.Request) {
	}
}

// WithNodeID a comma-separated list of node IDs or names to limit the returned information; use _local to return information from the node you're connecting to, leave empty to get information from all nodes.
func WithNodeID(nodeID []string) Option {
	return func(r *http.Request) {
	}
}

// WithRetryFailed retries allocation of shards that are blocked due to too many subsequent allocation failures.
func WithRetryFailed(retryFailed bool) Option {
	return func(r *http.Request) {
	}
}

// WithTimeout explicit operation timeout.
func WithTimeout(timeout time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithWaitForActiveShards wait until the specified number of shards is active.
func WithWaitForActiveShards(waitForActiveShards string) Option {
	return func(r *http.Request) {
	}
}

// WithWaitForEvents wait until all currently queued events with the given priority are processed.
func WithWaitForEvents(waitForEvents struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithWaitForNoRelocatingShards whether to wait until there are no relocating shards in the cluster.
func WithWaitForNoRelocatingShards(waitForNoRelocatingShards bool) Option {
	return func(r *http.Request) {
	}
}

// WithWaitForNodes wait until the specified number of nodes is available.
func WithWaitForNodes(waitForNodes string) Option {
	return func(r *http.Request) {
	}
}

// WithWaitForStatus wait until cluster is in a specific state.
func WithWaitForStatus(waitForStatus struct{}) Option {
	return func(r *http.Request) {
	}
}

// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package nodes

import (
	"time"

	"github.com/elastic/go-elasticsearch/transport"
)

// DocumentType - the type to sample (default: cpu).
type DocumentType int

const (
	// DocumentTypeCPU can be used to set DocumentType to "cpu"
	DocumentTypeCPU = iota
	// DocumentTypeWait can be used to set DocumentType to "wait"
	DocumentTypeWait = iota
	// DocumentTypeBlock can be used to set DocumentType to "block"
	DocumentTypeBlock = iota
)

// Level - return indices stats aggregated at index, node or shard level.
type Level int

const (
	// LevelIndices can be used to set Level to "indices"
	LevelIndices = iota
	// LevelNode can be used to set Level to "node"
	LevelNode = iota
	// LevelShards can be used to set Level to "shards"
	LevelShards = iota
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option struct {
	name  string
	apply func(r *transport.Request)
}

// WithCompletionFields - a comma-separated list of fields for "fielddata" and "suggest" index metric (supports wildcards).
func WithCompletionFields(completionFields []string) Option {
	return Option{
		name: "WithCompletionFields",
		apply: func(r *transport.Request) {
		},
	}
}

// WithType - the type to sample (default: cpu).
func WithType(documentType DocumentType) Option {
	return Option{
		name: "WithType",
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

// WithFielddataFields - a comma-separated list of fields for "fielddata" index metric (supports wildcards).
func WithFielddataFields(fielddataFields []string) Option {
	return Option{
		name: "WithFielddataFields",
		apply: func(r *transport.Request) {
		},
	}
}

// WithFields - a comma-separated list of fields for "fielddata" and "completion" index metric (supports wildcards).
func WithFields(fields []string) Option {
	return Option{
		name: "WithFields",
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

// WithGroups - a comma-separated list of search groups for "search" index metric.
func WithGroups(groups bool) Option {
	return Option{
		name: "WithGroups",
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

// WithIgnoreIdleThreads - don't show threads that are in known-idle places, such as waiting on a socket select or pulling from an empty task queue (default: true).
func WithIgnoreIdleThreads(ignoreIdleThreads bool) Option {
	return Option{
		name: "WithIgnoreIdleThreads",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIncludeSegmentFileSizes - whether to report the aggregated disk usage of each one of the Lucene index files (only applies if segment stats are requested).
func WithIncludeSegmentFileSizes(includeSegmentFileSizes bool) Option {
	return Option{
		name: "WithIncludeSegmentFileSizes",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIndexMetric - limit the information returned for "indices" metric to the specific index metrics. Isn't used if "indices" (or "all") metric isn't specified.
func WithIndexMetric(indexMetric []string) Option {
	return Option{
		name: "WithIndexMetric",
		apply: func(r *transport.Request) {
		},
	}
}

// WithInterval - the interval for the second sampling of threads.
func WithInterval(interval time.Duration) Option {
	return Option{
		name: "WithInterval",
		apply: func(r *transport.Request) {
		},
	}
}

// WithLevel - return indices stats aggregated at index, node or shard level.
func WithLevel(level Level) Option {
	return Option{
		name: "WithLevel",
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

// WithSnapshots - number of samples of thread stacktrace (default: 10).
func WithSnapshots(snapshots int) Option {
	return Option{
		name: "WithSnapshots",
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

// WithThreads - specify the number of threads to provide information for (default: 3).
func WithThreads(threads int) Option {
	return Option{
		name: "WithThreads",
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

// WithTypes - a comma-separated list of document types for the "indexing" index metric.
func WithTypes(types []string) Option {
	return Option{
		name: "WithTypes",
		apply: func(r *transport.Request) {
		},
	}
}

var (
	supportedOptions = map[string]map[string]struct{}{
		"Stats": map[string]struct{}{
			"WithIndexMetric":             struct{}{},
			"WithMetric":                  struct{}{},
			"WithNodeID":                  struct{}{},
			"WithCompletionFields":        struct{}{},
			"WithFielddataFields":         struct{}{},
			"WithFields":                  struct{}{},
			"WithGroups":                  struct{}{},
			"WithIncludeSegmentFileSizes": struct{}{},
			"WithLevel":                   struct{}{},
			"WithTimeout":                 struct{}{},
			"WithTypes":                   struct{}{},
			"WithErrorTrace":              struct{}{},
			"WithFilterPath":              struct{}{},
			"WithHuman":                   struct{}{},
			"WithIgnore":                  struct{}{},
			"WithPretty":                  struct{}{},
			"WithSourceParam":             struct{}{},
		},
		"Info": map[string]struct{}{
			"WithMetric":       struct{}{},
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
		"HotThreads": map[string]struct{}{
			"WithNodeID":            struct{}{},
			"WithIgnoreIdleThreads": struct{}{},
			"WithInterval":          struct{}{},
			"WithSnapshots":         struct{}{},
			"WithThreads":           struct{}{},
			"WithTimeout":           struct{}{},
			"WithType":              struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
	}
)

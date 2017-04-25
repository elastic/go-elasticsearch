// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package nodes

import (
	"net/http"
	"time"
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option func(r *http.Request)

// WithCompletionFields a comma-separated list of fields for fielddata and suggest index metric (supports wildcards).
func WithCompletionFields(completionFields []string) Option {
	return func(r *http.Request) {
	}
}

// WithType the type to sample (default: cpu).
func WithType(documentType struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithErrorTrace include the stack trace of returned errors.
func WithErrorTrace(errorTrace bool) Option {
	return func(r *http.Request) {
	}
}

// WithFielddataFields a comma-separated list of fields for fielddata index metric (supports wildcards).
func WithFielddataFields(fielddataFields []string) Option {
	return func(r *http.Request) {
	}
}

// WithFields a comma-separated list of fields for fielddata and completion index metric (supports wildcards).
func WithFields(fields []string) Option {
	return func(r *http.Request) {
	}
}

// WithFilterPath a comma-separated list of filters used to reduce the respone.
func WithFilterPath(filterPath []string) Option {
	return func(r *http.Request) {
	}
}

// WithFlatSettings return settings in flat format (default: false).
func WithFlatSettings(flatSettings bool) Option {
	return func(r *http.Request) {
	}
}

// WithGroups a comma-separated list of search groups for search index metric.
func WithGroups(groups bool) Option {
	return func(r *http.Request) {
	}
}

// WithHuman return human readable values for statistics.
func WithHuman(human bool) Option {
	return func(r *http.Request) {
	}
}

// WithIgnoreIdleThreads don't show threads that are in known-idle places, such as waiting on a socket select or pulling from an empty task queue (default: true).
func WithIgnoreIdleThreads(ignoreIdleThreads bool) Option {
	return func(r *http.Request) {
	}
}

// WithIncludeSegmentFileSizes whether to report the aggregated disk usage of each one of the Lucene index files (only applies if segment stats are requested).
func WithIncludeSegmentFileSizes(includeSegmentFileSizes bool) Option {
	return func(r *http.Request) {
	}
}

// WithIndexMetric limit the information returned for indices metric to the specific index metrics. Isn't used if indices (or all) metric isn't specified.
func WithIndexMetric(indexMetric []string) Option {
	return func(r *http.Request) {
	}
}

// WithInterval the interval for the second sampling of threads.
func WithInterval(interval time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithLevel return indices stats aggregated at index, node or shard level.
func WithLevel(level struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithMetric a comma-separated list of metrics you wish returned. Leave empty to return all.
func WithMetric(metric []string) Option {
	return func(r *http.Request) {
	}
}

// WithNodeID a comma-separated list of node IDs or names to limit the returned information; use _local to return information from the node you're connecting to, leave empty to get information from all nodes.
func WithNodeID(nodeID []string) Option {
	return func(r *http.Request) {
	}
}

// WithPretty pretty format the returned JSON response.
func WithPretty(pretty bool) Option {
	return func(r *http.Request) {
	}
}

// WithSnapshots number of samples of thread stacktrace (default: 10).
func WithSnapshots(snapshots int) Option {
	return func(r *http.Request) {
	}
}

// WithSourceParam the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSourceParam(sourceParam string) Option {
	return func(r *http.Request) {
	}
}

// WithThreads specify the number of threads to provide information for (default: 3).
func WithThreads(threads int) Option {
	return func(r *http.Request) {
	}
}

// WithTimeout explicit operation timeout.
func WithTimeout(timeout time.Time) Option {
	return func(r *http.Request) {
	}
}

// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"net/http"
	"time"
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option func(r *http.Request)

// WithActiveOnly display only those recoveries that are currently on-going.
func WithActiveOnly(activeOnly bool) Option {
	return func(r *http.Request) {
	}
}

// WithAllShards execute validation on all shards instead of one random shard per index.
func WithAllShards(allShards bool) Option {
	return func(r *http.Request) {
	}
}

// WithAllowNoIndices whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes _all string or when no indices have been specified).
func WithAllowNoIndices(allowNoIndices bool) Option {
	return func(r *http.Request) {
	}
}

// WithAnalyzeWildcard specify whether wildcard and prefix queries should be analyzed (default: false).
func WithAnalyzeWildcard(analyzeWildcard bool) Option {
	return func(r *http.Request) {
	}
}

// WithAnalyzer the analyzer to use for the query string.
func WithAnalyzer(analyzer string) Option {
	return func(r *http.Request) {
	}
}

// WithCompletionFields a comma-separated list of fields for fielddata and suggest index metric (supports wildcards).
func WithCompletionFields(completionFields []string) Option {
	return func(r *http.Request) {
	}
}

// WithCreate whether the index template should only be added if new or can also replace an existing one.
func WithCreate(create bool) Option {
	return func(r *http.Request) {
	}
}

// WithDefaultOperator the default operator for query string query (AND or OR).
func WithDefaultOperator(defaultOperator struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithDetailed whether to display detailed information about shard recovery.
func WithDetailed(detailed bool) Option {
	return func(r *http.Request) {
	}
}

// WithDf the field to use as default where no field prefix is given in the query string.
func WithDf(df string) Option {
	return func(r *http.Request) {
	}
}

// WithType a comma-separated list of document types.
func WithType(documentType []string) Option {
	return func(r *http.Request) {
	}
}

// WithDryRun if set to true the rollover action will only be validated but not actually performed even if a condition matches. The default is false.
func WithDryRun(dryRun bool) Option {
	return func(r *http.Request) {
	}
}

// WithErrorTrace include the stack trace of returned errors.
func WithErrorTrace(errorTrace bool) Option {
	return func(r *http.Request) {
	}
}

// WithExpandWildcards whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithExpandWildcards(expandWildcards struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithFieldData clear field data.
func WithFieldData(fieldData bool) Option {
	return func(r *http.Request) {
	}
}

// WithFielddata clear field data.
func WithFielddata(fielddata bool) Option {
	return func(r *http.Request) {
	}
}

// WithFielddataFields a comma-separated list of fields for fielddata index metric (supports wildcards).
func WithFielddataFields(fielddataFields []string) Option {
	return func(r *http.Request) {
	}
}

// WithFields a comma-separated list of fields to clear when using the field_data parameter (default: all).
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

// WithFlush specify whether the index should be flushed after performing the operation (default: true).
func WithFlush(flush bool) Option {
	return func(r *http.Request) {
	}
}

// WithForce whether a flush should be forced even if it is not necessarily needed ie. if no changes will be committed to the index. This is useful if transaction log IDs should be incremented even if no uncommitted changes are present. (This setting can be considered as internal).
func WithForce(force bool) Option {
	return func(r *http.Request) {
	}
}

// WithFormat format of the output.
func WithFormat(format struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithHuman return human readable values for statistics.
func WithHuman(human bool) Option {
	return func(r *http.Request) {
	}
}

// WithIgnoreUnavailable whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithIgnoreUnavailable(ignoreUnavailable bool) Option {
	return func(r *http.Request) {
	}
}

// WithIncludeDefaults whether the default mapping values should be returned as well.
func WithIncludeDefaults(includeDefaults bool) Option {
	return func(r *http.Request) {
	}
}

// WithIndex the name of the index to scope the operation.
func WithIndex(index string) Option {
	return func(r *http.Request) {
	}
}

// WithIndexParam the name of the index to scope the operation.
func WithIndexParam(indexParam string) Option {
	return func(r *http.Request) {
	}
}

// WithLocal return local information, do not retrieve the state from master node (default: false).
func WithLocal(local bool) Option {
	return func(r *http.Request) {
	}
}

// WithMasterTimeout specify timeout for connection to master.
func WithMasterTimeout(masterTimeout time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithNewIndex the name of the rollover index.
func WithNewIndex(newIndex string) Option {
	return func(r *http.Request) {
	}
}

// WithOrder the order for this template when merging multiple matching ones (higher numbers are merged later, overriding the lower numbers).
func WithOrder(order int) Option {
	return func(r *http.Request) {
	}
}

// WithPreferLocal with true, specify that a local shard should be used if available, with false, use a random shard (default: true).
func WithPreferLocal(preferLocal bool) Option {
	return func(r *http.Request) {
	}
}

// WithPretty pretty format the returned JSON response.
func WithPretty(pretty bool) Option {
	return func(r *http.Request) {
	}
}

// WithSourceParam the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSourceParam(sourceParam string) Option {
	return func(r *http.Request) {
	}
}

// WithTimeout explicit operation timeout.
func WithTimeout(timeout time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithUpdateAllTypes whether to update the mapping for all fields with the same name across all types or not.
func WithUpdateAllTypes(updateAllTypes bool) Option {
	return func(r *http.Request) {
	}
}

// WithWaitForActiveShards set the number of active shards to wait for before the operation returns.
func WithWaitForActiveShards(waitForActiveShards string) Option {
	return func(r *http.Request) {
	}
}

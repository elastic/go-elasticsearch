// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"net/http"
	"time"
)

// Conflicts - what to do when the delete-by-query hits version conflicts?
type Conflicts int

const (
	// ConflictsAbort can be used to set Conflicts to "abort"
	ConflictsAbort = iota
	// ConflictsProceed can be used to set Conflicts to "proceed"
	ConflictsProceed = iota
)

// DefaultOperator - the default operator for query string query (AND or OR).
type DefaultOperator int

const (
	// DefaultOperatorAND can be used to set DefaultOperator to "AND"
	DefaultOperatorAND = iota
	// DefaultOperatorOR can be used to set DefaultOperator to "OR"
	DefaultOperatorOR = iota
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

// OpType - explicit operation type.
type OpType int

const (
	// OpTypeIndex can be used to set OpType to "index"
	OpTypeIndex = iota
	// OpTypeCreate can be used to set OpType to "create"
	OpTypeCreate = iota
)

// Refresh - if "true" then refresh the effected shards to make this operation visible to search, if "wait_for" then wait for a refresh to make this operation visible to search, if "false" (the default) then do nothing with refreshes.
type Refresh int

const (
	// RefreshTrue can be used to set Refresh to "true"
	RefreshTrue = iota
	// RefreshFalse can be used to set Refresh to "false"
	RefreshFalse = iota
	// RefreshWaitFor can be used to set Refresh to "wait_for"
	RefreshWaitFor = iota
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option func(r *http.Request)

// WithAllowNoIndices whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
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

// WithBatchedReduceSize the number of shard results that should be reduced at once on the coordinating node. This value should be used as a protection mechanism to reduce the memory overhead per search request if the potential number of shards in the request can be large.
func WithBatchedReduceSize(batchedReduceSize int) Option {
	return func(r *http.Request) {
	}
}

// WithConflicts what to do when the delete-by-query hits version conflicts?
func WithConflicts(conflicts Conflicts) Option {
	return func(r *http.Request) {
	}
}

// WithDefaultOperator the default operator for query string query (AND or OR).
func WithDefaultOperator(defaultOperator DefaultOperator) Option {
	return func(r *http.Request) {
	}
}

// WithDf the field to use as default where no field prefix is given in the query string.
func WithDf(df string) Option {
	return func(r *http.Request) {
	}
}

// WithType default document type for items which don't provide one.
func WithType(documentType string) Option {
	return func(r *http.Request) {
	}
}

// WithTypeParam default document type for items which don't provide one.
func WithTypeParam(documentTypeParam string) Option {
	return func(r *http.Request) {
	}
}

// WithErrorTrace include the stack trace of returned errors.
func WithErrorTrace(errorTrace bool) Option {
	return func(r *http.Request) {
	}
}

// WithExpandWildcards whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithExpandWildcards(expandWildcards ExpandWildcards) Option {
	return func(r *http.Request) {
	}
}

// WithFieldStatistics specifies if document count, sum of document frequencies and sum of total term frequencies should be returned.
func WithFieldStatistics(fieldStatistics bool) Option {
	return func(r *http.Request) {
	}
}

// WithFields default comma-separated list of fields to return in the response for updates, can be overridden on each sub-request.
func WithFields(fields []string) Option {
	return func(r *http.Request) {
	}
}

// WithFilterPath a comma-separated list of filters used to reduce the respone.
func WithFilterPath(filterPath []string) Option {
	return func(r *http.Request) {
	}
}

// WithHuman return human readable values for statistics.
func WithHuman(human bool) Option {
	return func(r *http.Request) {
	}
}

// WithID document ID.
func WithID(id string) Option {
	return func(r *http.Request) {
	}
}

// WithIgnoreUnavailable whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithIgnoreUnavailable(ignoreUnavailable bool) Option {
	return func(r *http.Request) {
	}
}

// WithIndex default index for items which don't provide one.
func WithIndex(index string) Option {
	return func(r *http.Request) {
	}
}

// WithOpType explicit operation type.
func WithOpType(opType OpType) Option {
	return func(r *http.Request) {
	}
}

// WithParent ID of the parent document.
func WithParent(parent string) Option {
	return func(r *http.Request) {
	}
}

// WithPipeline the pipeline id to preprocess incoming documents with.
func WithPipeline(pipeline string) Option {
	return func(r *http.Request) {
	}
}

// WithPretty pretty format the returned JSON response.
func WithPretty(pretty bool) Option {
	return func(r *http.Request) {
	}
}

// WithRefresh if "true" then refresh the effected shards to make this operation visible to search, if "wait_for" then wait for a refresh to make this operation visible to search, if "false" (the default) then do nothing with refreshes.
func WithRefresh(refresh Refresh) Option {
	return func(r *http.Request) {
	}
}

// WithRouting specific routing value.
func WithRouting(routing string) Option {
	return func(r *http.Request) {
	}
}

// WithScroll specify how long a consistent view of the index should be maintained for scrolled search.
func WithScroll(scroll time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithScrollID a comma-separated list of scroll IDs to clear.
func WithScrollID(scrollID []string) Option {
	return func(r *http.Request) {
	}
}

// WithSource true or false to return the _source field or not, or default list of fields to return, can be overridden on each sub-request.
func WithSource(source []string) Option {
	return func(r *http.Request) {
	}
}

// WithSourceExclude default list of fields to exclude from the returned _source field, can be overridden on each sub-request.
func WithSourceExclude(sourceExclude []string) Option {
	return func(r *http.Request) {
	}
}

// WithSourceInclude default list of fields to extract and return from the _source field, can be overridden on each sub-request.
func WithSourceInclude(sourceInclude []string) Option {
	return func(r *http.Request) {
	}
}

// WithSourceParam the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSourceParam(sourceParam string) Option {
	return func(r *http.Request) {
	}
}

// WithTaskID the task id to rethrottle.
func WithTaskID(taskID string) Option {
	return func(r *http.Request) {
	}
}

// WithTimeout explicit operation timeout.
func WithTimeout(timeout time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithWaitForActiveShards sets the number of shard copies that must be active before proceeding with the bulk operation. Defaults to 1, meaning the primary shard only. Set to "all" for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas + 1).
func WithWaitForActiveShards(waitForActiveShards string) Option {
	return func(r *http.Request) {
	}
}

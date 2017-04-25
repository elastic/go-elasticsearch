// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"net/http"
	"time"
)

// Bytes - the unit in which to display byte values.
type Bytes int

const (
	// BytesB can be used to set Bytes to "b"
	BytesB = iota
	// BytesK can be used to set Bytes to "k"
	BytesK = iota
	// BytesKb can be used to set Bytes to "kb"
	BytesKb = iota
	// BytesM can be used to set Bytes to "m"
	BytesM = iota
	// BytesMb can be used to set Bytes to "mb"
	BytesMb = iota
	// BytesG can be used to set Bytes to "g"
	BytesG = iota
	// BytesGb can be used to set Bytes to "gb"
	BytesGb = iota
	// BytesT can be used to set Bytes to "t"
	BytesT = iota
	// BytesTb can be used to set Bytes to "tb"
	BytesTb = iota
	// BytesP can be used to set Bytes to "p"
	BytesP = iota
	// BytesPb can be used to set Bytes to "pb"
	BytesPb = iota
)

// Size - the multiplier in which to display values.
type Size int

const (
	// SizeZero can be used to set Size to "zero"
	SizeZero = iota
	// SizeK can be used to set Size to "k"
	SizeK = iota
	// SizeM can be used to set Size to "m"
	SizeM = iota
	// SizeG can be used to set Size to "g"
	SizeG = iota
	// SizeT can be used to set Size to "t"
	SizeT = iota
	// SizeP can be used to set Size to "p"
	SizeP = iota
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option func(r *http.Request)

// WithActions a comma-separated list of actions that should be returned. Leave empty to return all.
func WithActions(actions []string) Option {
	return func(r *http.Request) {
	}
}

// WithBytes the unit in which to display byte values.
func WithBytes(bytes Bytes) Option {
	return func(r *http.Request) {
	}
}

// WithDetailed return detailed task information (default: false).
func WithDetailed(detailed bool) Option {
	return func(r *http.Request) {
	}
}

// WithErrorTrace include the stack trace of returned errors.
func WithErrorTrace(errorTrace bool) Option {
	return func(r *http.Request) {
	}
}

// WithFields a comma-separated list of fields to return the fielddata size.
func WithFields(fields []string) Option {
	return func(r *http.Request) {
	}
}

// WithFieldsParam a comma-separated list of fields to return in the output.
func WithFieldsParam(fieldsParam []string) Option {
	return func(r *http.Request) {
	}
}

// WithFilterPath a comma-separated list of filters used to reduce the respone.
func WithFilterPath(filterPath []string) Option {
	return func(r *http.Request) {
	}
}

// WithFormat a short version of the Accept header, e.g. json, yaml.
func WithFormat(format string) Option {
	return func(r *http.Request) {
	}
}

// WithFullID return the full node ID instead of the shortened version (default: false).
func WithFullID(fullID bool) Option {
	return func(r *http.Request) {
	}
}

// WithH comma-separated list of column names to display.
func WithH(h []string) Option {
	return func(r *http.Request) {
	}
}

// WithHelp return help information.
func WithHelp(help bool) Option {
	return func(r *http.Request) {
	}
}

// WithHuman return human readable values for statistics.
func WithHuman(human bool) Option {
	return func(r *http.Request) {
	}
}

// WithIgnoreUnavailable set to true to ignore unavailable snapshots.
func WithIgnoreUnavailable(ignoreUnavailable bool) Option {
	return func(r *http.Request) {
	}
}

// WithIndex a comma-separated list of index names to limit the returned information.
func WithIndex(index []string) Option {
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

// WithName a comma-separated list of alias names to return.
func WithName(name []string) Option {
	return func(r *http.Request) {
	}
}

// WithNodeID a comma-separated list of node IDs or names to limit the returned information.
func WithNodeID(nodeID []string) Option {
	return func(r *http.Request) {
	}
}

// WithPretty pretty format the returned JSON response.
func WithPretty(pretty bool) Option {
	return func(r *http.Request) {
	}
}

// WithS comma-separated list of column names or column aliases to sort by.
func WithS(s []string) Option {
	return func(r *http.Request) {
	}
}

// WithSize the multiplier in which to display values.
func WithSize(size Size) Option {
	return func(r *http.Request) {
	}
}

// WithSourceParam the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSourceParam(sourceParam string) Option {
	return func(r *http.Request) {
	}
}

// WithThreadPoolPatterns a comma-separated list of regular-expressions to filter the thread pools in the output.
func WithThreadPoolPatterns(threadPoolPatterns []string) Option {
	return func(r *http.Request) {
	}
}

// WithTs set to false to disable timestamping.
func WithTs(ts bool) Option {
	return func(r *http.Request) {
	}
}

// WithV verbose mode. Display column headers.
func WithV(v bool) Option {
	return func(r *http.Request) {
	}
}

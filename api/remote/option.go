// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package remote

import (
	"github.com/elastic/go-elasticsearch/transport"
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option struct {
	name  string
	apply func(r *transport.Request)
}

// WithErrorTrace - include the stack trace of returned errors.
func WithErrorTrace(errorTrace bool) Option {
	return Option{
		name: "WithErrorTrace",
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

var (
	supportedOptions = map[string]map[string]struct{}{
		"Info": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
	}
)

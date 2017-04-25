// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// ValidateQuery - the validate API allows a user to validate a potentially expensive query without executing it. See http://www.elastic.co/guide/en/elasticsearch/reference/master/search-validate.html for more info.
//
// body: the query definition specified with the Query DSL.
//
// options: optional parameters. Supports the following functional options: WithAllShards, WithAllowNoIndices, WithAnalyzeWildcard, WithAnalyzer, WithDefaultOperator, WithDf, WithType, WithErrorTrace, WithExpandWildcards, WithExplain, WithFilterPath, WithHuman, WithIgnoreUnavailable, WithIndex, WithLenient, WithOperationThreading, WithPretty, WithQ, WithRewrite, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) ValidateQuery(body map[string]interface{}, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithAllShards":          struct{}{},
		"WithAllowNoIndices":     struct{}{},
		"WithAnalyzeWildcard":    struct{}{},
		"WithAnalyzer":           struct{}{},
		"WithDefaultOperator":    struct{}{},
		"WithDf":                 struct{}{},
		"WithType":               struct{}{},
		"WithErrorTrace":         struct{}{},
		"WithExpandWildcards":    struct{}{},
		"WithExplain":            struct{}{},
		"WithFilterPath":         struct{}{},
		"WithHuman":              struct{}{},
		"WithIgnoreUnavailable":  struct{}{},
		"WithIndex":              struct{}{},
		"WithLenient":            struct{}{},
		"WithOperationThreading": struct{}{},
		"WithPretty":             struct{}{},
		"WithQ":                  struct{}{},
		"WithRewrite":            struct{}{},
		"WithSourceParam":        struct{}{},
	}
	for _, option := range options {
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "GET",
	}
	return i.client.Do(req)
}

// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// Search allows you to execute a search query and get back search hits that match the query. See http://www.elastic.co/guide/en/elasticsearch/reference/master/search-search.html for more info.
//
// body: the search definition using the Query DSL.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithAnalyzeWildcard, WithAnalyzer, WithBatchedReduceSize, WithDefaultOperator, WithDf, WithType, WithDocvalueFields, WithErrorTrace, WithExpandWildcards, WithExplain, WithFielddataFields, WithFilterPath, WithFrom, WithHuman, WithIgnoreUnavailable, WithIndex, WithLenient, WithPreference, WithPretty, WithQ, WithRequestCache, WithRouting, WithScroll, WithSearchType, WithSize, WithSort, WithSource, WithSourceExclude, WithSourceInclude, WithSourceParam, WithStats, WithStoredFields, WithSuggestField, WithSuggestMode, WithSuggestSize, WithSuggestText, WithTerminateAfter, WithTimeout, WithTrackScores, WithTypedKeys, WithVersion, see the Option type in this package for more info.
func (c *Client) Search(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithAllowNoIndices":    struct{}{},
		"WithAnalyzeWildcard":   struct{}{},
		"WithAnalyzer":          struct{}{},
		"WithBatchedReduceSize": struct{}{},
		"WithDefaultOperator":   struct{}{},
		"WithDf":                struct{}{},
		"WithType":              struct{}{},
		"WithDocvalueFields":    struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithExpandWildcards":   struct{}{},
		"WithExplain":           struct{}{},
		"WithFielddataFields":   struct{}{},
		"WithFilterPath":        struct{}{},
		"WithFrom":              struct{}{},
		"WithHuman":             struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithIndex":             struct{}{},
		"WithLenient":           struct{}{},
		"WithPreference":        struct{}{},
		"WithPretty":            struct{}{},
		"WithQ":                 struct{}{},
		"WithRequestCache":      struct{}{},
		"WithRouting":           struct{}{},
		"WithScroll":            struct{}{},
		"WithSearchType":        struct{}{},
		"WithSize":              struct{}{},
		"WithSort":              struct{}{},
		"WithSource":            struct{}{},
		"WithSourceExclude":     struct{}{},
		"WithSourceInclude":     struct{}{},
		"WithSourceParam":       struct{}{},
		"WithStats":             struct{}{},
		"WithStoredFields":      struct{}{},
		"WithSuggestField":      struct{}{},
		"WithSuggestMode":       struct{}{},
		"WithSuggestSize":       struct{}{},
		"WithSuggestText":       struct{}{},
		"WithTerminateAfter":    struct{}{},
		"WithTimeout":           struct{}{},
		"WithTrackScores":       struct{}{},
		"WithTypedKeys":         struct{}{},
		"WithVersion":           struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.Scheme,
			Host:   c.transport.Host,
		},
		Method: "GET",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return c.transport.Do(req)
}

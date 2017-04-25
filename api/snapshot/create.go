// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package snapshot

import (
	"fmt"
	"net/http"
	"net/url"
)

// Create - the snapshot and restore module allows to create snapshots of individual indices or an entire cluster into a remote repository like shared file system, S3, or HDFS. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-snapshots.html for more info.
//
// repository: a repository name.
//
// snapshot: a snapshot name.
//
// body: the snapshot definition.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithMasterTimeout, WithPretty, WithSourceParam, WithWaitForCompletion, see the Option type in this package for more info.
func (s *Snapshot) Create(repository string, snapshot string, body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":        struct{}{},
		"WithFilterPath":        struct{}{},
		"WithHuman":             struct{}{},
		"WithMasterTimeout":     struct{}{},
		"WithPretty":            struct{}{},
		"WithSourceParam":       struct{}{},
		"WithWaitForCompletion": struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: s.transport.Scheme,
			Host:   s.transport.Host,
		},
		Method: "PUT",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return s.transport.Do(req)
}

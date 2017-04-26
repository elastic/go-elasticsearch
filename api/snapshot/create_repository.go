// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package snapshot

import (
	"fmt"
	"net/http"
	"net/url"
)

// CreateRepository - the snapshot and restore module allows to create snapshots of individual indices or an entire cluster into a remote repository like shared file system, S3, or HDFS. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-snapshots.html for more info.
//
// repository: a repository name.
//
// body: the repository definition.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithTimeout, WithVerify, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (s *Snapshot) CreateRepository(repository string, body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithMasterTimeout": struct{}{},
		"WithTimeout":       struct{}{},
		"WithVerify":        struct{}{},
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithPretty":        struct{}{},
		"WithSourceParam":   struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: s.transport.URL.Scheme,
			Host:   s.transport.URL.Host,
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

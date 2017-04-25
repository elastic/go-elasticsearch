// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package snapshot

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Delete - the snapshot and restore module allows to create snapshots of individual indices or an entire cluster into a remote repository like shared file system, S3, or HDFS. See http://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html for more info.
//
// repository: a repository name.
//
// snapshot: a snapshot name.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithMasterTimeout, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (s *Snapshot) Delete(repository string, snapshot string, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithPretty":        struct{}{},
		"WithSourceParam":   struct{}{},
	}
	for _, option := range options {
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "DELETE",
	}
	return s.client.Do(req)
}

// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// PutTemplate - see http://www.elastic.co/guide/en/elasticsearch/reference/master/search-template.html for more info.
//
// id: template ID.
//
// body: the document.
func (c *Client) PutTemplate(id string, body map[string]interface{}, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{}
	for _, option := range options {
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "PUT",
	}
	return c.client.Do(req)
}

// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"net/url"
)

// Templates - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-templates.html for more info.
//
// options: optional parameters. Supports the following functional options: WithName, WithFormat, WithH, WithHelp, WithLocal, WithMasterTimeout, WithS, WithV, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cat) Templates(options ...*Option) (*TemplatesResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Templates"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &TemplatesResponse{resp}, err
}

// TemplatesResponse is the response for Templates
type TemplatesResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

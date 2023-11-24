// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Adds a node to be shut down. Designed for indirect use by ECE/ESS and ECK.
// Direct use is not supported.
package putnode

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/type_"
)

const (
	nodeidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutNode struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	nodeid string
}

// NewPutNode type alias for index.
type NewPutNode func(nodeid string) *PutNode

// NewPutNodeFunc returns a new instance of PutNode with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutNodeFunc(tp elastictransport.Interface) NewPutNode {
	return func(nodeid string) *PutNode {
		n := New(tp)

		n._nodeid(nodeid)

		return n
	}
}

// Adds a node to be shut down. Designed for indirect use by ECE/ESS and ECK.
// Direct use is not supported.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current
func New(tp elastictransport.Interface) *PutNode {
	r := &PutNode{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutNode) Raw(raw io.Reader) *PutNode {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutNode) Request(req *Request) *PutNode {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutNode) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutNode: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nodeidMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("shutdown")

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PutNode) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutNode query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putnode.Response
func (r PutNode) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the PutNode headers map.
func (r *PutNode) Header(key, value string) *PutNode {
	r.headers.Set(key, value)

	return r
}

// NodeId The node id of node to be shut down
// API Name: nodeid
func (r *PutNode) _nodeid(nodeid string) *PutNode {
	r.paramSet |= nodeidMask
	r.nodeid = nodeid

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *PutNode) MasterTimeout(mastertimeout timeunit.TimeUnit) *PutNode {
	r.values.Set("master_timeout", mastertimeout.String())

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *PutNode) Timeout(timeout timeunit.TimeUnit) *PutNode {
	r.values.Set("timeout", timeout.String())

	return r
}

// AllocationDelay Only valid if type is restart.
// Controls how long Elasticsearch will wait for the node to restart and join
// the cluster before reassigning its shards to other nodes.
// This works the same as delaying allocation with the
// index.unassigned.node_left.delayed_timeout setting.
// If you specify both a restart allocation delay and an index-level allocation
// delay, the longer of the two is used.
// API name: allocation_delay
func (r *PutNode) AllocationDelay(allocationdelay string) *PutNode {

	r.req.AllocationDelay = &allocationdelay

	return r
}

// Reason A human-readable reason that the node is being shut down.
// This field provides information for other cluster operators; it does not
// affect the shut down process.
// API name: reason
func (r *PutNode) Reason(reason string) *PutNode {

	r.req.Reason = reason

	return r
}

// TargetNodeName Only valid if type is replace.
// Specifies the name of the node that is replacing the node being shut down.
// Shards from the shut down node are only allowed to be allocated to the target
// node, and no other data will be allocated to the target node.
// During relocation of data certain allocation rules are ignored, such as disk
// watermarks or user attribute filtering rules.
// API name: target_node_name
func (r *PutNode) TargetNodeName(targetnodename string) *PutNode {

	r.req.TargetNodeName = &targetnodename

	return r
}

// Type Valid values are restart, remove, or replace.
// Use restart when you need to temporarily shut down a node to perform an
// upgrade, make configuration changes, or perform other maintenance.
// Because the node is expected to rejoin the cluster, data is not migrated off
// of the node.
// Use remove when you need to permanently remove a node from the cluster.
// The node is not marked ready for shutdown until data is migrated off of the
// node Use replace to do a 1:1 replacement of a node with another node.
// Certain allocation decisions will be ignored (such as disk watermarks) in the
// interest of true replacement of the source node with the target node.
// During a replace-type shutdown, rollover and index creation may result in
// unassigned shards, and shrink may fail until the replacement is complete.
// API name: type
func (r *PutNode) Type(type_ type_.Type) *PutNode {
	r.req.Type = type_

	return r
}

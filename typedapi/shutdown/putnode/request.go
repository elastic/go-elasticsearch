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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package putnode

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/type_"
)

// Request holds the request body struct for the package putnode
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/shutdown/put_node/ShutdownPutNodeRequest.ts#L25-L77
type Request struct {

	// AllocationDelay Only valid if type is restart.
	// Controls how long Elasticsearch will wait for the node to restart and join
	// the cluster before reassigning its shards to other nodes.
	// This works the same as delaying allocation with the
	// index.unassigned.node_left.delayed_timeout setting.
	// If you specify both a restart allocation delay and an index-level allocation
	// delay, the longer of the two is used.
	AllocationDelay *string `json:"allocation_delay,omitempty"`

	// Reason A human-readable reason that the node is being shut down.
	// This field provides information for other cluster operators; it does not
	// affect the shut down process.
	Reason string `json:"reason"`

	// TargetNodeName Only valid if type is replace.
	// Specifies the name of the node that is replacing the node being shut down.
	// Shards from the shut down node are only allowed to be allocated to the target
	// node, and no other data will be allocated to the target node.
	// During relocation of data certain allocation rules are ignored, such as disk
	// watermarks or user attribute filtering rules.
	TargetNodeName *string `json:"target_node_name,omitempty"`

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
	Type type_.Type `json:"type"`
}

// RequestBuilder is the builder API for the putnode.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putnode request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) AllocationDelay(allocationdelay string) *RequestBuilder {
	rb.v.AllocationDelay = &allocationdelay
	return rb
}

func (rb *RequestBuilder) Reason(reason string) *RequestBuilder {
	rb.v.Reason = reason
	return rb
}

func (rb *RequestBuilder) TargetNodeName(targetnodename string) *RequestBuilder {
	rb.v.TargetNodeName = &targetnodename
	return rb
}

func (rb *RequestBuilder) Type_(type_ type_.Type) *RequestBuilder {
	rb.v.Type = type_
	return rb
}

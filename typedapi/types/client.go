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


package types

// Client type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L266-L278
type Client struct {
	Agent                 *string `json:"agent,omitempty"`
	ClosedTimeMillis      *int64  `json:"closed_time_millis,omitempty"`
	Id                    *int64  `json:"id,omitempty"`
	LastRequestTimeMillis *int64  `json:"last_request_time_millis,omitempty"`
	LastUri               *string `json:"last_uri,omitempty"`
	LocalAddress          *string `json:"local_address,omitempty"`
	OpenedTimeMillis      *int64  `json:"opened_time_millis,omitempty"`
	RemoteAddress         *string `json:"remote_address,omitempty"`
	RequestCount          *int64  `json:"request_count,omitempty"`
	RequestSizeBytes      *int64  `json:"request_size_bytes,omitempty"`
	XOpaqueId             *string `json:"x_opaque_id,omitempty"`
}

// ClientBuilder holds Client struct and provides a builder API.
type ClientBuilder struct {
	v *Client
}

// NewClient provides a builder for the Client struct.
func NewClientBuilder() *ClientBuilder {
	r := ClientBuilder{
		&Client{},
	}

	return &r
}

// Build finalize the chain and returns the Client struct
func (rb *ClientBuilder) Build() Client {
	return *rb.v
}

func (rb *ClientBuilder) Agent(agent string) *ClientBuilder {
	rb.v.Agent = &agent
	return rb
}

func (rb *ClientBuilder) ClosedTimeMillis(closedtimemillis int64) *ClientBuilder {
	rb.v.ClosedTimeMillis = &closedtimemillis
	return rb
}

func (rb *ClientBuilder) Id(id int64) *ClientBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *ClientBuilder) LastRequestTimeMillis(lastrequesttimemillis int64) *ClientBuilder {
	rb.v.LastRequestTimeMillis = &lastrequesttimemillis
	return rb
}

func (rb *ClientBuilder) LastUri(lasturi string) *ClientBuilder {
	rb.v.LastUri = &lasturi
	return rb
}

func (rb *ClientBuilder) LocalAddress(localaddress string) *ClientBuilder {
	rb.v.LocalAddress = &localaddress
	return rb
}

func (rb *ClientBuilder) OpenedTimeMillis(openedtimemillis int64) *ClientBuilder {
	rb.v.OpenedTimeMillis = &openedtimemillis
	return rb
}

func (rb *ClientBuilder) RemoteAddress(remoteaddress string) *ClientBuilder {
	rb.v.RemoteAddress = &remoteaddress
	return rb
}

func (rb *ClientBuilder) RequestCount(requestcount int64) *ClientBuilder {
	rb.v.RequestCount = &requestcount
	return rb
}

func (rb *ClientBuilder) RequestSizeBytes(requestsizebytes int64) *ClientBuilder {
	rb.v.RequestSizeBytes = &requestsizebytes
	return rb
}

func (rb *ClientBuilder) XOpaqueId(xopaqueid string) *ClientBuilder {
	rb.v.XOpaqueId = &xopaqueid
	return rb
}

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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package enrollnode

// Response holds the response body struct for the package enrollnode
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/enroll_node/Response.ts#L20-L47
type Response struct {

	// HttpCaCert The CA certificate that can be used by the new node in order to sign its
	// certificate for the HTTP layer, as a Base64 encoded string of the ASN.1 DER
	// encoding of the certificate.
	HttpCaCert string `json:"http_ca_cert"`
	// HttpCaKey The CA private key that can be used by the new node in order to sign its
	// certificate for the HTTP layer, as a Base64 encoded string of the ASN.1 DER
	// encoding of the key.
	HttpCaKey string `json:"http_ca_key"`
	// NodesAddresses A list of transport addresses in the form of `host:port` for the nodes that
	// are already members of the cluster.
	NodesAddresses []string `json:"nodes_addresses"`
	// TransportCaCert The CA certificate that is used to sign the TLS certificate for the transport
	// layer, as a Base64 encoded string of the ASN.1 DER encoding of the
	// certificate.
	TransportCaCert string `json:"transport_ca_cert"`
	// TransportCert The certificate that the node can use for TLS for its transport layer, as a
	// Base64 encoded string of the ASN.1 DER encoding of the certificate.
	TransportCert string `json:"transport_cert"`
	// TransportKey The private key that the node can use for TLS for its transport layer, as a
	// Base64 encoded string of the ASN.1 DER encoding of the key.
	TransportKey string `json:"transport_key"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

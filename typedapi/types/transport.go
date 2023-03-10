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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// Transport type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/nodes/_types/Stats.ts#L420-L431
type Transport struct {
	InboundHandlingTimeHistogram  []TransportHistogram `json:"inbound_handling_time_histogram,omitempty"`
	OutboundHandlingTimeHistogram []TransportHistogram `json:"outbound_handling_time_histogram,omitempty"`
	RxCount                       *int64               `json:"rx_count,omitempty"`
	RxSize                        *string              `json:"rx_size,omitempty"`
	RxSizeInBytes                 *int64               `json:"rx_size_in_bytes,omitempty"`
	ServerOpen                    *int                 `json:"server_open,omitempty"`
	TotalOutboundConnections      *int64               `json:"total_outbound_connections,omitempty"`
	TxCount                       *int64               `json:"tx_count,omitempty"`
	TxSize                        *string              `json:"tx_size,omitempty"`
	TxSizeInBytes                 *int64               `json:"tx_size_in_bytes,omitempty"`
}

// NewTransport returns a Transport.
func NewTransport() *Transport {
	r := &Transport{}

	return r
}

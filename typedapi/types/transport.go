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

// Transport type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L413-L424
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

// TransportBuilder holds Transport struct and provides a builder API.
type TransportBuilder struct {
	v *Transport
}

// NewTransport provides a builder for the Transport struct.
func NewTransportBuilder() *TransportBuilder {
	r := TransportBuilder{
		&Transport{},
	}

	return &r
}

// Build finalize the chain and returns the Transport struct
func (rb *TransportBuilder) Build() Transport {
	return *rb.v
}

func (rb *TransportBuilder) InboundHandlingTimeHistogram(inbound_handling_time_histogram []TransportHistogramBuilder) *TransportBuilder {
	tmp := make([]TransportHistogram, len(inbound_handling_time_histogram))
	for _, value := range inbound_handling_time_histogram {
		tmp = append(tmp, value.Build())
	}
	rb.v.InboundHandlingTimeHistogram = tmp
	return rb
}

func (rb *TransportBuilder) OutboundHandlingTimeHistogram(outbound_handling_time_histogram []TransportHistogramBuilder) *TransportBuilder {
	tmp := make([]TransportHistogram, len(outbound_handling_time_histogram))
	for _, value := range outbound_handling_time_histogram {
		tmp = append(tmp, value.Build())
	}
	rb.v.OutboundHandlingTimeHistogram = tmp
	return rb
}

func (rb *TransportBuilder) RxCount(rxcount int64) *TransportBuilder {
	rb.v.RxCount = &rxcount
	return rb
}

func (rb *TransportBuilder) RxSize(rxsize string) *TransportBuilder {
	rb.v.RxSize = &rxsize
	return rb
}

func (rb *TransportBuilder) RxSizeInBytes(rxsizeinbytes int64) *TransportBuilder {
	rb.v.RxSizeInBytes = &rxsizeinbytes
	return rb
}

func (rb *TransportBuilder) ServerOpen(serveropen int) *TransportBuilder {
	rb.v.ServerOpen = &serveropen
	return rb
}

func (rb *TransportBuilder) TotalOutboundConnections(totaloutboundconnections int64) *TransportBuilder {
	rb.v.TotalOutboundConnections = &totaloutboundconnections
	return rb
}

func (rb *TransportBuilder) TxCount(txcount int64) *TransportBuilder {
	rb.v.TxCount = &txcount
	return rb
}

func (rb *TransportBuilder) TxSize(txsize string) *TransportBuilder {
	rb.v.TxSize = &txsize
	return rb
}

func (rb *TransportBuilder) TxSizeInBytes(txsizeinbytes int64) *TransportBuilder {
	rb.v.TxSizeInBytes = &txsizeinbytes
	return rb
}

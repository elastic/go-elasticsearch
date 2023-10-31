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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Transport type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/_types/Stats.ts#L1047-L1090
type Transport struct {
	// InboundHandlingTimeHistogram The distribution of the time spent handling each inbound message on a
	// transport thread, represented as a histogram.
	InboundHandlingTimeHistogram []TransportHistogram `json:"inbound_handling_time_histogram,omitempty"`
	// OutboundHandlingTimeHistogram The distribution of the time spent sending each outbound transport message on
	// a transport thread, represented as a histogram.
	OutboundHandlingTimeHistogram []TransportHistogram `json:"outbound_handling_time_histogram,omitempty"`
	// RxCount Total number of RX (receive) packets received by the node during internal
	// cluster communication.
	RxCount *int64 `json:"rx_count,omitempty"`
	// RxSize Size of RX packets received by the node during internal cluster
	// communication.
	RxSize *string `json:"rx_size,omitempty"`
	// RxSizeInBytes Size, in bytes, of RX packets received by the node during internal cluster
	// communication.
	RxSizeInBytes *int64 `json:"rx_size_in_bytes,omitempty"`
	// ServerOpen Current number of inbound TCP connections used for internal communication
	// between nodes.
	ServerOpen *int `json:"server_open,omitempty"`
	// TotalOutboundConnections The cumulative number of outbound transport connections that this node has
	// opened since it started.
	// Each transport connection may comprise multiple TCP connections but is only
	// counted once in this statistic.
	// Transport connections are typically long-lived so this statistic should
	// remain constant in a stable cluster.
	TotalOutboundConnections *int64 `json:"total_outbound_connections,omitempty"`
	// TxCount Total number of TX (transmit) packets sent by the node during internal
	// cluster communication.
	TxCount *int64 `json:"tx_count,omitempty"`
	// TxSize Size of TX packets sent by the node during internal cluster communication.
	TxSize *string `json:"tx_size,omitempty"`
	// TxSizeInBytes Size, in bytes, of TX packets sent by the node during internal cluster
	// communication.
	TxSizeInBytes *int64 `json:"tx_size_in_bytes,omitempty"`
}

func (s *Transport) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "inbound_handling_time_histogram":
			if err := dec.Decode(&s.InboundHandlingTimeHistogram); err != nil {
				return err
			}

		case "outbound_handling_time_histogram":
			if err := dec.Decode(&s.OutboundHandlingTimeHistogram); err != nil {
				return err
			}

		case "rx_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RxCount = &value
			case float64:
				f := int64(v)
				s.RxCount = &f
			}

		case "rx_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RxSize = &o

		case "rx_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RxSizeInBytes = &value
			case float64:
				f := int64(v)
				s.RxSizeInBytes = &f
			}

		case "server_open":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ServerOpen = &value
			case float64:
				f := int(v)
				s.ServerOpen = &f
			}

		case "total_outbound_connections":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalOutboundConnections = &value
			case float64:
				f := int64(v)
				s.TotalOutboundConnections = &f
			}

		case "tx_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TxCount = &value
			case float64:
				f := int64(v)
				s.TxCount = &f
			}

		case "tx_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TxSize = &o

		case "tx_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TxSizeInBytes = &value
			case float64:
				f := int64(v)
				s.TxSizeInBytes = &f
			}

		}
	}
	return nil
}

// NewTransport returns a Transport.
func NewTransport() *Transport {
	r := &Transport{}

	return r
}

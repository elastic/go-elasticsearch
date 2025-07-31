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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"
)

// Stats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/_types/Stats.ts#L30-L114
type Stats struct {
	// AdaptiveSelection Statistics about adaptive replica selection.
	AdaptiveSelection map[string]AdaptiveSelection `json:"adaptive_selection,omitempty"`
	// Attributes Contains a list of attributes for the node.
	Attributes map[string]string `json:"attributes,omitempty"`
	// Breakers Statistics about the field data circuit breaker.
	Breakers map[string]Breaker `json:"breakers,omitempty"`
	// Discovery Contains node discovery statistics for the node.
	Discovery *Discovery `json:"discovery,omitempty"`
	// Fs File system information, data path, free disk space, read/write stats.
	Fs *FileSystem `json:"fs,omitempty"`
	// Host Network host for the node, based on the network host setting.
	Host *string `json:"host,omitempty"`
	// Http HTTP connection information.
	Http *Http `json:"http,omitempty"`
	// IndexingPressure Contains indexing pressure statistics for the node.
	IndexingPressure *NodesIndexingPressure `json:"indexing_pressure,omitempty"`
	// Indices Indices stats about size, document count, indexing and deletion times, search
	// times, field cache size, merges and flushes.
	Indices *IndicesShardStats `json:"indices,omitempty"`
	// Ingest Statistics about ingest preprocessing.
	Ingest *NodesIngest `json:"ingest,omitempty"`
	// Ip IP address and port for the node.
	Ip []string `json:"ip,omitempty"`
	// Jvm JVM stats, memory pool information, garbage collection, buffer pools, number
	// of loaded/unloaded classes.
	Jvm *Jvm `json:"jvm,omitempty"`
	// Name Human-readable identifier for the node.
	// Based on the node name setting.
	Name *string `json:"name,omitempty"`
	// Os Operating system stats, load average, mem, swap.
	Os *OperatingSystem `json:"os,omitempty"`
	// Process Process statistics, memory consumption, cpu usage, open file descriptors.
	Process *Process `json:"process,omitempty"`
	// Roles Roles assigned to the node.
	Roles []noderole.NodeRole `json:"roles,omitempty"`
	// Script Contains script statistics for the node.
	Script      *Scripting               `json:"script,omitempty"`
	ScriptCache map[string][]ScriptCache `json:"script_cache,omitempty"`
	// ThreadPool Statistics about each thread pool, including current size, queue and rejected
	// tasks.
	ThreadPool map[string]ThreadCount `json:"thread_pool,omitempty"`
	Timestamp  *int64                 `json:"timestamp,omitempty"`
	// Transport Transport statistics about sent and received bytes in cluster communication.
	Transport *Transport `json:"transport,omitempty"`
	// TransportAddress Host and port for the transport layer, used for internal communication
	// between nodes in a cluster.
	TransportAddress *string `json:"transport_address,omitempty"`
}

func (s *Stats) UnmarshalJSON(data []byte) error {

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

		case "adaptive_selection":
			if s.AdaptiveSelection == nil {
				s.AdaptiveSelection = make(map[string]AdaptiveSelection, 0)
			}
			if err := dec.Decode(&s.AdaptiveSelection); err != nil {
				return fmt.Errorf("%s | %w", "AdaptiveSelection", err)
			}

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return fmt.Errorf("%s | %w", "Attributes", err)
			}

		case "breakers":
			if s.Breakers == nil {
				s.Breakers = make(map[string]Breaker, 0)
			}
			if err := dec.Decode(&s.Breakers); err != nil {
				return fmt.Errorf("%s | %w", "Breakers", err)
			}

		case "discovery":
			if err := dec.Decode(&s.Discovery); err != nil {
				return fmt.Errorf("%s | %w", "Discovery", err)
			}

		case "fs":
			if err := dec.Decode(&s.Fs); err != nil {
				return fmt.Errorf("%s | %w", "Fs", err)
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}

		case "http":
			if err := dec.Decode(&s.Http); err != nil {
				return fmt.Errorf("%s | %w", "Http", err)
			}

		case "indexing_pressure":
			if err := dec.Decode(&s.IndexingPressure); err != nil {
				return fmt.Errorf("%s | %w", "IndexingPressure", err)
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return fmt.Errorf("%s | %w", "Indices", err)
			}

		case "ingest":
			if err := dec.Decode(&s.Ingest); err != nil {
				return fmt.Errorf("%s | %w", "Ingest", err)
			}

		case "ip":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Ip", err)
				}

				s.Ip = append(s.Ip, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Ip); err != nil {
					return fmt.Errorf("%s | %w", "Ip", err)
				}
			}

		case "jvm":
			if err := dec.Decode(&s.Jvm); err != nil {
				return fmt.Errorf("%s | %w", "Jvm", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "os":
			if err := dec.Decode(&s.Os); err != nil {
				return fmt.Errorf("%s | %w", "Os", err)
			}

		case "process":
			if err := dec.Decode(&s.Process); err != nil {
				return fmt.Errorf("%s | %w", "Process", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "script_cache":
			if s.ScriptCache == nil {
				s.ScriptCache = make(map[string][]ScriptCache, 0)
			}
			rawMsg := make(map[string]json.RawMessage, 0)
			dec.Decode(&rawMsg)
			for key, value := range rawMsg {
				switch {
				case bytes.HasPrefix(value, []byte("\"")), bytes.HasPrefix(value, []byte("{")):
					o := NewScriptCache()
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return fmt.Errorf("%s | %w", "ScriptCache", err)
					}
					s.ScriptCache[key] = append(s.ScriptCache[key], *o)
				default:
					o := []ScriptCache{}
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return fmt.Errorf("%s | %w", "ScriptCache", err)
					}
					s.ScriptCache[key] = o
				}
			}

		case "thread_pool":
			if s.ThreadPool == nil {
				s.ThreadPool = make(map[string]ThreadCount, 0)
			}
			if err := dec.Decode(&s.ThreadPool); err != nil {
				return fmt.Errorf("%s | %w", "ThreadPool", err)
			}

		case "timestamp":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Timestamp", err)
				}
				s.Timestamp = &value
			case float64:
				f := int64(v)
				s.Timestamp = &f
			}

		case "transport":
			if err := dec.Decode(&s.Transport); err != nil {
				return fmt.Errorf("%s | %w", "Transport", err)
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return fmt.Errorf("%s | %w", "TransportAddress", err)
			}

		}
	}
	return nil
}

// NewStats returns a Stats.
func NewStats() *Stats {
	r := &Stats{
		AdaptiveSelection: make(map[string]AdaptiveSelection),
		Attributes:        make(map[string]string),
		Breakers:          make(map[string]Breaker),
		ScriptCache:       make(map[string][]ScriptCache),
		ThreadPool:        make(map[string]ThreadCount),
	}

	return r
}

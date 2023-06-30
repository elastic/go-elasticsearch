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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"
)

// Stats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/_types/Stats.ts#L30-L53
type Stats struct {
	AdaptiveSelection map[string]AdaptiveSelection `json:"adaptive_selection,omitempty"`
	Attributes        map[string]string            `json:"attributes,omitempty"`
	Breakers          map[string]Breaker           `json:"breakers,omitempty"`
	Discovery         *Discovery                   `json:"discovery,omitempty"`
	Fs                *FileSystem                  `json:"fs,omitempty"`
	Host              *string                      `json:"host,omitempty"`
	Http              *Http                        `json:"http,omitempty"`
	IndexingPressure  *NodesIndexingPressure       `json:"indexing_pressure,omitempty"`
	Indices           *IndicesShardStats           `json:"indices,omitempty"`
	Ingest            *NodesIngest                 `json:"ingest,omitempty"`
	Ip                []string                     `json:"ip,omitempty"`
	Jvm               *Jvm                         `json:"jvm,omitempty"`
	Name              *string                      `json:"name,omitempty"`
	Os                *OperatingSystem             `json:"os,omitempty"`
	Process           *Process                     `json:"process,omitempty"`
	Roles             []noderole.NodeRole          `json:"roles,omitempty"`
	Script            *Scripting                   `json:"script,omitempty"`
	ScriptCache       map[string][]ScriptCache     `json:"script_cache,omitempty"`
	ThreadPool        map[string]ThreadCount       `json:"thread_pool,omitempty"`
	Timestamp         *int64                       `json:"timestamp,omitempty"`
	Transport         *Transport                   `json:"transport,omitempty"`
	TransportAddress  *string                      `json:"transport_address,omitempty"`
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
				return err
			}

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return err
			}

		case "breakers":
			if s.Breakers == nil {
				s.Breakers = make(map[string]Breaker, 0)
			}
			if err := dec.Decode(&s.Breakers); err != nil {
				return err
			}

		case "discovery":
			if err := dec.Decode(&s.Discovery); err != nil {
				return err
			}

		case "fs":
			if err := dec.Decode(&s.Fs); err != nil {
				return err
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return err
			}

		case "http":
			if err := dec.Decode(&s.Http); err != nil {
				return err
			}

		case "indexing_pressure":
			if err := dec.Decode(&s.IndexingPressure); err != nil {
				return err
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "ingest":
			if err := dec.Decode(&s.Ingest); err != nil {
				return err
			}

		case "ip":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Ip = append(s.Ip, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Ip); err != nil {
					return err
				}
			}

		case "jvm":
			if err := dec.Decode(&s.Jvm); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "os":
			if err := dec.Decode(&s.Os); err != nil {
				return err
			}

		case "process":
			if err := dec.Decode(&s.Process); err != nil {
				return err
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
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
						return err
					}
					s.ScriptCache[key] = append(s.ScriptCache[key], *o)
				default:
					o := []ScriptCache{}
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return err
					}
					s.ScriptCache[key] = o
				}
			}

		case "thread_pool":
			if s.ThreadPool == nil {
				s.ThreadPool = make(map[string]ThreadCount, 0)
			}
			if err := dec.Decode(&s.ThreadPool); err != nil {
				return err
			}

		case "timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Timestamp = &value
			case float64:
				f := int64(v)
				s.Timestamp = &f
			}

		case "transport":
			if err := dec.Decode(&s.Transport); err != nil {
				return err
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewStats returns a Stats.
func NewStats() *Stats {
	r := &Stats{
		AdaptiveSelection: make(map[string]AdaptiveSelection, 0),
		Attributes:        make(map[string]string, 0),
		Breakers:          make(map[string]Breaker, 0),
		ScriptCache:       make(map[string][]ScriptCache, 0),
		ThreadPool:        make(map[string]ThreadCount, 0),
	}

	return r
}

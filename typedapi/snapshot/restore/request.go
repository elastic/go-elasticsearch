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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package restore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package restore
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/snapshot/restore/SnapshotRestoreRequest.ts#L25-L175
type Request struct {

	// FeatureStates The feature states to restore.
	// If `include_global_state` is `true`, the request restores all feature states
	// in the snapshot by default.
	// If `include_global_state` is `false`, the request restores no feature states
	// by default.
	// Note that specifying an empty array will result in the default behavior.
	// To restore no feature states, regardless of the `include_global_state` value,
	// specify an array containing only the value `none` (`["none"]`).
	FeatureStates []string `json:"feature_states,omitempty"`
	// IgnoreIndexSettings The index settings to not restore from the snapshot.
	// You can't use this option to ignore `index.number_of_shards`.
	//
	// For data streams, this option applies only to restored backing indices.
	// New backing indices are configured using the data stream's matching index
	// template.
	IgnoreIndexSettings []string `json:"ignore_index_settings,omitempty"`
	// IgnoreUnavailable If `true`, the request ignores any index or data stream in indices that's
	// missing from the snapshot.
	// If `false`, the request returns an error for any missing index or data
	// stream.
	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`
	// IncludeAliases If `true`, the request restores aliases for any restored data streams and
	// indices.
	// If `false`, the request doesn’t restore aliases.
	IncludeAliases *bool `json:"include_aliases,omitempty"`
	// IncludeGlobalState If `true`, restore the cluster state. The cluster state includes:
	//
	// * Persistent cluster settings
	// * Index templates
	// * Legacy index templates
	// * Ingest pipelines
	// * Index lifecycle management (ILM) policies
	// * Stored scripts
	// * For snapshots taken after 7.12.0, feature states
	//
	// If `include_global_state` is `true`, the restore operation merges the legacy
	// index templates in your cluster with the templates contained in the snapshot,
	// replacing any existing ones whose name matches one in the snapshot.
	// It completely removes all persistent settings, non-legacy index templates,
	// ingest pipelines, and ILM lifecycle policies that exist in your cluster and
	// replaces them with the corresponding items from the snapshot.
	//
	// Use the `feature_states` parameter to configure how feature states are
	// restored.
	//
	// If `include_global_state` is `true` and a snapshot was created without a
	// global state then the restore request will fail.
	IncludeGlobalState *bool `json:"include_global_state,omitempty"`
	// IndexSettings Index settings to add or change in restored indices, including backing
	// indices.
	// You can't use this option to change `index.number_of_shards`.
	//
	// For data streams, this option applies only to restored backing indices.
	// New backing indices are configured using the data stream's matching index
	// template.
	IndexSettings *types.IndexSettings `json:"index_settings,omitempty"`
	// Indices A comma-separated list of indices and data streams to restore.
	// It supports a multi-target syntax.
	// The default behavior is all regular indices and regular data streams in the
	// snapshot.
	//
	// You can't use this parameter to restore system indices or system data
	// streams.
	// Use `feature_states` instead.
	Indices []string `json:"indices,omitempty"`
	// Partial If `false`, the entire restore operation will fail if one or more indices
	// included in the snapshot do not have all primary shards available.
	//
	// If true, it allows restoring a partial snapshot of indices with unavailable
	// shards.
	// Only shards that were successfully included in the snapshot will be restored.
	// All missing shards will be recreated as empty.
	Partial *bool `json:"partial,omitempty"`
	// RenamePattern A rename pattern to apply to restored data streams and indices.
	// Data streams and indices matching the rename pattern will be renamed
	// according to `rename_replacement`.
	//
	// The rename pattern is applied as defined by the regular expression that
	// supports referencing the original text, according to the `appendReplacement`
	// logic.
	RenamePattern *string `json:"rename_pattern,omitempty"`
	// RenameReplacement The rename replacement string that is used with the `rename_pattern`.
	RenameReplacement *string `json:"rename_replacement,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Restore request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "feature_states":
			if err := dec.Decode(&s.FeatureStates); err != nil {
				return fmt.Errorf("%s | %w", "FeatureStates", err)
			}

		case "ignore_index_settings":
			if err := dec.Decode(&s.IgnoreIndexSettings); err != nil {
				return fmt.Errorf("%s | %w", "IgnoreIndexSettings", err)
			}

		case "ignore_unavailable":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreUnavailable", err)
				}
				s.IgnoreUnavailable = &value
			case bool:
				s.IgnoreUnavailable = &v
			}

		case "include_aliases":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeAliases", err)
				}
				s.IncludeAliases = &value
			case bool:
				s.IncludeAliases = &v
			}

		case "include_global_state":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeGlobalState", err)
				}
				s.IncludeGlobalState = &value
			case bool:
				s.IncludeGlobalState = &v
			}

		case "index_settings":
			if err := dec.Decode(&s.IndexSettings); err != nil {
				return fmt.Errorf("%s | %w", "IndexSettings", err)
			}

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Indices", err)
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return fmt.Errorf("%s | %w", "Indices", err)
				}
			}

		case "partial":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Partial", err)
				}
				s.Partial = &value
			case bool:
				s.Partial = &v
			}

		case "rename_pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RenamePattern", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RenamePattern = &o

		case "rename_replacement":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RenameReplacement", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RenameReplacement = &o

		}
	}
	return nil
}

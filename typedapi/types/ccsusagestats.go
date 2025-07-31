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
)

// CCSUsageStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L819-L844
type CCSUsageStats struct {
	// Clients Statistics about the clients that executed cross-cluster search requests. The
	// keys are the names of the clients, and the values are the number of requests
	// that were executed by that client. Only known clients (such as `kibana` or
	// `elasticsearch`) are counted.
	Clients map[string]int `json:"clients"`
	// Clusters Statistics about the clusters that were queried in cross-cluster search
	// requests. The keys are cluster names, and the values are per-cluster
	// telemetry data. This also includes the local cluster itself, which uses the
	// name `(local)`.
	Clusters map[string]CCSUsageClusterStats `json:"clusters"`
	// FailureReasons Statistics about the reasons for cross-cluster search request failures. The
	// keys are the failure reason names and the values are the number of requests
	// that failed for that reason.
	FailureReasons map[string]int `json:"failure_reasons"`
	// Features The keys are the names of the search feature, and the values are the number
	// of requests that used that feature. Single request can use more than one
	// feature (e.g. both `async` and `wildcard`).
	Features map[string]int `json:"features"`
	// RemotesPerSearchAvg The average number of remote clusters that were queried in a single
	// cross-cluster search request.
	RemotesPerSearchAvg Float64 `json:"remotes_per_search_avg"`
	// RemotesPerSearchMax The maximum number of remote clusters that were queried in a single
	// cross-cluster search request.
	RemotesPerSearchMax int `json:"remotes_per_search_max"`
	// Skipped The total number of cross-cluster search requests (successful or failed) that
	// had at least one remote cluster skipped.
	Skipped int `json:"skipped"`
	// Success The total number of cross-cluster search requests that have been successfully
	// executed by the cluster.
	Success int `json:"success"`
	// Took Statistics about the time taken to execute cross-cluster search requests.
	Took CCSUsageTimeValue `json:"took"`
	// TookMrtFalse Statistics about the time taken to execute cross-cluster search requests for
	// which the `ccs_minimize_roundtrips` setting was set to `false`.
	TookMrtFalse *CCSUsageTimeValue `json:"took_mrt_false,omitempty"`
	// TookMrtTrue Statistics about the time taken to execute cross-cluster search requests for
	// which the `ccs_minimize_roundtrips` setting was set to `true`.
	TookMrtTrue *CCSUsageTimeValue `json:"took_mrt_true,omitempty"`
	// Total The total number of cross-cluster search requests that have been executed by
	// the cluster.
	Total int `json:"total"`
}

func (s *CCSUsageStats) UnmarshalJSON(data []byte) error {

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

		case "clients":
			if s.Clients == nil {
				s.Clients = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.Clients); err != nil {
				return fmt.Errorf("%s | %w", "Clients", err)
			}

		case "clusters":
			if s.Clusters == nil {
				s.Clusters = make(map[string]CCSUsageClusterStats, 0)
			}
			if err := dec.Decode(&s.Clusters); err != nil {
				return fmt.Errorf("%s | %w", "Clusters", err)
			}

		case "failure_reasons":
			if s.FailureReasons == nil {
				s.FailureReasons = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.FailureReasons); err != nil {
				return fmt.Errorf("%s | %w", "FailureReasons", err)
			}

		case "features":
			if s.Features == nil {
				s.Features = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.Features); err != nil {
				return fmt.Errorf("%s | %w", "Features", err)
			}

		case "remotes_per_search_avg":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RemotesPerSearchAvg", err)
				}
				f := Float64(value)
				s.RemotesPerSearchAvg = f
			case float64:
				f := Float64(v)
				s.RemotesPerSearchAvg = f
			}

		case "remotes_per_search_max":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RemotesPerSearchMax", err)
				}
				s.RemotesPerSearchMax = value
			case float64:
				f := int(v)
				s.RemotesPerSearchMax = f
			}

		case "skipped":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Skipped", err)
				}
				s.Skipped = value
			case float64:
				f := int(v)
				s.Skipped = f
			}

		case "success":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Success", err)
				}
				s.Success = value
			case float64:
				f := int(v)
				s.Success = f
			}

		case "took":
			if err := dec.Decode(&s.Took); err != nil {
				return fmt.Errorf("%s | %w", "Took", err)
			}

		case "took_mrt_false":
			if err := dec.Decode(&s.TookMrtFalse); err != nil {
				return fmt.Errorf("%s | %w", "TookMrtFalse", err)
			}

		case "took_mrt_true":
			if err := dec.Decode(&s.TookMrtTrue); err != nil {
				return fmt.Errorf("%s | %w", "TookMrtTrue", err)
			}

		case "total":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewCCSUsageStats returns a CCSUsageStats.
func NewCCSUsageStats() *CCSUsageStats {
	r := &CCSUsageStats{
		Clients:        make(map[string]int),
		Clusters:       make(map[string]CCSUsageClusterStats),
		FailureReasons: make(map[string]int),
		Features:       make(map[string]int),
	}

	return r
}

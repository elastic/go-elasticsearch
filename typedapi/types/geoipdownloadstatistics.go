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

// GeoIpDownloadStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ingest/geo_ip_stats/types.ts#L24-L35
type GeoIpDownloadStatistics struct {
	// DatabaseCount Current number of databases available for use.
	DatabaseCount int `json:"database_count"`
	// FailedDownloads Total number of failed database downloads.
	FailedDownloads int `json:"failed_downloads"`
	// SkippedUpdates Total number of database updates skipped.
	SkippedUpdates int `json:"skipped_updates"`
	// SuccessfulDownloads Total number of successful database downloads.
	SuccessfulDownloads int `json:"successful_downloads"`
	// TotalDownloadTime Total milliseconds spent downloading databases.
	TotalDownloadTime int64 `json:"total_download_time"`
}

func (s *GeoIpDownloadStatistics) UnmarshalJSON(data []byte) error {

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

		case "database_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DatabaseCount = value
			case float64:
				f := int(v)
				s.DatabaseCount = f
			}

		case "failed_downloads":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FailedDownloads = value
			case float64:
				f := int(v)
				s.FailedDownloads = f
			}

		case "skipped_updates":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SkippedUpdates = value
			case float64:
				f := int(v)
				s.SkippedUpdates = f
			}

		case "successful_downloads":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SuccessfulDownloads = value
			case float64:
				f := int(v)
				s.SuccessfulDownloads = f
			}

		case "total_download_time":
			if err := dec.Decode(&s.TotalDownloadTime); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewGeoIpDownloadStatistics returns a GeoIpDownloadStatistics.
func NewGeoIpDownloadStatistics() *GeoIpDownloadStatistics {
	r := &GeoIpDownloadStatistics{}

	return r
}

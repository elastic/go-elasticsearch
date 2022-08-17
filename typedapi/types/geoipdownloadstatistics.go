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

// GeoIpDownloadStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/geo_ip_stats/types.ts#L24-L35
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
	TotalDownloadTime DurationValueUnitMillis `json:"total_download_time"`
}

// GeoIpDownloadStatisticsBuilder holds GeoIpDownloadStatistics struct and provides a builder API.
type GeoIpDownloadStatisticsBuilder struct {
	v *GeoIpDownloadStatistics
}

// NewGeoIpDownloadStatistics provides a builder for the GeoIpDownloadStatistics struct.
func NewGeoIpDownloadStatisticsBuilder() *GeoIpDownloadStatisticsBuilder {
	r := GeoIpDownloadStatisticsBuilder{
		&GeoIpDownloadStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the GeoIpDownloadStatistics struct
func (rb *GeoIpDownloadStatisticsBuilder) Build() GeoIpDownloadStatistics {
	return *rb.v
}

// DatabaseCount Current number of databases available for use.

func (rb *GeoIpDownloadStatisticsBuilder) DatabaseCount(databasecount int) *GeoIpDownloadStatisticsBuilder {
	rb.v.DatabaseCount = databasecount
	return rb
}

// FailedDownloads Total number of failed database downloads.

func (rb *GeoIpDownloadStatisticsBuilder) FailedDownloads(faileddownloads int) *GeoIpDownloadStatisticsBuilder {
	rb.v.FailedDownloads = faileddownloads
	return rb
}

// SkippedUpdates Total number of database updates skipped.

func (rb *GeoIpDownloadStatisticsBuilder) SkippedUpdates(skippedupdates int) *GeoIpDownloadStatisticsBuilder {
	rb.v.SkippedUpdates = skippedupdates
	return rb
}

// SuccessfulDownloads Total number of successful database downloads.

func (rb *GeoIpDownloadStatisticsBuilder) SuccessfulDownloads(successfuldownloads int) *GeoIpDownloadStatisticsBuilder {
	rb.v.SuccessfulDownloads = successfuldownloads
	return rb
}

// TotalDownloadTime Total milliseconds spent downloading databases.

func (rb *GeoIpDownloadStatisticsBuilder) TotalDownloadTime(totaldownloadtime *DurationValueUnitMillisBuilder) *GeoIpDownloadStatisticsBuilder {
	v := totaldownloadtime.Build()
	rb.v.TotalDownloadTime = v
	return rb
}

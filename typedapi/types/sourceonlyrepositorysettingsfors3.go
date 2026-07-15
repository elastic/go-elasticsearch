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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SourceOnlyRepositorySettingsForS3 type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/snapshot/_types/SnapshotRepository.ts#L452-L456
type SourceOnlyRepositorySettingsForS3 struct {
	// BasePath The path to the repository data within its bucket. It defaults to an empty
	// string, meaning that the repository is at the root of the bucket. The value
	// of this setting should not start or end with a forward slash (`/`).
	//
	// NOTE: Don't set base_path when configuring a snapshot repository for Elastic
	// Cloud Enterprise. Elastic Cloud Enterprise automatically generates the
	// `base_path` for each deployment so that multiple deployments may share the
	// same bucket.
	BasePath *string `json:"base_path,omitempty"`
	// Bucket The name of the S3 bucket to use for snapshots. The bucket name must adhere
	// to Amazon's S3 bucket naming rules.
	Bucket string `json:"bucket"`
	// BufferSize The minimum threshold below which the chunk is uploaded using a single
	// request. Beyond this threshold, the S3 repository will use the AWS Multipart
	// Upload API to split the chunk into several parts, each of `buffer_size`
	// length, and to upload each part in its own request. Note that setting a
	// buffer size lower than 5mb is not allowed since it will prevent the use of
	// the Multipart API and may result in upload errors. It is also not possible to
	// set a buffer size greater than 5gb as it is the maximum upload size allowed
	// by S3. Defaults to `100mb` or 5% of JVM heap, whichever is smaller.
	BufferSize ByteSize `json:"buffer_size,omitempty"`
	// CannedAcl The S3 repository supports all S3 canned ACLs: `private`, `public-read`,
	// `public-read-write`, `authenticated-read`, `log-delivery-write`,
	// `bucket-owner-read`, `bucket-owner-full-control`. You could specify a canned
	// ACL using the `canned_acl` setting. When the S3 repository creates buckets
	// and objects, it adds the canned ACL into the buckets and objects.
	CannedAcl *string `json:"canned_acl,omitempty"`
	// Client The name of the S3 client to use to connect to S3.
	Client       *string `json:"client,omitempty"`
	DelegateType string  `json:"delegate_type,omitempty"`
	// DeleteObjectsMaxSize The maxmimum batch size, between 1 and 1000, used for `DeleteObjects`
	// requests. Defaults to 1000 which is the maximum number supported by the AWS
	// DeleteObjects API.
	DeleteObjectsMaxSize *int `json:"delete_objects_max_size,omitempty"`
	// GetRegisterRetryDelay The time to wait before trying again if an attempt to read a linearizable
	// register fails.
	GetRegisterRetryDelay Duration `json:"get_register_retry_delay,omitempty"`
	// MaxMultipartParts The maximum number of parts that Elasticsearch will write during a multipart
	// upload of a single object. Files which are larger than `buffer_size ×
	// max_multipart_parts` will be chunked into several smaller objects.
	// Elasticsearch may also split a file across multiple objects to satisfy other
	// constraints such as the `chunk_size` limit. Defaults to `10000` which is the
	// maximum number of parts in a multipart upload in AWS S3.
	MaxMultipartParts *int `json:"max_multipart_parts,omitempty"`
	// MaxMultipartUploadCleanupSize The maximum number of possibly-dangling multipart uploads to clean up in each
	// batch of snapshot deletions. Defaults to 1000 which is the maximum number
	// supported by the AWS ListMultipartUploads API. If set to `0`, Elasticsearch
	// will not attempt to clean up dangling multipart uploads.
	MaxMultipartUploadCleanupSize *int `json:"max_multipart_upload_cleanup_size,omitempty"`
	// Readonly If true, the repository is read-only. The cluster can retrieve and restore
	// snapshots from the repository but not write to the repository or create
	// snapshots in it.
	//
	// Only a cluster with write access can create snapshots in the repository. All
	// other clusters connected to the repository should have the `readonly`
	// parameter set to `true`.
	//
	// If `false`, the cluster can write to the repository and create snapshots in
	// it.
	//
	// IMPORTANT: If you register the same snapshot repository with multiple
	// clusters, only one cluster should have write access to the repository. Having
	// multiple clusters write to the repository at the same time risks corrupting
	// the contents of the repository.
	Readonly *bool `json:"readonly,omitempty"`
	// ServerSideEncryption When set to `true`, files are encrypted on server side using an AES256
	// algorithm.
	ServerSideEncryption *bool `json:"server_side_encryption,omitempty"`
	// StorageClass The S3 storage class for objects written to the repository. Values may be
	// `standard`, `reduced_redundancy`, `standard_ia`, `onezone_ia`, and
	// `intelligent_tiering`.
	StorageClass *string `json:"storage_class,omitempty"`
	// ThrottledDeleteRetryDelayIncrement The delay before the first retry and the amount the delay is incremented by
	// on each subsequent retry. The default is 50ms and the minimum is 0ms.
	ThrottledDeleteRetryDelayIncrement Duration `json:"throttled_delete_retry.delay_increment,omitempty"`
	// ThrottledDeleteRetryMaximumDelay The upper bound on how long the delays between retries will grow to. The
	// default is 5s and the minimum is 0ms.
	ThrottledDeleteRetryMaximumDelay Duration `json:"throttled_delete_retry.maximum_delay,omitempty"`
	// ThrottledDeleteRetryMaximumNumberOfRetries The number times to retry a throttled snapshot deletion. The default is 10
	// and the minimum value is 0 which will disable retries altogether. Note that
	// if retries are enabled in the Azure client, each of these retries comprises
	// that many client-level retries.
	ThrottledDeleteRetryMaximumNumberOfRetries *int `json:"throttled_delete_retry.maximum_number_of_retries,omitempty"`
}

func (s *SourceOnlyRepositorySettingsForS3) UnmarshalJSON(data []byte) error {

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

		case "base_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "BasePath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BasePath = &o

		case "bucket":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Bucket", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Bucket = o

		case "buffer_size":
			if err := dec.Decode(&s.BufferSize); err != nil {
				return fmt.Errorf("%s | %w", "BufferSize", err)
			}

		case "canned_acl":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "CannedAcl", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CannedAcl = &o

		case "client":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Client", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Client = &o

		case "delegate_type":
			if err := dec.Decode(&s.DelegateType); err != nil {
				return fmt.Errorf("%s | %w", "DelegateType", err)
			}

		case "delete_objects_max_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DeleteObjectsMaxSize", err)
				}
				s.DeleteObjectsMaxSize = &value
			case float64:
				f := int(v)
				s.DeleteObjectsMaxSize = &f
			}

		case "get_register_retry_delay":
			if err := dec.Decode(&s.GetRegisterRetryDelay); err != nil {
				return fmt.Errorf("%s | %w", "GetRegisterRetryDelay", err)
			}

		case "max_multipart_parts":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxMultipartParts", err)
				}
				s.MaxMultipartParts = &value
			case float64:
				f := int(v)
				s.MaxMultipartParts = &f
			}

		case "max_multipart_upload_cleanup_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxMultipartUploadCleanupSize", err)
				}
				s.MaxMultipartUploadCleanupSize = &value
			case float64:
				f := int(v)
				s.MaxMultipartUploadCleanupSize = &f
			}

		case "readonly":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Readonly", err)
				}
				s.Readonly = &value
			case bool:
				s.Readonly = &v
			}

		case "server_side_encryption":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ServerSideEncryption", err)
				}
				s.ServerSideEncryption = &value
			case bool:
				s.ServerSideEncryption = &v
			}

		case "storage_class":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "StorageClass", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StorageClass = &o

		case "throttled_delete_retry.delay_increment":
			if err := dec.Decode(&s.ThrottledDeleteRetryDelayIncrement); err != nil {
				return fmt.Errorf("%s | %w", "ThrottledDeleteRetryDelayIncrement", err)
			}

		case "throttled_delete_retry.maximum_delay":
			if err := dec.Decode(&s.ThrottledDeleteRetryMaximumDelay); err != nil {
				return fmt.Errorf("%s | %w", "ThrottledDeleteRetryMaximumDelay", err)
			}

		case "throttled_delete_retry.maximum_number_of_retries":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ThrottledDeleteRetryMaximumNumberOfRetries", err)
				}
				s.ThrottledDeleteRetryMaximumNumberOfRetries = &value
			case float64:
				f := int(v)
				s.ThrottledDeleteRetryMaximumNumberOfRetries = &f
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s SourceOnlyRepositorySettingsForS3) MarshalJSON() ([]byte, error) {
	type innerSourceOnlyRepositorySettingsForS3 SourceOnlyRepositorySettingsForS3
	tmp := innerSourceOnlyRepositorySettingsForS3{
		BasePath:                           s.BasePath,
		Bucket:                             s.Bucket,
		BufferSize:                         s.BufferSize,
		CannedAcl:                          s.CannedAcl,
		Client:                             s.Client,
		DelegateType:                       s.DelegateType,
		DeleteObjectsMaxSize:               s.DeleteObjectsMaxSize,
		GetRegisterRetryDelay:              s.GetRegisterRetryDelay,
		MaxMultipartParts:                  s.MaxMultipartParts,
		MaxMultipartUploadCleanupSize:      s.MaxMultipartUploadCleanupSize,
		Readonly:                           s.Readonly,
		ServerSideEncryption:               s.ServerSideEncryption,
		StorageClass:                       s.StorageClass,
		ThrottledDeleteRetryDelayIncrement: s.ThrottledDeleteRetryDelayIncrement,
		ThrottledDeleteRetryMaximumDelay:   s.ThrottledDeleteRetryMaximumDelay,
		ThrottledDeleteRetryMaximumNumberOfRetries: s.ThrottledDeleteRetryMaximumNumberOfRetries,
	}

	tmp.DelegateType = "s3"

	return json.Marshal(tmp)
}

// NewSourceOnlyRepositorySettingsForS3 returns a SourceOnlyRepositorySettingsForS3.
func NewSourceOnlyRepositorySettingsForS3() *SourceOnlyRepositorySettingsForS3 {
	r := &SourceOnlyRepositorySettingsForS3{}

	return r
}

type SourceOnlyRepositorySettingsForS3Variant interface {
	SourceOnlyRepositorySettingsForS3Caster() *SourceOnlyRepositorySettingsForS3
}

func (s *SourceOnlyRepositorySettingsForS3) SourceOnlyRepositorySettingsForS3Caster() *SourceOnlyRepositorySettingsForS3 {
	return s
}

func (s *SourceOnlyRepositorySettingsForS3) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	if s == nil {
		return nil
	}
	o := SourceOnlyRepositorySettings(s)
	return &o
}

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

// RequestCounts type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/_types/RepositoryMeteringInformation.ts#L76-L103
type RequestCounts struct {
	// GetBlob Number of Get Blob requests (Azure)
	GetBlob *int64 `json:"GetBlob,omitempty"`
	// GetBlobProperties Number of Get Blob Properties requests (Azure)
	GetBlobProperties *int64 `json:"GetBlobProperties,omitempty"`
	// GetObject Number of get object requests (GCP, S3)
	GetObject *int64 `json:"GetObject,omitempty"`
	// InsertObject Number of insert object requests, including simple, multipart and resumable
	// uploads. Resumable uploads
	// can perform multiple http requests to insert a single object but they are
	// considered as a single request
	// since they are billed as an individual operation. (GCP)
	InsertObject *int64 `json:"InsertObject,omitempty"`
	// ListBlobs Number of List Blobs requests (Azure)
	ListBlobs *int64 `json:"ListBlobs,omitempty"`
	// ListObjects Number of list objects requests (GCP, S3)
	ListObjects *int64 `json:"ListObjects,omitempty"`
	// PutBlob Number of Put Blob requests (Azure)
	PutBlob *int64 `json:"PutBlob,omitempty"`
	// PutBlock Number of Put Block (Azure)
	PutBlock *int64 `json:"PutBlock,omitempty"`
	// PutBlockList Number of Put Block List requests
	PutBlockList *int64 `json:"PutBlockList,omitempty"`
	// PutMultipartObject Number of Multipart requests, including CreateMultipartUpload, UploadPart and
	// CompleteMultipartUpload requests (S3)
	PutMultipartObject *int64 `json:"PutMultipartObject,omitempty"`
	// PutObject Number of PutObject requests (S3)
	PutObject *int64 `json:"PutObject,omitempty"`
}

func (s *RequestCounts) UnmarshalJSON(data []byte) error {

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

		case "GetBlob":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.GetBlob = &value
			case float64:
				f := int64(v)
				s.GetBlob = &f
			}

		case "GetBlobProperties":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.GetBlobProperties = &value
			case float64:
				f := int64(v)
				s.GetBlobProperties = &f
			}

		case "GetObject":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.GetObject = &value
			case float64:
				f := int64(v)
				s.GetObject = &f
			}

		case "InsertObject":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.InsertObject = &value
			case float64:
				f := int64(v)
				s.InsertObject = &f
			}

		case "ListBlobs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ListBlobs = &value
			case float64:
				f := int64(v)
				s.ListBlobs = &f
			}

		case "ListObjects":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ListObjects = &value
			case float64:
				f := int64(v)
				s.ListObjects = &f
			}

		case "PutBlob":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PutBlob = &value
			case float64:
				f := int64(v)
				s.PutBlob = &f
			}

		case "PutBlock":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PutBlock = &value
			case float64:
				f := int64(v)
				s.PutBlock = &f
			}

		case "PutBlockList":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PutBlockList = &value
			case float64:
				f := int64(v)
				s.PutBlockList = &f
			}

		case "PutMultipartObject":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PutMultipartObject = &value
			case float64:
				f := int64(v)
				s.PutMultipartObject = &f
			}

		case "PutObject":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PutObject = &value
			case float64:
				f := int64(v)
				s.PutObject = &f
			}

		}
	}
	return nil
}

// NewRequestCounts returns a RequestCounts.
func NewRequestCounts() *RequestCounts {
	r := &RequestCounts{}

	return r
}

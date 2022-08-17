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

// RequestCounts type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/RepositoryMeteringInformation.ts#L76-L103
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

// RequestCountsBuilder holds RequestCounts struct and provides a builder API.
type RequestCountsBuilder struct {
	v *RequestCounts
}

// NewRequestCounts provides a builder for the RequestCounts struct.
func NewRequestCountsBuilder() *RequestCountsBuilder {
	r := RequestCountsBuilder{
		&RequestCounts{},
	}

	return &r
}

// Build finalize the chain and returns the RequestCounts struct
func (rb *RequestCountsBuilder) Build() RequestCounts {
	return *rb.v
}

// GetBlob Number of Get Blob requests (Azure)

func (rb *RequestCountsBuilder) GetBlob(getblob int64) *RequestCountsBuilder {
	rb.v.GetBlob = &getblob
	return rb
}

// GetBlobProperties Number of Get Blob Properties requests (Azure)

func (rb *RequestCountsBuilder) GetBlobProperties(getblobproperties int64) *RequestCountsBuilder {
	rb.v.GetBlobProperties = &getblobproperties
	return rb
}

// GetObject Number of get object requests (GCP, S3)

func (rb *RequestCountsBuilder) GetObject(getobject int64) *RequestCountsBuilder {
	rb.v.GetObject = &getobject
	return rb
}

// InsertObject Number of insert object requests, including simple, multipart and resumable
// uploads. Resumable uploads
// can perform multiple http requests to insert a single object but they are
// considered as a single request
// since they are billed as an individual operation. (GCP)

func (rb *RequestCountsBuilder) InsertObject(insertobject int64) *RequestCountsBuilder {
	rb.v.InsertObject = &insertobject
	return rb
}

// ListBlobs Number of List Blobs requests (Azure)

func (rb *RequestCountsBuilder) ListBlobs(listblobs int64) *RequestCountsBuilder {
	rb.v.ListBlobs = &listblobs
	return rb
}

// ListObjects Number of list objects requests (GCP, S3)

func (rb *RequestCountsBuilder) ListObjects(listobjects int64) *RequestCountsBuilder {
	rb.v.ListObjects = &listobjects
	return rb
}

// PutBlob Number of Put Blob requests (Azure)

func (rb *RequestCountsBuilder) PutBlob(putblob int64) *RequestCountsBuilder {
	rb.v.PutBlob = &putblob
	return rb
}

// PutBlock Number of Put Block (Azure)

func (rb *RequestCountsBuilder) PutBlock(putblock int64) *RequestCountsBuilder {
	rb.v.PutBlock = &putblock
	return rb
}

// PutBlockList Number of Put Block List requests

func (rb *RequestCountsBuilder) PutBlockList(putblocklist int64) *RequestCountsBuilder {
	rb.v.PutBlockList = &putblocklist
	return rb
}

// PutMultipartObject Number of Multipart requests, including CreateMultipartUpload, UploadPart and
// CompleteMultipartUpload requests (S3)

func (rb *RequestCountsBuilder) PutMultipartObject(putmultipartobject int64) *RequestCountsBuilder {
	rb.v.PutMultipartObject = &putmultipartobject
	return rb
}

// PutObject Number of PutObject requests (S3)

func (rb *RequestCountsBuilder) PutObject(putobject int64) *RequestCountsBuilder {
	rb.v.PutObject = &putobject
	return rb
}

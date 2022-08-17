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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/storagetype"
)

// Storage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L496-L505
type Storage struct {
	// AllowMmap You can restrict the use of the mmapfs and the related hybridfs store type
	// via the setting node.store.allow_mmap.
	// This is a boolean setting indicating whether or not memory-mapping is
	// allowed. The default is to allow it. This
	// setting is useful, for example, if you are in an environment where you can
	// not control the ability to create a lot
	// of memory maps so you need disable the ability to use memory-mapping.
	AllowMmap *bool                   `json:"allow_mmap,omitempty"`
	Type      storagetype.StorageType `json:"type"`
}

// StorageBuilder holds Storage struct and provides a builder API.
type StorageBuilder struct {
	v *Storage
}

// NewStorage provides a builder for the Storage struct.
func NewStorageBuilder() *StorageBuilder {
	r := StorageBuilder{
		&Storage{},
	}

	return &r
}

// Build finalize the chain and returns the Storage struct
func (rb *StorageBuilder) Build() Storage {
	return *rb.v
}

// AllowMmap You can restrict the use of the mmapfs and the related hybridfs store type
// via the setting node.store.allow_mmap.
// This is a boolean setting indicating whether or not memory-mapping is
// allowed. The default is to allow it. This
// setting is useful, for example, if you are in an environment where you can
// not control the ability to create a lot
// of memory maps so you need disable the ability to use memory-mapping.

func (rb *StorageBuilder) AllowMmap(allowmmap bool) *StorageBuilder {
	rb.v.AllowMmap = &allowmmap
	return rb
}

func (rb *StorageBuilder) Type_(type_ storagetype.StorageType) *StorageBuilder {
	rb.v.Type = type_
	return rb
}

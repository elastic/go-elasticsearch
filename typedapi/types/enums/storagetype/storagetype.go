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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package storagetype
package storagetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/indices/_types/IndexSettings.ts#L570-L598
type StorageType struct {
	Name string
}

var (

	// Fs Default file system implementation. This will pick the best implementation
	// depending on the operating environment, which is currently hybridfs on all
	// supported systems but is subject to change.
	Fs = StorageType{"fs"}

	// Niofs The NIO FS type stores the shard index on the file system (maps to Lucene
	// NIOFSDirectory) using NIO. It allows multiple threads to read from the same
	// file concurrently. It is not recommended on Windows because of a bug in the
	// SUN Java implementation and disables some optimizations for heap memory
	// usage.
	Niofs = StorageType{"niofs"}

	// Mmapfs The MMap FS type stores the shard index on the file system (maps to Lucene
	// MMapDirectory) by mapping a file into memory (mmap). Memory mapping uses up a
	// portion of the virtual memory address space in your process equal to the size
	// of the file being mapped. Before using this class, be sure you have allowed
	// plenty of virtual address space.
	Mmapfs = StorageType{"mmapfs"}

	// Hybridfs The hybridfs type is a hybrid of niofs and mmapfs, which chooses the best
	// file system type for each type of file based on the read access pattern.
	// Currently only the Lucene term dictionary, norms and doc values files are
	// memory mapped. All other files are opened using Lucene NIOFSDirectory.
	// Similarly to mmapfs be sure you have allowed plenty of virtual address space.
	Hybridfs = StorageType{"hybridfs"}
)

func (s StorageType) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *StorageType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "fs":
		*s = Fs
	case "niofs":
		*s = Niofs
	case "mmapfs":
		*s = Mmapfs
	case "hybridfs":
		*s = Hybridfs
	default:
		*s = StorageType{string(text)}
	}

	return nil
}

func (s StorageType) String() string {
	return s.Name
}

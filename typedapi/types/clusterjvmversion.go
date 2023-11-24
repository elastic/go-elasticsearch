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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ClusterJvmVersion type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L305-L335
type ClusterJvmVersion struct {
	// BundledJdk Always `true`. All distributions come with a bundled Java Development Kit
	// (JDK).
	BundledJdk bool `json:"bundled_jdk"`
	// Count Total number of selected nodes using JVM.
	Count int `json:"count"`
	// UsingBundledJdk If `true`, a bundled JDK is in use by JVM.
	UsingBundledJdk bool `json:"using_bundled_jdk"`
	// Version Version of JVM used by one or more selected nodes.
	Version string `json:"version"`
	// VmName Name of the JVM.
	VmName string `json:"vm_name"`
	// VmVendor Vendor of the JVM.
	VmVendor string `json:"vm_vendor"`
	// VmVersion Full version number of JVM.
	// The full version number includes a plus sign (+) followed by the build
	// number.
	VmVersion string `json:"vm_version"`
}

func (s *ClusterJvmVersion) UnmarshalJSON(data []byte) error {

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

		case "bundled_jdk":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.BundledJdk = value
			case bool:
				s.BundledJdk = v
			}

		case "count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "using_bundled_jdk":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.UsingBundledJdk = value
			case bool:
				s.UsingBundledJdk = v
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		case "vm_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VmName = o

		case "vm_vendor":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VmVendor = o

		case "vm_version":
			if err := dec.Decode(&s.VmVersion); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterJvmVersion returns a ClusterJvmVersion.
func NewClusterJvmVersion() *ClusterJvmVersion {
	r := &ClusterJvmVersion{}

	return r
}

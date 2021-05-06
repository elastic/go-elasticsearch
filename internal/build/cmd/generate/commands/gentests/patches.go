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

package gentests

import (
	"bytes"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	rePatchShardKeys       = regexp.MustCompile(`shards\.(\d+)\.(\d+)`)
	rePatchShardStoresKeys = regexp.MustCompile(`shards\.(\d+)\.stores\.(\d+)`)
	rePatchBucketKeys      = regexp.MustCompile(`aggregations.(\d+).(\d+).buckets`)
)

// PatchTestSource performs a regex based patching of the input.
//
func PatchTestSource(fpath string, r io.Reader) (io.Reader, error) {
	c, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// Patch the path to a shard by adding quotes around the shard number,
	// which is a map key, not an array index.
	for _, p := range []string{
		"indices.flush/10_basic.yml",
		"indices.put_template/10_basic.yml",
		"indices.recovery/10_basic.yml",
		"indices.segments/10_basic.yml",
		"indices.shard_stores/10_basic.yml",
		"indices.stats/12_level.yml",
	} {
		if strings.Contains(fpath, p) {
			c = rePatchShardKeys.ReplaceAll(c, []byte(`shards."$1".$2`))
			c = rePatchShardStoresKeys.ReplaceAll(c, []byte(`shards."$1".stores.$2`))
		}
	}

	for _, p := range []string{
		"search.aggregation/230_composite.yml",
	} {
		if strings.Contains(fpath, p) {
			c = rePatchBucketKeys.ReplaceAll(c, []byte(`aggregations."$1"."$2".buckets`))
		}
	}

	return bytes.NewReader(c), nil
}

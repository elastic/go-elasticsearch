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

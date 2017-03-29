package generator

import (
	"path/filepath"
	"testing"

	"github.com/golang/glog"
)

func TestUnmarshal(t *testing.T) {
	api := "indices.create"
	filename := api + ".json"
	spec, err := unmarshal(filepath.Join(apiDir, filename))
	if err != nil {
		t.Fatal(err)
	}
	expectedMethod := "PUT"
	methods := spec["indices.create"].(map[string]interface{})["methods"].([]interface{})
	if len(methods) != 1 || methods[0].(string) != expectedMethod {
		glog.Fatalf("Unexected methods: %s (expected %q)", methods, expectedMethod)
	}
}

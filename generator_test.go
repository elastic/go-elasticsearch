package main

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/golang/glog"
)

func TestUnmarshal(t *testing.T) {
	api := "indices.create"
	filename := api + ".json"
	spec, err := unmarshalSpec(filepath.Join(specDir, filename))
	if err != nil {
		t.Fatal(err)
	}
	expectedMethod := "PUT"
	methods := spec["indices.create"].(map[string]interface{})["methods"].([]interface{})
	if len(methods) != 1 || methods[0].(string) != expectedMethod {
		glog.Fatalf("Unexected methods: %s (expected %q)", methods, expectedMethod)
	}
}

func TestExecuteTemplate(t *testing.T) {
	api := "index"
	spec := map[string]interface{}{api: map[string]interface{}{}}
	apiDir, err := ioutil.TempDir("", "api")
	if err != nil {
		t.Fatal(err)
	}
	err = executeTemplate(spec, apiDir)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := "func (c *client) index() {}\n"
	code, err := ioutil.ReadFile(filepath.Join(apiDir, api) + ".go")
	if err != nil {
		t.Fatal(err)
	}
	if string(code) != expectedCode {
		t.Fatalf("Expected the generation of %q, got %q", expectedCode, code)
	}
}

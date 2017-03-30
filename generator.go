package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/golang/glog"
)

const (
	specDir         = "spec/elasticsearch/rest-api-spec/src/main/resources/rest-api-spec/api"
	specExt         = ".json"
	apiDir          = "api"
	templatesDir    = "templates"
	templateFile    = "method.tmpl"
	defaultTypeName = "client"
)

type method struct {
	TypeVar    string
	TypeName   string
	MethodName string
}

func executeTemplate(spec map[string]interface{}, outputDir string) error {
	if len(spec) > 1 {
		return fmt.Errorf("Each API spec is expected to contain a single API (found %d)", len(spec))
	}
	var api string
	for k := range spec {
		api = k
	}
	typeName := defaultTypeName
	apiParts := strings.Split(api, ".")
	switch len(apiParts) {
	case 1:
		api = apiParts[0]
	case 2:
		typeName = apiParts[0]
		api = apiParts[1]
	default:
		return fmt.Errorf("Unexpected API format: %s", api)
	}
	m := &method{
		TypeVar:    string(typeName[0]),
		TypeName:   typeName,
		MethodName: api,
	}
	t, err := template.ParseFiles(filepath.Join(templatesDir, templateFile))
	if err != nil {
		return fmt.Errorf("Failed to parse template in %q: %s", templateFile, err)
	}
	goFile, err := os.Create(filepath.Join(outputDir, api) + ".go")
	if err != nil {
		return err
	}
	defer goFile.Close()
	err = t.Execute(goFile, m)
	if err != nil {
		return fmt.Errorf("Failed to execute template in %q: %s", templateFile, err)
	}
	return err
}

func unmarshalSpec(filename string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var spec map[string]interface{}
	err = json.Unmarshal(bytes, &spec)
	return spec, err
}

func main() {
	flag.Parse()
	files, err := ioutil.ReadDir(specDir)
	if err != nil {
		glog.Fatal(err)
	}
	for _, file := range files {
		spec, err := unmarshalSpec(filepath.Join(specDir, file.Name()))
		if err != nil {
			glog.Error(err)
			continue
		}
		err = executeTemplate(spec, apiDir)
		if err != nil {
			glog.Error(err)
			continue
		}
	}
}

package generator

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/golang/glog"
)

const (
	apiDir = "spec/elasticsearch/rest-api-spec/src/main/resources/rest-api-spec/api"
)

func unmarshal(filename string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var spec map[string]interface{}
	err = json.Unmarshal(bytes, &spec)
	return spec, err
}

func main() {
	files, err := ioutil.ReadDir(apiDir)
	if err != nil {
		glog.Fatal(err)
	}
	for _, file := range files {
		if _, err = unmarshal(filepath.Join(apiDir, file.Name())); err != nil {
			glog.Error(err)
			continue
		}
	}
}

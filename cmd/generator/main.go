/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"

	"github.com/elastic/elasticsearch-go/generator"
	"github.com/golang/glog"
)

func main() {
	specDir := flag.String("specdir", "spec/elasticsearch/rest-api-spec/src/main/resources/rest-api-spec/api",
		"directory containing the JSON spec for the REST API")
	specFile := flag.String("specfile", "",
		"limit the generation to this JSON spec file")
	flag.Parse()
	files, err := ioutil.ReadDir(*specDir)
	if err != nil {
		glog.Fatal(err)
	}
	for _, file := range files {
		if *specFile != "" && file.Name() != *specFile {
			continue
		}
		g, err := generator.New(filepath.Join(*specDir, file.Name()))
		if err != nil {
			glog.Error(err)
			continue
		}
		err = g.Run()
		if err != nil {
			glog.Error(err)
			continue
		}
	}
}

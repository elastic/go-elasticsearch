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
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/elastic/elasticsearch-go/generator"
	"github.com/golang/glog"
)

type fileMap map[string]struct{}

func (f *fileMap) String() string {
	return fmt.Sprint(*f)
}

func (f *fileMap) Set(value string) error {
	for _, v := range strings.Split(value, ",") {
		(*f)[v] = struct{}{}
	}
	return nil
}

func main() {
	specDirFlag := flag.String("specdir", "spec/elasticsearch/rest-api-spec/src/main/resources/rest-api-spec/api",
		"directory containing the JSON spec for the REST API")
	skipFlag := fileMap{}
	flag.Var(&skipFlag, "skip",
		"comma-separated list of spec files to skip")
	cleanFlag := flag.Bool("clean", false,
		"delete the generated code")

	flag.Parse()
	files, err := ioutil.ReadDir(*specDirFlag)
	if err != nil {
		glog.Fatal(err)
	}
	for _, file := range files {
		if _, ok := skipFlag[file.Name()]; ok {
			continue
		}
		g, err := generator.New(filepath.Join(*specDirFlag, file.Name()))
		if err != nil {
			glog.Error(err)
			continue
		}
		if *cleanFlag {
			err = g.Clean()
		} else {
			err = g.Run()
		}
		if err != nil {
			glog.Error(err)
			continue
		}
	}
}

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
	"os"
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
	specDirFlag := flag.String("specdir", generator.DefaultSpecDir,
		"directory containing the JSON spec for the REST API")

	flag.Parse()
	g, err := generator.New(*specDirFlag)
	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}
	if err := g.Run(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}

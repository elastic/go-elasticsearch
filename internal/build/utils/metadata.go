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

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var resVersion = regexp.MustCompile(`(\d+\.(x|\d+.\d+)).*`)

// EsVersion returns the Elasticsearch from environment variable, Java property file, or an error.
//
func EsVersion(fpath string) (string, error) {
	if envEsVersion := os.Getenv("ELASTICSEARCH_BUILD_VERSION"); envEsVersion != "" {
		splitVersion := resVersion.FindStringSubmatch(envEsVersion)
		return splitVersion[1], nil
	} else {
		return "", fmt.Errorf("ELASTICSEARCH_BUILD_VERSION is empty")
	}
}

// GitCommit returns the Git commit from environment variable or parsing information from fpath, or an error.
//
func GitCommit(fpath string) (string, error) {
	if esBuildHash := os.Getenv("ELASTICSEARCH_BUILD_HASH"); esBuildHash != "" {
		return esBuildHash[:7], nil
	}
	return "", fmt.Errorf("ELASTICSEARCH_BUILD_HASH is empty")
}

// GitTag returns the Git tag for fpath if available, or an error.
//
func GitTag(fpath string) (string, error) {
	basePath, err := basePathFromFilepath(fpath)
	if err != nil {
		return "", fmt.Errorf("GitCommit: %s", err)
	}

	commit, err := GitCommit(fpath)
	if err != nil {
		return "", fmt.Errorf("GitTag: %s", err)
	}

	args := strings.Split("git --git-dir="+basePath+".git tag --points-at "+commit, " ")
	cmd := exec.Command(args[0:1][0], args[1:]...)
	// fmt.Printf("> %s\n", strings.Join(cmd.Args, " "))

	out, err := cmd.Output()
	if err != nil {
		return "", nil
	}

	return strings.TrimSpace(string(out)), nil
}

func basePathFromFilepath(fpath string) (string, error) {
	var bpath strings.Builder

	fpath, err := filepath.Abs(fpath)
	if err != nil {
		return "", err
	}

	for _, p := range strings.Split(fpath, string(filepath.Separator)) {
		if p == "rest-api-spec" {
			break
		}
		bpath.WriteString(p)
		bpath.WriteRune(filepath.Separator)
	}

	return bpath.String(), nil
}

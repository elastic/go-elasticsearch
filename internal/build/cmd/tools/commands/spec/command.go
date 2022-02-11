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

package cmd

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/internal/build/cmd"
	"github.com/elastic/go-elasticsearch/v8/internal/build/utils"
	"github.com/elastic/go-elasticsearch/v8/internal/version"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	output     *string
	commitHash *string
	debug      *bool
	info       *bool
)

func init() {
	output = toolsCmd.Flags().StringP("output", "o", "", "Path to a folder for generated output")
	commitHash = toolsCmd.Flags().StringP("commit_hash", "c", "", "Elasticsearch commit hash")
	debug = toolsCmd.Flags().BoolP("debug", "d", false, "Print the generated source to terminal")
	info = toolsCmd.Flags().Bool("info", false, "Print the API details to terminal")

	cmd.RegisterCmd(toolsCmd)
}

var toolsCmd = &cobra.Command{
	Use:   "download-spec",
	Short: "Download specification artifact for code & tests generation",
	Run: func(cmd *cobra.Command, args []string) {
		command := &Command{
			Output:     *output,
			CommitHash: *commitHash,
			Debug:      *debug,
			Info:       *info,
		}
		err := command.Execute()
		if err != nil {
			utils.PrintErr(err)
			os.Exit(1)
		}
	},
}

type Command struct {
	Output     string
	CommitHash string
	Debug      bool
	Info       bool
}

// download-spec runs a query to the Elastic artifact API, retrieve the list of active artifacts
// downloads, extract and write to disk the rest-resources spec alongside a json with relevant build information.
func (c Command) Execute() (err error) {
	const artifactsUrl = "https://artifacts-api.elastic.co/v1/versions"

	esBuildVersion := os.Getenv("ELASTICSEARCH_BUILD_VERSION")
	if esBuildVersion == "" {
		esBuildVersion = version.Client
	}

	versionUrl := strings.Join([]string{artifactsUrl, esBuildVersion}, "/")

	res, err := http.Get(versionUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()

	var v Versions
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&v)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if c.Debug {
		log.Printf("%d builds found", len(v.Version.Builds))
	}

	var build Build
	if c.CommitHash != "" {
		if build, err = findBuildByCommitHash(c.CommitHash, v.Version.Builds); err != nil {
			build = findMostRecentBuild(v.Version.Builds)
		}
	} else {
		build = findMostRecentBuild(v.Version.Builds)
	}
	if c.Debug {
		log.Printf("Build found : %s", build.Projects.Elasticsearch.CommitHash)
	}

	data, err := c.downloadZip(build)
	if err != nil {
		log.Fatalf("Cannot download zip from %s, reason : %s", build.zipfileUrl(), err)
	}

	if err := c.extractZipToDest(data); err != nil {
		log.Fatalf(err.Error())
	}

	d, _ := json.Marshal(build)

	err = c.writeFileToDest("elasticsearch.json", d)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return nil
}

func (c Command) writeFileToDest(filename string, data []byte) error {
	path := filepath.Join(c.Output, filename)
	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("cannot write file: %s", err)
	}
	if c.Debug {
		log.Printf("Successfuly written file to : %s", path)
	}
	return nil
}

type Versions struct {
	Version struct {
		Builds []Build `json:"builds"`
	} `json:"version"`
}

type PackageUrl struct {
	*url.URL
}

func (p *PackageUrl) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	url, err := url.Parse(string(data[1 : len(data)-1]))
	if err == nil {
		p.URL = url
	}
	return err
}

type BuildStartTime struct {
	*time.Time
}

func (t *BuildStartTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	parsedTime, err := time.Parse(`"`+"Mon, 2 Jan 2006 15:04:05 MST"+`"`, string(data))
	t.Time = &parsedTime
	return err
}

type Build struct {
	StartTime BuildStartTime `json:"start_time"`
	Version   string         `json:"version"`
	BuildId   string         `json:"build_id"`
	Projects  struct {
		Elasticsearch struct {
			Branch     string `json:"branch"`
			CommitHash string `json:"commit_hash"`
			Packages   map[string]struct {
				Url  PackageUrl `json:"url"`
				Type string     `json:"type"`
			}
		} `json:"elasticsearch"`
	} `json:"projects"`
}

func NewBuild() Build {
	t := time.Date(1970, 0,0,0,0,0,0, time.UTC)
	startTime := BuildStartTime{Time: &t}
	return Build{StartTime: startTime}
}

// zipfileUrl return the file URL for the rest-resources artifact from Build
// There should be only one artifact matching the requirements par Build.
func (b Build) zipfileUrl() string {
	for _, pack := range b.Projects.Elasticsearch.Packages {
		if pack.Type == "zip" && strings.Contains(pack.Url.String(), "rest-resources") {
			return pack.Url.String()
		}
	}
	return ""
}

// extractZipToDest extract the data from a previously downloaded file loaded in memory to Output target.
func (c Command) extractZipToDest(data []byte) error {
	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return err
	}

	if err = os.MkdirAll(c.Output, 0744); err != nil {
		return fmt.Errorf("cannot created destination directory: %s", err)
	}

	for _, file := range zipReader.File {
		f, err := file.Open()
		if err != nil {
			return fmt.Errorf("cannot read file in zipfile: %s", err)
		}
		defer f.Close()

		if file.FileInfo().IsDir() {
			path := filepath.Join(c.Output, file.Name)
			_ = os.MkdirAll(path, 0744)
		} else {
			data, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}

			if err := c.writeFileToDest(file.Name, data); err != nil {
				return err
			}
		}
	}

	if c.Debug {
		log.Printf("Zipfile successfully extracted to %s", c.Output)
	}

	return nil
}

// downloadZip fetches the rest-resources artifact from a Build and return its content as []byte.
func (c Command) downloadZip(b Build) ([]byte, error) {
	url := b.zipfileUrl()
	if c.Debug {
		log.Printf("Zipfile url : %s", b.zipfileUrl())
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept-Content", "gzip")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	return data, err
}

// findMostRecentBuild iterates through the builds retrieved from the api
// and return the latest one based on the StartTime of each Build.
func findMostRecentBuild(builds []Build) Build {
	var latestBuild Build
	latestBuild = NewBuild()
	for _, build := range builds {
		if build.StartTime.After(*latestBuild.StartTime.Time) {
			latestBuild = build
		}
	}
	return latestBuild
}

// findBuildByCommitHash iterates through the builds and returns the first occurrence of Build
// that matches the provided commitHash.
func findBuildByCommitHash(commitHash string, builds []Build) (Build, error) {
	for _, build := range builds {
		if build.Projects.Elasticsearch.CommitHash == commitHash {
			return build, nil
		}
	}
	return Build{}, fmt.Errorf("Build with commit hash %s not found", commitHash)
}

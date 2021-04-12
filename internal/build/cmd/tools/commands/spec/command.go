// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

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
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	output *string
	debug  *bool
	info   *bool
)

func init() {
	output = toolsCmd.Flags().StringP("output", "o", "", "Path to a folder for generated output")
	debug = toolsCmd.Flags().BoolP("debug", "d", false, "Print the generated source to terminal")
	info = toolsCmd.Flags().Bool("info", false, "Print the API details to terminal")

	cmd.RegisterCmd(toolsCmd)
}

var toolsCmd = &cobra.Command{
	Use:   "download-spec",
	Short: "Downdload specification artifact for code & tests generation",
	Run: func(cmd *cobra.Command, args []string) {
		command := &Command{
			Output: *output,
			Debug:  *debug,
			Info:   *info,
		}
		err := command.Execute()
		if err != nil {
			utils.PrintErr(err)
			os.Exit(1)
		}
	},
}

type Command struct {
	Output string
	Debug  bool
	Info   bool
}

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

	build := findMostRecentBuild(v.Version.Builds)
	if err := c.extractZipToDest(build); err != nil {
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
	return nil
}

type Versions struct {
	Version struct {
		Builds []Build `json:"builds"`
	} `json:"version"`
}

type Build struct {
	StartTime string `json:"start_time"`
	Version   string `json:"version"`
	BuildId   string `json:"build_id"`
	Projects  struct {
		Elasticsearch struct {
			Branch     string `json:"branch"`
			CommitHash string `json:"commit_hash"`
			Packages   map[string]struct {
				Url  string `json:"url"`
				Type string `json:"type"`
			}
		} `json:"elasticsearch"`
	} `json:"projects"`
}

func (b Build) zipfileUrl() string {
	for _, pack := range b.Projects.Elasticsearch.Packages {
		if pack.Type == "zip" && strings.Contains(pack.Url, "rest-resources") {
			return pack.Url
		}
	}
	return ""
}

func (c Command) extractZipToDest(b Build) error {
	url := b.zipfileUrl()
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept-Content", "gzip")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
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

	return nil
}

func findMostRecentBuild(builds []Build) Build {
	var latestBuild Build
	for _, build := range builds {
		t, _ := time.Parse(time.RFC1123, build.StartTime)
		lt, _ := time.Parse(time.RFC1123, latestBuild.StartTime)
		if t.After(lt) {
			latestBuild = build
		}
	}
	return latestBuild
}

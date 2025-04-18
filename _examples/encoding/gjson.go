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

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

var (
	faint = color.New(color.Faint)
	bold  = color.New(color.Bold)
)

func init() {
	log.SetFlags(0)
}

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	res, err := es.Cluster.Stats(es.Cluster.Stats.WithHuman())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	json := read(res.Body)

	fmt.Println(strings.Repeat("─", 50))
	faint.Print("cluster ")
	// Get cluster name
	bold.Print(gjson.Get(json, "cluster_name"))

	faint.Print(" status=")
	// Get cluster health status
	status := gjson.Get(json, "status")
	switch status.Str {
	case "green":
		bold.Add(color.FgHiGreen).Print(status)
	case "yellow":
		bold.Add(color.FgHiYellow).Print(status)
	case "red":
		bold.Add(color.FgHiRed).Print(status)
	default:
		bold.Add(color.FgHiRed, color.Underline).Print(status)
	}
	fmt.Println("\n" + strings.Repeat("─", 50))

	stats := []string{
		"indices.count",
		"indices.docs.count",
		"indices.store.size",
		"nodes.count.total",
		"nodes.os.mem.used_percent",
		"nodes.process.cpu.percent",
		"nodes.jvm.versions.#.version",
		"nodes.jvm.mem.heap_used",
		"nodes.jvm.mem.heap_max",
		"nodes.fs.free",
	}

	var maxwidth int
	for _, item := range stats {
		if len(item) > maxwidth {
			maxwidth = len(item)
		}
	}

	for _, item := range stats {
		pad := maxwidth - len(item)
		fmt.Print(strings.Repeat(" ", pad))
		faint.Printf("%s |", item)
		// Get stat dynamically from json
		fmt.Printf(" %s\n", gjson.Get(json, item))
	}
	fmt.Println()
}

func read(r io.Reader) string {
	var b bytes.Buffer
	b.ReadFrom(r)
	return b.String()
}

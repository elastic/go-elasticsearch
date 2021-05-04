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

package genexamples

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/internal/build/utils"
)

const testCheck = "\t" + `if err != nil {            // SKIP
		t.Fatalf("Error getting the response: %s", err)	 // SKIP
	}                                                  // SKIP
	defer res.Body.Close()                             // SKIP
`

var _ = log.Print

// ConsoleToGo contains translation rules.
//
var ConsoleToGo = []TranslateRule{

	{ // ----- Info() -----------------------------------------------------------
		Pattern: `^GET /\s*$`,
		Func: func(in string) (string, error) {
			return "res, err := es.Info()", nil
		}},

	{ // ----- Cat.Health() -----------------------------------------------------
		Pattern: `^GET /?_cat/health\?v`,
		Func: func(in string) (string, error) {
			return "\tres, err := es.Cat.Health(es.Cat.Health.WithV(true))", nil
		}},

	{ // ----- Cat.Indices() -----------------------------------------------------
		Pattern: `^GET /?_cat/indices`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?_cat/indices/?(?P<index>[^/\s\?]+)?(?P<params>\??[\S]+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			fmt.Fprint(&src, "\tres, err := es.Cat.Indices(\n")

			if matches[1] != "" {
				fmt.Fprintf(&src, "\tes.Cat.Indices.WithIndex([]string{%q}...),\n", matches[1])
			}

			if matches[2] != "" {
				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Cat.Indices", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Cluster.Health() -----------------------------------------------------
		Pattern: `^GET /?_cluster/health`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^(?P<method>GET) /?_cluster/health/?(?P<index>[^/\s\?]+)?(?P<params>\??[\S]+)?\s*$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			fmt.Fprint(&src, "\tres, err := es.Cluster.Health(\n")

			if matches[2] != "" {
				fmt.Fprintf(&src, "\tes.Cluster.Health.WithIndex([]string{%q}...),\n", matches[2])
			}

			if matches[3] != "" {
				params, err := queryToParams(matches[3])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Cluster.Health", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Cluster.State() -----------------------------------------------------
		Pattern: `^GET /?_cluster/state`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^(?P<method>GET) /?_cluster/state/?(?P<index>[^/\s\?]+)?(?P<params>\??[\S]+)?\s*$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			fmt.Fprint(&src, "\tres, err := es.Cluster.State(\n")

			if matches[2] != "" {
				fmt.Fprintf(&src, "\tes.Cluster.State.WithIndex([]string{%q}...),\n", matches[2])
			}

			if matches[3] != "" {
				params, err := queryToParams(matches[3])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Cluster.State", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Cluster.PutSettings() --------------------------------------------
		Pattern: `^PUT /?_cluster/settings`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^(?P<method>PUT) /?_cluster/settings/?(?P<params>\??[\S]+)?\s?(?P<body>.*)`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			fmt.Fprint(&src, "\tres, err := es.Cluster.PutSettings(\n")
			body, err := bodyStringToReader(matches[3])
			if err != nil {
				return "", fmt.Errorf("error converting body: %s", err)
			}
			fmt.Fprintf(&src, "\t%s", body)

			if matches[2] != "" {
				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Cluster.PutSettings", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil

		}},

	{ // ----- Nodes.Stats() -----------------------------------------------------
		Pattern: `^GET /?_nodes/.*/?stats`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^(?P<method>GET) /?_nodes/(?P<node_id>[[:alnum:]]+)?/?stats/?(?P<metric>[^/\s\?]+)?/?(?P<index_metric>[^/\s?]+)?(?P<params>\??[\S]+)?\s*$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}
			fmt.Println("matches")
			for i, m := range matches {
				fmt.Println(i, ":", m)
			}
			fmt.Fprint(&src, "\tres, err := es.Nodes.Stats(\n")

			if matches[2] != "" {
				fmt.Fprintf(&src, "\tes.Nodes.Stats.WithNodeID([]string{%q}...),\n", matches[2])
			}

			if matches[3] != "" {
				fmt.Fprintf(&src, "\tes.Nodes.Stats.WithMetric([]string{%q}...),\n", matches[3])
			}

			if matches[4] != "" {
				fmt.Fprintf(&src, "\tes.Nodes.Stats.WithIndexMetric([]string{%q}...),\n", matches[4])
			}

			if matches[5] != "" {
				params, err := queryToParams(matches[5])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Nodes.Stats", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- UpdateByQueryRethrottle() ---------------------------------------------------------
		Pattern: `^POST /?_update_by_query.+_rethrottle`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^POST /?_update_by_query/(?P<task_id>\S+)/_rethrottle(?P<params>\??[\S/]+)?`)

			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.UpdateByQueryRethrottle(\n")
			fmt.Fprintf(&src, "%q,\n", matches[1])
			if matches[2] != "" {
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}

				var requestsPerSecond []string
				for key, value := range params {
					if key == "requests_per_second" {
						requestsPerSecond = value
						delete(params, key)
					}
				}
				fmt.Fprintf(&src, "esapi.IntPtr(%v),\n", strings.Join(requestsPerSecond, ""))

				args, err := paramsToArguments("UpdateByQueryRethrottle", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- UpdateByQuery() ---------------------------------------------------------
		Pattern: `^(GET|POST) /?(\S+)?/?_update_by_query`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^(GET|POST) /?(?P<index>[^/]*)?/?_update_by_query(?P<params>\??[\S]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.UpdateByQuery(\n")
			if matches[2] != "" {
				fmt.Fprintf(&src, "\t[]string{%q},\n", matches[2])
			}

			if strings.TrimSpace(matches[4]) != "" {
				body, err := bodyStringToReader(matches[4])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				fmt.Fprintf(&src, "\tes.UpdateByQuery.WithBody(%s),\n", body)
			}

			var params url.Values
			if matches[3] != "" {
				var err error
				params, err = url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[3], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("UpdateByQuery", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprint(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- PutIngestPipeline() ------------------------------------------------------
		Pattern: `^PUT /?_ingest/pipeline.*`,
		Func: func(in string) (string, error) {
			var (
				src strings.Builder
			)

			re := regexp.MustCompile(`(?ms)^PUT /?_ingest/pipeline/(?P<id>[^/\?\s]+)/?(?P<params>\??[\S]+)?\s?(?P<body>.*)`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Ingest.PutPipeline(\n")
			fmt.Fprintf(&src, "\t%q,\n", matches[1])

			if matches[3] != "" {
				body, err := bodyStringToReader(matches[3])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				fmt.Fprintf(&src, "\t%s,\n", body)
			}

			if matches[2] != "" {
				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Ingest.PutPipeline", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, "\t%s", args)
			}

			src.WriteString(")")
			return src.String(), nil
		},
	},

	{ // ----- Index() or Create() ----------------------------------------------
		Pattern: `^(PUT|POST) /?.+/(_doc|_create|_update)/?.*`,
		Func: func(in string) (string, error) {
			var (
				src     strings.Builder
				apiName string
			)

			re := regexp.MustCompile(`(?ms)^(?P<method>PUT|POST) /?(?P<index>.+)/(?P<api>_doc|_create|_update)/?(?P<id>\w+)?(?P<params>\??[\S]+)?\s?(?P<body>.*)`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			switch matches[3] {
			case "_doc":
				apiName = "Index"
			case "_create":
				apiName = "Create"
			case "_update":
				apiName = "Update"
			default:
				return "", fmt.Errorf("Unknown API [%s]", matches[3])
			}

			src.WriteString("\tres, err := es." + apiName + "(\n")

			fmt.Fprintf(&src, "\t%q,\n", matches[2])

			if apiName == "Create" || apiName == "Update" {
				fmt.Fprintf(&src, "\t%q,\n", matches[4])
			}

			body, err := bodyStringToReader(matches[6])
			if err != nil {
				return "", fmt.Errorf("error converting body: %s", err)
			}
			fmt.Fprintf(&src, "\t%s,\n", body)

			if apiName == "Index" {
				if matches[4] != "" {
					fmt.Fprintf(&src, "\tes."+apiName+".WithDocumentID(%q),\n", matches[4])
				}
			}

			if matches[5] != "" {
				params, err := queryToParams(matches[5])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments(apiName, params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString("\tes." + apiName + ".WithPretty(),\n")
			src.WriteString("\t)")

			return src.String(), nil
		}},

	{ // ----- Indices.GetMapping() -------------------------------------------------
		Pattern: `^GET /?.*_mapping`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?(?P<index>[^/\s]+)?(?P<params>\??[\S/]+)?_mapping\s?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.GetMapping(")
			if matches[1] != "" {
				fmt.Fprintf(&src, "\tes.Indices.GetMapping.WithIndex(%q),\n", matches[1])
			}

			if matches[2] != "" {
				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Indices.PutTemplate", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.PutMapping() -------------------------------------------------
		Pattern: `^PUT /?[^/\s]+/_mapping`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^PUT /?(?P<index>[^/\s]+)(?P<params>\??[\S/]+)?_mapping\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.PutMapping(")

			if matches[2] != "" || matches[3] != "" {
				fmt.Fprintf(&src, "\n\t[]string{%q},\n", matches[1])

				if matches[3] != "" {
					body, err := bodyStringToReader(matches[3])
					if err != nil {
						return "", fmt.Errorf("error converting body: %s", err)
					}
					fmt.Fprintf(&src, "\t%s,\n", body)
				}

				if matches[2] != "" {
					params, err := queryToParams(matches[2])
					if err != nil {
						return "", fmt.Errorf("error parsing URL params: %s", err)
					}
					args, err := paramsToArguments("Indices.PutMapping", params)
					if err != nil {
						return "", fmt.Errorf("error converting params to arguments: %s", err)
					}
					fmt.Fprintf(&src, "\t%s", args)
				}
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.GetSettings() -------------------------------------------------
		Pattern: `^GET /?[^/\s]+/_settings`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?(?P<index>[^/\s]+)/_settings(?P<params>\??[\S/]+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.GetSettings(")

			if matches[1] != "" {
				fmt.Fprintf(&src, "\n\t\tes.Indices.GetSettings.WithIndex([]string{%q}...),\n", matches[1])
			}

			if matches[2] != "" {
				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Indices.GetSettings", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, "\t\t%s", args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.PutSettings() -------------------------------------------------
		Pattern: `^PUT /?[^/\s]+/_settings`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^PUT /?(?P<index>[^/\s]+)(?P<params>\??[\S/]+)?_settings\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.PutSettings(")

			if matches[2] != "" || matches[3] != "" {
				if matches[3] != "" {
					body, err := bodyStringToReader(matches[3])
					if err != nil {
						return "", fmt.Errorf("error converting body: %s", err)
					}
					fmt.Fprintf(&src, "\t%s,\n", body)
				}

				if matches[1] != "" {
					fmt.Fprintf(&src, "\n\tes.Indices.PutSettings.WithIndex([]string{%q}...),\n", matches[1])
				}

				if matches[2] != "" {
					params, err := queryToParams(matches[2])
					if err != nil {
						return "", fmt.Errorf("error parsing URL params: %s", err)
					}
					args, err := paramsToArguments("Indices.PutSettings", params)
					if err != nil {
						return "", fmt.Errorf("error converting params to arguments: %s", err)
					}
					fmt.Fprintf(&src, "\t%s", args)
				}
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.PutTemplate() -------------------------------------------------
		Pattern: `^PUT /?_template/[^/\s]+`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^PUT /?_template/(?P<name>[^/\s]+)/?(?P<params>\??[\S/]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.PutTemplate(")
			fmt.Fprintf(&src, "\n\t%q,\n", matches[1])

			if matches[2] != "" || matches[3] != "" {
				if strings.TrimSpace(matches[3]) != "" {
					body, err := bodyStringToReader(matches[3])
					if err != nil {
						return "", fmt.Errorf("error converting body: %s", err)
					}
					fmt.Fprintf(&src, "\t%s,\n", body)
				}

				if matches[2] != "" {
					params, err := queryToParams(matches[2])
					if err != nil {
						return "", fmt.Errorf("error parsing URL params: %s", err)
					}
					args, err := paramsToArguments("Indices.PutTemplate", params)
					if err != nil {
						return "", fmt.Errorf("error converting params to arguments: %s", err)
					}
					fmt.Fprintf(&src, args)
				}
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.GetTemplate() -------------------------------------------------
		Pattern: `^GET /?_template/[^/\s]+`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?_template/(?P<name>[^/\?\s]+)/?(?P<params>\??[\S/]+)?\s?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.GetTemplate(")
			if matches[1] != "" {
				fmt.Fprintf(&src, "\n\tes.Indices.GetTemplate.WithName(%q),\n", matches[1])
			}

			if matches[2] != "" {
				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Indices.GetTemplate", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.Create() -------------------------------------------------
		Pattern: `^PUT /?[^/\s]+\s?`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^PUT /?(?P<index>[^/\s]+)(?P<params>\??[\S/]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.Create(")
			if matches[2] != "" || matches[3] != "" {
				fmt.Fprintf(&src, "\n\t%q,\n", matches[1])

				if strings.TrimSpace(matches[3]) != "" {
					body, err := bodyStringToReader(matches[3])
					if err != nil {
						return "", fmt.Errorf("error converting body: %s", err)
					}
					fmt.Fprintf(&src, "\tes.Indices.Create.WithBody(%s),\n", body)
				}

				if matches[2] != "" {
					params, err := queryToParams(matches[2])
					if err != nil {
						return "", fmt.Errorf("error parsing URL params: %s", err)
					}
					args, err := paramsToArguments("Indices.Create", params)
					if err != nil {
						return "", fmt.Errorf("error converting params to arguments: %s", err)
					}
					fmt.Fprintf(&src, args)
				}
			} else {
				fmt.Fprintf(&src, "%q", matches[1])
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- ClearScroll() ---------------------------------------------------------
		Pattern: `^DELETE /?_search/scroll`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^DELETE /?_search/scroll/?(?P<ids>[[:alnum:]=_,]*)?(?P<params>\??[\S]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.ClearScroll(\n")

			if strings.TrimSpace(matches[1]) != "" {
				src.WriteString("\t\tes.ClearScroll.WithScrollID(")
				scrollIDs := strings.Split(matches[1], ",")
				for i, s := range scrollIDs {
					fmt.Fprintf(&src, "%q", s)
					if i < len(scrollIDs)-1 {
						src.WriteString(",")
					}
				}
				src.WriteString("),\n")
			}

			if strings.TrimSpace(matches[3]) != "" {
				body, err := bodyStringToReader(matches[3])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				fmt.Fprintf(&src, "\t\tes.ClearScroll.WithBody(%s),\n", body)
			}

			var params url.Values
			if matches[2] != "" {
				var err error
				params, err = url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("ClearScroll", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprint(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Delete() ---------------------------------------------------------
		Pattern: `^DELETE /?\w+/_doc/\w+`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^DELETE /?(?P<index>\w+)/_doc/(?P<id>\w+)(?P<params>\??\S+)?\s*$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			fmt.Fprint(&src, "\tres, err := es.Delete(")

			if matches[3] != "" {
				fmt.Fprintf(&src, "\t%q,\n\t%q,\n", matches[1], matches[2])
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[3], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Delete", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprint(&src, args)

				src.WriteString("\tes.Delete.WithPretty(),\n")
			} else {
				fmt.Fprintf(&src, "\t%q, %q, es.Delete.WithPretty()", matches[1], matches[2])
			}
			src.WriteString(")")

			return src.String(), nil
		}},

	{ // ----- Indices.Delete() -------------------------------------------------
		Pattern: `^DELETE /?[^/\s]+[^_/]*`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^DELETE /?(?P<index>[^/\s]+)(?P<params>\??[\S/]+)?$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.Delete(")

			if matches[2] != "" {
				fmt.Fprintf(&src, "\n\t[]string{%q},\n", matches[1])

				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Indices.Delete", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			} else {
				fmt.Fprintf(&src, "[]string{%q}", matches[1])
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.Refresh() -------------------------------------------------
		Pattern: `^GET /?[^/\s]*/?_refresh`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?(?P<index>[^/\s]*)/?_refresh$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.Refresh(")

			if matches[1] != "" {
				fmt.Fprintf(&src, "[]string{%q}", matches[1])
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.Open() + Indices.Close() ---------------------------------
		Pattern: `^POST /?[^/\s]*/?(?P<operation>_close|_open)\s*`,
		Func: func(in string) (string, error) {
			var src strings.Builder
			var apiName string

			re := regexp.MustCompile(`(?ms)^POST /?(?P<index>[^/\s]*)/?(?P<operation>_open|_close)\s*$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			switch matches[2] {
			case "_open":
				apiName = "Open"
			case "_close":
				apiName = "Close"
			default:
				return "", fmt.Errorf("unknown operation [%s]", matches[2])
			}

			fmt.Fprintf(&src, "\tres, err := es.Indices.%s(", apiName)

			if matches[1] != "" {
				fmt.Fprintf(&src, "[]string{%q}", matches[1])
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.Aliases() + Indices.PutAlias() ---------------------------
		Pattern: `^POST /?[^/\s]*/?_aliases`,
		Func: func(in string) (string, error) {
			var src strings.Builder
			var apiName string

			re := regexp.MustCompile(`(?ms)^POST /?(?P<index>[^/\s]*)/?_aliases/?(?P<params>\??[\S/]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			if matches[1] != "" {
				apiName = "PutAlias"
				src.WriteString("\tres, err := es.Indices.PutAlias(")
				fmt.Fprintf(&src, "[]string{%q}", matches[1])
			} else {
				apiName = "UpdateAliases"
				src.WriteString("\tres, err := es.Indices.UpdateAliases(")
			}

			if strings.TrimSpace(matches[3]) != "" {
				body, err := bodyStringToReader(matches[3])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				if apiName == "PutAlias" {
					fmt.Fprintf(&src, "\tes.Indices.%s.WithBody(%s),\n", apiName, body)
				} else {
					fmt.Fprintf(&src, "\t%s,\n", body)
				}
			}

			if matches[2] != "" {
				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments(apiName, params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Indices.Forcemerge() -------------------------------------------------
		Pattern: `^POST /?[^/\s]*/?_forcemerge`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^POST /?(?P<index>[^/\s]*)/?_forcemerge/?(?P<params>\??[\S/]+)?$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.Forcemerge(")

			if matches[1] != "" {
				fmt.Fprintf(&src, "\tes.Indices.Forcemerge.WithIndex([]string{%q}...),\n", matches[1])
			}

			if matches[2] != "" {
				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Indices.Forcemerge", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString(")")
			return src.String(), nil
		}},

	{ // ----- Get() or GetSource() ---------------------------------------------
		Pattern: `^GET /?.+/(_doc|_source)/\w+`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?(?P<index>.+)/(?P<api>_doc|_source)/(?P<id>\w+)(?P<params>\??\S+)?\s*$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			var apiName string
			switch matches[2] {
			case "_doc":
				apiName = "Get"
			case "_source":
				apiName = "GetSource"
			default:
				return "", fmt.Errorf("unknown API variant %q", matches[2])
			}

			if matches[4] == "" {
				fmt.Fprintf(&src, "\tres, err := es."+apiName+"(%q, %q, es."+apiName+".WithPretty()", matches[1], matches[3])
			} else {
				fmt.Fprintf(&src, "\tres, err := es."+apiName+"(\n\t%q,\n\t%q,\n\t", matches[1], matches[3])

				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[4], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments(apiName, params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)

				src.WriteString("\tes." + apiName + ".WithPretty(),\n")
			}

			src.WriteString(")")

			return src.String(), nil
		}},

	{ // ----- Exists() or ExistsSource() ---------------------------------------
		Pattern: `^HEAD /?\w+/(_doc|_source)/\w+`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^HEAD /?(?P<index>\w+)/(?P<api>_doc|_source)/(?P<id>\w+)(?P<params>\??[\S]+)?\s*$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			var apiName string
			switch matches[2] {
			case "_doc":
				apiName = "Exists"
			case "_source":
				apiName = "ExistsSource"
			default:
				return "", fmt.Errorf("unknown API variant %q", matches[2])
			}

			if matches[4] == "" {
				fmt.Fprintf(&src, "\tres, err := es."+apiName+"(%q, %q, es."+apiName+".WithPretty()", matches[1], matches[2])
			} else {
				fmt.Fprintf(&src, "\tres, err := es."+apiName+"(\n\t%q,\n\t%q,\n\t", matches[1], matches[2])
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[4], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments(apiName, params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)

				src.WriteString("\tes." + apiName + ".WithPretty(),\n")
			}

			src.WriteString(")")

			return src.String(), nil
		}},

	{ // ----- Scroll() ---------------------------------------------------------
		Pattern: `^(GET|POST) /?_search/scroll`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^(GET|POST) /?_search/scroll(?P<params>\??[\S]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Scroll(\n")

			if strings.TrimSpace(matches[3]) != "" {
				body, err := bodyStringToReader(matches[3])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				fmt.Fprintf(&src, "\tes.Scroll.WithBody(%s),\n", body)
			}

			var params url.Values
			if matches[2] != "" {
				var err error
				params, err = url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Search", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprint(&src, args)
			}

			var hasPretty bool
			for k, _ := range params {
				if k == "pretty" {
					hasPretty = true
				}
			}

			if !hasPretty {
				src.WriteString("\tes.Scroll.WithPretty(),\n")
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Search() ---------------------------------------------------------
		Pattern: `^(GET|POST) /?[^/]*/?_search`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^(GET|POST) /?(?P<index>[^/]*)?/?_search(?P<params>\??[\S]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Search(\n")
			if matches[2] != "" {
				fmt.Fprintf(&src, "\tes.Search.WithIndex(%q),\n", matches[2])
			}

			if strings.TrimSpace(matches[4]) != "" {
				body, err := bodyStringToReader(matches[4])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				fmt.Fprintf(&src, "\tes.Search.WithBody(%s),\n", body)
			}

			var params url.Values
			if matches[3] != "" {
				var err error
				params, err = url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[3], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Search", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprint(&src, args)
			}

			var hasPretty bool
			for k, _ := range params {
				if k == "pretty" {
					hasPretty = true
				}
			}

			if !hasPretty {
				src.WriteString("\tes.Search.WithPretty(),\n")
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Count() ---------------------------------------------------------
		Pattern: `^GET /?[^/]*/?_count`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?(?P<index>[^/]*)?/?_count(?P<params>\??[\S]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Count(\n")
			if matches[1] != "" {
				fmt.Fprintf(&src, "\tes.Count.WithIndex(%q),\n", matches[2])
			}

			if strings.TrimSpace(matches[3]) != "" {
				body, err := bodyStringToReader(matches[3])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				fmt.Fprintf(&src, "\tes.Count.WithBody(%s),\n", body)
			}

			var params url.Values
			if matches[2] != "" {
				var err error
				params, err = url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Count", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprint(&src, args)
			}

			var hasPretty bool
			for k, _ := range params {
				if k == "pretty" {
					hasPretty = true
				}
			}

			if !hasPretty {
				src.WriteString("\tes.Count.WithPretty(),\n")
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- DeleteByQueryRethrottle() ---------------------------------------------------------
		Pattern: `^POST /?_delete_by_query.+_rethrottle`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^POST /?_delete_by_query/(?P<task_id>\S+)/_rethrottle(?P<params>\??[\S/]+)?`)

			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.DeleteByQueryRethrottle(\n")
			fmt.Fprintf(&src, "%q,\n", matches[1])

			if matches[2] != "" {
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}

				var requestsPerSecond []string
				for key, value := range params {
					if key == "requests_per_second" {
						requestsPerSecond = value
						delete(params, key)
					}
				}
				fmt.Fprintf(&src, "esapi.IntPtr(%v),\n", strings.Join(requestsPerSecond, ""))

				args, err := paramsToArguments("DeleteByQueryRethrottle", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- DeleteByQuery() ---------------------------------------------------------
		Pattern: `^(GET|POST) /?(\S+)?/?_delete_by_query`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^(GET|POST) /?(?P<index>[^/]*)?/?_delete_by_query(?P<params>\??[\S]+)?\s?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.DeleteByQuery(\n")
			if matches[2] != "" {
				fmt.Fprintf(&src, "\t[]string{%q},\n", matches[2])
			}

			if strings.TrimSpace(matches[4]) != "" {
				body, err := bodyStringToReader(matches[4])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				fmt.Fprintf(&src, "\t%s,\n", body)
			}

			var params url.Values
			if matches[3] != "" {
				var err error
				params, err = url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[3], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("DeleteByQuery", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprint(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- ReindexRethrottle() ---------------------------------------------------------
		Pattern: `^POST /?_reindex.+_rethrottle`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^POST /?_reindex/(?P<task_id>\S+)/_rethrottle(?P<params>\??[\S/]+)?`)

			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.ReindexRethrottle(\n")
			fmt.Fprintf(&src, "%q,\n", matches[1])

			if matches[2] != "" {
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}

				var requestsPerSecond []string
				for key, value := range params {
					if key == "requests_per_second" {
						requestsPerSecond = value
						delete(params, key)
					}
				}
				fmt.Fprintf(&src, "esapi.IntPtr(%v),\n", strings.Join(requestsPerSecond, ""))

				args, err := paramsToArguments("ReindexRethrottle", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Reindex() ---------------------------------------------------------
		Pattern: `^POST /?_reindex`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^POST /?_reindex(?P<params>\??[\S/]+)?\s?(?P<body>.+)?`)

			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Reindex(\n")

			if strings.TrimSpace(matches[2]) != "" {
				body, err := bodyStringToReader(matches[2])
				if err != nil {
					return "", fmt.Errorf("error converting body: %s", err)
				}
				fmt.Fprint(&src, body)
			}

			if matches[1] != "" {
				src.WriteString(",\n")
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[1], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Reindex", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Bulk() ---------------------------------------------------------
		Pattern: `^POST /?(\S+)?/?_bulk`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^POST /?(?P<index>\S+)?/?_bulk(?P<params>\??[\S/]+)?\s?\n?(?P<body>.+)?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Bulk(\n")

			if matches[3] != "" {
				fmt.Fprintf(&src, "\tstrings.NewReader(`\n%s`),\n", matches[3])
			}

			if matches[1] != "" {
				fmt.Fprintf(&src, "\tes.Bulk.WithIndex(%q),\n", matches[1])
			}

			if matches[2] != "" {
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Bulk", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprint(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Tasks.Get() ------------------------------------------------------------
		Pattern: `^GET /?_tasks/.+`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?_tasks/?(?P<task_id>\S+)?/?(?P<params>\??[\S/]+)?`)

			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Tasks.Get(\n")

			if matches[1] != "" {
				fmt.Fprintf(&src, "\t%q,\n", matches[1])
			}

			if matches[2] != "" {
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}

				args, err := paramsToArguments("Tasks.Get", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Tasks.List() ------------------------------------------------------------
		Pattern: `^GET /?_tasks`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?_tasks/?(?P<params>\??[\S/]+)?`)

			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Tasks.List(\n")

			if matches[1] != "" {
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[1], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}

				args, err := paramsToArguments("Tasks.List", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Tasks.Cancel() ---------------------------------------------------------
		Pattern: `^POST /?_tasks.+_cancel`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^POST /?_tasks/(?P<task_id>\S+)/_cancel(?P<params>\??[\S/]+)?`)

			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Tasks.Cancel(\n")
			if matches[1] != "" {
				fmt.Fprintf(&src, "es.Tasks.Cancel.WithTaskID(%q),\n", matches[1])
			}

			if matches[2] != "" {
				params, err := url.ParseQuery(strings.TrimPrefix(strings.TrimPrefix(matches[2], "/"), "?"))
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}

				args, err := paramsToArguments("Tasks.Cancel", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			}

			src.WriteString("\t)")
			return src.String(), nil
		}},

	{ // ----- Indices.Get() -------------------------------------------------
		Pattern: `^GET /?[^/_\s]+`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?(?P<index>[^/\s]+)(?P<params>\??[\S/]+)?$`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.Get(")

			if matches[2] != "" {
				fmt.Fprintf(&src, "\n\t[]string{%q},\n", matches[1])

				params, err := queryToParams(matches[2])
				if err != nil {
					return "", fmt.Errorf("error parsing URL params: %s", err)
				}
				args, err := paramsToArguments("Indices.Get", params)
				if err != nil {
					return "", fmt.Errorf("error converting params to arguments: %s", err)
				}
				fmt.Fprintf(&src, args)
			} else {
				fmt.Fprintf(&src, "[]string{%q}", matches[1])
			}

			src.WriteString(")")
			return src.String(), nil
		}},
}

// Translator represents converter from Console source code to Go source code.
//
type Translator struct {
	Example Example
}

// TranslateRule represents a rule for translating code from Console to Go.
//
type TranslateRule struct {
	Pattern string
	Func    func(string) (string, error)
}

// IsTranslated returns true when a rule for translating the Console example to Go source code exists.
//
func (t Translator) IsTranslated() bool {
	var translated bool

	cmds, err := t.Example.Commands()
	if err != nil {
		panic(fmt.Sprintf("error getting example commands: %s", err))
	}

	for _, c := range cmds {
		for _, r := range ConsoleToGo {
			if r.Match(c) {
				translated = true
			}
		}
	}

	return translated
}

// Translate returns the Console code translated to Go.
//
func (t Translator) Translate() (string, error) {
	var out strings.Builder

	cmds, err := t.Example.Commands()
	if err != nil {
		return "", fmt.Errorf("error getting example commands: %s", err)
	}

	out.WriteRune('\n')
	fmt.Fprintf(&out, "\t// tag:%s[]\n", t.Example.Digest)
	for i, c := range cmds {
		var translated bool
		for _, r := range ConsoleToGo {
			if translated {
				break
			}

			if r.Match(c) {
				src, err := r.Func(c)
				if err != nil {
					return "", fmt.Errorf("error translating the example: %s", err)
				}
				translated = true

				if len(cmds) > 1 {
					out.WriteString("\t{\n")
				}
				out.WriteString(src)
				out.WriteRune('\n')
				out.WriteString("\tfmt.Println(res, err)\n")
				out.WriteString(testCheck)
				if len(cmds) > 1 {
					out.WriteString("\t}\n")
					if i != len(cmds)-1 {
						out.WriteString("\n")
					}
				}
			}
		}
		if !translated {
			return "", errors.New("no rule to translate the example")
		}
	}
	fmt.Fprintf(&out, "\t// end:%s[]\n", t.Example.Digest)
	return out.String(), nil
}

// Match returns true when the input matches the rule pattern.
//
func (r TranslateRule) Match(input string) bool {
	matched, _ := regexp.MatchString(r.Pattern, input)
	return matched
}

// queryToParams extracts the URL params.
//
func queryToParams(input string) (url.Values, error) {
	input = strings.TrimPrefix(input, "/")
	input = strings.TrimPrefix(input, "?")
	return url.ParseQuery(input)
}

// paramsToArguments converts params to API arguments.
//
func paramsToArguments(api string, params url.Values) (string, error) {
	var b strings.Builder

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		var (
			name  string
			value string
		)

		value = strings.Join(params[k], ",")

		switch k {
		case "q":
			name = "Query"
		default:
			name = utils.NameToGo(k)
		}

		switch k {
		case "timeout", "scroll": // duration
			if params[k][0] == "" {
				break
			}
			dur, err := time.ParseDuration(params[k][0])
			if err != nil {
				return "", fmt.Errorf("error parsing duration [%s]: %s", params[k][0], err)
			}
			value = fmt.Sprintf("time.Duration(%d)", time.Duration(dur))
		case "from", "size", "terminate_after", "version", "requests_per_second", "scroll_size", "max_num_segments": // numeric
			value = fmt.Sprintf("%s", value)
		case "pretty", "error_trace":
			value = "" // Helper methods don't take any value
		case "refresh":
			switch api {
			case "Index":
				if params[k][0] == "" {
					value = `"true"` // WithRefresh() needs a string
				} else {
					value = strconv.Quote(value)
				}
			case "Reindex", "DeleteByQuery", "UpdateByQuery":
				if params[k][0] == "" {
					value = "true" // WithRefresh() needs a boolean
				}
			default:
				value = strconv.Quote(value)
			}
		case "detailed", "flat_settings": // bool
			value = string(value)
		case "v": // Blank bool
			value = "true"
		default:
			value = strconv.Quote(value)
		}
		fmt.Fprintf(&b, "\tes.%s.With%s(%s),\n", api, name, value)
	}

	return b.String(), nil
}

// bodyStringToReader reformats input JSON string and returns it wrapped in strings.NewReader.
//
func bodyStringToReader(input string) (string, error) {
	var body bytes.Buffer
	err := json.Indent(&body, []byte(input), "\t\t", "  ")
	if err != nil {
		fmt.Printf("Input: %q", input)
		return "", err
	}
	return fmt.Sprintf("strings.NewReader(`%s`)", strings.TrimRight(body.String(), "\n")), nil
}

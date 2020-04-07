// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

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

	"github.com/elastic/go-elasticsearch/v8/internal/cmd/generate/utils"
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
		Pattern: "^GET /$",
		Func: func(in string) (string, error) {
			return "res, err := es.Info()", nil
		}},

	{ // ----- Cat.Health() -----------------------------------------------------
		Pattern: `^GET /_cat/health\?v`,
		Func: func(in string) (string, error) {
			return "\tres, err := es.Cat.Health(es.Cat.Health.WithV(true))", nil
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
			src.WriteString(")")

			return src.String(), nil

		}},

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
		Pattern: `^GET /?[^/\s]+/_mapping\s?(?P<body>.+)?`,
		Func: func(in string) (string, error) {
			var src strings.Builder

			re := regexp.MustCompile(`(?ms)^GET /?(?P<index>[^/\s]+)(?P<params>\??[\S/]+)?_mapping\s?`)
			matches := re.FindStringSubmatch(in)
			if matches == nil {
				return "", errors.New("cannot match example source to pattern")
			}

			src.WriteString("\tres, err := es.Indices.GetMapping(")
			if matches[2] != "" || matches[3] != "" {
				if matches[1] != "" {
					fmt.Fprintf(&src, "\tes.Indices.GetMapping.WithIndex(%q),\n", matches[1])
				}
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

	{ // ----- Indices.Delete() -------------------------------------------------
		Pattern: `^DELETE /?[^/\s]+`,
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

	{ // ----- Tasks.Get() ---------------------------------------------------------
		Pattern: `^GET /?_tasks`,
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
		case "timeout": // duration
			dur, err := time.ParseDuration(params[k][0])
			if err != nil {
				return "", fmt.Errorf("error parsing duration: %s", err)
			}
			value = fmt.Sprintf("time.Duration(%d)", time.Duration(dur))
		case "from", "size", "terminate_after", "version", "requests_per_second", "scroll_size": // numeric
			value = fmt.Sprintf("%s", value)
		case "pretty":
			value = "" // WithPretty() doesn't take any value
		case "refresh":
			switch api {
			case "Index":
				if params[k][0] == "" {
					value = `"true"` // WithRefresh() needs a string
				} else {
					value = strconv.Quote(value)
				}
			case "Reindex", "DeleteByQuery":
				if params[k][0] == "" {
					value = "true" // WithRefresh() needs a boolean
				}
			default:
				value = strconv.Quote(value)
			}
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

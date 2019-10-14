// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package gentests

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/tools/imports"
)

var (
	goruncache map[string]string
)

func init() {
	goruncache = make(map[string]string)
}

// Generator represents the "gentests" generator.
//
type Generator struct {
	b bytes.Buffer

	TestSuite TestSuite
}

// Output returns the generator output.
//
func (g *Generator) Output() (io.Reader, error) {
	name := g.TestSuite.Name()
	if g.TestSuite.Type == "xpack" {
		name = "XPack_" + name
	}

	g.genFileHeader()
	g.w("func Test" + name + "(t *testing.T) {\n")
	g.genInitializeClient()
	g.genHelpers()
	g.genCommonSetup()
	if g.TestSuite.Type == "xpack" {
		g.genXPackSetup()
	}
	if len(g.TestSuite.Setup) > 0 {
		g.w("// ----- Test Suite Setup --------------------------------------------------------\n")
		g.w("testSuiteSetup := func() {\n")
		g.genSetupTeardown(g.TestSuite.Setup)
		g.w("}\n")
		g.w("_ = testSuiteSetup\n")
		g.w("// --------------------------------------------------------------------------------\n")
		g.w("\n")
	}
	if len(g.TestSuite.Teardown) > 0 {
		g.w("\t// Teardown\n")
		g.w("\tdefer func(t *testing.T) {\n")
		g.genSetupTeardown(g.TestSuite.Teardown)
		g.w("\t}(t)\n")
	}
	for i, t := range g.TestSuite.Tests {
		g.w("\n")
		g.genLocationYAML(t)
		g.w("\t" + `t.Run("` + strings.Title(t.Name) + `", ` + "func(t *testing.T) {\n")
		if !g.genSkip(t) {
			g.w("\tdefer recoverPanic(t)\n")
			g.w("\tcommonSetup()\n")
			if g.TestSuite.Type == "xpack" {
				g.w("\txpackSetup()\n")
			}
			if len(g.TestSuite.Setup) > 0 {
				g.w("\ttestSuiteSetup()\n")
			}
			g.w("\n")
			if len(t.Setup) > 0 {
				g.w("\t// Test setup\n")
				g.genSetupTeardown(t.Setup)
			}
			if len(t.Teardown) > 0 {
				g.w("\t// Test teardown\n")
				g.w("\tdefer func(t) {\n")
				g.genSetupTeardown(t.Teardown)
				g.w("\t}(t *testing.T)\n")
			}
			if len(t.Setup) > 0 || len(t.Teardown) > 0 {
				g.w("\n")
			}
			g.genSteps(t)
		}
		g.w("\t})\n")
		if i < len(g.TestSuite.Tests)-1 {
			g.w("\n")
		}
	}
	g.w("}\n")
	return bytes.NewReader(g.b.Bytes()), nil
}

// OutputFormatted returns a formatted generator output.
//
func (g *Generator) OutputFormatted() (io.Reader, error) {
	out, err := g.Output()
	if err != nil {
		return bytes.NewReader(g.b.Bytes()), err
	}

	var b bytes.Buffer
	if _, err := io.Copy(&b, out); err != nil {
		return bytes.NewReader(g.b.Bytes()), err
	}

	fout, err := imports.Process(
		"",
		g.b.Bytes(),
		&imports.Options{
			AllErrors:  true,
			Comments:   true,
			FormatOnly: false,
			TabIndent:  true,
			TabWidth:   1,
		})
	if err != nil {
		return bytes.NewReader(b.Bytes()), err
	}

	g.b.Reset()
	g.b.Write(fout)

	return bytes.NewReader(fout), nil
}

func (g *Generator) w(s string) {
	g.b.WriteString(s)
}

func (g *Generator) gorun(code string) (string, error) {
	if goruncache[code] != "" {
		return goruncache[code], nil
	}

	dir, err := ioutil.TempDir("tmp", "gorun")
	if err != nil {
		return "", fmt.Errorf("gorun: %s", err)
	}
	f, err := os.Create(filepath.Join(dir, "type_for_struct_field.go"))
	if err != nil {
		return "", fmt.Errorf("gorun: %s", err)
	}
	defer func() {
		f.Close()
		os.RemoveAll(dir)
	}()

	// fmt.Println(code)
	if err := ioutil.WriteFile(f.Name(), []byte(code), 0644); err != nil {
		return "", fmt.Errorf("gorun: %s", err)
	}

	cmd := exec.Command("go", "run", f.Name())
	out, err := cmd.Output()
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("gorun: %s", e.Stderr)
		}
		return "", fmt.Errorf("gorun: %s", err)
	}

	goruncache[code] = string(out)

	return string(out), nil
}

func (g *Generator) genFileHeader() {
	g.w("// Code generated")
	if EsVersion != "" || GitCommit != "" || GitTag != "" {
		g.w(" from YAML test suite version")
		if GitCommit != "" {
			g.w(fmt.Sprintf(" %s", GitCommit))
			if GitTag != "" {
				g.w(fmt.Sprintf("|%s", GitTag))
			}
		}
	}
	g.w(" -- DO NOT EDIT\n")
	g.w("\n")
	g.w("package esapi_test\n")
	g.w(`
import (
	encjson "encoding/json"
	encyaml "gopkg.in/yaml.v2"
	"context"
	"crypto/tls"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var (
	// Prevent compilation errors for unused packages
	_ = encjson.NewDecoder
	_ = encyaml.NewDecoder
	_ = tls.Certificate{}
	_ = fmt.Printf
)` + "\n")
}

func (g *Generator) genInitializeClient() {
	if g.TestSuite.Type == "xpack" {
		g.w(`
	cfg := elasticsearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		// Logger: &estransport.TextLogger{
		// 	Output: os.Stdout,
		// 	// EnableRequestBody:  true,
		// 	// EnableResponseBody: true,
		// },
	}
	es, eserr := elasticsearch.NewClient(cfg)
	if eserr != nil {
		t.Fatalf("Error creating the client: %s\n", eserr)
	}

`)
	} else {
		g.w(`	es, eserr := elasticsearch.NewDefaultClient()
	if eserr != nil {
		t.Fatalf("Error creating the client: %s\n", eserr)
	}

`)
	}
}

func (g *Generator) genHelpers() {
	g.w(`recoverPanic := func(t *testing.T) {
	reLocation := regexp.MustCompile("(.*_test.go:\\d+).*")
	if rec := recover(); rec != nil {
		var loc string
		s := strings.Split(string(debug.Stack()), "\n")
		for i := len(s) - 1; i >= 0; i-- {
			if reLocation.MatchString(s[i]) {
				loc = strings.TrimSpace(s[i])
				break
			}
		}
		t.Fatalf("Panic: %s in %s", rec, reLocation.ReplaceAllString(loc, "$1"))
	}
}
_ = recoverPanic
` + "\n")

	g.w(`
	handleResponseError := func(t *testing.T, res *esapi.Response) {
		if res.IsError() {
			reLocation := regexp.MustCompile("(.*_test.go:\\d+).*")
			var loc string
			s := strings.Split(string(debug.Stack()), "\n")
			for i := len(s) - 1; i >= 0; i-- {
				if reLocation.MatchString(s[i]) {
					loc = strings.TrimSpace(s[i])
					break
				}
			}
			t.Logf("Response error: %s in %s", res, reLocation.ReplaceAllString(loc, "$1"))
		}
	}
	_ = handleResponseError
`)
	g.w("\n\n")
}

// Reference: https://github.com/elastic/elasticsearch/blob/master/test/framework/src/main/java/org/elasticsearch/test/rest/ESRestTestCase.java
//
func (g *Generator) genCommonSetup() {
	g.w(`
	// ----- Common Setup -------------------------------------------------------------
	commonSetup := func() {
		var res *esapi.Response

		{
			res, _ = es.Cluster.Health(es.Cluster.Health.WithWaitForNoInitializingShards(true))
			if res != nil && res.Body != nil {
				defer res.Body.Close()
			}
		}

		{
			res, _ = es.Indices.Delete([]string{"_all"})
			if res != nil && res.Body != nil { defer res.Body.Close() }
		}

		{
			var r map[string]interface{}
			res, _ = es.Indices.GetTemplate()
			if res != nil && res.Body != nil {
				defer res.Body.Close()
				json.NewDecoder(res.Body).Decode(&r)
				for templateName, _ := range r {
					if strings.HasPrefix(templateName, ".") {
						continue
					}
					if templateName == "security_audit_log" {
						continue
					}
					if templateName == "logstash-index-template" {
						continue
					}
					es.Indices.DeleteTemplate(templateName)
				}
			}
		}

		{
			res, _ = es.Indices.DeleteAlias([]string{"_all"}, []string{"_all"})
			if res != nil && res.Body != nil { defer res.Body.Close() }
		}

		{
			var r map[string]interface{}
			res, _ = es.Snapshot.GetRepository()
			if res != nil && res.Body != nil {
				defer res.Body.Close()
				json.NewDecoder(res.Body).Decode(&r)
				for repositoryID, _ := range r {
					var r map[string]interface{}
					res, _ = es.Snapshot.Get(repositoryID, []string{"_all"})
					json.NewDecoder(res.Body).Decode(&r)
					if r["responses"] != nil {
						for _, vv := range r["responses"].([]interface{}) {
							for _, v := range vv.(map[string]interface{})["snapshots"].([]interface{}) {
								snapshotID, ok := v.(map[string]interface{})["snapshot"]
								if !ok {
									continue
								}
								es.Snapshot.Delete(repositoryID, fmt.Sprintf("%s", snapshotID))
							}
						}
					}
					es.Snapshot.DeleteRepository([]string{fmt.Sprintf("%s", repositoryID)})
				}
			}
		}

		{
			res, _ = es.Cluster.Health(es.Cluster.Health.WithWaitForStatus("yellow"))
			if res != nil && res.Body != nil {
				defer res.Body.Close()
			}
		}
	}
	_ = commonSetup

	`)
}

// Reference: https://github.com/elastic/elasticsearch/blob/master/x-pack/plugin/src/test/java/org/elasticsearch/xpack/test/rest/XPackRestIT.java
// Reference: https://github.com/elastic/elasticsearch/blob/master/x-pack/plugin/core/src/test/java/org/elasticsearch/xpack/core/ml/integration/MlRestTestStateCleaner.java
//
func (g *Generator) genXPackSetup() {
	g.w(`
		// ----- XPack Setup -------------------------------------------------------------
		xpackSetup := func() {
			var res *esapi.Response

			{
				var r map[string]interface{}
				res, _ = es.Indices.GetTemplate()
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for templateName, _ := range r {
						if strings.HasPrefix(templateName, ".") {
							continue
						}
						es.Indices.DeleteTemplate(templateName)
					}
				}
			}

			{
				res, _ = es.Watcher.DeleteWatch("my_watch")
				if res != nil && res.Body != nil {
					defer res.Body.Close()
				}
			}

			{
				var r map[string]interface{}
				res, _ = es.Security.GetRole(es.Security.GetRole.WithPretty())
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for k, v := range r {
						reserved, ok := v.(map[string]interface{})["metadata"].(map[string]interface{})["_reserved"].(bool)
						if ok && reserved {
							continue
						}
						es.Security.DeleteRole(k)
					}
				}
			}

			{
				var r map[string]interface{}
				res, _ = es.Security.GetUser(es.Security.GetUser.WithPretty())
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for k, v := range r {
						reserved, ok := v.(map[string]interface{})["metadata"].(map[string]interface{})["_reserved"].(bool)
						if ok && reserved {
							continue
						}
						es.Security.DeleteUser(k)
					}
				}
			}

			{
				var r map[string]interface{}
				res, _ = es.Security.GetPrivileges(es.Security.GetPrivileges.WithPretty())
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for k, v := range r {
						reserved, ok := v.(map[string]interface{})["metadata"].(map[string]interface{})["_reserved"].(bool)
						if ok && reserved {
							continue
						}
						es.Security.DeletePrivileges(k, "_all")
					}
				}
			}

			{
				var r map[string]interface{}
				ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()
				es.ML.StopDatafeed("_all", es.ML.StopDatafeed.WithContext(ctx))
				res, _ = es.ML.GetDatafeeds()
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for _, v := range r["datafeeds"].([]interface{}) {
						datafeedID, ok := v.(map[string]interface{})["datafeed_id"]
						if !ok {
							continue
						}
						es.ML.DeleteDatafeed(datafeedID.(string))
					}
				}
			}

			{
				var r map[string]interface{}
				ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()
				es.ML.CloseJob("_all", es.ML.CloseJob.WithContext(ctx))
				res, _ = es.ML.GetJobs()
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for _, v := range r["jobs"].([]interface{}) {
						jobID, ok := v.(map[string]interface{})["job_id"]
						if !ok {
							continue
						}
						es.ML.DeleteJob(jobID.(string))
					}
				}
			}

			{
				var r map[string]interface{}
				res, _ = es.Rollup.GetJobs(es.Rollup.GetJobs.WithJobID("_all"))
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for _, v := range r["jobs"].([]interface{}) {
						jobID, ok := v.(map[string]interface{})["config"].(map[string]interface{})["id"]
						if !ok {
							continue
						}
						es.Rollup.StopJob(jobID.(string), es.Rollup.StopJob.WithWaitForCompletion(true))
						es.Rollup.DeleteJob(jobID.(string))
					}
				}
			}

			{
				var r map[string]interface{}
				res, _ = es.Tasks.List()
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for _, vv := range r["nodes"].(map[string]interface{}) {
						for _, v := range vv.(map[string]interface{})["tasks"].(map[string]interface{}) {
							cancellable, ok := v.(map[string]interface{})["cancellable"]
							if !ok || !cancellable.(bool) {
								continue
							}
							taskID := fmt.Sprintf("%v:%v", v.(map[string]interface{})["node"], v.(map[string]interface{})["id"])
							ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
							defer cancel()
							es.Tasks.Cancel(es.Tasks.Cancel.WithTaskID(taskID), es.Tasks.Cancel.WithContext(ctx))
						}
					}
				}
			}

			{
				var r map[string]interface{}
				res, _ = es.Snapshot.GetRepository()
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					json.NewDecoder(res.Body).Decode(&r)
					for repositoryID, _ := range r {
						var r map[string]interface{}
						res, _ = es.Snapshot.Get(repositoryID, []string{"_all"})
						json.NewDecoder(res.Body).Decode(&r)
						for _, vv := range r["responses"].([]interface{}) {
							for _, v := range vv.(map[string]interface{})["snapshots"].([]interface{}) {
								snapshotID, ok := v.(map[string]interface{})["snapshot"]
								if !ok {
									continue
								}
								es.Snapshot.Delete(repositoryID, fmt.Sprintf("%s", snapshotID))
							}
						}
						es.Snapshot.DeleteRepository([]string{fmt.Sprintf("%s", repositoryID)})
					}
				}
			}

			{
				res, _ = es.ILM.RemovePolicy("_all")
				if res != nil && res.Body != nil {
					defer res.Body.Close()
				}
			}

			{
				res, _ = es.Cluster.Health(es.Cluster.Health.WithWaitForStatus("yellow"))
				if res != nil && res.Body != nil {
					defer res.Body.Close()
				}
			}

			{
				res, _ = es.Security.PutUser("x_pack_rest_user", strings.NewReader(` + "`" + `{"password":"x-pack-test-password", "roles":["superuser"]}` + "`" + `), es.Security.PutUser.WithPretty())
				if res != nil && res.Body != nil {
					defer res.Body.Close()
				}
			}

			{
				res, _ = es.Indices.Refresh(es.Indices.Refresh.WithIndex("_all"))
				if res != nil && res.Body != nil {
					defer res.Body.Close()
				}
			}

			{
				res, _ = es.Cluster.Health(es.Cluster.Health.WithWaitForStatus("yellow"))
				if res != nil && res.Body != nil {
					defer res.Body.Close()
				}
			}

			{
				var i int
				for {
					i++
					var r map[string]interface{}
					res, _ = es.Cluster.PendingTasks()
					if res != nil && res.Body != nil {
						defer res.Body.Close()
						json.NewDecoder(res.Body).Decode(&r)
						if len(r["tasks"].([]interface{})) < 1 {
							break
						}
					}
					if i > 30 {
						break
					}
					time.Sleep(time.Second)
				}
			}
		}
		_ = xpackSetup

	`)
}

func (g *Generator) genLocationYAML(t Test) {
	f, err := os.Open(t.Filepath)
	if err != nil {
		g.w(fmt.Sprintf("// Error opening file: %s\n", err))
	}

	scanner := bufio.NewScanner(f)
	var i int
	for scanner.Scan() {
		i++
		tname := scanner.Text()
		tname = strings.TrimRight(tname, `:`)
		tname = strings.NewReplacer(`\"`, `"`).Replace(tname)
		tname = strings.TrimPrefix(tname, `"`)
		tname = strings.TrimSuffix(tname, `"`)

		if tname == t.OrigName {
			// TODO: Github URL (with proper branch/commit/etc)
			g.w("\t// Source: " + t.Filepath + fmt.Sprintf(":%d", i) + "\n\t//\n")
		}
	}
	if err := scanner.Err(); err != nil {
		g.w(fmt.Sprintf("// Error reading file: %s\n", err))
	}
}

func (g *Generator) genSkip(t Test) (skipped bool) {
	// Check the custom skip list
	if skips, ok := skipTests[t.BaseFilename()]; ok {
		if len(skips) < 1 {
			g.w("\t// Skipping all tests in '" + t.BaseFilename() + "'\n")
			g.w("\tt.SkipNow()\n\n")
			return true
		}

		for _, skip := range skips {
			if skip == t.OrigName {
				g.w("\tt.SkipNow()\n\n")
				return true
			}
		}
	}

	// Check the skip property coming from YAML
	if t.Skip {
		if t.SkipInfo != "" {
			g.w("\tt.Skip(" + strconv.Quote(t.SkipInfo) + ")\n\n")
			return true
		} else {
			g.w("\tt.SkipNow()\n\n")
			return true
		}
	}

	return false
}

func (g *Generator) genSetupTeardown(actions []Action) {
	g.genVarSection(Test{})

	for _, a := range actions {
		g.genAction(a, false)
		g.w("\n")
	}
}

func (g *Generator) genSteps(t Test) {
	var skipBody bool
	if !t.Steps.ContainsAssertion() && !t.Steps.ContainsCatch() && !t.Steps.ContainsStash() {
		skipBody = true
	}
	g.genVarSection(t, skipBody)

	for _, step := range t.Steps {
		switch step.(type) {
		case Action:
			// Generate debug info
			var dbg strings.Builder
			dbg.WriteString("\t\t// => " + step.(Action).Method() + "(")
			var j int
			for k, v := range step.(Action).Params() {
				j++
				dbg.WriteString(k + ": " + strings.Replace(fmt.Sprintf("%v", v), "\n", "|", -1))
				if j < len(step.(Action).Params()) {
					dbg.WriteString(", ")
				}
			}
			dbg.WriteString(") ")
			pad := 101 - dbg.Len()
			if pad < 0 {
				pad = 0
			}
			g.w(dbg.String() + strings.Repeat("-", pad) + "\n\t\t//\n")

			// Generate the action
			g.genAction(step.(Action), skipBody)
			g.w("\t\t// " + strings.Repeat("-", 96) + "\n\n")
		case Assertion:
			// Generate debug info
			g.w("\t\t// ~> ")
			g.w(fmt.Sprintf("%q: ", step.(Assertion).operation))
			g.w(strings.Replace(fmt.Sprintf("%s", step.(Assertion).payload), "\n", "|", -1))
			g.w("\n")
			// Generate the assertion
			g.genAssertion(step.(Assertion))
			g.w("\n")
		case Stash:
			// Generate setting the stash
			g.genStashSet(step.(Stash))
			g.w("\n")
		default:
			panic(fmt.Sprintf("Unknown step %T", step))
		}
	}
}

func (g *Generator) genVarSection(t Test, skipBody ...bool) {
	g.w("\t\tvar (\n")
	g.w("\t\t\treq esapi.Request\n")
	g.w("\t\t\tres *esapi.Response\n")
	g.w("\t\t\terr error\n\n")

	g.w("\t\t\tstash = make(map[string]interface{}, 0)\n\n")

	if (len(skipBody) < 1 || (len(skipBody) > 0 && skipBody[0] == false)) &&
		(t.Steps.ContainsAssertion() || t.Steps.ContainsCatch() || true) {
		g.w("\t\t\tbody []byte\n")
		g.w("\t\t\tmapi map[string]interface{}\n")
		g.w("\t\t\tslic []interface{}\n")
	}

	if t.Steps.ContainsAssertion("is_false", "is_true") {
		g.w("\n\t\t\tvalue reflect.Value\n")
	}

	g.w("\n")
	g.w("\t\t\tassertion bool\n")

	g.w("\t\t\tactual   interface{}\n")
	g.w("\t\t\texpected interface{}\n")
	g.w("\n")

	if t.Steps.ContainsAssertion("match", "match-regexp") {
		g.w("\n\t\t\tre *regexp.Regexp\n")
		g.w("\t\t\tmatch bool\n")
	}

	g.w("\t\t)\n\n")

	if (len(skipBody) < 1 || (len(skipBody) > 0 && skipBody[0] == false)) &&
		(t.Steps.ContainsAssertion() || t.Steps.ContainsCatch() || true) {
		g.w("\t\t_ = mapi\n")
		g.w("\t\t_ = slic\n")
		g.w("\n")
		g.w(`handleResponseBody := func(t *testing.T, res *esapi.Response) {
			// Reset deserialized structures
			mapi = make(map[string]interface{})
			slic = make([]interface{}, 0)

			var err error
			body, err = ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Error reading body: %s", err)
			}
			res.Body.Close()
			res.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			if len(body) < 1 {
				// FIXME: Hack to prevent EOF errors
				return
			}

			if len(res.Header) > 0 {
				if strings.Contains(res.Header["Content-Type"][0], "text/plain") {
					return
				}

				if strings.Contains(res.Header["Content-Type"][0], "yaml") {
					if strings.HasPrefix(string(body), "---\n-") {
						if err := encyaml.NewDecoder(res.Body).Decode(&slic); err != nil {
							t.Fatalf("Error parsing the response body: %s", err)
						}
					} else {
						if err := encyaml.NewDecoder(res.Body).Decode(&mapi); err != nil {
							t.Fatalf("Error parsing the response body: %s", err)
						}
					}
					return
				}
			}

			d := encjson.NewDecoder(res.Body)
			d.UseNumber()

			if strings.HasPrefix(string(body), "[") {
				if err := d.Decode(&slic); err != nil {
					t.Fatalf("Error parsing the response body: %s", err)
				}
			} else {
				if err := d.Decode(&mapi); err != nil {
					t.Fatalf("Error parsing the response body: %s", err)
				}
			}
		}` + "\n")
	}

	g.w("\n")

	g.w("\t\t_ = stash\n")

	if t.Steps.ContainsAssertion("is_false", "is_true") {
		g.w("\t\t_ = value\n")
	}

	g.w("\t\t_ = assertion\n")

	g.w("\t\t_ = actual\n")
	g.w("\t\t_ = expected\n")

	if t.Steps.ContainsAssertion("match", "match-regexp") {
		g.w("\n")
		g.w("\t\t_ = re\n")
		g.w("\t\t_ = match\n")
	}

	g.w("\n")
}

func (g *Generator) genAction(a Action, skipBody ...bool) {
	// Initialize the request
	g.w("\t\treq = esapi." + a.Request() + "{\n")

	// Pass the parameters
	for k, v := range a.Params() {
		// fmt.Printf("%s.%s: <%T> %v\n", a.Request(), k, v, v)

		if strings.HasPrefix(fmt.Sprintf("%s", v), "$") {
			v = `stash[` + strconv.Quote(fmt.Sprintf("%s", v)) + `]`
		}

		switch v.(type) {
		case bool:
			g.w("\t\t\t" + k + ": ")

			typ, ok := apiRegistry[a.Request()][k]
			if !ok {
				panic(fmt.Sprintf("%s.%s: field not found", a.Request(), k))
			}

			switch typ {
			case "bool":
				g.w(strconv.FormatBool(v.(bool)))
			case "*bool":
				g.w(`esapi.BoolPtr(` + strconv.FormatBool(v.(bool)) + `)`)
			case "string":
				g.w(`"` + strconv.FormatBool(v.(bool)) + `"`)
			case "[]string":
				// TODO: Listify
				g.w(`[]string{"` + strconv.FormatBool(v.(bool)) + `"}`)
			default:
				g.w(strconv.FormatBool(v.(bool)))
			}
			g.w(",\n")

		case string:
			if k == "Body" {
				g.w("\t\t\t" + k + ": ")
				body := v.(string)
				if !strings.HasSuffix(body, "\n") {
					body = body + "\n"
				}
				g.w("strings.NewReader(`" + body + "`)")
			} else {
				g.w("\t\t\t" + k + ": ")
				// TODO: Handle comma separated strings as lists

				// fmt.Printf("%s: %#v\n", a.Request(), apiRegistry[a.Request()])
				// fmt.Printf("%s: %#v\n", k, apiRegistry[a.Request()][k])
				typ, ok := apiRegistry[a.Request()][k]
				if !ok {
					panic(fmt.Sprintf("%s.%s: field not found", a.Request(), k))
				}

				var value string
				if strings.HasPrefix(v.(string), "stash[") {
					switch typ {
					case "bool":
						value = `fmt.Sprintf("%v", ` + v.(string) + `)`
					case "string":
						value = fmt.Sprintf("%s.(string)", v)
					case "[]string":
						// TODO: Comma-separated list => Quoted list
						value = fmt.Sprintf(`[]string{%s.(string)}`, v)
					case "int":
						value = `func() int {
				switch ` + v.(string) + `.(type) {
				case int:
					return ` + v.(string) + `.(int)
				case float64:
					return int(` + v.(string) + `.(float64))
				}
				case json.Number:
					v, _ := ` + v.(string) + `.(encjson.Number).Int64()
					vv := int(v)
					return vv
				panic(fmt.Sprintf(` + "`" + `Unexpected type %T for ` + v.(string) + "`" + `, ` + v.(string) + `))
			}()`
					case "*int":
						value = `func() *int {
				switch ` + v.(string) + `.(type) {
				case int:
					v := ` + v.(string) + `.(int)
					return &v
				case float64:
					v := int(` + v.(string) + `.(float64))
					return &v
				case json.Number:
					v, _ := ` + v.(string) + `.(encjson.Number).Int64()
					vv := int(v)
					return &vv
				}
				panic(fmt.Sprintf(` + "`" + `Unexpected type %T for ` + v.(string) + "`" + `, ` + v.(string) + `))
			}()`
					case "time.Duration":
						value = `fmt.Sprintf("%d", ` + v.(string) + `)`
					default:
						panic(fmt.Sprintf("Unexpected type %q for value %v", typ, v))
					}
				} else {
					switch typ {
					case "[]string":
						value = `[]string{` + fmt.Sprintf("%q", v) + `}`
					case "time.Duration":
						// re := regexp.MustCompile("^(\\d+).*")
						// value = re.ReplaceAllString(fmt.Sprintf("%s", v), "$1")
						inputValue := v.(string)
						if strings.HasSuffix(inputValue, "d") {
							inputValue = inputValue[:len(inputValue)-1]
							numericValue, err := strconv.Atoi(inputValue)
							if err != nil {
								panic(fmt.Sprintf("Cannot convert duration [%s]: %s", inputValue, err))
							}
							// Convert to hours
							inputValue = fmt.Sprintf("%dh", numericValue*24)
						}

						dur, err := time.ParseDuration(inputValue)
						if err != nil {
							panic(fmt.Sprintf("Cannot parse duration [%s]: %s", v, err))
						}
						value = fmt.Sprintf("%d", dur.Nanoseconds())
					default:
						value = fmt.Sprintf("%q", v)
					}
				}
				g.w(value)
			}
			g.w(",\n")

		case int, *int, float64:
			g.w("\t\t\t" + k + ": ")

			typ, ok := apiRegistry[a.Request()][k]
			if !ok {
				panic(fmt.Sprintf("%s.%s: field not found", a.Request(), k))
			}

			var value string
			switch typ {
			case "string":
				value = `"` + fmt.Sprintf("%d", v) + `"`
			case "[]string":
				value = `[]string{"` + fmt.Sprintf("%d", v) + `"}`
			case "time.Duration":
				re := regexp.MustCompile("^(\\d+).*")
				value = re.ReplaceAllString(fmt.Sprintf("%d", v), "$1")
			case "*int":
				g.w(`esapi.IntPtr(` + fmt.Sprintf("%d", v) + `)`)
			default:
				value = fmt.Sprintf("%v", v)
			}
			g.w(value)
			g.w(",\n")

		case []interface{}:
			g.w("\t\t\t" + k + ": ")

			typ, ok := apiRegistry[a.Request()][k]
			if !ok {
				panic(fmt.Sprintf("%s.%s: field not found", a.Request(), k))
			}

			switch typ {
			case "string":
				switch v.(type) {
				case string:
					g.w("`" + v.(string) + "`")
				case []interface{}:
					vvv := make([]string, 0)
					for _, vv := range v.([]interface{}) {
						vvv = append(vvv, fmt.Sprintf("%s", vv))
					}
					g.w("`" + strings.Join(vvv, ",") + "`")
				default:
					panic(fmt.Sprintf("<%s> %s{}.%s: unexpected value <%T> %#v", typ, a.Request(), k, v, v))
				}
			case "[]string":
				qv := make([]string, 0)
				for _, vv := range v.([]interface{}) {
					// TODO: Check type
					qv = append(qv, fmt.Sprintf("%q", vv.(string)))
				}
				g.w(`[]string{` + strings.Join(qv, ",") + `}`)
			case "io.Reader":
				// Serialize Bulk payloads ...
				if k == "Body" {
					var b strings.Builder
					for _, vv := range v.([]interface{}) {
						switch vv.(type) {
						case string:
							b.WriteString(vv.(string))
						default:
							j, err := json.Marshal(convert(vv))
							if err != nil {
								panic(fmt.Sprintf("%s{}.%s: %s (%s)", a.Request(), k, err, v))
							}
							b.WriteString(string(j))
						}
						b.WriteString("\n")
					}
					b.WriteString("\n")
					g.w("\t\tstrings.NewReader(`" + b.String() + "`)")
					// ... or just convert the value to JSON
				} else {
					j, err := json.Marshal(convert(v))
					if err != nil {
						panic(fmt.Sprintf("%s{}.%s: %s (%s)", a.Request(), k, err, v))
					}
					g.w("\t\tstrings.NewReader(`" + fmt.Sprintf("%s", j) + "`)")
				}
			}
			g.w(",\n")

		case map[interface{}]interface{}:
			g.w("\t\t\t" + k + ": ")
			// vv := unstash(convert(v).(map[string]interface{}))
			// fmt.Println(vv)
			j, err := json.Marshal(convert(v))
			if err != nil {
				panic(fmt.Sprintf("JSON parse error: %s; %s", err, v))
			} else {
				// Unstash values
				reStash := regexp.MustCompile(`("\$[^"]+")`)
				j = reStash.ReplaceAll(j, []byte("` + strconv.Quote(fmt.Sprintf(\"%v\", stash[$1])) + `"))

				g.w("\t\tstrings.NewReader(`" + fmt.Sprintf("%s", j) + "`)")
				g.w(",\n")
			}

		default:
			g.w(fmt.Sprintf("\t\t// TODO: %s (%v)\n", k, v))
		}
	}

	if len(a.headers) > 0 {
		if strings.Contains(a.headers["Accept"], "yaml") && strings.HasPrefix(a.Request(), "Cat") {
			g.w("\t\t" + `Format: "yaml",` + "\n")
		}
		if auth_header, ok := a.headers["Authorization"]; ok {
			auth_fields := strings.Split(auth_header, " ")
			auth_name := auth_fields[0]
			auth_value := auth_fields[1]
			if strings.HasPrefix(auth_value, "$") {
				auth_value = `fmt.Sprintf("%s", stash["` + strings.ReplaceAll(strings.ReplaceAll(auth_value, "{", ""), "}", "") + `"])`
			} else {
				auth_value = `"` + auth_value + `"`
			}
			g.w("\t\t" + `Header: http.Header{` + "\n")
			g.w("\t\t\t" + `"Authorization": []string{"` + auth_name + ` " + ` + auth_value + `},` + "\n")
			g.w("\t\t" + `},` + "\n")
		}
	}

	g.w("\t\t}\n\n")

	// Get response
	g.w("\t\tres, err = req.Do(context.Background(), es)\n")

	g.w(`		if err != nil {
			t.Fatalf("ERROR: %s", err)
		}
		defer res.Body.Close()
	`)

	g.w("\n\n")

	if len(a.catch) < 1 {
		// Handle error responses
		g.w(`		handleResponseError(t, res)` + "\n")
	} else {
		// TODO: Test catch
	}

	if len(skipBody) < 1 || (len(skipBody) > 0 && skipBody[0] == false) {
		// Read and parse the body
		g.w(`		handleResponseBody(t, res)` + "\n")
	}
}

func (g *Generator) genAssertion(a Assertion) {
	g.w(a.Condition())
	g.w(a.Error() + "\n")
	g.w("}\n") // Close the condition
}

func (g *Generator) genStashSet(s Stash) {
	g.w(fmt.Sprintf("// Set %q\n", s.Key()))

	value := s.Value()
	if strings.HasPrefix(value, `mapi["#`) {
		switch {
		case strings.HasPrefix(value, `mapi["#base64EncodeCredentials`):
			i, j := strings.Index(value, "("), strings.Index(value, ")")
			values := strings.Split(value[i+1:j], ",")
			value = `base64.StdEncoding.EncodeToString([]byte(`
			value += `strings.Join([]string{`
			for n, v := range values {
				value += `mapi["` + v + `"].(string)`
				if n < len(values)-1 {
					value += ","
				}
			}
			value += `}, ":")`
			value += `))`
		default:
			panic(fmt.Sprintf("Unknown transformation: %s", value))
		}
	}

	g.w(fmt.Sprintf("stash[%q] = %s\n", s.Key(), value))
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			var ks string
			switch k.(type) {
			case string:
				ks = k.(string)
			case int:
				ks = fmt.Sprintf("%d", k)
			default:
				ks = fmt.Sprintf("%v", k)
			}
			m2[ks] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

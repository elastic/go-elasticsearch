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
	g.genFileHeader()
	g.w("func Test" + g.TestSuite.Name() + "(t *testing.T) {\n")
	g.genInitializeClient()
	g.genHelpers()
	g.genCommonSetup()
	if len(g.TestSuite.Setup) > 0 {
		g.w("// ----- Test Suite Setup --------------------------------------------------------\n")
		g.w("testSuiteSetup := func() {\n")
		g.genSetupTeardown(g.TestSuite.Setup)
		g.w("}\n")
		g.w("// --------------------------------------------------------------------------------\n")
		g.w("\n")
	}
	if len(g.TestSuite.Teardown) > 0 {
		g.w("\t// Teardown\n")
		g.w("\tdefer func() {\n")
		g.genSetupTeardown(g.TestSuite.Teardown)
		g.w("\t}()\n")
	}
	for i, t := range g.TestSuite.Tests {
		g.w("\n")
		g.genLocationYAML(t)
		g.w("\t" + `t.Run("` + strings.Title(t.Name) + `", ` + "func(t *testing.T) {\n")
		g.genSkip(t)
		g.w("\tdefer recoverPanic(t)\n")
		g.w("\tcommonSetup()\n")
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
			g.w("\tdefer func() {\n")
			g.genSetupTeardown(t.Teardown)
			g.w("\t}()\n")
		}
		if len(t.Setup) > 0 || len(t.Teardown) > 0 {
			g.w("\n")
		}
		g.genSteps(t)
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
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	// Prevent compilation errors for unused packages
	_ = encjson.NewDecoder
	_ = encyaml.NewDecoder
	_ = fmt.Printf
)` + "\n")
}

func (g *Generator) genInitializeClient() {
	g.w(`	es, eserr := elasticsearch.NewDefaultClient()
	if eserr != nil {
		t.Fatalf("Error creating the client: %s\n", eserr)
	}

`)
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
` + "\n")

	g.w(`debug := func(v interface{}) {
		if os.Getenv("DEBUG") != "" {
			t.Logf("%s", v)
		}
	}
`)
	g.w(`
	handleResponseError := func(t *testing.T, res *esapi.Response) {
		if res.IsError() {
			t.Logf("Response error: %s", res)
			// t.Fatalf("Response error: %s", res.Status())
		}
	}
	_ = handleResponseError
`)
	g.w("\n\n")
}

func (g *Generator) genCommonSetup() {
	// TODO: Full setup
	// https://github.com/elastic/elasticsearch-ruby/blob/master/elasticsearch-api/test/integration/yaml_test_runner.rb
	g.w(`
	// ----- Common Setup -------------------------------------------------------------
	commonSetup := func() {
		var res *esapi.Response
		res, _ = es.Indices.Delete([]string{"_all"})
		res.Body.Close()

		res, _ = es.Indices.DeleteTemplate("*")
		res.Body.Close()

		res, _ = es.Indices.DeleteAlias([]string{"_all"}, []string{"_all"})
		res.Body.Close()

		res, _ = es.Snapshot.Delete("test_repo_create_1", "test_snapshot")
		res.Body.Close()
		res, _ = es.Snapshot.Delete("test_repo_restore_1", "test_snapshot")
		res.Body.Close()
		res, _ = es.Snapshot.Delete("test_cat_snapshots_1", "snap1")
		res.Body.Close()
		res, _ = es.Snapshot.Delete("test_cat_snapshots_1", "snap2")
		res.Body.Close()
		for _, n := range []string{"test_repo_create_1", "test_repo_restore_1", "test_repo_get_1", "test_repo_get_2", "test_repo_status_1", "test_cat_repo_1", "test_cat_repo_2", "test_cat_snapshots_1"} {
			res, _ = es.Snapshot.DeleteRepository([]string{n})
			res.Body.Close()
		}
	}
	commonSetup()
	// --------------------------------------------------------------------------------

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

func (g *Generator) genSkip(t Test) {
	// Check the custom skip list
	if skips, ok := skipTests[t.BaseFilename()]; ok {
		if len(skips) < 1 {
			g.w("\t// Skipping all tests in '" + t.BaseFilename() + "'\n")
			g.w("\tt.SkipNow()\n\n")
			return
		}

		for _, skip := range skips {
			if skip == t.OrigName {
				g.w("\tt.SkipNow()\n\n")
			}
		}
	}

	// Check the skip property coming from YAML
	if t.Skip {
		if t.SkipInfo != "" {
			g.w("\tt.Skip(" + strconv.Quote(t.SkipInfo) + ")\n\n")
		} else {
			g.w("\tt.SkipNow()\n\n")
		}
	}
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
		g.w(`handleResponseBody := func(res *esapi.Response) {
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
						dur, err := time.ParseDuration(v.(string))
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
				value = fmt.Sprintf("%d", v)
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

	if len(a.headers) > 0 && strings.Contains(a.headers["Accept"], "yaml") && strings.HasPrefix(a.Request(), "Cat") {
		g.w("\t\t" + `Format: "yaml",` + "\n")
	}

	g.w("\t\t}\n\n")

	// Get response
	g.w("\t\tres, err = req.Do(context.Background(), es)\n")

	g.w(`		if err != nil {
			t.Fatalf("ERROR: %s", err)
		}
		defer res.Body.Close()
		debug(res)
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
		g.w(`		handleResponseBody(res)` + "\n")
	}
}

func (g *Generator) genAssertion(a Assertion) {
	g.w(a.Condition())
	g.w(a.Error() + "\n")
	g.w("}\n") // Close the condition
}

func (g *Generator) genStashSet(s Stash) {
	g.w(fmt.Sprintf("// Set %q\n", s.Key()))
	g.w(fmt.Sprintf("stash[%q] = %s\n", s.Key(), s.Value()))
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

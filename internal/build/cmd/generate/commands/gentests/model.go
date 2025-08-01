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

package gentests

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v9/internal/build/utils"
)

var reFilename = regexp.MustCompile(`\d*_?(.+)\.ya?ml`)
var reNumber = regexp.MustCompile(`^\d+$`)
var reBaseFilename = regexp.MustCompile(`rest-api-spec/test/\w+/(.*$)`)

// TestPayload represents a single raw section (`---`) from the YAML file.
type TestPayload struct {
	Filepath string
	Payload  interface{}
}

// TestSuite represents a group of tests (a "section") in one file.
type TestSuite struct {
	Dir      string
	Filepath string
	Skip     bool
	SkipInfo string
	Type     string

	Setup    []Action
	Teardown []Action

	Tests []Test
}

// Test represents a single named test.
type Test struct {
	Name     string
	Filepath string
	OrigName string

	Skip     bool
	SkipInfo string

	Setup    []Action
	Teardown []Action

	Steps Steps
}

// Steps represents the test steps.
type Steps []Step

// Step represents an Action or an Assertion.
type Step interface{}

// Action represents an API action (`do`).
type Action struct {
	payload interface{}

	method  string
	params  map[interface{}]interface{}
	catch   string
	headers map[string]string
}

// Assertion represents a test assertion (`is_true`, `match`, etc).
type Assertion struct {
	payload interface{}

	operation string
}

// Stash represents the registry for `set` operations.
type Stash struct {
	payload interface{}
}

// NewTestSuite returns a test group from the payloads.
func NewTestSuite(fpath string, payloads []TestPayload) TestSuite {
	ts := TestSuite{
		Dir:      strings.Title(filepath.Base(filepath.Dir(fpath))),
		Filepath: fpath,
	}

	if strings.Contains(fpath, "platinum") {
		ts.Type = "xpack"
	}
	if ts.Type == "" {
		ts.Type = "free"
	}

	for _, payload := range payloads {
		secKeys := utils.MapKeys(payload.Payload)
		switch {
		case len(secKeys) > 0 && strings.Contains(strings.Join(secKeys, ","), "setup") || strings.Contains(strings.Join(secKeys, ","), "teardown"):
			for k, v := range payload.Payload.(map[interface{}]interface{}) {
				switch k {
				case "setup":
					for _, vv := range v.([]interface{}) {
						for k, vvv := range vv.(map[interface{}]interface{}) {
							if k == "do" {
								ts.Setup = append(ts.Setup, NewAction(vvv))
							}
						}
					}
				case "teardown":
					for _, vv := range v.([]interface{}) {
						for k, vvv := range vv.(map[interface{}]interface{}) {
							if k == "do" {
								ts.Teardown = append(ts.Teardown, NewAction(vvv))
							}
						}
					}
				}
			}
		default:
			for k, v := range payload.Payload.(map[interface{}]interface{}) {
				var t Test
				var steps []Step

				for _, vv := range v.([]interface{}) {
					switch utils.MapKeys(vv)[0] {
					case "skip":
						var (
							ok    bool
							skipV string
							skipR string
						)

						skip := vv.(map[interface{}]interface{})["skip"]

						if skipV, ok = skip.(map[interface{}]interface{})["version"].(string); ok {
							if skipRR, ok := skip.(map[interface{}]interface{})["reason"].(string); ok {
								skipR = skipRR
							}
							if skipV == "all" {
								t.Skip = true
								t.SkipInfo = "Skipping all versions"
								break
							}
							if ts.SkipEsVersion(skipV) {
								// fmt.Printf("Skip: %q, EsVersion: %s, Skip: %v, Reason: %s\n", skipV, EsVersion, ts.SkipEsVersion(skipV), skipR)
								t.Skip = true
								t.SkipInfo = skipR
							}
						}
					case "setup":
						for _, vvv := range vv.(map[interface{}]interface{}) {
							for _, vvvv := range vvv.([]interface{}) {
								t.Setup = append(t.Setup, NewAction(vvvv.(map[interface{}]interface{})["do"]))
							}
						}
					case "teardown":
						for _, vvv := range vv.(map[interface{}]interface{}) {
							for _, vvvv := range vvv.([]interface{}) {
								t.Teardown = append(t.Teardown, NewAction(vvvv.(map[interface{}]interface{})["do"]))
							}
						}
					case "set":
						for _, vvv := range vv.(map[interface{}]interface{}) {
							steps = append(steps, NewStash(vvv))
						}
					case "transform_and_set":
						for _, vvv := range vv.(map[interface{}]interface{}) {
							// NOTE: `set_and_transform` has flipped ordering of key and value, compared to `set`
							key := utils.MapValues(vvv)[0]
							val := utils.MapKeys(vvv)[0]
							payload := make(map[interface{}]interface{})
							payload[key] = val
							// fmt.Println(payload)
							steps = append(steps, NewStash(payload))
						}
					case "do":
						for _, vvv := range vv.(map[interface{}]interface{}) {
							key := utils.MapKeys(vvv)[0]
							if strings.HasPrefix(key, "_internal") {
								ts.Skip = true
							}
							steps = append(steps, NewAction(vvv))
						}
					case "is_true", "is_false", "match", "lt", "gt", "lte", "gte":
						for kkk, vvv := range vv.(map[interface{}]interface{}) {
							steps = append(steps, NewAssertion(kkk.(string), vvv))
						}
						// TODO: "length" operator
					}
				}

				if !ts.Skip {
					t.Name = strings.ReplaceAll(k.(string), `"`, `'`)
					t.Filepath = payload.Filepath
					t.OrigName = k.(string)
					t.Steps = steps

					ts.Tests = append(ts.Tests, t)
				}
			}
		}
	}

	return ts
}

// NewAction returns a new action from the payload.
func NewAction(payload interface{}) Action {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", r)
			fmt.Fprintf(os.Stderr, "Payload: %v\n", payload)
		}
	}()

	var a Action
	a.payload = payload
	a.headers = make(map[string]string)

	for k, v := range payload.(map[interface{}]interface{}) {
		switch k {
		case "catch":
			a.catch = v.(string)
		case "warnings", "allowed_warnings", "allowed_warnings_regex", "node_selector", "arbitrary_key":
			continue
		case "headers":
			for kk, vv := range v.(map[interface{}]interface{}) {
				a.headers[kk.(string)] = vv.(string)
			}
		default:
			a.method = k.(string)
			a.params = v.(map[interface{}]interface{})
		}
	}

	return a
}

// NewAssertion returns a new assertion from the payload.
func NewAssertion(operation string, payload interface{}) Assertion {
	return Assertion{operation: operation, payload: payload}
}

// NewStash returns a new stash from the payload.
func NewStash(payload interface{}) Stash {
	return Stash{payload: payload}
}

// Name returns a human name for the test suite.
func (ts TestSuite) Name() string {
	var b strings.Builder
	for _, v := range strings.Split(ts.Dir, ".") {
		b.WriteString(strings.Title(v))
	}

	bname := reFilename.ReplaceAllString(filepath.Base(ts.Filepath), "$1")
	for _, v := range strings.Split(bname, "_") {
		b.WriteString(strings.Title(v))
	}
	return strings.ReplaceAll(b.String(), "-", "")
}

// Filename returns a suitable filename for the test suite.
func (ts TestSuite) Filename() string {
	var b strings.Builder

	if ts.Type == "xpack" {
		b.WriteString("xpack_")
	}

	b.WriteString(strings.ToLower(strings.Replace(ts.Dir, ".", "_", -1)))
	b.WriteString("__")

	bname := reFilename.ReplaceAllString(filepath.Base(ts.Filepath), "$1")
	b.WriteString(strings.ToLower(bname))

	return b.String()
}

// SkipEsVersion returns true if the test suite should be skipped.
func (ts TestSuite) SkipEsVersion(minmax string) bool {
	return skipVersion(minmax)
}

// BaseFilename extracts and returns the test filename in form of `foo/bar/10_qux.yml`.
func (t Test) BaseFilename() string {
	parts := reBaseFilename.FindStringSubmatch(t.Filepath)
	if len(parts) < 1 {
		panic(fmt.Sprintf("Unexpected parts for path [%s]: %s", t.Filepath, parts))
	}
	return strings.TrimPrefix(parts[1], string(filepath.Separator))
}

// SkipEsVersion returns true if the test should be skipped.
func (t Test) SkipEsVersion(minmax string) bool {
	return skipVersion(minmax)
}

// ContainsAssertion returns true when the set of steps
// contain an assertion, or a specific assertion.
func (s Steps) ContainsAssertion(keys ...string) bool {
	for _, step := range s {
		if a, ok := step.(Assertion); ok {
			// No keys, any assertion matches
			if len(keys) < 1 {
				return true
			}
			for _, key := range keys {
				// Operation matches the key
				if a.operation == key {
					return true
				}
				if a.operation == "match" && key == "match-regexp" {
					// Does any value look like a regexp?
					val := utils.MapValues(a.payload)[0]
					if val, ok := val.(string); ok {
						if strings.HasPrefix(val, "/") {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

// ContainsCatch returns true when the set of steps contains the "catch" clause.
func (s Steps) ContainsCatch(keys ...string) bool {
	for _, step := range s {
		if a, ok := step.(Action); ok {
			if a.catch != "" {
				return true
			}
		}
	}
	return false
}

// ContainsStash returns true when the set of steps contains the "set" clause.
func (s Steps) ContainsStash(keys ...string) bool {
	for _, step := range s {
		if _, ok := step.(Stash); ok {
			return true
		}
	}
	return false
}

// ContainsCallToEndpoint returns true when the set of steps contains a call to an endpoint.
func (s Steps) ContainsCallToEndpoint(endpoint string) bool {
	for _, step := range s {
		if a, ok := step.(Action); ok {
			if a.Method() == utils.NameToGo(endpoint) {
				return true
			}
		}
	}
	return false
}

// Method returns the API method name for the action.
func (a Action) Method() string {
	return utils.NameToGo(a.method)
}

// Request returns the API request name for the action.
func (a Action) Request() string {
	var rParts []string
	parts := strings.Split(a.method, ".")
	for _, p := range parts {
		rParts = append(rParts, utils.NameToGo(p))
	}
	return strings.Join(rParts, "") + "Request"
}

// Params returns a map of parameters for the action.
func (a Action) Params() map[string]interface{} {
	var kk string
	out := make(map[string]interface{})
	for k, v := range a.params {
		switch k.(string) {
		case "context":
			kk = "ScriptContext"
		case "q":
			kk = "Query"
		case "ignore":
			// TODO: Properly handle ignoring status codes
			continue
		default:
			kk = utils.NameToGo(k.(string), utils.APIToGo(a.method))
		}
		switch v.(type) {
		case bool:
			out[kk] = v.(bool)
		case string:
			out[kk] = v.(string)
		case int:
			out[kk] = v.(int)
		case float64:
			out[kk] = v.(float64)
		default:
			out[kk] = v
		}
	}

	return out
}

// Condition returns the condition for the assertion.
//
// TODO: Evaulate https://pkg.go.dev/github.com/google/go-cmp/cmp
func (a Assertion) Condition() string {
	var (
		output string

		subject  string
		expected string
		operator string
	)

	switch a.operation {

	// ================================================================================
	case "is_true":
		subject = expand(a.payload.(string))
		nilGuard := catchnil(subject)

		if subject != "mapi" && subject != "slic" {
			output = `if !(` + nilGuard + ") {\n"
			output += `switch ` + escape(subject) + `.(type) {
case bool:
	assertion = (` + escape(subject) + ` == true)
case string:
	assertion = (` + escape(subject) + `.(string) != ` + `""` + `)
case int:
	assertion = (` + escape(subject) + `.(int) != 0)
case float64:
	assertion = (` + escape(subject) + `.(float64) != 0)
case encjson.Number:
	assertion = (` + escape(subject) + `.(encjson.Number).String() != "0")
default:
	assertion = (` + escape(subject) + ` != nil)
}` + "\n"
			output += "} else {\n"
			output += "assertion = true}\n"
			output += `if !assertion` + " {\n"
		} else {
			// if subject == "slic" {
			// 	output = `if fmt.Sprintf("%v", ` + subject + `) == fmt.Sprintf("%v", []interface{}{}) {` + "\n"
			// } else {
			// 	output = `if fmt.Sprintf("%v", ` + subject + `) == fmt.Sprintf("%v", map[string]interface{}{}) {` + "\n"
			// }
			output = `if ` + subject + ` == nil {` + "\n"
		}
		return output

	// ================================================================================
	case "is_false":
		subject = expand(a.payload.(string))
		nilGuard := catchnil(subject)

		if subject != "mapi" && subject != "slic" {
			output = `if !(` + nilGuard + ") {\n"
			output += `switch ` + escape(subject) + `.(type) {
case bool:
	assertion = (` + escape(subject) + ` == false)
case string:
	assertion = (` + escape(subject) + `.(string) == ` + `""` + `)
case int:
	assertion = (` + escape(subject) + `.(int) == 0)
case float64:
	assertion = (` + escape(subject) + `.(float64) == 0)
case encjson.Number:
	assertion = (` + escape(subject) + `.(encjson.Number).String() == "0")
default:
	assertion = (` + escape(subject) + ` == nil)
}` + "\n"
			output += "} else {\n"
			output += "assertion = true}\n"
			output += `if !assertion` + " {\n"
		} else {
			if subject == "slic" {
				output = `if fmt.Sprintf("%v", ` + subject + `) != fmt.Sprintf("%v", []interface{}{}) {` + "\n"
			} else {
				output = `if fmt.Sprintf("%v", ` + subject + `) != fmt.Sprintf("%v", map[string]interface{}{}) {` + "\n"
			}
		}
		return output

	// ================================================================================
	case "lt", "gt", "lte", "gte":
		key := utils.MapKeys(a.payload)[0]

		subject = expand(key)
		expected = fmt.Sprintf("%v", utils.MapValues(a.payload)[0])

		nilGuard := catchnil(subject)

		// Is the value a stash value?
		if strings.HasPrefix(fmt.Sprintf("%v", utils.MapValues(a.payload)[0]), "$") {
			expected = `stash[` + strconv.Quote(expected) + `].(encjson.Number).Float64()`
		} else {
			expected = `float64(` + expected + `), struct{}{}`
		}

		switch a.operation {
		case "lt":
			operator = "<"
		case "gt":
			operator = ">"
		case "lte":
			operator = "<="
		case "gte":
			operator = ">="
		}

		r := strings.NewReplacer(`"`, "", `\`, "")
		rkey := r.Replace(key)
		output += `		if ` + nilGuard + ` { t.Error("Expected [` + rkey + `] to not be nil") }
		actual, _ = ` + escape(subject) + `.(encjson.Number).Float64()
		expected, _ = ` + expected + `
		assertion = actual.(float64) ` + operator + ` expected.(float64)
		if !assertion {
			t.Logf("%v != %v", actual, expected)` + "\n"

		return output

	// ================================================================================
	case "match":
		var b strings.Builder

		key := utils.MapKeys(a.payload)[0]
		val := utils.MapValues(a.payload)[0]

		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// 1. Match on body
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		if key == "$body" {
			expected = strings.TrimSpace(fmt.Sprintf("%v", val))

			// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
			// 1a. Is the expected value a regex pattern?
			// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
			if strings.HasPrefix(expected, "/") {
				// TODO: Refactor
				var pattern string
				pattern = expected

				// Remove all comments
				reComment := regexp.MustCompile(`#\s*[^\n]+`)
				pattern = reComment.ReplaceAllString(pattern, "")

				// Remove the prefix and suffix slashes
				pattern = strings.TrimPrefix(pattern, "/")
				pattern = strings.TrimSuffix(pattern, "/")

				// Remove the new lines
				pattern = strings.Replace(pattern, "\n", "", -1)

				// Remove whitespace (can be preceded by slash, eg. cat.templates/10_basic.yml:255)
				reWhitespace := regexp.MustCompile(`\\*\s+`)
				pattern = reWhitespace.ReplaceAllString(pattern, `\s*`)

				b.WriteString(`re, err = regexp.Compile("(?m)" + ` + "`" + pattern + "`" + `)
		if err != nil {
			t.Fatalf("Regex compile error: %s", err)
		}

		match = re.Match(body)` + "\n")
				b.WriteString(`		if !match` + " {\n")
				return b.String()
			}

			// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
			// 1b. Is the expected value a literal value?
			// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
			// TODO: Match on literal value
			// output = fmt.Sprintf("if true { // TODO, MatchBodyLiteral: %s => %v\n", a.operation, a.payload)
			output = `if strings.HasPrefix(string(body), "[") {
					match = !reflect.DeepEqual(fmt.Sprintf("%#v", slic), ` + fmt.Sprintf("%#v", val) + `)
				} else {
					match = !reflect.DeepEqual(fmt.Sprintf("%#v", mapi), ` + fmt.Sprintf("%#v", val) + `)
				}
				if !match {` + "\n"
			return output

			// panic(fmt.Sprintf("MatchUnknown: %q => %v\n", a.operation, a.payload))
		}

		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// 2. Match on JSON/YAML field or numbered item
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

		subject = expand(key)
		expected = strings.TrimSpace(fmt.Sprintf("%v", val))

		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// 2a. Is the expected value a regex pattern?
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		if strings.HasPrefix(expected, "/") {
			var pattern string
			pattern = expected
			pattern = strings.Replace(pattern, "/", "", -1)
			pattern = strings.Replace(pattern, "\n", "", -1)
			pattern = strings.Replace(pattern, " ", `\s*`, -1)

			b.WriteString(`re, err = regexp.Compile("(?m)" + ` + "`" + pattern + "`" + `)
		if err != nil {
			t.Fatalf("Regex compile error: %s", err)
		}

		match = re.MatchString(fmt.Sprintf("%v", ` + escape(subject) + `))` + "\n")
			b.WriteString(`		if !match` + " {\n")
			return b.String()
		}

		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// 2b. Is the value a regular value?
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

		// fmt.Printf("subject: %s, val: %T, %v \n", subject, val, val)

		nilGuard := catchnil(subject)
		switch val.(type) {

		// --------------------------------------------------------------------------------
		case nil:
			output += `		if ` + escape(subject) + ` != nil` + " {\n"

		// --------------------------------------------------------------------------------
		case bool:
			output += `		if ` +
				nilGuard +
				" || \n" +
				`fmt.Sprintf("%v", ` + escape(subject) + `) != ` +
				`fmt.Sprintf("%v", ` + expected + `)` +
				" {\n"
		case int, float64:
			// Date range special case
			// From & to are set in yaml tests as integer
			// but ES returns a scientific notation float
			var specialCase string
			if strings.HasSuffix(key, "from") || strings.HasSuffix(key, "to") {
				specialCase = `else if floatValue, err := actual.(encjson.Number).Float64(); err == nil {
							actual = strconv.FormatFloat(floatValue, 'f', -1, 64)
						} `
			}

			r := strings.NewReplacer(`"`, "", `\`, "")
			rkey := r.Replace(key)
			output += `		if ` + nilGuard + ` { t.Error("Expected [` + rkey + `] to not be nil") }
						actual = ` + escape(subject) + `
						if intValue, ok := actual.(int); ok {
							actual = fmt.Sprint(intValue)
						}` + specialCase + `else {
							actual = actual.(encjson.Number).String()
						}
						expected = ` + expected + `
						// TEMP: Hack to prevent 1.0 != 1 errors
						if strings.HasSuffix(actual.(string), ".0") {
							actual = strings.TrimSuffix(actual.(string), ".0")
							expected = strings.TrimSuffix(` + strconv.Quote(expected) + `, ".0")
						}
						if ok, _ := regexp.Match(` + "`^\\d\\.\\d+E\\d+$`" + `, []byte(actual.(string))); ok {
							replacer := strings.NewReplacer("e+0", "E", "e+", "E")
							expected = replacer.Replace(` + strconv.Quote(expected) + `)
						}
						assertion = fmt.Sprintf("%v", actual) == fmt.Sprintf("%v", expected)
					if !assertion {
						t.Logf("%v != %v", actual, expected)` + "\n"

		// --------------------------------------------------------------------------------
		case string:
			output += `		if ` +
				nilGuard +
				" || \n" +
				`strings.TrimSpace(fmt.Sprintf("%s", ` + escape(subject) + `)) != `
			if strings.HasPrefix(expected, "$") {
				// Remove brackets if we compare to a stashed value replaced in the body.
				expected = strings.NewReplacer("{", "", "}", "").Replace(expected)
				output += `strings.TrimSpace(fmt.Sprintf("%s", ` + `stash["` + expected + `"]` + `))`
			} else {
				output += `strings.TrimSpace(fmt.Sprintf("%s", ` + strconv.Quote(expected) + `))`
			}
			output += " {\n"

		// --------------------------------------------------------------------------------
		case map[interface{}]interface{}, map[string]interface{}:
			// We cannot reliably serialize to json and compare the json outputs: YAML responses are parsed as
			// a map[interface{}]interface{} that encoding/json fails to marshall
			// See https://play.golang.org/p/jhcXwg5dIrn
			expectedOutput := flattenPayload(val)
			expectedPayload := fmt.Sprintf("%#v", expectedOutput)

			expectedPayload = strings.ReplaceAll(expectedPayload, "map[interface {}]interface {}", "map[string]interface {}")
			output = `		actual = fmt.Sprintf("%v",` + escape(subject) + `)
				expected = fmt.Sprintf("%v",` + expectedPayload + `)
				if actual != expected {` + "\n"

		// --------------------------------------------------------------------------------
		case []interface{}:
			expectedPayload := fmt.Sprintf("%#v", val)
			expectedPayload = strings.ReplaceAll(expectedPayload, "map[interface {}]interface {}", "map[string]interface {}")
			output = `		actual, _ = encjson.Marshal(` + escape(subject) + `)
				expected, _ = encjson.Marshal(` + expectedPayload + `)
				if fmt.Sprintf("%s", actual) != fmt.Sprintf("%s", expected) {` + "\n"

		// --------------------------------------------------------------------------------
		default:
			output += `		if true { // TODO, MatchMissingType: ` + fmt.Sprintf("<%[1]T>%[1]v", val) + "\n"
			output += `		// Subject: ` + subject + "\n"
			output += `		// Value:   ` + fmt.Sprintf("%#v", val) + "\n"
		}

		return output

	// ================================================================================
	default:
		return fmt.Sprintf("if true { // TODO, Unimplemented: %q => %v\n", a.operation, a.payload)
	}

	// panic(fmt.Sprintf("Unreachable: %q => %v", a.operation, a.payload))
}

// Error returns an error handling for the failed assertion.
func (a Assertion) Error() string {
	var (
		output string

		subject  string
		expected string
	)

	output = `Unexpected value; <` + a.operation + `> : ` + fmt.Sprintf("%s", escape(a.payload))

	switch a.operation {
	case "is_true":
		subject = expand(a.payload.(string))
		output = `Expected ` + escape(subject) + ` to be truthy`

	case "is_false":
		subject = expand(a.payload.(string))
		output = `Expected ` + escape(subject) + ` to be falsey`

	case "lt", "gt", "lte", "gte":
		subject = expand(utils.MapKeys(a.payload)[0])
		expected = fmt.Sprintf("%v", utils.MapValues(a.payload)[0])
		output = `Expected ` + escape(utils.MapKeys(a.payload)[0]) + ` ` + a.operation + ` ` + escape(expected)

	case "length":
		panic(fmt.Sprintf("TODO: length: %s", a.payload))

	case "match":
		key := utils.MapKeys(a.payload)[0]
		val := utils.MapValues(a.payload)[0]

		expected := strings.TrimSpace(fmt.Sprintf("%v", val))

		if key == "$body" { // 1. Match on body

			if strings.HasPrefix(expected, "/") {
				return `t.Errorf("Expected [$body] to match pattern: %s", re)`
			}
			output = `Expected [$body] to match value: ` + escape(expected)

		} else { // 2. Match on JSON/YAML field or numbered item
			if key == "" {
				key = "response"
			}

			if strings.HasPrefix(expected, "/") { // 2a. Is the value to match a regex?
				expected = strings.Replace(expected, "\n", "<br>", -1)
				expected = strings.Replace(expected, `/`, `|`, -1)

				return `t.Errorf("Expected [` + strings.Trim(strconv.Quote(escape(key)), `"`) + `] to match pattern: %s", re)`
			}

			// 2b. Is the value to match a regular value?
			output = `Expected [` + strings.Trim(strconv.Quote(escape(key)), `"`) + `] to match '` + escape(expected) + `'`
		}
	}

	return `t.Error(` + strconv.Quote(output) + `)`
}

// Key returns the stash key as a string.
func (s Stash) Key() string {
	vals := utils.MapValues(s.payload)
	if len(vals) < 1 {
		return ""
	}

	return fmt.Sprintf("$%s", vals[0])
}

// Value return the stash value as a string.
func (s Stash) FirstValue() string {
	vals := utils.MapKeys(s.payload)
	if len(vals) < 1 {
		return ""
	}

	return vals[0]
}

// ExpandedValue returns the stash value in a container as a string.
func (s Stash) ExpandedValue() string {
	vals := utils.MapKeys(s.payload)
	if len(vals) < 1 {
		return ""
	}

	return expand(fmt.Sprintf("%v", vals[0]))
}

// expand "unwraps" a field name like `foo.bar.0.baz` into a proper Go structure
//
// See https://github.com/elastic/elasticsearch/tree/master/rest-api-spec/src/main/resources/rest-api-spec/test#dot-notation
func expand(s string, format ...string) string {
	var (
		b bytes.Buffer

		container string
	)

	// Handle the "literal dot" in a fieldname
	s = strings.Replace(s, `\.`, `_~_|_~_`, -1)

	parts := strings.Split(s, ".")

	if len(parts) > 0 {
		if reNumber.MatchString(parts[0]) {
			container = "slic"
		} else {
			container = "mapi"
		}
	} else {
		container = "mapi"
	}

	b.WriteString(container)

	if len(parts) > 0 && parts[0] == "" {
		return b.String()
	}

	for i, v := range parts {
		if reNumber.MatchString(v) {
			if i > 0 {
				b.WriteString(`.([]interface{})`)
			}
			b.WriteString(`[`)
			b.WriteString(v)
			b.WriteString(`]`)
		} else {
			if i > 0 {
				if len(format) > 0 && format[0] == "yaml" {
					b.WriteString(`.(map[interface{}]interface{})`)
				} else {
					b.WriteString(`.(map[string]interface{})`)
				}
			}
			b.WriteString("[")
			if strings.HasPrefix(v, "$") {
				b.WriteString(fmt.Sprintf(`stash["%s"].(string)`, strings.Trim(v, `"`))) // Remove the quotes from keys
			} else {
				b.WriteString(`"`)
				b.WriteString(strings.Trim(v, `"`)) // Remove the quotes from keys
				b.WriteString(`"`)
			}
			b.WriteString("]")
		}

	}
	return strings.Replace(b.String(), `_~_|_~_`, `\.`, -1)
}

// catchnil returns a condition which expands the input and checks if any part is not nil.
func catchnil(input string) string {
	var output string

	parts := strings.Split(strings.Replace(input, `\.`, `_~_|_~_`, -1), ".")
	for i := range parts {
		// Skip casting for runtime stash variable replacement nil escape
		if parts[i] == "(string)]" {
			continue
		}

		part := parts[:i+1]
		if strings.Contains(parts[i], "$") {
			part = append(part, parts[i+1])
		}
		output += strings.Join(part, ".") + " == nil"
		if i+1 < len(parts) {
			output += " ||\n\t\t"
		}
	}
	output = strings.Replace(output, `_~_|_~_`, `.`, -1)

	return output
}

// escape replaces unsafe characters in strings
func escape(s interface{}) string {
	if s, ok := s.(string); ok {
		// s = strings.Replace(s, `"`, `'`, -1)
		// s = strings.Replace(s, `\`, `\\`, -1)
		s = strings.Replace(s, `\.`, `.`, -1)
		return s
	}
	return fmt.Sprintf("%s", s)
}

// skipVersion parses minmax string and returns
// true when EsVersion is in the range.
func skipVersion(minmax string) bool {
	versions := strings.Split(minmax, "-")
	if len(versions) < 2 {
		panic(fmt.Sprintf("skipVersion: Unexpected input: %q", minmax))
	}

	min, max := strings.TrimSpace(versions[0]), strings.TrimSpace(versions[1])

	if max == "" { // Skip when max version is not set
		return true
	}

	if EsVersion >= min && EsVersion <= max {
		return true
	}

	return false
}

// flattenPayload serializes the expected payload as a map tree to compare within tests.
func flattenPayload(val interface{}) map[string]interface{} {
	expectedOutput := make(map[string]interface{})
	if cast, ok := val.(map[interface{}]interface{}); ok {
		for k, v := range cast {
			if _, ok := v.(map[interface{}]interface{}); ok {
				v = flattenPayload(v)
			}
			expectedOutput[fmt.Sprintf("%v", k)] = v
		}
	} else {
		expectedOutput = val.(map[string]interface{})
	}
	return expectedOutput
}

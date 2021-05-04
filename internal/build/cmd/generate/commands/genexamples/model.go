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
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func init() {
	sort.Strings(EnabledFiles)
}

// EnabledFiles contains a list of files where documentation should be generated.
//
var EnabledFiles = []string{
	"aggregations/bucket/datehistogram-aggregation.asciidoc",
	"aggregations/bucket/filter-aggregation.asciidoc",
	"aggregations/bucket/terms-aggregation.asciidoc",
	"aggregations/metrics/valuecount-aggregation.asciidoc",
	"api-conventions.asciidoc",
	"cat/indices.asciidoc",
	"cluster/health.asciidoc",
	"docs/bulk.asciidoc",
	"docs/delete-by-query.asciidoc",
	"docs/delete.asciidoc",
	"docs/get.asciidoc",
	"docs/index_.asciidoc",
	"docs/reindex.asciidoc",
	"docs/update-by-query.asciidoc",
	"docs/update.asciidoc",
	"getting-started.asciidoc",
	"indices/aliases.asciidoc",
	"indices/create-index.asciidoc",
	"indices/delete-index.asciidoc",
	"indices/get-index.asciidoc",
	"indices/get-mapping.asciidoc",
	"indices/put-mapping.asciidoc",
	"indices/templates.asciidoc",
	"indices/update-settings.asciidoc",
	"mapping.asciidoc",
	"mapping/fields/id-field.asciidoc",
	"mapping/params/fielddata.asciidoc",
	"mapping/params/format.asciidoc",
	"mapping/params/multi-fields.asciidoc",
	"mapping/types/array.asciidoc",
	"mapping/types/date.asciidoc",
	"mapping/types/keyword.asciidoc",
	"mapping/types/nested.asciidoc",
	"mapping/types/numeric.asciidoc",
	"query-dsl.asciidoc",
	"query-dsl/bool-query.asciidoc",
	"query-dsl/exists-query.asciidoc",
	"query-dsl/function-score-query.asciidoc",
	"query-dsl/match-all-query.asciidoc",
	"query-dsl/match-phrase-query.asciidoc",
	"query-dsl/match-query.asciidoc",
	"query-dsl/multi-match-query.asciidoc",
	"query-dsl/nested-query.asciidoc",
	"query-dsl/query-string-query.asciidoc",
	"query-dsl/query_filter_context.asciidoc",
	"query-dsl/range-query.asciidoc",
	"query-dsl/regexp-query.asciidoc",
	"query-dsl/term-query.asciidoc",
	"query-dsl/terms-query.asciidoc",
	"query-dsl/wildcard-query.asciidoc",
	"search.asciidoc",
	"search/count.asciidoc",
	"search/request-body.asciidoc",
	"search/request/from-size.asciidoc",
	"search/request/scroll.asciidoc",
	"search/request/sort.asciidoc",
	"search/search.asciidoc",
	"search/suggesters.asciidoc",
	"setup/install/check-running.asciidoc",
	"setup/logging-config.asciidoc",
}

var (
	reHTTPMethod = regexp.MustCompile(`^HEAD|GET|PUT|DELETE|POST`)
	reAnnotation = regexp.MustCompile(`^(?P<line>.*)\s*\<\d\>\s*$`)
	reComment    = regexp.MustCompile(`^#\s?`)
)

// Example represents the code example in documentation.
//
// See: https://github.com/elastic/built-docs/blob/master/raw/en/elasticsearch/reference/master/alternatives_report.json
//
type Example struct {
	SourceLocation struct {
		File string
		Line int
	} `json:"source_location"`

	Digest string
	Source string
}

// IsEnabled returns true when the example should be processed.
//
func (e Example) IsEnabled() bool {
	// TODO(karmi): Use "filepatch.Match()" to support glob patterns

	index := sort.SearchStrings(EnabledFiles, e.SourceLocation.File)

	if index > len(EnabledFiles)-1 || EnabledFiles[index] != e.SourceLocation.File {
		return false
	}

	return true
}

// IsExecutable returns true when the example contains a request.
//
func (e Example) IsExecutable() bool {
	return reHTTPMethod.MatchString(e.Source)
}

// IsTranslated returns true when the example can be converted to Go source code.
//
func (e Example) IsTranslated() bool {
	return Translator{Example: e}.IsTranslated()
}

// ID returns example identifier.
//
func (e Example) ID() string {
	return fmt.Sprintf("%s:%d", e.SourceLocation.File, e.SourceLocation.Line)
}

// Chapter returns the example chapter.
//
func (e Example) Chapter() string {
	r := strings.NewReplacer("/", "_", "-", "_", ".asciidoc", "")
	return r.Replace(e.SourceLocation.File)
}

// GithubURL returns a link for the example source.
//
func (e Example) GithubURL() string {
	return fmt.Sprintf("https://github.com/elastic/elasticsearch/blob/master/docs/reference/%s#L%d", e.SourceLocation.File, e.SourceLocation.Line)
}

// Commands returns the list of commands from source.
//
func (e Example) Commands() ([]string, error) {
	var (
		buf  strings.Builder
		list []string
		scan = bufio.NewScanner(strings.NewReader(e.Source))
	)

	for scan.Scan() {
		line := scan.Text()

		if reComment.MatchString(line) {
			continue
		}

		line = reAnnotation.ReplaceAllString(line, "$1")

		if reHTTPMethod.MatchString(line) {
			if buf.Len() > 0 {
				list = append(list, buf.String())
			}
			buf.Reset()
		}

		buf.WriteString(line)
		buf.WriteString("\n")
	}

	if err := scan.Err(); err != nil {
		return list, err
	}

	if buf.Len() > 0 {
		list = append(list, buf.String())
	}

	return list, nil
}

// Translated returns the code translated from Console to Go.
//
func (e Example) Translated() (string, error) {
	return Translator{Example: e}.Translate()
}

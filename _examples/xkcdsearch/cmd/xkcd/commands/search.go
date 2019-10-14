// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package commands

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/elastic/go-elasticsearch/v8/_examples/xkcdsearch"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search xkcd.com",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "\x1b[2m%s\x1b[0m\n", strings.Repeat("━", tWidth))
		fmt.Fprintf(os.Stdout, "\x1b[2m?q=\x1b[0m\x1b[1m%s\x1b[0m\n", strings.Join(args, " "))
		fmt.Fprintf(os.Stdout, "\x1b[2m%s\x1b[0m\n", strings.Repeat("━", tWidth))

		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;107;41mERROR: %s\x1b[0m\n", err)
		}

		config := xkcdsearch.StoreConfig{Client: es, IndexName: IndexName}
		store, err := xkcdsearch.NewStore(config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;107;41mERROR: %s\x1b[0m\n", err)
			os.Exit(1)
		}
		search := Search{store: store, reHighlight: regexp.MustCompile("<em>(.+?)</em>")}

		results, err := search.getResults(strings.Join(args, " "))
		if err != nil {
			fmt.Fprintf(os.Stderr, "\x1b[1;107;41mERROR: %s\x1b[0m\n", err)
			os.Exit(1)
		}

		if results.Total < 1 {
			fmt.Fprintln(os.Stdout, "⨯ No results")
			fmt.Fprintf(os.Stdout, "\x1b[2m%s\x1b[0m\n", strings.Repeat("─", tWidth))
			os.Exit(0)
		}

		for _, result := range results.Hits {
			search.displayResult(os.Stdout, result)
		}
	},
}

// Search allows to get and display results matching query.
//
type Search struct {
	store       *xkcdsearch.Store
	reHighlight *regexp.Regexp
}

func (s *Search) getResults(query string) (*xkcdsearch.SearchResults, error) {
	return s.store.Search(query)
}

func (s *Search) displayResult(w io.Writer, hit *xkcdsearch.Hit) {
	var title string
	if len(hit.Highlights.Title) > 0 {
		title = hit.Highlights.Title[0]
	} else {
		title = hit.Title
	}
	fmt.Fprintf(w, "\x1b[2m•\x1b[0m \x1b[1m%s\x1b[0m", s.highlightString(title))
	fmt.Fprintf(w, " [%s] ", hit.Published)
	fmt.Fprintf(w, "<%s>\n", hit.URL)

	if len(hit.Highlights.Alt) > 0 {
		fmt.Fprintf(w, "  \x1b[2m%s\x1b[0m", "Alt: ")
		fmt.Fprintf(w, "%s\n", s.highlightString(hit.Highlights.Alt[0]))
	}

	if len(hit.Highlights.Transcript) > 0 {
		t := strings.Join(hit.Highlights.Transcript, " … ")
		t = strings.Replace(t, "\n", " ", -1)
		t = s.highlightString(t)
		fmt.Fprintf(w, "  \x1b[2m%s\x1b[0m", "Transcript: ")
		fmt.Fprintf(w, "%s\n", t)
	}
	fmt.Fprintf(os.Stdout, "\x1b[2m%s\x1b[0m\n", strings.Repeat("─", tWidth))
}

func (s *Search) highlightString(input string) string {
	return s.reHighlight.ReplaceAllString(input, "\x1b[30;47m$1\x1b[0m")
}

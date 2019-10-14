// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package gensource

import (
	"fmt"
	"math"
	"strings"
	"text/tabwriter"

	"github.com/elastic/go-elasticsearch/v7/internal/cmd/generate/utils"
)

// DebugInfo returns information about the endpoint as a string.
//
func (e *Endpoint) DebugInfo() string {
	var out strings.Builder
	w := tabwriter.NewWriter(&out, 0, 0, 1, ' ', 0)

	fmt.Fprintln(&out, strings.Repeat("─", utils.TerminalWidth()))
	fmt.Fprintf(&out, "API: %s (%s:%s)\n", e.MethodWithNamespace(), e.Type, e.Name)
	fmt.Fprintf(&out, "%s\n", e.Documentation.Description[0:int(math.Min(float64(80), float64(len(e.Documentation.Description))))])
	fmt.Fprintln(&out, strings.Repeat("─", utils.TerminalWidth()))

	fmt.Fprintln(&out, "Paths:")
	for _, path := range e.URL.Paths {
		fmt.Fprintf(w, "%6s\t%s", path.Methods[0], path.Path)
		if path.Deprecated.Version != "" {
			fmt.Fprintf(w, "\t*deprecated*")
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()

	longestPath := e.URL.Paths[0]
	for _, v := range e.URL.Paths {
		if len(v.Path) > len(longestPath.Path) {
			longestPath = v
		}
	}

	if len(longestPath.Parts) > 0 {
		fmt.Fprintln(&out, "Parts:")
		for _, part := range longestPath.Parts {
			fmt.Fprintf(w, "  • %s\t", part.Name)
			fmt.Fprintf(w, "  %s", part.Type)
			if part.Required {
				fmt.Fprint(w, ", required")
			}
			if part.Default != nil {
				fmt.Fprintf(w, ", default: %s", part.Default)
			}
			fmt.Fprint(w, "\n")
		}
		w.Flush()
	}

	if len(e.URL.Params) > 0 {
		fmt.Fprintln(&out, "Params:")
		for _, param := range e.URL.Params {
			fmt.Fprintf(w, "  • %s\t %s", param.Name, param.Type)
			if param.Required {
				fmt.Fprint(w, ", required")
			}
			if param.Type == "enum" {
				fmt.Fprintf(w, ": %s", strings.Join(param.Options, ", "))
			}
			if param.Default != nil {
				fmt.Fprintf(w, ", default: %s", param.Default)
			}
			fmt.Fprint(w, "\n")
		}
		w.Flush()
	}

	if e.Body != nil {
		fmt.Fprintln(&out, "Body:")
		if e.Body.Description != "" {
			fmt.Fprintf(&out, "  %s.", e.Body.Description)
		}
		if e.Body.Required {
			fmt.Fprintf(&out, " *Required*")
		} else {
			fmt.Fprintf(&out, " Optional")
		}
		if e.Body.ContentType != "" {
			fmt.Fprintf(&out, ", format: %s", e.Body.ContentType)
		}
		fmt.Fprintf(&out, "\n")
	}

	return out.String()
}

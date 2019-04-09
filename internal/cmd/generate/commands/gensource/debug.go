package gensource

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/elastic/go-elasticsearch/v8/internal/cmd/generate/utils"
)

// DebugInfo returns information about the endpoint as a string.
//
func (e *Endpoint) DebugInfo() string {
	var out strings.Builder
	w := tabwriter.NewWriter(&out, 0, 0, 1, ' ', 0)

	fmt.Fprintln(&out, strings.Repeat("─", utils.TerminalWidth()))
	fmt.Fprintf(&out, "API: %s (%s)\n", e.MethodWithNamespace(), e.Name)
	// fmt.Fprintf(&out, "<%s>\n", e.Documentation)
	fmt.Fprintln(&out, strings.Repeat("─", utils.TerminalWidth()))

	fmt.Fprintf(&out, "Methods:\n  %s\n", e.Methods)

	fmt.Fprintln(&out, "Paths:")
	for _, path := range e.URL.Paths {
		fmt.Fprintf(&out, "  • %s\n", path)
	}

	if len(e.URL.Parts) > 0 {
		fmt.Fprintln(&out, "Parts:")
		for _, part := range e.URL.Parts {
			fmt.Fprintf(w, "  • %s\t", part.Name)
			fmt.Fprintf(w, "  %s", part.Type)
			if part.Required {
				fmt.Fprint(w, ", required")
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
			fmt.Fprint(w, "\n")
		}
		w.Flush()
	}

	if e.Body != nil {
		fmt.Fprint(&out, "Body: ")
		if e.Body.Required {
			fmt.Fprintf(&out, "required")
		} else {
			fmt.Fprintf(&out, "optional")
		}
		if e.Body.ContentType != "" {
			fmt.Fprintf(&out, " (format: %s)", e.Body.ContentType)
		}
		fmt.Fprintf(&out, "\n")
	}

	return out.String()
}

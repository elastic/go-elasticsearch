package gentests

import (
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/internal/cmd/generate/utils"
)

// DebugInfo returns information about the endpoint as a string.
//
func (tg TestSuite) DebugInfo() string {
	var out strings.Builder

	fmt.Fprintln(&out, strings.Repeat("─", utils.TerminalWidth()))
	fmt.Fprint(&out, "["+tg.Name()+"]\n")
	fmt.Fprintln(&out, strings.Repeat("─", utils.TerminalWidth()))

	if len(tg.Setup) > 0 {
		for _, a := range tg.Setup {
			fmt.Fprint(&out, "[setup] ")
			fmt.Fprint(&out, a.Method()+"(")
			var i int
			for k, v := range a.Params() {
				i++
				fmt.Fprint(&out, ""+k+": "+fmt.Sprintf("%v", v))
				if i < len(a.Params()) {
					fmt.Fprint(&out, ", ")
				}
			}
			fmt.Fprint(&out, ")")
			fmt.Fprint(&out, "\n")
		}
	}

	if len(tg.Teardown) > 0 {
		for _, a := range tg.Teardown {
			fmt.Fprint(&out, "[tdown] ")
			fmt.Fprint(&out, a.Method()+"(")
			var i int
			for k, v := range a.Params() {
				i++
				fmt.Fprint(&out, ""+k+": "+fmt.Sprintf("%v", v))
				if i < len(a.Params()) {
					fmt.Fprint(&out, ", ")
				}
			}
			fmt.Fprint(&out, ")\n")
		}
	}

	for _, t := range tg.Tests {
		if utils.IsTTY() {
			fmt.Fprint(&out, "\x1b[1;2m")
		}
		fmt.Fprintln(&out, t.Name+":")
		if utils.IsTTY() {
			fmt.Fprint(&out, "\x1b[0;2m")
		}
		for _, a := range t.Setup {
			fmt.Fprintf(&out, "  [setup] ")
			fmt.Fprint(&out, a.Method()+"()\n")
		}
		for _, a := range t.Teardown {
			fmt.Fprintf(&out, "  [tdown] ")
			fmt.Fprint(&out, a.Method()+"()\n")
		}
		for _, a := range t.Steps {
			switch a.(type) {
			case Action:
				aa := a.(Action)
				fmt.Fprintf(&out, "  ==> ")
				fmt.Fprint(&out, aa.Method()+"(")
				var i int
				for k, v := range aa.Params() {
					i++
					fmt.Fprint(&out, ""+k+": "+fmt.Sprintf("%v", v))
					if i < len(aa.Params()) {
						fmt.Fprint(&out, ", ")
					}
				}
				fmt.Fprint(&out, ")\n")
			case Assertion:
				aa := a.(Assertion)
				fmt.Fprintf(&out, "  ~~> ")
				fmt.Fprintf(&out, "%q: %s", aa.operation, aa.payload)
				fmt.Fprint(&out, "\n")
			default:
				panic(fmt.Sprintf("Unknown step %T", a))
			}
		}
	}

	fmt.Fprintln(&out, strings.Repeat("─", utils.TerminalWidth()))
	return out.String()
}

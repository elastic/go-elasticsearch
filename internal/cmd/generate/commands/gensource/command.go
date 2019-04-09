package gensource

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/tools/imports"

	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v8/internal/cmd/generate/commands"
	"github.com/elastic/go-elasticsearch/v8/internal/cmd/generate/utils"
)

var (
	input  *string
	output *string
	gofmt  *bool
	color  *bool
	debug  *bool
	info   *bool

	skipAPIRegistry *bool
)

var (
	GitCommit string
	GitTag    string
	EsVersion string
)

func init() {
	input = gensourceCmd.Flags().StringP("input", "i", "", "Path to a folder or files with Elasticsearch REST API specification")
	output = gensourceCmd.Flags().StringP("output", "o", "", "Path to a folder for generated output")
	gofmt = gensourceCmd.Flags().BoolP("gofmt", "f", true, "Format generated output with 'gofmt'")
	skipAPIRegistry = gensourceCmd.Flags().BoolP("skip-registry", "s", false, "Skip generating the API registry (api._.go)")
	color = gensourceCmd.Flags().BoolP("color", "c", true, "Syntax highlight the debug output")
	debug = gensourceCmd.Flags().BoolP("debug", "d", false, "Print the generated source to terminal")
	info = gensourceCmd.Flags().Bool("info", false, "Print the API details to terminal")

	gensourceCmd.MarkFlagRequired("input")
	gensourceCmd.MarkFlagRequired("output")
	gensourceCmd.Flags().SortFlags = false

	commands.RegisterCmd(gensourceCmd)
}

var gensourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Generate Go API from Elasticsearch REST API spec",
	Run: func(cmd *cobra.Command, args []string) {
		command := &Command{
			Input:           *input,
			Output:          *output,
			Gofmt:           *gofmt,
			DebugInfo:       *info,
			DebugSource:     *debug,
			ColorizeSource:  *color,
			SkipAPIRegistry: *skipAPIRegistry,
		}
		err := command.Execute()
		if err != nil {
			utils.PrintErr(err)
			os.Exit(1)
		}
	},
}

// Command represents the "gensource" command.
//
type Command struct {
	Input           string
	Output          string
	Gofmt           bool
	DebugInfo       bool
	DebugSource     bool
	ColorizeSource  bool
	SkipAPIRegistry bool
}

// Execute runs the command.
//
func (cmd *Command) Execute() (err error) {
	var inputFiles []string

	if strings.Contains(cmd.Input, ",") {
		inputFiles = strings.Split(cmd.Input, ",")
	} else {
		inputFiles, err = filepath.Glob(cmd.Input)
		if err != nil {
			return fmt.Errorf("Failed to glob input %q: %s", cmd.Input, err)
		}
	}

	if len(inputFiles) < 1 {
		return fmt.Errorf("No files matching input %q", cmd.Input)
	}

	EsVersion, err = utils.EsVersion(filepath.Dir(inputFiles[0]))
	if err != nil {
		return err
	}
	GitCommit, err = utils.GitCommit(filepath.Dir(inputFiles[0]))
	if err != nil {
		return err
	}
	GitTag, err = utils.GitTag(filepath.Dir(inputFiles[0]))
	if err != nil {
		return err
	}

	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[2m")
	}
	fmt.Fprintln(os.Stderr, strings.Repeat("━", utils.TerminalWidth()))
	fmt.Fprintf(os.Stderr, "Specification: %s", EsVersion)
	fmt.Fprintf(os.Stderr, " (%s", GitCommit)
	if GitTag != "" {
		fmt.Fprintf(os.Stderr, ", %s", GitTag)
	}
	fmt.Fprintf(os.Stderr, ")\n")
	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[0m")
	}

	var endpoints []*Endpoint

	stats := struct {
		n     int
		start time.Time
	}{start: time.Now()}

	for _, fpath := range inputFiles {
		fname := filepath.Base(fpath)
		if fname == "_common.json" {
			continue
		}

		f, err := os.Open(fpath)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := cmd.processFile(f); err != nil {
			return fmt.Errorf("Processing file %q: %s", fname, err)
		}

		f.Seek(0, 0)
		e, err := NewEndpoint(f)
		if err != nil {
			return fmt.Errorf("error creating endpoint for %q: %s", fname, err)
		}
		endpoints = append(endpoints, e)
		stats.n++
	}

	if cmd.Output != "-" && !cmd.SkipAPIRegistry {
		if err := cmd.processAPIConstructor(endpoints); err != nil {
			return fmt.Errorf("error handling API constructor: %s", err)
		}
	}

	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[2m")
	}
	fmt.Fprintln(os.Stderr, strings.Repeat("━", utils.TerminalWidth()))
	fmt.Fprintf(os.Stderr, "Processed %d files in %s\n", stats.n, time.Since(stats.start).Truncate(time.Millisecond))
	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[0m")
	}

	return nil
}

func (cmd *Command) processFile(f *os.File) (err error) {
	var out io.Reader

	fname := filepath.Base(f.Name())

	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[2m")
	}
	fmt.Fprintln(os.Stderr, strings.Repeat("━", utils.TerminalWidth()))
	fmt.Fprintf(os.Stderr, "Processing file %q\n", fname)
	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[0m")
	}

	endpoint, err := NewEndpoint(f)
	if err != nil {
		return fmt.Errorf("error creating endpoint for %q: %s", fname, err)
	}

	gen := Generator{Endpoint: endpoint}

	if cmd.Gofmt {
		out, err = gen.OutputFormatted()
	} else {
		out, err = gen.Output()
	}
	if err != nil {
		if cmd.DebugSource {
			utils.PrintSourceWithErr(out, err)
		}
		return fmt.Errorf("error generating output for %q: %s", fname, err)
	}

	if cmd.DebugInfo {
		if utils.IsTTY() {
			fmt.Fprint(os.Stderr, "\x1b[2m")
		}
		fmt.Fprintf(os.Stderr, gen.Endpoint.DebugInfo())
		if utils.IsTTY() {
			fmt.Fprint(os.Stderr, "\x1b[0m")
		}
	}

	if cmd.DebugSource {
		var src io.Reader
		var buf bytes.Buffer
		tee := io.TeeReader(out, &buf)

		if utils.IsTTY() {
			fmt.Fprint(os.Stderr, "\x1b[2m")
		}
		fmt.Fprintln(os.Stderr, strings.Repeat("━", utils.TerminalWidth()))
		if utils.IsTTY() {
			fmt.Fprint(os.Stderr, "\x1b[0m")
		}
		if cmd.ColorizeSource {
			src, err = utils.Chromatize(tee)
			if err != nil {
				return fmt.Errorf("error syntax highligting the output: %s", err)
			}

			_, err = io.Copy(os.Stderr, src)
			if err != nil {
				return fmt.Errorf("error copying output: %s", err)
			}
		}

		_, err = io.Copy(os.Stderr, out)
		if err != nil {
			return fmt.Errorf("error copying output: %s", err)
		}

		fmt.Fprintf(os.Stderr, "\n\n")

		out = &buf
	}

	if cmd.Output == "-" {
		_, err = io.Copy(os.Stdout, out)
		if err != nil {
			return fmt.Errorf("error copying output: %s", err)
		}
	} else {
		if err := os.MkdirAll(cmd.Output, 0775); err != nil {
			return fmt.Errorf("error creating directory: %s", err)
		}

		fName := filepath.Join(cmd.Output, "api."+gen.Endpoint.Name+".go")
		f, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return fmt.Errorf("error creating file: %s", err)
		}
		_, err = io.Copy(f, out)
		if err != nil {
			return fmt.Errorf("error copying output: %s", err)
		}
		if err := f.Close(); err != nil {
			return fmt.Errorf("error closing file: %s", err)
		}
	}

	return nil
}

func (cmd *Command) processAPIConstructor(endpoints []*Endpoint) (err error) {
	var b bytes.Buffer

	var namespaces = []string{"Cat", "Cluster", "Indices", "Ingest", "Nodes", "Remote", "Snapshot", "Tasks"}

	b.WriteString("// Code generated")
	if EsVersion != "" || GitCommit != "" || GitTag != "" {
		b.WriteString(fmt.Sprintf(" from specification version %s", EsVersion))
		if GitCommit != "" {
			b.WriteString(fmt.Sprintf(" (%s", GitCommit))
			if GitTag != "" {
				b.WriteString(fmt.Sprintf("|%s", GitTag))
			}
			b.WriteString(")")
		}
	}
	b.WriteString(": DO NOT EDIT\n")
	b.WriteString("\n")

	b.WriteString(`package esapi

// API contains the Elasticsearch APIs
//
type API struct {
`)
	for _, n := range namespaces {
		b.WriteString(fmt.Sprintf("\t%[1]s *%[1]s\n", n))
	}

	b.WriteString("\n")

	for _, e := range endpoints {
		skip := false
		name := e.MethodWithNamespace()

		for _, n := range namespaces {
			if strings.HasPrefix(name, n) {
				skip = true
			}
		}
		if !skip {
			b.WriteString(fmt.Sprintf("\t%[1]s %[1]s\n", name))
		}
	}

	b.WriteString(`}` + "\n\n")

	for _, n := range namespaces {
		b.WriteString(`// ` + n + ` contains the ` + n + ` APIs` + "\n")
		b.WriteString(`type ` + n + ` struct {` + "\n")
		for _, e := range endpoints {
			ns := e.Namespace()
			if ns == n {
				b.WriteString(fmt.Sprintf("\t%s %s\n", e.MethodName(), e.MethodWithNamespace()))
			}
		}
		b.WriteString("}\n\n")
	}

	b.WriteString(`// New creates new API
//
func New(t Transport) *API {
	return &API{
`)

	for _, e := range endpoints {
		skip := false
		name := e.MethodWithNamespace()

		for _, n := range namespaces {
			if strings.HasPrefix(name, n) {
				skip = true
			}
		}
		if !skip {
			b.WriteString(fmt.Sprintf("\t\t%[1]s: new%[1]sFunc(t),\n", name))
		}
	}

	for _, n := range namespaces {
		b.WriteString(fmt.Sprintf("\t\t%[1]s: &%[1]s{\n", n))
		for _, e := range endpoints {
			ns := e.Namespace()
			if ns == n {
				b.WriteString(fmt.Sprintf("\t\t\t%s: new%sFunc(t),\n", e.MethodName(), e.MethodWithNamespace()))
			}
		}
		b.WriteString("\t\t},\n")
	}

	b.WriteString(`	}
}
`)

	if cmd.Gofmt {
		out, err := imports.Process(
			"esapi/api._.go",
			b.Bytes(),
			&imports.Options{
				AllErrors:  true,
				Comments:   true,
				FormatOnly: false,
				TabIndent:  true,
				TabWidth:   1,
			})
		if err != nil {
			return err
		}
		b = *bytes.NewBuffer(out)
	}

	fname := filepath.Join(cmd.Output, "api._.go")
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	_, err = io.Copy(f, &b)
	if err != nil {
		return fmt.Errorf("error copying output: %s", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("error closing file: %s", err)
	}
	return nil
}

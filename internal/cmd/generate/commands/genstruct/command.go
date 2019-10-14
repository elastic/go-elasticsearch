// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package genstruct

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"

	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v7/internal/cmd/generate/commands"
	"github.com/elastic/go-elasticsearch/v7/internal/cmd/generate/utils"
)

var (
	GitCommit string
	GitTag    string
	EsVersion string
)

var (
	output *string
	gofmt  *bool
	color  *bool
	debug  *bool
	info   *bool

	pkgNames []string
)

func init() {
	if pkgNamesEnv := os.Getenv("PACKAGE_NAMES"); pkgNamesEnv != "" {
		pkgNames = strings.Split(pkgNamesEnv, ",")
	} else {
		pkgNames = []string{
			"github.com/elastic/go-elasticsearch/v7/esapi",
		}
	}

	output = genapiCmd.Flags().StringP("output", "o", "", "Path to a folder for generated output")
	gofmt = genapiCmd.Flags().BoolP("gofmt", "f", true, "Format generated output with 'gofmt'")
	color = genapiCmd.Flags().BoolP("color", "c", true, "Syntax highlight the debug output")
	debug = genapiCmd.Flags().BoolP("debug", "d", false, "Print the generated source to terminal")
	info = genapiCmd.Flags().Bool("info", false, "Print the API details to terminal")
	genapiCmd.Flags().SortFlags = false

	commands.RegisterCmd(genapiCmd)
}

var genapiCmd = &cobra.Command{
	Use:   "apistruct",
	Short: "Generate the main Go API struct and constructor",
	Run: func(cmd *cobra.Command, args []string) {
		command := &Command{
			Output:         *output,
			Gofmt:          *gofmt,
			DebugSource:    *debug,
			ColorizeSource: *color,
		}
		err := command.Execute()
		if err != nil {
			utils.PrintErr(err)
			os.Exit(1)
		}
	},
}

// Command represents the "genapi" command.
//
type Command struct {
	Output         string
	Gofmt          bool
	DebugSource    bool
	ColorizeSource bool
}

// Execute runs the command.
//
func (cmd *Command) Execute() (err error) {
	EsVersion, err = utils.EsVersion("")
	if err != nil {
		return err
	}
	GitCommit, err = utils.GitCommit("")
	if err != nil {
		return err
	}
	GitTag, err = utils.GitTag("")
	if err != nil {
		return err
	}
	return cmd.processAPIConstructor()
}

func (cmd *Command) processAPIConstructor() (err error) {
	var (
		namespaces []string
		endpoints  []*types.TypeName

		b bytes.Buffer
		i int
		n int
		m int
	)

	namespaces = []string{
		// Core APIs
		//
		"Cat",
		"Cluster",
		"Indices",
		"Ingest",
		"Nodes",
		"Remote",
		"Snapshot",
		"Tasks",

		// XPack APIs
		//
		"CCR",
		"ILM",
		"License",
		"Migration",
		"ML",
		"Monitoring",
		"Rollup",
		"Security",
		"SQL",
		"SSL",
		"Watcher",
		"XPack",
	}

	lpkgs, err := packages.Load(&packages.Config{Mode: packages.LoadTypes}, pkgNames...)
	if err != nil {
		fmt.Printf("Error loading packages: %s\n", err)
		return err
	}

	for _, lpkg := range lpkgs {
		n++

		if cmd.DebugSource {
			fmt.Println(lpkg.Types)
		}

		scope := lpkg.Types.Scope()
		for _, name := range scope.Names() {
			m++
			obj := scope.Lookup(name)
			// Skip unexported objects
			if !obj.Exported() {
				continue
			}

			// Skip non-structs
			structObj, ok := obj.Type().Underlying().(*types.Struct)
			if !ok {
				continue
			}

			// Skip non-request objects
			if !strings.HasSuffix(obj.Name(), "Request") {
				continue
			}

			// Append API struct to endpoints
			endpoints = append(endpoints, obj.(*types.TypeName))

			i++
			fmt.Printf("%-3d | %s{}\n", i, obj.Name())
			for j := 0; j < structObj.NumFields(); j++ {
				field := structObj.Field(j)
				if cmd.DebugSource {
					fmt.Printf("        %s %s\n", field.Name(), field.Type())
				}
			}
		}
	}

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
		name := strings.ReplaceAll(e.Name(), "Request", "")

		// Skip APIs in namespace
		for _, n := range namespaces {
			if strings.HasPrefix(strings.ToLower(name), strings.ToLower(n)) {
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
			name := strings.ReplaceAll(e.Name(), "Request", "")
			if strings.HasPrefix(strings.ToLower(name), strings.ToLower(n)) {
				methodName := strings.ReplaceAll(name, n, "")
				b.WriteString(fmt.Sprintf("\t%s %s\n", methodName, name))
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
		name := strings.ReplaceAll(e.Name(), "Request", "")

		for _, n := range namespaces {
			if strings.HasPrefix(strings.ToLower(name), strings.ToLower(n)) {
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
			name := strings.ReplaceAll(e.Name(), "Request", "")
			if strings.HasPrefix(strings.ToLower(name), strings.ToLower(n)) {
				methodName := strings.ReplaceAll(name, n, "")
				b.WriteString(fmt.Sprintf("\t\t\t%s: new%sFunc(t),\n", methodName, name))
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

	if cmd.DebugSource {
		var src io.Reader
		var buf bytes.Buffer
		tee := io.TeeReader(bytes.NewReader(b.Bytes()), &buf)

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
		} else {
			src = bytes.NewReader(b.Bytes())
		}

		_, err = io.Copy(os.Stderr, src)
		if err != nil {
			return fmt.Errorf("error copying output: %s", err)
		}

		fmt.Fprintf(os.Stderr, "\n\n")
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

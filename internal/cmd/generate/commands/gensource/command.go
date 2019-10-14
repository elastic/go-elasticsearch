// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package gensource

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v7/internal/cmd/generate/commands"
	"github.com/elastic/go-elasticsearch/v7/internal/cmd/generate/utils"
)

var (
	input  *string
	output *string
	gofmt  *bool
	color  *bool
	debug  *bool
	info   *bool
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
	color = gensourceCmd.Flags().BoolP("color", "c", true, "Syntax highlight the debug output")
	debug = gensourceCmd.Flags().BoolP("debug", "d", false, "Print the generated source to terminal")
	info = gensourceCmd.Flags().Bool("info", false, "Print the API details to terminal")

	gensourceCmd.MarkFlagRequired("input")
	gensourceCmd.MarkFlagRequired("output")
	gensourceCmd.Flags().SortFlags = false

	commands.RegisterCmd(gensourceCmd)
}

var gensourceCmd = &cobra.Command{
	Use:   "apisource",
	Short: "Generate the Go APIs from the Elasticsearch REST API specification",
	Run: func(cmd *cobra.Command, args []string) {
		command := &Command{
			Input:          *input,
			Output:         *output,
			Gofmt:          *gofmt,
			DebugInfo:      *info,
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

// Command represents the "gensource" command.
//
type Command struct {
	Input          string
	Output         string
	Gofmt          bool
	DebugInfo      bool
	DebugSource    bool
	ColorizeSource bool
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

		var fName string
		if gen.Endpoint.Type == "core" {
			fName = filepath.Join(cmd.Output, "api."+gen.Endpoint.Name+".go")
		} else {
			fName = filepath.Join(cmd.Output, "api."+gen.Endpoint.Type+"."+gen.Endpoint.Name+".go")
		}

		f, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

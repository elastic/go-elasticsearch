package gentests

//go:generate go run gen_api_registry.go

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

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

	// Populated by 'go generate'
	apiRegistry map[string]map[string]string
)

var (
	GitCommit string
	GitTag    string
	EsVersion string
)

func init() {
	input = gentestsCmd.Flags().StringP("input", "i", "", "Path to a folder or files with Elasticsearch REST API specification")
	output = gentestsCmd.Flags().StringP("output", "o", "", "Path to a folder for generated output")
	gofmt = gentestsCmd.Flags().BoolP("gofmt", "f", true, "Format generated output with 'gofmt'")
	color = gentestsCmd.Flags().BoolP("color", "c", true, "Syntax highlight the debug output")
	debug = gentestsCmd.Flags().BoolP("debug", "d", false, "Print the generated source to terminal")
	info = gentestsCmd.Flags().Bool("info", false, "Print the API details to terminal")

	gentestsCmd.MarkFlagRequired("input")
	gentestsCmd.MarkFlagRequired("output")

	commands.RegisterCmd(gentestsCmd)
}

var gentestsCmd = &cobra.Command{
	Use:   "tests",
	Short: "Generate Go integration tests from Elasticsearch common test suite",
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

// Command represents the "gentests" command.
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
func (cmd *Command) Execute() error {
	if len(apiRegistry) < 1 {
		return fmt.Errorf("API registry in 'api_registry.gen.go' is empty: Did you run go generate?")
	}

	inputFiles, err := filepath.Glob(cmd.Input)
	if err != nil {
		return fmt.Errorf("Failed to glob input %q: %s", cmd.Input, err)
	}

	if len(inputFiles) < 1 {
		return fmt.Errorf("No files matching input %q", cmd.Input)
	}

	EsVersion, err = utils.EsVersion(filepath.Dir(inputFiles[0]))
	if err != nil {
		return err
	}
	if EsVersion == "" {
		return errors.New("Elasticsearch version is empty")
	}

	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[2;1m")
	}
	fmt.Fprintf(os.Stderr, "Elasticsearch %s\n", EsVersion)
	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[0m")
	}

	GitCommit, err = utils.GitCommit(filepath.Dir(inputFiles[0]))
	if err != nil {
		return fmt.Errorf("Failed to get Git commit: %s", err)
	}
	GitTag, _ = utils.GitTag(filepath.Dir(inputFiles[0]))

	stats := struct {
		n     int
		start time.Time
	}{start: time.Now()}

	for _, fpath := range inputFiles {
		var skip bool

		if filepath.Ext(fpath) != ".yml" && filepath.Ext(fpath) != ".yaml" {
			skip = true
		}

		for _, skipFile := range skipFiles {
			if strings.HasSuffix(fpath, skipFile) {
				if utils.IsTTY() {
					fmt.Fprint(os.Stderr, "\x1b[2m")
				}
				fmt.Fprintln(os.Stderr, strings.Repeat("━", utils.TerminalWidth()))
				if utils.IsTTY() {
					fmt.Fprint(os.Stderr, "\x1b[0;33m")
				}
				fmt.Fprintf(os.Stderr, "Skipping %q\n", fpath)
				if utils.IsTTY() {
					fmt.Fprint(os.Stderr, "\x1b[0m")
				}
				skip = true
			}
		}

		if skip {
			continue
		}

		if err := cmd.processFile(fpath); err != nil {
			return fmt.Errorf("Processing file %q: %s", fpath, err)
		}
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

func (cmd *Command) processFile(fpath string) (err error) {
	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[2m")
	}
	fmt.Fprintln(os.Stderr, strings.Repeat("━", utils.TerminalWidth()))
	fmt.Fprintf(os.Stderr, "Processing %q\n", fpath)
	if utils.IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[0m")
	}

	var out io.Reader
	var payloads []TestPayload

	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	patched, err := PatchTestSource(fpath, f)
	if err != nil {
		return fmt.Errorf("cannot patch source: %s", err)
	}

	dec := yaml.NewDecoder(patched)

	for {
		var payload interface{}
		err := dec.Decode(&payload)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("parse error: %s", err)
		}
		payloads = append(payloads, TestPayload{fpath, payload})
	}

	ts := NewTestSuite(fpath, payloads)

	if cmd.DebugInfo {
		if utils.IsTTY() {
			fmt.Fprint(os.Stderr, "\x1b[2m")
		}
		fmt.Fprint(os.Stderr, ts.DebugInfo())
		if utils.IsTTY() {
			fmt.Fprint(os.Stderr, "\n\x1b[0m")
		}
	}

	if ts.Skip {
		if utils.IsTTY() {
			fmt.Fprint(os.Stderr, "\x1b[2m")
		}
		fmt.Fprintf(os.Stderr, "Skipping test suite %q", ts.Name())
		if utils.IsTTY() {
			fmt.Fprint(os.Stderr, "\n\x1b[0m")
		}
		return nil
	}

	gen := Generator{TestSuite: ts}

	if cmd.Gofmt {
		out, err = gen.OutputFormatted()
	} else {
		out, err = gen.Output()
	}
	if err != nil {
		if cmd.DebugSource {
			utils.PrintSourceWithErr(out, err)
		}
		return fmt.Errorf("error generating output: %s", err)
	}

	if cmd.DebugSource {
		var src io.Reader
		var buf bytes.Buffer
		tee := io.TeeReader(out, &buf)

		src, err = utils.Chromatize(tee)
		if err != nil {
			return fmt.Errorf("error syntax highligting the output: %s", err)
		}

		_, err = io.Copy(os.Stderr, src)
		if err != nil {
			return fmt.Errorf("error copying output: %s", err)
		}

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

		fName := filepath.Join(cmd.Output, gen.TestSuite.Filename()+"__test.go")
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

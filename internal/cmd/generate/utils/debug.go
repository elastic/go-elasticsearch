package utils

import (
	"bufio"
	"fmt"
	"go/scanner"
	"go/token"
	"io"
	"os"
	"strings"
)

// PrintSourceWithErr returns source code annotated with location of an error.
//
func PrintSourceWithErr(out io.Reader, err error) {
	if IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[2m")
	}
	if e, ok := err.(scanner.ErrorList); ok {
		fmt.Fprintf(os.Stderr, "\x1b[2m✖ %s\n", e.Error())
	}
	fmt.Fprintln(os.Stderr, strings.Repeat("━", TerminalWidth()))
	if IsTTY() {
		fmt.Fprint(os.Stderr, "\x1b[0m")
	}
	if _, ok := err.(scanner.ErrorList); ok {
		cur := 0
		ferr := err.(scanner.ErrorList)
		type ParseError struct {
			Pos token.Position
			Msg string
		}
		errlines := make(map[int]ParseError)
		for _, e := range ferr {
			errlines[e.Pos.Line] = ParseError{Pos: e.Pos, Msg: e.Msg}
		}
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			cur++

			if _, ok := errlines[cur]; ok {
				if IsTTY() {
					fmt.Fprint(os.Stderr, "\x1b[1;31m")
				}
				fmt.Fprintf(os.Stderr, "%3d| ", cur)
				fmt.Fprintln(os.Stderr, scanner.Text())
			} else {
				if IsTTY() {
					fmt.Fprint(os.Stderr, "\x1b[2m")
				}
				fmt.Fprintf(os.Stderr, "%3d| ", cur)
				if IsTTY() {
					fmt.Fprint(os.Stderr, "\x1b[0m")
				}
				fmt.Fprintln(os.Stderr, scanner.Text())
			}
			if IsTTY() {
				fmt.Fprint(os.Stderr, "\x1b[0m")
			}
			if e, ok := errlines[cur]; ok {
				if IsTTY() {
					fmt.Fprint(os.Stderr, "\x1b[31m")
				}
				fmt.Fprintf(os.Stderr, strings.Repeat(" ", 4))
				for i := 0; i < e.Pos.Column; i++ {
					fmt.Fprintf(os.Stderr, "-")
				}
				fmt.Fprintf(os.Stderr, "^ ")
				fmt.Fprintf(os.Stderr, e.Msg)
				fmt.Fprintf(os.Stderr, "\n")
				if IsTTY() {
					fmt.Fprint(os.Stderr, "\x1b[0m")
				}
			}
		}
		fmt.Fprintln(os.Stderr, "")
	}
}

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

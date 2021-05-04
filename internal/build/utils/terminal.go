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
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	isTTY  bool
	tWidth int
)

func init() {
	isTTY = terminal.IsTerminal(int(os.Stderr.Fd()))
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))
}

// PrintErr prints an error to STDERR.
//
func PrintErr(err error) {
	if isTTY {
		fmt.Fprint(os.Stderr, "\x1b[1;37;41m")
	}
	fmt.Fprintf(os.Stderr, "ERROR: %s", err)
	if isTTY {
		fmt.Fprint(os.Stderr, "\x1b[0m")
	}
	fmt.Fprint(os.Stderr, "\n")
}

// IsTTY returns true when os.Stderr is a terminal.
//
func IsTTY() bool {
	return isTTY
}

// TerminalWidth returns the width of terminal, or zero.
//
func TerminalWidth() int {
	if tWidth < 0 {
		return 0
	}

	return tWidth
}

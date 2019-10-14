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

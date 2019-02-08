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

func IsTTY() bool {
	return isTTY
}

func TerminalWidth() int {
	if tWidth < 0 {
		return 0
	}

	return tWidth
}

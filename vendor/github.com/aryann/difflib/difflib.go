// Copyright 2012 Aryan Naraghi (aryan.naraghi@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package difflib provides functionality for computing the difference
// between two sequences of strings.
package difflib

import (
	"bytes"
	"fmt"
	"math"
)

// DeltaType describes the relationship of elements in two
// sequences. The following table provides a summary:
//
//    Constant    Code   Meaning
//   ----------  ------ ---------------------------------------
//    Common      " "    The element occurs in both sequences.
//    LeftOnly    "-"    The element is unique to sequence 1.
//    RightOnly   "+"    The element is unique to sequence 2.
type DeltaType int

const (
	Common DeltaType = iota
	LeftOnly
	RightOnly
)

// String returns a string representation for DeltaType.
func (t DeltaType) String() string {
	switch t {
	case Common:
		return " "
	case LeftOnly:
		return "-"
	case RightOnly:
		return "+"
	}
	return "?"
}

type DiffRecord struct {
	Payload string
	Delta   DeltaType
}

// String returns a string representation of d. The string is a
// concatenation of the delta type and the payload.
func (d DiffRecord) String() string {
	return fmt.Sprintf("%s %s", d.Delta, d.Payload)
}

// Diff returns the result of diffing the seq1 and seq2.
func Diff(seq1, seq2 []string) (diff []DiffRecord) {
	// Trims any common elements at the heads and tails of the
	// sequences before running the diff algorithm. This is an
	// optimization.
	start, end := numEqualStartAndEndElements(seq1, seq2)

	for _, content := range seq1[:start] {
		diff = append(diff, DiffRecord{content, Common})
	}

	diffRes := compute(seq1[start:len(seq1)-end], seq2[start:len(seq2)-end])
	diff = append(diff, diffRes...)

	for _, content := range seq1[len(seq1)-end:] {
		diff = append(diff, DiffRecord{content, Common})
	}
	return
}

// HTMLDiff returns the results of diffing seq1 and seq2 as an HTML
// string. The resulting HTML is a table without the opening and
// closing table tags. Each table row represents a DiffRecord. The
// first and last columns contain the "line numbers" for seq1 and
// seq2, respectively (the function assumes that seq1 and seq2
// represent the lines in a file). The second and third columns
// contain the actual file contents.
//
// The cells that contain line numbers are decorated with the class
// "line-num". The cells that contain deleted elements are decorated
// with "deleted" and the cells that contain added elements are
// decorated with "added".
func HTMLDiff(seq1, seq2 []string) string {
	buf := bytes.NewBufferString("")
	i, j := 0, 0
	for _, d := range Diff(seq1, seq2) {
		buf.WriteString(`<tr><td class="line-num">`)
		if d.Delta == Common || d.Delta == LeftOnly {
			i++
			fmt.Fprintf(buf, "%d</td><td", i)
			if d.Delta == LeftOnly {
				fmt.Fprint(buf, ` class="deleted"`)
			}
			fmt.Fprintf(buf, "><pre>%s</pre>", d.Payload)
		} else {
			buf.WriteString("</td><td>")
		}
		buf.WriteString("</td><td")
		if d.Delta == Common || d.Delta == RightOnly {
			j++
			if d.Delta == RightOnly {
				fmt.Fprint(buf, ` class="added"`)
			}
			fmt.Fprintf(buf, `><pre>%s</pre></td><td class="line-num">%d`, d.Payload, j)
		} else {
			buf.WriteString(`></td><td class="line-num">`)
		}
		buf.WriteString("</td></tr>\n")
	}
	return buf.String()
}

// numEqualStartAndEndElements returns the number of elements a and b
// have in common from the beginning and from the end. If a and b are
// equal, start will equal len(a) == len(b) and end will be zero.
func numEqualStartAndEndElements(seq1, seq2 []string) (start, end int) {
	for start < len(seq1) && start < len(seq2) && seq1[start] == seq2[start] {
		start++
	}
	i, j := len(seq1)-1, len(seq2)-1
	for i > start && j > start && seq1[i] == seq2[j] {
		i--
		j--
		end++
	}
	return
}

// intMatrix returns a 2-dimensional slice of ints with the given
// number of rows and columns.
func intMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

// longestCommonSubsequenceMatrix returns the table that results from
// applying the dynamic programming approach for finding the longest
// common subsequence of seq1 and seq2.
func longestCommonSubsequenceMatrix(seq1, seq2 []string) [][]int {
	matrix := intMatrix(len(seq1)+1, len(seq2)+1)
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if seq1[len(seq1)-i] == seq2[len(seq2)-j] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = int(math.Max(float64(matrix[i-1][j]),
					float64(matrix[i][j-1])))
			}
		}
	}
	return matrix
}

// compute is the unexported helper for Diff that returns the results of
// diffing left and right.
func compute(seq1, seq2 []string) (diff []DiffRecord) {
	matrix := longestCommonSubsequenceMatrix(seq1, seq2)
	i, j := len(seq1), len(seq2)
	for i > 0 || j > 0 {
		if i > 0 && matrix[i][j] == matrix[i-1][j] {
			diff = append(diff, DiffRecord{seq1[len(seq1)-i], LeftOnly})
			i--
		} else if j > 0 && matrix[i][j] == matrix[i][j-1] {
			diff = append(diff, DiffRecord{seq2[len(seq2)-j], RightOnly})
			j--
		} else if i > 0 && j > 0 {
			diff = append(diff, DiffRecord{seq1[len(seq1)-i], Common})
			i--
			j--
		}
	}
	return
}

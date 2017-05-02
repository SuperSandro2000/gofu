/*******************************************************************************
*
* Copyright 2017 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

//ShowSorted sorts the given lines and prints them on stdout.
func ShowSorted(lines []string) {
	sort.Strings(lines)
	fmt.Println(strings.Join(lines, "\n"))
}

//ShowError prints the given error on stderr if it is non-nil, or returns false otherwise.
func ShowError(err error) bool {
	if err == nil {
		return false
	}
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	return true
}

//FatalIfError prints the given error on stderr and exits with an error code.
func FatalIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: %s\n", err.Error())
		os.Exit(255)
	}
}

var stdin = bufio.NewReader(os.Stdin)

//Prompt prints the question, then waits for the user to press one of the
//possible answer keys. Answer keys will automatically be converted to lower
//case and returned as such.
//
//    choice := Prompt("(y)es or (n)o", []string{"y","n"})
//    //choice is either "y" or "n"
func Prompt(out io.Writer, question string, answers []string) string {
	for idx, answer := range answers {
		answers[idx] = strings.ToLower(answer)
	}

	out.Write([]byte(">> " + strings.TrimSpace(question) + " "))
	for {
		input, err := stdin.ReadString('\n')
		FatalIfError(err)
		input = strings.TrimSpace(input)
		for _, answer := range answers {
			if strings.ToLower(input) == answer {
				return answer
			}
		}

		//user typed gibberish - ask again
		out.Write([]byte("Please type "))
		for idx, answer := range answers {
			if idx > 0 {
				if idx == len(answers)-1 {
					out.Write([]byte(" or "))
				} else {
					out.Write([]byte(", "))
				}
			}
			out.Write([]byte("'" + answer + "'"))
		}
		out.Write([]byte(": "))
	}
}

//ReadLine reads a line from stdin, with whitespace already trimmed.
func ReadLine() string {
	input, err := stdin.ReadString('\n')
	FatalIfError(err)
	return strings.TrimSpace(input)
}

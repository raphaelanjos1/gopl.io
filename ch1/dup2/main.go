// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesWithDuplicates := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesWithDuplicates, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesWithDuplicates, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
			for filename := range filesWithDuplicates[line] {
				fmt.Printf("\t%s", filename)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesWithDuplicates map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if counts[line] > 1 {
			if filesWithDuplicates[line] == nil {
				filesWithDuplicates[line] = make(map[string]bool)
			}
			filesWithDuplicates[line][filename] = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-

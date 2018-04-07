// Dup2 prints the count and text of lines
// that appear more than once in the input.
// It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dup_files := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dup_files, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dup_files, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("line: %s\tfile: %s\n", line, dup_files[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, dup_files map[string]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		dup_files[input.Text()] += filename + " "
	}
}

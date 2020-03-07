/*

练习 1.4： 修改dup2，出现重复的行时打印文件名称。

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LineAttribute ..
type LineAttribute struct {
	Count     int
	FileNames []string
}

func countLines(f *os.File, counts map[string]*LineAttribute) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		_, ok := counts[key]
		if ok {
			counts[key].Count++
		} else {
			counts[key] = new(LineAttribute)
			counts[key].Count = 1
		}

		// uniq FileNames
		if !contains(counts[key].FileNames, f.Name()) {
			counts[key].FileNames = append(counts[key].FileNames, f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

// Contains ..
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	counts := make(map[string]*LineAttribute)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.Count > 1 {
			fmt.Printf("filenames: %v\trepeats: %d\t%s\n", strings.Join(n.FileNames, ","), n.Count, line)
		}
	}
}

//!-

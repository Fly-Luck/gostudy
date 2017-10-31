// 1. use of stdin, files and ioutil
// 2. use of make and map
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//func main() {
//dup1()
//dup2()
//	dup3()
//}

//read stdin
func dup1() {
	counts := make(map[string]int) //map: <string, int>
	countLines(os.Stdin, counts)
	output(counts)
}

//read from files
func dup2() {
	counts := make(map[string]int)
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
	output(counts)
}

//read from files using iotuil & strings
func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) //file->byte slice
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") { //split by newline
			counts[line]++
		}
	}
	output(counts)
}

// count from stdin or files
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() { //use Ctrl + D to stop input
		counts[input.Text()]++
	}
}

//print the map
func output(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			//%d decimal, %x %o %b hex octal binary
			//%f %g %e floats, %t bool, %c char
			//%s string %q quoted string
			//%v any natural format, %T any value
			//%% escape '%'
		}
	}
}

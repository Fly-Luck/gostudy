// flag
// pointer usages with flag pkg
package main

import (
	"flag"
	"fmt"
	"strings"
)

// only when -n=false, the newline is printed
// n aliases the returned pointer that points to -n value
// if not specified, nv = false, n = &nv
var n = flag.Bool("n", false, "omit trailing newline")

// -s sets separator for each input args
// sep aliases the returned pointer that points to -s value
// if not specified, sv = " ", sep = &sv
var sep = flag.String("s", " ", "separator")

/*
func main() {
	echo4()
}
*/

func echo4() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

// 1. both way of imports
// 2. first use of variables and iteration
// 2. first use of os.Args and slice
package main

import (
	"fmt"
	"strings"
)
import "os"

func main() {
	echo1()
	echo2()
	echo3()
}

func echo1() {
	var s string
	var sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("=====Echo version 1")
	fmt.Println("seperated by space:", s)
	fmt.Println("len of args:", len(os.Args))
	fmt.Println("from index 1 to index len:", os.Args[1:len(os.Args)])
	fmt.Println("from index 1, missing end:", os.Args[1:])
	fmt.Println("missing start, to index len(will print Args[0]):", os.Args[:len(os.Args)])
	fmt.Println("missing both start and end(will print Args[0]):", os.Args[:])
}
func echo2() {
	var s, sep string // short for "var s string; var sep string"
	for _, arg := range os.Args[1:] {
		// range iterates two variables: index and item
		// if index is not needed, use blank identifier "_" to avoid compilation error
		s += sep + arg
		sep = " "
	}
	fmt.Println("======Echo version 2")
	fmt.Println(s)
}
func echo3() {
	fmt.Println("======Echo version 3")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

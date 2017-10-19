// 1. both way of imports
// 2. first use of variables and iteration
// 2. first use of os.Args and slice
package main

import (
	"fmt"
)
import "os"

func main() {

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("seperated by space:", s)
	fmt.Println("len of args:", len(os.Args))
	fmt.Println("from index 1 to index len:", os.Args[1:len(os.Args)])
	fmt.Println("from index 1, missing end:", os.Args[1:])
	fmt.Println("missing start, to index len(will print Args[0]):", os.Args[:len(os.Args)])
	fmt.Println("missing both start and end(will print Args[0]):", os.Args[:])
}

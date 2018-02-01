// boiling
// declarations, variables and constants
package main

/*
import (
	"fmt"
)
*/

const boilingF = 212.0

/*
func main() {
	// variable full format: var <name> <type> = <expression>
	var f = boilingF
	// short variable full format: <name> := <expression>, type is determined by expression
	// short variable expression can pass value to variables defined before, but at least a new variable must be defined & pass value
	// e.g. 1. f, err := os.Open(file)
	// anotherF, err := os.Open(anotherFile) //OK, new variable anotherF defined
	// 2. f, err := os.Open(file)
	// f, err := os.Open(anotherFile) //compilation error, no new variable defined
	c := (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)

	const freezingF, boilingF = 32.0, 212.0
	// multiple var/const in one go
	fmt.Printf("%gF = %fC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %fC\n", boilingF, fToC(boilingF))

}
*/

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

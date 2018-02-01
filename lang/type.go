// type
// usage of type declaration
package main

//type <name> <underlying type>, usally package level, if starts with UPPER case, it's accessible across packages

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func main() {
	fmt.Println("Hello World!")
}

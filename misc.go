// 1. switch
// 2. type
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var heads, tails int

func main() {
	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Printf("edge lander!")
	}
	fmt.Printf("%d heads, %d tails", heads, tails)
}

//flip coin
func coinflip() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	if n > 50 {
		return "heads"
	} else if n < 50 {
		return "tails"
	} else {
		return "edge"
	}
}
func signum(x int) int {
	switch { // equivalent to switch true
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}

// "type" <name> "struct"
type Point struct {
	X, Y int
}

package main

import (
	"fmt"
	"math"
)

const val string = "PI Val"
const PI = 3.14

func main() {
	fmt.Println(val, PI)

	const n = 2000000000;
	const d = 3e20 / n;
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
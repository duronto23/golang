package main

import (
	"fmt"
	"math"
)

const cnst string = "Constant String"

func main() {
	fmt.Println(cnst)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

package main

import "fmt"

func main() {
	fmt.Println("First loop")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("Second loop")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("Third loop")
	for {
		fmt.Println("This is a loop!")
		break
	}

	fmt.Println("Fourth loop")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

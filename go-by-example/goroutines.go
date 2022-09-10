package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")
	go f("another")
	time.Sleep(time.Millisecond)
	for j := 0; j < 4; j++ {
		fmt.Println("main")
	}
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

package main

import "fmt"

func accum() func(int) int {
	accumulator := 0 // memoised accumulator

	return func(x int) int {
		accumulator += x
		return accumulator
	}
}

func main() {
	accumulate := accum() // accumulate: func(int) int

	for i := 0; i < 10; i++ {
		fmt.Println(accumulate(i))
	}
}

package main

import "fmt"

func add(nums ...int) int {
	//we can pass as many numbers as we want
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2, 3))
}

package main

import "fmt"

func counter() func() int {
	count := 0 //storing 0 dynamically
	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter()
	fmt.Println(c()) // Output: 1
	fmt.Println(c()) // Output: 2
	fmt.Println(c()) // Output: 3
}

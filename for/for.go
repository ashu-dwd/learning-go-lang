package main

import "fmt"

//for-> only construct in go for looping
func main() {
	//while loop
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}

	//classic for loop
	for i:=1; i<=3; i++{
		fmt.Println(i)
	}

	//range
	for i:= range 10{
		fmt.Println(`value`,i)
	}


}
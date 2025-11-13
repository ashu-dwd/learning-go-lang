package main

import "fmt"

func main() {
	const name = "Raghav Dwivedi"
	const age = 25

	// const name = "Raghav Dwivedi" ///cant redefine the name

	//we can also define multiple constants in group
	const (
		name1 = "Raghav "
		name2 = "Raghav "
		age1  = 25
		age2  = 25
	)


	fmt.Println(name1)
	fmt.Println(age)
}
package main

import (
	"fmt"
)

func main() {
	var nums = []int{1,2,3,4,5}
	for i := 0; i < len(nums); i++ { //simple for loop
		fmt.Println(nums[i])
		fmt.Println(i)
	}

	sum := 0

	for _, num := range nums { //range loop
		sum += num
		fmt.Println(num)
	}

	fmt.Println(sum)

	///iterating over maps

	var m = map[string]int{
		"Raghav": 10,
		"Dwivedi": 20,
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	//iterating over strings
	s := "Raghav Dwivedi"
	for i, letter := range s {
		fmt.Println(i, letter) //letter prints unicode value  
	}

}
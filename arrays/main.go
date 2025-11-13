package main

import "fmt"


func main(){
	var nums[4]int

	nums[0] = 2

	//use len() function to get lenght of array

	fmt.Println(len(nums))
	fmt.Println(nums[0])
	//if you put only one value in array so other index will be zero see line no 9
	fmt.Println(nums)

	//we can even put boolean values in array
	var bools[4]bool
	bools[0] = true
	bools[1] = false
	//as line 15 shows other values will be zero so in boolean 0 is false and 1 is true
	fmt.Println(bools)

	// we can even put string values in array
	var strs[4]string
	strs[0] = "Raghav"
	strs[1] = "Dwivedi"
	// the reason behind empty string is that in go string is a pointer
	fmt.Println(strs)

	// putting int
	 nums4:= [4]int{1,2,3,4}

 	fmt.Println(nums4)

	//2d array
	 arr2d := [2][3]int{{1,2,3},{4,5,6}}

 	fmt.Println(arr2d) 
}
package main

import (
	"fmt"
	"slices"
)

//slices are dynamic arrays
//most used construct in go

func main() {
	// var nums []int
	// fmt.Println(nums)

	// var nums2 = make([]string, 2)
	// fmt.Println(nums2)
	// nums2 = append(nums2, "Raghav")
	// nums2 = append(nums2, "Dwivedi")

	// //capacity means how many elements can be stored in the slice
	// //len means how many elements are stored in the slice
	// fmt.Println(nums2)
	// fmt.Println(len(nums2))
	// fmt.Println(cap(nums2))

	// nums2 = append(nums2, "Raghav")
	// fmt.Println(nums2)
	// fmt.Println(len(nums2))
	// fmt.Println(cap(nums2))

	//copy function copies the elements from one slice to another
	// var nums = make([]int, 0,5)
	// nums = append(nums, 1,2,3,4,5)
	// var nums2 = make([]int,len(nums))
	// copy(nums2,nums)
	// fmt.Println(nums2)

	//slice operator
	var nums = []int{1,2,3,4,5}
	fmt.Println(nums[1:3]) // prints [2,3] means from index 1 to index 3
	fmt.Println(nums[:3]) // prints [1,2,3] means from index 0 to index 3
	fmt.Println(nums[2:]) // prints [3,4,5] means from index 2 to the end
	fmt.Println(nums[:]) // prints [1,2,3,4,5] means from index 0 to the end

	//comparison of arrays using slices
	var nums1 = []int{1,2,3,4,5}
	var nums2 = []int{1,2,3,4,5}
	fmt.Println(slices.Equal(nums1,nums2))

}

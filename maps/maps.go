package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]int) //means make a map with key as string and value as int

	// m["Raghav"] = 10
	// m["Dwivedi"] = 20
	// m["Raghav"] = 30

	// for key, value := range m {
	// 	println(key, value)
	// }
	// println(m) // print something abnormal like 0xc0000626b0 means its a pointer
	// println(m["raghavdwd"]) // if key is not found then it will return zero value
	// fmt.Println(m) // prints map[Raghav:30 Dwivedi:20]
	// delete(m, "Raghav") // deletes the key from the map
	// fmt.Println(m) // prints map[Dwivedi:20]

	//short hand for map
	// m := map[string]int{
	// 	"Raghav": 10,
	// 	"Dwivedi": 20,
	// }
	// //fmt.Println(m)

	// //checking if key exists
	// k, ok := m["Raghav"] // _ means discard the value and ok means if key exists
	// //fmt.Println(ok)
	// fmt.Println(k) // prints 10 means value of key Raghav
	// if ok {
	// 	fmt.Println("key exists")
	// } else {
	// 	fmt.Println("key does not exist")
	// }

	//comparison of maps
	m1 := map[string]int{
		"Raghav": 10,
		"Dwivedi": 20,
	}
	m2 := map[string]int{
		"Raghav": 10,
		"Dwivedi": 20,
	}
	fmt.Println(maps.Equal(m1,m2))
}
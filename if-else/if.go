package main

import "fmt"

func main(){
	age:= 15

	if age>16 {
		fmt.Println("You can hang out!")
	} else {
		fmt.Println("Go and study!")
	}

	var role = "admin"
	var hasPermissions = true
	if role == "admin" && hasPermissions {
		fmt.Println("You are your own boss! You can do everything!")
	}
}
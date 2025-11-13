package main

import "fmt"

func main() {
	var a int = 1
	var b float64 = 1.0
	var c bool = true
	var d byte = 'b'
	var e string = "Hello World"
	var f string = `Hello World`
	var g bool = true

	//if we dont assign data type to variable, go will infer the data type
	//var a = 1
	//var b = 1.0
	//var c = true
	//var d = 'b'

	//shorthand syntax

	name := "Raghav Dwivedi"
	age := 25

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
} 

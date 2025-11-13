package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func getLang() (string, string, string, bool) {
	return "Go", "Python", "JavaScript", true
}

func main() {
	addResult := add(1, 2)
	subtractResult := subtract(3, 4)
	multiplyResult := multiply(5, 6)
	divideResult := divide(7, 8)

	fmt.Println("Addition:", addResult)
	fmt.Println("Subtraction:", subtractResult)
	fmt.Println("Multiplication:", multiplyResult)
	fmt.Println("Division:", divideResult)

	goLang, python, js, isSupported := getLang()
	fmt.Println("Languages:", goLang, python, js)
	if isSupported {
		fmt.Println("Supported")
	} else {
		fmt.Println("Not Supported")
	}
	fmt.Println("Hello, World!")
}

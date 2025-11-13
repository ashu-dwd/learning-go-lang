package main

import (
	"fmt"
	"time"
)

func main(){
	i:=4

 switch i {
	case 1:
		fmt.Println("1-sunday")
	case 2: 
	  fmt.Println("2-Monday")
	default:
		fmt.Println("Other")
	}

	//multiple condition switch case

	switch time.Now().Weekday() {
	case time.Monday, time.Friday:
		fmt.Println("Enjoy your Monday and Friday")
	default:
		fmt.Println("Dont Enjoy! You are here to study.")
	}

	//type switch
	whoAmI:= func (i interface{})  {
		switch t:= i.(type){
		case int:
			fmt.Println("I am Int")
		case float64:
			fmt.Println("I am float64")
		default:
			fmt.Println("I am none!", t)
		}

	}

	whoAmI(5.0)

}
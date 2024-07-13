package main

import (
	"fmt"
)

func typeDetermine(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	case bool:
		fmt.Println("Type is bool")
	case chan int:
		fmt.Println("Type is channel of int")
	case chan string:
		fmt.Println("Type is channel of string")
	case chan bool:
		fmt.Println("Type is channel of bool")
	default:
		fmt.Println("Unknown type")
	}

}

func main() {
	v1 := 42
	v2 := "hello"
	v3 := true
	v4 := make(chan int)
	v5 := make(chan string)
	v6 := make(chan bool)

	typeDetermine(v1)
	typeDetermine(v2)
	typeDetermine(v3)
	typeDetermine(v4)
	typeDetermine(v5)
	typeDetermine(v6)
}

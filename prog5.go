package main

import (
	"fmt"
)

func divide(a, b int) int {
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()
	if b == 0 {

		panic("division by zero")
	}
	return a / b
}

func main() {
	result := divide(5, 0)
	fmt.Println(result)
}

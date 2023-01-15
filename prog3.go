package main

import "fmt"

func modifyValue(p *int) {
	defer func() {
		fmt.Println("Inside 2nd output: *p =", *p)
	}()
	*p = 20
	fmt.Println("Inside 1st output: *p =", *p)
}

func main() {
	var i int = 10
	var p *int = &i
	defer func() {
		fmt.Println("Inside main's 3rd output: i =", i)
		fmt.Println("Inside main's 4th: *p =", *p)
	}()
	modifyValue(p)
	fmt.Println("Inside main's 1st output: i =", i)
	fmt.Println("Inside main's 2nd output: *p =", *p)
}

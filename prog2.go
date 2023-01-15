package main

import "fmt"

func main() {
	var i int = 10
	var p *int = &i
	defer func() {
		fmt.Println("defer i:", i)
		fmt.Println(" defer *p:", *p)
	}()
	*p = 20
	fmt.Println("i:", i)
	fmt.Println("*p:", *p)
}

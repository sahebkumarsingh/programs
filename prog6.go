package main

import (
	"fmt"
	"time"
)

func printHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")

	}
}

func printWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("World")

	}
}

func main() {

	go printHello() //this two functions acts as a thread that will run concurrently with main program
	go printWorld()
	time.Sleep(11 * time.Second)
	fmt.Println("hii")
}

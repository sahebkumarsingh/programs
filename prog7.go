package main

import (
	"fmt"
	"sync"
	"time"
)

func myfunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("finished executing goroutine")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("GO waitgrup tutorial")
	wg.Add(1)
	go myfunc(&wg)
	wg.Wait() //blocks until zero
	fmt.Println("finished executing my program")
}

package main

import (
	"fmt"
	"sync"
)

var sharedData []int
var mutex sync.Mutex
var wg sync.WaitGroup

func addData(data int) {
	mutex.Lock()
	defer mutex.Unlock()
	sharedData = append(sharedData, data)
	fmt.Println("Added:", data)
}

func main() {

	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			addData(i)
		}
		wg.Done()
	}()

	go func() {
		for i := 10; i < 20; i++ {
			addData(i)
		}
		wg.Done()
	}()
	wg.Wait()

}

package main

import (
	"fmt"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)

	}
}
func main() {
	say("hey")
	say("there")

}

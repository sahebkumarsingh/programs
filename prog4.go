package main

import (
	"fmt"
	"io/ioutil" //used for making the response
	"net/http"  //used for making GET request
)

func main() {
	fmt.Println(" web req")
	response, err := http.Get("https://lco.dev") //makes GET request to the given URL and assign the response in response variable
	if err != nil {
		panic(err) //panic shows the error messg and panic is always at last even after defer
	}
	fmt.Println("response : ", response)
	defer response.Body.Close()
	datatypes, err := ioutil.ReadAll(response.Body) //reads the response of the body and assigns to datatype

	if err != nil {
		panic(err)
	}

	fmt.Println(string(datatypes))
}

package main

import (
	"net/http"

	"github.com/pluralsight/webservices/controllers"
)

func main() {

	controllers.Registercontrollers()
	http.ListenAndServe(":4000", nil)
}

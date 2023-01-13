package controllers

import "net/http"

func Registercontrollers() {
	uc := newusercontroller()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

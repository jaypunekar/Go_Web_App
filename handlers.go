package main

import (
	"net/http"
)

// Home is home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

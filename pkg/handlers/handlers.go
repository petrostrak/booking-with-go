package handlers

import (
	"fmt"
	"net/http"
)

// Home handler will return a simple msg to Writer
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

// About handler will return the info regarding the site
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A bookings and reservation system")
}

package controllers

import (
	"fmt"
	"net/http"
)

// GetHome controller
func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

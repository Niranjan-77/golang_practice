package main

import (
	"fmt"
	"net/http"
)

type APIHandler struct{}

func (APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've reached the API endpoint!\n")

}
func main() {
	http.ListenAndServe(":8080", &APIHandler{})
}

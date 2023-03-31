package main

import (
	"fmt"
	"net/http"
)

// This is the function that is called anytime someone comes to the web server
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to this awesome website</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)

	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", nil)

}

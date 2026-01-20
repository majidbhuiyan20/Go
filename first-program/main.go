package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Majid Bhuiyan. I am software Engineer")
}

func main() {
	mux := http.NewServeMux()              // mux call router
	mux.HandleFunc("/hello", helloHandler) //hello is route
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

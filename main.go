package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// root path calls getRoot handler function
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got root request\n")
	io.WriteString(w, "This is my website!\n")
}

// /hello path calls getHello handler fucntion
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "Hello HTTP!\n")
}

func main() {
	// each HandleFunc call sets up a specific request path
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
		// checks for any non-ErrServerClosed error
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

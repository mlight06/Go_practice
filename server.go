package main

import (
	"fmt"
	"errors"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
} else if err != nil {
	fmt.Printf("error starting server: %s\n", err)
	os.Exit(1)
	}
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "this is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request \n")
	io.WriteString(w, "Hello, Http!\n")
}
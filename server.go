package main

import (
	"fmt"
	"errors"
	"io"
	"net/http"
	"os"
)

func main() {

	// creates a custom, server multiplexer
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	// mux.HandleFunc("/", serveRoot)
	mux.HandleFunc("/hello", getHello)
	mux.Handle("/", fileServer)

	err := http.ListenAndServe(":3333", mux)

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
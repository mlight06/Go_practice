package main

import (
	"fmt"
	"errors"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "this is my website!]n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request \n")
	io.WriteString(w, "Hello, Http!\n")
}
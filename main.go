package main

import (
	"fmt"
	"net/http"
)

// main heap
func main() {
	mux := http.NewServeMux()       // creates a new request multiplexer
	mux.HandleFunc("/", handleRoot) // template used to control the traffic

	// at this point, the server as not started yet.

	fmt.Println("Server listening to :8080")
	http.ListenAndServe(":8080", mux) // starts the server at localhost:8080. takes two parameters; port and mux

}

func handleRoot(
	w http.ResponseWriter, // responsable for constructing a response for the client; (example: send header; send response;)
	r *http.Request, // contains information of the request; (example: body; headers; url;)
) {
	fmt.Fprintf(w, "Hello World.")
}

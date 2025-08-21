package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User) // create map to sim a user db

// this creates a mutual exclusive flag
var cacheMutex sync.RWMutex // blocks all reading/writing when locked; mutex is generally a safe way to synchronize data in a multithreaded application.

// main heap
func main() {
	mux := http.NewServeMux()       // creates a new request multiplexer
	mux.HandleFunc("/", handleRoot) // template used to control the traffic

	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users/{id}", getUser)

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

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// (1) create a new decoder based off body info in the request. (2) decode this body info for the user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// (3) validate username
	if user.Name == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	// (4) new id; insert user into local memory db
	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id")) // (1) obtain id string. (2) turn value into an int
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	cacheMutex.RLock()
	user, ok := userCache[id] // (3) check for user
	cacheMutex.RUnlock()

	if !ok {
		http.Error(
			w,
			"user not found",
			http.StatusNotFound,
		)
		return
	}

	j, err := json.Marshal(user) // (4) retrieve json representation of this user for this endpoint
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK) // (5) returns 200 + json representation of retrieved user
	w.Write(j)                   // (6) write the Marshall user to the response writer
}

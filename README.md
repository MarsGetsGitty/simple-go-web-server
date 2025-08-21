# simple-go-web-server

A minimal, no-frills web server written in Go—designed as a personal template or starting point for learning web services in Go.

---

##  Overview

This project provides a basic HTTP server implemented using Go's standard library. It demonstrates:

- `http.NewServeMux()` for request routing
- `HandleFunc` for mapping endpoints
- JSON payload parsing and validation
- In-memory storage using a `map[int]User`

It's perfect as a reference or skeleton for your own Go web projects.

---

##  Features

| Feature                           | Description                                        |
|----------------------------------|----------------------------------------------------|
| HTTP routing                     | Multiplexer (`ServeMux`) with route handlers      |
| POST `/users` endpoint           | Accepts JSON input and adds users to local memory |
| Simple root handler              | `"Hello World."` response on `GET /`              |
| No external dependencies         | Implements everything with Go standard library     |

---

##  Getting Started

Clone the repo and switch to the `dev` branch:

```bash
git clone https://github.com/MarsGetsGitty/simple-go-web-server.git
cd simple-go-web-server
git checkout dev
````

Build and run:

```bash
go run main.go
```

You should see:

```
Server listening on :8080
```

### Test the endpoints

* **GET /**
  Open `http://localhost:8080/` — you’ll get:

  ```
  Hello World.
  ```

* **POST /users**
  Create a user:

  ```bash
  curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{"name": "Alice"}'
  ```

  Expect a `201 Created` with JSON:

  ```json
  {
    "id": 1,
    "name": "Alice"
  }
  ```

---

## Example Code Snippet

```go
type User struct {
  Name string `json:"name"`
}

var userCache = make(map[int]User)

func createUser(w http.ResponseWriter, r *http.Request) {
  var user User
  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  if user.Name == "" {
    http.Error(w, "username is required", http.StatusBadRequest)
    return
  }
  id := len(userCache) + 1
  userCache[id] = user
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  _ = json.NewEncoder(w).Encode(map[string]any{"id": id, "name": user.Name})
}
```

---

## Why Use This?

* **Educational**: A clean reference for Go web basics.
* **Lightweight**: No frameworks, no external dependencies—just the standard library.
* **Boilerplate Ready**: A straightforward template you can expand (e.g., add middleware, database support, route grouping).

---

## Next Steps

Some ideas to build on:

* Add persistent storage (e.g., file or SQLite)
* Support more HTTP methods (GET, PUT, DELETE)
* Implement logging and request validation
* Use `go.mod` to turn this into a reusable module
* Add graceful shutdown and middleware (e.g., CORS, JSON validation)

---
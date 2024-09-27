package main

import (
	"Backend-Go/internal/db"
	"Backend-Go/internal/router"
	"Backend-Go/internal/user"
	"log"
	"net/http"
)

func main() {

	// Initialize user repository and service
	dbInstance := db.InitDB()
	userRepo := user.UserRepository{DB: dbInstance}
	userService := user.UserService{Repo: userRepo}
	userHandler := user.UserHandler{Service: userService}

	// Set up the router
	r := router.NewRouter(userHandler)

	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request at /")
	w.Write([]byte("hello"))
}

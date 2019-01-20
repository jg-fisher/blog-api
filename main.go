package main

import (
	"encoding/json"
	"github/jg-fisher/blog-api/models"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Initialize database context for package models
	var connString string
	connString = "john:1234@tcp(127.0.0.1:3306)/blog"
	models.InitDB(connString)

	// CORS config
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// TODO: set this allowed origin in env, called with os.Getenv
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, methodsOk)(router)))
}

func getPosts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	posts, err := models.AllPosts()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

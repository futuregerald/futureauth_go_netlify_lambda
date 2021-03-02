package main

import (
	"log"
	"net/http"

	"github.com/futuregerald/futureauth-go/src/functions/signup/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Print("Error loading .env file")
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/.netlify/functions/signup", api.LambdaHandler)
	log.Fatal(http.ListenAndServe(":3000", r))
}

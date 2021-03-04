package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/futuregerald/futureauth-go/src/functions/signup/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Print("Error loading .env file")
		}
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	client := api.New()
	log.Print(client.MongoURI)
	r.Post("/.netlify/functions/signup", client.LambdaHandler)
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		log.Fatal(http.ListenAndServe(":3030", r))
	} else {
		log.Fatal(gateway.ListenAndServe(":3000", r))
	}
}

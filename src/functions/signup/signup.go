package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/futuregerald/futureauth-go/src/functions/db"
	"github.com/futuregerald/futureauth-go/src/functions/signup/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
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
	dbClient, err := db.New()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Unable to create database connnection"))
	}
	a := api.New(dbClient)
	if err != nil {
		log.Print(errors.Wrap(err, "Unable to start API"))
	}
	r.Post("/.netlify/functions/signup", a.LambdaHandler)
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		log.Print("starting signup function")
		log.Fatal(http.ListenAndServe(":3030", r))
	} else {
		log.Fatal(gateway.ListenAndServe(":3000", r))
	}
}

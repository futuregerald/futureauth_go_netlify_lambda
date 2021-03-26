package db

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO create method on user to validate password

func Connect(uri string) error {
	return mgm.SetDefaultConfig(nil, "local_dev", options.Client().ApplyURI(uri))
}

func New(mongoURI string) error {
	if mongoURI != "" {
		return Connect(mongoURI)
	}
	log.Print("No mongoDB URI provided. Api starting without a data store")
	return nil
}

package db

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO create method on user to validate password

func Connect(uri string) {
	if err := mgm.SetDefaultConfig(nil, "local_dev", options.Client().ApplyURI(uri)); err != nil {
		log.Print(err)
	}
}

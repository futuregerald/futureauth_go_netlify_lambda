package db

import (
	"encoding/json"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Email            string             `json:"name" bson:"name"`
	Tenant           primitive.ObjectID `json:"tenantID" bson:"tenantID,omitempty"`
	Password         string             `json:"password" bson:"password"`
	appMetaData      json.RawMessage    `json:"appMetaData" bson:"appMetaData"`
	userMetaData     json.RawMessage    `json:"userMetaData" bson:"appMetaData"`
	confirmed        bool               `json:"confirmed" bson:"confirmed"`
	isAdmin          bool               `json:"isAdmin" bson:"isAdmin"`
	disabled         bool               `json:"disabled" bson:"disabled"`
	roles            []string           `json:"roles" bson:"roles"`
}

// TODO populate NewUser method as per User struct
func NewUser(name string, pages int) *User {
	return &User{}
}

// TODO fill out creating to use argon2 to create password hash
func (model *User) Creating() error {
	// Call to DefaultModel Creating hook
	if err := model.DefaultModel.Creating(); err != nil {
		return err
	}

	return nil
}

// TODO create method on user to validate password

func Connect(uri string) {
	if err := mgm.SetDefaultConfig(nil, "local_dev", options.Client().ApplyURI(uri)); err != nil {
		log.Print(err)
	}
}

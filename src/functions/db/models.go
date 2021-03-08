package db

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PasswordConfig is used to generate the argon2 password hash
type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

type User struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Email            string             `json:"name" bson:"name"`
	Tenant           primitive.ObjectID `json:"tenantID" bson:"tenantID,omitempty"`
	Password         string             `json:"-" bson:"password"`
	AppMetaData      string             `json:"appMetaData" bson:"appMetaData"`
	UserMetaData     string             `json:"userMetaData" bson:"userMetaData"`
	Confirmed        bool               `json:"confirmed" bson:"confirmed"`
	IsAdmin          bool               `json:"isAdmin" bson:"isAdmin"`
	Disabled         bool               `json:"disabled" bson:"disabled"`
	Roles            []string           `json:"roles" bson:"roles"`
}

package api

import (
	"encoding/json"
)

// SignupData is what's sent to the signup endpoint and used to create the User model
type SignupData struct {
	ID           string          `json:"id,omitempty"`
	Email        string          `json:"email" bson:"name" validate:"required,email,max=30,min=6"`
	Tenant       string          `json:"tenantID" bson:"tenantID,omitempty"`
	Password     string          `json:"password" bson:"password" validate:"required,max=30,min=6"`
	AppMetaData  json.RawMessage `json:"appMetaData" bson:"appMetaData"`
	UserMetaData json.RawMessage `json:"userMetaData" bson:"appMetaData"`
	Confirmed    bool            `json:"confirmed" bson:"confirmed"`
	IsAdmin      bool            `json:"isAdmin" bson:"isAdmin"`
	Disabled     bool            `json:"disabled" bson:"disabled"`
	Roles        []string        `json:"roles" bson:"roles"`
}

package futureauth

import (
	"encoding/json"
)

// SignupData is what's sent to the signup endpoint and used to create the User model
type SignupData struct {
	ID           string          `json:"id,omitempty"`
	Email        string          `json:"email" bson:"name" validate:"required,email,max=30,min=6"`
	Tenant       string          `json:"tenantID"`
	Password     string          `json:"password"validate:"required,max=30,min=6"`
	AppMetaData  json.RawMessage `json:"appMetaData"`
	UserMetaData json.RawMessage `json:"userMetaData"`
	Confirmed    bool            `json:"confirmed"`
	IsAdmin      bool            `json:"isAdmin"`
	Disabled     bool            `json:"disabled"`
	Roles        []string        `json:"roles"`
}

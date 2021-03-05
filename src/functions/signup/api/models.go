package api

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
)

// RequestData is the inbound json body this endpoint expects
type SignupData struct {
	AppMetadata  json.RawMessage `json:"appMetadata"`
	Email        string          `json:"email"`
	Password     string          `json:"password"`
	Roles        []string        `json:"roles"`
	UserMetadata json.RawMessage `json:"userMetadata"`
}

type Client struct {
	Db       *mongo.Client
	MongoURI string
}

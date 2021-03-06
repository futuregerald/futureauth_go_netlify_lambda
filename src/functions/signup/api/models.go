package api

import (
	"encoding/json"
)

// RequestData is the inbound json body this endpoint expects
type SignupData struct {
	AppMetadata  json.RawMessage `json:"appMetadata"`
	Email        string          `json:"email"`
	Password     string          `json:"password"`
	Roles        []string        `json:"roles"`
	UserMetadata json.RawMessage `json:"userMetadata"`
}

package api

// RequestData is the inbound json body this endpoint expects
type SignupData struct {
	AppMetadata struct {
		Random string `json:"random"`
	} `json:"appMetadata"`
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Roles        []string `json:"roles"`
	UserMetadata struct {
		Random string `json:"random"`
	} `json:"userMetadata"`
}

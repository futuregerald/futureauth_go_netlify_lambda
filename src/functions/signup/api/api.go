package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/futuregerald/futureauth-go/src/functions/db"
	"github.com/futuregerald/futureauth-go/src/functions/helpers"
	"github.com/go-playground/validator/v10"
)

func (c *Client) LambdaHandler(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	reqBody, err := ioutil.ReadAll(r.Body)
	log.Print(len(reqBody))
	if err != nil {
		log.Print(err)
		if err := helpers.SendJSON(w, http.StatusUnprocessableEntity, "Invalid body"); err != nil {
			log.Print(err)
		}
		return
	}
	if len(reqBody) == 0 {
		if err := helpers.SendJSON(w, http.StatusUnprocessableEntity, "No payload present. Please pass at least a username and password."); err != nil {
			log.Print(err)
		}
		return
	}

	var req SignupData
	if err := json.Unmarshal(reqBody, &req); err != nil {
		log.Print(err)
		if err := helpers.SendJSON(w, http.StatusUnprocessableEntity, "Invalid body"); err != nil {
			log.Print(err)
		}
		return
	}
	if err := validate.Struct(&req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if err != nil {
			if err := helpers.SendJSON(w, http.StatusUnprocessableEntity, validationErrors.Error()); err != nil {
				log.Print(err)
			}
			return
		}
	}

	if newUser, err := db.NewUser(c.dbClient, req.Email, req.Tenant, req.Password, req.Confirmed, req.IsAdmin, req.Disabled, req.AppMetaData, req.UserMetaData, req.Roles); err != nil {
		log.Print(err)
		if err := helpers.SendJSON(w, http.StatusInternalServerError, "Unable to create new user!"); err != nil {
			log.Print(err)
		}
	} else {
		log.Print(newUser)
		if err := helpers.SendJSON(w, http.StatusOK, "Successfully created the new user!"); err != nil {
			log.Print(err)
		}
	}

	if err := helpers.SendJSON(w, http.StatusOK, "this endpoint works!"); err != nil {
		log.Print(err)
	}

}

func New(dbc db.DBClient) *Client {
	return &Client{
		dbClient: dbc,
	}
}

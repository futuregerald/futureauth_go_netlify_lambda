package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/futuregerald/futureauth-go/src/functions/futureauth"
	"github.com/futuregerald/futureauth-go/src/functions/helpers"
)

func LambdaHandler(w http.ResponseWriter, r *http.Request) {

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

	var req futureauth.SignupData
	if err := json.Unmarshal(reqBody, &req); err != nil {
		log.Print(err)
		if err := helpers.SendJSON(w, http.StatusUnprocessableEntity, "Invalid body"); err != nil {
			log.Print(err)
		}
		return
	}

	if err, user := futureauth.CreateUser(req); err != nil {
		log.Print("create user error: ", err)
		if err := helpers.SendJSON(w, http.StatusUnprocessableEntity, "User created successfully!"); err != nil {
			log.Print("making response error: ", err)
		}
		return
	} else {
		if err := helpers.SendJSON(w, http.StatusOK, user); err != nil {
			log.Print("making response error: ", err)
		}
	}

}

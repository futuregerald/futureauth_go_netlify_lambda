package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/futuregerald/futureauth-go/src/functions/db"
	"github.com/futuregerald/futureauth-go/src/functions/helpers"
)

func LambdaHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		if err := helpers.SendJSON(w, http.StatusUnprocessableEntity, "Invalid body"); err != nil {
			log.Print(err)
		}
		return
	}

	var parsedBody SignupData
	if err := json.Unmarshal(reqBody, &parsedBody); err != nil {
		log.Print(err)
		if err := helpers.SendJSON(w, http.StatusUnprocessableEntity, "Invalid body"); err != nil {
			log.Print(err)
		}
		return
	}
	if err := helpers.SendJSON(w, http.StatusOK, "this endpoint works!"); err != nil {
		log.Print(err)
	}

}

func New() error {
	mongoURI := os.Getenv("MONGO_URI")
	db.Connect(mongoURI)
	if mongoURI == "" {
		return db.Connect(mongoURI)
	}
	log.Print("No mongoDB URI provided. Api starting without a data store")
	return nil
}

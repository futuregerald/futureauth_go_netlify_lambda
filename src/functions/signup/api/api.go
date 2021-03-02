package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

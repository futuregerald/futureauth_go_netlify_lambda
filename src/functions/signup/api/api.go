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

func (*Client) LambdaHandler(w http.ResponseWriter, r *http.Request) {
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

func New() *Client {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return &Client{}
	}
	return &Client{
		Db: db.Connect(mongoURI),
	}
}

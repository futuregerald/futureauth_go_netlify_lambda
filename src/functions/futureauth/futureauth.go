package futureauth

import "github.com/futuregerald/futureauth-go/src/functions/futureauth/db"

func New(mongoURI string) error {
	return db.New(mongoURI)
}

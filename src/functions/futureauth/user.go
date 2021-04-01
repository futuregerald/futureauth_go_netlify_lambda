package futureauth

import (
	"github.com/futuregerald/futureauth-go/src/functions/futureauth/db"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func CreateUser(signupData SignupData) (error, *db.User) {
	validate := validator.New()
	if err := validate.Struct(&signupData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if err != nil {
			return errors.Wrap(err, validationErrors.Error()), &db.User{}
		}
	}

	if newUser, err := db.NewUser(signupData.Email, signupData.Tenant, signupData.Password, signupData.Confirmed, signupData.IsAdmin, signupData.Disabled, signupData.AppMetaData, signupData.UserMetaData, signupData.Roles); err != nil {
		return errors.Wrap(err, "User creation failed"), &db.User{}
	} else {
		return nil, newUser
	}
}

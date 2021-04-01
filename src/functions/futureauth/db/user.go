package db

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/kamva/mgm/v3"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
)

func getArgonConfig() *PasswordConfig {
	return &PasswordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}
}

func (u *User) generatePasswordHash(c *PasswordConfig) (string, error) {

	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(u.Password), salt, c.time, c.memory, c.threads, c.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
	return full, nil
}

func (u *User) verifyPassword(password string) (bool, error) {

	parts := strings.Split(u.Password, "$")

	c := &PasswordConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	c.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	return (subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1), nil
}

func NewUser(email, tenant, password string, confirmed, isAdmin, disabled bool, appMetaData, userMetaData json.RawMessage, roles []string) (*User, error) {
	newUser := User{
		Email:     email,
		Password:  password,
		Confirmed: confirmed,
		IsAdmin:   isAdmin,
		Disabled:  disabled,
		Roles:     roles,
	}
	if appMetaData != nil {
		stringAppMetaData, err := json.Marshal(&appMetaData)
		if err != nil {
			return &User{}, errors.Wrap(err, "Unable to marshal appMetaData")
		}
		log.Print("stringAppMetaData", stringAppMetaData)
		newUser.AppMetaData = string(stringAppMetaData)
	}
	if userMetaData != nil {
		stringUserMetaData, err := json.Marshal(&userMetaData)
		if err != nil {
			return &User{}, errors.Wrap(err, "Unable to marshal userMetadata")
		}
		newUser.UserMetaData = string(stringUserMetaData)
	}

	if tenant != "" {
		tenantID, err := primitive.ObjectIDFromHex(tenant)
		if err != nil {
			return &User{}, errors.Wrap(err, "Invalid Tenant ID")
		}
		newUser.Tenant = tenantID
	}
	if err := mgm.Coll(&newUser).Create(&newUser); err != nil {
		return &User{}, errors.Wrap(err, "Unable to create new user")
	}
	return &newUser, nil
}

// This is a pre-save hook that hashes the password
func (user *User) Saving() error {
	// Call to DefaultModel Creating hook
	if err := user.DefaultModel.Creating(); err != nil {
		return err
	}
	passwordHash, err := user.generatePasswordHash(getArgonConfig())
	if err != nil {
		log.Print("Unable to hash password")
		user.Password = ""
		return errors.Wrap(err, "Unable to hash password")
	}
	user.Password = passwordHash
	return nil
}

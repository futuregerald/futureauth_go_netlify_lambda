package db

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
)

func (*User) getArgonConfig() *PasswordConfig {
	return &PasswordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}
}

func (*User) GeneratePassword(c *PasswordConfig, password string) (string, error) {

	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
	return full, nil
}

func (*User) ComparePassword(password, hash string) (bool, error) {

	parts := strings.Split(hash, "$")

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

// TODO populate NewUser method as per User struct
func NewUser(email, tenant string, Tenant int) *User {
	var tenantID primitive.ObjectID
	var err error
	if tenant != "" {
		tenantID, err = primitive.ObjectIDFromHex(tenant)
		if err != nil {

		}
	}

	return &User{
		Email:  email,
		Tenant: tenantID,
	}
}

// TODO fill out creating to use argon2 to create password hash
func (user *User) Creating() error {
	// Call to DefaultModel Creating hook
	if err := user.DefaultModel.Creating(); err != nil {
		return err
	}
	passwordHash, err := user.GeneratePassword(user.getArgonConfig(), user.Password)
	if err != nil {
		log.Print("Unable to hash password")
		user.Password = ""
		return errors.Wrap(err, "Unable to hash password")
	}
	user.Password = passwordHash
	return nil
}

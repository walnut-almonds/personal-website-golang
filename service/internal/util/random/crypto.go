package random

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/xerrors"
)

func GenUserPasswordSalt() (string, error) {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", xerrors.Errorf("Error generating random bytes: %s", err.Error())
	}

	return base64.URLEncoding.EncodeToString(randomBytes), nil
}

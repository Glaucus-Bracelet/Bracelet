package token

import (
	"fmt"
	"github.com/Glaucus/Bracelet/users"
	"github.com/brianvoe/sjwt"
	"os"
)

func ParseAuthentication(jwt string) (users.User, error) {
	// Verify that the secret signature is valid
	secretKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	isSigned := sjwt.Verify(jwt, secretKey)
	if !isSigned {
		return users.User{}, fmt.Errorf("authentication is not properly signed")
	}

	claims, err := sjwt.Parse(jwt)
	if err != nil {
		return users.User{}, fmt.Errorf("authentication is not properly formatted")
	}

	err = claims.Validate()
	if err != nil {
		return users.User{}, fmt.Errorf("authentication is not valid")
	}

	// Get claims
	username, err := claims.GetStr("username")
	if err != nil {
		return users.User{}, fmt.Errorf("authentication username is not valid")
	}

	user, err := users.FindUser(username)
	if err != nil {
		return users.User{}, fmt.Errorf("authentication user is not found")
	}

	// Finally it should work
	return user, nil
}

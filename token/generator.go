package token

import (
	"github.com/Glaucus/Bracelet/users"
	"github.com/brianvoe/sjwt"
	"os"
	"time"
)

func GenerateJWT(user users.User) string {
	// Set Claims
	claims := sjwt.New()
	claims.Set("username", user.Username)
	claims.Set("user_id", user.Id)
	claims.Set("iat", time.Now().Unix())

	// Generate jwt
	secretKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	return claims.Generate(secretKey)
}

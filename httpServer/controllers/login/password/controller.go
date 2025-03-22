package password

import (
	"github.com/Glaucus/Bracelet/users"
	"time"

	"github.com/brianvoe/sjwt"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var request Request
	json.Unmarshal(reqBody, &request)

	if !isValid(request) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid request"})
		return
	}

	user, err := users.Login(request.Username, request.Password)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Error: "Invalid credentials: " + err.Error()})
		return
	}

	// Set Claims
	claims := sjwt.New()
	claims.Set("username", user.Username)
	claims.Set("user_id", user.Id)
	claims.Set("iat", time.Now().Unix())

	// Generate jwt
	secretKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	jwt := claims.Generate(secretKey)

	json.NewEncoder(w).Encode(Response{Token: jwt})
}

func isValid(request Request) bool {
	if request.Username == "" {
		return false
	}
	if request.Password == "" {
		return false
	}
	return true
}

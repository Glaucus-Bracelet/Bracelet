package password

import (
	"encoding/json"
	"github.com/Glaucus/Bracelet/token"
	"github.com/Glaucus/Bracelet/users"
	"io/ioutil"
	"net/http"
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

	user, err := users.Register(request.Username, request.Password)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Error: "Invalid credentials: " + err.Error()})
		return
	}

	// Generate a JWT!
	jwt := token.GenerateJWT(user)
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

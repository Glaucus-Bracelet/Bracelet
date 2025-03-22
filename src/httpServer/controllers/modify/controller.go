package modify

import (
	"encoding/json"
	"github.com/Glaucus/Bracelet/token"
	"github.com/Glaucus/Bracelet/users"
	"io"
	"net/http"
	"strings"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	authentication := r.Header.Get("Authorization")
	if authentication == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Authentication required"})
		return
	}

	jwt := strings.Split(authentication, " ")[1]
	user, err := token.ParseAuthentication(jwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid authentication: " + err.Error()})
		return
	}

	reqBody, _ := io.ReadAll(r.Body)
	var request Request
	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		return
	}

	if !isValid(request) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid request"})
		return
	}

	user, err = users.Modify(user.Id, request.Username, request.Password)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Error: "Invalid credentials: " + err.Error()})
		return
	}

	// Generate a JWT!
	jwt = token.GenerateJWT(user)
	json.NewEncoder(w).Encode(Response{Token: jwt})
}

func isValid(request Request) bool {
	if request.Username != "" {
		return true
	}
	if request.Password != "" {
		return true
	}
	return false
}

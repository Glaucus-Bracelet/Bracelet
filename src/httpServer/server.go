package httpServer

import (
	loginPassword "github.com/Glaucus/Bracelet/httpServer/controllers/login/password"
	accountModify "github.com/Glaucus/Bracelet/httpServer/controllers/modify"
	registerPassword "github.com/Glaucus/Bracelet/httpServer/controllers/register/password"

	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func Start() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.Use(contentTypeApplicationJsonMiddleware)
	apiRouter.HandleFunc("/login/password", loginPassword.Controller)
	apiRouter.HandleFunc("/register/password", registerPassword.Controller)
	apiRouter.HandleFunc("/modify", accountModify.Controller)
	log.Fatal(http.ListenAndServe(":80", apiRouter))
}

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sanjog-pariyar/user-service/controller"
	"github.com/sanjog-pariyar/user-service/utils"
	"github.com/gorilla/handlers"
)

func router() http.Handler {
	r := mux.NewRouter()

	userPath := r.PathPrefix("/api/v1/user").Subrouter()

	userPath.HandleFunc("/signup", controller.Instance().SignUpHandler).Methods("POST")
	userPath.HandleFunc("/login", controller.Instance().LoginHandler).Methods("POST")

	r.HandleFunc("/health", handlePing).Methods("GET")
	r.HandleFunc("/auth/google/login", controller.Instance().GoogleLogin).Methods("GET")
	r.HandleFunc("/auth/google/callback", controller.Instance().GoogleAuthCallback)
	r.HandleFunc("/image-transform", controller.Instance().GetAssetInfo)

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	return cors(r)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sanjog-pariyar/user-service/controller"
	"github.com/sanjog-pariyar/user-service/utils"
)

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/signup", controller.Instance().Signup).Methods("POST")
	r.HandleFunc("/login", controller.Instance().Login).Methods("POST")
	r.HandleFunc("/health", handlePing).Methods("GET")
	r.HandleFunc("/auth/google/login", controller.Instance().GoogleLogin).Methods("GET")
	r.HandleFunc("/auth/google/callback", controller.Instance().GoogleAuthCallback)
	r.HandleFunc("/image-transform", controller.Instance().ImageTransform)

	return r
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
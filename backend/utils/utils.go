package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sanjog-pariyar/user-service/errorhandler"
)

func RespondWithError(w http.ResponseWriter, err error) {
	var userServiceError *errorhandler.UserServiceError

	if errors.As(err, &userServiceError) {
		switch userServiceError.ErrorType {
		case errorhandler.AlreadyExist:
			RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": userServiceError.ClientMessage})
			return
		case errorhandler.Invalid:
			RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": userServiceError.ClientMessage})
			return
		case errorhandler.NotFound:
			RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": userServiceError.ClientMessage})
			return
		case errorhandler.Unknown:
			RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": userServiceError.ClientMessage})
			return
		case errorhandler.Internal:
			RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": userServiceError.ClientMessage})
			return
		}
	}

	RespondWithJSON(w, http.StatusInternalServerError, map[string]string{
		"error": "internal server error",
	})

}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

package util

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Adarsh-Kmt/EndServer/types"
)

type HttpError struct {
	Error  any
	Status int
}

type HttpFunc func(w http.ResponseWriter, r *http.Request) *HttpError

func MakeHttpHandlerFunc(f HttpFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if HttpError := f(w, r); HttpError != nil {

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(HttpError.Status)
			json.NewEncoder(w).Encode(map[string]any{"error": HttpError.Error})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, body any) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func ValidateLoginRequest(rq types.UserLoginRequest) *map[string]string {

	errorMap := make(map[string]string)
	if len(rq.Username) == 0 {
		log.Println("userId cannot be empty.")
		errorMap["userId"] = "userId cannot be empty."

	}

	if len(rq.Password) == 0 {
		log.Println("password cannot be empty.")
		errorMap["password"] = "password cannot be empty."
	}

	if len(errorMap) == 0 {
		return nil
	}
	return &errorMap
}

func ValidateRegisterRequest(rq types.UserRegisterRequest) *map[string]string {

	errorMap := make(map[string]string)
	if len(rq.Password) == 0 {

		errorMap["userId"] = "userId cannot be empty."

	}

	if len(rq.Username) == 0 {
		errorMap["password"] = "password cannot be empty."
	}

	if len(errorMap) == 0 {
		return nil
	}
	return &errorMap
}

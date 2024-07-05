package util

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Error  string
	Status int
}

type HttpFunc func(w http.ResponseWriter, r *http.Request) *HttpError

func MakeHttpHandlerFunc(f HttpFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if HttpError := f(w, r); HttpError != nil {

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(HttpError.Status)
			json.NewEncoder(w).Encode(HttpError.Error)
		}
	}
}

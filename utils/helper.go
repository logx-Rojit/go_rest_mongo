package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func GetAuthorization(r *http.Request) error {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return errors.New("please provide token to get api data")
	}
	tokenString := authorization[len("Bearer "):]
	err := VerifyToken(tokenString, "access")
	if err != nil {
		return err
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func MakeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

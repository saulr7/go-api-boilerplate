package middleware

import (
	"api-boilerplate/auth"
	"net/http"

	"log"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Request %q, method: %q", r.URL.Path, r.Method)
		f(w, r)
	}

}

func Authenticated(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		_, err := auth.ValidateToken(token)

		if err != nil {
			forbidden(w, r)
			return
		}

		log.Printf("Request %q, method: %q", r.URL.Path, r.Method)
		f(w, r)
	}

}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Unauthorize"))
}

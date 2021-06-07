package server

import (
	"net/http"

	"api-boilerplate/handlers"
	mw "api-boilerplate/middleware"
	"api-boilerplate/storage"
)

func RoutePerson(mux *http.ServeMux, storage storage.Storage) {

	handler := handlers.NewPerson(&storage)

	mux.HandleFunc("/v1/persons", mw.Log(handler.GetAll))
	mux.HandleFunc("/v1/persons/create", mw.Authenticated(mw.Log(handler.Create)))
	mux.HandleFunc("/v1/persons/update", mw.Log(handler.Update))
	mux.HandleFunc("/v1/persons/delete", mw.Log(handler.Delete))

}

func RouteLogin(mux *http.ServeMux, storage storage.Storage) {

	h := handlers.NewLogin(&storage)

	mux.HandleFunc("/v1/login", h.Login)
}

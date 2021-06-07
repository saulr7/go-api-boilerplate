package main

import (
	"fmt"
	"log"
	"net/http"

	"api-boilerplate/auth"
	"api-boilerplate/server"
	"api-boilerplate/storage"
)

func main() {

	err := auth.LoadFiles("./certs/app.rsa", "./certs/app.rsa.pub")

	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	server.RoutePerson(mux, &store)
	server.RouteLogin(mux, &store)
	fmt.Println("Server running at port", 8080)

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}

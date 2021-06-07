package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello world")
}

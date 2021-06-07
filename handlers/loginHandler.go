package handlers

import (
	"api-boilerplate/auth"
	"api-boilerplate/models"
	"api-boilerplate/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

type login struct {
	storage storage.Storage
}

func NewLogin(s *storage.Storage) login {
	return login{*s}
}

func (l *login) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not allow", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	data := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := newResponse(Error, "Structure incorrect", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	if !isLoginValid(&data) {
		response := newResponse(Error, "Credentials incorrects", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	token, err := auth.GenerateToken(&data)

	if err != nil {
		fmt.Println(err)
		response := newResponse(Error, "Couldn't sign token", nil)
		responseJson(w, http.StatusInternalServerError, response)
		return
	}

	dataToken := map[string]string{token: token}

	response := newResponse(Message, "Ok", dataToken)
	responseJson(w, http.StatusOK, response)

}

func isLoginValid(data *models.Login) bool {

	return data.Email == "saulr" && data.Password == "123456"

}

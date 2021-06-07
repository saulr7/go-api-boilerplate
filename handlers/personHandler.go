package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"api-boilerplate/models"
	"api-boilerplate/storage"
)

type person struct {
	storage storage.Storage
}

func NewPerson(storage *storage.Storage) person {
	return person{storage: *storage}
}

func (p *person) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not allow", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	data := models.Person{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := newResponse(Error, "Structure no complete", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Create(&data)

	if err != nil {
		response := newResponse(Error, "Something went wrong", nil)
		responseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Created", nil)
	responseJson(w, http.StatusCreated, response)

}

func (p *person) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {

		response := newResponse(Error, "Method not allow", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response := newResponse(Error, "No correct ID", nil)
		responseJson(w, http.StatusInternalServerError, response)
		return
	}

	data := models.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := newResponse(Error, "Structure no complete", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)

	if err != nil {
		response := newResponse(Error, "Something went wrong", nil)
		responseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Created", nil)
	responseJson(w, http.StatusCreated, response)

}

func (p *person) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Method not allow", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response := newResponse(Error, "No correct ID", nil)
		responseJson(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Delete(ID)

	if errors.Is(err, models.ErrIDPersonNoExits) {
		response := newResponse(Error, "Person doesn't exist", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := newResponse(Error, "Something went wrong", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "Deleted", nil)
	responseJson(w, http.StatusOK, response)

}

func (p *person) GetAll(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method not allow", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	people, err := p.storage.GetAll()

	if err != nil {
		response := newResponse(Error, "Something went wrong", nil)
		responseJson(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "Updated", people)
	responseJson(w, http.StatusOK, response)

}

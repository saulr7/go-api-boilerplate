package storage

import "api-boilerplate/models"

type Storage interface {
	Create(person *models.Person) error
	Update(ID int, p *models.Person) error
	Delete(ID int) error
	GetById(ID int) (models.Person, error)
	GetAll() (models.Persons, error)
}

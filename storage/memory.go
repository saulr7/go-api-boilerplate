package storage

import (
	"api-boilerplate/models"
)

type Memory struct {
	currentId int
	Persons   map[int]models.Person
}

func NewMemory() Memory {
	persons := make(map[int]models.Person)

	return Memory{
		currentId: 0,
		Persons:   persons,
	}
}

func (m *Memory) Create(person *models.Person) error {

	if person == nil {
		return models.ErrPersonCannotBeNil
	}

	m.currentId++
	m.Persons[m.currentId] = *person

	return nil
}

func (m *Memory) Update(ID int, p *models.Person) error {

	if p == nil {
		return models.ErrPersonCannotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return models.ErrIDPersonNoExits
	}

	m.Persons[ID] = *p

	return nil

}

func (m *Memory) Delete(ID int) error {

	if _, ok := m.Persons[ID]; !ok {
		return models.ErrIDPersonNoExits
	}

	delete(m.Persons, ID)

	return nil
}

func (m *Memory) GetById(ID int) (models.Person, error) {

	p, ok := m.Persons[ID]

	if !ok {
		return p, models.ErrIDPersonNoExits
	}

	return p, nil
}

func (m *Memory) GetAll() (models.Persons, error) {

	var ps models.Persons

	for _, v := range m.Persons {
		ps = append(ps, v)
	}

	return ps, nil

}

package models

type Person struct {
	Name        string      `json:"name"`
	Age         uint        `json:"age"`
	Communities []Community `json:"communities"`
}

type Persons = []Person

type Community struct {
	Name string `json:"Name"`
}

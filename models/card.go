package models

import "fmt"

type Card interface {
	GetName() string
	GetNumberCard() int
}

type KTP struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Gender             bool   `json:"gender"`
	NumberCivilization int    `json:"nik"`
}

func (k KTP) GetName() string {
	return k.FirstName + "" + k.LastName
}
func (k KTP) GetNumberCard() int {
	return k.NumberCivilization
}

type SIM struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       bool   `json:"gender"`
	NumberDriver int    `json:"sim"`
}

func (s SIM) GetName() string {
	return s.FirstName + " " + s.LastName
}

func (s SIM) GetNumberCard() int {
	return s.NumberDriver
}

func GetCardInformation(c Card) string {

	return fmt.Sprintf(fmt.Sprintf(`{"Name" :%s,"Number" :%d}`, c.GetName(), c.GetNumberCard()))

}

package models

import (
	"github.com/go-playground/validator/v10"
)

type Person struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" validate:"required,alpha,min=2,max=50"`
	Surname     string `json:"surname" validate:"required,alpha,min=2,max=50"`
	Patronymic  string `json:"patronymic,omitempty" validate:"omitempty,alpha,max=50"`
	Age         int    `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty" gorm:"size:10"`
	Nationality string `json:"nationality,omitempty" gorm:"size:50"`
}

func (p *Person) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Nome    string    `json:"nome" validate:"nonzero"`
	CPF     string    `json:"cpf" validate:"len=11"`
	RG      string    `json:"rg" validate:"len=9"`
	Address Address `gorm:"polymorphic:Customer;"`
}

type Address struct {
	gorm.Model
	Street       string `json:"street"`
	Neigborhood  string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipcode"`
	CustomerID   int    `json:"customerId"`
	CustomerType string `json:"customerType"`
}

func ValidateCustomer(customer *Customer) error {
	if err := validator.Validate(customer); err != nil {
		return err
	}
	return nil
}

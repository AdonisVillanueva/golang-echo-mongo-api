package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	FirstName    string             `json:"firstname,omitempty" validate:"required"`
	LastName     string             `json:"lastname,omitempty" validate:"required"`
	UserName     string             `json:"username,omitempty" validate:"required"`
	OrgName      string             `json:"orgname"`
	Location     string             `json:"location,omitempty" validate:"required"`
	Title        string             `json:"title,omitempty" validate:"required"`
	Address1     string             `json:"address1"`
	Address2     string             `json:"address2"`
	City         string             `json:"city"`
	State        string             `json:"State"`
	ZipCode      string             `json:"zip"`
	Country      string             `json:"country"`
	Gender       string             `json:"gender"`
	Account      string             `json:"account"`
	DOB          string             `json:"dob,omitempty" validate:"required"`
	Email        string             `json:"email" validate:"required,email"`
	Password     string             `json:"password"`
	Active       bool               `json:"active"`
	LastAccessed string             `json:"lastaccessed"`
	CreatedDate  string             `json:"createddate"`
}

// type User struct {
// 	Id       primitive.ObjectID `json:"id,omitempty"`
// 	Name     string             `json:"name,omitempty" validate:"required"`
// 	Location string             `json:"location,omitempty" validate:"required"`
// 	Title    string             `json:"title,omitempty" validate:"required"`
// }

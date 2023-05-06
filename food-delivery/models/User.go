package models

import (
	gender "github.com/nikzayn/lld-golang/food-delivery/enums"
)

type User struct {
	Name          string
	Gender        gender.Gender
	ContactNumber string
	Pincode       string
}

func (u *User) createUser() *User {
	user := User{
		Name:          "Nikhil",
		Gender:        "Male",
		ContactNumber: "7042224751",
		Pincode:       "110063",
	}
	return &user
}

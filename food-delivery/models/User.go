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

func (u User) getName() string {
	return u.Name
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u User) getGender() gender.Gender {
	return u.Gender
}

func (u *User) setGender(gender gender.Gender) {
	u.Gender = gender
}

func (u User) getContactNumber() string {
	return u.ContactNumber
}

func (u *User) setContactNumber(contactNumber string) {
	u.ContactNumber = contactNumber
}

func (u User) getPincode() string {
	return u.Pincode
}

func (u *User) setPincode(pincode string) {
	u.Pincode = pincode
}

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

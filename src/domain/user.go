package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

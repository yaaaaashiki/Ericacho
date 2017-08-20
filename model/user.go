package model

import (
	"github.com/jinzhu/gorm"
)

func (u *User) Create(db *gorm.DB) {
	db.Create(u)
}

func (u *User) FetchAllUsers(db *gorm.DB) ([]User, error) {
	users := make([]User, 0, 16)
	db.Find(&users)
	return users, nil
}

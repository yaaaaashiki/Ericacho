//go:generate scaneo $GOFILE

package model

import "time"

// User returns model object for user.
type User struct {
	Id      int64      `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Salt    string     `json:"salt"`
	Salted  string     `json:"salted"`
	Created *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	Updated *time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
}

type Comment struct {
	Id      int64      `json:"id"`
	Message string     `json:"message"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

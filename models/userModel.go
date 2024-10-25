package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null;size:63"`
	Password string `gorm:"not null;size:63"`
	Note     datatypes.JSON
}

// BeforeSave is a GORM hook that could be used to hash the password before saving
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// Here you could add password hashing logic before saving
	return
}

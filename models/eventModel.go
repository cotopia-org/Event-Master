package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	OwnerID  uint `gorm:"not null"`           // Foreign key to User model
	Owner    User `gorm:"foreignKey:OwnerID"` // Establishes relationship to User as Owner
	Epoch    int  `gorm:"not null"`
	Kind     string
	Doer     string
	IsPair   bool `gorm:"default:false"`
	PairId   uint
	Duration int `gorm:"default:-1"`
	Note     datatypes.JSON
}

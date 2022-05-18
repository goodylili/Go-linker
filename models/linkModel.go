package models

import (
	"gorm.io/gorm"
)

type LinkModel struct {
	gorm.Model `json:"-"`
	Identifier string `gorm:"primaryKey;not null" json:"shortened"`
	Link       string `gorm:"unique;not null" json:"link"`
}

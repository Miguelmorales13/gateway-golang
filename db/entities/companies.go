package entities

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);not null"`
	Logo       string `gorm:"type:varchar(60);not null"`
	Background string `gorm:"type:varchar(60);not null"`
	NamePublic string `gorm:"type:varchar(100);not null"`
	Color      string `gorm:"type:varchar(100);not null"`
	Active     bool   `gorm:"default:true;not null"`
	Users      []User
}

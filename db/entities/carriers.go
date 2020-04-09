package entities

import (
	"github.com/jinzhu/gorm"
)

type Carrier struct {
	gorm.Model
	Name   string `gorm:"type:varchar(200);not null"`
	Active bool   `gorm:"default:true"`
}

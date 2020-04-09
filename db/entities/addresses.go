package entities

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Url       string `gorm:"type:varchar(200);not null"`
	Domain    string `gorm:"type:varchar(200);not null"`
	ApiKey    string `gorm:"type:varchar(200);not null"`
	SecretKey string `gorm:"type:varchar(200);not null"`
	Active    bool   `gorm:"default:true;not null"`
}

/**
type Address struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Url      string             	`json:"url,omitempty" bson:"url,omitempty"`
	Domain        string             `json:"domain,omitempty" bson:"domain,omitempty"`
	User        User             		`json:"user,omitempty" bson:"user,omitempty"`
	Active        bool             `json:"active,omitempty" bson:"active,omitempty"`
}
*/

package entities

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"type:varchar(60);not null"`
	Email     string `gorm:"type:varchar(100);not null"`
	User      string `gorm:"type:varchar(100);not null"`
	Active    bool   `gorm:"default:true;not null"`
	CompanyID string `gorm:"type:int;not null;"`
	Apis      []Api
}

/**type User struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BusisnessName string             `json:"busisnessName,omitempty" bson:"busisnessName,omitempty"`
	CompanyName   string             `json:"companyName,omitempty" bson:"companyName,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone         string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	CreateAt      string             `json:"createAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt     string             `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt     string             `json:"deletedAt,omitempty" bson:"deletedAt,omitempty" `
	Active        string             `json:"active,omitempty" bson:"active,omitempty"`
}*/

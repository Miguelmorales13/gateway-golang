package entities

import (
	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	IpClient       string `gorm:"type:char(30);not null;"`
	TimeResponse   int    `gorm:"type:int;not null;"`
	HttpStatusCode string `gorm:"type:varchar(10);not null;"`
	HttpReplySize  int    `gorm:"type:int;not null"`
	HttpMethod     string `gorm:"type:varchar(20);not null;"`
	HttpUrl        string `gorm:"type:text;not null"`
	HttpMimeType   string `gorm:"type:varchar(50);not null;"`
	ProxyStatus    string `gorm:"type:varchar(30);not null;"`
	AddressID      string `gorm:"type:int;not null;"`
	Active         bool   `gorm:"default:true;"`
	CarrierID      string `gorm:"type:int;not null;"`
}

/**
type Api struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ScretKey      string             `json:"secretKey,omitempty" bson:"secretKey,omitempty"`
	ApiKey        string             `json:"apiKey,omitempty" bson:"apiKey,omitempty"`
	Url        string             `json:"url,omitempty" bson:"url,omitempty"`
	Active        bool             `json:"active,omitempty" bson:"active,omitempty"`
}
*/

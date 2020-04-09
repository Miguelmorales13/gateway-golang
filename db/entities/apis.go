package entities

type Api struct {
	ApiKey    string `gorm:"primary_key;index;type:varchar(400);not null"`
	UserID    string `gorm:"type:int;not null"`
	Active    bool   `gorm:"default:true"`
	Addresses []Address
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

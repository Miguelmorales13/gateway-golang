package models

import (
	db "../conexion"
	"../entities"
)

func FindAll() (apis []entities.Address) {
	my := db.DbConn()
	my.Find(&apis)
	defer my.Close()

	return
}
func FindByHost(host string, url string, apikey string) (address entities.Address) {
	my := db.DbConn()
	my.Where("url LIKE ? and api_key  = ? and active  = ?", "%"+host+"%", apikey, true).First(&address)
	defer my.Close()

	return
}

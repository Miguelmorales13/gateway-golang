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
	println(host,apikey)
	my.Where("(url LIKE ? or domain LIKE ?) and api_key  = ? and active  = ? ", "%"+host+"%", "%"+host+"%", apikey, true).First(&address)
	defer my.Close()

	return
}

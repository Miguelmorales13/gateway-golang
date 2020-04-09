package models

import (
	db "../conexion"
	"../entities"
)

func FindAllReports() (apis []entities.Report) {
	my := db.DbConn()
	my.Where(map[string]interface{}{"active": true}).Find(&apis)
	defer my.Close()
	return
}
func Create(log *entities.Report) {
	my := db.DbConn()
	my.Create(&log)

	defer my.Close()

}

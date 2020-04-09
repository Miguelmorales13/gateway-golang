package conexion

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func DbConn() *gorm.DB {
	db, err := gorm.Open(os.Getenv("GORM_DIALECT"), "host="+os.Getenv("GORM_HOST")+" port="+os.Getenv("GORM_PORT")+" user="+os.Getenv("GORM_USER")+" dbname="+os.Getenv("GORM_DATABASE")+" sslmode="+os.Getenv("GORM_SSLMODE")+" password="+os.Getenv("GORM_PASSWORD"))
	if err != nil {
		log.Println(err, "Logger")
		return nil
	} else {
		log.Print("Data base connected")
	}
	// "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	return db
}

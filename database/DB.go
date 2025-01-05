package database

import (
	"log"
	"url_shortner/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConnect() {
	dns := "root:@tcp(127.0.0.1:3306)/url_shortner?charset=utf8mb4"

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectarse a la base de datos: %v", err)
	}
}

func Migrate() error {
	err = DB.AutoMigrate(&models.URL{})
	if err != nil {
		return err
	}
	return nil
}

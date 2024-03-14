package db

import (
	"log"
	"os"
	"strconv"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error

	Host := os.Getenv("DB_HOST")
	Port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	Database := os.Getenv("DB_NAME")
	IntegratedAuth, _ := strconv.ParseBool(os.Getenv("DB_INTEGRATED_AUTH"))

	var dsn string
	if IntegratedAuth {
		dsn = "server=" + Host + ";user id=;password=;port=" + strconv.Itoa(Port) + ";database=" + Database + ";integratedSecurity=true"
	} else {
		// Aquí puedes proporcionar detalles de autenticación si no usas autenticación integrada
	}

	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection successful")
	}
}

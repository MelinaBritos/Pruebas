package baseDeDatos

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func Conexiondb() {
	var err error

	DSN, err := ObtenerDSN()

	if err != nil{
		log.Fatal(err)
		return
	}

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Base de datos conectada")
	}
}

func ObtenerDSN() (string, error) {
	err := godotenv.Load("../TP-Principal-AMAZONA/.env.example")
	if err != nil{
		return "", err
	}
	return os.Getenv("DSN"), nil
}

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/elwicho/APIREST-GO/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Dbinstance representa una instancia de la base de datos.
type Dbinstance struct {
	Db *gorm.DB
}

// DB es una variable global que contiene una instancia de la base de datos.
var DB Dbinstance

// ConnectDb establece la conexión con la base de datos y realiza las migraciones necesarias.
func ConnectDb() {
	// Construye la cadena de conexión utilizando los valores de entorno.
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Mexico_City",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Abre la conexión a la base de datos.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	db.Exec("UPDATE sabores SET complemento = 'valor_predeterminado' WHERE complemento IS NULL")

	log.Println("Running migrations")
	// Aplica las migraciones necesarias en la base de datos para los modelos definidos.
	db.AutoMigrate(&models.Sabores{})

	// Asigna la instancia de la base de datos a la variable global DB.
	DB = Dbinstance{
		Db: db,
	}
}

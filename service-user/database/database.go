package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//CreateConnection
func CreateConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbName),
	)
}

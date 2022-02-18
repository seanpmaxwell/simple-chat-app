package daos

import (
	"fmt"
	"os"
	"simple-chat-app/server/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB = nil
)

/**
https://github.com/go-gorm/postgres
*/
func init() {
	// Setup env vars
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbName := os.Getenv("DATABASE_NAME")
	dbPwd := os.Getenv("DATABASE_PASSWORD")
	// Open connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPwd, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Setup GORM models
	db.AutoMigrate(&models.User{}, &models.UserCreds{})
	// Init connection
	conn = db
}

/**
Get the database connection.
*/
func GetDbConn() *gorm.DB {
	return conn
}

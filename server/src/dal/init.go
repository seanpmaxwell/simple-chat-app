package dal

import (
	"fmt"
	"simple-chat-app/server/src/models"
	"simple-chat-app/server/src/shared"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dnsStr = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
)

var (
	db *gorm.DB = nil
)

/**
https://github.com/go-gorm/postgres
*/
func Init() {
	// Don't setup if already connected
	if db != nil {
		return
	}
	// Setup connection string
	dbParams := shared.GetDbParams()
	dsn := fmt.Sprintf(dnsStr, dbParams.Host, dbParams.User, dbParams.Pwd, dbParams.Name,
		dbParams.Port)
	// Open connection
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Migrate GORM models
	conn.AutoMigrate(&models.User{}, &models.UserCreds{})
	// Init connection
	db = conn
}

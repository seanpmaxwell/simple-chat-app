package main

import (
	"fmt"
	"simple-chat-app/server/src/models"
	"simple-chat-app/server/src/shared"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
https://github.com/go-gorm/postgres
*/
func newDbConn(envVars *shared.EnvVars) *gorm.DB {
	// Setup connection string
	dbParams := envVars.DbParams
	dsn := fmt.Sprintf(dnsStr, dbParams.Host, dbParams.User, dbParams.Pwd, dbParams.Name,
		dbParams.Port)
	// Open connection
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// Migrate GORM models
	conn.AutoMigrate(&models.User{}, &models.UserCreds{})
	// Init connection
	return conn
}

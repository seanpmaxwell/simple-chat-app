package simplechat

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	dnsStr = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
)

var (
	db *gorm.DB = nil
)

// initPostgres gets the Postgres DB connection
// See https://github.com/go-gorm/postgres
func initPostgres() {

	// Don't initialize if already connected
	if db != nil {
		return
	}

	// Setup connection string
	p := GetDBParams()
	dsn := fmt.Sprintf(dnsStr, p.Host, p.User, p.Pwd, p.Name, p.Port)

	// Open connection
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to open DB connection at %s; %+v", p.Host, err)
	}

	// Migrate GORM models
	err = conn.AutoMigrate(&User{}, &UserCreds{})
	if err != nil {
		log.Fatalf("fail to run automatic DB migration; %+v", err)
	}

	// Init connection
	db = conn
}

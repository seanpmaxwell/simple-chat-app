package shared

import (
	"fmt"
	"os"
	"strconv"
)

type DbParams struct {
	Host string
	Port string
	Name string
	User string
	Pwd  string
}

type JwtParams struct {
	Secret []byte
	Exp    int
}

type CookieParams struct {
	Name   string
	Domain string
	Path   string
	Secure bool
	Exp    int
}

type EnvVars struct {
	DbParams     *DbParams
	JwtParams    *JwtParams
	CookieParams *CookieParams
}

// Wire()
func WireEnvVars() *EnvVars {
	var err error
	// Database
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPwd := os.Getenv("DATABASE_PASSWORD")
	dbParams := DbParams{dbHost, dbPort, dbName, dbUser, dbPwd}
	// Cookie
	cookieName := os.Getenv("COOKIE_NAME")
	cookieDomain := os.Getenv("COOKIE_DOMAIN")
	cookiePath := os.Getenv("COOKIE_PATH")
	secureCookie, err := strconv.ParseBool(os.Getenv("SECURE_COOKIE"))
	if err != nil {
		fmt.Println(err.Error())
	}
	cookieExp, err := strconv.Atoi(os.Getenv("COOKIE_EXP"))
	if err != nil {
		fmt.Println(err.Error())
	}
	cookieParams := CookieParams{cookieName, cookieDomain, cookiePath, secureCookie, cookieExp}
	// Json-Web-Token
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	jwtParams := JwtParams{jwtSecret, cookieExp}
	return &EnvVars{&dbParams, &jwtParams, &cookieParams}
}

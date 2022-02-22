package simplechat

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mitchellh/go-homedir"
)

const DefaultPort = 5432

type DBParams struct {
	Host string
	Port int
	Name string
	User string
	Pwd  string
}

type CookieParams struct {
	Name   string
	Domain string
	Path   string
	Secure bool
	Exp    int
}

var (
	//go:embed env/*.env
	envs embed.FS
)

const envSubdir = "env"
const envExt = ".env"

var (
	configPath = fmt.Sprintf(".%s", Identifier)
	envDir     = filepath.Join(mustGetHomeDir(), configPath, envSubdir)
)

var (
	// Database
	dbParams *DBParams

	// Json-Web-Token
	jwtSecret = make([]byte, 0)

	// Cookies
	cookieParams *CookieParams
)

// mustPopulateEnvDir copies the embedded .env files
// and writes them to an
func mustPopulateEnvDir() {
	err := os.MkdirAll(envDir, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to make directory %s; %+v", envDir, err)
	}

	envFiles, _ := envs.ReadDir(envSubdir)

	for _, f := range envFiles {
		name := f.Name()
		content, _ := envs.ReadFile(fmt.Sprintf("%s/%s", envSubdir, name))
		fp := mustGetEnvFilepath(name)
		err = ioutil.WriteFile(fp, content, os.ModePerm)
		if err != nil {
			log.Fatalf("unable to create %s; %+v", fp, err)
		}
	}
}

// GetEnvs returns a list of available environments.
func GetEnvs() (envs []string) {
	files, err := ioutil.ReadDir(envDir)
	if err != nil {
		log.Fatalf("unable to read directory %s; %+v", envDir, err)
	}

	envs = make([]string, 0, 3)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		name := f.Name()
		if filepath.Ext(name) != envExt {
			continue
		}
		envs = append(envs, name[:len(name)-len(envExt)])
	}
	return envs
}

// initEnv loads env file and call functions that require parsing.
func initEnv() {
	loadEnv()

	// Database
	dbParams = &DBParams{
		Host: os.Getenv("DATABASE_HOST"),
		Port: getPort(),
		Name: os.Getenv("DATABASE_NAME"),
		User: os.Getenv("DATABASE_USER"),
		Pwd:  os.Getenv("DATABASE_PASSWORD"),
	}

	// Json-Web-Token
	jwtSecret = getJWTSecret()

	// Cookie
	cookieParams = &CookieParams{
		Name:   os.Getenv("COOKIE_NAME"),
		Domain: os.Getenv("COOKIE_DOMAIN"),
		Path:   os.Getenv("COOKIE_PATH"),
		Secure: getSecureCookie(),
		Exp:    getCookieExp(),
	}
}

// GetDBParams looks in daos/db.go for connection setup.
func GetDBParams() *DBParams {
	return dbParams
}

// GetJWTSecret is used in jwtUtil to sign the token and the session middleware.
func GetJWTSecret() []byte {
	return jwtSecret
}

// GetJWTExp is used in jwtUtil to sign the token. Exp jwt and same time as cookie.
func GetJWTExp() int {
	return cookieParams.Exp
}

// GetCookieParams are used to set the cookie in the auth router
func GetCookieParams() *CookieParams {
	return cookieParams
}

// GetCookieName is used to look up the cookie in the middleware.
func GetCookieName() string {
	return cookieParams.Name
}

func getPort() (port int) {
	var err error
	var i64Port int64

	sPort := os.Getenv("DATABASE_PORT")
	if sPort == "" {
		port = DefaultPort
		goto end
	}

	i64Port, err = strconv.ParseInt(sPort, 10, 0)
	if err != nil {
		log.Printf("Value for environment variable `DATABASE_PORT` [%d] is invalid, using default value [%d] instead", i64Port, DefaultPort)
		port = DefaultPort
		goto end
	}

end:
	return port
}

func getJWTSecret() (secret []byte) {
	secret = []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		fmt.Println("Unable to get JWT secret. Environment variable `JWT_SECRET` not set or has an empty value.")
		os.Exit(ExitNoJWTSecretCookieEnvVar)
	}
	return secret
}

func getSecureCookie() bool {
	cookie := os.Getenv("SECURE_COOKIE")
	exists := cookie != ""
	if !exists {
		fmt.Println("Unable to get secure cookie. Environment variable `SECURE_COOKIE` not set or has an empty value.")
		os.Exit(ExitNoSecureCookieEnvVar)
	}
	secure, err := strconv.ParseBool(cookie)
	if err != nil {
		fmt.Printf("The `SECURE_COOKIE` environment variable whose value is '%s' cannot be parsed as true or false.\n", cookie)
		os.Exit(ExitInvalidSecureCookieEnvVar)
	}
	return secure
}

func getCookieExp() int {
	sExp := os.Getenv("COOKIE_EXP")
	exp, err := strconv.Atoi(sExp)
	if err != nil {
		prefix := "Unable to get cookie expiration. Environment variable `COOKIE_EXP`"
		if sExp == "" {
			fmt.Printf("%s is not set or is empty.", prefix)
			os.Exit(ExitNoCookieExpEnvVar)
		} else {
			fmt.Printf("%s has an invalid value: '%s.'", prefix, sExp)
			os.Exit(ExitInvalidCookieExpEnvVar)
		}
	}
	return exp
}

func mustGetHomeDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Unable to get home directory; %+v", err)
	}
	return dir
}

func entryExists(entryPath string) bool {
	_, err := os.Stat(entryPath)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	if err != nil {
		log.Fatalf("Unable to access %s; %+v", entryPath, err)
	}
	return true
}

// mustGetEnvFilepath returns the filepath for the desired .env file.
func mustGetEnvFilepath(baseFile string) string {
	return filepath.Join(envDir, fmt.Sprintf("%s%s", baseFile, envExt))
}

// loadEnv loads environment variables from ".env" files.
func loadEnv() {
	var err error

	if !hasEnvFiles() {
		mustPopulateEnvDir()
	}

	if len(os.Args) == 1 {
		fmt.Printf("You must specify one of these environments as a parameter: %v",
			GetEnvs())
		os.Exit(ExitNoEnvParam)
	}

	env := os.Args[1]
	fp := mustGetEnvFilepath(env)
	if !entryExists(fp) {
		fmt.Printf("The environment '%s' does not exist as a file at %s", env, fp)
		os.Exit(ExitNoEnvFile)
	}

	err = godotenv.Load(fp)
	if err != nil {
		log.Fatalf("Unable to load %s file; %+v", fp, err)
	}
}

// hasEnvFiles returns true if the .simplechat/env directory has .env files
func hasEnvFiles() (has bool) {
	var files []fs.FileInfo
	var err error

	if !entryExists(envDir) {
		goto end
	}

	files, err = ioutil.ReadDir(envDir)
	if err != nil {
		log.Fatalf("Unable to read directory %s; %+v", envDir, err)
	}
	if len(files) == 0 {
		goto end
	}
	has = true
end:
	return has

}

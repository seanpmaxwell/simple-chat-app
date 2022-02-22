package simplechat

const (
	Identifier = "simplechat"
	Name       = "Simple Chat Server"
	Version    = "0.0.1"
)

type ExitCode = int

const (
	ExitNoEnvParam int = iota
	ExitNoEnvFile
	ExitNoSecureCookieEnvVar
	ExitInvalidSecureCookieEnvVar
	ExitNoCookieExpEnvVar
	ExitInvalidCookieExpEnvVar
	ExitNoJWTSecretCookieEnvVar
)

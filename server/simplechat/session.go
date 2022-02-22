package simplechat

const (
	sessionDataKey = "session-data"
)

func GetSessionDataKey() string {
	return sessionDataKey
}

type Session struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

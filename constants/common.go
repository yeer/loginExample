package constants

import "time"

const (
	DBName    = "user"
	CacheName = "user"
	CACHE_TTL = 24 * time.Hour

	LogLogin    = "login"
	LogRegister = "register"
)

package constants

/**
2000+ : 系统级别错误
3000+ : 业务解决错误
*/
const (
	StatusOk           = 0 // RFC 7231, 6.2.1
	StatusParamsError  = 2001
	StatusResultEmpty  = 2002
	StatusDbNoResult   = 2003
	StatusProgramError = 2500

	StatusCreateFailure = 3001
	StatusLoginFailure  = 3002
)

var statusText = map[int]string{
	StatusOk:            "OK",
	StatusParamsError:   "params error",
	StatusResultEmpty:   "result empty",
	StatusDbNoResult:    "No record in the database",
	StatusProgramError:  "Internal Server Error",
	StatusCreateFailure: "Create Data Error",
	StatusLoginFailure:  "Incorrect username or password",
}

func StatusText(code int) string {
	return statusText[code]
}

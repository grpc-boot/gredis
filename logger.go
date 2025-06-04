package gredis

var (
	errorLog ErrLog
)

type (
	ErrLog func(err error, cmd string, opt *Option)
)

func SetErrorLog(l ErrLog) {
	errorLog = l
}

func WriteLog(err error, cmd string, opt *Option) {
	if errorLog == nil || IsNil(err) {
		return
	}

	errorLog(err, cmd, opt)
}

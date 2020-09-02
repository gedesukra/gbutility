package gbutility

import (
	"fmt"
	"runtime"
)

func LogAllError(logger *CustomLogger, strErr string, err error) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		logger.Println("LogAllError : " + strErr + " -> " + err.Error() + "\nStacktrace:" + errStrFormat)
	} else {
		logger.Println("LogAllError : " + strErr + " -> " + err.Error())
	}
}

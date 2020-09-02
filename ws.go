package gbutility

import (
	"fmt"
	"net/http"
	"runtime"
)

func GetErrorWs(completeUrl string, logger *CustomLogger, err error, strErr string) string {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		logger.Println("Handle Ws Error : " + strErr + " -> " + err.Error() + "\nStacktrace:" + errStrFormat)
	} else {
		logger.Println("Handle Ws Error : " + strErr + " -> " + err.Error())
	}

	return GetJsonStrRs(completeUrl, http.StatusInternalServerError, false, strErr)
}

func GetErrorReqWs(completeUrl string, logger *CustomLogger, err error) string {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		logger.Println("Request Ws Error : " + err.Error() + "\nStacktrace:" + errStrFormat)
	} else {
		logger.Println("Request Ws Error : " + err.Error())
	}

	return GetJsonStrRs(completeUrl, http.StatusInternalServerError, false, "Invalid request: "+AddSlashesDoubleQuote(err.Error()))
}

func GetErrorUnmarshallWs(completeUrl string, logger *CustomLogger, err error) string {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		logger.Println("Request Ws Error : " + err.Error() + "\nStacktrace:" + errStrFormat)
	} else {
		logger.Println("Request Ws Error : " + err.Error())
	}

	return GetJsonStrRs(completeUrl, http.StatusInternalServerError, false, "Invalid request json :"+AddSlashesDoubleQuote(err.Error()))
}

func GetErrorPanicWs(completeUrl string, logger *CustomLogger, err error, strErr string) string {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		logger.Println("Handle Ws Error : " + strErr + " -> " + err.Error() + "\nStacktrace:" + errStrFormat)
	} else {
		logger.Println("Handle Ws Error : " + strErr + " -> " + err.Error())
	}

	return GetJsonStrRs(completeUrl, http.StatusInternalServerError, false, "There is something wrong, please contact administrator")
}

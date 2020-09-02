package gbutility

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
)

func EnableCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*res).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func ReturnPanicToHttp(completeUrl string, res http.ResponseWriter, req *http.Request, logger *CustomLogger) {
	if err := recover(); err != nil {
		if err != nil {
			//	pc, fn, line, _ := runtime.Caller(1)
			// errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
			errStr := string(debug.Stack())
			logger.Println("Panic Error : " + err.(error).Error() + "\nStacktrace:" + errStr)
		} else {
			logger.Println("Panic Error :  Unknown error")
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		defResponse := GetJsonStrRs(completeUrl, http.StatusInternalServerError, false, fmt.Sprintf("There is something wrong, please contact administrator"))
		res.Write([]byte(defResponse))
		return
	}
}

func HandleErrReq(completeUrl string, res http.ResponseWriter, req *http.Request, logger *CustomLogger, err error) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		logger.Println("Request Error : " + err.Error() + "\nStacktrace:" + errStrFormat)
	} else {
		logger.Println("Request Error :  Unknown error")
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)
	errJsonStr := AddSlashesDoubleQuote(err.Error())
	defResponse := GetJsonStrRs(completeUrl, http.StatusInternalServerError, false, fmt.Sprintf("Invalid request : %v", errJsonStr))
	res.Write([]byte(defResponse))
	return
}

func HandleErrUnmarshal(completeUrl string, res http.ResponseWriter, req *http.Request, logger *CustomLogger, err error) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		logger.Println("Unmarshall Error : " + err.Error() + "\nStacktrace:" + errStrFormat)
	} else {
		logger.Println("Unmarshall Error : Unknown error")
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)
	errJsonStr := AddSlashesDoubleQuote(err.Error())
	defResponse := GetJsonStrRs(completeUrl, http.StatusInternalServerError, false, fmt.Sprintf("Invalid request json : %v", errJsonStr))
	res.Write([]byte(defResponse))
	return
}

func HandleError(completeUrl string, res http.ResponseWriter, req *http.Request, logger *CustomLogger, err error, strErr string) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		errStrFormat := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		logger.Println("HandleError : " + strErr + " -> " + err.Error() + "\nStacktrace:" + errStrFormat)
	} else {
		logger.Println("HandleError : " + strErr)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)
	defResponse := GetJsonStrRs(completeUrl, http.StatusInternalServerError, false, fmt.Sprintf("There is something wrong, please contact administrator"))
	res.Write([]byte(defResponse))
	return
}

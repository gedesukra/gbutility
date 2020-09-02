package gbutility

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
)

/*CustomLogger struct hold filename and logger pointer*/
type CustomLogger struct {
	filename string
	*log.Logger
}

var loggerInstance *CustomLogger
var onceLogger sync.Once

var appLog *log.Logger

/*GetInstanceLogger create file log from paramter app name & path */
func GetInstanceLogger(appName string, loggerPath string) *CustomLogger {
	onceLogger.Do(func() {
		// loggerInstance = createLogger(GetLoggerPath("app.log"))
		loggerInstance = createLogger(appName, loggerPath)
	})
	return loggerInstance
}

func createLogger(appName string, fname string) *CustomLogger {
	// file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	file, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		os.Exit(1)
	}

	appLog = log.New(file, appName+" : ", log.Ldate|log.Ltime)
	appLog.SetOutput(&lumberjack.Logger{
		Filename:   fname,
		MaxSize:    10, // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})

	return &CustomLogger{
		filename: fname,
		Logger:   appLog,
	}

}

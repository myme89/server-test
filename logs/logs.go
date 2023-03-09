package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("./logs/log_server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	Logger = log.New()
	Logger.SetOutput(file)
	Logger.SetLevel(log.DebugLevel)
	Logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceQuote:      true,
	})
}

func FunctionWithLogging(namePack, nameFun, status string) {
	Logger.WithFields(log.Fields{
		"package":  namePack,
		"function": nameFun,
		"status":   status,
	}).Info()
}

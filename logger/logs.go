package logger

import (
	"os"
)

var writingLogs = true
var logsDir = getLogDir()

func getLogDir() string {
	result := os.Getenv("LOG_DIR")
	if os.Getenv("LOG_DIR") == "" {
		result = "/var/log/goffer"
	}
	return result
}

func writeLog(log string) {
	if !writingLogs {
		return
	}
	if _, err := os.Stat(logsDir); err != nil {
		err = os.Mkdir(logsDir, 0777)
		if err != nil {
			Error(err)
		}
	}
	logFile, err := os.OpenFile(logsDir+"/"+startTime+".log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err == nil {
		_, err = logFile.WriteString(log)
		if err != nil {
			Error(err)
		}
		err = logFile.Close()
		if err != nil {
			Error(err)
		}
	} else {
		logFile, err := os.Create(logsDir + "/" + startTime + ".log")
		if err != nil {
			Error(err)
		}
		_, err = logFile.WriteString(log)
		if err != nil {
			Error(err)
		}
		err = logFile.Close()
		if err != nil {
			Error(err)
		}
	}
}

func SetLogs(value bool) {
	writingLogs = value
}

func SetLogsDir(path string) {
	logsDir = path
}

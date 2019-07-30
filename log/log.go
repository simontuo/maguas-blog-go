package log

import (
	"io"
	"maguas-blog-go/config"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Log() {
	gin.DisableConsoleColor()

	var fileName string
	if config.LogDaily {
		fileName = CreateDateFileName() + ".log"
	} else {
		fileName = "log.log"
	}
	f, _ := os.Create(fileName)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func CreateDateFileName() string {
	logFilePath := config.LogFilePath
	fileName := time.Now().Format("20060102")

	return logFilePath + "/" + fileName
}

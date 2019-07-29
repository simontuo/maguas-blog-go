package log

import (
	"github.com/gin-gonic/gin"
	"io"
	"maguas-blog-go/config"
	"os"
	"time"
)

func Log()  {
	gin.DisableConsoleColor()

	fileName := CreateDateFileName()
	f, _ := os.Create(fileName)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func CreateDateFileName() string {
	logFilePath := config.LogFilePath
	fileName := time.Now().Format("20060102")

	return logFilePath + "/" +fileName + ".log"
}
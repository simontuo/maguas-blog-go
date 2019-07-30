package config

const (
	// 服务端口
	Port = 8080

	// 数据库参数
	DatabaseHost     = "127.0.0.1"
	DatabasePort     = "3306"
	DatabaseUser     = "root"
	DatabasePassword = "root"
	DatabaseName     = "maguas-blog-go"

	// jwt
	TokenKey = "simontuo"

	// log
	LogDaily    = true
	LogFilePath = "./log/data"
)

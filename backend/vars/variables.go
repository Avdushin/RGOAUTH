package vars

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// app
	VERSION = "1.1"
	_       = godotenv.Load()
	// .env vars
	PORT = os.Getenv("PORT")
	// Получение переменных окружения
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	// DB connection...
	// unux (Linux/Mac OS) connection without pass
	// DBConn = fmt.Sprintf("%s@tcp(%s:%s)/%s", DBUser, DBHost, DBPort, DBName)
	// other type of connection with pass
	DBConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)
	GormDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPass, DBHost, DBPort, DBName)
)

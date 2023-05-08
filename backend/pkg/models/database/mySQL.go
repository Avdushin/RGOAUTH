package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/avdushin/rgoauth/vars"
	_ "github.com/go-sql-driver/mysql"
)

/* work with MySQL DataBase */

// Make querry...
func DBQuerry(querry, comment string) {
	// DBConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)

	// Соедененеие с базой данных
	db, err := sql.Open("mysql", vars.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // закрываем подключение к БД

	// Делаем запрос
	if _, err = db.Exec(fmt.Sprint(querry)); err != nil {
		log.Fatal(err)
	}

	// Выводим лог
	log.Println(comment)
}

// Создаем нужные таблицы
func InitTables() {
	// Create users table
	DBQuerry(`
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	  );`, "[DB]: `users` table is ready...")

	// Create sessions table
	// DBQuerry(`
	// CREATE TABLE IF NOT EXISTS sessions (
	// 	id VARCHAR(350) NOT NULL PRIMARY KEY,
	// 	user_id VARCHAR(350) NOT NULL,
	// 	token VARCHAR(128) DEFAULT NULL,
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// );`, "Таблица sessions создана или уже существует")

	log.Println("[DB]: Таблицы созданы...")
}

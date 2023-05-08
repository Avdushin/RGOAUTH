package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// .env vars
	PORT   = os.Getenv("PORT")
	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	// DB connection...
	// DBConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", DBUser, DBPass, DBName)
	DBConn = "root:MySQL123@tcp(127.0.0.1:3306)/rgoauth"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if PORT == "" {
		PORT = ":8080"
	}
	InitTables()
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding user registration request:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", DBConn)

	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Проверяем, что пользователь с таким же именем не существует
	stmt, err := db.Prepare("SELECT id FROM users WHERE username = ?")
	if err != nil {
		log.Println("Error preparing database query:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(user.Username).Scan(&id)
	if err == nil {
		// Пользователь уже существует
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Регистрируем нового пользователя
	stmt, err = db.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")
	if err != nil {
		log.Println("Error preparing database statement:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		log.Println("Error inserting user into database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Получаем ID только что созданного пользователя
	id, err = result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Отправляем ID созданного пользователя в ответе
	response := map[string]interface{}{"id": id}
	json.NewEncoder(w).Encode(response)

	fmt.Printf("Создан пользователь: %s\n", user.Username)
}

// Login heandler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding user login request:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", DBConn)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Проверяем, что пользователь существует и пароль совпадает
	// SELECT id FROM users WHERE username = ? AND password = ?
	stmt, err := db.Prepare("SELECT username FROM users WHERE email = ? AND password = ?")
	if err != nil {
		log.Println("Error preparing database query:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// var id int64
	err = stmt.QueryRow(user.Email, user.Password).Scan(&user.Username)
	if err != nil {
		log.Fatal(err)
	}

	// Возвращаем ответ с именем пользователя в теле
	resp := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	fmt.Println("Выполнен вход под именем пользователя:", user.Username)

	fmt.Println(resp)
}

// Функция для создания новой сессии
func CreateSession(db *sql.DB, userID int64) (string, error) {
	// Проверяем, существует ли сессия для данного пользователя
	var sessionID string
	err := db.QueryRow("SELECT session_id FROM sessions WHERE user_id = ?", userID).Scan(&sessionID)
	if err == nil {
		// Сессия уже существует, возвращаем ее идентификатор
		return sessionID, nil
	}

	// Генерируем идентификатор сессии
	sessionID = generateSessionID(32)

	// Создаем новую сессию в базе данных
	stmt, err := db.Prepare("INSERT INTO sessions(session_id, user_id) VALUES(?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionID, userID)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

// Функция для генерации случайной строки заданной длины
func generateSessionID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Creating Tables
func InitTables() {
	CreateUsers()
	CreateSessions()
	log.Println("БД подключена...")
}

// Session table
func CreateSessions() {
	// Инициализация соединения с базой данных
	db, err := sql.Open("mysql", DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы sessions, если ее нет
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS sessions (
            id VARCHAR(350) NOT NULL PRIMARY KEY,
            user_id VARCHAR(350) NOT NULL,
            token VARCHAR(128) DEFAULT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Таблица sessions создана или уже существует")
}

// Users table
func CreateUsers() {
	// Инициализация соединения с базой данных
	db, err := sql.Open("mysql", DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы users, если ее нет
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(50) NOT NULL
	  );
    `)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Таблица users создана или уже существует")
}

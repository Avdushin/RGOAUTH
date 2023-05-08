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
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var (
	_ = godotenv.Load()
	// .env vars
	PORT = os.Getenv("PORT")
	// Получение переменных окружения
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	// DB connection...
	DBConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	fmt.Println(DBConn)
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

// SignUp handler
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result, err := stmt.Exec(user.Username, user.Email, string(hashedPassword))

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
	stmt, err := db.Prepare("SELECT username, password FROM users WHERE email = ?")
	if err != nil {
		log.Println("Error preparing database query:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var hashedPassword string
	err = stmt.QueryRow(user.Email).Scan(&user.Username, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusBadRequest)
			log.Printf("Email не найден\nОшибка: %v", err)
			return
		}
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("User Pass: ", user.Password)
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		log.Printf("Неверный пароль\nОшибка: %v", err)
		return
	}

	fmt.Println("User Pass: ", user.Password)

	// var id int64
	fmt.Println("Username: ", user.Username)
	err = stmt.QueryRow(user.Email).Scan(&user.Username, &user.Password)
	// err = stmt.QueryRow(user.Email).Scan(&user.Username)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Username: ", user.Username)

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

	fmt.Println("User Pass: ", user.Password)
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

// DB

func DBQuerry(querry, comment string) {
	// защищаем переменные БД
	// DBConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)

	// Соедененеие с базой данных
	db, err := sql.Open("mysql", DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprint(querry))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(comment)
}

// Creating Tables
func InitTables() {
	// Create users table
	DBQuerry(`
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	  );`, "Таблица users создана или уже существует")

	// Create sessions table
	DBQuerry(`
	CREATE TABLE IF NOT EXISTS sessions (
		id VARCHAR(350) NOT NULL PRIMARY KEY,
		user_id VARCHAR(350) NOT NULL,
		token VARCHAR(128) DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`, "Таблица sessions создана или уже существует")

	log.Println("Таблицы users, session созданы...")
}

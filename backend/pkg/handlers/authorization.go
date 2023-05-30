package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/avdushin/rgoauth/pkg/utils"
	"github.com/avdushin/rgoauth/vars"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        string `json:"age"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

// SignUp handler
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
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

	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)

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

	log.Printf("Создан пользователь: %s\n", user.Username)
}

// Login heandler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding user login request:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Проверяем, что пользователь существует и пароль совпадает
	// SELECT id FROM users WHERE username = ? AND password = ?
	stmt, err := db.Prepare("SELECT id, username, password FROM users WHERE email = ?")
	if err != nil {
		log.Println("Error preparing database query:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var hashedPassword string
	if err = stmt.QueryRow(user.Email).Scan(&user.ID, &user.Username, &hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusBadRequest)
			log.Printf("Email не найден\nОшибка: %v", err)
			return
		}
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized) // StatusBadRequest
		log.Printf("Неверный пароль\nОшибка: %v", err)
		return
	}

	if err = stmt.QueryRow(user.Email).Scan(&user.ID, &user.Username, &user.Password); err != nil {
		log.Panicln(err)
	}

	// Возвращаем ответ с именем пользователя в теле
	resp := struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	log.Println("Выполнен вход под именем пользователя:", user.Username)

	// fmt.Println(resp) // json response
}

// func RegisterHandler(c *gin.Context) {
// 	var user User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		log.Println("Error decoding user registration request:", err)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
// 	if err != nil {
// 		log.Println("Error connecting to database:", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}
// 	defer db.Close()

// 	stmt, err := db.Prepare("SELECT id FROM users WHERE username = ?")
// 	if err != nil {
// 		log.Println("Error preparing database query:", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}
// 	defer stmt.Close()

// 	var id int64
// 	err = stmt.QueryRow(user.Username).Scan(&id)
// 	if err == nil {
// 		// User already exists
// 		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "User already exists"})
// 		return
// 	}

// 	stmt, err = db.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")
// 	if err != nil {
// 		log.Println("Error preparing database statement:", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}
// 	defer stmt.Close()

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Println("Error hashing password:", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	result, err := stmt.Exec(user.Username, user.Email, string(hashedPassword))
// 	if err != nil {
// 		log.Println("Error inserting user into database:", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	id, err = result.LastInsertId()
// 	if err != nil {
// 		log.Println("Error getting last insert ID:", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	response := map[string]interface{}{"id": id}
// 	c.JSON(http.StatusOK, response)

// 	log.Printf("Создан пользователь: %s\n", user.Username)
// }

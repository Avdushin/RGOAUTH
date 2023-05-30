package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	h "github.com/avdushin/rgoauth/pkg/handlers"
	"github.com/avdushin/rgoauth/pkg/models/database"
	"github.com/avdushin/rgoauth/vars"
)

func main() {
	// Установка соединения с базой данных MySQL
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("[DB] SQL adress: ", vars.DBConn+vars.DBName)
	if vars.PORT == "" {
		vars.PORT = ":4000"
	}
	// Init DB tables
	database.InitTables()

	router := mux.NewRouter()

	// Обработчик OPTIONS запросов
	router.Methods(http.MethodOptions).HandlerFunc(handleOptions)

	// Handlers
	router.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}", h.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", h.UpdateUser).Methods(http.MethodPut)

	// Настройте обработку CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	/*
	* Start backup DB
	* Every 4 hours
	 */
	go func() {
		for range time.Tick(4 * time.Hour) {
			database.BackupDB()
			database.SaveStorage()
		}
	}()

	log.Printf("Server run at the %s%s", vars.DBHost, vars.PORT)

	// Start server
	log.Fatal(http.ListenAndServe(vars.PORT, corsHandler(router)))

	// Start SSL server (Production Server)
	// log.Fatal(http.ListenAndServeTLS(vars.PORT, vars.Cert, vars.Key, corsHandler(router)))
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

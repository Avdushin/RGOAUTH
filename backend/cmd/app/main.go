package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/avdushin/rgoauth/pkg/handlers"
	"github.com/avdushin/rgoauth/pkg/models/database"
	"github.com/avdushin/rgoauth/vars"
)

func main() {
	// Logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("[DB] SQL adress: ", vars.DBConn)
	if vars.PORT == "" {
		vars.PORT = ":4000"
	}
	// Init DB tables
	database.InitTables()
	// Handlers
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	/*
	* Start backup DB
	* Every 5 hours
	 */
	go func() {
		for range time.Tick(4 * time.Hour) {
			database.BackupDB()
			database.SaveStorage()
		}
	}()

	// Start server
	log.Fatal(http.ListenAndServe(vars.PORT, nil))

	// Start SSL server
	// log.Fatal(http.ListenAndServeTLS(vars.PORT, "var/certs/cert.pem", "var/certs/key.pem", nil))
}

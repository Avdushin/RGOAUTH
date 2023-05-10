package main

import (
	"log"
	"net/http"

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
		vars.PORT = ":8080"
	}
	// Init DB tables
	database.InitTables()
	// Handlers
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	// Start server
	log.Fatal(http.ListenAndServe(vars.PORT, nil))

	// Start SSL server
	// log.Fatal(http.ListenAndServeTLS(vars.PORT, "var/certs/cert.pem", "var/certs/key.pem", nil))
}

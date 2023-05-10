package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// SQL Injection Vulnerability
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("mysql", "user:password@/dbname")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		username := r.URL.Query().Get("username")

		query := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", username)
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Process and display query results
		// ...
	})

	// Insecure Random Number Generation
	rand.Seed(time.Now().UnixNano())
	token := rand.Int63()
	fmt.Println("Generated token:", token)

	// Plaintext Password Storage
	username := "testuser"
	password := "plaintextpassword"
	register(username, password)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func register(username, password string) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = db.Exec(query, username, password)
	if err != nil {
		log.Fatal(err)
	}
}

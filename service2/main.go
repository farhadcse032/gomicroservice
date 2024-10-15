// main.go
package main

import (
	"fmt"
	"log"
	"microServices/service1/infra/db"
	"net/http"

	"log/slog"

	"github.com/kenshaw/envcfg"
)

func main() {
	// Load the configuration
	config, err := envcfg.New()
	if err != nil {
		slog.Error("Failed to read config: ", err)
	}

	// Initialize the database connection
	db.Initialize()

	// Example of a simple handler that uses the DB
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		conn := db.GetDB()

		// Example query to fetch users
		var users []struct {
			ID   string    `db:"id"`
			Email string `db:"email"`
		}
		err := conn.Select(&users, "SELECT id, email FROM user_info LIMIT 10")
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		for _, user := range users {
			w.Write([]byte(fmt.Sprintf("CMS User: %s\n", user.Email)))
		}
	})

	// Start the server
	port := config.GetInt("server.port")
	log.Printf("Starting server on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

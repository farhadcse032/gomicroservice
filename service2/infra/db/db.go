// db/db.go
package db

import (
	"fmt"
	"log"
	"time"

	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/kenshaw/envcfg"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

// Initialize initializes the DB connection using configuration settings
func Initialize() {
	cfg, err := envcfg.New()
	if err != nil {
		slog.Error("Failed to read config: ", err)
	}

	// Create the DSN (Data Source Name) string for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.GetString("database.host"), cfg.GetInt("database.port"), cfg.GetString("database.user"), cfg.GetString("database.password"), cfg.GetString("database.dbname"))
	fmt.Println("*****************")
	fmt.Println(dsn)
	// Open the connection
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Set connection pool settings (optional but recommended)
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	DB = conn
	log.Println("Database connection established")
}

// GetDB returns the database instance
func GetDB() *sqlx.DB {
	if DB == nil {
		Initialize() // Ensure the connection is initialized
	}
	return DB
}

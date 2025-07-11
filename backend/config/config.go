package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Config holds the application configuration
type Config struct {
	Port   string
	DBPath string
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	return &Config{
		Port:   "8080",
		DBPath: "./vouchers.db",
	}
}

// InitDB initializes the SQLite database and creates the vouchers table
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create the vouchers table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS vouchers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		crew_name TEXT NOT NULL,
		crew_id TEXT NOT NULL,
		flight_number TEXT NOT NULL,
		flight_date TEXT NOT NULL,
		aircraft_type TEXT NOT NULL,
		seat1 TEXT NOT NULL,
		seat2 TEXT NOT NULL,
		seat3 TEXT NOT NULL,
		created_at TEXT NOT NULL
	);
	`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		db.Close()
		return nil, err
	}

	// Create an index on flight_number and flight_date for faster lookups
	createIndexQuery := `
	CREATE INDEX IF NOT EXISTS idx_flight_date ON vouchers(flight_number, flight_date);
	`

	_, err = db.Exec(createIndexQuery)
	if err != nil {
		log.Printf("Warning: Failed to create index: %v", err)
	}

	return db, nil
}

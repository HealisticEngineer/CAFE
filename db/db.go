package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	// Determine the environment (default to "development")
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Load the appropriate .env file
	envFile := fmt.Sprintf(".env.%s", env)
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}

	// Get the connection string from the environment
	connString := os.Getenv("SQLSERVER_CONN")
	if connString == "" {
		log.Fatalf("SQLSERVER_CONN is not set in the environment")
	}

	// Connect to the database
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}
}

// EnsureTables ensures the required tables exist
func EnsureTables() {
	scriptPath := "./sql_tables.sql" // Relative path to the SQL script
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Fatalf("Failed to read SQL script: %v", err)
	}

	_, err = DB.Exec(string(script))
	if err != nil {
		log.Fatalf("Failed to execute SQL script: %v", err)
	}

	log.Println("Database tables ensured.")
}

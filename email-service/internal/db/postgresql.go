package db

import (
	"database/sql"
	"fmt"
	"log"

	"email-service/internal/config"
	_ "github.com/lib/pq" // Import pq driver for side effects to register with database/sql
)

// PostgresClient wraps sql.DB for PostgreSQL
type PostgresClient struct {
	DB *sql.DB
}

// ConnectPostgres establishes a connection to PostgreSQL
func ConnectPostgres(config *config.Config) (*PostgresClient, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.DBUser, config.DBName, config.DBPassword, config.DBHost, config.DBAddr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully")
	return &PostgresClient{DB: db}, nil
}

// GetUserByEmail fetches a user by email from the database
func (pc *PostgresClient) GetUserByEmail(email string) (map[string]interface{}, error) {
	query := "SELECT username, email FROM \"user\" WHERE email = $1"
	row := pc.DB.QueryRow(query, email)

	var username, userEmail string
	if err := row.Scan(&username, &userEmail); err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	user := map[string]interface{}{
		"username": username,
		"email":    userEmail,
	}
	return user, nil
}

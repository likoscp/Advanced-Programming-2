package db

import (
	"database/sql"
	"fmt"
	"log"

	"email-service/internal/config"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	DB *sql.DB
}

func ConnectPostgres(config *config.Config) (*PostgresClient, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.DBUser, config.DBName, config.DBPassword, config.DBHost, config.DBAddr)
	log.Printf("Attempting to connect with connStr: %s", connStr)
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

func (pc *PostgresClient) GetAllUsers() ([]map[string]interface{}, error) {
	query := "SELECT username, email FROM \"user\""
	rows, err := pc.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var username, email string
		if err := rows.Scan(&username, &email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		user := map[string]interface{}{
			"username": username,
			"email":    email,
		}
		users = append(users, user)
	}
	return users, nil
}

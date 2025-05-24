package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/configs"
)

type Store struct {
	db *sql.DB
}

func NewStore(config configs.ConfigDB) (*Store, error) {
	cnn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.User, config.Name, config.Password, config.Host, config.Addr)
	db, err := sql.Open("postgres", cnn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

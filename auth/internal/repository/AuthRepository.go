package repository

import (
	"github.com/likoscp/Advanced-Programming-2/auth/internal/store/postgresql"
	"github.com/likoscp/Advanced-Programming-2/auth/models"
)

type AuthRepository struct {
	store *postgresql.Store
}

func NewAuthRepository(store *postgresql.Store) *AuthRepository {
	return &AuthRepository{
		store: store,
	}
}

func (ar *AuthRepository) Create(u *models.User) (string, error) {
	query := `INSERT INTO "user"(email, password) VALUES ($1, $2) RETURNING id`
	id := ""
	rows := ar.store.GetDB().QueryRow(query, u.Email, u.Password)
	if rows.Err() != nil {
		return "", rows.Err()
	}
	if err := rows.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
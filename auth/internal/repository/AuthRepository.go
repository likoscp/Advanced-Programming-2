package repository

import (
	"database/sql"
	"errors"

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

func (ar *AuthRepository) Get(email string) (models.User, error) {
	query := `SELECT id, password FROM "user" WHERE email = $1`
	u := models.User{Email: email}

	row := ar.store.GetDB().QueryRow(query, u.Email)
	if err := row.Scan(&u.ID, &u.Password); err != nil {
		return u, err
	}
	return u, nil
}
func (ar *AuthRepository) GetAdminId(id string) (bool, error) {
	query := `SELECT id FROM "admin" WHERE id = $1`

	row := ar.store.GetDB().QueryRow(query, id)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return false, nil
	}
	if row.Err() != nil {
		return false, row.Err()
	}
	return true, nil
}

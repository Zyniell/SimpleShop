package repository

import (
	"database/sql"
	"errors"
	"simpleshop/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, created_at`

	result := &model.User{}
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).
		Scan(&result.ID, &result.Name, &result.Email, &result.CreatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE email = $1`
	user := &model.User{}
	err := r.db.QueryRow(query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
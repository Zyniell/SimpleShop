package repository

import (
	"database/sql"
	"errors"
	"simpleshop/internal/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	query := `SELECT id, name, created_at FROM categories ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) FindByID(id int) (*model.Category, error) {
	query := `SELECT id, name, created_at FROM categories WHERE id = $1`
	c := &model.Category{}
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CategoryRepository) Create(name string) (*model.Category, error) {
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id, name, created_at`
	c := &model.Category{}
	err := r.db.QueryRow(query, name).Scan(&c.ID, &c.Name, &c.CreatedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CategoryRepository) Update(id int, name string) (*model.Category, error) {
	query := `UPDATE categories SET name = $1 WHERE id = $2 RETURNING id, name, created_at`
	c := &model.Category{}
	err := r.db.QueryRow(query, name, id).Scan(&c.ID, &c.Name, &c.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CategoryRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
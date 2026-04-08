package repository

import (
	"database/sql"
	"errors"
	"simpleshop/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindAll() ([]model.Product, error) {
	query := `
		SELECT p.id, p.name, p.description, p.price, p.stock, p.category_id, p.created_at,
		       c.id, c.name, c.created_at
		FROM products p
		JOIN categories c ON p.category_id = c.id
		ORDER BY p.id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		var cat model.Category
		if err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.CreatedAt,
			&cat.ID, &cat.Name, &cat.CreatedAt,
		); err != nil {
			return nil, err
		}
		p.Category = &cat
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id int) (*model.Product, error) {
	query := `
		SELECT p.id, p.name, p.description, p.price, p.stock, p.category_id, p.created_at,
		       c.id, c.name, c.created_at
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1`

	var p model.Product
	var cat model.Category
	err := r.db.QueryRow(query, id).Scan(
		&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.CreatedAt,
		&cat.ID, &cat.Name, &cat.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	p.Category = &cat
	return &p, nil
}

func (r *ProductRepository) Create(req *model.ProductRequest) (*model.Product, error) {
	query := `
		INSERT INTO products (name, description, price, stock, category_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	var id int
	err := r.db.QueryRow(query, req.Name, req.Description, req.Price, req.Stock, req.CategoryID).
		Scan(&id)
	if err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *ProductRepository) Update(id int, req *model.ProductRequest) (*model.Product, error) {
	query := `
		UPDATE products
		SET name=$1, description=$2, price=$3, stock=$4, category_id=$5
		WHERE id=$6`

	res, err := r.db.Exec(query, req.Name, req.Description, req.Price, req.Stock, req.CategoryID, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, nil
	}
	return r.FindByID(id)
}

func (r *ProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
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

func (r *ProductRepository) DeductStock(id, quantity int) error {
	query := `UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1`
	res, err := r.db.Exec(query, quantity, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("insufficient stock")
	}
	return nil
}
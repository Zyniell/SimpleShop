package repository

import (
	"database/sql"
	"simpleshop/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(userID, productID, quantity int, totalPrice float64) (*model.Order, error) {
	query := `
		INSERT INTO orders (user_id, product_id, quantity, total_price, status)
		VALUES ($1, $2, $3, $4, 'pending')
		RETURNING id, user_id, product_id, quantity, total_price, status, created_at`

	o := &model.Order{}
	err := r.db.QueryRow(query, userID, productID, quantity, totalPrice).
		Scan(&o.ID, &o.UserID, &o.ProductID, &o.Quantity, &o.TotalPrice, &o.Status, &o.CreatedAt)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (r *OrderRepository) FindByUserID(userID int) ([]model.Order, error) {
	query := `
		SELECT o.id, o.user_id, o.product_id, o.quantity, o.total_price, o.status, o.created_at,
		       p.id, p.name, p.description, p.price, p.stock, p.category_id, p.created_at
		FROM orders o
		JOIN products p ON o.product_id = p.id
		WHERE o.user_id = $1
		ORDER BY o.id DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		var p model.Product
		if err := rows.Scan(
			&o.ID, &o.UserID, &o.ProductID, &o.Quantity, &o.TotalPrice, &o.Status, &o.CreatedAt,
			&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		o.Product = &p
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrderRepository) FindAll() ([]model.Order, error) {
	query := `
		SELECT o.id, o.user_id, o.product_id, o.quantity, o.total_price, o.status, o.created_at,
		       p.id, p.name, p.description, p.price, p.stock, p.category_id, p.created_at
		FROM orders o
		JOIN products p ON o.product_id = p.id
		ORDER BY o.id DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		var p model.Product
		if err := rows.Scan(
			&o.ID, &o.UserID, &o.ProductID, &o.Quantity, &o.TotalPrice, &o.Status, &o.CreatedAt,
			&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		o.Product = &p
		orders = append(orders, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// UpdateStatus updates the status of a specific order.
func (r *OrderRepository) UpdateStatus(orderID int, status string) error {
	query := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(query, status, orderID)
	return err
}

// FindByID returns a single order by its id, joining product info.
func (r *OrderRepository) FindByID(orderID int) (*model.Order, error) {
	query := `
		SELECT o.id, o.user_id, o.product_id, o.quantity, o.total_price, o.status, o.created_at,
		       p.id, p.name, p.description, p.price, p.stock, p.category_id, p.created_at
		FROM orders o
		JOIN products p ON o.product_id = p.id
		WHERE o.id = $1`

	var o model.Order
	var p model.Product
	err := r.db.QueryRow(query, orderID).Scan(
		&o.ID, &o.UserID, &o.ProductID, &o.Quantity, &o.TotalPrice, &o.Status, &o.CreatedAt,
		&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	o.Product = &p
	return &o, nil
}
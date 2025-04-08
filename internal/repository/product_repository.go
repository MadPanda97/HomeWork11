package repository

import (
	"context"
	"database/sql"
	"internet-store/internal/entity"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *entity.Product) (int64, error)
	GetProductByID(ctx context.Context, id int64) (*entity.Product, error)
}

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepo{db}
}

func (r *ProductRepo) GetProductByID(ctx context.Context, id int64) (*entity.Product, error) {
	var product entity.Product
	err := r.db.QueryRowContext(ctx,
		`SELECT id, name, price, weight, description, sku, amount
		 FROM products WHERE id = $1`, id).
		Scan(&product.ID, &product.Name, &product.Price, &product.Weight,
			&product.Description, &product.SKU, &product.Amount)

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepo) CreateProduct(ctx context.Context, product *entity.Product) (int64, error) {
	var id int64
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO products (name, price, weight, description, sku, amount) 
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		product.Name, product.Price, product.Weight, product.Description,
		product.SKU, product.Amount).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

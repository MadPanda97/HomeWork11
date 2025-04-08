package service

import (
	"context"
	"fmt"
	"internet-store/internal/entity"
	"internet-store/internal/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, user *entity.CreateProductRequest) error
	GetProductByID(ctx context.Context, id int64) (*entity.Product, error)
}
type productService struct {
	r repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *productService {
	return &productService{r}
}

func (s *productService) GetProductByID(ctx context.Context, id int64) (*entity.Product, error) {
	product, err := s.r.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", entity.ErrorProductNotFound, err)
	}
	return product, nil
}

func (s *productService) CreateProduct(ctx context.Context, user *entity.CreateProductRequest) error {

}

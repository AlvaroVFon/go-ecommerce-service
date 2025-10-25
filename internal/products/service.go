package products

import (
	"context"
	"log"

	"ecommerce-service/internal/config"
)

type Repository interface {
	Create(ctx context.Context, data CreateProductRequest) error
	FindByID(ctx context.Context, id int) (*Product, error)
	FindAll(ctx context.Context, limit, page int) ([]Product, error)
	Update(ctx context.Context, id int, data UpdateProductRequest) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
}

type ProductService struct {
	productRepo Repository
	config      *config.Config
}

func NewProductService(productRepo Repository, c *config.Config) *ProductService {
	return &ProductService{productRepo: productRepo, config: c}
}

func (ps *ProductService) Create(ctx context.Context, p *CreateProductRequest) error {
	err := ps.productRepo.Create(ctx, *p)
	if err != nil {
		log.Printf("error creating product: %v", err)
		return err
	}
	return nil
}

func (ps *ProductService) FindByID(ctx context.Context, id int) (*Product, error) {
	return ps.productRepo.FindByID(ctx, id)
}

func (ps *ProductService) FindAll(ctx context.Context, limit, page int) ([]Product, error) {
	offset := (page - 1) * limit
	return ps.productRepo.FindAll(ctx, limit, offset)
}

func (ps *ProductService) Update(ctx context.Context, id int, p UpdateProductRequest) error {
	return ps.productRepo.Update(ctx, id, p)
}

func (ps *ProductService) Delete(ctx context.Context, id int) error {
	return ps.productRepo.Delete(ctx, id)
}

func (ps *ProductService) Count(ctx context.Context) (int, error) {
	return ps.productRepo.Count(ctx)
}

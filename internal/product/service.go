package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, data CreateProductRequest) error
	FindByID(ctx context.Context, id int) (*Product, error)
	FindAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, id int, data UpdateProductRequest) error
	Delete(ctx context.Context, id int) error
}

type ProductService struct {
	productRepo Repository
}

func NewProductService(productRepo Repository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (ps *ProductService) Create(ctx context.Context, product *CreateProductRequest) error {
	err := ps.productRepo.Create(ctx, *product)
	if err != nil {
		log.Printf("error creating product: %v", err)
		return err
	}
	return nil
}

func (ps *ProductService) FindByID(ctx context.Context, id int) (*Product, error) {
	return ps.productRepo.FindByID(ctx, id)
}

func (ps *ProductService) FindAll(ctx context.Context) ([]Product, error) {
	return ps.productRepo.FindAll(ctx)
}

func (ps *ProductService) Update(ctx context.Context, id int, product *UpdateProductRequest) error {
	return ps.productRepo.Update(ctx, id, *product)
}

func (ps *ProductService) Delete(ctx context.Context, id int) error {
	return ps.productRepo.Delete(ctx, id)
}

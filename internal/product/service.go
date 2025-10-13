package product

import "context"

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
		return err
	}
	return nil
}

func (ps *ProductService) FindById(ctx context.Context, id int) (*Product, error) {
	product, err := ps.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) FindAll(ctx context.Context) ([]Product, error) {
	products, err := ps.productRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) Update(ctx context.Context, id int, product *UpdateProductRequest) error {
	err := ps.productRepo.Update(ctx, id, *product)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) Delete(ctx context.Context, id int) error {
	err := ps.productRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

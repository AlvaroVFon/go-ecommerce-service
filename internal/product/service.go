package product

import "context"

type ProductService struct {
	productRepo *ProductRepository
}

func NewProductService(productRepo *ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (ps *ProductService) Create(ctx context.Context, product *CreateProductDTO) error {
	err := ps.productRepo.Create(ctx, *product)
	if err != nil {
		return err
	}
	return nil
}

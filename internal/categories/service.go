package categories

import (
	"context"
)

type (
	Repository interface {
		FindAll(ctx context.Context, limit, offset int) ([]Category, error)
		FindByID(ctx context.Context, id int) (*Category, error)
		Count(ctx context.Context) (int, error)
	}
	CategoryService struct {
		categoryRepo Repository
	}
)

func NewCategoryService(categoryRepo Repository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) FindAll(ctx context.Context, page, limit int) ([]Category, error) {
	offset := (page - 1) * limit
	return s.categoryRepo.FindAll(ctx, limit, offset)
}

func (s *CategoryService) FindByID(ctx context.Context, id int) (*Category, error) {
	return s.categoryRepo.FindByID(ctx, id)
}

func (s *CategoryService) Count(ctx context.Context) (int, error) {
	return s.categoryRepo.Count(ctx)
}

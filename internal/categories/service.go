package categories

import "context"

type (
	Repository interface {
		FindAll(ctx context.Context) ([]Category, error)
		FindByID(ctx context.Context, id int) (*Category, error)
		FindByName(ctx context.Context, name string) (*Category, error)
	}
	CategoryService struct {
		categoryRepo Repository
	}
)

func NewCategoryService(categoryRepo Repository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) FindAll(ctx context.Context) ([]Category, error) {
	return s.categoryRepo.FindAll(ctx)
}

func (s *CategoryService) FindByID(ctx context.Context, id int) (*Category, error) {
	return s.categoryRepo.FindByID(ctx, id)
}

func (s *CategoryService) FindByName(ctx context.Context, name string) (*Category, error) {
	return s.categoryRepo.FindByName(ctx, name)
}

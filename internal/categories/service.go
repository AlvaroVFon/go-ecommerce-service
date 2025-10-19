package categories

import "context"

type (
	Repository interface {
		FindAll(ctx context.Context) ([]Category, error)
		FindByID(ctx context.Context, id int) (*Category, error)
		FindByName(ctx context.Context, name string) (*Category, error)
	}
	CategoryServoce struct {
		categoryRepo Repository
	}
)

func NewCategoryService(categoryRepo Repository) *CategoryServoce {
	return &CategoryServoce{categoryRepo: categoryRepo}
}

func (s *CategoryServoce) FindAll(ctx context.Context) ([]Category, error) {
	return s.categoryRepo.FindAll(ctx)
}

func (s *CategoryServoce) FindByID(ctx context.Context, id int) (*Category, error) {
	return s.categoryRepo.FindByID(ctx, id)
}

func (s *CategoryServoce) FindByName(ctx context.Context, name string) (*Category, error) {
	return s.categoryRepo.FindByName(ctx, name)
}

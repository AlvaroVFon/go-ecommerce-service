package roles

import "context"

type (
	Repository interface {
		FindByID(ctx context.Context, id int) (*Role, error)
		FindByName(ctx context.Context, name string) (*Role, error)
		FindAll(ctx context.Context) ([]Role, error)
	}

	RoleService struct {
		roleRepo Repository
	}
)

func NewRoleService(roleRepo Repository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

func (s *RoleService) FindByID(ctx context.Context, id int) (*Role, error) {
	return s.roleRepo.FindByID(ctx, id)
}

func (s *RoleService) FindByName(ctx context.Context, name string) (*Role, error) {
	return s.roleRepo.FindByName(ctx, name)
}

func (s *RoleService) FindAll(ctx context.Context) ([]Role, error) {
	return s.roleRepo.FindAll(ctx)
}

package roles

type (
	Repository interface {
		FindyByID(id int) (*Role, error)
		FindyByName(name string) (*Role, error)
		FindAll() ([]Role, error)
	}

	RoleService struct {
		roleRepo Repository
	}
)

func NewRoleService(roleRepo Repository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

func (s *RoleService) FinfByID(id int) (*Role, error) {
	return s.roleRepo.FindyByID(id)
}

func (s *RoleService) FindByName(name string) (*Role, error) {
	return s.roleRepo.FindyByName(name)
}

func (s *RoleService) FindAll() ([]Role, error) {
	return s.roleRepo.FindAll()
}

package users

import (
	"context"
	"ecommerce-service/internal/utils"
)

type Repository interface {
	Create(ctx context.Context, u *CreateUserRequest) error
	FindByID(ctx context.Context, id int) (*PublicUser, error)
	FindAll(ctx context.Context) ([]PublicUser, error)
	Update(ctx context.Context, id int, u UpdateUserRequest) error
}

type UserService struct {
	userRepo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{userRepo: repo}
}

func (us *UserService) Create(ctx context.Context, u *CreateUserRequest) error {
	hashPassword, err := utils.HashPassword(u.Password, 10)
	if err != nil {
		return err
	}
	u.Password = hashPassword

	return us.userRepo.Create(ctx, u)
}

func (us *UserService) FindByID(ctx context.Context, id int) (*PublicUser, error) {
	return us.userRepo.FindByID(ctx, id)
}

func (us *UserService) FindAll(ctx context.Context) ([]PublicUser, error) {
	return us.userRepo.FindAll(ctx)
}

func (us *UserService) Update(ctx context.Context, id int, u UpdateUserRequest) error {
	if u.Password != nil {
		hashPassword, err := utils.HashPassword(*u.Password, 10)
		if err != nil {
			return err
		}
		u.Password = &hashPassword
	}

	err := us.userRepo.Update(ctx, id, u)
	if err != nil {
		return err
	}

	return nil
}

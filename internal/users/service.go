package users

import (
	"context"

	"ecommerce-service/pkg/cryptox"
)

type Repository interface {
	Create(ctx context.Context, u *CreateUserRequest) error
	FindByID(ctx context.Context, id int) (*User, error)
	FindAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, id int, u UpdateUserRequest) error
	Delete(ctx context.Context, id int) error
}

type UserService struct {
	userRepo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{userRepo: repo}
}

func (us *UserService) Create(ctx context.Context, u *CreateUserRequest) error {
	hashPassword, err := cryptox.HashPassword(u.Password, 10)
	if err != nil {
		return err
	}
	u.Password = hashPassword

	return us.userRepo.Create(ctx, u)
}

func (us *UserService) FindByID(ctx context.Context, id int) (*User, error) {
	return us.userRepo.FindByID(ctx, id)
}

func (us *UserService) FindAll(ctx context.Context) ([]User, error) {
	return us.userRepo.FindAll(ctx)
}

func (us *UserService) Update(ctx context.Context, id int, u UpdateUserRequest) error {
	if u.Password != nil {
		hashPassword, err := cryptox.HashPassword(*u.Password, 10)
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

func (us *UserService) Delete(ctx context.Context, id int) error {
	return us.userRepo.Delete(ctx, id)
}

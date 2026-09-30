package usecase

import (
	"context"
	"fmt"

	"github.com/DevLucasHenrique/mymock-backend/internal/entity"
	"github.com/DevLucasHenrique/mymock-backend/internal/repository"
)

func sendEmail() {
	fmt.Printf("Get users route used")
}

type UserUsecase struct {
	UserRepo *repository.UserRepository
}

func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepo: userRepo,
	}
}

func (us *UserUsecase) ListAllUsers(ctx context.Context) ([]entity.UserResponse, error) {
	go sendEmail()
	users, err := us.UserRepo.GetAllUsers(ctx)
	if err!=nil {
		return []entity.UserResponse{}, err
	}

	return users, nil
}
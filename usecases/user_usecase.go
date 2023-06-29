package usecases

import (
	"github.com/saracha-06422/go-tickets/domain"
	"github.com/saracha-06422/go-tickets/entity"
)

type userUseCase struct {
	userRepo domain.UserMemoryRepository
}

// จาก Repository มันกลายเป็น UseCase
func NewUserUseCase(repo domain.UserMemoryRepository) domain.UserUseCase {
	//must implements
	return &userUseCase{userRepo: repo}
}

func (u *userUseCase) GetAllUser() (users []entity.User, err error) {
	//ตัวใหม่
	// var user []entity.User
	userList, handleErr := u.userRepo.GetAllUser()
	return userList, handleErr
}

func (u *userUseCase) CreateUser(user *entity.User) (err error) {
	//ชั้น useCase ไป เรียก repo ต่อ
	handleErr := u.userRepo.CreateUser(user)
	return handleErr
}

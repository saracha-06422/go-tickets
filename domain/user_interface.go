package domain

import (
	"github.com/saracha-06422/go-tickets/entity"
)

type UserUseCase interface {
	GetAllUser() (u []entity.User, err error)
	CreateUser(u *entity.User) (err error)
}

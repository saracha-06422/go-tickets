package repositories

import (
	"github.com/saracha-06422/go-tickets/domain"
	"github.com/saracha-06422/go-tickets/entity"
)

// คือ struct ที่เราทำหาร implements method ด้านล้าง GetAllUser, CreateUser
type userMemoryRepository struct {
	//Memory
	userList *[]entity.User
}

// จะ return ตัว UserMemoryRepository ซึ่งมันเป็น interface ดังนั้นจะต้อง implements method ทั้งหมดใน UserMemoryRepository ก่อน

// ตัว domain.UserMemoryRepository กับ ตัว &userMemoryRepository{userList} เท่ากันได้ เพราะมีการ implements method GetAllUser, CreateUser
func NewUserMemoryRepository(userList *[]entity.User) domain.UserMemoryRepository {
	return &userMemoryRepository{userList}
}

func (um *userMemoryRepository) GetAllUser() (userList []entity.User, err error) {
	return *um.userList, err
}

func (um *userMemoryRepository) CreateUser(user *entity.User) (err error) {
	*um.userList = append(*um.userList, *user)
	return err
}

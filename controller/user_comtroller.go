package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saracha-06422/go-tickets/domain"
	"github.com/saracha-06422/go-tickets/entity"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(usecase domain.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// GetAll
func (u *UserHandler) GetAllUser(c *gin.Context) {
	resp, err := u.userUseCase.GetAllUser()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// CreateUser
func (u *UserHandler) CreateUser(c *gin.Context) {
	var newUser entity.User
	err := c.ShouldBind(&newUser)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	//จาก controller ไป usecase ต่อ
	errCreateUser := u.userUseCase.CreateUser(&newUser)
	if errCreateUser != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newUser)
	}
}

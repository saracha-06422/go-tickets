package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saracha-06422/go-tickets/controller"
	"github.com/saracha-06422/go-tickets/entity"
	"github.com/saracha-06422/go-tickets/repositories"
	"github.com/saracha-06422/go-tickets/usecases"
)

func SetupRouter() *gin.Engine {
	//valiable from memory
	userList := []entity.User{}

	userRepo := repositories.NewUserMemoryRepository(&userList)

	userUseCase := usecases.NewUserUseCase(userRepo)

	userController := controller.NewUserHandler(userUseCase)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("user", userController.GetAllUser)
		v1.POST("createUser", userController.CreateUser)
	}
	return r
}

package routes

import (
	"go_crud/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserControllerRoute(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (r *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/users")

	router.GET("/", r.userController.FindUsers)
	router.GET("/:userId", r.userController.FindUserById)
	router.POST("/", r.userController.CreateUser)
	router.PATCH("/:userId", r.userController.UpdateUser)
	router.DELETE("/:userId", r.userController.DeleteUser)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oGalhardo/EstudandoGo/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser/", controller.CreateUser)
	r.POST("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deletedUser/:userId", controller.DeleteUser)

}

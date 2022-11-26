package main

import (
	"Golang-API/config"
	"Golang-API/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	r := gin.Default()
	authRouter := r.Group("api/auth")
	{
		authRouter.POST("/Login", authController.Login)
		authRouter.POST("/Register", authController.Register)
	}
	r.Run()
}

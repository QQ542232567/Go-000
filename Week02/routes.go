package main

import (
	"week002/controller"
	"week002/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.RecoveryMiddleware())

	userRoutes := r.Group("/user")
	userController := controller.UserController{}
	userRoutes.GET(":id", userController.Show)

	return r
}

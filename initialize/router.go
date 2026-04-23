package initialize

import (
	"github.com/gin-gonic/gin"
)

func registerRouters(r *gin.RouterGroup, appManagement *AppManagement) {
	// Initialize User module dependencies

	// User routes
	{
		usersGroup := r.Group("/users")
		usersGroup.GET("", appManagement.UserController.GetUsers)
		usersGroup.POST("", appManagement.UserController.CreateUser)
	}

	// Products routes (placeholder)
	{
		productsGroup := r.Group("/products")
		productsGroup.GET("")
	}

}

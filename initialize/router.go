package initialize

import (
	"ca_ddd_solid_dip_di/initialize/dependency_injection"

	"github.com/gin-gonic/gin"
)

func registerRouters(r *gin.RouterGroup) {
	// Initialize User module dependencies
	userDI := dependency_injection.InitUserDependencies()

	// User routes
	{
		usersGroup := r.Group("/users")
		usersGroup.GET("", userDI.UserController.GetUsers)
		usersGroup.POST("", userDI.UserController.CreateUser)
	}

	// Products routes (placeholder)
	{
		productsGroup := r.Group("/products")
		productsGroup.GET("")
	}

}

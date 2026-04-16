package controller

import (
	"net/http"

	"ca_ddd_solid_dip_di/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUserUseCase *usecase.CreateUser
}

func NewUserController(createUserUseCase *usecase.CreateUser) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
	}
}

// GetUsers handles GET /users
func (uc *UserController) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all users",
	})
}

// CreateUser handles POST /users
func (uc *UserController) CreateUser(c *gin.Context) {
	uc.createUserUseCase.Execute()
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

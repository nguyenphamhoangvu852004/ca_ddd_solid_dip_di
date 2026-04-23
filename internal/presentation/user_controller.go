package presentation

import (
	usecase "ca_ddd_solid_dip_di/internal/application/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUserUseCase *usecase.CreateUserUseCase
	getUsersUseCase   *usecase.GetListUserUseCase
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

func NewUserController(createUserUseCase *usecase.CreateUserUseCase, getUsersUseCase *usecase.GetListUserUseCase) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
		getUsersUseCase:   getUsersUseCase,
	}
}

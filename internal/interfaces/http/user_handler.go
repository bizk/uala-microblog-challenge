package http

import (
	"net/http"
	"uala-microblog-challenge/internal/domain"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUser CreateUser
	getUser    GetUser
}

func NewUserHandler(c CreateUser, g GetUser) *UserHandler {
	return &UserHandler{
		createUser: c,
		getUser:    g,
	}
}

// ---- Create User ----

// CreateUser godoc
// @Summary Crear un nuevo usuario
// @Description Crear un nuevo usuario
// @Tags users
// @Accept json
// @Produce json
// @Success 201 {string} string "ok"
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	userID, err := h.createUser.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created", "user_id": userID})
}

type CreateUser interface {
	Create() (string, error)
}

// ---- Get User ----

// GetUser godoc
// @Summary Obtener un usuario
// @Description Obtener un usuario
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 201 {string} string "ok"
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /users/{user_id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	user, err := h.getUser.Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

type GetUser interface {
	Get(userID string) (*domain.User, error)
}

type GetUserRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

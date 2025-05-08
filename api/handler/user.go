package handler

import (
	"net/http"
	"todo/api/utils"
	"todo/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Description Create a new user with name, email, and password
// @Accept json
// @Produce json
// @Param user body models.CreateUserReq true "Create User"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	var req models.CreateUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	user, err := h.Storage.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// @Summary Login user
// @Description User logs in with their email and password
// @Accept json
// @Produce json
// @Param login body models.LoginReq true "Login credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var req models.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
		return
	}

	user, err := h.Storage.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid email or password"})
		return
	}

	// ✅ Token yaratish
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		User:  *user,
		Token: token,
	})
}

package v1

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imraushankr/gozen/src/configs"
	"github.com/imraushankr/gozen/src/internal/domain"
	"github.com/imraushankr/gozen/src/internal/pkg/logger"
	"github.com/imraushankr/gozen/src/internal/usecase"
	"go.uber.org/zap"
)

type userHandler struct {
	userUsercase usecase.UserUsecase
	cfg          *configs.JWTConfig
}

func NewUserHandler(userUsecase *usecase.UserUsecase, cfg *configs.JWTConfig) *userHandler {
	return &userHandler{
		userUsercase: *userUsecase,
		cfg:          cfg,
	}
}

type signupRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,min=10,max=15"`
	Password  string `json:"password" binding:"required,min=8"`
	Role      string `json:"role"`
}

type signinRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type userResponse struct {
	ID          string      `json:"id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	Role        domain.Role `json:"role"`
	Avatar      string      `json:"avatar,omitempty"`
	IsActive    bool        `json:"is_active"`
	IsVerified  bool        `json:"is_verified"`
	LastLoginAt *time.Time  `json:"last_login_at,omitempty"`
	CreatedAt   *time.Time  `json:"created_at,omitempty"`
}

type authResponse struct {
	User         userResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token,omitempty"`
}

func (h *userHandler) Signup(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Format phone number to E.164 if it doesn't start with +
	phone := req.Phone
	if phone[0] != '+' {
		phone = "+" + phone // Add + prefix if missing
	}

	user := &domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Phone:     phone,
		Password:  req.Password,
		Role:      domain.Role(req.Role),
	}

	if err := h.userUsercase.Signup(c.Request.Context(), user); err != nil {
		logger.Error("Failed to sign up user", zap.Error(err))
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp := authResponse{
		User: h.userToResponse(user),
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *userHandler) Signin(c *gin.Context) {
	var req signinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, accessToken, err := h.userUsercase.Signin(c.Request.Context(), req.Identifier, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid credientials")))
		return
	}

	resp := authResponse{
		User:         h.userToResponse(user),
		AccessToken:  accessToken,
		RefreshToken: user.RefreshToken,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) Signout(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
		return
	}

	if err := h.userUsercase.Signout(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully signed out"})
}

func (h *userHandler) RefreshToken(c *gin.Context) {
	var req refreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accessToken, err := h.userUsercase.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (h *userHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
		return
	}

	user, err := h.userUsercase.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, h.userToResponse(user))
}

func (h *userHandler) userToResponse(user *domain.User) userResponse {

	return userResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		Email:       user.Email,
		Phone:       user.Phone,
		Role:        user.Role,
		Avatar:      user.Avatar,
		IsActive:    user.IsActive,
		IsVerified:  user.IsVerified,
		LastLoginAt: user.LastLoginAt,
		CreatedAt:   &user.CreatedAt,
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

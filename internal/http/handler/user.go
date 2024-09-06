package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "todolist/internal/app/docs"
	entityuser "todolist/internal/entity/user"
	userusecase "todolist/internal/usecase/user"
)

type UserServer struct {
	user userusecase.UserUseCaseImpl
}

func NewUserServer(user userusecase.UserUseCaseImpl) *UserServer {
	return &UserServer{
		user: user,
	}
}

// Register godoc
// @Summary Register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body entityuser.CreateUserReq true "User registration information"
// @Success 201 {object} entityuser.StatusMessage
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /user/register [post]
func (u *UserServer) Register(c *gin.Context) {
	var req entityuser.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := u.user.RegisterUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// VerifyUser godoc
// @Summary VerifyUser
// @Description Confirm the code sent to the email
// @Tags auth
// @Accept json
// @Produce json
// @Param body body entityuser.VerifyUserReq true "User verification information"
// @Success 200 {object} entityuser.User
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/verify [post]
func (u *UserServer) VerifyUser(c *gin.Context) {
	var req entityuser.VerifyUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := u.user.VerifyUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// Login godoc
// @Summary Login
// @Description Authenticate a user and return a token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body entityuser.LoginReq true "User login information"
// @Success 200 {object} entityuser.LoginReq
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/login [post]
func (u *UserServer) Login(c *gin.Context) {
	var req entityuser.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := u.user.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}

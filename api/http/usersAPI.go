package api

import (
	"net/http"

	"bitbucket.com/AntonyIS/vision/app"
	"github.com/gin-gonic/gin"
)

type UsersHandler interface {
	CreateUser(*gin.Context)
	User(*gin.Context)
	Users(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userHandler struct {
	userService app.UserServiceInterface
}

func NewHandler(userService app.UserServiceInterface) UsersHandler {
	return &userHandler{
		userService,
	}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user app.UserModel

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	err := h.userService.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfuly",
	})
}

func (h *userHandler) User(c *gin.Context) {
	userID := c.Param("id")
	user, err := h.userService.User(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"todo": user,
	})
}

func (h *userHandler) Users(c *gin.Context) {

	users, err := h.userService.Users()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"todo": users,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var user app.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInternalServerError,
		})

		return
	}
	err := h.userService.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"todo": "User updated successfuly",
	})

}

func (h *userHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if err := h.userService.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrorInternalServerError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo deleted successfuly",
	})

}

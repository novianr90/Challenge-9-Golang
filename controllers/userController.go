package controllers

import (
	"challenge-9/helpers"
	"challenge-9/models"
	"challenge-9/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJson = "application/json"

type UserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserController struct {
	UserService *services.UserService
}

func (uc *UserController) Register(c *gin.Context) {
	var userDto UserDto

	contentType := helpers.GetContentType(c)

	if contentType == appJson {
		if err := c.ShouldBindJSON(&userDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	} else {

		if err := c.ShouldBind(&userDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

	}

	userPassword := helpers.HashPassword(userDto.Password)

	user := models.User{
		Email:    userDto.Email,
		Password: userPassword,
		Username: userDto.Username,
	}

	result, err := uc.UserService.CreateUser(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    result.ID,
		"email": result.Email,
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var userDto UserDto

	if contentType := helpers.GetContentType(c); contentType == appJson {
		if err := c.ShouldBindJSON(&userDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&userDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}

	email := userDto.Email

	user, err := uc.UserService.GetUserByEmail(email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You need to sign up first",
		})
		return
	}
	ok := helpers.CompareHashAndPass([]byte(user.Password), []byte(userDto.Password))

	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (uc *UserController) GetAllUser(c *gin.Context) {

	data := c.MustGet("data").(map[string]any)

	isAdmin := data["isAdmin"].(bool)

	if isAdmin {
		users, err := uc.UserService.GetAllUser()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credential",
		})
	}

}

func (uc *UserController) EditUser(c *gin.Context) {
	var (
		userDto UserDto

		data = c.MustGet("data").(map[string]any)

		userData = data["user"].(models.User)

		email = userData.Email

		contentType = helpers.GetContentType(c)

		err error
	)

	if contentType == appJson {

		if err = c.ShouldBindJSON(&userDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	} else {
		if err = c.ShouldBind(&userDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
	}

	editedUser := models.User{
		Email:    userDto.Email,
		Username: userDto.Username,
		Password: userDto.Password,
	}

	err = uc.UserService.UpdateUserByEmail(email, editedUser)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User sucessfully updated",
	})

}

func (uc *UserController) DeleteUserByEmail(c *gin.Context) {
	var (
		data     = c.MustGet("data").(map[string]any)
		userData = data["user"].(models.User)
	)

	if err := uc.UserService.DeleteUserByEmail(userData.Email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User sucesfully deleted",
	})
}

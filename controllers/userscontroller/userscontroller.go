package userscontroller

import (
	"net/http"

	"example.com/go-rest-api-gin/app"
	"example.com/go-rest-api-gin/database"
	"example.com/go-rest-api-gin/helpers"
	"example.com/go-rest-api-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var user models.Users
	var userLoginData app.UserLoginData

	if err := c.BindJSON(&userLoginData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	// check email from database
	if database.DB.Where("email = ?", userLoginData.Email).First(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Email not found"})

		return
	}

	// check password from database by email found before
	if database.DB.Where("password = ? ", userLoginData.Password).Preload("Photos").First(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Wrong password!"})

		return
	}

	token := helpers.GenerateToken(user)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetUsers(c *gin.Context) {
	var users []models.Users

	database.DB.Preload("Photos").Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserById(c *gin.Context) {
	var users models.Users

	id := c.Param("id")

	if err := database.DB.First(&users, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Id not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": users})
}

func AddUser(c *gin.Context) {
	var user models.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	var user models.Users
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if database.DB.Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Update failed!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update successfully!"})
}

func DeleteUser(c *gin.Context) {
	var user models.Users
	id := c.Param("id")

	if err := database.DB.Delete(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Id not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	if database.DB.Where("id = ?", id).Delete(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Delete user failed!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete user successfully!"})
}

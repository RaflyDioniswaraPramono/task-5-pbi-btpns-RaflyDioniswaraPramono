package photoscontroller

import (
	"net/http"

	"example.com/go-rest-api-gin/database"
	"example.com/go-rest-api-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPhotos(c *gin.Context) {
	var photos []models.Photos

	database.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

func GetPhotoById(c *gin.Context) {
	var photo models.Photos
	id := c.Param("id")

	if err := database.DB.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Photo id not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": photo})
}

func AddPhoto(c *gin.Context) {
	var photo models.Photos

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	database.DB.Create(&photo)
	c.JSON(http.StatusOK, gin.H{"photo": photo})
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photos
	id := c.Param("id")

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if database.DB.Where("id = ?", id).Updates(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Update photo failed!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update photo successfully!"})
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photos
	id := c.Param("id")

	if err := database.DB.Delete(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Photo Id not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	if database.DB.Where("id = ?", id).Delete(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Delete photo failed!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete photo successfully!"})
}

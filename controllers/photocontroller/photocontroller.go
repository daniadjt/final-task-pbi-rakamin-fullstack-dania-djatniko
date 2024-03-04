package photocontroller

import (
	"net/http"
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	models.ConnectDatabase()
}


func GetAllPhoto(c *gin.Context) {

	var photo []models.Photo

	models.DB.Find(&photo)
	c.JSON(http.StatusOK, gin.H{"photo": photo})

}

func GetPhotoById(c *gin.Context) {
	var photo models.Photo
	id := c.Param("id")

	if err := models.DB.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Photo not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"photo": photo})
}

func Create(c *gin.Context) {

	var photo models.Photo
	user_id, _ := c.Get("user_id")

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, ok := user_id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	photo.UserID = userID

	if err := models.DB.Create(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create Success!", "photo": photo})
}

func Update(c *gin.Context) {

	var photo models.Photo
	id := c.Param("id")
	user_id, _ := c.Get("user_id")

	if err := models.DB.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	userID, ok := user_id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	if userID != photo.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Success!", "photo": photo})
}


func Delete(c *gin.Context) {

	var photo models.Photo
	id := c.Param("id")
	user_id, _ := c.Get("user_id")

	if err := models.DB.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Photo not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	userID, ok := user_id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	if userID != photo.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := models.DB.Delete(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success delete the photo"})

}

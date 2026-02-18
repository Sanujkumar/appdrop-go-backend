package handlers

import (
	"net/http"

	"appdrop/config"
	"appdrop/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBrandInput struct {
	Name string `json:"name"`
}

func CreateBrand(c *gin.Context) {
	var input CreateBrandInput

	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand name required"})
		return
	}

	brand := models.Brand{
		ID:   uuid.New(),
		Name: input.Name,
	}

	config.DB.Create(&brand)
	c.JSON(http.StatusCreated, brand)
}

func GetBrands(c *gin.Context) {
	var brands []models.Brand
	config.DB.Find(&brands)
	c.JSON(http.StatusOK, brands)
}

func DeleteBrand(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Brand{}, "id = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Delete failed"})
		return
	}

	c.JSON(200, gin.H{"message": "Brand deleted"})
}
   
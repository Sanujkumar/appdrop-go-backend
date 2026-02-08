package handlers

import (
	"net/http"

	"appdrop/config"
	"appdrop/models"
	"appdrop/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


// ---------- POST /pages ----------
type CreatePageInput struct {
	Name   string `json:"name"`
	Route  string `json:"route"`
	IsHome bool   `json:"is_home"`
}

func CreatePage(c *gin.Context) {
	var input CreatePageInput
	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" || input.Route == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Name and route are required")
		return
	}

	// unique route check
	var existing models.Page
	if err := config.DB.Where("route = ?", input.Route).First(&existing).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "VALIDATION_ERROR", "Page route already exists")
		return
	}

	// only one home page
	if input.IsHome {
		config.DB.Model(&models.Page{}).Where("is_home = ?", true).Update("is_home", false)
	}

	page := models.Page{
		ID:     uuid.New(),
		Name:   input.Name,
		Route:  input.Route,
		IsHome: input.IsHome,
	}

	if err := config.DB.Create(&page).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "DB_ERROR", "Failed to create page")
		return
	}

	c.JSON(http.StatusCreated, page)
}

// ---------- GET /pages ----------
func ListPages(c *gin.Context) {
	var pages []models.Page
	config.DB.Find(&pages)
	c.JSON(http.StatusOK, pages)
}

// ---------- GET /pages/:id ----------
func GetPage(c *gin.Context) {
	id := c.Param("id")
	var page models.Page

	if err := config.DB.Preload("Widgets", func(db *gorm.DB) *gorm.DB {
		return db.Order("position asc")
	}).First(&page, "id = ?", id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "NOT_FOUND", "Page not found")
		return
	}

	c.JSON(http.StatusOK, page)
}

// ---------- PUT /pages/:id ----------
func UpdatePage(c *gin.Context) {
	id := c.Param("id")
	var page models.Page

	if err := config.DB.First(&page, "id = ?", id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "NOT_FOUND", "Page not found")
		return
	}

	var input CreatePageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid input")
		return
	}

	// route unique check
	if input.Route != "" && input.Route != page.Route {
		var existing models.Page
		if err := config.DB.Where("route = ?", input.Route).First(&existing).Error; err == nil {
			utils.ErrorResponse(c, http.StatusConflict, "VALIDATION_ERROR", "Page route already exists")
			return
		}
		page.Route = input.Route
	}

	if input.Name != "" {
		page.Name = input.Name
	}

	if input.IsHome {
		config.DB.Model(&models.Page{}).Where("is_home = ?", true).Update("is_home", false)
		page.IsHome = true
	}

	config.DB.Save(&page)
	c.JSON(http.StatusOK, page)
}

// ---------- DELETE /pages/:id ----------
func DeletePage(c *gin.Context) {
	id := c.Param("id")
	var page models.Page

	if err := config.DB.First(&page, "id = ?", id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "NOT_FOUND", "Page not found")
		return
	}

	if page.IsHome {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Cannot delete home page")
		return
	}

	config.DB.Delete(&page)
	c.Status(http.StatusOK)
}

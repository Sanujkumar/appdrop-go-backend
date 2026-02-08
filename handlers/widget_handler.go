package handlers

import (
	"net/http"

	"appdrop/config"
	"appdrop/models"
	"appdrop/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

var allowedTypes = map[string]bool{
	"banner": true,
	"product_grid": true,
	"text": true,
	"image": true,
	"spacer": true,
}

// ---------- POST /pages/:id/widgets ----------
type CreateWidgetInput struct {  
	Type     string          `json:"type"`
	Position int             `json:"position"`
	Config   datatypes.JSON  `json:"config"`
}

func AddWidget(c *gin.Context) {
	pageID := c.Param("id")

	var input CreateWidgetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid input")
		return
	}

	if !allowedTypes[input.Type] {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid widget type")
		return
	}

	w := models.Widget{
		ID:       uuid.New(),
		PageID:   uuid.MustParse(pageID),  
		Type:     input.Type,
		Position: input.Position,
		Config:   input.Config,
	}

	if err := config.DB.Create(&w).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "DB_ERROR", "Failed to create widget")
		return
	}

	c.JSON(http.StatusCreated, w)
}


// ---------- PUT /widgets/:id ----------
func UpdateWidget(c *gin.Context) {
	id := c.Param("id")
	var w models.Widget

	if err := config.DB.First(&w, "id = ?", id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "NOT_FOUND", "Widget not found")
		return
	}

	var input CreateWidgetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid input")
		return
	}

	if input.Type != "" && !allowedTypes[input.Type] {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid widget type")
		return
	}

	if input.Type != "" {
		w.Type = input.Type
	}
	w.Position = input.Position
	if len(input.Config) > 0 {
		w.Config = input.Config
	}

	config.DB.Save(&w)
	c.JSON(http.StatusOK, w)
}

// ---------- DELETE /widgets/:id ----------
func DeleteWidget(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Widget{}, "id = ?", id)
	c.Status(http.StatusOK)
}

// ---------- POST /pages/:id/widgets/reorder ----------


func ReorderWidgets(c *gin.Context) {
	pageID := c.Param("id")  

	var widgetIDs []string
	if err := c.ShouldBindJSON(&widgetIDs); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid input")
		return
	}

	if len(widgetIDs) == 0 {
		utils.ErrorResponse(c, http.StatusBadRequest, "VALIDATION_ERROR", "Widget list required")
		return
	}

	for i, wid := range widgetIDs {
		if err := config.DB.Model(&models.Widget{}).
			Where("id = ? AND page_id = ?", wid, pageID).
			Update("position", i+1).Error; err != nil {

			utils.ErrorResponse(c, http.StatusInternalServerError, "DB_ERROR", "Failed to reorder widgets")
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Widgets reordered successfully",
	})
}

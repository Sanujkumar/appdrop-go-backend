package routes

import (
	"appdrop/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Pages
	r.GET("/pages", handlers.GetPages)
	r.GET("/pages/:id", handlers.GetPageByID)
	r.POST("/pages", handlers.CreatePage)
	r.PUT("/pages/:id", handlers.UpdatePage)
	r.DELETE("/pages/:id", handlers.DeletePage)

	// Widgets
	r.POST("/pages/:id/widgets", handlers.AddWidget)  
	r.PUT("/widgets/:id", handlers.UpdateWidget)
	r.DELETE("/widgets/:id", handlers.DeleteWidget)
	r.POST("/pages/:id/widgets/reorder", handlers.ReorderWidgets)
}

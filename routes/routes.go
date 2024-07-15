package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	fxRateController := controllers.NewFxRateController(db)

	router.GET("/fx-rates/:date", fxRateController.ReadFxRatesByDate)

	return router
}

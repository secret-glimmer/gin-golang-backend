package controllers

import (
	"app/services"
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FxRateController struct {
	DB *gorm.DB
}

func NewFxRateController(db *gorm.DB) *FxRateController {
	return &FxRateController{
		DB: db,
	}
}

func (controller *FxRateController) ReadFxRatesByDate(context *gin.Context) {
	date, err := utils.ValidateDate(context.Param("date"))
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, "Invalid date format: "+err.Error())
		return
	}

	fxRates := [][]any{}

	pricingService := services.NewPricingService(controller.DB)
	err = pricingService.ReadByDate(&fxRates, date)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, "Failed to process data: "+err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, fxRates)
}

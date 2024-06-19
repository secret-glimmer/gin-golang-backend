package services

import (
	"app/models"
	"app/repositories"
	"slices"
	"time"

	"gorm.io/gorm"
)

type PricingService struct {
	DB *gorm.DB
}

func NewPricingService(db *gorm.DB) *PricingService {
	return &PricingService{
		DB: db,
	}
}

func TransformPricingsToFxRates(fxRates *[][]any, pricings *[]models.Pricing) {
	fxRateMap := map[int]*map[string]float64{}
	transferAmounts := []int{}
	organizationNames := []string{}
	organizationNameFlags := map[string]bool{}
	for _, pricing := range *pricings {
		transferAmount := pricing.TransferAmount
		organizationName := pricing.OrganizationName
		organizationMap, ok := fxRateMap[transferAmount]
		if !ok {
			transferAmounts = append(transferAmounts, transferAmount)
			organizationMap = &map[string]float64{}
			fxRateMap[transferAmount] = organizationMap
		}
		if !organizationNameFlags[organizationName] {
			organizationNameFlags[organizationName] = true
			organizationNames = append(organizationNames, organizationName)
		}
		(*organizationMap)[organizationName] = pricing.FxRate
	}
	(*fxRates) = append((*fxRates), []any{"Transfer Amount"})
	slices.Sort(organizationNames)
	for _, organizationName := range organizationNames {
		(*fxRates)[0] = append((*fxRates)[0], organizationName)
	}
	for _, transferAmount := range transferAmounts {
		fxRateLine := []any{transferAmount}
		for _, organizationName := range organizationNames {
			fxRate, ok := (*fxRateMap[transferAmount])[organizationName]
			if !ok {
				fxRateLine = append(fxRateLine, nil)
				continue
			}
			fxRateLine = append(fxRateLine, fxRate)
		}
		(*fxRates) = append((*fxRates), fxRateLine)
	}
}

func (service *PricingService) ReadByDate(fxRates *[][]any, date time.Time) error {
	pricings := []models.Pricing{}
	pricingRepository := repositories.NewPricingRepository(service.DB)
	err := pricingRepository.ReadByDate(&pricings, date)
	if err != nil {
		return err
	}

	TransformPricingsToFxRates(fxRates, &pricings)

	return nil
}

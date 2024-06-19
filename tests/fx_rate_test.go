package tests

import (
	"app/models"
	"app/services"
	"app/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	samplePricings = []models.Pricing{
		{Id: 311, OrganizationName: "GlobalSettle", TransferAmount: 100, FxRate: 4.21, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Id: 26, OrganizationName: "RapidRemit", TransferAmount: 100, FxRate: 4.94, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Id: 757, OrganizationName: "CoinGate", TransferAmount: 300, FxRate: 3.48, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Id: 312, OrganizationName: "SecureShift", TransferAmount: 300, FxRate: 0.65, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Id: 231, OrganizationName: "TransactFX", TransferAmount: 300, FxRate: 4.47, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Id: 286, OrganizationName: "GlobalSettle", TransferAmount: 400, FxRate: 4.2, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Id: 567, OrganizationName: "SecureShift", TransferAmount: 400, FxRate: 0.71, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Id: 869, OrganizationName: "SecureShift", TransferAmount: 400, FxRate: 1.25, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Id: 970, OrganizationName: "GlobalSettle", TransferAmount: 500, FxRate: 1.3, Date: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)},
	}
	sampleFxRates = [][]any{
		{"Transfer Amount", "GlobalSettle", "RapidRemit", "CoinGate", "SecureShift", "TransactFX"},
		{100, 4.21, 4.94, nil, nil, nil},
		{300, nil, nil, 3.48, 0.65, 4.47},
		{400, 4.2, nil, nil, 0.98, nil},
		{500, 1.3, nil, nil, nil, nil},
	}
)

func TestValidateDate(t *testing.T) {
	_, err := utils.ValidateDate("2024-01-13")
	assert.Equal(t, nil, err)
}

func TestTransformPricingsToFxRates(t *testing.T) {
	fxRates := [][]any{}

	services.TransformPricingsToFxRates(&fxRates, &samplePricings)
	assert.Equal(t, sampleFxRates, fxRates)
}

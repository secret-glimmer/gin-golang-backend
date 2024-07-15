package repositories

import (
	"app/models"
	"time"

	"gorm.io/gorm"
)

type PricingRepository struct {
	DB *gorm.DB
}

func NewPricingRepository(db *gorm.DB) *PricingRepository {
	return &PricingRepository{
		DB: db,
	}
}

func (repository *PricingRepository) ReadByDate(pricings *[]models.Pricing, date time.Time) error {
	return repository.DB.
		Where("date = ?", date).
		Order("transfer_amount ASC").
		Find(pricings).
		Error
}

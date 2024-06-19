package models

import "time"

type Pricing struct {
	Id               int       `gorm:"column:id; type:int"`
	Date             time.Time `gorm:"column:date; type:datetime"`
	OrganizationName string    `gorm:"column:organization_name; type:varchar(255)"`
	TransferAmount   int       `gorm:"column:transfer_amount; type:int"`
	FxRate           float64   `gorm:"column:fx_rate; type:float"`
}

func (Pricing) TableName() string {
	return "pricing"
}

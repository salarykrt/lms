package masters

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MasterFoir struct {
	gorm.Model
	CompanyID            uint8  `gorm:"NOT NULL"`
	ProductID            uint8  `gorm:"NOT NULL"`
	CityCategory         string `gorm:"NOT NULL"`
	MonthlySalary        string `gorm:"NOT NULL"`
	MonthlySalary2       string
	FoirPercentageLow    decimal.Decimal `gorm:";type:decimal(10,2);NOT NULL;comment:FOIR%(No official mail id and rented/Relative house)"` // FOIR%(No official mail id and rented/Relative house)
	FoirPercentageMedium decimal.Decimal `gorm:"type:decimal(10,2);NOT NULL;comment:FOIR % ( Official mail ID or Owned house)"`             // FOIR % ( Official mail ID or Owned house)
	FoirPercentageHigh   decimal.Decimal `gorm:"type:decimal(10,2);NOT NULL;comment:FOIR % (official Mail id+ Owned House)"`                // FOIR % (official Mail id+ Owned House)
	CreatedBy            uint32          `gorm:"NOT NULL"`
	UpdatedBy            uint32
}

func (m *MasterFoir) TableName() string {
	return "master_foir"
}

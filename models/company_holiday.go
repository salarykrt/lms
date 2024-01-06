package models

import (
	"time"

	"gorm.io/gorm"
)

type CompanyHoliday struct {
	gorm.Model
	ChHolidayDate time.Time `gorm:"NOT NULL"`
	ChHolidayName string    `gorm:"NOT NULL"`
	ChCreatedBy   uint32
	ChDeletedBy   uint32
	ChActive      uint8 `gorm:"default:1;NOT NULL"`
	ChDeleted     uint8 `gorm:"default:0;NOT NULL"`
}

func (m *CompanyHoliday) TableName() string {
	return "company_holiday"
}

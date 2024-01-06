package models

import (
	"time"

	"gorm.io/gorm"
)

type LeadsOtpTran struct {
	gorm.Model
	LeadID            uint64
	LotUserID         uint32
	LotMobileNo       uint64    `gorm:"NOT NULL"`
	LotMobileOtp      string    `gorm:"NOT NULL"`
	LotMobileOtpType  uint32    `gorm:"comment:1=>QDE Form"` // 1=>QDE Form
	LotOtpTriggerTime time.Time `gorm:"NOT NULL"`
	LotOtpVerifyFlag  uint8     `gorm:"default:0;NOT NULL"`
	LotOtpVerifyTime  time.Time
	LotOtpValidTime   int8
	LotActive         uint8 `gorm:"default:1;NOT NULL"`
	LotDeleted        uint8 `gorm:"default:0;NOT NULL"`
}

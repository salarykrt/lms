package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerVerification struct {
	gorm.Model
	LeadID                         int64 `gorm:"index;NOT NULL"`
	PancardVerifiedStatus          uint8
	PancardVerifiedOn              time.Time
	EmailVerifiedStatus            uint8
	EmailVerifiedOn                time.Time
	AlternateEmailVerifiedStatus   uint8
	AlternateEmailVerifiedOn       time.Time
	PancardOcrVerifiedStatus       uint8 `gorm:"comment:1=>Verified"` // 1=>Verified
	PancardOcrVerifiedOn           time.Time
	AadhaarOcrVerifiedStatus       uint8
	AadhaarOcrVerifiedOn           time.Time
	CustomerEkycRequestIP          string
	CustomerEkycRequestInitiatedOn time.Time
	CustomerDigitalEkycFlag        uint8
	CustomerDigitalEkycDoneOn      time.Time
	CustomerBreRunFlag             uint8
	CustomerBreRunDatetime         time.Time
	Active                         uint8 `gorm:"default:1;NOT NULL"`
	Deleted                        uint8 `gorm:"default:0;NOT NULL"`
}

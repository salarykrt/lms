package models

import (
	"time"

	"gorm.io/gorm"
)

type LeadFollowup struct {
	gorm.Model
	LeadID               uint64 `gorm:"NOT NULL"`
	UserID               uint32
	Remarks              string
	Reason               string
	ScheduledDate        time.Time
	Stage                string `gorm:"NOT NULL;comment:master_stage"`
	LeadFollowupStatusID uint16 `gorm:"NOT NULL;comment:master_status"`
	LeadFollowupActive   uint8  `gorm:"default:1;NOT NULL"`
	LeadFollowupDeleted  uint8  `gorm:"default:0;NOT NULL"`
}

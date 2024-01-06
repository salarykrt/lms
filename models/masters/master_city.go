package masters

import "gorm.io/gorm"

type MasterCity struct {
	gorm.Model
	MCityName          string `gorm:"NOT NULL"`
	MCityCode          string
	MCityCategory      string
	MCityStateID       uint16 `gorm:"NOT NULL"`
	MCityBranchID      uint16 `gorm:"default:0;NOT NULL"`
	MCityIsSourcing    uint8  `gorm:"default:0"`
	MCityTrialSourcing uint8  `gorm:"default:0;NOT NULL;comment:Used for new opening city"` // Used for new opening city
	MCityActive        uint8  `gorm:"default:1;NOT NULL"`
	MCityDeleted       uint8  `gorm:"default:0;NOT NULL"`
}

func (m *MasterCity) TableName() string {
	return "master_city"
}

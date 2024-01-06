package masters

import (
	"gorm.io/gorm"
)

type MasterRelationType struct {
	gorm.Model
	MrtName    string
	MrtActive  uint8 `gorm:"default:1;NOT NULL"`
	MrtDeleted uint8 `gorm:"default:0;NOT NULL"`
}

func (m *MasterRelationType) TableName() string {
	return "master_relation_types"
}

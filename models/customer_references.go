package models

import (
	"gorm.io/gorm"
)

type CustomerReferences struct {
	gorm.Model
	LeadID          uint64 `gorm:"NOT NULL"`
	LcrName         string
	LcrRelationType uint8 `gorm:"comment:master_relation_type"` // master_relation_type
	LcrRefType      uint8 `gorm:"comment:1=>Home,2=>Office"`    // 1=>Home,2=>Office
	LcrMobile       int64 `gorm:"NOT NULL"`
	LcrCreatedBy    uint32
	LcrUdpatedBy    uint32
	LcrActive       uint8 `gorm:"default:1;NOT NULL"`
	LcrDeleted      uint8 `gorm:"default:0;NOT NULL"`
}

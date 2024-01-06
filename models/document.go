package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	LeadID            uint64 `gorm:"NOT NULL"`
	DocPancard        string
	DocMobile         uint64
	DocType           string `gorm:"NOT NULL"`
	SubDocType        string `gorm:"NOT NULL"`
	DocPassword       string
	DocFileName       string
	DocIP             string
	UploadBy          int32
	RemovedBy         int32
	DocsNovelReturnID string
	DocsMasterID      uint8
	DocsAadhaarMasked uint8 `gorm:"default:0;comment:1=>Yes"` // 1=>Yes
	DocsActive        uint8 `gorm:"default:1;NOT NULL"`
	DocsDeleted       uint8 `gorm:"default:0;NOT NULL"`
}

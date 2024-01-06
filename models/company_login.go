package models

import "gorm.io/gorm"

type CompanyLogin struct {
	gorm.Model
	CompanyName    string `gorm:"NOT NULL"`
	CompanyCode    string `gorm:"NOT NULL"`
	CompanyType    string `gorm:"NOT NULL"`
	CompanyUrl     string `gorm:"NOT NULL"`
	CompanyAddress string `gorm:"NOT NULL"`
	CompanyContact string `gorm:"NOT NULL"`
	CreatedBy      uint32 `gorm:"NOT NULL"`
	UpdatedBy      uint32 `gorm:"NOT NULL"`
	Active         uint8  `gorm:"default:1;NOT NULL"`
	Delected       uint8  `gorm:"default:0;NOT NULL"`
}

func (m *CompanyLogin) TableName() string {
	return "company_login"
}

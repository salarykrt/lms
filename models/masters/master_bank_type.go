package masters

import "gorm.io/gorm"

type MasterBankType struct {
	gorm.Model
	MBankTypeName    string `gorm:"NOT NULL"`
	MBankTypeActive  uint8  `gorm:"default:1;NOT NULL"`
	MBankTypeDeleted uint8  `gorm:"default:0;NOT NULL"`
}

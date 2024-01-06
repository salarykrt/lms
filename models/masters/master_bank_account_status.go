package masters

import "gorm.io/gorm"

type MasterBankAccountStatus struct {
	gorm.Model
	BasName    string
	BasActive  uint8 `gorm:"default:1;NOT NULL"`
	BasDeleted uint8 `gorm:"default:0;NOT NULL"`
}

func (m *MasterBankAccountStatus) TableName() string {
	return "master_bank_account_status"
}
